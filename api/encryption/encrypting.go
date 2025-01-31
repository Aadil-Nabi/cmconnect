package encryption

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/auth/jwtauth"
	"github.com/Aadil-Nabi/cmconnect/internal/config"
	"github.com/Aadil-Nabi/cmconnect/internal/pkg/cmhttpclient"
)

// Encrypting method to encrypt the data using the provided key
func Encrypting() {

	jwt_details := jwtauth.GetAuthDetails()
	bearer := jwt_details.Token_type + " " + jwt_details.Jwt

	configs := config.MustLoad()
	url := configs.Base_Url + configs.Version + "/crypto/encrypt"

	// Encode the data to be encrypted in base64 string as CM only accepts a valid base64 string
	plaintext := "Asdf@1234"
	plaintext = base64.StdEncoding.EncodeToString([]byte(plaintext))
	payload := map[string]string{
		"id":        configs.Encryption_key,
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

	fmt.Println(string(data))

}
