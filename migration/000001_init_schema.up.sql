CREATE TABLE "driver" (
  "id" bigserial PRIMARY KEY,
  "phone" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" date,
  "avatar_url" varchar,
  "verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "driver" ("name");

CREATE INDEX ON "driver" ("phone");
