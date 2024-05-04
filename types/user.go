package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 13
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p UpdateUserParams) ToBson() bson.M {
	m := bson.M{}

	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}

	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params UpdateUserParams) Validate() error {
	if len(params.FirstName) > 0 {
		if len(params.FirstName) < minFirstNameLen {
			return fmt.Errorf(
				"firstName length should be at least %d Characters long.",
				minFirstNameLen,
			)
		}
	}
	if len(params.LastName) > 0 {
		if len(params.LastName) < minLastNameLen {
			return fmt.Errorf(
				"lastName length should be at least %d Characters long.",
				minLastNameLen,
			)
		}
	}
	return nil
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf(
			"firstName length should be at least %d Characters long.",
			minFirstNameLen,
		)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf(
			"lastName length should be at least %d Characters long.",
			minLastNameLen,
		)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf(
			"password length should be at least %d Characters long.",
			minPasswordLen,
		)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is not valid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)
	return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"     json:"id,omitempty"`
	FirstName         string             `bson:"firstName"         json:"firstName"`
	LastName          string             `bson:"lastName"          json:"lastName"`
	Email             string             `bson:"email"             json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
