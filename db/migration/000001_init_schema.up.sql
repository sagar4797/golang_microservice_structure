CREATE EXTENSION IF NOT EXISTS "uuid-ossp" schema pg_catalog;
CREATE TABLE "student"(
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar ,
  "standard" varchar,
  "subjects" varchar[],
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);
