package avito

import (
	"apricewatcher/internal/request"
)

func GetPrice(url string) uint64 {
	pageContent := request.Get(url)

	price := extractPrice(pageContent)

	return price
}
