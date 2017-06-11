// Code generated by protoc-gen-go. DO NOT EDIT.
// source: projects/projects.proto

/*
Package projects is a generated protocol buffer package.

It is generated from these files:
	projects/projects.proto

It has these top-level messages:
	Project
	ListProjectsRequest
	ListProjectsResponse
	GetProjectRequest
	CreateProjectRequest
	UpdateProjectRequest
	DeleteProjectRequest
*/
package projects

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

type Project struct {
	Id             string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrganizationId string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId" json:"organization_id,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListProjectsRequest struct {
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListProjectsRequest) Reset()                    { *m = ListProjectsRequest{} }
func (m *ListProjectsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsRequest) ProtoMessage()               {}
func (*ListProjectsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListProjectsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListProjectsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListProjectsResponse struct {
	Projects      []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListProjectsResponse) Reset()                    { *m = ListProjectsResponse{} }
func (m *ListProjectsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProjectsResponse) ProtoMessage()               {}
func (*ListProjectsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListProjectsResponse) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *ListProjectsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetProjectRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetProjectRequest) Reset()                    { *m = GetProjectRequest{} }
func (m *GetProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProjectRequest) ProtoMessage()               {}
func (*GetProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetProjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateProjectRequest struct {
	Project *Project `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
}

func (m *CreateProjectRequest) Reset()                    { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()               {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateProjectRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateProjectRequest struct {
	Id      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Project *Project `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
}

func (m *UpdateProjectRequest) Reset()                    { *m = UpdateProjectRequest{} }
func (m *UpdateProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateProjectRequest) ProtoMessage()               {}
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateProjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateProjectRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type DeleteProjectRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteProjectRequest) Reset()                    { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()               {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteProjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "projects.Project")
	proto.RegisterType((*ListProjectsRequest)(nil), "projects.ListProjectsRequest")
	proto.RegisterType((*ListProjectsResponse)(nil), "projects.ListProjectsResponse")
	proto.RegisterType((*GetProjectRequest)(nil), "projects.GetProjectRequest")
	proto.RegisterType((*CreateProjectRequest)(nil), "projects.CreateProjectRequest")
	proto.RegisterType((*UpdateProjectRequest)(nil), "projects.UpdateProjectRequest")
	proto.RegisterType((*DeleteProjectRequest)(nil), "projects.DeleteProjectRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Projects service

type ProjectsClient interface {
	List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error)
	Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type projectsClient struct {
	cc *grpc.ClientConn
}

func NewProjectsClient(cc *grpc.ClientConn) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := grpc.Invoke(ctx, "/projects.Projects/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := grpc.Invoke(ctx, "/projects.Projects/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := grpc.Invoke(ctx, "/projects.Projects/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := grpc.Invoke(ctx, "/projects.Projects/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/projects.Projects/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Projects service

type ProjectsServer interface {
	List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	Get(context.Context, *GetProjectRequest) (*Project, error)
	Create(context.Context, *CreateProjectRequest) (*Project, error)
	Update(context.Context, *UpdateProjectRequest) (*Project, error)
	Delete(context.Context, *DeleteProjectRequest) (*google_protobuf1.Empty, error)
}

func RegisterProjectsServer(s *grpc.Server, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/projects.Projects/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).List(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/projects.Projects/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Get(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/projects.Projects/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Create(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/projects.Projects/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Update(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/projects.Projects/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).Delete(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "projects.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Projects_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Projects_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Projects_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Projects_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Projects_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "projects/projects.proto",
}

func init() { proto.RegisterFile("projects/projects.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x55, 0xbb, 0xb8, 0xe9, 0x00, 0xa9, 0x3a, 0x58, 0xb4, 0xd8, 0x04, 0xc1, 0x82, 0x4a,
	0x55, 0x84, 0x2d, 0xca, 0x8d, 0x6b, 0x41, 0x08, 0x89, 0x43, 0x71, 0x01, 0x21, 0x38, 0x44, 0x4e,
	0x32, 0x58, 0x86, 0xc4, 0x6b, 0xe2, 0x0d, 0x82, 0x20, 0x2e, 0xbc, 0x02, 0x8f, 0xc6, 0x2b, 0x70,
	0xe3, 0x25, 0x58, 0xaf, 0xbd, 0x8e, 0x13, 0x3b, 0xca, 0x6d, 0xbc, 0x33, 0xfe, 0xfe, 0xf1, 0xff,
	0x7b, 0xe1, 0x20, 0x9d, 0xf2, 0x4f, 0x34, 0x14, 0x99, 0xaf, 0x0b, 0x4f, 0x16, 0x82, 0x63, 0x47,
	0x3f, 0x3b, 0x37, 0x23, 0xce, 0xa3, 0x31, 0xf9, 0x61, 0x1a, 0xfb, 0x61, 0x92, 0x70, 0x11, 0x8a,
	0x98, 0x27, 0xe5, 0x9c, 0xe3, 0x96, 0x5d, 0xf5, 0x34, 0x98, 0x7d, 0xf4, 0x69, 0x92, 0x8a, 0xef,
	0x45, 0x93, 0xbd, 0x85, 0x9d, 0xf3, 0x02, 0x83, 0x5d, 0x30, 0xe2, 0xd1, 0xe1, 0xd6, 0xed, 0xad,
	0xe3, 0xdd, 0x40, 0x56, 0x78, 0x1f, 0xf6, 0xf8, 0x34, 0x0a, 0x93, 0x78, 0xae, 0x70, 0x7d, 0xd9,
	0x34, 0x54, 0xb3, 0x5b, 0x3f, 0x7e, 0x31, 0x42, 0x84, 0xed, 0x24, 0x9c, 0xd0, 0xa1, 0xa9, 0xba,
	0xaa, 0x66, 0xaf, 0xe0, 0xda, 0xcb, 0x38, 0x13, 0x25, 0x3b, 0x0b, 0xe8, 0xcb, 0x8c, 0x32, 0x81,
	0x2e, 0xec, 0xa6, 0x61, 0x44, 0xfd, 0x2c, 0x9e, 0x93, 0x92, 0xba, 0x14, 0x74, 0xf2, 0x83, 0x0b,
	0xf9, 0x8c, 0x3d, 0x00, 0xd5, 0x14, 0xfc, 0x33, 0x25, 0xa5, 0x96, 0x1a, 0x7f, 0x9d, 0x1f, 0xb0,
	0x09, 0xd8, 0xcb, 0xc8, 0x2c, 0x95, 0x1f, 0x49, 0xf8, 0x10, 0x2a, 0x27, 0x24, 0xd2, 0x3c, 0xbe,
	0x7c, 0xba, 0xef, 0x55, 0x56, 0x95, 0xd3, 0x41, 0x35, 0x82, 0x47, 0xb0, 0x97, 0xd0, 0x37, 0xd1,
	0x6f, 0x48, 0x5d, 0xcd, 0x8f, 0xcf, 0x2b, 0xb9, 0xbb, 0xb0, 0xff, 0x9c, 0xb4, 0x9a, 0xde, 0x7f,
	0xc5, 0x23, 0x76, 0x06, 0xf6, 0xd9, 0x94, 0x42, 0x41, 0x2b, 0x73, 0x0f, 0x60, 0xa7, 0x14, 0x54,
	0xc3, 0xad, 0x2b, 0xe9, 0x09, 0x76, 0x01, 0xf6, 0x9b, 0x74, 0xd4, 0x84, 0xac, 0x06, 0x52, 0x83,
	0x1a, 0x1b, 0xa1, 0x47, 0x60, 0x3f, 0xa5, 0x31, 0x6d, 0x82, 0x9e, 0xfe, 0x33, 0xa1, 0xa3, 0x2d,
	0xc5, 0x0f, 0xb0, 0x9d, 0x5b, 0x8c, 0xbd, 0x05, 0xb8, 0x25, 0x45, 0xe7, 0xd6, 0xba, 0x76, 0x91,
	0x08, 0xb3, 0x7f, 0xfd, 0xf9, 0xfb, 0xdb, 0xe8, 0xe2, 0x15, 0xff, 0xeb, 0xa3, 0xea, 0xaf, 0xc5,
	0x77, 0x60, 0x4a, 0x43, 0xd1, 0x5d, 0xbc, 0xdc, 0xf0, 0xd7, 0x69, 0x7e, 0x11, 0xbb, 0xa3, 0x60,
	0x2e, 0xde, 0xa8, 0xc3, 0xfc, 0x1f, 0x65, 0x25, 0xff, 0xca, 0x9f, 0xf8, 0x1e, 0xac, 0x22, 0x05,
	0xac, 0x6d, 0xd6, 0x96, 0x4b, 0x1b, 0xbf, 0xa7, 0xf8, 0x07, 0x6c, 0x69, 0xd9, 0x27, 0xda, 0x47,
	0x8c, 0xc0, 0x2a, 0xc2, 0xa9, 0xb3, 0xdb, 0xe2, 0x6a, 0x63, 0x9f, 0x28, 0xf6, 0x3d, 0x67, 0xfd,
	0xee, 0x0b, 0xa1, 0x21, 0x58, 0x45, 0x60, 0x75, 0xa1, 0xb6, 0x08, 0x9d, 0xeb, 0x5e, 0x71, 0xa3,
	0x3d, 0x7d, 0xa3, 0xbd, 0x67, 0xf9, 0x8d, 0xd6, 0x4e, 0x9d, 0xac, 0x57, 0x1b, 0x58, 0xea, 0x95,
	0xc7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x92, 0x4a, 0xdf, 0x55, 0x04, 0x00, 0x00,
}
