package server

import "context"

type Server interface {
	Start()
	Stop(ctx context.Context)
}
