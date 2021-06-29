package main 

import (
	"encoding/json"
	"strconv"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"github.com/gocolly/colly"
)

type Fact struct {
    ID int `json:"id"` 
	Description string `json:"description"`
}

func main() {

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
			ID: factID,
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
	writeJSON(Facts)

}

func writeJSON(data []Fact) {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Impossível criar o arquivo JSON")
		return
	}

	_ = ioutil.WriteFile("brazilfacts.json", file, 0644)

}
