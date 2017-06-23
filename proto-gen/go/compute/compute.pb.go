// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compute/compute.proto

/*
Package compute is a generated protocol buffer package.

It is generated from these files:
	compute/compute.proto

It has these top-level messages:
	CreateComputationRequest
	ComputationStatusRequest
	ComputationStatusResponse
*/
package compute

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import operators "github.com/stupschwartz/qubit/proto-gen/go/operators"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateComputationRequest struct {
	RootOperatorId string                         `protobuf:"bytes,1,opt,name=root_operator_id,json=rootOperatorId" json:"root_operator_id,omitempty"`
	OperatorMap    map[string]*operators.Operator `protobuf:"bytes,2,rep,name=operator_map,json=operatorMap" json:"operator_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateComputationRequest) Reset()                    { *m = CreateComputationRequest{} }
func (m *CreateComputationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateComputationRequest) ProtoMessage()               {}
func (*CreateComputationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateComputationRequest) GetRootOperatorId() string {
	if m != nil {
		return m.RootOperatorId
	}
	return ""
}

func (m *CreateComputationRequest) GetOperatorMap() map[string]*operators.Operator {
	if m != nil {
		return m.OperatorMap
	}
	return nil
}

type ComputationStatusRequest struct {
	ComputationId string `protobuf:"bytes,1,opt,name=computation_id,json=computationId" json:"computation_id,omitempty"`
}

func (m *ComputationStatusRequest) Reset()                    { *m = ComputationStatusRequest{} }
func (m *ComputationStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ComputationStatusRequest) ProtoMessage()               {}
func (*ComputationStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ComputationStatusRequest) GetComputationId() string {
	if m != nil {
		return m.ComputationId
	}
	return ""
}

type ComputationStatusResponse struct {
	ComputationId string `protobuf:"bytes,1,opt,name=computation_id,json=computationId" json:"computation_id,omitempty"`
	Status        string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ComputationStatusResponse) Reset()                    { *m = ComputationStatusResponse{} }
func (m *ComputationStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ComputationStatusResponse) ProtoMessage()               {}
func (*ComputationStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ComputationStatusResponse) GetComputationId() string {
	if m != nil {
		return m.ComputationId
	}
	return ""
}

func (m *ComputationStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateComputationRequest)(nil), "compute.CreateComputationRequest")
	proto.RegisterType((*ComputationStatusRequest)(nil), "compute.ComputationStatusRequest")
	proto.RegisterType((*ComputationStatusResponse)(nil), "compute.ComputationStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Compute service

type ComputeClient interface {
	CreateComputation(ctx context.Context, in *CreateComputationRequest, opts ...grpc.CallOption) (*ComputationStatusResponse, error)
	GetComputationStatus(ctx context.Context, in *ComputationStatusRequest, opts ...grpc.CallOption) (*ComputationStatusResponse, error)
}

type computeClient struct {
	cc *grpc.ClientConn
}

func NewComputeClient(cc *grpc.ClientConn) ComputeClient {
	return &computeClient{cc}
}

func (c *computeClient) CreateComputation(ctx context.Context, in *CreateComputationRequest, opts ...grpc.CallOption) (*ComputationStatusResponse, error) {
	out := new(ComputationStatusResponse)
	err := grpc.Invoke(ctx, "/compute.Compute/CreateComputation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeClient) GetComputationStatus(ctx context.Context, in *ComputationStatusRequest, opts ...grpc.CallOption) (*ComputationStatusResponse, error) {
	out := new(ComputationStatusResponse)
	err := grpc.Invoke(ctx, "/compute.Compute/GetComputationStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Compute service

type ComputeServer interface {
	CreateComputation(context.Context, *CreateComputationRequest) (*ComputationStatusResponse, error)
	GetComputationStatus(context.Context, *ComputationStatusRequest) (*ComputationStatusResponse, error)
}

func RegisterComputeServer(s *grpc.Server, srv ComputeServer) {
	s.RegisterService(&_Compute_serviceDesc, srv)
}

func _Compute_CreateComputation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComputationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).CreateComputation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compute.Compute/CreateComputation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).CreateComputation(ctx, req.(*CreateComputationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Compute_GetComputationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).GetComputationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compute.Compute/GetComputationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).GetComputationStatus(ctx, req.(*ComputationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Compute_serviceDesc = grpc.ServiceDesc{
	ServiceName: "compute.Compute",
	HandlerType: (*ComputeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComputation",
			Handler:    _Compute_CreateComputation_Handler,
		},
		{
			MethodName: "GetComputationStatus",
			Handler:    _Compute_GetComputationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compute/compute.proto",
}

func init() { proto.RegisterFile("compute/compute.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2d,
	0x28, 0x2d, 0x49, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x64, 0x7e, 0x41, 0x6a, 0x51, 0x62, 0x49, 0x7e, 0x51, 0xb1, 0x3e, 0x9c, 0x05, 0x51, 0xa3,
	0xf4, 0x81, 0x91, 0x4b, 0xc2, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xd5, 0x19, 0xac, 0x38, 0xb1, 0x24,
	0x33, 0x3f, 0x2f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x83, 0x4b, 0xa0, 0x28, 0x3f,
	0xbf, 0x24, 0x1e, 0xa6, 0x29, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x0f,
	0x24, 0xee, 0x0f, 0x15, 0xf6, 0x4c, 0x11, 0x0a, 0xe5, 0xe2, 0x81, 0x2b, 0xca, 0x4d, 0x2c, 0x90,
	0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x32, 0xd2, 0x83, 0x39, 0x08, 0x97, 0x15, 0x7a, 0x30, 0x33,
	0x7c, 0x13, 0x0b, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xb8, 0xf3, 0x11, 0x22, 0x52, 0xc1, 0x5c,
	0x02, 0xe8, 0x0a, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0xa1, 0xee, 0x00, 0x31, 0x85, 0x34,
	0xb9, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x81, 0xb6, 0x32, 0x02, 0x6d, 0x15, 0xd6, 0x43, 0x78,
	0x12, 0xa6, 0x3b, 0x08, 0xa2, 0xc2, 0x8a, 0xc9, 0x82, 0x51, 0xc9, 0x11, 0xe8, 0x63, 0x84, 0x43,
	0x82, 0x81, 0x54, 0x69, 0x31, 0xcc, 0xc7, 0xaa, 0x5c, 0x7c, 0xc9, 0x08, 0x39, 0x84, 0x7f, 0x79,
	0x91, 0x44, 0x3d, 0x53, 0x94, 0xa2, 0xb8, 0x24, 0xb1, 0x18, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x4a, 0xa4, 0x19, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x60, 0x8d, 0x60, 0x67, 0x73, 0x06, 0x41, 0x79,
	0x46, 0x27, 0x18, 0xb9, 0xd8, 0x21, 0x86, 0xa7, 0x0a, 0xc5, 0x70, 0x09, 0x62, 0x84, 0x9c, 0x90,
	0x22, 0xc1, 0x50, 0x95, 0x52, 0x42, 0x28, 0xc1, 0xe5, 0x4c, 0x25, 0x06, 0xa1, 0x78, 0x2e, 0x11,
	0xf7, 0xd4, 0x12, 0x0c, 0x15, 0xc8, 0x16, 0xe0, 0x08, 0x27, 0xe2, 0x2c, 0x48, 0x62, 0x03, 0xa7,
	0x31, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xfc, 0x7c, 0x65, 0xa0, 0x02, 0x00, 0x00,
}
