# Desserted: The Dessert Card Game

## Overview
Desserted is a strategic, fun, and engaging online card game where players combine ingredients to craft delicious desserts and score points. It's a game of wit, planning, and a dash of luck!

## Key Features
- **Multiplayer Gameplay**: Compete against friends or online opponents in real-time.
- **Strategic Depth**: Mix and match ingredients to create high-scoring desserts.
- **Dynamic Play**: Each game is unique with a variety of ingredients and special cards.

## Getting Started

### Prerequisites
- Go (for the backend)
- PostgreSQL (for the database)
- gRPC and Protobuf (for API design and communication)
- WebSocket (for real-time interaction)
- dbdocs (for database documentation)
- sqlc (for generating Go code from SQL)
- Node.js (required for development environment setup)

### Installation
- Clone the repository: `git clone [github.com/PlatosCodes/Desserted]`
- Backend setup: `cd backend && go install`
- Frontend setup (React): `cd frontend && npm install`

### Running the Application
- Start the backend server: `go run server.go`
- Start the frontend application (React): `npm start`

## Development

### Project Structure
- `backend/`: Contains all backend Go code.
- `frontend/`: Contains all frontend React code.
- `docs/`: Additional project documentation.
- `tests/`: Test scripts and files.

### Advanced Tools
- **gRPC & Protobuf**: Used for defining and implementing highly efficient and scalable API services.
- **WebSocket**: Enables real-time, bi-directional communication between client and server.
- **Swagger**: Utilized for API documentation and exploration.
- **dbdocs**: For visualizing and sharing the database schema. (https://dbdocs.io/codingplato/Desserted)
- **sqlc**: Generating type-safe Go code from SQL.

### Deployment
- **Initial Testing**: Deployed on AWS free tier for the beta version, ensuring cost-effectiveness and ease of setup.
- **Full Version**: Planned deployment on AWS for scalability and robustness to handle a larger user base.

## Authors
- Alexander Merola - *Initial work* - [PlatosCodes](https://github.com/yourprofile)

## License
This project is licensed under the GNU AFFERO GENERAL PUBLIC LICENSE - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments
- Special thanks to my wife, who came up with the idea for this game, for her ongoing support to pursue my dreams of being a professional software engineer.
