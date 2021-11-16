package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_clients "github.com/mayerkv/bff/grpc-clients"
	"github.com/mayerkv/go-notifications/grpc-service"
	"net/http"
	"time"
)

type NotificationController struct {
	client grpc_service.NotificationsServiceClient
}

func NewNotificationController(client grpc_service.NotificationsServiceClient) *NotificationController {
	return &NotificationController{client: client}
}

func (c *NotificationController) CreateTemplate(ctx *gin.Context) {
	var dto CreateTemplateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.CreateTemplateRequest{
		Name:     dto.Name,
		Template: dto.Template,
	}

	response, err := c.client.CreateTemplate(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": response.Id,
	})
}

func (c *NotificationController) SearchTemplates(ctx *gin.Context) {
	var dto SearchTemplatesDto
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.SearchTemplatesRequest{
		Name:           dto.Name,
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapTemplatesOrderBy(dto.OrderBy),
		OrderDirection: mapTemplatesOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.SearchTemplates(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapTemplates(response.List),
		"count": response.Count,
	})
}

func (c *NotificationController) SearchNotifications(ctx *gin.Context) {
	var dto SearchNotificationsDto
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.SearchNotificationsRequest{
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapNotificationsOrderBy(dto.OrderBy),
		OrderDirection: mapNotificationsOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.SearchNotifications(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapNotificationDtoList(response.List),
		"count": response.Count,
	})
}
