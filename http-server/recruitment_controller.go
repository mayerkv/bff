package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-recruitmens/grpc-service"
)

type RecruitmentController struct {
	client grpc_service.RecruitmentServiceClient
}

func NewRecruitmentController(client grpc_service.RecruitmentServiceClient) *RecruitmentController {
	return &RecruitmentController{client: client}
}

func (c *RecruitmentController) ConsiderCandidate(ctx *gin.Context) {

}

func (c *RecruitmentController) ConsiderCandidateAnotherVacancy(ctx *gin.Context) {

}

func (c *RecruitmentController) AcceptRecruitmentStage(ctx *gin.Context) {

}

func (c *RecruitmentController) DenyRecruitment(ctx *gin.Context) {

}

func (c *RecruitmentController) GetRecruitment(ctx *gin.Context) {

}

func (c *RecruitmentController) ShowRecruitments(ctx *gin.Context) {

}


