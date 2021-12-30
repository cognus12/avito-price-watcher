package avito

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetPrice(url string) uint64 {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	dataInBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	pageContent := string(dataInBytes)

	price := extractPrice(pageContent)

	return price
}
