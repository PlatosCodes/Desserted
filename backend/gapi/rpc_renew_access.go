// rpc_renew_access.go

package gapi

import (
	"context"
	"database/sql"
	"time"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RenewAccessToken(ctx context.Context, req *pb.RenewAccessRequest) (*pb.RenewAccessResponse, error) {
	refreshToken := req.GetRefreshToken()

	// Verify the refresh token and get the payload
	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	// Retrieve the session from the database
	session, err := server.Store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "session not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve session")
	}

	// Check if the session is valid
	if session.IsBlocked || time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid session")
	}

	// Create a new access token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.UserID,
		session.Username, // Assuming username is part of your session or payload
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	// Prepare the response
	rsp := &pb.RenewAccessResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiresAt.Time),
	}

	return rsp, nil
}
