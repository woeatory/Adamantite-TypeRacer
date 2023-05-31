package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo() *Repo {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	repoDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	err = repoDB.Ping()
	if err != nil {
		panic(err)
	}
	query, err := os.ReadFile("sql/create_tables.sql")
	if err != nil {
		panic(err)
	}
	rows, err := repoDB.Query(string(query))
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	return &Repo{
		DB: repoDB,
	}
}

func (repo *Repo) CloseRepo() error {
	err := repo.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
