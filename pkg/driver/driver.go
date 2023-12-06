package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type DB struct {
	SQL *sql.DB
}

var DBConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDblifeTime = 5 * time.Minute

func ConnectDB() (*DB, error) {
	//Load the enviornment Variables to use them
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	//Open a connection pool
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	//testing whether the databasea is connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//setting constants for the database connections
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetConnMaxLifetime(maxDblifeTime)
	//assigning the DB struct to the database
	DBConn.SQL = db
	return DBConn, nil
}
