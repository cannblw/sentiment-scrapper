package main

import (
	"edgarchirivella.com/sentiment-scrapper/entity"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := "https://www.nytimes.com/2020/09/26/opinion/sunday/trump-cuomo-new-york-revenge.html"

	err, content := GetNewsItemContent(url)
	if err != nil {
		log.Fatal("Error scrapping news item content")
	}

	newsItem := entity.NewsItem{
		Url:     content,
		Content: "Lorem Ipsum",
	}

	err = SaveNewsItem(newsItem)
	if err != nil {
		log.Fatal("Error saving news item to DynamoDB")
	}
}
