package otp

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	userName    = "root"
	accountName = "root"
)

func GenerateOTPKey() *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      userName,
		AccountName: accountName,
	})

	if err != nil {
		panic(err)
	}

	return key
}

func GenerateOTPCode(key *otp.Key) string {
	otpCode, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		panic(err)
	}
	return otpCode
}
