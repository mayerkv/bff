package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-recruitmens/grpc-service"
)

type VacancyController struct {
	client grpc_service.RecruitmentServiceClient
}

func NewVacancyController(client grpc_service.RecruitmentServiceClient) *VacancyController {
	return &VacancyController{client: client}
}

func (c *VacancyController) PostVacancy(ctx *gin.Context) {

}

func (c *VacancyController) ShowVacancies(ctx *gin.Context) {

}

func (c *VacancyController) SearchVacancies(ctx *gin.Context) {

}

func (c *VacancyController) ChangeVacancyPosition(ctx *gin.Context) {

}

func (c *VacancyController) ApproveVacancy(ctx *gin.Context) {

}

func (c *VacancyController) CloseVacancy(ctx *gin.Context) {

}

func (c *VacancyController) RejectVacancy(ctx *gin.Context) {

}

func (c *VacancyController) TakeInWorkVacancy(ctx *gin.Context) {

}

func (c *VacancyController) GetVacancy(ctx *gin.Context) {

}
