package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbName   = "calhounio_demo"
)

func Connect() (*sql.DB, error) {
	// postgres credentials
	host = os.Getenv("DB_HOST")
	// converting port number to integer
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")

	connUri := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connUri)
	if err != nil {
		log.Println(err)
		return db, err
	}

	// close db connection at the end
	// defer db.Close() -> db will be closed in other package

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return db, err
	}

	return db, nil
}
