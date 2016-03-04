package db

import (
	"database/sql"
	"encoding/base64"
	"net"
	"time"

	"github.com/zaninime/btracker/udp"
)

// Peer represents the peer relation inside the database
type Peer struct {
	ID          []byte    `db:"id"`
	TorrentID   []byte    `db:"torrent_id"`
	State       int       `db:"state"`
	IP          string    `db:"ip"`
	Port        uint16    `db:"port"`
	Downloaded  int64     `db:"downloaded"`
	Uploaded    int64     `db:"uploaded"`
	Left        int64     `db:"left"`
	LastUpdated time.Time `db:"last_updated"`
}

// CheckConnectionID queries the database for connection id validity
func CheckConnectionID(connectionID []byte, addr net.IP) (bool, error) {
	var ipAddr string
	logger.Debug("running query", "q", stmtCheckConnectionID.String, "connectionID", connectionID, "addr", addr.String())
	err := stmtCheckConnectionID.Stmt.QueryRow(base64.StdEncoding.EncodeToString(connectionID), addr.String(), time.Now()).Scan(&ipAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logger.Error("couldn't query database for connection id", "err", err)
		return false, err
	}
	return true, nil
}

// PopulatePeersFields set peers list and seeders and leechers stats
func PopulatePeersFields(response *udp.AnnounceResponse, peerID []byte, torrentID []byte, limit int32) error {
	// stats first
	logger.Debug("running query", "q", stmtGetLeecherPeers.String, "torrent_id", torrentID)
	if err := stmtGetLeecherPeers.Stmt.QueryRow(base64.StdEncoding.EncodeToString(torrentID)).Scan(&response.Leechers); err != nil {
		logger.Error("couldn't query the database for leechers", "err", err)
		return err
	}
	logger.Debug("running query", "q", stmtGetSeederPeers.String, "torrent_id", torrentID)
	if err := stmtGetSeederPeers.Stmt.QueryRow(base64.StdEncoding.EncodeToString(torrentID)).Scan(&response.Seeders); err != nil {
		logger.Error("couldn't query the database for seeders", "err", err)
		return err
	}

	// peer list
	logger.Debug("running query", "q", stmtGetPeers.String, "torrent_id", torrentID, "id", peerID, "limit", limit)
	result, err := stmtGetPeers.Stmt.Query(base64.StdEncoding.EncodeToString(peerID), base64.StdEncoding.EncodeToString(torrentID), limit)
	defer result.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("couldn't query the database for peers", "err", err)
		return err
	}
	for result.Next() {
		var ipString string
		var port uint16
		err := result.Scan(&ipString, &port)
		if err != nil {
			logger.Error("couldn't scan results", "err", err)
			return err
		}
		ip := net.ParseIP(ipString)
		response.Peers = append(response.Peers, udp.Peer{IP: ip, Port: port})
	}
	return nil
}

// GetPeer returns a pointer to a Peer struct populated with the requested data
// Peer will be nil if error occurred or if not found inside the database
func GetPeer(peerID, torrentID []byte) (*Peer, error) {
	var peer Peer
	if err := stmtGetPeer.Stmt.QueryRowx(base64.StdEncoding.EncodeToString(peerID), base64.StdEncoding.EncodeToString(torrentID)).StructScan(&peer); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("couldn't query database for peer", "err", err)
		return nil, err
	}
	return &peer, nil
}

func InsertPeer(peer *Peer) error {
	logger.Debug("running query",
		"q", stmtInsertNewPeer.String,
		"id", base64.StdEncoding.EncodeToString(peer.ID),
		"torrent_id", base64.StdEncoding.EncodeToString(peer.TorrentID),
		"state", peer.State,
		"ip", peer.IP,
		"port", peer.Port,
		"downloaded", peer.Downloaded,
		"uploaded", peer.Uploaded,
		"left", peer.Left,
	)
	if _, err := stmtInsertNewPeer.Stmt.Exec(base64.StdEncoding.EncodeToString(peer.ID),
		base64.StdEncoding.EncodeToString(peer.TorrentID),
		peer.State,
		peer.IP,
		peer.Port,
		peer.Downloaded,
		peer.Uploaded,
		peer.Left); err != nil {
		logger.Error("couldn't insert peer into database", "err", err)
		return err
	}
	return nil
}

func UpdatePeer(peer *Peer) error {
	logger.Debug("running query",
		"q", stmtUpdatePeer.String,
		"id", base64.StdEncoding.EncodeToString(peer.ID),
		"torrent_id", base64.StdEncoding.EncodeToString(peer.TorrentID),
		"state", peer.State,
		"ip", peer.IP,
		"port", peer.Port,
		"downloaded", peer.Downloaded,
		"uploaded", peer.Uploaded,
		"left", peer.Left,
	)
	if _, err := stmtUpdatePeer.Stmt.Exec(
		base64.StdEncoding.EncodeToString(peer.ID),
		base64.StdEncoding.EncodeToString(peer.TorrentID),
		peer.State,
		peer.IP,
		peer.Port,
		peer.Downloaded,
		peer.Uploaded,
		peer.Left); err != nil {
		logger.Error("couldn't insert peer into database", "err", err)
		return err
	}
	return nil
}

