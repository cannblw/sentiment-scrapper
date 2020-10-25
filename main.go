package main

import "fmt"

func main() {
	MigrateDatabase()

	err, content := GetNewsItemContent("https://www.nytimes.com/2020/09/26/opinion/sunday/trump-cuomo-new-york-revenge.html")

	if err != nil {
		fmt.Println("Error scrapping news item content")
	}

	fmt.Println(content)
}
