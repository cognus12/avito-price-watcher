package crawler

import (
	"context"

	"github.com/chromedp/chromedp"
)

type crawler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type Crawler interface {
	GetAttribute(url string, selector string, attr string) (string, error)
	Stop()
}

func NewCrawler() Crawler {
	ctx, cancel := chromedp.NewContext(context.Background())
	return &crawler{ctx: ctx, cancel: cancel}
}

func (c *crawler) GetAttribute(url string, selector string, attr string) (string, error) {
	var value string

	var err error

	var ok bool

	err = chromedp.Run(c.ctx,
		chromedp.Navigate(url),
		chromedp.AttributeValue(selector, attr, &value, &ok),
	)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *crawler) Stop() {
	c.cancel()
}
