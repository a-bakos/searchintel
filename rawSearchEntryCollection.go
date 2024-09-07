package main

type RawSearchEntryCollection []SearchEntry

func (rawCollection RawSearchEntryCollection) cleanEntries() (SearchEntryCollection, SkippedEntryCollection) {
	var cleanCollection SearchEntryCollection
	var skippedCollection SkippedEntryCollection

	for _, searchItem := range rawCollection {

		searchItem = searchItem.clean()

		if searchItem.isValid() {
			cleanCollection = append(cleanCollection, searchItem)

			// TODO if searchItem.isSearchDoi // handle DOI search
			// This will probably live elsewhere
			// DOI's start with 10.xxxx, where x = a digit
			// (?i)10.\d{4}

		} else {
			if STORE_INVALID_ITEMS {
				skippedCollection = append(skippedCollection, searchItem)
			}
		}
	}

	return cleanCollection, skippedCollection
}

func (collection RawSearchEntryCollection) countKeywords() KeywordCountMap {
	topKeywords := make(map[string]int)

	for _, searchItem := range collection {
		topKeywords[searchItem.Query]++
	}

	return topKeywords
}
