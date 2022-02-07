package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

func processResponse(response *http.Response) string {
	dataInBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	pageContent := string(dataInBytes)

	return pageContent
}
