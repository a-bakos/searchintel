package main

import "strings"

type SearchEntry struct {
	Id    int32
	Time  string
	Query string
	Url   string
	Hits  int32
	//user   string
	//email  string
	Target string
	//ip     string
}

func (searchEntry SearchEntry) isValid() bool {

	// todo see if searchEntry.Query contains invalid things

	return isValidLength(searchEntry.Query)
}

func (searchEntry SearchEntry) clean() SearchEntry {

	keyword := strings.TrimSpace(searchEntry.Query)
	keyword = strings.ToLower(keyword)
	keyword = maybeStripReplace(keyword)

	urlSource := urlDecode(searchEntry.Url)
	urlTarget := urlDecode(searchEntry.Target)

	return SearchEntry{
		Id:     searchEntry.Id,
		Time:   searchEntry.Time,
		Query:  keyword,
		Url:    urlSource,
		Target: urlTarget}
}

type EnumSearchSpecific int

const (
	TIME EnumSearchSpecific = iota
	QUERY
	URL
	HITS
	TARGET
)
