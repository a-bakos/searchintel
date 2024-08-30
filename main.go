package main

import "fmt"

type Enum int

func main() {
	var searchEntryCollection SearchEntryCollection = parseCsv(FILE_WITH_PATH)

	searchEntryCollection.getAllFieldValues(QUERY)

	target := searchEntryCollection.getAllFieldValues(TARGET)
	fmt.Println(target)
}
