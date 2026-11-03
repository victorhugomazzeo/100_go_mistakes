// Package redis is a fake client: just enough to collide with a variable
// named redis in main.go.
package redis

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(key string) string {
	return "bar"
}
