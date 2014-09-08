package service

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresService struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Db       string `json:"db"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func (pg *PostgresService) State() state {
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", pg.Host, pg.Port, pg.User, pg.Password, pg.Db)
	db, err := sql.Open("postgres", uri)
	defer db.Close()
	if err != nil {
		return DOWN
	}
	_, err = db.Query("select now();")
	if err != nil {
		return DOWN
	}
	return ALIVE
}

func (pg *PostgresService) Name() string {
	return "postgres"
}

func (pg *PostgresService) String() string {
	str := fmt.Sprintf("%s %s:%d", pg.Name(), pg.Host, pg.Port)
	return str
}
