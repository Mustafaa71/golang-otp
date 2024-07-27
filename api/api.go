package api

import (
	"example/otpAPI/otp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	router.GET("/", displayOTPCode)
	router.POST("/verify", verifyOTPCode)

	router.Run()
}

type otpStruct struct {
	OtpCode string `json:"otpCode"`
}

var storedOTP string

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
	userInput := c.Query("otpCode")
	if userInput == storedOTP {
		c.String(http.StatusOK, "Authentication successful!")
	} else {

		c.String(http.StatusUnauthorized, "Authentication failed. Invalid OTP code.")
	}

}
