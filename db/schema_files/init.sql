CREATE TABLE "public"."torrent" (
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

INSERT INTO "public"."schema" VALUES ('version', 1);
