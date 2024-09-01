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

type EnumSearchSpecific int

const (
	TIME EnumSearchSpecific = iota
	QUERY
	URL
	HITS
	TARGET
)
