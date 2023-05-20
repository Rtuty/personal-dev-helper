package db

import (
	"context"
)

type db struct {
	client client.Client
}

func NewStorage(client client.Client) db.Storage {
	return &db{client: client}
}

func (d db) Create(ctx context.Context) error {
	return nil
}

func (d db) Get(ctx context.Context) error {
	return nil
}

func (d db) Update(ctx context.Context) error {
	return nil
}
