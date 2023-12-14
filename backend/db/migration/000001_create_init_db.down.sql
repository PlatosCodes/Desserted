-- SQL Down Migration for Desserted

-- Drop all foreign keys before dropping tables
ALTER TABLE player_hand DROP CONSTRAINT IF EXISTS player_hand_player_game_id_fkey;
ALTER TABLE player_hand DROP CONSTRAINT IF EXISTS player_hand_card_id_fkey;
ALTER TABLE played_cards DROP CONSTRAINT IF EXISTS played_cards_player_game_id_fkey;
ALTER TABLE played_cards DROP CONSTRAINT IF EXISTS played_cards_card_id_fkey;
ALTER TABLE dessert_played DROP CONSTRAINT IF EXISTS dessert_played_player_game_id_fkey;
ALTER TABLE dessert_played DROP CONSTRAINT IF EXISTS dessert_played_dessert_id_fkey;
ALTER TABLE player_game DROP CONSTRAINT IF EXISTS player_game_player_id_fkey;
ALTER TABLE player_game DROP CONSTRAINT IF EXISTS player_game_game_id_fkey;
ALTER TABLE games DROP CONSTRAINT IF EXISTS games_created_by_fkey;

-- Drop tables in reverse order
DROP TABLE IF EXISTS dessert_played;
DROP TABLE IF EXISTS played_cards;
DROP TABLE IF EXISTS player_hand;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS player_game;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS users;
