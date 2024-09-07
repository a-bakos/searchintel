package main

import "fmt"

func main() {
	var rawCollection RawSearchEntryCollection = parseCsv(RAW_INPUT_FILE_WITH_PATH)

	var cleanCollection SearchEntryCollection
	var skippedCollection SkippedEntryCollection
	cleanCollection, skippedCollection = rawCollection.cleanEntries()

	if STORE_INVALID_ITEMS {
		fmt.Println(" SKIPPED ITEMS ")
		fmt.Println(skippedCollection)
	}

	fmt.Println(" CLEAN COLLECTION ")
	fmt.Println(cleanCollection)

	//rawSearchEntryCollection.getAllFieldValues(QUERY)

	//target := searchEntryCollection.getAllFieldValues(TARGET)
	//fmt.Println(target)

	//topkws := searchEntryCollection.countKeywords()
	//fmt.Println(topkws)

	//topkws.orderByCount()
}
