package models

type publicError string

func (p publicError) Error() string {
	return string(p)
}

type internalError string

func (i internalError) Error() string {
	return string(i)
}

const (
	ErrPasswordShort     publicError = "ERROR:Password should contain at least 8 characters"
	ErrPasswordRequired  publicError = "ERROR:Password required please fill in the password field"
	ErrPasswordMatch     publicError = "ERROR: Password must be a combination of lower and upper case letter and contain at least one special character ie @,/ { ... "
	ErrPasswordTooLong   publicError = "ERROR:Password length should be at max 32 characters"
	ErrPasswordIncorrect publicError = "ERROR:Password incorrect"

	ErrEmailRequired publicError   = "ERROR: Email required ,please make sure to provide your email"
	ErrInvalidEmail  publicError   = "ERROR:Email invalid please provide a valid email "
	ErrEmailTaken    publicError   = "ERROR: this email is already in use ,please try another email"
	ErrNoAccount     publicError   = "ERROR:you don't have an account "
	ErrTokenNotSet   publicError   = "ERROR: Remember not set"
	ErrIdInvalid     publicError   = "ERROR: Id can't be negative"
	ErrApp           internalError = "Something went wrong,please try again later.If the problem persist please contact us "
)
