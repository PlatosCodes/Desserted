package gapi

import (
	"context"
	"database/sql"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ListFriendRequests lists all pending friend requests for a user.
func (server *Server) ListFriendRequests(ctx context.Context, req *pb.ListFriendRequestsRequest) (*pb.ListFriendRequestsResponse, error) {
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

	friendRequests, err := server.Store.ListPendingFriendRequests(ctx, userID)
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

	return &pb.ListFriendRequestsResponse{
		FriendRequests: pbFriendRequests,
	}, nil
}
