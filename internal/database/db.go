package database

import (
	"context"
	"fmt"
	"log"
	"portfolio/config"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// Getting the Conn Object from teh databse instance
	GetConn() (*pgx.Conn, error)
}

type service struct {
	conn *pgx.Conn
}

// Close implements Service.
func (s *service) Close() error {
	return s.conn.Close(context.Background())
}

// GetConn implements Service.
func (s *service) GetConn() (*pgx.Conn, error) {
	return s.conn, nil
}

func New() Service {
	connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DbHost, config.Envs.DbUsername, config.Envs.DbPassword, config.Envs.DbName)
	return &service{
		conn: connectDb(connStr),
	}
}

// Connects the db with the deadline of 5 seconds
func connectDb(cfg string) *pgx.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return conn
}
