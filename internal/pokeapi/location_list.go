package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) RequestLocations(url string) (LocationResponse, error) {
	data, ok := c.cache.Get(url)
	if !ok {
		var err error
		data, err = c.httpGet(url)
		if err != nil {
			return LocationResponse{}, err
		}
	}

	locationResponse := LocationResponse{}
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationResponse{}, err
	}

	return locationResponse, nil
}

func (c *Client) httpGet(url string) ([]byte, error) {
	res, err := c.httpClient.Get(url)

	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		errMsg := fmt.Sprintf("Response failed with status code: %d\nbody:%s\n", res.StatusCode, body)
		return []byte{}, errors.New(errMsg)
	}

	if err != nil {
		return []byte{}, err
	}

	c.cache.Add(url, body)

	return body, nil
}
