package db

import "time"

func ClearConnections() error {
	logger.Debug("clearing expired connections")
	stmt := `DELETE FROM "public"."connection" WHERE "expiry" < $1`
	now := time.Now()
	logger.Debug("running query", "q", stmt, "expiry", now)
	if _, err := DB.Exec(stmt, now); err != nil {
		logger.Error("ClearConnections: couldn't clear old connections from database", "err", err)
		return err
	}
	return nil
}

func ClearPeers(olderThan time.Duration) error {
	logger.Debug("clearing expired peer updates")
	stmt := `DELETE FROM "public"."peer" WHERE "last_updated" < $1`
	age := time.Now().Add(-olderThan * time.Second)
	logger.Debug("running query", "q", stmt, "last_updated", age)
	if _, err := DB.Exec(stmt, age); err != nil {
		logger.Error("ClearPeers: couldn't clear old connections from database", "err", err)
		return err
	}
	return nil
}
