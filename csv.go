package main

import (
	"encoding/csv"
	"os"
)

func parseCsv(fileWithPath string) []SearchEntry {
	var searchEntries []SearchEntry
	file, err := os.Open(fileWithPath)
	if err != nil {
		panic("Error")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error")
	}

	//fmt.Println(records)

	for i, line := range records {
		if i == INDEX_FIRST_ROW { // skip first row
			continue
		}

		searchEntries = append(searchEntries, SearchEntry{
			getTheId(line[0]),
			line[1],
			line[2],
			line[3],
			getTheId(line[4]),
			line[8]})
	}

	return searchEntries
}
