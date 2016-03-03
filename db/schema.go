package db

import (
	"strings"

	"github.com/lib/pq"
)

const lastSchemaVersion = 1

const (
	schemaInit = `CREATE TABLE "public"."torrent" (
	  "id" serial PRIMARY KEY,
	  "completed" integer NOT NULL DEFAULT 0,
	  "hash" bytea NOT NULL
	);

	-- \run\

	CREATE TABLE "public"."peer" (
	  "id" bytea NOT NULL,
	  "torrent_id" integer REFERENCES "torrent"("id") NOT NULL,
	  "state" integer NOT NULL,
	  "ip" inet NOT NULL,
	  "port" integer NOT NULL,
	  "downloaded" integer NOT NULL,
	  "uploaded" integer NOT NULL,
	  "left" integer NOT NULL,
	  "last_updated" timestamp NOT NULL DEFAULT NOW(),
	  PRIMARY KEY ("id", "torrent_id")
	);

	-- \run\

	CREATE TABLE "public"."connection" (
	  "id" bytea PRIMARY KEY,
	  "ip" inet NOT NULL,
	  "expiry" timestamp NOT NULL
	);

	-- \run\

	CREATE TABLE "public"."schema" (
	  key varchar PRIMARY KEY,
	  value integer NOT NULL
	);

	-- \run\

	CREATE INDEX ON "public"."torrent" ("hash");

	-- \run\

	CREATE INDEX ON "public"."peer" ("id", "torrent_id", "state");

	-- \run\

	INSERT INTO "public"."schema" VALUES ('version', 1);`
)

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
					logger.Fatal("cannot create schema", "err", err)
				}
				schemaVersion = lastSchemaVersion
			} else {
				logger.Fatal("unexpected postgres error", "err", err)
			}
		} else {
			logger.Fatal("unexpected error", "err", err)
		}
	}

	if schemaVersion < lastSchemaVersion {
		logger.Info("schema version is old, upgrading is required")
	} else {
		logger.Debug("running the latest schema version")
	}
	return nil
}

func runSchemaInit() error {
	logger.Debug("running schema initialization")
	queries := strings.Split(schemaInit, "-- \\run\\")
	for _, q := range queries {
		q = strings.TrimSpace(q)
		logger.Debug("running query", "q", q)
		_, err := DB.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}
