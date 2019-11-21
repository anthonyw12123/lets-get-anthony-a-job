package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("data/title.basics.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	re := regexp.MustCompile(`^t{2}\d{7}$`)
	for {
		record, err := reader.Read()
		if err != nil || !re.MatchString(record[0]) {
			continue
		}
		fmt.Println(strings.Join(record, "\t"))
		m := map[string]string{"tconst": record[0], "titleType": record[1], "primaryTitle": record[2], "originalTitle": record[3], "isAdult": record[4], "startYear": record[5], "endYear": record[6], "runtimeMinutes": record[7], "genres": record[8]}
		enc := json.NewEncoder(log.Writer())
		enc.Encode(m)
	}
}
