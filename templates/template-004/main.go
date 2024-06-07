package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func parseCSV(filePath string) []Record {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	records := make([]Record, 0, len(rows))
	for _, row := range rows {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}

	return records
}

func main() {
	records := parseCSV("table.csv")

	tmpl, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	err = tmpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
