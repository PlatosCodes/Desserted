# API Endpoints for Desserted

## User Authentication
- `POST /api/auth/register`: User registration.
- `POST /api/auth/login`: User login.

## Game Management
- `POST /api/game/create`: Create a new game session.
- `POST /api/game/join`: Join an existing game session.
- `GET /api/game/status`: Get the current status of the game.

## Player Actions
- `POST /api/game/draw`: Draw a card.
- `POST /api/game/play`: Play a dessert or a special card.
- `POST /api/game/endTurn`: End the current player's turn.
