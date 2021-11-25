package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func CreateRouter(
	vacancyController *VacancyController,
	recruitmentController *RecruitmentController,
	userController *UserController,
	candidateController *CandidateController,
	catalogController *CatalogController,
	notificationController *NotificationController,
) *gin.Engine {
	responseTimeHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "response_time_bucket",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.DefBuckets,
	}, []string{"service", "method", "code"})

	requestCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "Request counts.",
	}, []string{"service", "method", "code"})

	prometheus.MustRegister(requestCount, responseTimeHistogram)

	r := gin.New()
	r.Use(gin.Recovery())

	vacanciesGroup := r.Group("/vacancies", gin.Logger())
	{
		vacanciesGroup.POST("", metricsMiddleware("vacancies", "PostVacancy", responseTimeHistogram, requestCount), vacancyController.PostVacancy)
		vacanciesGroup.GET("", metricsMiddleware("vacancies", "ShowVacancies", responseTimeHistogram, requestCount), vacancyController.ShowVacancies)
		vacanciesGroup.GET("/my", metricsMiddleware("vacancies", "SearchVacancies", responseTimeHistogram, requestCount), vacancyController.SearchVacancies)
		vacanciesGroup.PATCH("/:id/position", metricsMiddleware("vacancies", "ChangeVacancyPosition", responseTimeHistogram, requestCount), vacancyController.ChangeVacancyPosition)
		vacanciesGroup.POST("/:id/approve", metricsMiddleware("vacancies", "ApproveVacancy", responseTimeHistogram, requestCount), vacancyController.ApproveVacancy)
		vacanciesGroup.POST("/:id/close", metricsMiddleware("vacancies", "CloseVacancy", responseTimeHistogram, requestCount), vacancyController.CloseVacancy)
		vacanciesGroup.POST("/:id/reject", metricsMiddleware("vacancies", "RejectVacancy", responseTimeHistogram, requestCount), vacancyController.RejectVacancy)
		vacanciesGroup.POST("/:id/take_in_work", metricsMiddleware("vacancies", "TakeInWorkVacancy", responseTimeHistogram, requestCount), vacancyController.TakeInWorkVacancy)
		vacanciesGroup.GET("/:id", metricsMiddleware("vacancies", "GetVacancy", responseTimeHistogram, requestCount), vacancyController.GetVacancy)
	}

	recruitmentGroup := r.Group("/recruitment", gin.Logger())
	{
		recruitmentGroup.POST("", metricsMiddleware("recruitment", "ConsiderCandidate", responseTimeHistogram, requestCount), recruitmentController.ConsiderCandidate)
		recruitmentGroup.GET("/:id", metricsMiddleware("recruitment", "GetRecruitment", responseTimeHistogram, requestCount), recruitmentController.GetRecruitment)
		recruitmentGroup.GET("", metricsMiddleware("recruitment", "ShowRecruitments", responseTimeHistogram, requestCount), recruitmentController.ShowRecruitments)
		recruitmentGroup.POST("/vacancy/:id", metricsMiddleware("recruitment", "ConsiderCandidateAnotherVacancy", responseTimeHistogram, requestCount), recruitmentController.ConsiderCandidateAnotherVacancy)
		recruitmentGroup.POST("/:id/stage/:stageId/accept", metricsMiddleware("recruitment", "AcceptRecruitmentStage", responseTimeHistogram, requestCount), recruitmentController.AcceptRecruitmentStage)
		recruitmentGroup.POST("/:id/status/deny", metricsMiddleware("recruitment", "DenyRecruitment", responseTimeHistogram, requestCount), recruitmentController.DenyRecruitment)
	}

	usersGroup := r.Group("/users", gin.Logger())
	{
		usersGroup.POST("", metricsMiddleware("users", "CreateUser", responseTimeHistogram, requestCount), userController.CreateUser)
	}

	candidatesGroup := r.Group("/candidates", gin.Logger())
	{
		candidatesGroup.POST("", metricsMiddleware("candidates", "CreateCandidate", responseTimeHistogram, requestCount), candidateController.CreateCandidate)
		candidatesGroup.GET("/:id", metricsMiddleware("candidates", "GetCandidate", responseTimeHistogram, requestCount), candidateController.GetCandidate)
		candidatesGroup.GET("", metricsMiddleware("candidates", "SearchCandidates", responseTimeHistogram, requestCount), candidateController.SearchCandidates)
	}

	catalogsGroup := r.Group("/catalogs", gin.Logger())
	{
		catalogsGroup.POST("", metricsMiddleware("catalogs", "CreateCatalog", responseTimeHistogram, requestCount), catalogController.CreateCatalog)
		catalogsGroup.POST("/:id/items", metricsMiddleware("catalogs", "AddCatalogItem", responseTimeHistogram, requestCount), catalogController.AddCatalogItem)
		catalogsGroup.GET("/:id/items", metricsMiddleware("catalogs", "GetCatalogItems", responseTimeHistogram, requestCount), catalogController.GetCatalogItems)
	}

	notificationsGroup := r.Group("/notifications", gin.Logger())
	{
		notificationsGroup.POST("/templates", metricsMiddleware("notifications", "CreateTemplate", responseTimeHistogram, requestCount), notificationController.CreateTemplate)
		notificationsGroup.GET("/templates", metricsMiddleware("notifications", "SearchTemplates", responseTimeHistogram, requestCount), notificationController.SearchTemplates)
		notificationsGroup.GET("", metricsMiddleware("notifications", "SearchNotifications", responseTimeHistogram, requestCount), notificationController.SearchNotifications)
	}

	handler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})

	r.GET("/metrics", func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	return r
}
