package request

import (
	"log"
	"net/http"
	"net/url"
)

func configureClient(proxyStr string) *http.Client {

	var transport *http.Transport

	if proxyStr == "" {
		transport = &http.Transport{}
	} else {
		proxyURL, err := url.Parse(proxyStr)

		if err != nil {
			log.Fatal(err)
		}

		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	client := &http.Client{Transport: transport}

	return client
}
