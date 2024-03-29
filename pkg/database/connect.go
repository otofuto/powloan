package database

import (
	//"log"
	"os"
	"strings"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	_ = godotenv.Load()

	connectionstring := os.Getenv("DATABASE_URL")
	if connectionstring == "" {
		connectionstring = "host=127.0.0.1 port=5432 user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	}
	db, err := sql.Open("postgres", connectionstring)
	if err != nil {
		panic(err.Error())
	}
	//log.Println("connect db");

	return db
}

func Escape(str string) string {
	ret := strings.Replace(str, "\\", "\\\\", -1)
	ret = strings.Replace(ret, "\"", "\\\"", -1)
	ret = strings.Replace(ret, "'", "\\'", -1)
	ret = strings.Replace(ret, "\t", "\\t", -1)
	ret = strings.Replace(ret, "\r", "\\r", -1)
	ret = strings.Replace(ret, "\n", "\\n", -1)

	return ret
}