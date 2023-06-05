package config

import "os"

type SSLconf struct {
	SSL_CERT        string
	SSL_PRIVATE_KEY string
}

func(ssl *SSLconf) InitSSL() {

	err := os.WriteFile("./ssl/certificate.crt", []byte(ssl.SSL_CERT), 0644)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./ssl/private.key", []byte(ssl.SSL_PRIVATE_KEY), 0644)

	if err != nil {
		panic(err)
	}

}