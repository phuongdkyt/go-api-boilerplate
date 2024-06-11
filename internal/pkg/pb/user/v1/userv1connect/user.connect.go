// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/anhnmt/go-api-boilerplate/internal/pkg/pb/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceSayHelloProcedure is the fully-qualified name of the UserService's SayHello RPC.
	UserServiceSayHelloProcedure = "/user.v1.UserService/SayHello"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor        = v1.File_user_v1_user_proto.Services().ByName("UserService")
	userServiceSayHelloMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("SayHello")
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	SayHello(context.Context, *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloReply], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		sayHello: connect.NewClient[v1.HelloRequest, v1.HelloReply](
			httpClient,
			baseURL+UserServiceSayHelloProcedure,
			connect.WithSchema(userServiceSayHelloMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	sayHello *connect.Client[v1.HelloRequest, v1.HelloReply]
}

// SayHello calls user.v1.UserService.SayHello.
func (c *userServiceClient) SayHello(ctx context.Context, req *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloReply], error) {
	return c.sayHello.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	SayHello(context.Context, *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloReply], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceSayHelloHandler := connect.NewUnaryHandler(
		UserServiceSayHelloProcedure,
		svc.SayHello,
		connect.WithSchema(userServiceSayHelloMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/user.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceSayHelloProcedure:
			userServiceSayHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) SayHello(context.Context, *connect.Request[v1.HelloRequest]) (*connect.Response[v1.HelloReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.SayHello is not implemented"))
}
