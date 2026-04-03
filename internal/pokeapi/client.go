package pokeapi

import (
	"net/http"
	"time"

	"github.com/PicusB/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout, cacheExpiry time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(cacheExpiry),
	}
}
