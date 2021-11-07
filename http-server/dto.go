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
	OrderBy        int `json:"order_by,omitempty" form:"order_by" binding:"numeric"`
	OrderDirection int `json:"order_direction,omitempty" form:"order_direction" binding:"numeric"`
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