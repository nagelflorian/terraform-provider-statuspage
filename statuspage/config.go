package statuspage

import (
	statuspageAPI "github.com/nagelflorian/statuspage-go"
)

// Config is used to configure the creation of a Statuspage client.
type Config struct {
	APIToken string
}

// Meta is used by the provider to access the Statuspage API.
type Meta struct {
	client *statuspageAPI.Client
}

// Client returns a configured Statuspage client.
func (c *Config) Client() (interface{}, error) {
	return &Meta{client: statuspageAPI.NewClient(c.APIToken, nil)}, nil
}
