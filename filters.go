package main

import (
	"fmt"
	"net/url"
	"strconv"
)

// Parse string as int32
func stringToInt32(theString string) int32 {
	idInt, err := strconv.ParseInt(theString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	return int32(idInt)
}

func urlDecode(encodedUrl string) string {
	decodedUrl, err := url.QueryUnescape(encodedUrl)

	if err != nil {
		panic(err)
	}

	return decodedUrl
}
