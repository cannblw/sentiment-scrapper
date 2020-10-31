package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const paragraphSeparator = " "

func GetNewsItemContent(url string) (string, error) {
	var newsItemContent string

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("div.StoryBodyCompanionColumn", func(e *colly.HTMLElement) {
		paragraph := e.Text

		newsItemContent += paragraph
		newsItemContent += paragraphSeparator
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(url)

	return newsItemContent, err
}
