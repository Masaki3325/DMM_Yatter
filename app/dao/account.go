package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (r *account) FindByID(ctx context.Context, id int) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (r *account) CreateNewAccount(account *object.Account) error {
	tx, _ := r.db.Begin()
	var err error
	defer func() {
		switch r := recover(); {
		case r != nil:
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec(`INSERT INTO account (username,password_hash) VALUES (?,?)`, account.Username, account.PasswordHash); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *account) UpdateAccount(account *object.Account) error {
	tx, _ := r.db.Begin()
	var err error
	defer func() {
		switch r := recover(); {
		case r != nil:
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	query := `UPDATE account SET display_name = ?, avatar = ?, header = ?, note = ? WHERE username = ?`
	if _, err = tx.Exec(query, account.DisplayName, account.Avatar, account.Header, account.Note, account.Username); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *account) FollowAccount(accountid int, followaccountid int) error {
	tx, _ := r.db.Begin()
	var err error
	defer func() {
		switch r := recover(); {
		case r != nil:
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec(`INSERT INTO relationship (follower_id,followee_id) VALUES (?,?)`, accountid, followaccountid); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *account) CheckFollow(ctx context.Context, accountid int, followaccountid int) (bool, error) {
	var result bool
	query := "SELECT EXISTS(SELECT 1 FROM relationship WHERE follower_id = ? AND followee_id = ?)"
	err := r.db.QueryRowContext(ctx, query, followaccountid, accountid).Scan(&result) // クエリのパラメータを逆にする
	if err != nil {
		return false, fmt.Errorf("failed to check follow relationship: %w", err)
	}
	return result, nil
}
