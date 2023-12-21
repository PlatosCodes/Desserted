package gapi

import (
	"context"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) InvitePlayersToGame(ctx context.Context, req *pb.InvitePlayersToGameRequest) (*emptypb.Empty, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	log.Println("yooohooo:", req)

	inviter := req.GetInviterPlayerId()

	if authPayload.UserID != inviter {
		return nil, status.Errorf(codes.PermissionDenied, "you are not authorized to add a game for this user")
	}

	creator, err := server.Store.IsUserGameCreator(ctx, db.IsUserGameCreatorParams{
		CreatedBy: inviter,
		GameID:    req.GetGameId(),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, status.Errorf(codes.Internal, "database error in checking game creator: %s, code: %s", err, pqErr.Code)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	if !creator {
		return nil, status.Errorf(codes.PermissionDenied, "you are not authorized to invite players to this game")
	}

	invitees := req.GetInviteeUsernames()

	log.Println(invitees)

	for _, username := range invitees {
		arg := db.CreateGameInvitationWithUsernameParams{
			InviterPlayerID: inviter,
			Username:        username,
			GameID:          req.GetGameId(),
		}
		if err := server.Store.CreateGameInvitationWithUsername(ctx, arg); err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					return nil, status.Errorf(codes.AlreadyExists, "player has already been invited to game: %s", err)
				default:
					return nil, status.Errorf(codes.Internal, "database error in inviting player to game: %s", err)
				}
			}
			return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
		}
	}

	return &emptypb.Empty{}, nil

}
