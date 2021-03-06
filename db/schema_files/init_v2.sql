CREATE TABLE "public"."torrent" (
  "hash" bytea PRIMARY KEY,
  "downloaded" integer NOT NULL DEFAULT 0
);

-- \run\

CREATE TABLE "public"."peer" (
  "id" bytea NOT NULL,
  "torrent_id" bytea NOT NULL,
  "state" integer NOT NULL,
  "ip" inet NOT NULL,
  "port" integer NOT NULL,
  "downloaded" bigint NOT NULL,
  "uploaded" bigint NOT NULL,
  "left" bigint NOT NULL,
  "last_updated" timestamp with time zone NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("id", "torrent_id")
);

-- \run\

CREATE TABLE "public"."connection" (
  "id" bytea NOT NULL,
  "ip" inet NOT NULL,
  "expiry" timestamp with time zone NOT NULL,
  PRIMARY KEY ("id", "ip")
);

-- \run\

CREATE TABLE "public"."schema" (
  key varchar PRIMARY KEY,
  value integer NOT NULL
);

-- \run\

CREATE INDEX ON "public"."peer" ("id", "torrent_id", "state");

-- \run\

INSERT INTO "public"."schema" VALUES ('version', 2);
