package subscription

import (
	"apricescrapper/internal/apperror"
	"apricescrapper/internal/crawler"
	"fmt"
	"strconv"
	"strings"
)

type service struct {
	crawler crawler.Crawler
	store   Repository
}

type Service interface {
	GetAdInfo(url string) (adInfo, error)
	Subscribe(url string, email string) error
	Unsubscribe(url string, email string) error
}

type urlParams struct {
	city     string
	category string
	slug     string
}

type adInfo struct {
	Url   string `json:"url"`
	Price uint64 `json:"price"`
}

func NewService(crawler crawler.Crawler, repository Repository) Service {
	return &service{
		crawler: crawler,
		store:   repository,
	}
}

func (s *service) Subscribe(url, email string) error {
	err := s.store.CreateSubscibtion(url, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Unsubscribe(url, email string) error {
	err := s.store.DeleteSubscibtion(url, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAdInfo(url string) (adInfo, error) {
	result := adInfo{}

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
