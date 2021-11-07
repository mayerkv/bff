package grpc_clients

import "google.golang.org/grpc"

func grpcDial(addr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	return grpc.Dial(addr, opts...)
}
