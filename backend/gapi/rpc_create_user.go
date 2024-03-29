package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/PlatosCodes/desserted/backend/util"
	"github.com/PlatosCodes/desserted/backend/val"

	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {

		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		Password: hashedPassword,
		Email:    req.GetEmail(),
	}

	user, err := server.Store.RegisterTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	activationToken, activationPayload, err := server.tokenMaker.CreateToken(
		user.User.ID,
		user.User.Username,
		server.config.RefreshTokenDuration*7,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create activation token: %s", err)
	}

	activationInfo, err := server.Store.InsertActivationToken(ctx, db.InsertActivationTokenParams{
		UserID:          user.User.ID,
		ActivationToken: activationToken,
		ExpiresAt:       activationPayload.ExpiresAt.Time,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert activation token into the db: %s", err)
	}

	dbs := map[string]interface{}{
		"activationToken": activationInfo.ActivationToken,
		"userID":          activationPayload.UserID,
		"username":        user.User.Username,
	}

	// make async later
	err = server.mailer.Send(user.User.Email, "user_welcome.tmpl", dbs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to send activation email to user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertRegisterUser(user),
	}
	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
