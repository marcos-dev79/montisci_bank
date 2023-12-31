// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64
	Name      string
	Balance   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountTransaction struct {
	ID        int64
	AccountID int64
	Amount    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountUser struct {
	ID        int64
	AccountID int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profession struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID           int64
	Name         string
	Bio          sql.NullString
	Birthdate    time.Time
	Email        string
	LastLogin    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProfessionID int64
}
