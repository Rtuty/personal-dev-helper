package db

import (
	"context"
	"dev-helper/pkg/client"
	"dev-helper/pkg/logging"
)

type db struct {
	client client.Client
	logger *logging.Logger
}

func NewStorage(client client.Client, logger *logging.Logger) Storage {
	return &db{client: client, logger: logger}
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
