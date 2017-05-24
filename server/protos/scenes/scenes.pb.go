// Code generated by protoc-gen-go.
// source: scenes.proto
// DO NOT EDIT!

/*
Package scenes is a generated protocol buffer package.

It is generated from these files:
	scenes.proto

It has these top-level messages:
	Scene
	ListScenesRequest
	ListScenesResponse
	GetSceneRequest
	CreateSceneRequest
	UpdateSceneRequest
	DeleteSceneRequest
*/
package scenes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

type Scene struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Scene) Reset()                    { *m = Scene{} }
func (m *Scene) String() string            { return proto.CompactTextString(m) }
func (*Scene) ProtoMessage()               {}
func (*Scene) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Scene) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListScenesRequest struct {
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListScenesRequest) Reset()                    { *m = ListScenesRequest{} }
func (m *ListScenesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListScenesRequest) ProtoMessage()               {}
func (*ListScenesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListScenesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListScenesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListScenesResponse struct {
	Scenes        []*Scene `protobuf:"bytes,1,rep,name=scenes" json:"scenes,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListScenesResponse) Reset()                    { *m = ListScenesResponse{} }
func (m *ListScenesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListScenesResponse) ProtoMessage()               {}
func (*ListScenesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListScenesResponse) GetScenes() []*Scene {
	if m != nil {
		return m.Scenes
	}
	return nil
}

func (m *ListScenesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetSceneRequest struct {
	SceneId int64 `protobuf:"varint,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
}

func (m *GetSceneRequest) Reset()                    { *m = GetSceneRequest{} }
func (m *GetSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSceneRequest) ProtoMessage()               {}
func (*GetSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetSceneRequest) GetSceneId() int64 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

type CreateSceneRequest struct {
	Scene *Scene `protobuf:"bytes,1,opt,name=scene" json:"scene,omitempty"`
}

func (m *CreateSceneRequest) Reset()                    { *m = CreateSceneRequest{} }
func (m *CreateSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSceneRequest) ProtoMessage()               {}
func (*CreateSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateSceneRequest) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type UpdateSceneRequest struct {
	SceneId int64  `protobuf:"varint,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	Scene   *Scene `protobuf:"bytes,2,opt,name=scene" json:"scene,omitempty"`
}

func (m *UpdateSceneRequest) Reset()                    { *m = UpdateSceneRequest{} }
func (m *UpdateSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSceneRequest) ProtoMessage()               {}
func (*UpdateSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateSceneRequest) GetSceneId() int64 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

func (m *UpdateSceneRequest) GetScene() *Scene {
	if m != nil {
		return m.Scene
	}
	return nil
}

type DeleteSceneRequest struct {
	SceneId int64 `protobuf:"varint,1,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
}

func (m *DeleteSceneRequest) Reset()                    { *m = DeleteSceneRequest{} }
func (m *DeleteSceneRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSceneRequest) ProtoMessage()               {}
func (*DeleteSceneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteSceneRequest) GetSceneId() int64 {
	if m != nil {
		return m.SceneId
	}
	return 0
}

func init() {
	proto.RegisterType((*Scene)(nil), "scenes.Scene")
	proto.RegisterType((*ListScenesRequest)(nil), "scenes.ListScenesRequest")
	proto.RegisterType((*ListScenesResponse)(nil), "scenes.ListScenesResponse")
	proto.RegisterType((*GetSceneRequest)(nil), "scenes.GetSceneRequest")
	proto.RegisterType((*CreateSceneRequest)(nil), "scenes.CreateSceneRequest")
	proto.RegisterType((*UpdateSceneRequest)(nil), "scenes.UpdateSceneRequest")
	proto.RegisterType((*DeleteSceneRequest)(nil), "scenes.DeleteSceneRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scenes service

type ScenesClient interface {
	List(ctx context.Context, in *ListScenesRequest, opts ...grpc.CallOption) (*ListScenesResponse, error)
	Get(ctx context.Context, in *GetSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Create(ctx context.Context, in *CreateSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Update(ctx context.Context, in *UpdateSceneRequest, opts ...grpc.CallOption) (*Scene, error)
	Delete(ctx context.Context, in *DeleteSceneRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type scenesClient struct {
	cc *grpc.ClientConn
}

func NewScenesClient(cc *grpc.ClientConn) ScenesClient {
	return &scenesClient{cc}
}

func (c *scenesClient) List(ctx context.Context, in *ListScenesRequest, opts ...grpc.CallOption) (*ListScenesResponse, error) {
	out := new(ListScenesResponse)
	err := grpc.Invoke(ctx, "/scenes.Scenes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Get(ctx context.Context, in *GetSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Create(ctx context.Context, in *CreateSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Update(ctx context.Context, in *UpdateSceneRequest, opts ...grpc.CallOption) (*Scene, error) {
	out := new(Scene)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scenesClient) Delete(ctx context.Context, in *DeleteSceneRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/scenes.Scenes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scenes service

type ScenesServer interface {
	List(context.Context, *ListScenesRequest) (*ListScenesResponse, error)
	Get(context.Context, *GetSceneRequest) (*Scene, error)
	Create(context.Context, *CreateSceneRequest) (*Scene, error)
	Update(context.Context, *UpdateSceneRequest) (*Scene, error)
	Delete(context.Context, *DeleteSceneRequest) (*google_protobuf1.Empty, error)
}

func RegisterScenesServer(s *grpc.Server, srv ScenesServer) {
	s.RegisterService(&_Scenes_serviceDesc, srv)
}

func _Scenes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScenesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).List(ctx, req.(*ListScenesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Get(ctx, req.(*GetSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Create(ctx, req.(*CreateSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Update(ctx, req.(*UpdateSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scenes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSceneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScenesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scenes.Scenes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScenesServer).Delete(ctx, req.(*DeleteSceneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scenes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scenes.Scenes",
	HandlerType: (*ScenesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Scenes_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Scenes_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Scenes_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Scenes_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Scenes_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scenes.proto",
}

func init() { proto.RegisterFile("scenes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xc7, 0x85, 0x5d, 0x5c, 0x98, 0x42, 0x51, 0xa7, 0x6a, 0x01, 0x53, 0xaa, 0x6a, 0xfb, 0x21,
	0x54, 0x55, 0xb6, 0x4a, 0x4f, 0xed, 0xb5, 0xad, 0xaa, 0x4a, 0xa8, 0x4d, 0x0c, 0x39, 0x44, 0x39,
	0x20, 0x03, 0x13, 0x64, 0x85, 0xd8, 0x0e, 0x5e, 0xa2, 0x7c, 0x28, 0x97, 0xbc, 0x42, 0x9e, 0x24,
	0xcf, 0x92, 0x57, 0xc8, 0x83, 0xc4, 0xde, 0xb5, 0x31, 0xd8, 0x24, 0xca, 0x6d, 0x67, 0xff, 0xb3,
	0xbf, 0x99, 0xf9, 0x8f, 0x16, 0x2a, 0xc1, 0x98, 0x5c, 0x0a, 0x0c, 0x7f, 0xee, 0x71, 0x0f, 0x35,
	0x19, 0xe9, 0x6f, 0xa6, 0x9e, 0x37, 0x9d, 0x91, 0x69, 0xfb, 0x8e, 0x69, 0xbb, 0xae, 0xc7, 0x6d,
	0xee, 0x78, 0x6e, 0x9c, 0xa5, 0xb7, 0x62, 0x55, 0x44, 0xa3, 0xc5, 0xbe, 0x49, 0x87, 0x3e, 0x3f,
	0x95, 0x22, 0xab, 0x43, 0xb1, 0x1f, 0x41, 0xf0, 0x39, 0x28, 0xce, 0xa4, 0x51, 0x78, 0x57, 0xe8,
	0xa8, 0x56, 0x78, 0x62, 0xff, 0xe1, 0x45, 0xcf, 0x09, 0xb8, 0x10, 0x03, 0x8b, 0x8e, 0x16, 0x14,
	0x70, 0x6c, 0x41, 0xd9, 0xb7, 0xa7, 0x34, 0x0c, 0x9c, 0x33, 0x6a, 0x28, 0x61, 0x6e, 0xd1, 0x2a,
	0x45, 0x17, 0xfd, 0x30, 0xc6, 0x36, 0x80, 0x10, 0xb9, 0x77, 0x40, 0x6e, 0x43, 0x0d, 0xd5, 0xb2,
	0x25, 0xd2, 0x07, 0xd1, 0x05, 0x1b, 0x03, 0xae, 0x02, 0x03, 0x3f, 0xec, 0x90, 0xf0, 0x23, 0xc4,
	0x43, 0x84, 0xa5, 0xd5, 0xce, 0xb3, 0x6e, 0xd5, 0x88, 0x27, 0x14, 0x79, 0x56, 0x2c, 0xe2, 0x27,
	0xa8, 0xb9, 0x74, 0xc2, 0x87, 0x2b, 0x05, 0x14, 0x51, 0xa0, 0x1a, 0x5d, 0x6f, 0x2d, 0x8b, 0x7c,
	0x81, 0xda, 0x1f, 0x92, 0x35, 0x92, 0x9e, 0x9b, 0x50, 0x12, 0x90, 0xe1, 0x72, 0xbc, 0xa7, 0x22,
	0xfe, 0x3b, 0x61, 0xdf, 0x01, 0x7f, 0xce, 0xc9, 0xe6, 0xb4, 0xf6, 0xe0, 0x3d, 0x14, 0x45, 0x82,
	0xc8, 0xce, 0x75, 0x24, 0x35, 0x36, 0x00, 0xdc, 0xf1, 0x27, 0xd9, 0xa7, 0xf7, 0xd7, 0x4a, 0xa9,
	0xca, 0x03, 0x54, 0x13, 0xf0, 0x17, 0xcd, 0xe8, 0xd1, 0xd4, 0xee, 0xb5, 0x0a, 0x9a, 0x74, 0x14,
	0xb7, 0xe1, 0x49, 0xe4, 0x2f, 0x36, 0x13, 0x72, 0x6e, 0x7d, 0xba, 0xbe, 0x49, 0x92, 0x8b, 0x60,
	0x78, 0x79, 0x73, 0x7b, 0xa5, 0x54, 0x10, 0xcc, 0xe3, 0xaf, 0x66, 0xec, 0x7a, 0x0f, 0xd4, 0xd0,
	0x4d, 0xac, 0x27, 0xcf, 0x32, 0xd6, 0xea, 0xeb, 0x43, 0xb0, 0xb6, 0x40, 0xd4, 0xf1, 0x55, 0x8a,
	0x30, 0xcf, 0x93, 0xce, 0x2f, 0xf0, 0x1f, 0x68, 0xd2, 0x6d, 0x5c, 0xf6, 0x91, 0x77, 0x3f, 0xcb,
	0x6c, 0x0a, 0xe6, 0x4b, 0xb6, 0xd2, 0xd6, 0x0f, 0x69, 0x16, 0xee, 0x82, 0x26, 0x57, 0x90, 0xf2,
	0xf2, 0x2b, 0xc9, 0xf2, 0x3e, 0x08, 0xde, 0x5b, 0x7d, 0x73, 0x8f, 0x09, 0x7a, 0x0f, 0x34, 0xb9,
	0x87, 0x14, 0x9d, 0xdf, 0x8b, 0xfe, 0xda, 0x90, 0x3f, 0xcb, 0x48, 0x7e, 0x96, 0xf1, 0x3b, 0xfa,
	0x59, 0x89, 0x0f, 0x9f, 0x37, 0xd7, 0x18, 0x69, 0x22, 0xfd, 0xdb, 0x5d, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x52, 0xa4, 0xa0, 0xcc, 0x03, 0x00, 0x00,
}
