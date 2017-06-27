// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compute-processor/compute-processor.proto

/*
Package processor is a generated protocol buffer package.

It is generated from these files:
	compute-processor/compute-processor.proto

It has these top-level messages:
*/
package processor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import renders "github.com/stupschwartz/qubit/proto-gen/go/renders"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RenderRequest from public import renders/renders.proto
type RenderRequest renders.RenderRequest

func (m *RenderRequest) Reset()                 { (*renders.RenderRequest)(m).Reset() }
func (m *RenderRequest) String() string         { return (*renders.RenderRequest)(m).String() }
func (*RenderRequest) ProtoMessage()            {}
func (m *RenderRequest) GetOperatorId() string  { return (*renders.RenderRequest)(m).GetOperatorId() }
func (m *RenderRequest) GetParameterId() string { return (*renders.RenderRequest)(m).GetParameterId() }
func (m *RenderRequest) GetTime() int32         { return (*renders.RenderRequest)(m).GetTime() }

// RenderResponse from public import renders/renders.proto
type RenderResponse renders.RenderResponse

func (m *RenderResponse) Reset()          { (*renders.RenderResponse)(m).Reset() }
func (m *RenderResponse) String() string  { return (*renders.RenderResponse)(m).String() }
func (*RenderResponse) ProtoMessage()     {}
func (m *RenderResponse) GetData() []byte { return (*renders.RenderResponse)(m).GetData() }

func init() { proto.RegisterFile("compute-processor/compute-processor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0xcf, 0x2d,
	0x28, 0x2d, 0x49, 0xd5, 0x2d, 0x28, 0xca, 0x4f, 0x4e, 0x2d, 0x2e, 0xce, 0x2f, 0xd2, 0xc7, 0x10,
	0xd1, 0x03, 0xb2, 0x4a, 0xf2, 0x85, 0x38, 0xe1, 0x02, 0x52, 0xa2, 0x45, 0xa9, 0x79, 0x29, 0xa9,
	0x45, 0xc5, 0xfa, 0x50, 0x1a, 0xa2, 0x22, 0x80, 0x21, 0x89, 0x0d, 0xcc, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0x21, 0x14, 0xdf, 0x57, 0x00, 0x00, 0x00,
}
