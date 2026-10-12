package pokeapi

import (
	"net/http"
	"time"

	"gitlab.insa-rouen.fr/tsavary/goPokedexCLI/internal/pokecache"
)

const (
	BaseLocationURL = "https://pokeapi.co/api/v2/location-area"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
