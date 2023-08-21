package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func VerifyAuth(c *gin.Context) {
	tokenFromRequest := c.GetHeader("Authentication")
	if tokenFromRequest == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}
	log.Println(tokenFromRequest)
	ciphertext, err := hex.DecodeString(tokenFromRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}

	authToken := os.Getenv("AUTH_TOKEN")
	secretKey, err := hex.DecodeString(authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(ciphertext) < aes.BlockSize {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	log.Println(string(ciphertext))

	c.Next()
}
