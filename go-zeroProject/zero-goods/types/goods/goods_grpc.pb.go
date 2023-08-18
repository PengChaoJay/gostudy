// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: goods.proto

package goods

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

const (
	Goods_GetGoods_FullMethodName = "/goods.Goods/getGoods"
)

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	// rpc方法
	GetGoods(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGoods(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsResponse, error) {
	out := new(GoodsResponse)
	err := c.cc.Invoke(ctx, Goods_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	// rpc方法
	GetGoods(context.Context, *GoodsRequest) (*GoodsResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GetGoods(context.Context, *GoodsRequest) (*GoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Goods_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoods(ctx, req.(*GoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getGoods",
			Handler:    _Goods_GetGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods.proto",
}
