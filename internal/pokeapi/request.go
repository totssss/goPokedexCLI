package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) httpGet(requestUrl string) ([]byte, error) {
	response, err := c.httpClient.Get(requestUrl)

	if err != nil {
		return []byte{}, err
	}

	responseBody, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	if response.StatusCode > 299 {
		errMsg := fmt.Sprintf("Response failed with status code: %d\nbody:%s\n", response.StatusCode, responseBody)
		return []byte{}, errors.New(errMsg)
	}

	if err != nil {
		return []byte{}, err
	}

	c.cache.Add(requestUrl, responseBody)

	return responseBody, nil
}

func getJSON[Response any](c *Client, requestUrl string) (Response, error) {
	var result Response
	responseBody, cacheHit := c.cache.Get(requestUrl)

	if !cacheHit {
		var err error
		responseBody, err = c.httpGet(requestUrl)
		if err != nil {
			return result, err
		}
	}

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return result, err
	}

	return result, nil
}
