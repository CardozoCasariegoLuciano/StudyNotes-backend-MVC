package errortypes

import "errors"

var (
	MailAlreadyTaken = errors.New("Email already taken")
	InternalError    = errors.New("Server internal error")
	WrongPassOrEmail = errors.New("Wrong email or password")
	UserNotFound     = errors.New("User not found")
)
