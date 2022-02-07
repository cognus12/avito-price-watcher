package request

import (
	"log"
)

func Get(urlStr string) string {
	client := configureClient("")

	req := configureRequest(urlStr)

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// TODO improve error handling
	// if response.StatusCode == 403 {
	// 	pageContent := processResponse(response)
	// 	fmt.Println(pageContent)
	// 	log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	// }

	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	pageContent := processResponse(response)

	return pageContent
}
