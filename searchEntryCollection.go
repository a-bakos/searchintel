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

func (collection SearchEntryCollection) getTopKeywords() map[string]int {
	topKeywords := make(map[string]int)

	for _, searchItem := range collection {
		topKeywords[searchItem.Query]++
	}

	byCount := make(map[int][]string)
	// sort map by value
	for key, count := range topKeywords {
		// todo
	}

	return topKeywords
}
