package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		driver   = os.Getenv("DB_DRIVER")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	driverURL := driver + "://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	db, err := sql.Open("postgres", driverURL)
	if err != nil {
		log.Println(err)
	}

	m, err := migrate.New("file://migrations", driverURL)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	return &Repo{
		DB: db,
	}
}

func (repo *Repo) CloseRepo() error {
	err := repo.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
