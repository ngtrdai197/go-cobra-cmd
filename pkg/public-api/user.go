package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/ngtrdai197/cobra-cmd/db/sqlc"
	"github.com/rs/zerolog/log"
)

type createUserRequest struct {
	Username    string `json:"username"  binding:"required,alphanum"`
	FullName    string `json:"full_name" binding:"required"`
	PhoneNumber string `json:"phone_number"     binding:"required"`
}

type getUserRequest struct {
	Username string `form:"username"  binding:"required,alphanum"`
}

type userResponse struct {
	Username    string    `json:"username"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:    user.Username,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}
}

func (s *PublicApiServer) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.CreateUserParams{
		Username:    req.Username,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
	}
	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		if pgError, ok := err.(*pq.Error); ok {
			log.Printf("PG Error: %v", pgError.Code.Name())
			switch pgError.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

func (s *PublicApiServer) GetUser(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	user, err := s.store.GetUser(ctx, req.Username)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusInternalServerError, errorResponse(pqErr))
		}
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}
