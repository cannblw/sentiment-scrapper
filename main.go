package main

import (
	"edgarchirivella.com/sentiment-scrapper/entity"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatal(err)
	}

	InitDynamoDb()

	url := "https://www.nytimes.com/2020/09/26/opinion/sunday/trump-cuomo-new-york-revenge.html"

	err, content := GetNewsItemContent(url)
	if err != nil {
		fmt.Println("Error scrapping news item content")
		log.Fatal(err)
	}

	newsItem := entity.NewsItem{
		Url:     content,
		Content: "Lorem Ipsum",
	}

	err = SaveNewsItem(newsItem)
	if err != nil {
		fmt.Println("Error saving news item to DynamoDB")
		log.Fatal(err)
	}
}
