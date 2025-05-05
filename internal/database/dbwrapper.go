package database

import "database/sql"

type DBWrapper struct {
	*sql.DB
}

func (db *DBWrapper) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DBWrapper) Query(qurry string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(qurry, args...)
}

func (db *DBWrapper) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}