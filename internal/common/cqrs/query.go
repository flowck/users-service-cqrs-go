package cqrs

import "context"

type Query[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
