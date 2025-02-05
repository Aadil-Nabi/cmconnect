package controllers

import (
	"crypto/hmac"
	"fmt"
	"net/http"

	"github.com/Aadil-Nabi/cmconnect/configs"
	"github.com/Aadil-Nabi/cmconnect/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// securityPin struct to store the value from the request.
type securityPin struct {
	SecurityPin string
}

func ReadPostHandler(c *gin.Context) {

	var secretPin securityPin
	c.Bind(&secretPin)

	// Load DB and other configurations required.
	DB := configs.ConnectDB()
	// cnfs := configs.MustLoad()

	var identity models.Identity
	// var identityResult identityResult

	result := DB.Where("employee_name=?", "Bz3uZQg=").First(&identity)
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

	fmt.Println(" identity object is : >", identity)

	// verifyMAC(identityDetail.IdentityNumber, cnfs.Encryption_key, identityDetail.Mac)

}

// verifyMAC checks if the provided MAC is valid
func verifyMAC(message, secretKey, receivedMAC string) bool {
	expectedMAC := generateMac(message, secretKey)
	return hmac.Equal([]byte(receivedMAC), []byte(expectedMAC))
}
