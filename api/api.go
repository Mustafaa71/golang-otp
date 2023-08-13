package api

import (
	"example/otpAPI/otp"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	port := os.Getenv("PORT")

	router.GET("/", displayOTPCode)
	router.GET("/env", env)
	router.POST("/verify", verifyOTPCode)

	router.Run("0.0.0.0:" + port)
}

type otpStruct struct {
	OtpCode string `json:"otpCode"`
}

var storedOTP string

func env(c *gin.Context) {
	testENV := os.Getenv("TEST_ENV")
	c.String(http.StatusOK, testENV)
}

func getOTPcode() string {
	key := otp.GenerateOTPKey()
	storedOTP = otp.GenerateOTPCode(key)
	return storedOTP
}

func displayOTPCode(c *gin.Context) {
	otpCode := getOTPcode()
	otpStruct := otpStruct{
		OtpCode: otpCode,
	}
	c.JSON(http.StatusOK, otpStruct)
}

func verifyOTPCode(c *gin.Context) {
	userInput := c.PostForm("otpCode")
	if userInput == storedOTP {
		fmt.Println("Authentication successful!")
		c.String(http.StatusOK, "Authentication successful!")
	} else {
		fmt.Println("Authentication failed. Invalid OTP code.")
		c.String(http.StatusUnauthorized, "Authentication failed. Invalid OTP code.")
	}
}
