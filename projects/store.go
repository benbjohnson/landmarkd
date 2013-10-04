package projects

import (
	"fmt"
	"errors"
    "github.com/skydb/sky.go"
	"net/url"
)

type Store interface {
	Open() error
	Close()
	FindByApiKey(string) (*Project, error)
}

// Creates new store based on a connection URI.
func NewStore(client sky.Client, uri string) (Store, error) {
	if client == nil {
		return nil, errors.New("Sky client required for project store")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "redis":
		return NewRedisStore(client, u), nil
	default:
		return nil, fmt.Errorf("Invalid store scheme: %s", u.Scheme)
	}
}
