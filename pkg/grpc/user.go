package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/ngtrdai197/cobra-cmd/db/sqlc"
	"github.com/ngtrdai197/cobra-cmd/pkg/grpc/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {

	arg := db.CreateUserParams{
		Username:    req.GetUsername(),
		FullName:    req.GetFullName(),
		PhoneNumber: req.GetPhoneNumber(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	r := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return r, nil
}

func (server *Server) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {

	log.Debug().Msgf("request %s", req.GetUsername())
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, status.Errorf(codes.Internal, "failed to get user with error=%s", pqErr)
		}
		return nil, status.Errorf(codes.NotFound, "cannot found user with error=%s", err)
	}

	r := &pb.GetUserResponse{
		User: convertUser(user),
	}
	return r, nil
}
