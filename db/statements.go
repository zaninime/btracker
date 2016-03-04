package db

import (
	"database/sql"
	"encoding/base64"
	"net"
	"time"

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
	StmtCheckConnectionID = preparedStmt{
		`SELECT "ip" FROM "public"."connection" WHERE id = decode($1, 'base64') AND ip = $2 AND expiry >= $3`,
		nil,
	}
	StmtGetPeers = preparedStmt{
		`SELECT "ip", "port" FROM "public"."peer" WHERE id = decode($1, 'base64') AND torrent_id = decode($2, 'base64') AND "state" > 0 LIMIT $3`,
		nil,
	}
	StmtGetLeecherPeers = preparedStmt{
		`SELECT COUNT(*) FROM "public"."peer" WHERE torrent_id = decode($1, 'base64') AND "left" > 0`,
		nil,
	}
	StmtGetSeederPeers = preparedStmt{
		`SELECT COUNT(*) FROM "public"."peer" WHERE torrent_id = decode($1, 'base64') AND "left" = 0`,
		nil,
	}
	StmtGetPeer = preparedStmtx{
		`SELECT * FROM "public"."peer" WHERE id = decode($1, 'base64') AND torrent_id = decode($2, 'base64')`,
		nil,
	}
	StmtInsertNewPeer = preparedStmt{
		`INSERT INTO "public"."peer" ("id", "torrent_id", "state", "ip", "port", "downloaded", "uploaded", "left") VALUES (decode($1, 'base64'), decode($2, 'base64'), $3, $4, $5, $6, $7, $8)`,
		nil,
	}
	StmtUpdatePeer = preparedStmt{
		`UPDATE "public"."peer" SET "state" = $3, "ip" = $4, "port" = $5, "downloaded" = $6, "uploaded" = $7, "left" = $8, "last_updated" = NOW() WHERE id = decode($1, 'base64') AND torrent_id = decode($2, 'base64')`,
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

	StmtCheckConnectionID.Stmt, err = DB.Prepare(StmtCheckConnectionID.String)
	if err != nil {
		logger.Error("error while preparing StmtCheckConnectionID", "err", err)
		return err
	}

	StmtGetPeers.Stmt, err = DB.Prepare(StmtGetPeers.String)
	if err != nil {
		logger.Error("error while preparing StmtGetPeers", "err", err)
		return err
	}

	StmtGetLeecherPeers.Stmt, err = DB.Prepare(StmtGetLeecherPeers.String)
	if err != nil {
		logger.Error("error while preparing StmtGetLeecherPeers", "err", err)
		return err
	}

	StmtGetSeederPeers.Stmt, err = DB.Prepare(StmtGetSeederPeers.String)
	if err != nil {
		logger.Error("error while preparing StmtGetSeederPeers", "err", err)
		return err
	}

	StmtGetPeer.Stmt, err = DB.Preparex(StmtGetPeer.String)
	if err != nil {
		logger.Error("error while preparing StmtGetPeer", "err", err)
		return err
	}

	StmtInsertNewPeer.Stmt, err = DB.Prepare(StmtInsertNewPeer.String)
	if err != nil {
		logger.Error("error while preparing StmtInsertNewPeer", "err", err)
		return err
	}

	StmtUpdatePeer.Stmt, err = DB.Prepare(StmtUpdatePeer.String)
	if err != nil {
		logger.Error("error while preparing StmtUpdatePeer", "err", err)
		return err
	}
	return nil
}

// CheckConnectionID queries the database for connection id validity
func CheckConnectionID(connectionID []byte, addr net.IP) (bool, error) {
	var ipAddr string
	logger.Debug("running query", "q", StmtCheckConnectionID.String, "connectionID", connectionID, "addr", addr.String())
	err := StmtCheckConnectionID.Stmt.QueryRow(base64.StdEncoding.EncodeToString(connectionID), addr.String(), time.Now()).Scan(&ipAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
