ALTER TABLE "public"."peer" ALTER COLUMN "downloaded" TYPE bigint, ALTER COLUMN "uploaded" TYPE bigint, ALTER COLUMN "left" TYPE bigint;

-- \run\

UPDATE "public"."schema" SET "value"=2 WHERE "key"='version';
