package main

import (
	"fmt"
	xhtml "golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func parsePrice(node string) uint64 {
	reader := strings.NewReader(node)
	z := xhtml.NewTokenizer(reader)

	var price string

	for {
		tt := z.Next()
		switch {
		case tt == xhtml.StartTagToken:
			t := z.Token()
			for _, a := range t.Attr {
				if a.Key == "content" {
					price = a.Val
					break
				}
			}

		}
		break
	}

	parsed, err := strconv.ParseUint(price, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return parsed
}

func extractPrice(html string) uint64 {
	node := findPriceNode(html)

	return parsePrice(node)
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
	requestPrice("https://www.avito.ru/vladimir/kvartiry/2-k._kvartira_522_m_25_et._2285971912")
}
