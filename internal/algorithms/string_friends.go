package algorithms

import (
	"strings"
)

// IsStringFriends valid if the two chains are friends
func IsStringFriends(x, y string) bool {
	var builder strings.Builder
	if x == "" || y == "" || len(x) != len(y) {
		return false
	}

	for i := 1; i < len(x); i++ {
		builder.Write([]byte(x[i:]))
		builder.Write([]byte(x[:i]))
		if y == builder.String() {
			return true
		}

		builder.Reset()
	}

	return false
}
