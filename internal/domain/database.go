package domain

import "context"

type User struct {
	Id     string
	Fname  string
	Lname  string
	Avatar string
	Email  string
}

type Database interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, fname string, lname string, avatar string, email string) (*User, error)
}
