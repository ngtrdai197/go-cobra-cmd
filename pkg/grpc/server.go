package gapi

import (
	"database/sql"
	"net"

	_ "github.com/lib/pq"
	"github.com/ngtrdai197/cobra-cmd/config"
	db "github.com/ngtrdai197/cobra-cmd/db/sqlc"
	"github.com/ngtrdai197/cobra-cmd/pkg/grpc/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config *config.Config
	store  db.Store
}

// NewServer creates a new gRPC server.
func NewServer(c *config.Config) *Server {
	conn, err := sql.Open(c.DbDriver, c.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect db")
	}
	store := db.NewStore(conn)

	server := &Server{
		config: c,
		store:  store,
	}

	return server
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.config.GrpcAddress)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}
