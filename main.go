package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitializeDynamodb()

	err, content := GetNewsItemContent("https://www.nytimes.com/2020/09/26/opinion/sunday/trump-cuomo-new-york-revenge.html")

	if err != nil {
		fmt.Println("Error scrapping news item content")
	}

	fmt.Println(content)
}
