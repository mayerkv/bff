package main

import (
	"github.com/gin-gonic/gin"
	grpc_clients "github.com/mayerkv/bff/grpc-clients"
	http_server "github.com/mayerkv/bff/http-server"
	"log"
	"net/http"
	"time"
)

func main() {
	usersClient, usersConn, err := grpc_clients.CreateUsersClient("users:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer usersConn.Close()

	recruitmentsClient, recruitmentsConn, err := grpc_clients.CreateRecruitmentsClient("recruitments:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer recruitmentsConn.Close()

	candidatesClient, candidatesConn, err := grpc_clients.CreateCandidatesClient("candidates:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer candidatesConn.Close()

	catalogsClient, catalogsConn, err := grpc_clients.CreateCatalogsClient("catalogs:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer catalogsConn.Close()

	notificationsClient, notificationsConn, err := grpc_clients.CreateNotificationsClient("notifications:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer notificationsConn.Close()

	vacancyController := http_server.NewVacancyController(recruitmentsClient)
	recruitmentController := http_server.NewRecruitmentController(recruitmentsClient)
	userController := http_server.NewUserController(usersClient)
	candidateController := http_server.NewCandidateController(candidatesClient)
	catalogController := http_server.NewCatalogController(catalogsClient)
	notificationController := http_server.NewNotificationController(notificationsClient)

	r := http_server.CreateRouter(vacancyController, recruitmentController, userController, candidateController, catalogController, notificationController)

	if err := runHttpServer(r); err != nil {
		log.Fatal(err)
	}
}

func runHttpServer(r *gin.Engine) error {
	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return server.ListenAndServe()
}
