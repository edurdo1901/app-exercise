package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = time.Second * 10
	defaultHost    = "https://pokeapi.co/"
)

var (
	ErrNotFound = errors.New("client: pokemon not found")
	ErrProvider = errors.New("client: error consuming pokemon api")
)

// Option modifies values within the client
type Option func(client *http.Client)

type Client struct {
	host   string
	Detail *http.Client
}

// New create pokemon client
func New(options ...Option) *Client {
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	for _, op := range options {
		op(client)
	}

	return &Client{
		host:   defaultHost,
		Detail: client,
	}
}

// GetDetail get the pokemon detail
func (c *Client) GetDetail(id int) (Pokemon, error) {
	url, err := url.Parse(fmt.Sprintf("%s%s", c.host, fmt.Sprintf("api/v2/pokemon-form/%d", id)))
	if err != nil {
		return Pokemon{}, err
	}

	request := &http.Request{
		Method: "GET",
		URL:    url,
	}

	res, err := c.Detail.Do(request)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	if err = handlerError(res); err != nil {
		return Pokemon{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err = json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

// handlerError validates if api response is greater than 400 and generates an error
func handlerError(res *http.Response) error {
	if res.StatusCode >= http.StatusBadRequest {
		if res.StatusCode == http.StatusNotFound {
			return ErrNotFound
		}

		return ErrProvider
	}

	return nil
}
