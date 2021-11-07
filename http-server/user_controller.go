package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-users/grpc-service"
)

type UserController struct {
	client grpc_service.UsersServiceClient
}

func NewUserController(client grpc_service.UsersServiceClient) *UserController {
	return &UserController{client: client}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

}
