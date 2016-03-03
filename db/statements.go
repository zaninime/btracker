package db

import "database/sql"

const (
	stmtCheckConnectionID = ``
)

var (
	StmtCheckConnectionID *sql.Stmt
)

func prepareStatements() error {
	var err error
	StmtCheckConnectionID, err = DB.Prepare(stmtCheckConnectionID)
	if err != nil {
		return err
	}
	return nil
}
