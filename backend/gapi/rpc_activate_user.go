// File: gapi/rpc_activate_user.go

package gapi

import (
	"context"
	"log"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ActivateUser(ctx context.Context, req *pb.ActivateUserRequest) (*pb.ActivateUserResponse, error) {
	// Verify the activation token
	log.Println("ACTIVATE MEEEE", req)
	_, err := server.tokenMaker.VerifyToken(req.GetActivationToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid activation token: %s", err)
	}

	// Get the stored activation token
	storedToken, err := server.Store.GetActivationToken(ctx, req.ActivationToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get activation token from the db: %s", err)
	}

	if storedToken.UserID != req.GetUserId() {
		return nil, status.Errorf(codes.Internal, "user id mismatch: %s", err)
	}

	// Activate the user
	err = server.Store.ActivateUser(ctx, storedToken.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to activate user: %s", err)
	}

	// Delete the activation token
	err = server.Store.DeleteActivationToken(ctx, storedToken.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete activation token: %s", err)
	}

	rsp := &pb.ActivateUserResponse{
		Message: "User successfully activated",
	}

	return rsp, nil
}
