package helper

import (
	"crypto/rand"
	"io"
)

type OtpGenerator interface {
	GetOtp() (string, error)
}

type otpGenerator struct {}

func NewOtpGenerator() *otpGenerator {
	return &otpGenerator{}
}

func (o *otpGenerator) GetOtp() (string, error) {
	code := make([]byte, OTP_SIZE)

	n, err := io.ReadAtLeast(rand.Reader, code, OTP_SIZE)
	if n != OTP_SIZE || err != nil {
		return "", err
	}

	for i := 0; i < len(code); i++ {
		// iterate each code byte val after randomizing,
		// then change the value by modulo the val by len of num list
		// then we get random index of table value
		code[i] = table[int(code[i])%len(table)]
	}

	return string(code), nil
}

const OTP_SIZE = 4

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
