package repository

import (
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateNewStatus(status *object.Status) error
}
