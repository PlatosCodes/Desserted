package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListUserFriends lists the friendships of a given user.
func (server *Server) ListUserFriends(ctx context.Context, req *pb.ListUserFriendsRequest) (*pb.ListUserFriendsResponse, error) {
	// Validate the request
	if req.GetUserId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user ID is required")
	}

	// Fetch the friendships from the database
	friendships, err := server.Store.ListUserFriends(ctx, db.ListUserFriendsParams{
		FrienderID: req.GetUserId(),
		Limit:      req.GetLimit(),
		Offset:     req.GetOffset(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list friendships: %v", err)
	}

	// Convert the database friendships to protobuf friendships
	var pbFriendships []*pb.Friend
	for _, f := range friendships {
		pbFriendships = append(pbFriendships, &pb.Friend{
			FriendshipId: f.FriendshipID,
			FrienderId:   f.FrienderID,
			FriendeeId:   f.FriendeeID,
			FriendedAt:   timestamppb.New(f.FriendedAt),
		})
	}

	// Create and return the response
	resp := &pb.ListUserFriendsResponse{
		Friendships: pbFriendships,
	}
	return resp, nil
}
