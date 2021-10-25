package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mayerkv/go-recruitmens/recruitment-service"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	recruitmentsAddr := os.Getenv("RECRUITMENTS_ADDRESS")
	if recruitmentsAddr == "" {
		recruitmentsAddr = "localhost:9090"
	}

	conn, err := grpc.Dial(recruitmentsAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := recruitment_service.NewRecruitmentServiceClient(conn)

	r := gin.Default()
	r.POST("/vacancies", postVacancy(client))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func postVacancy(client recruitment_service.RecruitmentServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &recruitment_service.PostVacancyRequest{Message: os.Getenv("HOSTNAME")}

		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := client.PostVacancy(c, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": res.Message,
		})
	}
}
