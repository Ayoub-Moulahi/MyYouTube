package models

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	CreateUser(ctx context.Context, arg User) (*User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int32) (*User, error)
	GetUserByRemember(ctx context.Context, remember string) (*User, error)
	UpdateUserEmail(ctx context.Context, id int32, email string) error
	UpdateUserPassword(ctx context.Context, id int32, password string) error
	UpdateUserRemember(ctx context.Context, id int32, remember string) error
}
type userStruct struct {
	UserInterface
}

func NewUserInterface(db *sql.DB) UserInterface {
	udb := &userDB{
		db: db,
	}
	uv := newUserValidator(udb)
	return &userStruct{
		uv,
	}
}

func (ui *userStruct) Authenticate(email, password string) (*User, error) {
	var ctx = context.Background()
	tmp, err := ui.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(tmp.PasswordHash), []byte(password+pwd_pepper))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, ErrPasswordIncorrect
	} else if err != nil {
		return nil, err
	}
	return tmp, nil
}
