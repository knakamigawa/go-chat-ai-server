// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package chat_db

import (
	"database/sql"
)

type Character struct {
	ID   int64
	Name string
	Bio  string
	Note sql.NullString
}