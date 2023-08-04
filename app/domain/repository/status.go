package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateNewStatus(status *object.Status) error
	FindByID(ctx context.Context, id int) (*object.Status, error)
	LastInserted(ctx context.Context) (int, error)
}
