package http_server

import "github.com/gin-gonic/gin"

func CreateRouter(
	vacancyController *VacancyController,
	recruitmentController *RecruitmentController,
	userController *UserController,
	candidateController *CandidateController,
	catalogController *CatalogController,
	notificationController *NotificationController,
) *gin.Engine {
	r := gin.Default()

	vacanciesGroup := r.Group("/vacancies")
	{
		vacanciesGroup.POST("", vacancyController.PostVacancy)
		vacanciesGroup.GET("", vacancyController.ShowVacancies)
		vacanciesGroup.GET("/my", vacancyController.SearchVacancies)
		vacanciesGroup.PATCH("/:id/position", vacancyController.ChangeVacancyPosition)
		vacanciesGroup.POST("/:id/approve", vacancyController.ApproveVacancy)
		vacanciesGroup.POST("/:id/close", vacancyController.CloseVacancy)
		vacanciesGroup.POST("/:id/reject", vacancyController.RejectVacancy)
		vacanciesGroup.POST("/:id/take_in_work", vacancyController.TakeInWorkVacancy)
		vacanciesGroup.GET("/:id", vacancyController.GetVacancy)
	}

	recruitmentGroup := r.Group("/recruitment")
	{
		recruitmentGroup.POST("", recruitmentController.ConsiderCandidate)
		recruitmentGroup.GET("/:id", recruitmentController.GetRecruitment)
		recruitmentGroup.GET("", recruitmentController.ShowRecruitments)
		recruitmentGroup.POST("/vacancy/:id", recruitmentController.ConsiderCandidateAnotherVacancy)
		recruitmentGroup.POST("/:id/stage/:stageId/accept", recruitmentController.AcceptRecruitmentStage)
		recruitmentGroup.POST("/:id/status/deny", recruitmentController.DenyRecruitment)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("", userController.CreateUser)
	}

	candidatesGroup := r.Group("/candidates")
	{
		candidatesGroup.POST("", candidateController.CreateCandidate)
		candidatesGroup.GET("/:id", candidateController.GetCandidate)
		candidatesGroup.GET("", candidateController.SearchCandidates)
	}

	catalogsGroup := r.Group("/catalogs")
	{
		catalogsGroup.POST("", catalogController.CreateCatalog)
		catalogsGroup.POST("/:id/items", catalogController.AddCatalogItem)
		catalogsGroup.GET("/:id/items", catalogController.GetCatalogItems)
	}

	notificationsGroup := r.Group("/notifications")
	{
		notificationsGroup.POST("/templates", notificationController.CreateTemplate)
		notificationsGroup.GET("/templates", notificationController.SearchTemplates)
		notificationsGroup.GET("", notificationController.SearchNotifications)
	}

	return r
}
