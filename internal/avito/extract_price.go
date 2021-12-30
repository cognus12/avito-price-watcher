package avito

import (
	xhtml "golang.org/x/net/html"
	"log"
	"strconv"
	"strings"
)

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
