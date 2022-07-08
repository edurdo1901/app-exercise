package algorithms

import (
	"sort"
	"strings"
)

// Order separates the names by , and organizes them by alphabetical name and returns the number of elements and the organized names
func Order(value string) ([]string, int) {
	if value == "" {
		return []string{}, 0
	}

	values := strings.Split(value, ",")
	sort.Strings(values)
	return values, len(values)
}
