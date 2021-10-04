// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MoviesServiceClient is the client API for MoviesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesServiceClient interface {
	FetchMovie(ctx context.Context, in *FetchMovieReq, opts ...grpc.CallOption) (*FetchMovieResponse, error)
	DetailMovie(ctx context.Context, in *DetailMovieReq, opts ...grpc.CallOption) (*DetailMovieResponse, error)
}

type moviesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesServiceClient(cc grpc.ClientConnInterface) MoviesServiceClient {
	return &moviesServiceClient{cc}
}

func (c *moviesServiceClient) FetchMovie(ctx context.Context, in *FetchMovieReq, opts ...grpc.CallOption) (*FetchMovieResponse, error) {
	out := new(FetchMovieResponse)
	err := c.cc.Invoke(ctx, "/movie.MoviesService/FetchMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesServiceClient) DetailMovie(ctx context.Context, in *DetailMovieReq, opts ...grpc.CallOption) (*DetailMovieResponse, error) {
	out := new(DetailMovieResponse)
	err := c.cc.Invoke(ctx, "/movie.MoviesService/DetailMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServiceServer is the server API for MoviesService service.
// All implementations must embed UnimplementedMoviesServiceServer
// for forward compatibility
type MoviesServiceServer interface {
	FetchMovie(context.Context, *FetchMovieReq) (*FetchMovieResponse, error)
	DetailMovie(context.Context, *DetailMovieReq) (*DetailMovieResponse, error)
	mustEmbedUnimplementedMoviesServiceServer()
}

// UnimplementedMoviesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMoviesServiceServer struct {
}

func (UnimplementedMoviesServiceServer) FetchMovie(context.Context, *FetchMovieReq) (*FetchMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMovie not implemented")
}
func (UnimplementedMoviesServiceServer) DetailMovie(context.Context, *DetailMovieReq) (*DetailMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailMovie not implemented")
}
func (UnimplementedMoviesServiceServer) mustEmbedUnimplementedMoviesServiceServer() {}

// UnsafeMoviesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesServiceServer will
// result in compilation errors.
type UnsafeMoviesServiceServer interface {
	mustEmbedUnimplementedMoviesServiceServer()
}

func RegisterMoviesServiceServer(s grpc.ServiceRegistrar, srv MoviesServiceServer) {
	s.RegisterService(&MoviesService_ServiceDesc, srv)
}

func _MoviesService_FetchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).FetchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.MoviesService/FetchMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).FetchMovie(ctx, req.(*FetchMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoviesService_DetailMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServiceServer).DetailMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.MoviesService/DetailMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServiceServer).DetailMovie(ctx, req.(*DetailMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MoviesService_ServiceDesc is the grpc.ServiceDesc for MoviesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoviesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movie.MoviesService",
	HandlerType: (*MoviesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMovie",
			Handler:    _MoviesService_FetchMovie_Handler,
		},
		{
			MethodName: "DetailMovie",
			Handler:    _MoviesService_DetailMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
