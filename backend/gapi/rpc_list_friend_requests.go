package gapi

import (
	"context"
	"database/sql"
	"log"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListFriendRequests lists all pending friend requests for a user.
func (server *Server) ListFriendRequests(ctx context.Context, req *pb.ListFriendRequestsRequest) (*pb.ListFriendRequestsResponse, error) {
	user_id := req.GetUserId()
	if user_id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user ID is required")
	}

	friendRequests, err := server.Store.ListPendingFriendRequests(ctx, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "friend requests not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve friend requests")
	}

	var pbFriendRequests []*pb.FriendRequest
	for _, fr := range friendRequests {
		pbFriendRequests = append(pbFriendRequests, &pb.FriendRequest{
			FriendshipId:     fr.FriendshipID,
			FrienderId:       fr.ID,
			FrienderUsername: fr.Username,
			FriendedAt:       timestamppb.New(fr.FriendedAt),
		})
	}

	log.Println("HEY PAL:", &pb.ListFriendRequestsResponse{
		FriendRequests: pbFriendRequests,
	})

	return &pb.ListFriendRequestsResponse{
		FriendRequests: pbFriendRequests,
	}, nil
}
