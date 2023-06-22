package helper

import (
	"errors"
	"net/url"
)

func IsValidUrl(rawURL string, validUrlList map[string]bool) error {
	url, err := url.Parse(rawURL)
	// log.Println(url)
	if err != nil {
		return ErrInvalidUrl
	}

	host := url.Hostname()
	isValidHost := validUrlList[host]
	if url.Scheme != "https" {
		return ErrInvalidUrl
	}

	if !isValidHost {
		return ErrInvalidUrlHost
	}

	return nil
}

// Error Message
var (
	ErrInvalidUrl     = errors.New("error invalid url")
	ErrInvalidUrlHost   = errors.New("invalid url host, use telegram or zoom link")
)
