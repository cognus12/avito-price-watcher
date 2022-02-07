package request

import (
	"log"
	"net/http"
)

func configureRequest(urlStr string) *http.Request {

	req, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		log.Fatalln(err)
	}

	ua := randUserAgent()

	//req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("accept-language", "ru,en;q=0.9,ru-RU;q=0.8,en-US;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", "n\" Not;A Brandn\";v=n\"99n\", n\"Google Chromen\";v=n\"97n\", n\"Chromiumn\";v=n\"97n\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "n\"Windowsn\"")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", ua)

	return req
}
