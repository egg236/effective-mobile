package db

import (
	"database/sql"
	"effective-mobile/config"
	"effective-mobile/entities"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB interface {
	ReadRecordByID(id int) (*entities.Record, error)
	ReadRecords(userID, serviceName string) ([]*entities.Record, error)

	CreateRecord(record *entities.Record) (int, error)

	UpdateRecord(record *entities.Record) error

	DeleteRecord(id int) error
}

type database struct {
	*sql.DB
}

func NewDB() (DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.Cfg.DBUser(),
		config.Cfg.DBPass(),
		config.Cfg.DBHost(),
		config.Cfg.DBPort(),
		config.Cfg.DBName(),
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &database{
		conn,
	}, nil
}
