CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" bytea NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "games" (
  "id" bigserial PRIMARY KEY,
  "status" varchar DEFAULT 'active',
  "created_by" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "ended_at" timestamp
);

CREATE TABLE "players" (
  "user_id" bigint,
  "game_id" bigint,
  "score" integer DEFAULT 0,
  "hand_cards" text,
  "played_cards" text
);

CREATE TABLE "desserts" (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "points" INT NOT NULL,
  "type" TEXT NOT NULL
);

CREATE TABLE "cards" (
  "id" bigserial PRIMARY KEY,
  "type" varchar,
  "name" varchar
);

COMMENT ON TABLE "users" IS 'Stores user account information';

COMMENT ON TABLE "games" IS 'Represents a game session';

COMMENT ON TABLE "players" IS 'Associates users with their game sessions and tracks their progress';

COMMENT ON TABLE "cards" IS 'Details of each card used in the game';

ALTER TABLE "games" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "players" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "players" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id");
