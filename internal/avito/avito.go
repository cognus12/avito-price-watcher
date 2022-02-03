package avito

import (
	"apricescrapper/internal/scrapper"
)

func GetPrice(url string) uint64 {
	pageContent := scrapper.GetHTML(url)

	price := extractPrice(pageContent)

	return price
}
