package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-candidates/grpc-service"
	"net/http"
	"time"
)

type CandidateController struct {
	client grpc_service.CandidatesServiceClient
}

func NewCandidateController(client grpc_service.CandidatesServiceClient) *CandidateController {
	return &CandidateController{client: client}
}

func (c *CandidateController) CreateCandidate(ctx *gin.Context) {
	var dto CreateCandidateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.CreateCandidateRequest{
		Name:     dto.Name,
		Surname:  dto.Surname,
		Contacts: mapDtoToContacts(dto.Contacts),
	}
	candidate, err := c.client.CreateCandidate(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": candidate.Id,
	})
}

func (c *CandidateController) GetCandidate(ctx *gin.Context) {
	var dto GetCandidateDto
	if err := ctx.BindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := &grpc_service.GetCandidateRequest{Id: dto.Id}
	response, err := c.client.GetCandidate(reqCtx, request)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, mapCandidateToDto(response.Candidate))
}

func (c *CandidateController) SearchCandidates(ctx *gin.Context) {
	var dto SearchCandidatesDto
	if err := ctx.ShouldBind(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.SearchCandidatesRequest{
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapCandidateOrder(dto.OrderBy),
		OrderDirection: mapOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.SearchCandidates(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapCandidatesToDto(response.List),
		"count": response.Count,
	})
}
