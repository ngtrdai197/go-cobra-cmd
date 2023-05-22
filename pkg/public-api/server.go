package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ngtrdai197/cobra-cmd/config"
	db "github.com/ngtrdai197/cobra-cmd/db/sqlc"
	"github.com/rs/zerolog/log"
)

type PublicApiServer struct {
	config *config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(c *config.Config) *PublicApiServer {
	conn, err := sql.Open(c.DbDriver, c.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect db")
	}
	store := db.NewStore(conn)
	server := &PublicApiServer{config: c, store: store}
	server.setupRouter(c)
	if err := server.Start(c.PublicApiAddress); err != nil {
		log.Fatal().Err(err).Msg("cannot create public api")
	}
	return server
}

func (s *PublicApiServer) setupRouter(c *config.Config) {
	r := gin.Default()
	if c.AppEnv != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user", s.CreateUser)

	s.router = r
}

func (s *PublicApiServer) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
