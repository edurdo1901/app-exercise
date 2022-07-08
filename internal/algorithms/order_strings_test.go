package algorithms_test

import (
	"testing"

	"github.com/exercise/internal/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestGetNames(t *testing.T) {
	tt := []struct {
		name          string
		values        string
		namesExpected []string
	}{
		{
			name:   "four names",
			values: "Luis,Camilo,Andres,Laura",
			namesExpected: []string{
				"Andres",
				"Camilo",
				"Laura",
				"Luis",
			},
		},
		{
			name:          "empty name",
			values:        "",
			namesExpected: []string{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			resultNames, countNames := algorithms.Order(tc.values)
			assert.Equal(t, len(tc.namesExpected), countNames)
			assert.Equal(t, tc.namesExpected, resultNames)
		})
	}
}
