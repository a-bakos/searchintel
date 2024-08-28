package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type SearchEntry struct {
	Id    int32
	Time  string
	Query string
	Url   string
	Hits  int32
	//user   string
	//email  string
	Target string
	//ip     string
}

func main() {
	//testSearch := SearchEntry{1, "2024/03/23 19:06:04", "cats", "http://example.com/home", 2, "http://example.com/kittens"}

	var searchEntries []SearchEntry

	file, err := os.Open("testsheet.csv")
	if err != nil {
		panic("Error")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error")
	}

	fmt.Println(records)

	for i, line := range records {
		if i == 0 { // skip first row
			continue
		}
		if err != nil {
			panic("Error")
		}

		// Parse string as int32
		convert := func(theString string) int32 {
			idInt, err := strconv.ParseInt(theString, 10, 64)
			if err != nil {
				fmt.Println(err)
			}

			return int32(idInt)
		}

		searchEntries = append(searchEntries, SearchEntry{
			convert(line[0]),
			line[1],
			line[2],
			line[3],
			convert(line[4]),
			line[7]})
	}

	fmt.Println(searchEntries)
}
