package main

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

	return false
}

func (searchEntry SearchEntry) clean() SearchEntry {

	return SearchEntry{
		Id:     searchEntry.Id,
		Time:   searchEntry.Time,
		Query:  searchEntry.Query,
		Url:    searchEntry.Url,
		Target: searchEntry.Target}
}

type EnumSearchSpecific int

const (
	TIME EnumSearchSpecific = iota
	QUERY
	URL
	HITS
	TARGET
)
