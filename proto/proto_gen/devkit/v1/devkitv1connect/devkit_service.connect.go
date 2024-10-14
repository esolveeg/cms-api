// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: devkit/v1/devkit_service.proto

package devkitv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/darwishdev/devkit-api/proto_gen/devkit/v1"
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
	// DevkitServiceName is the fully-qualified name of the DevkitService service.
	DevkitServiceName = "devkit.v1.DevkitService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DevkitServiceHelloWorldProcedure is the fully-qualified name of the DevkitService's HelloWorld
	// RPC.
	DevkitServiceHelloWorldProcedure = "/devkit.v1.DevkitService/HelloWorld"
	// DevkitServiceRoleCreateProcedure is the fully-qualified name of the DevkitService's RoleCreate
	// RPC.
	DevkitServiceRoleCreateProcedure = "/devkit.v1.DevkitService/RoleCreate"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	devkitServiceServiceDescriptor          = v1.File_devkit_v1_devkit_service_proto.Services().ByName("DevkitService")
	devkitServiceHelloWorldMethodDescriptor = devkitServiceServiceDescriptor.Methods().ByName("HelloWorld")
	devkitServiceRoleCreateMethodDescriptor = devkitServiceServiceDescriptor.Methods().ByName("RoleCreate")
)

// DevkitServiceClient is a client for the devkit.v1.DevkitService service.
type DevkitServiceClient interface {
	HelloWorld(context.Context, *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error)
	RoleCreate(context.Context, *connect.Request[v1.RoleCreateRequest]) (*connect.Response[v1.RoleCreateResponse], error)
}

// NewDevkitServiceClient constructs a client for the devkit.v1.DevkitService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDevkitServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DevkitServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &devkitServiceClient{
		helloWorld: connect.NewClient[v1.HelloWorldRequest, v1.HelloWorldResponse](
			httpClient,
			baseURL+DevkitServiceHelloWorldProcedure,
			connect.WithSchema(devkitServiceHelloWorldMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		roleCreate: connect.NewClient[v1.RoleCreateRequest, v1.RoleCreateResponse](
			httpClient,
			baseURL+DevkitServiceRoleCreateProcedure,
			connect.WithSchema(devkitServiceRoleCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// devkitServiceClient implements DevkitServiceClient.
type devkitServiceClient struct {
	helloWorld *connect.Client[v1.HelloWorldRequest, v1.HelloWorldResponse]
	roleCreate *connect.Client[v1.RoleCreateRequest, v1.RoleCreateResponse]
}

// HelloWorld calls devkit.v1.DevkitService.HelloWorld.
func (c *devkitServiceClient) HelloWorld(ctx context.Context, req *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
	return c.helloWorld.CallUnary(ctx, req)
}

// RoleCreate calls devkit.v1.DevkitService.RoleCreate.
func (c *devkitServiceClient) RoleCreate(ctx context.Context, req *connect.Request[v1.RoleCreateRequest]) (*connect.Response[v1.RoleCreateResponse], error) {
	return c.roleCreate.CallUnary(ctx, req)
}

// DevkitServiceHandler is an implementation of the devkit.v1.DevkitService service.
type DevkitServiceHandler interface {
	HelloWorld(context.Context, *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error)
	RoleCreate(context.Context, *connect.Request[v1.RoleCreateRequest]) (*connect.Response[v1.RoleCreateResponse], error)
}

// NewDevkitServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDevkitServiceHandler(svc DevkitServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	devkitServiceHelloWorldHandler := connect.NewUnaryHandler(
		DevkitServiceHelloWorldProcedure,
		svc.HelloWorld,
		connect.WithSchema(devkitServiceHelloWorldMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	devkitServiceRoleCreateHandler := connect.NewUnaryHandler(
		DevkitServiceRoleCreateProcedure,
		svc.RoleCreate,
		connect.WithSchema(devkitServiceRoleCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/devkit.v1.DevkitService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DevkitServiceHelloWorldProcedure:
			devkitServiceHelloWorldHandler.ServeHTTP(w, r)
		case DevkitServiceRoleCreateProcedure:
			devkitServiceRoleCreateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDevkitServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDevkitServiceHandler struct{}

func (UnimplementedDevkitServiceHandler) HelloWorld(context.Context, *connect.Request[v1.HelloWorldRequest]) (*connect.Response[v1.HelloWorldResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("devkit.v1.DevkitService.HelloWorld is not implemented"))
}

func (UnimplementedDevkitServiceHandler) RoleCreate(context.Context, *connect.Request[v1.RoleCreateRequest]) (*connect.Response[v1.RoleCreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("devkit.v1.DevkitService.RoleCreate is not implemented"))
}
