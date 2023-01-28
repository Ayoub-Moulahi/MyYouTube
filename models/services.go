package models

import (
	"database/sql"
)

type Services struct {
	UserInter UserInterface
}

func NewService(dialect, connInfo string) (*Services, error) {
	db, err := sql.Open(dialect, connInfo)
	if err != nil {
		return nil, err
	}
	return &Services{
		UserInter: NewUserInterface(db),
	}, nil

}
