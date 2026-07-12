package sqldb

import (
	"database/sql"
)

var (
	sqlDb *sql.DB
)

func SetDb(db *sql.DB) {
	sqlDb = db
}

func GetDb() *sql.DB {
	return sqlDb
}
