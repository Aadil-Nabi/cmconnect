package cmhttpclient

import (
	"crypto/tls"
	"net/http"
)

// getClient Creates a custom HTTP client with a custom Transport, here we skip the TLS
func GetClient() http.Client {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return client
}
