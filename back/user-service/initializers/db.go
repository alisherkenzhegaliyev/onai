package initializers

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // posgresql driver
)

var DB *sql.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database is not responding:", err)
	}
}
