package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/exercise/cmd/api/handlers"
	"github.com/exercise/internal/pkg/pokemon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	pokemonUrl        = "/pokemon/1"
	pokemonUrlInvalid = "/pokemon/p"
	orderNamesUrl     = "/order/names"
	stringFriendsUrl  = "/string/friends"
)

type clientPokemonMock struct {
	mock.Mock
}

type mockArgs struct {
	methodName string
	inputArgs  []interface{}
	returnArgs []interface{}
	times      int
}

func (c *clientPokemonMock) GetDetail(id int) (pokemon.Pokemon, error) {
	args := c.Called(id)
	return args.Get(0).(pokemon.Pokemon), args.Error(1)
}

func TestGetPokemon(t *testing.T) {
	tt := map[string]struct {
		statusCode       int
		mock             mockArgs
		expectedResponse string
		url              string
	}{
		"Get pokemon ok": {
			statusCode: http.StatusOK,
			mock: mockArgs{
				methodName: "GetDetail",
				inputArgs:  []interface{}{1},
				returnArgs: []interface{}{pokemon.Pokemon{
					Name: "test",
				}, nil},
				times: 1,
			},
			expectedResponse: "test_data/response/pokemon_ok.json",
			url:              pokemonUrl,
		},
		"Get pokemon not found": {
			statusCode: http.StatusNotFound,
			mock: mockArgs{
				methodName: "GetDetail",
				inputArgs:  []interface{}{1},
				returnArgs: []interface{}{pokemon.Pokemon{}, pokemon.ErrNotFound},
				times:      1,
			},
			expectedResponse: "test_data/response/pokemon_not_found.json",
			url:              pokemonUrl,
		},
		"Get pokemon provider error": {
			statusCode: http.StatusInternalServerError,
			mock: mockArgs{
				methodName: "GetDetail",
				inputArgs:  []interface{}{1},
				returnArgs: []interface{}{pokemon.Pokemon{}, pokemon.ErrProvider},
				times:      1,
			},
			expectedResponse: "test_data/response/pokemon_provider_error.json",
			url:              pokemonUrl,
		},
		"Get pokemon invalid identifier": {
			statusCode:       http.StatusUnprocessableEntity,
			expectedResponse: "test_data/response/pokemon_invalid_url.json",
			url:              pokemonUrlInvalid,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := setUpRouter()
			var cmock clientPokemonMock
			setupMock(tc.mock, &cmock.Mock)
			handler := handlers.New(&cmock)
			handler.API(r)

			req, _ := http.NewRequest("GET", tc.url, nil)
			req.Header.Add("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			responseData, _ := ioutil.ReadAll(w.Body)
			if tc.statusCode == http.StatusOK {
				assert.Equal(t, string(readFile(t, tc.expectedResponse)), string(responseData))
			} else {
				assert.JSONEq(t, string(readFile(t, tc.expectedResponse)), string(responseData))
			}
			assert.Equal(t, tc.statusCode, w.Code)
			mock.AssertExpectationsForObjects(t, &cmock)
		})
	}
}

func TestOrderNames(t *testing.T) {
	tt := map[string]struct {
		statusCode       int
		payload          string
		expectedResponse string
		url              string
	}{
		"order name correct": {
			statusCode:       http.StatusOK,
			payload:          "test_data/request/order_name_ok.json",
			expectedResponse: "test_data/response/order_name_ok.json",
			url:              orderNamesUrl,
		},
		"order name error request": {
			statusCode:       http.StatusUnprocessableEntity,
			payload:          "test_data/request/order_name_error.json",
			expectedResponse: "test_data/response/order_name_error.json",
			url:              orderNamesUrl,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := setUpRouter()
			var cmock clientPokemonMock
			handler := handlers.New(&cmock)
			handler.API(r)

			req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer(readFile(t, tc.payload)))
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			responseData, _ := ioutil.ReadAll(w.Body)
			assert.JSONEq(t, string(readFile(t, tc.expectedResponse)), string(responseData))
			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestStringFriend(t *testing.T) {
	tt := map[string]struct {
		statusCode       int
		payload          string
		expectedResponse string
		url              string
	}{
		"String friend ok": {
			statusCode:       http.StatusOK,
			payload:          "test_data/request/string_friends_ok.json",
			expectedResponse: "test_data/response/string_friends_ok.json",
			url:              stringFriendsUrl,
		},
		"String friend error validator": {
			statusCode:       http.StatusUnprocessableEntity,
			payload:          "test_data/request/string_friends_error.json",
			expectedResponse: "test_data/response/string_friends_error.json",
			url:              stringFriendsUrl,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := setUpRouter()
			var cmock clientPokemonMock
			handler := handlers.New(&cmock)
			handler.API(r)

			req, _ := http.NewRequest("POST", tc.url, bytes.NewBuffer(readFile(t, tc.payload)))
			req.Header.Add("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			responseData, _ := ioutil.ReadAll(w.Body)
			if tc.statusCode == http.StatusOK {
				assert.Equal(t, string(readFile(t, tc.expectedResponse)), string(responseData))
			} else {
				assert.JSONEq(t, string(readFile(t, tc.expectedResponse)), string(responseData))
			}
			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func readFile(t *testing.T, fileName string) []byte {
	if fileName == "" {
		return []byte("")
	}

	bytes, err := os.ReadFile(fileName)
	require.NoError(t, err)
	return bytes
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupMock(params mockArgs, mock *mock.Mock) {
	if params.methodName != "" {
		mock.On(params.methodName, params.inputArgs...).
			Return(params.returnArgs...).
			Times(1)
	}
}
