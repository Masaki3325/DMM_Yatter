package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	FindByID(ctx context.Context, id int) (*object.Account, error)
	CheckFollow(ctx context.Context, int, followaccountid int) (bool, error)
	// TODO: Add Other APIs
	CreateNewAccount(account *object.Account) error
	UpdateAccount(account *object.Account) error
	FollowAccount(accountid int, followaccountid int) error
}
