package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-notifications/grpc-service"
)

type NotificationController struct {
	client grpc_service.NotificationsServiceClient
}

func NewNotificationController(client grpc_service.NotificationsServiceClient) *NotificationController {
	return &NotificationController{client: client}
}

func (c *NotificationController) CreateTemplate(ctx *gin.Context) {
	
}

func (c *NotificationController) SearchTemplates(ctx *gin.Context) {
	
}

func (c *NotificationController) SearchNotifications(ctx *gin.Context) {
	
}
