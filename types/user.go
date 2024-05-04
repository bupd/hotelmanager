package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 13
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() error {
	if len(params.FirstName) < minFirstNameLen {
		return fmt.Errorf(
			"firstName length should be at least %d Characters long.",
			minFirstNameLen,
		)
	}
	if len(params.LastName) < minLastNameLen {
		return fmt.Errorf("lastName length should be at least %d Characters long.", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		return fmt.Errorf("password length should be at least %d Characters long.", minPasswordLen)
	}
  if !isEmailValid(params.Email) {
    return fmt.Errorf("email is not valid")
  }
  return nil
}

func isEmailValid(e string) bool {
  emailRegex := regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)
  return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"     json:"id,omitempty"`
	FirstName         string             `bson:"firstName"         json:"firstName"`
	LastName          string             `bson:"LastName"          json:"LastName"`
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
