package http_server

import (
	candidates "github.com/mayerkv/go-candidates/grpc-service"
	catalogs "github.com/mayerkv/go-catalogs/grpc-service"
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
	res := make([]*catalogs.CatalogItem, 0)

	for _, i := range items {
		res = append(res, mapCatalogItemDto(i))
	}

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
