package models

import (
	"context"
	"database/sql"
)

func newUserDB(db *sql.DB) *userDB {
	return &userDB{db: db}
}

type userDB struct {
	db *sql.DB
}

const createUser = `
INSERT INTO users (
   username,email,birthdate,password_hash,remember_hash
) VALUES (
            $1, $2,$3,$4,$5
        )
   RETURNING id, created_at, updated_at, username, email, birthdate, password_hash, remember_hash
`

// CreateUser inserts a user to the database
func (q *userDB) CreateUser(ctx context.Context, arg User) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Birthdate,
		arg.PasswordHash,
		arg.RememberHash,
	)
	var u User
	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Username,
		&u.Email,
		&u.Birthdate,
		&u.PasswordHash,
		&u.RememberHash,
	)
	return &u, err
}

const deleteUser = `
DELETE FROM users WHERE id = $1
`

// DeleteUser deletes a user from the database
func (q *userDB) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `
SELECT id, created_at, updated_at, username, email, birthdate, password_hash, remember_hash FROM users
WHERE email = $1 LIMIT 1
`

// GetUserByEmail retrives a user from the database by his email
func (q *userDB) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var u User
	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Username,
		&u.Email,
		&u.Birthdate,
		&u.PasswordHash,
		&u.RememberHash,
	)
	return &u, err
}

const getUserByID = `
SELECT id, created_at, updated_at, username, email, birthdate, password_hash, remember_hash FROM users
WHERE id = $1 LIMIT 1
`

// GetUserByID retrives a user from the database by his id
func (q *userDB) GetUserByID(ctx context.Context, id int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var u User
	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Username,
		&u.Email,
		&u.Birthdate,
		&u.PasswordHash,
		&u.RememberHash,
	)
	return &u, err
}

const getUserByRememberHash = `
SELECT id, created_at, updated_at, username, email, birthdate, password_hash, remember_hash FROM users
WHERE remember_hash = $1 LIMIT 1
`

func (q *userDB) GetUserByRemember(ctx context.Context, rememberHash string) (*User, error) {

	row := q.db.QueryRowContext(ctx, getUserByRememberHash, rememberHash)
	var u User
	err := row.Scan(
		&u.ID,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Username,
		&u.Email,
		&u.Birthdate,
		&u.PasswordHash,
		&u.RememberHash,
	)
	return &u, err
}

const updateUserEmail = `
UPDATE users
SET email = $2
WHERE id =  $1
`

// UpdateUserEmail update the user's email
func (q *userDB) UpdateUserEmail(ctx context.Context, id int32, email string) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail, id, email)
	return err
}

const updateUserPassword = `
UPDATE users
SET password_hash = $2
WHERE id =  $1
`

// UpdateUserPassword updates the user's password with the new one
func (q *userDB) UpdateUserPassword(ctx context.Context, id int32, password string) error {

	_, err := q.db.ExecContext(ctx, updateUserPassword, id, password)
	return err
}

const updateUserRememberHash = `
UPDATE users
SET remember_hash = $2
WHERE id =  $1
`

// UpdateUserRemember updates the user's password with the new one
func (q *userDB) UpdateUserRemember(ctx context.Context, id int32, remember string) error {

	_, err := q.db.ExecContext(ctx, updateUserRememberHash, id, remember)
	return err
}
