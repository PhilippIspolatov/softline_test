package tools

import "errors"

var (
	AlreadyExist = errors.New("already exist")
	ErrorCreatingUser = errors.New("error creating user")

	HttpBadRequest = errors.New("bad request")
	HttpConflict = errors.New("user already exist")
)