package main

import (
	"net/url"
	"strconv"
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
		return ""
		// fmt.Println(encodedUrl)
		// panic(err)
	}

	return decodedUrl
}
