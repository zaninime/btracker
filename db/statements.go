package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type preparedStmt struct {
	String string
	Stmt   *sql.Stmt
}

type preparedStmtx struct {
	String string
	Stmt   *sqlx.Stmt
}

// Prepared statements used a lot in the application
var (
	StmtWriteConnectionID = preparedStmt{
		`INSERT INTO "public"."connection" ("id", "ip", "expiry") VALUES (decode($1, 'base64'), $2, $3)`,
		nil,
	}
	stmtCheckConnectionID = preparedStmt{
		`SELECT "ip" FROM "public"."connection" WHERE id = decode($1, 'base64') AND ip = $2 AND expiry >= $3`,
		nil,
	}
	stmtGetPeers = preparedStmt{
		`SELECT "ip", "port" FROM "public"."peer" WHERE id != decode($1, 'base64') AND torrent_id = decode($2, 'base64') AND "state" > 0 LIMIT $3`,
		nil,
	}
	stmtGetLeecherPeers = preparedStmt{
		`SELECT COUNT(*) FROM "public"."peer" WHERE torrent_id = decode($1, 'base64') AND "left" > 0`,
		nil,
	}
	stmtGetSeederPeers = preparedStmt{
		`SELECT COUNT(*) FROM "public"."peer" WHERE torrent_id = decode($1, 'base64') AND "left" = 0`,
		nil,
	}
	stmtGetPeer = preparedStmtx{
		`SELECT * FROM "public"."peer" WHERE id = decode($1, 'base64') AND torrent_id = decode($2, 'base64')`,
		nil,
	}
	stmtInsertNewPeer = preparedStmt{
		`INSERT INTO "public"."peer" ("id", "torrent_id", "state", "ip", "port", "downloaded", "uploaded", "left") VALUES (decode($1, 'base64'), decode($2, 'base64'), $3, $4, $5, $6, $7, $8)`,
		nil,
	}
	stmtUpdatePeer = preparedStmt{
		`UPDATE "public"."peer" SET "state" = $3, "ip" = $4, "port" = $5, "downloaded" = $6, "uploaded" = $7, "left" = $8, "last_updated" = NOW() WHERE id = decode($1, 'base64') AND torrent_id = decode($2, 'base64')`,
		nil,
	}
	stmtGetTorrentStats = preparedStmt{
		`SELECT "completed", "downloaded" FROM "public"."torrent" WHERE "hash" = decode($1, 'base64')`,
		nil,
	}
)

func prepareStatements() error {
	var err error
	StmtWriteConnectionID.Stmt, err = DB.Prepare(StmtWriteConnectionID.String)
	if err != nil {
		logger.Error("error while preparing StmtWriteConnectionID", "err", err)
		return err
	}

	stmtCheckConnectionID.Stmt, err = DB.Prepare(stmtCheckConnectionID.String)
	if err != nil {
		logger.Error("error while preparing stmtCheckConnectionID", "err", err)
		return err
	}

	stmtGetPeers.Stmt, err = DB.Prepare(stmtGetPeers.String)
	if err != nil {
		logger.Error("error while preparing stmtGetPeers", "err", err)
		return err
	}

	stmtGetLeecherPeers.Stmt, err = DB.Prepare(stmtGetLeecherPeers.String)
	if err != nil {
		logger.Error("error while preparing stmtGetLeecherPeers", "err", err)
		return err
	}

	stmtGetSeederPeers.Stmt, err = DB.Prepare(stmtGetSeederPeers.String)
	if err != nil {
		logger.Error("error while preparing stmtGetSeederPeers", "err", err)
		return err
	}

	stmtGetPeer.Stmt, err = DB.Preparex(stmtGetPeer.String)
	if err != nil {
		logger.Error("error while preparing stmtGetPeer", "err", err)
		return err
	}

	stmtInsertNewPeer.Stmt, err = DB.Prepare(stmtInsertNewPeer.String)
	if err != nil {
		logger.Error("error while preparing stmtInsertNewPeer", "err", err)
		return err
	}

	stmtUpdatePeer.Stmt, err = DB.Prepare(stmtUpdatePeer.String)
	if err != nil {
		logger.Error("error while preparing stmtUpdatePeer", "err", err)
		return err
	}

	stmtGetTorrentStats.Stmt, err = DB.Prepare(stmtGetTorrentStats.String)
	if err != nil {
		logger.Error("error while preparing stmtGetTorrentStats", "err", err)
		return err
	}
	return nil
}
