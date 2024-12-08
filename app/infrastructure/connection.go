package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/kajiLabTeam/walking-information-storage-server-api/app/config"
)

func Connect() (*sql.DB, error) {
	env := config.NewPostgresEnv()

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetHostOfPrivateValue(),
		env.GetPortOfPrivateValue(),
		env.GetUserOfPrivateValue(),
		env.GetPasswordOfPrivateValue(),
		env.GetDatabaseOfPrivateValue(),
	)

	fmt.Println(connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
