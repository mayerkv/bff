package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-candidates/grpc-service"
)

type CandidateController struct {
	client grpc_service.CandidatesServiceClient
}

func NewCandidateController(client grpc_service.CandidatesServiceClient) *CandidateController {
	return &CandidateController{client: client}
}

func (c *CandidateController) CreateCandidate(ctx *gin.Context) {

}

func (c *CandidateController) GetCandidate(ctx *gin.Context) {

}

func (c *CandidateController) SearchCandidates(ctx *gin.Context) {

}
