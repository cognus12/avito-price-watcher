package scrapper

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func GetHTML(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string

	err := chromedp.Run(ctx, chromedp.Navigate(url), chromedp.OuterHTML("html", &html, chromedp.ByQuery))

	if err != nil {
		log.Fatal(err)
	}

	return html
}
