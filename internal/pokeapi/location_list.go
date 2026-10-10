package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func RequestLocations(url string) (LocationResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return LocationResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		errMsg := fmt.Sprintf("Response failed with status code: %d\nbody:%s\n", res.StatusCode, body)
		return LocationResponse{}, errors.New(errMsg)
	}

	if err != nil {
		return LocationResponse{}, err
	}

	locationResponse := LocationResponse{}

	errUnmarshal := json.Unmarshal(body, &locationResponse)
	if errUnmarshal != nil {
		return LocationResponse{}, err
	}

	return locationResponse, nil
}
