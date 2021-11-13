package http_server

import (
	"fmt"
	candidates "github.com/mayerkv/go-candidates/grpc-service"
	catalogs "github.com/mayerkv/go-catalogs/grpc-service"
	notifications "github.com/mayerkv/go-notifications/grpc-service"
	recruitments "github.com/mayerkv/go-recruitmens/grpc-service"
	users "github.com/mayerkv/go-users/grpc-service"
)

func mapCandidateToDto(candidate *candidates.Candidate) CandidateDto {
	return CandidateDto{
		Id:       candidate.Id,
		Name:     candidate.Name,
		Surname:  candidate.Surname,
		Contacts: mapContactsToDto(candidate.Contacts),
	}
}

func mapContactsToDto(contacts []*candidates.Contact) []ContactDto {
	res := make([]ContactDto, 0)

	for _, c := range contacts {
		res = append(res, mapContactToDto(c))
	}

	return res
}

func mapContactToDto(c *candidates.Contact) ContactDto {
	return ContactDto{
		Type:  int(c.Type),
		Value: c.Value,
	}
}

func mapDtoToContacts(contacts []ContactDto) []*candidates.Contact {
	res := make([]*candidates.Contact, 0)

	for _, c := range contacts {
		res = append(res, mapContact(c))
	}

	return res
}

func mapContact(c ContactDto) *candidates.Contact {
	return &candidates.Contact{
		Type:  intToContactType(c.Type),
		Value: c.Value,
	}
}

func intToContactType(i int) candidates.Contact_Type {
	var t candidates.Contact_Type

	switch i {
	case 0:
		t = candidates.Contact_PHONE
	case 1:
		t = candidates.Contact_EMAIL
	}

	return t
}

func mapCatalogDto(dto CreateCatalogDto) *catalogs.Catalog {
	return &catalogs.Catalog{
		Id:    dto.Id,
		Title: dto.Title,
		Items: mapCatalogItemsDto(dto.Items),
	}
}

func mapCatalogItemsDto(items []CatalogItemDto) []*catalogs.CatalogItem {
	fmt.Println(items)

	res := make([]*catalogs.CatalogItem, 0, len(items))

	for _, i := range items {
		res = append(res, mapCatalogItemDto(i))
	}


	fmt.Println(res)

	return res
}

func mapCatalogItemDto(i CatalogItemDto) *catalogs.CatalogItem {
	return &catalogs.CatalogItem{
		Id:    i.Id,
		Value: i.Value,
	}
}

func mapCatalogItems(items []*catalogs.CatalogItem) []CatalogItemDto {
	res := make([]CatalogItemDto, 0)

	for _, i := range items {
		res = append(res, CatalogItemDto{
			Id:    i.Id,
			Value: i.Value,
		})
	}

	return res
}

func mapTemplates(list []*notifications.Template) []TemplateDto {
	res := make([]TemplateDto, 0)

	for _, i := range list {
		res = append(res, TemplateDto{
			Id:       i.Id,
			Name:     i.Name,
			Template: i.Template,
		})
	}

	return res
}

func mapTemplatesOrderDirection(direction int) notifications.OrderDirection {
	if direction == 1 {
		return notifications.OrderDirection_DESC
	}

	return notifications.OrderDirection_ASC
}

func mapTemplatesOrderBy(by int) notifications.SearchTemplatesRequest_OrderBy {
	return notifications.SearchTemplatesRequest_NAME
}

func mapNotificationDtoList(list []*notifications.Notification) []NotificationDto {
	res := make([]NotificationDto, 0)

	for _, i := range list {
		res = append(res, NotificationDto{
			Id:        i.Id,
			Addressee: i.Addressee,
			Message:   i.Message,
			Type:      int(i.Type),
			CreatedAt: i.CreatedAt,
		})
	}

	return res
}

func mapNotificationsOrderDirection(direction int) notifications.OrderDirection {
	if direction == 1 {
		return notifications.OrderDirection_DESC
	}

	return notifications.OrderDirection_ASC
}

func mapNotificationsOrderBy(by int) notifications.SearchNotificationsRequest_OrderBy {
	return notifications.SearchNotificationsRequest_CREATED_AT
}

func mapSettings(settings []SettingDto) []*recruitments.StageLineSettings {
	res := make([]*recruitments.StageLineSettings, 0, len(settings))

	for _, i := range settings {
		res = append(res, &recruitments.StageLineSettings{
			StageId:              i.StageId,
			DeadlineDurationSec:  int32(i.DeadlineDuration),
			ThresholdDurationSec: int32(i.ThresholdDuration),
		})
	}

	return res
}

