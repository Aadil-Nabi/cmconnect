package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/auth/jwtauth"
	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/internal/pkg/cmhttpclient"
	"github.com/Aadil-Nabi/cmconnect/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// Create a struct accessible from outside of this package
type IdentityDetails struct {
	IdentityNumber string
	SecurityPin    string
	Department     string
}

var identityPayload IdentityDetails

func CreatePostHandler(c *gin.Context) {

	// bind the requested input to the required struct, in short store the requested value in the variable.
	c.Bind(&identityPayload)

	fmt.Println("identityPayload is:", identityPayload)

	// Load configurations from the yaml file.
	cnfg := configs.MustLoad()

	// generate the mac value for the identity number.
	mac := generateMac(identityPayload.SecurityPin, cnfg.Encryption_key)

	fmt.Println("SecurityPinNumber", identityPayload.SecurityPin)
	fmt.Println("IdentityNUmber", identityPayload.IdentityNumber)
	// encrypt the identity number and insert in two columns (identity_number and cipher)
	cipherText := encrypting()
	fmt.Println("ciphertext is :: ", cipherText)
	cipher := cipherText["ciphertext"]

	fmt.Println("cipher is ::", cipher)

	identity := models.Identity{
		IdentityNumber: cipher,
		SecurityPin:    mac,
		Department:     identityPayload.Department,
	}

	// get DB connection and create an entry inside the table.
	DB := configs.ConnectDB()

	result := DB.Where("security_pin_number=?", mac).First(&identity)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			DB.Create(&identity)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "identity already exist in the database, try changin Security Pin Number",
		})
		// fmt.Println("entry already exisist ", identity)
		return
	}
	// Send JSON response to the front end.
	c.JSON(http.StatusOK, gin.H{
		"ciphertext": cipher,
	})

}

// encrypting method to encrypt the data using the provided key in the config.yaml file
func encrypting() map[string]string {
	cnfg := configs.MustLoad()

	identityNumber := identityPayload.IdentityNumber
	fmt.Println("identityNumber is :", identityNumber)

	// Get Jwt details like token type and actual token to create a bearer string
	jwt_details := jwtauth.GetAuthDetails()
	bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	url := cnfg.Base_Url + cnfg.Version + "/crypto/encrypt"

	// Encode the data to be encrypted in base64 string as CM only accepts a valid base64 string
	plaintext := identityNumber
	fmt.Println("plaintext is :", plaintext)
	plaintext = base64.StdEncoding.EncodeToString([]byte(plaintext))
	payload := map[string]string{
		"id":        cnfg.Encryption_key,
		"plaintext": plaintext,
	}

	// Convert data into JSON encoded byte array
	encodedBody, _ := json.Marshal(payload)

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	body := bytes.NewBuffer(encodedBody)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalf("Something went wrong in the request  %v", err)
	}

	// Add the required headers to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", bearer)

	//get client from a helper function
	client := cmhttpclient.GetClient()

	// Do method to send the http request to the CM to http response
	// this is used when we add headers to the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to Encrypt %v", err)
	}

	// close the response
	defer resp.Body.Close()

	// Read the response received from the CM
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output map[string]string

	yaml.Unmarshal(data, &output)

	fmt.Println("data is :", string(data))
	fmt.Println("output is: ", output)
	return output

}

// generateMac creates a MAC using HMAC-SHA256 of the passed message using the key.
func generateMac(message, key string) string {
	secretKey := []byte(key)
	hash := hmac.New(sha256.New, secretKey)
	hash.Write([]byte(message))
	mac := hash.Sum(nil)
	return hex.EncodeToString(mac)

}
