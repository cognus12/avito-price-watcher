package crawler

import (
	"context"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type crawler struct {
	ctx    context.Context
	cancel context.CancelFunc
	b      chromedp.Browser
}

type Crawler interface {
	GetAttribute(url string, selector string, attr string) (string, error)
	Close()
}

var lock = &sync.Mutex{}

var once sync.Once

var instance *crawler

func initialize() {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.DisableGPU,
		chromedp.Headless,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
	}

	//Initialization parameters, first pass an empty data
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(ctx)

	chromedp.Run(chromeCtx)

	instance = &crawler{ctx: chromeCtx, cancel: cancel, b: *chromedp.FromContext(chromeCtx).Browser}
}

func Instance() *crawler {
	once.Do(func() {
		initialize()
	})

	return instance
}

func (c *crawler) GetAttribute(url string, selector string, attr string) (string, error) {
	var value string

	var err error

	var ok bool

	//Create a context with a timeout of 60s
	ctx, cancel := context.WithTimeout(c.ctx, 60*time.Second)

	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)

	defer cancel()

	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.AttributeValue(selector, attr, &value, &ok),
	)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *crawler) Close() {
	c.b.Process().Kill()
	c.cancel()
}
