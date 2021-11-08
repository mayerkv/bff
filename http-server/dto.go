package http_server

type CreateUserDto struct {
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
	Role     int    `json:"role,omitempty" binding:"numeric"`
}

type CreateCandidateDto struct {
	Name     string       `json:"name,omitempty" binding:"required"`
	Surname  string       `json:"surname,omitempty" binding:"required"`
	Contacts []ContactDto `json:"contacts,omitempty" binding:"required"`
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
	Page           int `json:"page,omitempty" form:"page" binding:"numeric"`
	Size           int `json:"size,omitempty" form:"size" binding:"numeric"`
	OrderBy        int `json:"orderBy,omitempty" form:"orderBy" binding:"numeric"`
	OrderDirection int `json:"orderDirection,omitempty" form:"orderDirection" binding:"numeric"`
}

type CreateCatalogDto struct {
	Id    string           `json:"id,omitempty"`
	Title string           `json:"title,omitempty" binding:"required"`
	Items []CatalogItemDto `json:"items,omitempty"`
}

type CatalogItemDto struct {
	Id    string `json:"id,omitempty"`
	Value string `json:"value,omitempty" binding:"required"`
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
	Name     string `json:"name,omitempty" binding:"required"`
	Template string `json:"template,omitempty" binding:"required"`
}

type SearchTemplatesDto struct {
	Name           string `form:"name"`
	Page           int    `form:"page" binding:"numeric"`
	Size           int    `form:"size" binding:"numeric"`
	OrderBy        int    `form:"orderBy" binding:"numeric"`
	OrderDirection int    `form:"orderDirection" binding:"numeric"`
}

type TemplateDto struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Template string `json:"template,omitempty"`
}

type SearchNotificationsDto struct {
	Page           int `form:"page" binding:"numeric"`
	Size           int `form:"size" binding:"numeric"`
	OrderBy        int `form:"orderBy" binding:"numeric"`
	OrderDirection int `form:"orderDirection" binding:"numeric"`
}

type NotificationDto struct {
	Id        string `json:"id,omitempty"`
	Addressee string `json:"addressee,omitempty"`
	Message   string `json:"message,omitempty"`
	Type      int    `json:"type,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type ConsiderCandidateDto struct {
	VacancyId     string       `json:"vacancyId,omitempty" binding:"required"`
	CandidateId   string       `json:"candidateId,omitempty" binding:"required"`
	ResponsibleId string       `json:"responsibleId,omitempty" binding:"required"`
	Settings      []SettingDto `json:"settings,omitempty" binding:"required"`
}

type SettingDto struct {
	StageId           string `json:"stageId,omitempty" binding:"required"`
	DeadlineDuration  int    `json:"deadlineDuration,omitempty" binding:"numeric"`
	ThresholdDuration int    `json:"thresholdDuration,omitempty" binding:"numeric"`
}

type ConsiderCandidateAnotherVacancyDto struct {
	VacancyId     string       `uri:"id" binding:"required"`
	RecruitmentId string       `json:"recruitmentId,omitempty" binding:"required"`
	Settings      []SettingDto `json:"settings,omitempty" binding:"required"`
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
	ReasonId string `json:"reasonId,omitempty" binding:"required"`
	Comment  string `json:"comment,omitempty" binding:"required"`
}

type GetRecruitmentDto struct {
	RecruitmentId string `uri:"id" binding:"required"`
}

type RecruitmentDto struct {
	Id            string          `json:"id,omitempty"`
	CandidateId   string          `json:"candidateId,omitempty"`
	ResponsibleId string          `json:"responsibleId,omitempty"`
	CreatedAt     string          `json:"createdAt,omitempty"`
	StageLine     StageLineDto    `json:"stageLine,omitempty"`
	Vacancy       VacancyDto      `json:"vacancy,omitempty"`
	RefuseReason  RefuseReasonDto `json:"refuseReason"`
}

type StageLineDto struct {
	StageId  string                      `json:"stageId,omitempty"`
	Settings []SettingDto                `json:"settings,omitempty"`
	History  map[string]StageLineItemDto `json:"history,omitempty"`
}

type StageLineItemDto struct {
	StageId       string `json:"stageId,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	FinishDate    string `json:"finishDate,omitempty"`
	DeadlineDate  string `json:"deadlineDate,omitempty"`
	ThresholdDate string `json:"thresholdDate,omitempty"`
}

type VacancyDto struct {
	Id         string `json:"id,omitempty"`
	PositionId string `json:"positionId,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	Status     int    `json:"status,omitempty"`
}

type RefuseReasonDto struct {
	ReasonId string `json:"reasonId,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

type ShowRecruitmentsDto struct {
	ResponsibleId  string `json:"responsibleId,omitempty"`
	Page           int    `json:"page,omitempty"`
	Size           int    `json:"size,omitempty"`
	OrderBy        int    `json:"orderBy,omitempty"`
	OrderDirection int    `json:"orderDirection,omitempty"`
}
