package models

import (
	"context"
	"database/sql"
	"github.com/Ayoub-Moulahi/MyYouTube/setting"
	"golang.org/x/crypto/bcrypt"
)

// UserDbInterface  used to define the interface for interacting with the user database
type UserDbInterface interface {
	CreateUser(ctx context.Context, arg User) (*User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int32) (*User, error)
	GetUserByRemember(ctx context.Context, remember string) (*User, error)
	UpdateUserEmail(ctx context.Context, id int32, email string) error
	UpdateUserPassword(ctx context.Context, id int32, password string) error
	UpdateUserRemember(ctx context.Context, id int32, remember string) error
}

// UserInterface used to define the interface to interact with user model
type UserInterface interface {
	UserDbInterface
	Authenticate(email, password string) (*User, error)
}
type userStruct struct {
	UserDbInterface
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
	config, err := setting.LoadConfig("../")
	if err != nil {
		return nil, ErrApp
	}
	var ctx = context.Background()
	tmp, err := ui.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, ErrNoAccount
	}
	err = bcrypt.CompareHashAndPassword([]byte(tmp.PasswordHash), []byte(password+config.PasswordPepper))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, ErrPasswordIncorrect
	} else if err != nil {
		return nil, err
	}
	return tmp, nil
}
