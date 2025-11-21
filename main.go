package main

import (
	"github.com/danny-lee-tech/aca-nevada-scraper/internal/acanevada"
)

var DefaultConfigLocation = "configs/config.yml"

func main() {
	acanevada.RetrievePlans()
}
