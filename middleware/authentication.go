package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

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

	decryptedString := (string(ciphertext))
	validToken := validateToken(decryptedString)
	if !validToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
		return
	}
	c.Next()
}

func validateToken(tokenString string) bool {
	var numbers, text strings.Builder

	for _, char := range tokenString {
		if unicode.IsDigit(char) {
			numbers.WriteRune(char)
		} else {
			text.WriteRune(char)
		}
	}

	separatedNumbers := numbers.String()
	separatedText := text.String()

	layout := "20060102150405"
	parsedTime, err := time.Parse(layout, separatedNumbers)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	currentTime := time.Now()
	timeDifference := currentTime.Sub(parsedTime)

	fiveMinutes := 5 * time.Minute
	if timeDifference > fiveMinutes || separatedText != os.Getenv("AUTH_USER") {
		log.Println("The parsed time ftom the token is greater than 5 minutes from the current time.")
		return false
	}
	return true
}
