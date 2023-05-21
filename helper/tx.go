package helper

import "database/sql"

func CommitOrRolback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRolback := tx.Rollback()
		PanicIfError(errorRolback)
		panic(err)
	} else {
		errorPanic := tx.Commit()
		PanicIfError(errorPanic)
	}
}
