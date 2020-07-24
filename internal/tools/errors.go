package tools

import "errors"

var (
	AlreadyExist = errors.New("already exist")
	ErrorCreatingUser = errors.New("error creating user")
	UserDoesNotExist = errors.New("user does not exist")
	WrongPassword = errors.New("wrong password")

	HttpBadRequest = errors.New("bad request")
	HttpConflict = errors.New("user already exist")
	HttpAlreadyAuthenticate = errors.New("user already authenticated")
)