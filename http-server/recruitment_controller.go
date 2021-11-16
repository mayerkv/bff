package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_clients "github.com/mayerkv/bff/grpc-clients"
	"github.com/mayerkv/go-recruitmens/grpc-service"
	"net/http"
	"time"
)

type RecruitmentController struct {
	client grpc_service.RecruitmentServiceClient
}

func NewRecruitmentController(client grpc_service.RecruitmentServiceClient) *RecruitmentController {
	return &RecruitmentController{client: client}
}

func (c *RecruitmentController) ConsiderCandidate(ctx *gin.Context) {
	var dto ConsiderCandidateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.ConsiderCandidateRequest{
		VacancyId:     dto.VacancyId,
		CandidateId:   dto.CandidateId,
		ResponsibleId: dto.ResponsibleId,
		Settings:      mapSettings(dto.Settings),
	}
	response, err := c.client.ConsiderCandidate(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": response.RecruitmentId,
	})
}

func (c *RecruitmentController) ConsiderCandidateAnotherVacancy(ctx *gin.Context) {
	var dto ConsiderCandidateAnotherVacancyDto
	ctx.ShouldBindUri(&dto)
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &grpc_service.ConsiderCandidateAnotherVacancyRequest{
		RecruitmentId: dto.RecruitmentId,
		VacancyId:     dto.VacancyId,
		Settings:      mapSettings(dto.Settings),
	}

	response, err := c.client.ConsiderCandidateAnotherVacancy(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"recruitmentId": response.RecruitmentId,
	})
}

func (c *RecruitmentController) AcceptRecruitmentStage(ctx *gin.Context) {
	var dto AcceptRecruitmentStageDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.AcceptRecruitmentStageRequest{
		RecruitmentId:    dto.RecruitmentId,
		RequestedStageId: dto.RequestedStageId,
	}
	_, err := c.client.AcceptRecruitmentStage(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *RecruitmentController) DenyRecruitment(ctx *gin.Context) {
	var dto DenyRecruitmentDto
	ctx.ShouldBindUri(&dto)
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.DenyRecruitmentRequest{
		RecruitmentId: dto.RecruitmentId,
		Reason:        mapReason(dto.Reason),
	}
	_, err := c.client.DenyRecruitment(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *RecruitmentController) GetRecruitment(ctx *gin.Context) {
	var dto GetRecruitmentDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.GetRecruitmentRequest{
		RecruitmentId: dto.RecruitmentId,
	}

	response, err := c.client.GetRecruitment(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, mapRecruitment(response.Recruitment))
}

func (c *RecruitmentController) ShowRecruitments(ctx *gin.Context) {
	var dto ShowRecruitmentsDto
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &grpc_service.ShowRecruitmentRequest{
		ResponsibleId:  dto.ResponsibleId,
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapRecruitmentOrderBy(dto.OrderBy),
		OrderDirection: mapRecruitmentOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.ShowRecruitments(reqCtx, req, grpc_clients.Headers(ctx.Request))
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapRecruitments(response.List),
		"count": response.Count,
	})
}
