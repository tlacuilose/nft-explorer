// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package nft_explorer_proto

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

// NFTExplorerClient is the client API for NFTExplorer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NFTExplorerClient interface {
	// Obtain the artworks owned by a given account.
	ListOwnedArtworks(ctx context.Context, in *Account, opts ...grpc.CallOption) (NFTExplorer_ListOwnedArtworksClient, error)
}

type nFTExplorerClient struct {
	cc grpc.ClientConnInterface
}

func NewNFTExplorerClient(cc grpc.ClientConnInterface) NFTExplorerClient {
	return &nFTExplorerClient{cc}
}

func (c *nFTExplorerClient) ListOwnedArtworks(ctx context.Context, in *Account, opts ...grpc.CallOption) (NFTExplorer_ListOwnedArtworksClient, error) {
	stream, err := c.cc.NewStream(ctx, &NFTExplorer_ServiceDesc.Streams[0], "/nft_explorer_proto.NFTExplorer/ListOwnedArtworks", opts...)
	if err != nil {
		return nil, err
	}
	x := &nFTExplorerListOwnedArtworksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NFTExplorer_ListOwnedArtworksClient interface {
	Recv() (*Artwork, error)
	grpc.ClientStream
}

type nFTExplorerListOwnedArtworksClient struct {
	grpc.ClientStream
}

func (x *nFTExplorerListOwnedArtworksClient) Recv() (*Artwork, error) {
	m := new(Artwork)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NFTExplorerServer is the server API for NFTExplorer service.
// All implementations must embed UnimplementedNFTExplorerServer
// for forward compatibility
type NFTExplorerServer interface {
	// Obtain the artworks owned by a given account.
	ListOwnedArtworks(*Account, NFTExplorer_ListOwnedArtworksServer) error
	mustEmbedUnimplementedNFTExplorerServer()
}

// UnimplementedNFTExplorerServer must be embedded to have forward compatible implementations.
type UnimplementedNFTExplorerServer struct {
}

func (UnimplementedNFTExplorerServer) ListOwnedArtworks(*Account, NFTExplorer_ListOwnedArtworksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOwnedArtworks not implemented")
}
func (UnimplementedNFTExplorerServer) mustEmbedUnimplementedNFTExplorerServer() {}

// UnsafeNFTExplorerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NFTExplorerServer will
// result in compilation errors.
type UnsafeNFTExplorerServer interface {
	mustEmbedUnimplementedNFTExplorerServer()
}

func RegisterNFTExplorerServer(s grpc.ServiceRegistrar, srv NFTExplorerServer) {
	s.RegisterService(&NFTExplorer_ServiceDesc, srv)
}

func _NFTExplorer_ListOwnedArtworks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Account)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NFTExplorerServer).ListOwnedArtworks(m, &nFTExplorerListOwnedArtworksServer{stream})
}

type NFTExplorer_ListOwnedArtworksServer interface {
	Send(*Artwork) error
	grpc.ServerStream
}

type nFTExplorerListOwnedArtworksServer struct {
	grpc.ServerStream
}

func (x *nFTExplorerListOwnedArtworksServer) Send(m *Artwork) error {
	return x.ServerStream.SendMsg(m)
}

// NFTExplorer_ServiceDesc is the grpc.ServiceDesc for NFTExplorer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NFTExplorer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nft_explorer_proto.NFTExplorer",
	HandlerType: (*NFTExplorerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOwnedArtworks",
			Handler:       _NFTExplorer_ListOwnedArtworks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nft_explorer.proto",
}
