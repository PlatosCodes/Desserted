# API Endpoints for Desserted

## User Authentication
- `POST /api/auth/register`: User registration.
- `POST /api/auth/login`: User login.
- `GET /user/{userId}`: Retrieve user profile and stats.

## Game Management
- `POST /api/game/create`: Create a new game session.
- `POST /api/game/join`: Join an existing game session.
- `GET /api/game/status`: Get the current status of the game.

## Player Actions
- `POST /api/game/draw`: Draw a card.
- `POST /api/game/play`: Play a dessert or a special card.
- `POST /api/game/endTurn`: End the current player's turn.
- `GET /api/game/score/{gameId}`: Retrieve the current score for all players.
- `POST /api/game/usecard`: Use a special card.
- `GET /api/game/end/{gameId}`: Check for game end conditions and final scores.