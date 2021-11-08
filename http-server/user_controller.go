package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-users/grpc-service"
	"net/http"
	"time"
)

type UserController struct {
	client grpc_service.UsersServiceClient
}

func NewUserController(client grpc_service.UsersServiceClient) *UserController {
	return &UserController{client: client}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var dto CreateUserDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	req := &grpc_service.CreateUserRequest{
		Email:    dto.Email,
		Password: dto.Password,
		Role:     intToUserRole(dto.Role),
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.client.CreateUser(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func intToUserRole(i int) grpc_service.UserRole {
	var res grpc_service.UserRole

	switch i {
	case 1:
		res = grpc_service.UserRole_ROLE_ADMIN
	}

	return res
}
