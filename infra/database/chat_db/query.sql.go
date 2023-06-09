// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package chat_db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCharacter = `-- name: CreateCharacter :one
INSERT INTO characters (name, bio, note)
VALUES ($1, $2, $3) RETURNING id, name, bio, note
`

type CreateCharacterParams struct {
	Name string
	Bio  string
	Note sql.NullString
}

func (q *Queries) CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error) {
	row := q.db.QueryRowContext(ctx, createCharacter, arg.Name, arg.Bio, arg.Note)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Note,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, login_name)
VALUES ($1, $2) RETURNING id, login_name
`

type CreateUserParams struct {
	ID        uuid.UUID
	LoginName string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.LoginName)
	var i User
	err := row.Scan(&i.ID, &i.LoginName)
	return i, err
}

const createUserEmailPassword = `-- name: CreateUserEmailPassword :one
INSERT INTO user_login_with_email_passwords (user_id, email, password_hash)
VALUES ($1, $2, $3) RETURNING id, user_id, email, password_hash
`

type CreateUserEmailPasswordParams struct {
	UserID       uuid.UUID
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUserEmailPassword(ctx context.Context, arg CreateUserEmailPasswordParams) (UserLoginWithEmailPassword, error) {
	row := q.db.QueryRowContext(ctx, createUserEmailPassword, arg.UserID, arg.Email, arg.PasswordHash)
	var i UserLoginWithEmailPassword
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}

const deleteCharacter = `-- name: DeleteCharacter :exec
DELETE
FROM characters
WHERE id = $1
`

func (q *Queries) DeleteCharacter(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCharacter, id)
	return err
}

const findCharacter = `-- name: FindCharacter :one
SELECT id, name, bio, note
FROM characters
WHERE name = $1 LIMIT 1
`

func (q *Queries) FindCharacter(ctx context.Context, name string) (Character, error) {
	row := q.db.QueryRowContext(ctx, findCharacter, name)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Note,
	)
	return i, err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT users.id, users.login_name, user_login_with_email_passwords.email, user_login_with_email_passwords.password_hash
FROM users
         JOIN user_login_with_email_passwords
              ON users.id = user_login_with_email_passwords.user_id
WHERE user_login_with_email_passwords.email = $1
LIMIT 1
`

type FindUserByEmailRow struct {
	ID           uuid.UUID
	LoginName    string
	Email        string
	PasswordHash string
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i FindUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.LoginName,
		&i.Email,
		&i.PasswordHash,
	)
	return i, err
}

const getCharacter = `-- name: GetCharacter :one
SELECT id, name, bio, note
FROM characters
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCharacter(ctx context.Context, id int64) (Character, error) {
	row := q.db.QueryRowContext(ctx, getCharacter, id)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.Note,
	)
	return i, err
}

const listCharacter = `-- name: ListCharacter :many
SELECT id, name, bio, note
FROM characters
ORDER BY name
`

func (q *Queries) ListCharacter(ctx context.Context) ([]Character, error) {
	rows, err := q.db.QueryContext(ctx, listCharacter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.Note,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
