package repository

import "database/sql"

const (
	driverName = "postgres"
)

type Database struct {
	url string
	Db  *sql.DB
}

func NewDatabase(url string) *Database {
	return &Database{
		url: url,
	}
}

func (d *Database) Connect() error {
	database, err := sql.Open(driverName, d.url)

	if err != nil {
		return err
	}

	d.Db = database
	return nil
}

func (d *Database) Close() error {
	return d.Db.Close()
}
