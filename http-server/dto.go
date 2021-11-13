package http_server

type CreateUserDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"role" binding:"numeric"`
}

type CreateCandidateDto struct {
	Name     string       `json:"name" binding:"required"`
	Surname  string       `json:"surname" binding:"required"`
	Contacts []ContactDto `json:"contacts" binding:"required"`
}

type ContactDto struct {
	Type  int    `json:"type" binding:"numeric"`
	Value string `json:"value" binding:"required"`
}

type GetCandidateDto struct {
	Id string `json:"id" binding:"required" uri:"id"`
}

type CandidateDto struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Surname  string       `json:"surname"`
	Contacts []ContactDto `json:"contacts"`
}

type SearchCandidatesDto struct {
	Page           int `json:"page" form:"page" binding:"numeric"`
	Size           int `json:"size" form:"size" binding:"numeric"`
	OrderBy        int `json:"orderBy" form:"orderBy" binding:"numeric"`
	OrderDirection int `json:"orderDirection" form:"orderDirection" binding:"numeric"`
}

type CreateCatalogDto struct {
	Id    string           `json:"id"`
	Title string           `json:"title" binding:"required"`
	Items []CatalogItemDto `json:"items"`
}

type CatalogItemDto struct {
	Id    string `json:"id"`
	Value string `json:"value" binding:"required"`
}

type AddCatalogItemDto struct {
	CatalogId string `uri:"id" binding:"required"`
	Id        string `json:"id"`
	Value     string `json:"value" binding:"required"`
}

type GetCatalogItemsDto struct {
	CatalogId string `uri:"id" binding:"required"`
}

type CreateTemplateDto struct {
	Name     string `json:"name" binding:"required"`
	Template string `json:"template" binding:"required"`
}

type SearchTemplatesDto struct {
	Name           string `form:"name"`
	Page           int    `form:"page" binding:"numeric"`
	Size           int    `form:"size" binding:"numeric"`
	OrderBy        int    `form:"orderBy" binding:"numeric"`
	OrderDirection int    `form:"orderDirection" binding:"numeric"`
}

type TemplateDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Template string `json:"template"`
}

type SearchNotificationsDto struct {
	Page           int `form:"page" binding:"numeric"`
	Size           int `form:"size" binding:"numeric"`
	OrderBy        int `form:"orderBy" binding:"numeric"`
	OrderDirection int `form:"orderDirection" binding:"numeric"`
}

type NotificationDto struct {
	Id        string `json:"id"`
	Addressee string `json:"addressee"`
	Message   string `json:"message"`
	Type      int    `json:"type"`
	CreatedAt string `json:"createdAt"`
}

type ConsiderCandidateDto struct {
	VacancyId     string       `json:"vacancyId" binding:"required"`
	CandidateId   string       `json:"candidateId" binding:"required"`
	ResponsibleId string       `json:"responsibleId" binding:"required"`
	Settings      []SettingDto `json:"settings" binding:"required"`
}

type SettingDto struct {
	StageId           string `json:"stageId" binding:"required"`
	DeadlineDuration  int    `json:"deadlineDuration" binding:"numeric"`
	ThresholdDuration int    `json:"thresholdDuration" binding:"numeric"`
}

type ConsiderCandidateAnotherVacancyDto struct {
	VacancyId     string       `uri:"id" binding:"required"`
	RecruitmentId string       `json:"recruitmentId" binding:"required"`
	Settings      []SettingDto `json:"settings" binding:"required"`
}

type AcceptRecruitmentStageDto struct {
	RecruitmentId    string `uri:"id" binding:"required"`
	RequestedStageId string `uri:"stageId" binding:"required"`
}

type DenyRecruitmentDto struct {
	RecruitmentId string    `uri:"id" binding:"required"`
	Reason        ReasonDto `json:"reason" binding:"required"`
}

type ReasonDto struct {
	ReasonId string `json:"reasonId" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
}

type GetRecruitmentDto struct {
	RecruitmentId string `uri:"id" binding:"required"`
}

type RecruitmentDto struct {
	Id            string          `json:"id"`
	CandidateId   string          `json:"candidateId"`
	ResponsibleId string          `json:"responsibleId"`
	CreatedAt     string          `json:"createdAt"`
	StageLine     *StageLineDto    `json:"stageLine"`
	Vacancy       *VacancyDto      `json:"vacancy"`
	RefuseReason  *RefuseReasonDto `json:"refuseReason"`
}

type StageLineDto struct {
	StageId  string                      `json:"stageId"`
	Settings []SettingDto                `json:"settings"`
	History  map[string]StageLineItemDto `json:"history"`
}

type StageLineItemDto struct {
	StageId       string `json:"stageId"`
	StartDate     string `json:"startDate"`
	FinishDate    string `json:"finishDate"`
	DeadlineDate  string `json:"deadlineDate"`
	ThresholdDate string `json:"thresholdDate"`
}

type VacancyDto struct {
	Id         string `json:"id"`
	PositionId string `json:"positionId"`
	CustomerId string `json:"customerId"`
	CreatedAt  string `json:"createdAt"`
	Status     int    `json:"status"`
}

type RefuseReasonDto struct {
	ReasonId string `json:"reasonId"`
	Comment  string `json:"comment"`
}

type ShowRecruitmentsDto struct {
	ResponsibleId  string `form:"responsibleId" binding:"required"`
	Page           int    `form:"page" binding:"numeric"`
	Size           int    `form:"size" binding:"numeric"`
	OrderBy        int    `form:"orderBy" binding:"numeric"`
	OrderDirection int    `form:"orderDirection" binding:"numeric"`
}

type PostVacancyDto struct {
	PositionId string `json:"positionId" binding:"required"`
	CustomerId string `json:"customerId" binding:"required"`
}

type ShowVacanciesDto struct {
	CustomerId     string `form:"customerId" binding:"required"`
	Page           int    `form:"page" binding:"numeric"`
	Size           int    `form:"size" binding:"numeric"`
	OrderBy        int    `form:"orderBy" binding:"numeric"`
	OrderDirection int    `form:"orderDirection" binding:"numeric"`
}

type SearchVacanciesDto struct {
	Page           int `form:"page" binding:"numeric"`
	Size           int `form:"size" binding:"numeric"`
	OrderBy        int `form:"orderBy" binding:"numeric"`
	OrderDirection int `form:"orderDirection" binding:"numeric"`
}

type ChangeVacancyPositionDto struct {
	VacancyId  string `uri:"id" binding:"required"`
	PositionId string `json:"positionId" binding:"required"`
}

type ApproveVacancyDto struct {
	VacancyId string `uri:"id" binding:"required"`
}

type CloseVacancyDto struct {
	VacancyId string `uri:"id" binding:"required"`
}

type RejectVacancyDto struct {
	VacancyId string `uri:"id" binding:"required"`
}

type TakeInWorkVacancyDto struct {
	VacancyId string `uri:"id" binding:"required"`
}

type GetVacancyDto struct {
	VacancyId string `uri:"id" binding:"required"`
}
