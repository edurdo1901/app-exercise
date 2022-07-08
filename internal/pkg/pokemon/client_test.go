package pokemon_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/exercise/internal/pkg/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const urlDefault = "/api/v2/pokemon-form/1"

func TestGetDetailPokemon(t *testing.T) {

	tt := []struct {
		name         string
		url          string
		status       int
		expectedName string
		payload      string
		err          error
	}{
		{
			name:         "successfull",
			url:          urlDefault,
			status:       200,
			expectedName: "bulbasaur",
			payload:      "testdata/ok.json",
		},
		{
			name:         "no found",
			url:          urlDefault,
			status:       404,
			expectedName: "",
			payload:      "",
			err:          pokemon.ErrNotFound,
		},
		{
			name:         "provider error",
			url:          urlDefault,
			status:       500,
			expectedName: "",
			payload:      "testdata/invalid.json",
			err:          pokemon.ErrProvider,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client := pokemon.New(setUpClient(t,
				getLoadFile(t, tc.payload),
				tc.url,
				tc.status))

			response, err := client.GetDetail(1)
			if tc.err != nil {
				assert.Error(t, tc.err, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedName, response.Name)
		})
	}
}

func getLoadFile(t *testing.T, fileName string) []byte {
	if fileName != "" {
		data, err := ioutil.ReadFile(fileName)
		require.NoError(t, err)
		require.NotEmpty(t, data)
		return data
	}

	return []byte("")
}

func setUpClient(t *testing.T, payload []byte, url string, statusCode int) pokemon.Option {
	return func(client *http.Client) {
		client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
			assert.Equal(t, req.URL.Path, url)
			return &http.Response{
				StatusCode: statusCode,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBuffer(payload)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})
	}
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
