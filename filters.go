package main

import (
	"net/url"
	"strconv"
	"strings"
)

func stringToInt32(theString string) (int32, error) {
	idInt, err := strconv.ParseInt(theString, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(idInt), nil
}

func getTheId(idAsString string) int32 {
	id, err := stringToInt32(idAsString)
	if err != nil {
		return DEFAULT_MISSING_ID
	}

	return id
}

func urlDecode(encodedUrl string) string {
	decodedUrl, err := url.QueryUnescape(encodedUrl)

	if err != nil {
		return DEFAULT_FAILED_URLDECODE
	}

	return decodedUrl
}

func maybeStripReplace(keyword string) string {
	for toReplace, replaceTo := range REPLACE_MAP {
		if strings.Contains(keyword, toReplace) {
			keyword = strings.Replace(keyword, toReplace, replaceTo, -1)
		}
	}

	return keyword
}

func isValidLength(keyword string) bool {
	keywordLength := len(keyword)
	if keywordLength >= KEYWORD_MIN_LENGTH && keywordLength <= KEYWORD_MAX_LENGTH {
		return true
	}
	return false
}
