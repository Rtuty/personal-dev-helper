package db

import "context"

type Storage interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) error
	Update(ctx context.Context) error
}
