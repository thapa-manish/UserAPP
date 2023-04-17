package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"
)

func RunMigration(db *sql.DB) {
	file, err := ioutil.ReadFile("./db/user.sql")
	if err != nil {
		log.Println("cant read sql file:", err)
		return
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			log.Println("cant execute the query: ", err)
		}
	}
}
