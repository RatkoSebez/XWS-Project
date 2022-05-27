// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package job_offer_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobOfferServiceClient is the client API for JobOfferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobOfferServiceClient interface {
	CreateOffer(ctx context.Context, in *OfferRequest, opts ...grpc.CallOption) (*Offer, error)
	GetOffersByCompany(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Offers, error)
	GetOffersByPosition(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Offers, error)
}

type jobOfferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobOfferServiceClient(cc grpc.ClientConnInterface) JobOfferServiceClient {
	return &jobOfferServiceClient{cc}
}

func (c *jobOfferServiceClient) CreateOffer(ctx context.Context, in *OfferRequest, opts ...grpc.CallOption) (*Offer, error) {
	out := new(Offer)
	err := c.cc.Invoke(ctx, "/job_offer.JobOfferService/CreateOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) GetOffersByCompany(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Offers, error) {
	out := new(Offers)
	err := c.cc.Invoke(ctx, "/job_offer.JobOfferService/GetOffersByCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) GetOffersByPosition(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Offers, error) {
	out := new(Offers)
	err := c.cc.Invoke(ctx, "/job_offer.JobOfferService/GetOffersByPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobOfferServiceServer is the server API for JobOfferService service.
// All implementations must embed UnimplementedJobOfferServiceServer
// for forward compatibility
type JobOfferServiceServer interface {
	CreateOffer(context.Context, *OfferRequest) (*Offer, error)
	GetOffersByCompany(context.Context, *CompanyID) (*Offers, error)
	GetOffersByPosition(context.Context, *Position) (*Offers, error)
	mustEmbedUnimplementedJobOfferServiceServer()
}

// UnimplementedJobOfferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobOfferServiceServer struct {
}

func (*UnimplementedJobOfferServiceServer) CreateOffer(context.Context, *OfferRequest) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffer not implemented")
}
func (*UnimplementedJobOfferServiceServer) GetOffersByCompany(context.Context, *CompanyID) (*Offers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffersByCompany not implemented")
}
func (*UnimplementedJobOfferServiceServer) GetOffersByPosition(context.Context, *Position) (*Offers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffersByPosition not implemented")
}
func (*UnimplementedJobOfferServiceServer) mustEmbedUnimplementedJobOfferServiceServer() {}

func RegisterJobOfferServiceServer(s *grpc.Server, srv JobOfferServiceServer) {
	s.RegisterService(&_JobOfferService_serviceDesc, srv)
}

func _JobOfferService_CreateOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).CreateOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offer.JobOfferService/CreateOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).CreateOffer(ctx, req.(*OfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_GetOffersByCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).GetOffersByCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offer.JobOfferService/GetOffersByCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).GetOffersByCompany(ctx, req.(*CompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_GetOffersByPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).GetOffersByPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job_offer.JobOfferService/GetOffersByPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).GetOffersByPosition(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobOfferService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "job_offer.JobOfferService",
	HandlerType: (*JobOfferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOffer",
			Handler:    _JobOfferService_CreateOffer_Handler,
		},
		{
			MethodName: "GetOffersByCompany",
			Handler:    _JobOfferService_GetOffersByCompany_Handler,
		},
		{
			MethodName: "GetOffersByPosition",
			Handler:    _JobOfferService_GetOffersByPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_offer_service.proto",
}
