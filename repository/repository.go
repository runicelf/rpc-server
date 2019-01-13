package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/runicelf/rpc-server/models"
	"github.com/satori/go.uuid"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func New(config models.Config) Repository {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open(config.DriverName, connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return Repository{DB: db}
}

func (r Repository) Add(login string) (string, error) {
	sqlStatement := "INSERT INTO users.users (uuid, login, date) VALUES ($1, $2, $3) RETURNING uuid"

	generatedUUID := uuid.Must(uuid.NewV4())

	uuidFromDB := ""
	err := r.DB.QueryRow(sqlStatement, generatedUUID, login, time.Now()).Scan(&uuidFromDB)
	if err != nil {
		return "", err
	}

	return uuidFromDB, nil
}

func (r Repository) Get(uuid string) (user models.DBModelUser, err error) {
	sqlStatement := "SELECT * FROM users.users WHERE uuid = $1;"
	err = r.DB.QueryRow(sqlStatement, uuid).Scan(&user.UUID, &user.Login, &user.Date)
	if err != nil {
		return models.DBModelUser{}, err
	}

	return user, nil
}

func (r Repository) Update(user models.RequestModelUser) error {
	sqlStatement := "UPDATE users.users SET login = $1 WHERE uuid = $2;"
	result, err := r.DB.Exec(sqlStatement, user.Login, user.UUID)
	if err != nil {
		return err
	}

	counter, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if counter == 0 {
		return errors.New("value with such uuid not found")
	}

	return nil
}
