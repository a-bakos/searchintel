package main

import "fmt"

func main() {
	var rawCollection RawSearchEntryCollection = parseCsv(RAW_INPUT_FILE_WITH_PATH)

	var cleanCollection SearchEntryCollection
	var skippedCollection SkippedEntryCollection
	cleanCollection, skippedCollection = rawCollection.cleanEntries()

	fmt.Println(cleanCollection)

	if STORE_INVALID_ITEMS {
		fmt.Println(skippedCollection)
	}

	//rawSearchEntryCollection.getAllFieldValues(QUERY)

	//target := searchEntryCollection.getAllFieldValues(TARGET)
	//fmt.Println(target)

	//topkws := searchEntryCollection.countKeywords()
	//fmt.Println(topkws)

	//topkws.orderByCount()
}
