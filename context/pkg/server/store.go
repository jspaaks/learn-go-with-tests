package server

import "context"

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
