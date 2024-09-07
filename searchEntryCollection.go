package main

type SearchEntryCollection []SearchEntry
type RawSearchEntryCollection []SearchEntry
type SkippedEntryCollection []SearchEntry

func (rawCollection RawSearchEntryCollection) cleanEntries() (SearchEntryCollection, SkippedEntryCollection) {
	var cleanCollection SearchEntryCollection
	var skippedCollection SkippedEntryCollection

	for _, searchItem := range rawCollection {

		// check validity of searchItem
		if searchItem.isValid() {
			cleanCollection = append(cleanCollection, searchItem)
		} else {
			// if invalid
			if STORE_INVALID_ITEMS {
				skippedCollection = append(skippedCollection, searchItem)
			}
		}
	}

	return cleanCollection, skippedCollection
}

func (collection SearchEntryCollection) getAllFieldValues(field EnumSearchSpecific) []string {
	var fieldValues []string

	for _, searchItem := range collection {
		var selected string
		switch field {
		case TIME:
			selected = searchItem.Time
		case QUERY:
			selected = searchItem.Query
		case URL:
			selected = searchItem.Url
		// case HITS:
		// 	selected = searchItem.Hits
		case TARGET:
			selected = searchItem.Target
		}

		fieldValues = append(fieldValues, selected)
	}

	return fieldValues
}

// TODO
//OrderByCount,
//OrderByTarget,
//TopKeywords,

type KeywordCountMap map[string]int

func (collection RawSearchEntryCollection) countKeywords() KeywordCountMap {
	topKeywords := make(map[string]int)

	for _, searchItem := range collection {
		topKeywords[searchItem.Query]++
	}

	return topKeywords
}

func (keywordCollection KeywordCountMap) orderByCount() map[int][]string {
	byCount := make(map[int][]string) // count - [kw]

	for keyword, count := range keywordCollection {
		byCount[count] = append(byCount[count], keyword)
	}

	return byCount
}
