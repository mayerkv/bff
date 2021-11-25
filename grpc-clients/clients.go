package grpc_clients

import (
	"context"
	candidates "github.com/mayerkv/go-candidates/grpc-service"
	catalogs "github.com/mayerkv/go-catalogs/grpc-service"
	notifications "github.com/mayerkv/go-notifications/grpc-service"
	recruitments "github.com/mayerkv/go-recruitmens/grpc-service"
	users "github.com/mayerkv/go-users/grpc-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"time"
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

const (
	prefixTracerState  = "x-b3-"
	zipkinTraceID      = prefixTracerState + "traceid"
	zipkinSpanID       = prefixTracerState + "spanid"
	zipkinParentSpanID = prefixTracerState + "parentspanid"
	zipkinSampled      = prefixTracerState + "sampled"
	zipkinFlags        = prefixTracerState + "flags"
)

var otHeaders = []string{
	zipkinTraceID,
	zipkinSpanID,
	zipkinParentSpanID,
	zipkinSampled,
	zipkinFlags,
}

func MetaDataFromHeaders(header http.Header) metadata.MD {
	pairs := make([]string, 0, len(otHeaders))
	for _, h := range otHeaders {
		if v := header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

func ContextWithCancel(header http.Header, t time.Duration) (context.Context, context.CancelFunc) {
	md := MetaDataFromHeaders(header)
	reqCtx := metadata.NewOutgoingContext(context.Background(), md)

	log.Println("meta:", md)

	return context.WithTimeout(reqCtx, t)
}