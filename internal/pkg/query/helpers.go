package query

import (
	"regexp"
	"strings"
)

func toSnakeCase(str string) string {
	if str == "-" || str == "" {
		return ""
	}

	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// trimJSONTag takes an json tag as a string and trim omitempty if exist.
func trimJSONTag(tag string) string {
	if strings.Contains(tag, "omitempty") {
		return strings.Replace(tag, ",omitempty", "", 1)
	}

	return tag
}

func shift(src []sort, elem []sort) []sort {
	if src == nil {
		return elem
	}

	return append(elem, src...)
}

func validSortType(sortType string) bool {
	return strings.ToLower(sortType) == ascending || strings.ToLower(sortType) == descending
}

// questionMarkSequence takes an integer and return a string containing
// n times questionMark  ',' separated.
func questionMarkSequence(n int) string {
	var res string

	for i := 0; i < n; i++ {
		res += "?,"
	}

	return strings.TrimRight(res, ",")
}
