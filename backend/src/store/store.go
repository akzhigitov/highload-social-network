package store

import "database/sql"

var db *sql.DB

func Connect() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(db:3306)/hl")

	if err != nil {
		panic(err.Error())
	}
}

func Disconnect() {
	db.Close()
}

func Migrate() {

}
