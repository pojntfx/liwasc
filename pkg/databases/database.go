package databases

import "database/sql"

type SQLiteDatabase struct {
	dbPath string
	db     *sql.DB
}

func (d *SQLiteDatabase) Open() error {
	db, err := sql.Open("sqlite3", d.dbPath)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1) // Prevent "database locked" errors

	d.db = db

	return nil
}
