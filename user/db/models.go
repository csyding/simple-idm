// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type SchemaMigration struct {
	Version int64 `json:"version"`
	Dirty   bool  `json:"dirty"`
}

type User struct {
	Uuid           uuid.UUID      `json:"uuid"`
	CreatedAt      time.Time      `json:"created_at"`
	LastModifiedAt time.Time      `json:"last_modified_at"`
	DeletedAt      sql.NullTime   `json:"deleted_at"`
	CreatedBy      sql.NullString `json:"created_by"`
	Email          string         `json:"email"`
	Name           sql.NullString `json:"name"`
}
