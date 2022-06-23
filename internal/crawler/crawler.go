package crawler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type crawler struct {
	ctx    *context.Context
	cancel context.CancelFunc
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
		chromedp.Flag("headless", true), // debug using
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
	}

	//Initialization parameters, first pass an empty data
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	ctx, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))

	//Execute an empty task to create a chrome instance in advance
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	instance = &crawler{ctx: &chromeCtx, cancel: cancel}
}

func Instance() *crawler {
	// if instance == nil {
	// 	lock.Lock()
	// 	defer lock.Unlock()

	// 	if instance == nil {
	// 		initialize()
	// 		return instance
	// 	}
	// }

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
	timeoutCtx, cancel := context.WithTimeout(*c.ctx, 60*time.Second)

	defer cancel()

	err = chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.AttributeValue(selector, attr, &value, &ok),
	)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *crawler) Close() {
	c.cancel()
}
