package main

import (
	"encoding/csv"
	"fmt"
	"os"
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
}
