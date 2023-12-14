-- SQL Schema for Desserted

-- Table: Users
-- Stores user account information
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password BYTEA NOT NULL,
  password_changed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  total_score INT NOT NULL DEFAULT 0,
  total_wins INT NOT NULL DEFAULT 0,
  total_losses INT NOT NULL DEFAULT 0
);
CREATE INDEX idx_user_username ON users(username);

-- Table: Games
-- Represents a game session
CREATE TABLE games (
  game_id SERIAL PRIMARY KEY,
  status VARCHAR(10) NOT NULL DEFAULT 'active',
  created_by BIGINT NOT NULL,
  start_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  end_time TIMESTAMPTZ,
  FOREIGN KEY (created_by) REFERENCES users(id)
);
CREATE INDEX idx_games_status ON games(status);

-- Table: PlayerGame
-- Associates users with their game sessions and tracks their progress
CREATE TABLE player_game (
  player_game_id SERIAL PRIMARY KEY,
  player_id BIGINT NOT NULL,
  game_id INT NOT NULL,
  player_score INT DEFAULT 0,
  player_status VARCHAR(50) DEFAULT 'active',
  hand_cards TEXT,
  played_cards TEXT,
  FOREIGN KEY (player_id) REFERENCES users(id),
  FOREIGN KEY (game_id) REFERENCES games(game_id)
);
CREATE INDEX idx_playergame_player_id ON player_game(player_id);
CREATE INDEX idx_playergame_game_id ON player_game(game_id);

-- Table: Cards
-- Details of each card used in the game
CREATE TABLE cards (
  card_id BIGSERIAL PRIMARY KEY,
  type VARCHAR NOT NULL,
  name VARCHAR NOT NULL
);

-- Table: Desserts
-- Stores information about desserts
CREATE TABLE desserts (
  dessert_id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  points INT NOT NULL
);

-- Table: PlayerHand
-- Tracks the cards in each player's hand for a game
CREATE TABLE player_hand (
  player_hand_id SERIAL PRIMARY KEY,
  player_game_id INT NOT NULL,
  card_id BIGINT NOT NULL,
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (card_id) REFERENCES cards(card_id)
);

-- Table: PlayedCards
-- Records the cards played by players in each game
CREATE TABLE played_cards (
  played_card_id SERIAL PRIMARY KEY,
  player_game_id INT NOT NULL,
  card_id BIGINT NOT NULL,
  play_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (card_id) REFERENCES cards(card_id)
);

-- Table: DessertPlayed
-- Records desserts created by players in each game
CREATE TABLE dessert_played (
  dessert_played_id SERIAL PRIMARY KEY,
  player_game_id INT NOT NULL,
  dessert_id INT NOT NULL,
  icon_path VARCHAR(255),
  timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (dessert_id) REFERENCES desserts(dessert_id)
);
CREATE INDEX idx_dessertplayed_timestamp ON dessert_played(timestamp);

COMMENT ON TABLE "users" IS 'Stores user account information';

COMMENT ON TABLE "games" IS 'Represents a game session';

COMMENT ON TABLE "player_game" IS 'Associates users with their game sessions and tracks their progress';

COMMENT ON TABLE "cards" IS 'Details of each card used in the game';