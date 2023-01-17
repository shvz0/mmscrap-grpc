// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: scrappers.proto

package scrappers

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

// ScrappersClient is the client API for Scrappers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScrappersClient interface {
	GetTodays24NewsArticles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Scrappers_GetTodays24NewsArticlesClient, error)
}

type scrappersClient struct {
	cc grpc.ClientConnInterface
}

func NewScrappersClient(cc grpc.ClientConnInterface) ScrappersClient {
	return &scrappersClient{cc}
}

func (c *scrappersClient) GetTodays24NewsArticles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Scrappers_GetTodays24NewsArticlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scrappers_ServiceDesc.Streams[0], "/Scrappers/GetTodays24NewsArticles", opts...)
	if err != nil {
		return nil, err
	}
	x := &scrappersGetTodays24NewsArticlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scrappers_GetTodays24NewsArticlesClient interface {
	Recv() (*Article, error)
	grpc.ClientStream
}

type scrappersGetTodays24NewsArticlesClient struct {
	grpc.ClientStream
}

func (x *scrappersGetTodays24NewsArticlesClient) Recv() (*Article, error) {
	m := new(Article)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ScrappersServer is the server API for Scrappers service.
// All implementations must embed UnimplementedScrappersServer
// for forward compatibility
type ScrappersServer interface {
	GetTodays24NewsArticles(*Empty, Scrappers_GetTodays24NewsArticlesServer) error
	mustEmbedUnimplementedScrappersServer()
}

// UnimplementedScrappersServer must be embedded to have forward compatible implementations.
type UnimplementedScrappersServer struct {
}

func (UnimplementedScrappersServer) GetTodays24NewsArticles(*Empty, Scrappers_GetTodays24NewsArticlesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTodays24NewsArticles not implemented")
}
func (UnimplementedScrappersServer) mustEmbedUnimplementedScrappersServer() {}

// UnsafeScrappersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScrappersServer will
// result in compilation errors.
type UnsafeScrappersServer interface {
	mustEmbedUnimplementedScrappersServer()
}

func RegisterScrappersServer(s grpc.ServiceRegistrar, srv ScrappersServer) {
	s.RegisterService(&Scrappers_ServiceDesc, srv)
}

func _Scrappers_GetTodays24NewsArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScrappersServer).GetTodays24NewsArticles(m, &scrappersGetTodays24NewsArticlesServer{stream})
}

type Scrappers_GetTodays24NewsArticlesServer interface {
	Send(*Article) error
	grpc.ServerStream
}

type scrappersGetTodays24NewsArticlesServer struct {
	grpc.ServerStream
}

func (x *scrappersGetTodays24NewsArticlesServer) Send(m *Article) error {
	return x.ServerStream.SendMsg(m)
}

// Scrappers_ServiceDesc is the grpc.ServiceDesc for Scrappers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scrappers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Scrappers",
	HandlerType: (*ScrappersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTodays24NewsArticles",
			Handler:       _Scrappers_GetTodays24NewsArticles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "scrappers.proto",
}