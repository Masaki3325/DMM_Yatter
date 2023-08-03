package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	timeline struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) GetStatuses(ctx context.Context, maxid string, sinceid string, limit string) ([]object.Status, error) {
	query := "SELECT * FROM status WHERE 1=1" // Basic query
	var args []interface{}                    // Arguments for the query

	// If maxid is provided, add it to the query
	if maxid != "" {
		maxIDValue, err := strconv.Atoi(maxid)
		if err != nil {
			return nil, fmt.Errorf("invalid maxid value: %w", err)
		}
		query += " AND id <= ?"
		args = append(args, maxIDValue)
	}

	// If sinceid is provided, add it to the query
	if sinceid != "" {
		sinceIDValue, err := strconv.Atoi(sinceid)
		if err != nil {
			return nil, fmt.Errorf("invalid sinceid value: %w", err)
		}
		query += " AND id > ?"
		args = append(args, sinceIDValue)
	}

	// Add ordering to make sure the results are consistent
	query += " ORDER BY id"

	// If limit is provided, add it to the query
	if limit != "" {
		limitValue, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("invalid limit value: %w", err)
		}
		query += " LIMIT ?"
		args = append(args, limitValue)
	}

	// Now execute the query
	var entity []object.Status
	err := r.db.SelectContext(ctx, &entity, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}
