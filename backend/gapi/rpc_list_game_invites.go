package gapi

import (
	"context"
	"database/sql"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListGameInvites lists all pending game invitations for a user.
func (server *Server) ListGameInvites(ctx context.Context, req *pb.ListGameInvitesRequest) (*pb.ListGameInvitesResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	userID := req.GetUserId()

	// Validate the request
	if userID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user ID is required")
	}

	// Ensure that the requestor is the player in question or has appropriate permissions
	if userID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you do not have permission to view this game data")
	}

	gameInvitations, err := server.Store.ListGameInvitationsForUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "friend requests not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve friend requests")
	}

	var pbGameInvitations []*pb.GameInvitation
	for _, gi := range gameInvitations {

		pbGameInvitations = append(pbGameInvitations, &pb.GameInvitation{
			GameInvitationId: gi.GameInvitationID,
			InviterPlayerId:  gi.InviterPlayerID,
			InviteePlayerId:  gi.InviteePlayerID,
			GameId:           gi.GameID,
			InvitationStatus: gi.InvitationStatus,
			Timestamp:        timestamppb.New(gi.Timestamp),
		})
	}

	return &pb.ListGameInvitesResponse{
		GameInvite: pbGameInvitations,
	}, nil
}
