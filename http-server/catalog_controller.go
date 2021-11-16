package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_clients "github.com/mayerkv/bff/grpc-clients"
	"github.com/mayerkv/go-catalogs/grpc-service"
	"net/http"
	"time"
)

type CatalogController struct {
	client grpc_service.CatalogsServiceClient
}

func NewCatalogController(client grpc_service.CatalogsServiceClient) *CatalogController {
	return &CatalogController{client: client}
}

func (c *CatalogController) CreateCatalog(ctx *gin.Context) {
	var dto CreateCatalogDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.CreateCatalogRequest{Catalog: mapCatalogDto(dto)}

	response, err := c.client.CreateCatalog(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": response.Id,
	})
}

func (c *CatalogController) AddCatalogItem(ctx *gin.Context) {
	var dto AddCatalogItemDto
	ctx.ShouldBindUri(&dto)

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.AddCatalogItemRequest{
		CatalogId: dto.CatalogId,
		Item: &grpc_service.CatalogItem{
			Id:    dto.Id,
			Value: dto.Value,
		},
	}

	_, err := c.client.AddCatalogItem(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c *CatalogController) GetCatalogItems(ctx *gin.Context) {
	var dto GetCatalogItemsDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.GetCatalogItemsRequest{
		CatalogId: dto.CatalogId,
	}

	response, err := c.client.GetCatalogItems(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, mapCatalogItems(response.Items))
}
