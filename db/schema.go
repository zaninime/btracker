package db

import (
	"bytes"

	"github.com/lib/pq"
)

const latestSchemaVersion = 2

func checkAndUpdateSchema() error {
	row := DB.QueryRow(`SELECT "value" FROM "public"."schema" WHERE "key"='version'`)
	var schemaVersion int
	err := row.Scan(&schemaVersion)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "42P01" {
				logger.Info("schema version not found, initializing schema")
				if err = runSchemaInit(); err != nil {
					logger.Error("cannot create schema", "err", err)
					return err
				}
				schemaVersion = latestSchemaVersion
			} else {
				logger.Error("unexpected postgres error", "err", err)
				return err
			}
		} else {
			logger.Error("unexpected error", "err", err)
			return err
		}
	}

	if schemaVersion < latestSchemaVersion {
		logger.Info("schema version is old, upgrading is required", "current", schemaVersion, "latest", latestSchemaVersion)
		for schemaVersion < latestSchemaVersion {
			err = updateSchema(schemaVersion)
			if err != nil {
				logger.Fatal("couldn't update database schema")
				return err
			}
			schemaVersion++
		}
	} else {
		logger.Debug("running the latest schema version", "version", schemaVersion)
	}
	return nil
}

func runSchemaInit() error {
	logger.Debug("running schema initialization", "version", latestSchemaVersion)
	return runTxQueries(MustAsset("init_v2.sql"))
}

func updateSchema(currentVersion int) error {
	logger.Debug("running schema update", "next", currentVersion+1)
	switch currentVersion {
	case 1:
		return runTxQueries(MustAsset("v1_to_v2.sql"))
	}
	return nil
}

func runTxQueries(queries []byte) error {
	queryList := bytes.Split(queries, []byte("-- \\run\\"))
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("couldn't begin transaction", "err", err)
		return err
	}
	for _, q := range queryList {
		q = bytes.TrimSpace(q)
		qStr := string(q)
		logger.Debug("running query", "q", qStr)
		_, err := tx.Exec(qStr)
		if err != nil {
			logger.Error("couldn't execute query, rolling back transaction", "err", err)
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		logger.Error("couldn't commit transaction", "err", err)
		return err
	}
	return nil
}
