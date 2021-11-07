package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-catalogs/grpc-service"
)

type CatalogController struct {
	client grpc_service.CatalogsServiceClient
}

func NewCatalogController(client grpc_service.CatalogsServiceClient) *CatalogController {
	return &CatalogController{client: client}
}

func (c *CatalogController) CreateCatalog(ctx *gin.Context) {

}

func (c *CatalogController) AddCatalogItem(ctx *gin.Context) {

}

func (c *CatalogController) GetCatalogItems(ctx *gin.Context) {

}
