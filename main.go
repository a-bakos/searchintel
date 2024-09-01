package main

import "fmt"

func main() {
	var searchEntryCollection SearchEntryCollection = parseCsv(RAW_INPUT_FILE_WITH_PATH)

	searchEntryCollection.getAllFieldValues(QUERY)

	target := searchEntryCollection.getAllFieldValues(TARGET)
	fmt.Println(target)
}
