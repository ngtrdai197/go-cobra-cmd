package gapi

import (
	db "github.com/ngtrdai197/cobra-cmd/db/sqlc"
	"github.com/ngtrdai197/cobra-cmd/pkg/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:    user.Username,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
}
