package gapi

import (
	"fmt"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/mailer"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/PlatosCodes/desserted/backend/token"
	"github.com/PlatosCodes/desserted/backend/util"
)

// Server serves gRPC requests for our desserted service.
type Server struct {
	pb.UnimplementedDessertedServer
	config     util.Config
	Store      db.Store
	tokenMaker token.Maker
	mailer     *mailer.Mailer
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store, mailer *mailer.Mailer) (*Server, token.Maker, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		Store:      store,
		tokenMaker: tokenMaker,
		mailer:     mailer,
	}

	return server, tokenMaker, nil
}
