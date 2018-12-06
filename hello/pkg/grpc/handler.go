package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "hello/hello/pkg/endpoint"
	pb "hello/hello/pkg/grpc/pb"
)

// makeSayHelloHandler creates the handler logic
func makeSayHelloHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SayHelloEndpoint, decodeSayHelloRequest, encodeSayHelloResponse, options...)
}

// decodeSayHelloResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeSayHelloRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Hello' Decoder is not impelemented")
}

// encodeSayHelloResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSayHelloResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Hello' Encoder is not impelemented")
}
func (g *grpcServer) SayHello(ctx context1.Context, req *pb.SayHelloRequest) (*pb.SayHelloReply, error) {
	_, rep, err := g.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SayHelloReply), nil
}
