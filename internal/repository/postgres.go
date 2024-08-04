package repository

import "github.com/jackc/pgx/v5/pgxpool"

type PostgresRepository struct {
	conn *pgxpool.Pool
}

var PostgresClient *PostgresRepository

func NewPostgresRepository(conn *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		conn: conn,
	}
}
