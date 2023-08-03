package object

import (
	"time"
)

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	Account *Account `json:"account,omitempty"`

	// The associated account ID
	AccountID int64 `json:"account_id,omitempty" db:"account_id"`

	// The content of the status
	Content string `json:"content" db:"content"`

	// The time the status was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
