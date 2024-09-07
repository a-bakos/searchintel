package main

type EnumSearchSpecific int

const (
	TIME EnumSearchSpecific = iota
	QUERY
	URL
	HITS
	TARGET
)
