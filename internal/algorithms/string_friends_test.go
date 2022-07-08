package algorithms_test

import (
	"testing"

	"github.com/exercise/internal/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestIsStringFriends(t *testing.T) {
	tt := []struct {
		name           string
		valueX         string
		valueY         string
		expectedResult bool
	}{
		{
			name:           "friends true",
			valueX:         "tokyo",
			valueY:         "kyoto",
			expectedResult: true,
		},
		{
			name:           "friends empty",
			valueX:         "",
			valueY:         "kyoto",
			expectedResult: false,
		},
		{
			name:           "no friends",
			valueX:         "amorg",
			valueY:         "kyoto",
			expectedResult: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := algorithms.IsStringFriends(tc.valueX, tc.valueY)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
