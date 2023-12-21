package gapi

import (
	"context"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateFriendship creates a new friendship between two users.
func (server *Server) CreateFriendship(ctx context.Context, req *pb.CreateFriendshipRequest) (*pb.CreateFriendshipResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Ensure that both user IDs are valid and not the same
	if req.GetFrienderId() <= 0 || len(req.GetFriendeeUsername()) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid users")
	}

	if req.GetFrienderId() != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you can only create friendships for yourself")
	}

	// Insert the new friendship into the database
	friendship, err := server.Store.CreateFriendship(ctx, db.CreateFriendshipParams{
		FrienderID: req.GetFrienderId(),
		Username:   req.GetFriendeeUsername(),
	})
	if err != nil {
		log.Printf("Failed to create friendship: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to create friendship: %v", err)
	}

	// Create and return the response
	resp := &pb.CreateFriendshipResponse{
		Friendship: &pb.Friend{
			FriendshipId: friendship.FriendshipID,
			FrienderId:   friendship.FrienderID,
			FriendeeId:   friendship.FriendeeID,
			FriendedAt:   timestamppb.New(friendship.FriendedAt),
		},
	}

	return resp, nil
}