func GetTorrentStats(torrentID []byte) (*udp.TorrentStats, error) {
	var completed, downloaded, leechers int32
	logger.Debug("running query", "q", stmtGetLeecherPeers.String, "torrent_id", torrentID)
	if err := stmtGetLeecherPeers.Stmt.QueryRow(base64.StdEncoding.EncodeToString(torrentID)).Scan(&leechers); err != nil {
		logger.Error("couldn't query the database for leechers", "err", err)
		return nil, err
	}
	logger.Debug("running query",
		"q", stmtGetTorrentStats.String,
		"torrent_id", base64.StdEncoding.EncodeToString(torrentID),
	)
	if err := stmtGetTorrentStats.Stmt.QueryRow(base64.StdEncoding.EncodeToString(torrentID)).Scan(&completed, &completed); err != nil {
		if err == sql.ErrNoRows {
			return &udp.TorrentStats{0, 0, leechers}, nil
		}
		logger.Error("couldn't retrieve torrent stats", "err", err)
		return nil, err
	}
	return &udp.TorrentStats{completed, downloaded, leechers}, nil
}

func IncrementTorrentCompletedStats(torrentID []byte) error {
	completed := 0
	torrentIDBase64 := base64.StdEncoding.EncodeToString(torrentID)
	getStmt := `SELECT "completed" FROM "public"."torrent" WHERE "hash" = decode($1, 'base64')`
	insertStmt := `INSERT INTO "public"."torrent" VALUES (decode($1, 'base64'), 1)`
	updateStmt := `UPDATE "public"."torrent" SET "completed" = $2 WHERE "hash" = decode($1, 'base64')`
	logger.Debug("beginning transaction")
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("couldn't begin transaction", "err", err)
		return err
	}
	logger.Debug("running query", "q", getStmt, "hash", torrentIDBase64)
	if err = tx.QueryRow(getStmt, torrentIDBase64).Scan(&completed); err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("running query", "q", insertStmt, "hash", torrentIDBase64)
			if _, err = tx.Exec(insertStmt, torrentIDBase64); err != nil {
				logger.Error("couldn't insert row into table torrent", "err", err)
				tx.Rollback()
				return err
			}
			if err = tx.Commit(); err != nil {
				logger.Error("couldn't commit transaction", "err", err)
			}
			logger.Debug("new row inserted", "hash", torrentIDBase64)
			return nil
		}
		logger.Error("couldn't read torrent stats from database", "err", err)
		tx.Rollback()
		return err
	}
	completed++
	logger.Debug("running query", "q", updateStmt, "hash", torrentIDBase64, "completed", completed)
	if _, err = tx.Exec(updateStmt, torrentIDBase64, completed); err != nil {
		logger.Error("couldn't update torrent stats", "err", err)
		tx.Rollback()
		return err
	}
	logger.Debug("torrent stats updated", "hash", torrentIDBase64, "completed", completed)
	tx.Commit()
	return nil
}

func IncrementTorrentDownloadedStats(torrentID []byte) error {
	downloaded := 0
	torrentIDBase64 := base64.StdEncoding.EncodeToString(torrentID)
	getStmt := `SELECT "downloaded" FROM "public"."torrent" WHERE "hash" = decode($1, 'base64')`
	insertStmt := `INSERT INTO "public"."torrent" VALUES (decode($1, 'base64'), 0, 1)`
	updateStmt := `UPDATE "public"."torrent" SET "downloaded" = $2 WHERE "hash" = decode($1, 'base64')`
	logger.Debug("beginning transaction")
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("couldn't begin transaction", "err", err)
		return err
	}
	logger.Debug("running query", "q", getStmt, "hash", torrentIDBase64)
	if err = tx.QueryRow(getStmt, torrentIDBase64).Scan(&downloaded); err != nil {
		if err == sql.ErrNoRows {
			logger.Debug("running query", "q", insertStmt, "hash", torrentIDBase64)
			if _, err = tx.Exec(insertStmt, torrentIDBase64); err != nil {
				logger.Error("couldn't insert row into table torrent", "err", err)
				tx.Rollback()
				return err
			}
			if err = tx.Commit(); err != nil {
				logger.Error("couldn't commit transaction", "err", err)
			}
			logger.Debug("new row inserted", "hash", torrentIDBase64)
			return nil
		}
		logger.Error("couldn't read torrent stats from database", "err", err)
		tx.Rollback()
		return err
	}
	downloaded++
	logger.Debug("running query", "q", updateStmt, "hash", torrentIDBase64, "downloaded", downloaded)
	if _, err = tx.Exec(updateStmt, torrentIDBase64, downloaded); err != nil {
		logger.Error("couldn't update torrent stats", "err", err)
		tx.Rollback()
		return err
	}
	logger.Debug("torrent stats updated", "hash", torrentIDBase64, "downloaded", downloaded)
	tx.Commit()
	return nil
}
