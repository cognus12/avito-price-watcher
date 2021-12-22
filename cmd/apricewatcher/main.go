package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func findPriceNode(content string) string {
	startIndex := strings.Index(content, "<span class=\"js-item-price\"")

	var node string

	node = content[startIndex:]

	closeIndex := strings.Index(node, "</span>")

	node = node[0 : closeIndex+7]

	return node
}

func extractPrice(html string) float64 {
	node := findPriceNode(html)

	fmt.Println(node)

	// parse html

	return 0
}

func requestPrice(url string) {
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

	fmt.Println(price)
}

func main() {
	requestPrice("https://www.avito.ru/vladimir/rasteniya/palma_2294526731")
}
