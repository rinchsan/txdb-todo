package transaction

import "database/sql"

// See also: https://stackoverflow.com/questions/16184238/database-sql-tx-detecting-commit-or-rollback
func Run(db *sql.DB, f func(tx *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return f(tx)
}
