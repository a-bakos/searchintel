package main

import (
	"fmt"
)

func main() {
	var searchEntries []SearchEntry = parseCsv(FILE_WITH_PATH)

	fmt.Println(searchEntries)
}
