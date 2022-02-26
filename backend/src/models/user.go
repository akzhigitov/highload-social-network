package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Birthday  string   `json:"birthday"`
	Sex       byte     `json:"sex"`
	Interests string   `json:"interests,omitempty"`
	Followers []Friend `json:"followers,omitempty"`
	Following []Friend `json:"following,omitempty"`
	City      string   `json:"city"`
	Email     string   `json:"-"`
	Password  []byte   `json:"-"`
	Token     string   `json:"token,omitempty"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
