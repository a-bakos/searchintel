package main

type SearchEntryCollection []SearchEntry

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

func (collection SearchEntryCollection) countKeywords() KeywordCountMap {
	topKeywords := make(map[string]int)

	for _, searchItem := range collection {
		topKeywords[searchItem.Query]++
	}

	return topKeywords
}

func (keywordCollection KeywordCountMap) orderByCount() map[int][]string {
	byCount := make(map[int][]string) // count - [kw]

	for key, count := range keywordCollection {
		byCount[count] = append(byCount[count], key)
	}

	return byCount
}
