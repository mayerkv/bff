package grpc_clients

import (
	candidates "github.com/mayerkv/go-candidates/grpc-service"
	catalogs "github.com/mayerkv/go-catalogs/grpc-service"
	notifications "github.com/mayerkv/go-notifications/grpc-service"
	recruitments "github.com/mayerkv/go-recruitmens/grpc-service"
	users "github.com/mayerkv/go-users/grpc-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func CreateUsersClient(addr string) (users.UsersServiceClient, *grpc.ClientConn, error) {
	conn, err := grpcDial(addr)
	if err != nil {
		return nil, nil, err
	}

	return users.NewUsersServiceClient(conn), conn, nil
}

func CreateCandidatesClient(addr string) (candidates.CandidatesServiceClient, *grpc.ClientConn, error) {
	conn, err := grpcDial(addr)
	if err != nil {
		return nil, nil, err
	}

	return candidates.NewCandidatesServiceClient(conn), conn, nil
}

func CreateCatalogsClient(addr string) (catalogs.CatalogsServiceClient, *grpc.ClientConn, error) {
	conn, err := grpcDial(addr)
	if err != nil {
		return nil, nil, err
	}

	return catalogs.NewCatalogsServiceClient(conn), conn, nil
}

func CreateNotificationsClient(addr string) (notifications.NotificationsServiceClient, *grpc.ClientConn, error) {
	conn, err := grpcDial(addr)
	if err != nil {
		return nil, nil, err
	}

	return notifications.NewNotificationsServiceClient(conn), conn, nil
}

func CreateRecruitmentsClient(addr string) (recruitments.RecruitmentServiceClient, *grpc.ClientConn, error) {
	conn, err := grpcDial(addr)
	if err != nil {
		return nil, nil, err
	}

	return recruitments.NewRecruitmentServiceClient(conn), conn, nil
}

func MetaDataFromHeaders(header http.Header) *metadata.MD {
	md := metadata.MD{}

	for name, values := range header {
		md[name] = values
	}

	return &md
}

func Headers(r *http.Request) grpc.CallOption {
	return grpc.Header(MetaDataFromHeaders(r.Header))
}