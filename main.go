package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {

	bibleFacts()
	brazilFacts()

}

func brazilFacts() {

	Facts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Não foi possível encontrar o ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		Facts = append(Facts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visitando", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/brazil-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "")
	writeBrazilJSON(Facts)

}

func bibleFacts() {

	Facts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Não foi possível encontrar o ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		Facts = append(Facts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visitando", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/bible-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "")
	writeBibleJSON(Facts)

}

func writeBrazilJSON(data []Fact) {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Impossível criar o arquivo JSON")
		return
	}

	_ = ioutil.WriteFile("brazilfacts.json", file, 0644)

}

func writeBibleJSON(data []Fact) {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Impossível criar o arquivo JSON")
		return
	}

	_ = ioutil.WriteFile("biblefacts.json", file, 0644)

}
