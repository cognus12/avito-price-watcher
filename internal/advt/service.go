package advt

import (
	"apricescrapper/internal/apperror"
	"apricescrapper/internal/crawler"
	"fmt"
	"strconv"
	"strings"
)

type AdInfo struct {
	Url   string `json:"url"`
	Price uint64 `json:"price"`
}

type service struct {
	crawler crawler.Crawler
}

type Service interface {
	GetAdInfo(url string) (AdInfo, error)
}

func NewService(crawler crawler.Crawler) Service {
	return &service{
		crawler: crawler,
	}
}

func (s *service) GetAdInfo(url string) (AdInfo, error) {
	result := AdInfo{}

	validUrl, err := s.sanitizeUrl(url)

	if err != nil {
		return result, err
	}

	priceStr, err := s.crawler.GetAttribute(validUrl, ".js-item-price", "content")

	if err != nil {
		return result, err
	}

	price, err := strconv.ParseUint(priceStr, 10, 64)

	if err != nil {
		return result, err
	}

	result.Price = price
	result.Url = url

	return result, nil
}

func (s *service) sanitizeUrl(url string) (string, error) {
	if url == "" {
		badRequest := apperror.BadRequest
		badRequest.Message = "No URL provided"

		return url, badRequest
	}

	if strings.HasPrefix(url, "avito.ru") || strings.HasPrefix(url, "www.avito.ru") {
		return fmt.Sprintf("https://%v", url), nil
	}

	if !strings.Contains(url, "avito.ru") {
		badRequest := apperror.BadRequest
		badRequest.Message = "Incorrect URL. URL should be to avito.ru"

		return url, badRequest
	}

	return url, nil
}
