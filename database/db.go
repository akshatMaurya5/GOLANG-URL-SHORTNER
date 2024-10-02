package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv" // Import the godotenv package
)

var DB *sql.DB

func ConnectDB() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	dbServer := os.Getenv("DB_SERVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbEncrypt := os.Getenv("DB_ENCRYPT")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=%s",
		dbServer,
		dbUser,
		dbPassword,
		dbPort,
		dbName,
		dbEncrypt)

	DB, err = sql.Open("mssql", connString)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("Connected to the database successfully!")
	return nil
}
