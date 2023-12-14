CREATE TABLE "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar,
  "email" varchar UNIQUE,
  "password_hash" varchar,
  "created_at" timestamp,
  "last_login" timestamp
);

CREATE TABLE "games" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "status" varchar,
  "created_at" timestamp,
  "ended_at" timestamp
);

CREATE TABLE "players" (
  "user_id" integer,
  "game_id" integer,
  "score" integer,
  "hand_cards" text,
  "played_cards" text
);

CREATE TABLE "cards" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "type" varchar,
  "name" varchar,
  "points" integer
);

COMMENT ON TABLE "users" IS 'Stores user account information';

COMMENT ON TABLE "games" IS 'Represents a game session';

COMMENT ON TABLE "players" IS 'Associates users with their game sessions and tracks their progress';

COMMENT ON TABLE "cards" IS 'Details of each card used in the game';

ALTER TABLE "players" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "players" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id");
