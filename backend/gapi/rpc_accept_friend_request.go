package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AcceptFriendRequest updates the status of a friend request to accepted.
func (server *Server) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*emptypb.Empty, error) {
	user_id := req.GetUserId()
	friendship_id := req.GetFriendshipId()

	if user_id == 0 || friendship_id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user ID and friendship ID are required")
	}

	err := server.Store.AcceptFriendRequest(ctx, db.AcceptFriendRequestParams{
		FriendeeID:   user_id,
		FriendshipID: friendship_id,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error in accepting friendship")
	}

	return &emptypb.Empty{}, nil
}
