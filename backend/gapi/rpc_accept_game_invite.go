package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AcceptGameInvite allows a user to accept an invitation to a game.
func (server *Server) AcceptGameInvite(ctx context.Context, req *pb.AcceptGameInviteRequest) (*emptypb.Empty, error) {
	// Authorize and authenticate the user making the request
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	invitee, gameID := req.GetInviteePlayerId(), req.GetGameId()

	// Validate if the user is the invitee
	if authPayload.UserID != invitee {
		return nil, status.Errorf(codes.PermissionDenied, "you can only accept invitations for yourself")
	}

	// Check if the invitation exists
	invitationExists, err := server.Store.DoesInvitationExist(ctx, db.DoesInvitationExistParams{
		InviteePlayerID: invitee,
		GameID:          gameID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error checking invitation existence: %v", err)
	}
	if !invitationExists {
		return nil, status.Errorf(codes.NotFound, "no invitation found for the specified game")
	}

	// Accept the game invitation
	err = server.Store.AcceptGameInvitation(ctx, db.AcceptGameInvitationParams{
		ID:     invitee,
		GameID: gameID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error accepting invitation: %v", err)
	}

	// Delete the invitation after accepting it
	err = server.Store.DeleteGameInvitation(ctx, db.DeleteGameInvitationParams{
		InviteePlayerID: invitee,
		GameID:          gameID,
	})
	if err != nil {
		// Log the error but do not fail the operation
		return nil, status.Errorf(codes.Internal, "failed to delete invitation: %v", err)
	}

	return &emptypb.Empty{}, nil
}
