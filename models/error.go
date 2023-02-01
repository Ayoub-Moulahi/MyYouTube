package models

type publicError string

func (p publicError) Error() string {
	return string(p)
}

func (p publicError) ToDisplay() string {
	return string(p)
}

const (
	ErrPasswordShort     publicError = "ERROR:Password should contain at least 8 characters"
	ErrPasswordRequired  publicError = "ERROR:Password required please fill in the password field"
	ErrPasswordMatch     publicError = "ERROR: Password must be a combination of lower and upper case letter and contain at least one digit number and one special character ie @,/ { ... "
	ErrPasswordTooLong   publicError = "ERROR:Password length should be at max 32 characters"
	ErrPasswordIncorrect publicError = "ERROR:Password incorrect"

	ErrEmailRequired publicError = "ERROR: Email required ,please make sure to provide your email"
	ErrInvalidEmail  publicError = "ERROR:Email invalid please provide a valid email "
	ErrEmailTaken    publicError = "ERROR: this email is already in use ,please try another email"
	ErrNoAccount     publicError = "ERROR:you don't have an account "
	ErrApp           publicError = "INTERNAL ERROR:Something went wrong,please try again later.If the problem persist please contact us "
)
