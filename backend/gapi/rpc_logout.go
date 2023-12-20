package gapi

import (
	"context"
	"database/sql"

	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Logout invalidates the user's session.
func (server *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*emptypb.Empty, error) {
	sessionIDString := req.GetSessionId()
	if sessionIDString == "" {
		return nil, status.Errorf(codes.InvalidArgument, "session ID is required")
	}

	// Parse the sessionID string into a UUID
	sessionID, err := uuid.Parse(sessionIDString)
	if err != nil {
		// Handle the error if the sessionID is not a valid UUID
		return nil, status.Errorf(codes.InvalidArgument, "invalid session ID")
	}

	// Verify if the session exists and is valid
	session, err := server.Store.GetSession(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "session not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get session: %v", err)
	}

	// Invalidate the session
	err = server.Store.BlockSession(ctx, session.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to invalidate session: %v", err)
	}

	return &emptypb.Empty{}, nil
}
