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
		m := map[string]string{"tconst": record[0]}
		enc := json.NewEncoder(log.Writer())
		enc.Encode(m)
	}

}
