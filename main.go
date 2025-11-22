package main

import (
	"log"

	"github.com/danny-lee-tech/aca-nevada-scraper/internal/acanevada"
)

var DefaultConfigLocation = "configs/config.yml"

func main() {
	_, err := acanevada.RetrievePlans()
	if err != nil {
		log.Fatal(err)
	}
}
