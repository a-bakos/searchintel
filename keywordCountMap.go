package main

type KeywordCountMap map[string]int

func (keywordCollection KeywordCountMap) orderByCount() map[int][]string {
	byCount := make(map[int][]string) // count - [kw]

	for keyword, count := range keywordCollection {
		byCount[count] = append(byCount[count], keyword)
	}

	return byCount
}
