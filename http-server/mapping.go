package http_server

import (
	grpc_service "github.com/mayerkv/go-candidates/grpc-service"
)

func mapCandidateToDto(candidate *grpc_service.Candidate) CandidateDto {
	return CandidateDto{
		Id:       candidate.Id,
		Name:     candidate.Name,
		Surname:  candidate.Surname,
		Contacts: mapContactsToDto(candidate.Contacts),
	}
}

func mapContactsToDto(contacts []*grpc_service.Contact) []ContactDto {
	var res []ContactDto

	for _, c := range contacts {
		res = append(res, mapContactToDto(c))
	}

	return res
}

func mapContactToDto(c *grpc_service.Contact) ContactDto {
	return ContactDto{
		Type:  int(c.Type),
		Value: c.Value,
	}
}

func mapDtoToContacts(contacts []ContactDto) []*grpc_service.Contact {
	var res []*grpc_service.Contact

	for _, c := range contacts {
		res = append(res, mapContact(c))
	}

	return res
}

func mapContact(c ContactDto) *grpc_service.Contact {
	return &grpc_service.Contact{
		Type:  intToContactType(c.Type),
		Value: c.Value,
	}
}

func intToContactType(i int) grpc_service.Contact_Type {
	var t grpc_service.Contact_Type

	switch i {
	case 0:
		t = grpc_service.Contact_PHONE
	case 1:
		t = grpc_service.Contact_EMAIL
	}

	return t
}
