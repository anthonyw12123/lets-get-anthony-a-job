package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	file, err := os.Open("title.basics.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	for {
		record, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		m := map[string]string{"tconst": record[0], "titleType": record[1], "primaryTitle": record[2], "originalTitle": record[3], "isAdult": record[4], "startYear": record[5], "endYear": record[6], "runtimeMinutes": record[7], "genres": record[8]}
		enc := json.NewEncoder(log.Writer())
		enc.Encode(m)
	}
}
