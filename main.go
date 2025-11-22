package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danny-lee-tech/aca-nevada-scraper/internal/acanevada"
)

var DefaultConfigLocation = "configs/config.yml"

func main() {
	planList, err := acanevada.RetrievePlans()
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.Create("plans.csv")
	if err != nil {
		log.Fatal("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprint(file, acanevada.PrintPlanCSVHeader())
	for _, plan := range planList {
		fmt.Fprint(file, plan.PrintPlanCSVRow())
	}
}
