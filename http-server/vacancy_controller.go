package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/bff/grpc-clients"
	"github.com/mayerkv/go-recruitmens/grpc-service"
	"net/http"
	"time"
)

type VacancyController struct {
	client grpc_service.RecruitmentServiceClient
}

func NewVacancyController(client grpc_service.RecruitmentServiceClient) *VacancyController {
	return &VacancyController{client: client}
}

func (c *VacancyController) PostVacancy(ctx *gin.Context) {
	var dto PostVacancyDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.PostVacancyRequest{
		PositionId: dto.PositionId,
		CustomerId: dto.CustomerId,
	}

	response, err := c.client.PostVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": response.VacancyId,
	})
}

func (c *VacancyController) ShowVacancies(ctx *gin.Context) {
	var dto ShowVacanciesDto
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.ShowVacanciesRequest{
		CustomerId:     dto.CustomerId,
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapVacancyOrder(dto.OrderBy),
		OrderDirection: mapVacancyOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.ShowVacancies(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapVacancyDtoList(response.List),
		"count": response.Count,
	})
}

func (c *VacancyController) SearchVacancies(ctx *gin.Context) {
	var dto SearchVacanciesDto
	if err := ctx.ShouldBindQuery(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.SearchVacanciesRequest{
		Page:           int32(dto.Page),
		Size:           int32(dto.Size),
		OrderBy:        mapVacancyOrder(dto.OrderBy),
		OrderDirection: mapVacancyOrderDirection(dto.OrderDirection),
	}

	response, err := c.client.SearchVacancies(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":  mapVacancyDtoList(response.List),
		"count": response.Count,
	})
}

func (c *VacancyController) ChangeVacancyPosition(ctx *gin.Context) {
	var dto ChangeVacancyPositionDto
	ctx.ShouldBindUri(&dto)
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.ChangeVacancyPositionRequest{
		VacancyId:  dto.VacancyId,
		PositionId: dto.PositionId,
	}

	_, err := c.client.ChangeVacancyPosition(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *VacancyController) ApproveVacancy(ctx *gin.Context) {
	var dto ApproveVacancyDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.ApproveVacancyRequest{
		VacancyId: dto.VacancyId,
	}

	_, err := c.client.ApproveVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *VacancyController) CloseVacancy(ctx *gin.Context) {
	var dto CloseVacancyDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.CloseVacancyRequest{
		VacancyId: dto.VacancyId,
	}

	_, err := c.client.CloseVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *VacancyController) RejectVacancy(ctx *gin.Context) {
	var dto RejectVacancyDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.RejectVacancyRequest{
		VacancyId: dto.VacancyId,
	}

	_, err := c.client.RejectVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *VacancyController) TakeInWorkVacancy(ctx *gin.Context) {
	var dto TakeInWorkVacancyDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.TakeInWorkVacancyRequest{
		VacancyId: dto.VacancyId,
	}

	_, err := c.client.TakeInWorkVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *VacancyController) GetVacancy(ctx *gin.Context) {
	var dto GetVacancyDto
	if err := ctx.ShouldBindUri(&dto); err != nil {
		handleError(ctx, err)
		return
	}

	reqCtx, cancel := grpc_clients.ContextWithCancel(ctx.Request.Header, 3*time.Second)
	defer cancel()

	req := &grpc_service.GetVacancyRequest{
		VacancyId: dto.VacancyId,
	}

	response, err := c.client.GetVacancy(reqCtx, req)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, mapVacancyDto(response.Vacancy))
}
