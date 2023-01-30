package supabase

import (
	"fmt"

	"github.com/carlmjohnson/requests"
)

type Client struct {
	url string
	key string

	accessToken  string
	refreshToken string
}

func NewClient(url string, key string) *Client {
	return &Client{
		url:         url,
		key:         key,
		accessToken: key,
	}
}

// Create a new requests.Builder with the API URL, the API key and the auth token if available
func (c *Client) api() *requests.Builder {
	builder := requests.URL(c.url).Header("apikey", c.key)
	if c.accessToken != "" {
		builder.Header("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}
	return builder
}
