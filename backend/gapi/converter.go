package gapi

import (
	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertRegisterUser(user db.RegisterTxResult) *pb.User {
	return &pb.User{
		Id:                user.User.ID,
		Username:          user.User.Username,
		Email:             user.User.Email,
		PasswordChangedAt: timestamppb.New(user.User.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.User.CreatedAt),
	}
}

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
