package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://rss.nytimes.com/services/xml/rss/nyt/sunday-review.xml")

	for _, item := range feed.Items {
		fmt.Println(item.Link)
		fmt.Println("----")
	}
}
