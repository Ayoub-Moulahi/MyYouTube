package models

import (
	"context"
	"database/sql"
	"github.com/Ayoub-Moulahi/MyYouTube/setting"
	"github.com/Ayoub-Moulahi/MyYouTube/token"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
	"unicode"
)

type validationFunction func(user *User) error

func runValidationFucnction(u *User, functions ...validationFunction) error {
	for _, function := range functions {
		err := function(u)
		if err != nil {
			return err
		}

	}
	return nil
}

type userValidator struct {
	ui          UserDbInterface
	emailRegExp *regexp.Regexp
}

func newUserValidator(ui UserDbInterface) *userValidator {
	return &userValidator{
		ui:          ui,
		emailRegExp: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}

}

//implementing validation

func (uv *userValidator) CreateUser(ctx context.Context, arg User) (*User, error) {

	err := runValidationFucnction(&arg, uv.requireEmail, uv.normalizeEmail, uv.checkValidEmail, uv.checkAvailableEmail, uv.requirePwd, uv.checkPwdLen, uv.checkPasswordMatch, uv.hashPassword, uv.pwdHashRequired, uv.setRemember, uv.hashRemember)
	if err != nil {
		return nil, err
	}
	return uv.ui.CreateUser(ctx, arg)

}

func (uv *userValidator) DeleteUser(ctx context.Context, id int32) error {
	if id < 0 {
		return ErrIdInvalid
	}
	return uv.ui.DeleteUser(ctx, id)
}

func (uv *userValidator) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{Email: email}
	err := runValidationFucnction(&u, uv.requireEmail, uv.normalizeEmail, uv.checkValidEmail)
	if err != nil {
		return nil, err
	}
	return uv.ui.GetUserByEmail(ctx, email)
}

func (uv *userValidator) GetUserByID(ctx context.Context, id int32) (*User, error) {
	if id < 0 {
		return nil, ErrIdInvalid
	}
	return uv.ui.GetUserByID(ctx, id)
}
func (uv *userValidator) GetUserByRemember(ctx context.Context, remember string) (*User, error) {
	u := User{Remember: remember}
	err := runValidationFucnction(&u, uv.hashRemember)
	if err != nil {
		return nil, err
	}
	return uv.ui.GetUserByRemember(ctx, u.RememberHash)
}
func (uv *userValidator) UpdateUserEmail(ctx context.Context, id int32, email string) error {
	if id < 0 {
		return ErrIdInvalid
	}
	u := User{Email: email}
	err := runValidationFucnction(&u, uv.requireEmail, uv.normalizeEmail, uv.checkValidEmail, uv.checkAvailableEmail)
	if err != nil {
		return err
	}
	return uv.ui.UpdateUserEmail(ctx, id, email)
}
func (uv *userValidator) UpdateUserPassword(ctx context.Context, id int32, password string) error {
	if id < 0 {
		return ErrIdInvalid
	}
	u := User{Password: password}
	err := runValidationFucnction(&u, uv.requirePwd, uv.checkPwdLen, uv.checkPasswordMatch, uv.hashPassword, uv.pwdHashRequired)
	if err != nil {
		return nil
	}
	return uv.ui.UpdateUserPassword(ctx, id, u.PasswordHash)

}

func (uv *userValidator) UpdateUserRemember(ctx context.Context, id int32, remember string) error {
	if id < 0 {
		return ErrIdInvalid
	}
	u := User{Remember: remember}
	err := runValidationFucnction(&u, uv.setRemember, uv.hashRemember)
	if err != nil {
		return err
	}
	return uv.ui.UpdateUserRemember(ctx, id, u.RememberHash)

}

// validating and normalizing password
func (uv *userValidator) requirePwd(u *User) error {
	if u.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func (uv *userValidator) checkPwdLen(u *User) error {
	if len(u.Password) < 8 {
		return ErrPasswordShort
	} else if len(u.Password) > 32 {
		return ErrPasswordTooLong
	}
	return nil
}

func (uv *userValidator) checkPasswordMatch(u *User) error {

	x, y, z := verifyPassword(u.Password)
	if x == true && y == true && z == true {
		return nil
	}
	return ErrPasswordMatch
}
func verifyPassword(s string) (number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
		}
	}

	return
}

func (uv *userValidator) hashPassword(u *User) error {
	config, err := setting.LoadConfig("../")
	if err != nil {
		return ErrApp
	}

	if u.Password == "" {
		return nil
	}
	tmp, err := bcrypt.GenerateFromPassword([]byte(u.Password+config.PasswordPepper), bcrypt.DefaultCost)
	if err != nil {
		return ErrApp
	}
	u.PasswordHash = string(tmp)
	u.Password = ""
	return nil
}

func (uv *userValidator) pwdHashRequired(u *User) error {
	if u.PasswordHash == "" {
		return ErrApp
	}
	return nil
}

// validating and normalizing email
func (uv *userValidator) requireEmail(u *User) error {
	if u.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (uv *userValidator) normalizeEmail(u *User) error {
	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)
	return nil
}

func (uv *userValidator) checkValidEmail(u *User) error {
	if uv.emailRegExp.MatchString(u.Email) == false {
		return ErrInvalidEmail
	}
	return nil
}
func (uv *userValidator) checkAvailableEmail(u *User) error {
	var ctx = context.Background()
	tmp, err := uv.ui.GetUserByEmail(ctx, u.Email)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	if tmp.ID != u.ID {
		return ErrEmailTaken
	}
	return nil
}

// other validation :
func (uv *userValidator) setRemember(u *User) error {
	if u.Remember != "" {
		return nil
	}

	newToken, err := token.GenerateToken(token.RememberTokenBytes)
	if err != nil {
		return ErrTokenNotSet
	}
	u.Remember = newToken
	return nil

}

func (uv *userValidator) hashRemember(u *User) error {
	config, err := setting.LoadConfig("../")
	if err != nil {
		return ErrApp
	}
	if u.Remember == "" {
		return nil
	}
	u.RememberHash = token.HashToken(u.Remember, config.HashKey)
	return nil
}
