package scrapper

import "apricescrapper/internal/web"

func GetPrice(url string) uint64 {
	pageContent := web.GetHTML(url)

	price := extractPrice(pageContent)

	return price
}
