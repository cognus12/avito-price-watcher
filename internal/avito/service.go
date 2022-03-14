package avito

import (
	"apricescrapper/internal/crawler"
	"errors"
	"strconv"
	"strings"
)

type service struct {
	crawler crawler.Crawler
}

type Service interface {
	GetAdInfo(args urlParams) (adInfo, error)
}

type urlParams struct {
	city     string
	category string
	slug     string
}

type adInfo struct {
	City     string `json:"city"`
	Category string `json:"catagory"`
	Price    uint64 `json:"price"`
}

func NewService(crawler crawler.Crawler) Service {
	return &service{
		crawler: crawler,
	}
}

func (s *service) GetAdInfo(args urlParams) (adInfo, error) {
	const baseUrl = "https://www.avito.ru/"

	failResult := adInfo{}

	errorMap := make(map[string]string)

	if args.city == "" {
		errorMap["city"] = "no city provided"
	}

	if args.category == "" {
		errorMap["category"] = "No category provided"
	}

	if args.slug == "" {
		errorMap["slug"] = "No slug provided"
	}

	if len(errorMap) > 0 {
		errSlice := []string{}

		for _, value := range errorMap {
			errSlice = append(errSlice, value)
		}

		return failResult, errors.New(strings.Join(errSlice, ", "))
	}

	url := baseUrl + args.city + "/" + args.category + "/" + args.slug

	priceStr, err := s.crawler.GetAttribute(url, ".js-item-price", "content")

	if err != nil {
		return failResult, err
	}

	price, err := strconv.ParseUint(priceStr, 10, 64)

	if err != nil {
		return failResult, err
	}

	successResul := adInfo{
		City:     args.city,
		Category: args.category,
		Price:    price,
	}

	return successResul, nil
}
