CREATE TABLE "driver" (
  "id" bigserial PRIMARY KEY,
  "phone" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" date,
  "avatar_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar NOT NULL DEFAULT 'finding',
  "latitude" float8 NOT NULL,
  "longitude" float8 NOT NULL
);

CREATE INDEX ON "driver" ("name");

CREATE INDEX ON "driver" ("phone");
