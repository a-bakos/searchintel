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
