package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/needl3/goreact-template/internal/domain"
)

func (p *PostgresRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := p.conn.Query(ctx, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error fetching from database")
	}
	for user.Next() {
		var id string
		var email string
		var fname string
		var lname string
		var avatar string
		err = user.Scan(&id, &fname, &lname, &avatar, &email, nil, nil)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Error marshalling row info from db")
		}
		return &domain.User{Id: id, Email: email, Fname: fname, Lname: lname, Avatar: avatar}, nil
	}
	return nil, errors.New("User not found")
}

func (p *PostgresRepository) CreateUser(ctx context.Context, fname string, lname string, avatar string, email string) (*domain.User, error) {
	_, err := p.conn.Exec(ctx, "INSERT INTO users (fname, lname, avatar, email) VALUES ($1, $2, $3, $4)", fname, lname, avatar, email)
	if err != nil {
		return nil, err
	}
	return &domain.User{Id: "", Email: email, Fname: fname, Lname: lname, Avatar: avatar}, nil
}
