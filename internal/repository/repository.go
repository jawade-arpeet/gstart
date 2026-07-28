package repository

import "gstart/internal/client"

type Repository struct{}

func New(client *client.Client) *Repository {
	return &Repository{}
}
