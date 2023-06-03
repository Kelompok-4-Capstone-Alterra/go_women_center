package helper

import (
	"crypto/rand"
	"io"
)

type OtpGenerator interface {
	GetOtp() (string, error)
}

type otpGenerator struct {
	OTP_SIZE int
	table []byte
}

func NewOtpGenerator(size int, table []byte) *otpGenerator {
	return &otpGenerator{
		OTP_SIZE: size,
		table: table,
	}
}

func (o *otpGenerator) GetOtp() (string, error) {
	code := make([]byte, o.OTP_SIZE)

	n, err := io.ReadAtLeast(rand.Reader, code, o.OTP_SIZE)
	if n != o.OTP_SIZE {
		return "", err
	}

	for i := 0; i < len(code); i++ {
		// iterate each code byte val after randomizing,
		// then change the value by modulo the val by len of num list
		// then we get random index of table value
		code[i] = o.table[int(code[i])%len(o.table)]
	}

	return string(code), nil
}
