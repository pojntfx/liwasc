package databases

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

type SQLiteDatabase struct {
	DBPath     string
	Migrations migrate.MigrationSource

	db *sql.DB
}

func (d *SQLiteDatabase) Open() error {
	// Open the DB
	db, err := sql.Open("sqlite3", d.DBPath)
	if err != nil {
		return err
	}

	// Configure the db
	db.SetMaxOpenConns(1) // Prevent "database locked" errors
	d.db = db

	// Run migrations if set
	if d.Migrations != nil {
		if _, err := migrate.Exec(d.db, "sqlite3", d.Migrations, migrate.Up); err != nil {
			return err
		}
	}

	return nil
}
