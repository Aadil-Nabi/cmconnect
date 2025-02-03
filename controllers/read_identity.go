package controllers

import (
	"crypto/hmac"
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadPostHandler(c *gin.Context) {

	// Load DB and other configurations required.
	DB := configs.ConnectDB()
	// cnfs := configs.MustLoad()

	var identity IdentityDetails

	result := DB.Where("cipher=?", "lrSNa13I").First(&identity)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// DB.Create(&identity)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "identitynumber not found",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": identity,
		})

	}

	// verifyMAC(identityDetail.IdentityNumber, cnfs.Encryption_key, identityDetail.Mac)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": "decryptyed data",
	// })

}

// verifyMAC checks if the provided MAC is valid
func verifyMAC(message, secretKey, receivedMAC string) bool {
	expectedMAC := generateMac(message, secretKey)
	return hmac.Equal([]byte(receivedMAC), []byte(expectedMAC))
}

func decrypting() {

	// jwt_details := jwtauth.GetAuthDetails()

	// bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	// configs := configs.MustLoad()
	// base_url := configs.Base_Url
	// version := configs.Version

	// url := base_url + version + "/crypto/decrypt"

	// // Encode the data to be encrypted in base64 string as CM only accepts a valid base64 string
	// cipherText := identityNumber
	// plaintext = base64.StdEncoding.EncodeToString([]byte(plaintext))
	// payload := map[string]string{
	// 	"id":        configs.Encryption_key,
	// 	"plaintext": plaintext,
	// }

	// // Convert data into JSON encoded byte array
	// encodedBody, _ := json.Marshal(payload)

	// // convert the encoded JSON data to a type implemented by the io.Reader interface
	// body := bytes.NewBuffer(encodedBody)

	// req, err := http.NewRequest("POST", url, body)
	// if err != nil {
	// 	log.Fatalf("Something went wrong in the request  %v", err)
	// }

	// // Add the required headers to the request
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Accept", "application/json")
	// req.Header.Set("Authorization", bearer)

	// //get client from a helper function
	// client := cmhttpclient.GetClient()

	// // Do method to send the http request to the CM to http response
	// // this is used when we add headers to the request
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Unable to Encrypt %v", err)
	// }

	// // close the response
	// defer resp.Body.Close()

	// // Read the response received from the CM
	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var output map[string]string

	// yaml.Unmarshal(data, &output)

	// return output

}
