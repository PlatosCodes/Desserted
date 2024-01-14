-- SQL Schema for Desserted

-- Table: Users
-- Stores user account information
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email CITEXT UNIQUE NOT NULL,
  password BYTEA NOT NULL,
  activated BOOLEAN NOT NULL DEFAULT false,
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
  game_id BIGSERIAL PRIMARY KEY,
  status VARCHAR(10) NOT NULL DEFAULT 'waiting',
  created_by BIGINT NOT NULL,
  number_of_players INT NOT NULL DEFAULT 0,
  current_turn INT NOT NULL DEFAULT 0,
  current_player_number INT,
  start_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  last_action_time TIMESTAMPTZ,
  end_time TIMESTAMPTZ,
  FOREIGN KEY (created_by) REFERENCES users(id)
);
CREATE INDEX idx_games_status ON games(status);

-- Table: Cards
-- Details of each card used in the game
CREATE TABLE cards (
  card_id BIGSERIAL PRIMARY KEY,
  type VARCHAR NOT NULL,
  name VARCHAR NOT NULL
);

-- Table: GameDeck
-- Stores the deck of cards for each game
CREATE TABLE game_deck (
  game_deck_id BIGSERIAL PRIMARY KEY,
  game_id BIGINT NOT NULL,
  card_id BIGINT NOT NULL,
  order_index INT NOT NULL,
  FOREIGN KEY (game_id) REFERENCES games(game_id),
  FOREIGN KEY (card_id) REFERENCES cards(card_id)
);
CREATE INDEX idx_game_deck_game_id ON game_deck(game_id);
CREATE INDEX idx_game_deck_order ON game_deck(order_index);

-- Table: PlayerGame
-- Associates users with their game sessions and tracks their progress
CREATE TABLE player_game (
  player_game_id BIGSERIAL PRIMARY KEY,
  player_id BIGINT NOT NULL,
  game_id BIGINT NOT NULL,
  player_number INT,
  player_score INT NOT NULL DEFAULT 0,
  player_status VARCHAR(10) NOT NULL DEFAULT 'active',
  FOREIGN KEY (player_id) REFERENCES users(id),
  FOREIGN KEY (game_id) REFERENCES games(game_id)
);
ALTER TABLE player_game ADD CONSTRAINT unique_player_game UNIQUE (player_id, game_id);
CREATE INDEX idx_playergame_player_id ON player_game(player_id);
CREATE INDEX idx_playergame_game_id ON player_game(game_id);

-- Table: PlayerHand
-- Tracks the cards in each player's hand for a game
CREATE TABLE player_hand (
  player_hand_id BIGSERIAL PRIMARY KEY,
  player_game_id BIGINT NOT NULL,
  card_id BIGINT NOT NULL,
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (card_id) REFERENCES cards(card_id)
);

-- Table: PlayedCards
-- Records the cards played by players in each game
CREATE TABLE played_cards (
  played_card_id BIGSERIAL PRIMARY KEY,
  player_game_id BIGINT NOT NULL,
  card_id BIGINT NOT NULL,
  play_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (card_id) REFERENCES cards(card_id)
);

-- Table: Desserts
-- Stores information about desserts
CREATE TABLE desserts (
  dessert_id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  points INT NOT NULL
);

-- Table: DessertPlayed
-- Records desserts created by players in each game
CREATE TABLE dessert_played (
  dessert_played_id BIGSERIAL PRIMARY KEY,
  player_game_id BIGINT NOT NULL,
  dessert_id BIGINT NOT NULL,
  icon_path VARCHAR(255),
  timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id),
  FOREIGN KEY (dessert_id) REFERENCES desserts(dessert_id)
);
CREATE INDEX idx_dessertplayed_timestamp ON dessert_played(timestamp);

-- Table: GameInvitations
-- Records game invitations created by players who created a game
CREATE TABLE game_invitations (
  game_invitation_id BIGSERIAL PRIMARY KEY,
  inviter_player_id BIGINT NOT NULL,
  invitee_player_id BIGINT NOT NULL,
  game_id BIGINT NOT NULL,
  invitation_status VARCHAR(10) NOT NULL DEFAULT 'pending',
  timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (inviter_player_id) REFERENCES users(id),
  FOREIGN KEY (invitee_player_id) REFERENCES users(id),
  FOREIGN KEY (game_id) REFERENCES games(game_id)
);

-- Table: PlayerTurnActionsTable
-- Records the actions player has taken on a turn
CREATE TABLE player_turn_actions (
  player_game_id BIGINT PRIMARY KEY,
  card_drawn BOOLEAN NOT NULL DEFAULT FALSE,
  dessert_played BOOLEAN NOT NULL DEFAULT FALSE,
  special_card_played BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY (player_game_id) REFERENCES player_game(player_game_id)
);

-- Table: Friends
-- Records friends
CREATE TABLE friends (
  friendship_id BIGSERIAL PRIMARY KEY,
  friender_id BIGINT NOT NULL,
  friendee_id BIGINT NOT NULL,
  status VARCHAR(10) NOT NULL DEFAULT 'pending',
  friended_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  accepted_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00',
  FOREIGN KEY (friender_id) REFERENCES users(id),
  FOREIGN KEY (friendee_id) REFERENCES users(id)
);
CREATE INDEX idx_friender_id ON friends(friender_id);
CREATE INDEX idx_friendee_id ON friends(friendee_id);

COMMENT ON TABLE "users" IS 'Stores user account information';

COMMENT ON TABLE "games" IS 'Represents a game session';

COMMENT ON TABLE "player_game" IS 'Associates users with their game sessions and tracks their progress';

COMMENT ON TABLE "cards" IS 'Details of each card used in the game';