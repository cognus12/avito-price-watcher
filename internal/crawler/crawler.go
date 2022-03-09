package crawler

import (
	"context"
	"strconv"

	"github.com/chromedp/chromedp"
)

type crawler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type Crawler interface {
	GetPrice(url string) (uint64, error)
	Stop()
}

func NewCrawler() Crawler {
	ctx, cancel := chromedp.NewContext(context.Background())
	return &crawler{ctx: ctx, cancel: cancel}
}

func (c *crawler) GetPrice(url string) (uint64, error) {

	var text string
	var res uint64
	var err error

	var ok bool

	err = chromedp.Run(c.ctx,
		chromedp.Navigate(url),
		chromedp.AttributeValue(`.js-item-price`, "content", &text, &ok),
	)

	if err != nil {
		return 0, err
	}

	res, err = strconv.ParseUint(text, 10, 64)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (c *crawler) Stop() {
	c.cancel()
}
