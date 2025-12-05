// Package scraper used to scrape
package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

func FetchRecipes() (string, error) {
	var data string
	domain := "satisfactory.wiki.gg"
	url := "https://satisfactory.wiki.gg/wiki/Template:DocsRecipes.json"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3.1 Safari/605.1.15"

	collector := colly.NewCollector(colly.UserAgent(userAgent))

	collector.OnRequest(func(r *colly.Request) {
		log.Printf("Request for: %s", domain)
	})
	collector.OnResponse(func(r *colly.Response) {
		log.Printf("Response from: %s", domain)
	})
	collector.OnError(func(r *colly.Response, e error) {
		log.Printf("Request for: %s failed with error - %v", domain, e)
	})
	collector.OnHTML("pre", func(f *colly.HTMLElement) {
		data = f.Text
	})

	err := collector.Visit(url)
	if err != nil {
		return "", err
	} else {
		return data, nil
	}
}
