package server

import (
	"context"
	"net/url"
)

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

// TODO: Check if this is needed
// Endpointer is registry endpoint.
type Endpointer interface {
	Endpoint() (*url.URL, error)
}
