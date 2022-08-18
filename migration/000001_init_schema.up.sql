CREATE TABLE "driver" (
  "id" bigserial PRIMARY KEY,
  "phone" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" date,
  "avatar_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar NOT NULL DEFAULT 'finding',
  "latitude" float8,
  "longitude" float8
);

CREATE INDEX ON "driver" ("name");

CREATE INDEX ON "driver" ("phone");
