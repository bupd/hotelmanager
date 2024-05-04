package db

import "github.com/bupd/hotelmanager/types"

type UserStore interface {
	GetUserById(string) (*types.User, error)
}

