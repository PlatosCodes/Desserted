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
	// Ensure that both user IDs are valid and not the same
	if req.GetFrienderId() == req.GetFriendeeId() || req.GetFrienderId() == 0 || req.GetFriendeeId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user IDs")
	}

	// Insert the new friendship into the database
	friendship, err := server.Store.CreateFriendship(ctx, db.CreateFriendshipParams{
		FrienderID: req.GetFrienderId(),
		FriendeeID: req.GetFriendeeId(),
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
