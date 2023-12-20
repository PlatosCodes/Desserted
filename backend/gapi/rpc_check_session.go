// rpc_check_session.go

package gapi

import (
	"context"
	"database/sql"
	"time"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (server *Server) CheckUserSession(ctx context.Context, req *pb.CheckUserSessionRequest) (*pb.CheckUserSessionResponse, error) {

	sessionIDString := req.GetSessionId()

	// Parse the sessionID string into a UUID
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		// Handle the error if the sessionID is not a valid UUID
		return nil, status.Errorf(codes.InvalidArgument, "invalid session ID")
	}

	session, err := server.Store.GetSession(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "session not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve session")
	}

	// Logic to check if the session is valid
	isValid := !session.IsBlocked && session.ExpiresAt.After(time.Now())

	return &pb.CheckUserSessionResponse{
		IsAuthenticated: isValid,
	}, nil
}
