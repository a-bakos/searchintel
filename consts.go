package main

const RAW_INPUT_FILE_WITH_PATH string = "testsheet.csv"
const OUTPUT_FILE_WITH_PATH string = ""

const INDEX_FIRST_ROW int = 0

const KEYWORD_MIN_LENGTH int = 3
const KEYWORD_MAX_LENGTH int = 100

const STORE_INVALID_ITEMS bool = true

const DEFAULT_MISSING_ID = 0
const DEFAULT_FAILED_URLDECODE = "FAILED_DECODE"

// Maps can't be consts because of dynamic assignment
// Case-sensitive keys - Use lowercase!
var REPLACE_MAP = map[string]string{
	"&amp;":  "&",
	"&test;": "[test]",
}
