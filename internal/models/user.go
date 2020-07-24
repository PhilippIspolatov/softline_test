package models

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Nickname string `json:"nickname" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Phone string `json:"phone" validate:"required,phone"`
}

func NewUser(nickname string, email string, password string, phone string) (*User, error) {
	v := validator.New()
	_ = v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		if fl.Field().String() == "" {
			return false
		}

		regExp := `^(\+7|7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$`

		match, err := regexp.Match(regExp, []byte(phone))
		if err != nil || !match {
			return false
		}

		return true
	})
	
	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		if fl.Field().String() == "" {
			return false
		}

		regExp := `^[a-zA-Z]\w{3,14}$`

		match, err := regexp.Match(regExp, []byte(password))

		if err != nil || !match {
			return false
		}

		return true
	})

	user := &User{
		Nickname: nickname,
		Email:    email,
		Password: password,
		Phone:    phone,
	}

	err := v.Struct(user)

	if err != nil {
		return nil, errors.New("bad params")
	}

	return user, nil
}

func (u *User) PasswordHash() {
	hash := sha512.New()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func (u *User) CheckPassword(password string) bool {
	hash := sha512.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)) == u.Password
}
