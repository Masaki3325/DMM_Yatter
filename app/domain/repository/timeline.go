package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch account which has specified username
	GetStatuses(ctx context.Context, maxid string, sinceid string, limit string) ([]object.Status, error)
}