func mapReason(reason ReasonDto) *recruitments.RefuseReason {
	return &recruitments.RefuseReason{
		ReasonId: reason.ReasonId,
		Comment:  reason.Comment,
	}
}

func mapRecruitment(recruitment *recruitments.Recruitment) RecruitmentDto {
	return RecruitmentDto{
		Id:            recruitment.Id,
		CandidateId:   recruitment.CandidateId,
		ResponsibleId: recruitment.ResponsibleId,
		CreatedAt:     recruitment.CreatedAt,
		StageLine:     mapStageLineDto(recruitment.StageLine),
		Vacancy:       mapVacancyDto(recruitment.Vacancy),
		RefuseReason:  mapRefuseReasonDto(recruitment.RefuseReason),
	}
}

func mapStageLineDto(line *recruitments.StageLine) *StageLineDto {
	if line == nil {
		return nil
	}

	return &StageLineDto{
		StageId:  line.StageId,
		Settings: mapSettingsDto(line.Settings),
		History:  mapHistoryDto(line.History),
	}
}

func mapHistoryDto(history map[string]*recruitments.StageLineItem) map[string]StageLineItemDto {
	res := map[string]StageLineItemDto{}

	for id, item := range history {
		res[id] = mapStageLineItemDto(item)
	}

	return res
}

func mapStageLineItemDto(item *recruitments.StageLineItem) StageLineItemDto {
	return StageLineItemDto{
		StageId:       item.StageId,
		StartDate:     item.StartDate,
		FinishDate:    item.FinishDate,
		DeadlineDate:  item.DeadlineDate,
		ThresholdDate: item.ThresholdDate,
	}
}

func mapSettingsDto(settings []*recruitments.StageLineSettings) []SettingDto {
	res := make([]SettingDto, 0, len(settings))

	for _, s := range settings {
		res = append(res, SettingDto{
			StageId:           s.StageId,
			DeadlineDuration:  int(s.DeadlineDurationSec),
			ThresholdDuration: int(s.ThresholdDurationSec),
		})
	}

	return res
}

func mapVacancyDto(vacancy *recruitments.Vacancy) *VacancyDto {
	if vacancy == nil {
		return nil
	}

	return &VacancyDto{
		Id:         vacancy.Id,
		PositionId: vacancy.PositionId,
		CustomerId: vacancy.CustomerId,
		CreatedAt:  vacancy.CreatedAt,
		Status:     int(vacancy.Status),
	}
}

func mapRefuseReasonDto(reason *recruitments.RefuseReason) *RefuseReasonDto {
	if reason == nil {
		return nil
	}

	return &RefuseReasonDto{
		ReasonId: reason.ReasonId,
		Comment:  reason.Comment,
	}
}

func mapRecruitments(list []*recruitments.Recruitment) []RecruitmentDto {
	res := make([]RecruitmentDto, 0, len(list))

	for _, i := range list {
		res = append(res, mapRecruitment(i))
	}

	return res
}

func mapRecruitmentOrderDirection(direction int) recruitments.OrderDirection {
	if direction == 1 {
		return recruitments.OrderDirection_DESC
	}

	return recruitments.OrderDirection_ASC
}

func mapRecruitmentOrderBy(by int) recruitments.Recruitment_Order {
	return recruitments.Recruitment_CREATED_AT
}

func mapVacancyDtoList(list []*recruitments.Vacancy) []VacancyDto {
	res := make([]VacancyDto, 0, len(list))

	for _, i := range list {
		res = append(res, *mapVacancyDto(i))
	}

	return res
}

func mapVacancyOrderDirection(direction int) recruitments.OrderDirection {
	if direction == 1 {
		return recruitments.OrderDirection_DESC
	}

	return recruitments.OrderDirection_ASC
}

func mapVacancyOrder(by int) recruitments.Vacancy_Order {
	return recruitments.Vacancy_CREATED_AT
}

func mapCandidatesToDto(list []*candidates.Candidate) []CandidateDto {
	var res []CandidateDto

	for _, i := range list {
		res = append(res, mapCandidateToDto(i))
	}

	return res
}

func mapOrderDirection(direction int) candidates.OrderDirection {
	if direction == 1 {
		return candidates.OrderDirection_DESC
	}

	return candidates.OrderDirection_ASC
}

func mapCandidateOrder(by int) candidates.Candidate_Order {
	res := candidates.Candidate_FULL_NAME

	return res
}

func intToUserRole(i int) users.UserRole {
	var res users.UserRole

	switch i {
	case 1:
		res = users.UserRole_ROLE_ADMIN
	}

	return res
}
