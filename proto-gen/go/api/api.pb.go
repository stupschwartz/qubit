// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import organizations "organizations"
import projects "projects"
import scenes "scenes"
import operators "github.com/stupschwartz/qubit/proto-gen/go/operators"
import image_sequences "image_sequences"
import images "github.com/stupschwartz/qubit/proto-gen/go/images"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Organization from public import organizations/organizations.proto
type Organization organizations.Organization

func (m *Organization) Reset()          { (*organizations.Organization)(m).Reset() }
func (m *Organization) String() string  { return (*organizations.Organization)(m).String() }
func (*Organization) ProtoMessage()     {}
func (m *Organization) GetId() string   { return (*organizations.Organization)(m).GetId() }
func (m *Organization) GetName() string { return (*organizations.Organization)(m).GetName() }

// ListOrganizationsRequest from public import organizations/organizations.proto
type ListOrganizationsRequest organizations.ListOrganizationsRequest

func (m *ListOrganizationsRequest) Reset() { (*organizations.ListOrganizationsRequest)(m).Reset() }
func (m *ListOrganizationsRequest) String() string {
	return (*organizations.ListOrganizationsRequest)(m).String()
}
func (*ListOrganizationsRequest) ProtoMessage() {}
func (m *ListOrganizationsRequest) GetPageSize() int32 {
	return (*organizations.ListOrganizationsRequest)(m).GetPageSize()
}
func (m *ListOrganizationsRequest) GetPageToken() string {
	return (*organizations.ListOrganizationsRequest)(m).GetPageToken()
}

// ListOrganizationsResponse from public import organizations/organizations.proto
type ListOrganizationsResponse organizations.ListOrganizationsResponse

func (m *ListOrganizationsResponse) Reset() { (*organizations.ListOrganizationsResponse)(m).Reset() }
func (m *ListOrganizationsResponse) String() string {
	return (*organizations.ListOrganizationsResponse)(m).String()
}
func (*ListOrganizationsResponse) ProtoMessage() {}
func (m *ListOrganizationsResponse) GetOrganizations() []*Organization {
	o := (*organizations.ListOrganizationsResponse)(m).GetOrganizations()
	if o == nil {
		return nil
	}
	s := make([]*Organization, len(o))
	for i, x := range o {
		s[i] = (*Organization)(x)
	}
	return s
}
func (m *ListOrganizationsResponse) GetNextPageToken() string {
	return (*organizations.ListOrganizationsResponse)(m).GetNextPageToken()
}

// GetOrganizationRequest from public import organizations/organizations.proto
type GetOrganizationRequest organizations.GetOrganizationRequest

func (m *GetOrganizationRequest) Reset() { (*organizations.GetOrganizationRequest)(m).Reset() }
func (m *GetOrganizationRequest) String() string {
	return (*organizations.GetOrganizationRequest)(m).String()
}
func (*GetOrganizationRequest) ProtoMessage() {}
func (m *GetOrganizationRequest) GetId() string {
	return (*organizations.GetOrganizationRequest)(m).GetId()
}

// CreateOrganizationRequest from public import organizations/organizations.proto
type CreateOrganizationRequest organizations.CreateOrganizationRequest

func (m *CreateOrganizationRequest) Reset() { (*organizations.CreateOrganizationRequest)(m).Reset() }
func (m *CreateOrganizationRequest) String() string {
	return (*organizations.CreateOrganizationRequest)(m).String()
}
func (*CreateOrganizationRequest) ProtoMessage() {}
func (m *CreateOrganizationRequest) GetOrganization() *Organization {
	return (*Organization)((*organizations.CreateOrganizationRequest)(m).GetOrganization())
}

// UpdateOrganizationRequest from public import organizations/organizations.proto
type UpdateOrganizationRequest organizations.UpdateOrganizationRequest

func (m *UpdateOrganizationRequest) Reset() { (*organizations.UpdateOrganizationRequest)(m).Reset() }
func (m *UpdateOrganizationRequest) String() string {
	return (*organizations.UpdateOrganizationRequest)(m).String()
}
func (*UpdateOrganizationRequest) ProtoMessage() {}
func (m *UpdateOrganizationRequest) GetId() string {
	return (*organizations.UpdateOrganizationRequest)(m).GetId()
}
func (m *UpdateOrganizationRequest) GetOrganization() *Organization {
	return (*Organization)((*organizations.UpdateOrganizationRequest)(m).GetOrganization())
}

// DeleteOrganizationRequest from public import organizations/organizations.proto
type DeleteOrganizationRequest organizations.DeleteOrganizationRequest

func (m *DeleteOrganizationRequest) Reset() { (*organizations.DeleteOrganizationRequest)(m).Reset() }
func (m *DeleteOrganizationRequest) String() string {
	return (*organizations.DeleteOrganizationRequest)(m).String()
}
func (*DeleteOrganizationRequest) ProtoMessage() {}
func (m *DeleteOrganizationRequest) GetId() string {
	return (*organizations.DeleteOrganizationRequest)(m).GetId()
}

// Project from public import projects/projects.proto
type Project projects.Project

func (m *Project) Reset()                    { (*projects.Project)(m).Reset() }
func (m *Project) String() string            { return (*projects.Project)(m).String() }
func (*Project) ProtoMessage()               {}
func (m *Project) GetId() string             { return (*projects.Project)(m).GetId() }
func (m *Project) GetOrganizationId() string { return (*projects.Project)(m).GetOrganizationId() }
func (m *Project) GetName() string           { return (*projects.Project)(m).GetName() }

// ListProjectsRequest from public import projects/projects.proto
type ListProjectsRequest projects.ListProjectsRequest

func (m *ListProjectsRequest) Reset()         { (*projects.ListProjectsRequest)(m).Reset() }
func (m *ListProjectsRequest) String() string { return (*projects.ListProjectsRequest)(m).String() }
func (*ListProjectsRequest) ProtoMessage()    {}
func (m *ListProjectsRequest) GetPageSize() int32 {
	return (*projects.ListProjectsRequest)(m).GetPageSize()
}
func (m *ListProjectsRequest) GetPageToken() string {
	return (*projects.ListProjectsRequest)(m).GetPageToken()
}

// ListProjectsResponse from public import projects/projects.proto
type ListProjectsResponse projects.ListProjectsResponse

func (m *ListProjectsResponse) Reset()         { (*projects.ListProjectsResponse)(m).Reset() }
func (m *ListProjectsResponse) String() string { return (*projects.ListProjectsResponse)(m).String() }
func (*ListProjectsResponse) ProtoMessage()    {}
func (m *ListProjectsResponse) GetProjects() []*Project {
	o := (*projects.ListProjectsResponse)(m).GetProjects()
	if o == nil {
		return nil
	}
	s := make([]*Project, len(o))
	for i, x := range o {
		s[i] = (*Project)(x)
	}
	return s
}
func (m *ListProjectsResponse) GetNextPageToken() string {
	return (*projects.ListProjectsResponse)(m).GetNextPageToken()
}

// GetProjectRequest from public import projects/projects.proto
type GetProjectRequest projects.GetProjectRequest

func (m *GetProjectRequest) Reset()         { (*projects.GetProjectRequest)(m).Reset() }
func (m *GetProjectRequest) String() string { return (*projects.GetProjectRequest)(m).String() }
func (*GetProjectRequest) ProtoMessage()    {}
func (m *GetProjectRequest) GetId() string  { return (*projects.GetProjectRequest)(m).GetId() }

// CreateProjectRequest from public import projects/projects.proto
type CreateProjectRequest projects.CreateProjectRequest

func (m *CreateProjectRequest) Reset()         { (*projects.CreateProjectRequest)(m).Reset() }
func (m *CreateProjectRequest) String() string { return (*projects.CreateProjectRequest)(m).String() }
func (*CreateProjectRequest) ProtoMessage()    {}
func (m *CreateProjectRequest) GetProject() *Project {
	return (*Project)((*projects.CreateProjectRequest)(m).GetProject())
}

// UpdateProjectRequest from public import projects/projects.proto
type UpdateProjectRequest projects.UpdateProjectRequest

func (m *UpdateProjectRequest) Reset()         { (*projects.UpdateProjectRequest)(m).Reset() }
func (m *UpdateProjectRequest) String() string { return (*projects.UpdateProjectRequest)(m).String() }
func (*UpdateProjectRequest) ProtoMessage()    {}
func (m *UpdateProjectRequest) GetId() string  { return (*projects.UpdateProjectRequest)(m).GetId() }
func (m *UpdateProjectRequest) GetProject() *Project {
	return (*Project)((*projects.UpdateProjectRequest)(m).GetProject())
}

// DeleteProjectRequest from public import projects/projects.proto
type DeleteProjectRequest projects.DeleteProjectRequest

func (m *DeleteProjectRequest) Reset()         { (*projects.DeleteProjectRequest)(m).Reset() }
func (m *DeleteProjectRequest) String() string { return (*projects.DeleteProjectRequest)(m).String() }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (m *DeleteProjectRequest) GetId() string  { return (*projects.DeleteProjectRequest)(m).GetId() }

// Scene from public import scenes/scenes.proto
type Scene scenes.Scene

func (m *Scene) Reset()               { (*scenes.Scene)(m).Reset() }
func (m *Scene) String() string       { return (*scenes.Scene)(m).String() }
func (*Scene) ProtoMessage()          {}
func (m *Scene) GetId() string        { return (*scenes.Scene)(m).GetId() }
func (m *Scene) GetProjectId() string { return (*scenes.Scene)(m).GetProjectId() }
func (m *Scene) GetName() string      { return (*scenes.Scene)(m).GetName() }

// ListScenesRequest from public import scenes/scenes.proto
type ListScenesRequest scenes.ListScenesRequest

func (m *ListScenesRequest) Reset()             { (*scenes.ListScenesRequest)(m).Reset() }
func (m *ListScenesRequest) String() string     { return (*scenes.ListScenesRequest)(m).String() }
func (*ListScenesRequest) ProtoMessage()        {}
func (m *ListScenesRequest) GetPageSize() int32 { return (*scenes.ListScenesRequest)(m).GetPageSize() }
func (m *ListScenesRequest) GetPageToken() string {
	return (*scenes.ListScenesRequest)(m).GetPageToken()
}

// ListScenesResponse from public import scenes/scenes.proto
type ListScenesResponse scenes.ListScenesResponse

func (m *ListScenesResponse) Reset()         { (*scenes.ListScenesResponse)(m).Reset() }
func (m *ListScenesResponse) String() string { return (*scenes.ListScenesResponse)(m).String() }
func (*ListScenesResponse) ProtoMessage()    {}
func (m *ListScenesResponse) GetScenes() []*Scene {
	o := (*scenes.ListScenesResponse)(m).GetScenes()
	if o == nil {
		return nil
	}
	s := make([]*Scene, len(o))
	for i, x := range o {
		s[i] = (*Scene)(x)
	}
	return s
}
func (m *ListScenesResponse) GetNextPageToken() string {
	return (*scenes.ListScenesResponse)(m).GetNextPageToken()
}

// GetSceneRequest from public import scenes/scenes.proto
type GetSceneRequest scenes.GetSceneRequest

func (m *GetSceneRequest) Reset()         { (*scenes.GetSceneRequest)(m).Reset() }
func (m *GetSceneRequest) String() string { return (*scenes.GetSceneRequest)(m).String() }
func (*GetSceneRequest) ProtoMessage()    {}
func (m *GetSceneRequest) GetId() string  { return (*scenes.GetSceneRequest)(m).GetId() }

// CreateSceneRequest from public import scenes/scenes.proto
type CreateSceneRequest scenes.CreateSceneRequest

func (m *CreateSceneRequest) Reset()         { (*scenes.CreateSceneRequest)(m).Reset() }
func (m *CreateSceneRequest) String() string { return (*scenes.CreateSceneRequest)(m).String() }
func (*CreateSceneRequest) ProtoMessage()    {}
func (m *CreateSceneRequest) GetScene() *Scene {
	return (*Scene)((*scenes.CreateSceneRequest)(m).GetScene())
}

// UpdateSceneRequest from public import scenes/scenes.proto
type UpdateSceneRequest scenes.UpdateSceneRequest

func (m *UpdateSceneRequest) Reset()         { (*scenes.UpdateSceneRequest)(m).Reset() }
func (m *UpdateSceneRequest) String() string { return (*scenes.UpdateSceneRequest)(m).String() }
func (*UpdateSceneRequest) ProtoMessage()    {}
func (m *UpdateSceneRequest) GetId() string  { return (*scenes.UpdateSceneRequest)(m).GetId() }
func (m *UpdateSceneRequest) GetScene() *Scene {
	return (*Scene)((*scenes.UpdateSceneRequest)(m).GetScene())
}

// DeleteSceneRequest from public import scenes/scenes.proto
type DeleteSceneRequest scenes.DeleteSceneRequest

func (m *DeleteSceneRequest) Reset()         { (*scenes.DeleteSceneRequest)(m).Reset() }
func (m *DeleteSceneRequest) String() string { return (*scenes.DeleteSceneRequest)(m).String() }
func (*DeleteSceneRequest) ProtoMessage()    {}
func (m *DeleteSceneRequest) GetId() string  { return (*scenes.DeleteSceneRequest)(m).GetId() }

// Component from public import operators/operators.proto
type Component operators.Component

func (m *Component) Reset()            { (*operators.Component)(m).Reset() }
func (m *Component) String() string    { return (*operators.Component)(m).String() }
func (*Component) ProtoMessage()       {}
func (m *Component) GetName() string   { return (*operators.Component)(m).GetName() }
func (m *Component) GetValue() float64 { return (*operators.Component)(m).GetValue() }

// Parameter from public import operators/operators.proto
type Parameter operators.Parameter

func (m *Parameter) Reset()          { (*operators.Parameter)(m).Reset() }
func (m *Parameter) String() string  { return (*operators.Parameter)(m).String() }
func (*Parameter) ProtoMessage()     {}
func (m *Parameter) GetName() string { return (*operators.Parameter)(m).GetName() }
func (m *Parameter) GetLabels() map[string]string {
	o := (*operators.Parameter)(m).GetLabels()
	if o == nil {
		return nil
	}
	s := make(map[string]string, len(o))
	for k, v := range o {
		s[k] = (string)(v)
	}
	return s
}
func (m *Parameter) GetComponents() []*Component {
	o := (*operators.Parameter)(m).GetComponents()
	if o == nil {
		return nil
	}
	s := make([]*Component, len(o))
	for i, x := range o {
		s[i] = (*Component)(x)
	}
	return s
}

// Operator from public import operators/operators.proto
type Operator operators.Operator

func (m *Operator) Reset()              { (*operators.Operator)(m).Reset() }
func (m *Operator) String() string      { return (*operators.Operator)(m).String() }
func (*Operator) ProtoMessage()         {}
func (m *Operator) GetId() string       { return (*operators.Operator)(m).GetId() }
func (m *Operator) GetSceneId() string  { return (*operators.Operator)(m).GetSceneId() }
func (m *Operator) GetType() string     { return (*operators.Operator)(m).GetType() }
func (m *Operator) GetName() string     { return (*operators.Operator)(m).GetName() }
func (m *Operator) GetContext() string  { return (*operators.Operator)(m).GetContext() }
func (m *Operator) GetInputs() []string { return (*operators.Operator)(m).GetInputs() }
func (m *Operator) GetParameters() []*Parameter {
	o := (*operators.Operator)(m).GetParameters()
	if o == nil {
		return nil
	}
	s := make([]*Parameter, len(o))
	for i, x := range o {
		s[i] = (*Parameter)(x)
	}
	return s
}

// ListOperatorsRequest from public import operators/operators.proto
type ListOperatorsRequest operators.ListOperatorsRequest

func (m *ListOperatorsRequest) Reset()         { (*operators.ListOperatorsRequest)(m).Reset() }
func (m *ListOperatorsRequest) String() string { return (*operators.ListOperatorsRequest)(m).String() }
func (*ListOperatorsRequest) ProtoMessage()    {}
func (m *ListOperatorsRequest) GetPageSize() int32 {
	return (*operators.ListOperatorsRequest)(m).GetPageSize()
}
func (m *ListOperatorsRequest) GetPageToken() string {
	return (*operators.ListOperatorsRequest)(m).GetPageToken()
}

// ListOperatorsResponse from public import operators/operators.proto
type ListOperatorsResponse operators.ListOperatorsResponse

func (m *ListOperatorsResponse) Reset()         { (*operators.ListOperatorsResponse)(m).Reset() }
func (m *ListOperatorsResponse) String() string { return (*operators.ListOperatorsResponse)(m).String() }
func (*ListOperatorsResponse) ProtoMessage()    {}
func (m *ListOperatorsResponse) GetOperators() []*Operator {
	o := (*operators.ListOperatorsResponse)(m).GetOperators()
	if o == nil {
		return nil
	}
	s := make([]*Operator, len(o))
	for i, x := range o {
		s[i] = (*Operator)(x)
	}
	return s
}
func (m *ListOperatorsResponse) GetNextPageToken() string {
	return (*operators.ListOperatorsResponse)(m).GetNextPageToken()
}

// GetOperatorRequest from public import operators/operators.proto
type GetOperatorRequest operators.GetOperatorRequest

func (m *GetOperatorRequest) Reset()         { (*operators.GetOperatorRequest)(m).Reset() }
func (m *GetOperatorRequest) String() string { return (*operators.GetOperatorRequest)(m).String() }
func (*GetOperatorRequest) ProtoMessage()    {}
func (m *GetOperatorRequest) GetId() string  { return (*operators.GetOperatorRequest)(m).GetId() }

// CreateOperatorRequest from public import operators/operators.proto
type CreateOperatorRequest operators.CreateOperatorRequest

func (m *CreateOperatorRequest) Reset()         { (*operators.CreateOperatorRequest)(m).Reset() }
func (m *CreateOperatorRequest) String() string { return (*operators.CreateOperatorRequest)(m).String() }
func (*CreateOperatorRequest) ProtoMessage()    {}
func (m *CreateOperatorRequest) GetOperator() *Operator {
	return (*Operator)((*operators.CreateOperatorRequest)(m).GetOperator())
}

// UpdateOperatorRequest from public import operators/operators.proto
type UpdateOperatorRequest operators.UpdateOperatorRequest

func (m *UpdateOperatorRequest) Reset()         { (*operators.UpdateOperatorRequest)(m).Reset() }
func (m *UpdateOperatorRequest) String() string { return (*operators.UpdateOperatorRequest)(m).String() }
func (*UpdateOperatorRequest) ProtoMessage()    {}
func (m *UpdateOperatorRequest) GetId() string  { return (*operators.UpdateOperatorRequest)(m).GetId() }
func (m *UpdateOperatorRequest) GetOperator() *Operator {
	return (*Operator)((*operators.UpdateOperatorRequest)(m).GetOperator())
}

// DeleteOperatorRequest from public import operators/operators.proto
type DeleteOperatorRequest operators.DeleteOperatorRequest

func (m *DeleteOperatorRequest) Reset()         { (*operators.DeleteOperatorRequest)(m).Reset() }
func (m *DeleteOperatorRequest) String() string { return (*operators.DeleteOperatorRequest)(m).String() }
func (*DeleteOperatorRequest) ProtoMessage()    {}
func (m *DeleteOperatorRequest) GetId() string  { return (*operators.DeleteOperatorRequest)(m).GetId() }

// RenderOperatorRequest from public import operators/operators.proto
type RenderOperatorRequest operators.RenderOperatorRequest

func (m *RenderOperatorRequest) Reset()         { (*operators.RenderOperatorRequest)(m).Reset() }
func (m *RenderOperatorRequest) String() string { return (*operators.RenderOperatorRequest)(m).String() }
func (*RenderOperatorRequest) ProtoMessage()    {}
func (m *RenderOperatorRequest) GetId() string  { return (*operators.RenderOperatorRequest)(m).GetId() }
func (m *RenderOperatorRequest) GetFrame() int32 {
	return (*operators.RenderOperatorRequest)(m).GetFrame()
}

// RenderOperatorResponse from public import operators/operators.proto
type RenderOperatorResponse operators.RenderOperatorResponse

func (m *RenderOperatorResponse) Reset() { (*operators.RenderOperatorResponse)(m).Reset() }
func (m *RenderOperatorResponse) String() string {
	return (*operators.RenderOperatorResponse)(m).String()
}
func (*RenderOperatorResponse) ProtoMessage() {}
func (m *RenderOperatorResponse) GetResultUrl() string {
	return (*operators.RenderOperatorResponse)(m).GetResultUrl()
}
func (m *RenderOperatorResponse) GetResultType() string {
	return (*operators.RenderOperatorResponse)(m).GetResultType()
}

// ImageSequence from public import image_sequences/image_sequences.proto
type ImageSequence image_sequences.ImageSequence

func (m *ImageSequence) Reset()         { (*image_sequences.ImageSequence)(m).Reset() }
func (m *ImageSequence) String() string { return (*image_sequences.ImageSequence)(m).String() }
func (*ImageSequence) ProtoMessage()    {}
func (m *ImageSequence) GetId() string  { return (*image_sequences.ImageSequence)(m).GetId() }
func (m *ImageSequence) GetProjectId() string {
	return (*image_sequences.ImageSequence)(m).GetProjectId()
}
func (m *ImageSequence) GetName() string { return (*image_sequences.ImageSequence)(m).GetName() }

// ListImageSequencesRequest from public import image_sequences/image_sequences.proto
type ListImageSequencesRequest image_sequences.ListImageSequencesRequest

func (m *ListImageSequencesRequest) Reset() { (*image_sequences.ListImageSequencesRequest)(m).Reset() }
func (m *ListImageSequencesRequest) String() string {
	return (*image_sequences.ListImageSequencesRequest)(m).String()
}
func (*ListImageSequencesRequest) ProtoMessage() {}
func (m *ListImageSequencesRequest) GetPageSize() int32 {
	return (*image_sequences.ListImageSequencesRequest)(m).GetPageSize()
}
func (m *ListImageSequencesRequest) GetPageToken() string {
	return (*image_sequences.ListImageSequencesRequest)(m).GetPageToken()
}

// ListImageSequencesResponse from public import image_sequences/image_sequences.proto
type ListImageSequencesResponse image_sequences.ListImageSequencesResponse

func (m *ListImageSequencesResponse) Reset() { (*image_sequences.ListImageSequencesResponse)(m).Reset() }
func (m *ListImageSequencesResponse) String() string {
	return (*image_sequences.ListImageSequencesResponse)(m).String()
}
func (*ListImageSequencesResponse) ProtoMessage() {}
func (m *ListImageSequencesResponse) GetImageSequences() []*ImageSequence {
	o := (*image_sequences.ListImageSequencesResponse)(m).GetImageSequences()
	if o == nil {
		return nil
	}
	s := make([]*ImageSequence, len(o))
	for i, x := range o {
		s[i] = (*ImageSequence)(x)
	}
	return s
}
func (m *ListImageSequencesResponse) GetNextPageToken() string {
	return (*image_sequences.ListImageSequencesResponse)(m).GetNextPageToken()
}

// GetImageSequenceRequest from public import image_sequences/image_sequences.proto
type GetImageSequenceRequest image_sequences.GetImageSequenceRequest

func (m *GetImageSequenceRequest) Reset() { (*image_sequences.GetImageSequenceRequest)(m).Reset() }
func (m *GetImageSequenceRequest) String() string {
	return (*image_sequences.GetImageSequenceRequest)(m).String()
}
func (*GetImageSequenceRequest) ProtoMessage() {}
func (m *GetImageSequenceRequest) GetId() string {
	return (*image_sequences.GetImageSequenceRequest)(m).GetId()
}

// CreateImageSequenceRequest from public import image_sequences/image_sequences.proto
type CreateImageSequenceRequest image_sequences.CreateImageSequenceRequest

func (m *CreateImageSequenceRequest) Reset() { (*image_sequences.CreateImageSequenceRequest)(m).Reset() }
func (m *CreateImageSequenceRequest) String() string {
	return (*image_sequences.CreateImageSequenceRequest)(m).String()
}
func (*CreateImageSequenceRequest) ProtoMessage() {}
func (m *CreateImageSequenceRequest) GetImageSequence() *ImageSequence {
	return (*ImageSequence)((*image_sequences.CreateImageSequenceRequest)(m).GetImageSequence())
}

// UpdateImageSequenceRequest from public import image_sequences/image_sequences.proto
type UpdateImageSequenceRequest image_sequences.UpdateImageSequenceRequest

func (m *UpdateImageSequenceRequest) Reset() { (*image_sequences.UpdateImageSequenceRequest)(m).Reset() }
func (m *UpdateImageSequenceRequest) String() string {
	return (*image_sequences.UpdateImageSequenceRequest)(m).String()
}
func (*UpdateImageSequenceRequest) ProtoMessage() {}
func (m *UpdateImageSequenceRequest) GetId() string {
	return (*image_sequences.UpdateImageSequenceRequest)(m).GetId()
}
func (m *UpdateImageSequenceRequest) GetImageSequence() *ImageSequence {
	return (*ImageSequence)((*image_sequences.UpdateImageSequenceRequest)(m).GetImageSequence())
}

// DeleteImageSequenceRequest from public import image_sequences/image_sequences.proto
type DeleteImageSequenceRequest image_sequences.DeleteImageSequenceRequest

func (m *DeleteImageSequenceRequest) Reset() { (*image_sequences.DeleteImageSequenceRequest)(m).Reset() }
func (m *DeleteImageSequenceRequest) String() string {
	return (*image_sequences.DeleteImageSequenceRequest)(m).String()
}
func (*DeleteImageSequenceRequest) ProtoMessage() {}
func (m *DeleteImageSequenceRequest) GetId() string {
	return (*image_sequences.DeleteImageSequenceRequest)(m).GetId()
}

// Image from public import images/images.proto
type Image images.Image

func (m *Image) Reset()                     { (*images.Image)(m).Reset() }
func (m *Image) String() string             { return (*images.Image)(m).String() }
func (*Image) ProtoMessage()                {}
func (m *Image) GetId() string              { return (*images.Image)(m).GetId() }
func (m *Image) GetImageSequenceId() string { return (*images.Image)(m).GetImageSequenceId() }
func (m *Image) GetName() string            { return (*images.Image)(m).GetName() }
func (m *Image) GetWidth() int32            { return (*images.Image)(m).GetWidth() }
func (m *Image) GetHeight() int32           { return (*images.Image)(m).GetHeight() }
func (m *Image) GetLabels() map[string]string {
	o := (*images.Image)(m).GetLabels()
	if o == nil {
		return nil
	}
	s := make(map[string]string, len(o))
	for k, v := range o {
		s[k] = (string)(v)
	}
	return s
}
func (m *Image) GetPlanes() []*Plane {
	o := (*images.Image)(m).GetPlanes()
	if o == nil {
		return nil
	}
	s := make([]*Plane, len(o))
	for i, x := range o {
		s[i] = (*Plane)(x)
	}
	return s
}

// Plane from public import images/images.proto
type Plane images.Plane

func (m *Plane) Reset()           { (*images.Plane)(m).Reset() }
func (m *Plane) String() string   { return (*images.Plane)(m).String() }
func (*Plane) ProtoMessage()      {}
func (m *Plane) GetName() string  { return (*images.Plane)(m).GetName() }
func (m *Plane) GetWidth() int32  { return (*images.Plane)(m).GetWidth() }
func (m *Plane) GetHeight() int32 { return (*images.Plane)(m).GetHeight() }
func (m *Plane) GetLabels() map[string]string {
	o := (*images.Plane)(m).GetLabels()
	if o == nil {
		return nil
	}
	s := make(map[string]string, len(o))
	for k, v := range o {
		s[k] = (string)(v)
	}
	return s
}
func (m *Plane) GetChannels() []*Channel {
	o := (*images.Plane)(m).GetChannels()
	if o == nil {
		return nil
	}
	s := make([]*Channel, len(o))
	for i, x := range o {
		s[i] = (*Channel)(x)
	}
	return s
}

// Channel from public import images/images.proto
type Channel images.Channel

func (m *Channel) Reset()          { (*images.Channel)(m).Reset() }
func (m *Channel) String() string  { return (*images.Channel)(m).String() }
func (*Channel) ProtoMessage()     {}
func (m *Channel) GetName() string { return (*images.Channel)(m).GetName() }
func (m *Channel) GetRows() []*Row {
	o := (*images.Channel)(m).GetRows()
	if o == nil {
		return nil
	}
	s := make([]*Row, len(o))
	for i, x := range o {
		s[i] = (*Row)(x)
	}
	return s
}

// Row from public import images/images.proto
type Row images.Row

func (m *Row) Reset()             { (*images.Row)(m).Reset() }
func (m *Row) String() string     { return (*images.Row)(m).String() }
func (*Row) ProtoMessage()        {}
func (m *Row) GetData() []float64 { return (*images.Row)(m).GetData() }

// ListImagesRequest from public import images/images.proto
type ListImagesRequest images.ListImagesRequest

func (m *ListImagesRequest) Reset()             { (*images.ListImagesRequest)(m).Reset() }
func (m *ListImagesRequest) String() string     { return (*images.ListImagesRequest)(m).String() }
func (*ListImagesRequest) ProtoMessage()        {}
func (m *ListImagesRequest) GetPageSize() int32 { return (*images.ListImagesRequest)(m).GetPageSize() }
func (m *ListImagesRequest) GetPageToken() string {
	return (*images.ListImagesRequest)(m).GetPageToken()
}

// ListImagesResponse from public import images/images.proto
type ListImagesResponse images.ListImagesResponse

func (m *ListImagesResponse) Reset()         { (*images.ListImagesResponse)(m).Reset() }
func (m *ListImagesResponse) String() string { return (*images.ListImagesResponse)(m).String() }
func (*ListImagesResponse) ProtoMessage()    {}
func (m *ListImagesResponse) GetImages() []*Image {
	o := (*images.ListImagesResponse)(m).GetImages()
	if o == nil {
		return nil
	}
	s := make([]*Image, len(o))
	for i, x := range o {
		s[i] = (*Image)(x)
	}
	return s
}
func (m *ListImagesResponse) GetNextPageToken() string {
	return (*images.ListImagesResponse)(m).GetNextPageToken()
}

// GetImageRequest from public import images/images.proto
type GetImageRequest images.GetImageRequest

func (m *GetImageRequest) Reset()         { (*images.GetImageRequest)(m).Reset() }
func (m *GetImageRequest) String() string { return (*images.GetImageRequest)(m).String() }
func (*GetImageRequest) ProtoMessage()    {}
func (m *GetImageRequest) GetId() string  { return (*images.GetImageRequest)(m).GetId() }

// CreateImageRequest from public import images/images.proto
type CreateImageRequest images.CreateImageRequest

func (m *CreateImageRequest) Reset()         { (*images.CreateImageRequest)(m).Reset() }
func (m *CreateImageRequest) String() string { return (*images.CreateImageRequest)(m).String() }
func (*CreateImageRequest) ProtoMessage()    {}
func (m *CreateImageRequest) GetImage() *Image {
	return (*Image)((*images.CreateImageRequest)(m).GetImage())
}

// UpdateImageRequest from public import images/images.proto
type UpdateImageRequest images.UpdateImageRequest

func (m *UpdateImageRequest) Reset()         { (*images.UpdateImageRequest)(m).Reset() }
func (m *UpdateImageRequest) String() string { return (*images.UpdateImageRequest)(m).String() }
func (*UpdateImageRequest) ProtoMessage()    {}
func (m *UpdateImageRequest) GetId() string  { return (*images.UpdateImageRequest)(m).GetId() }
func (m *UpdateImageRequest) GetImage() *Image {
	return (*Image)((*images.UpdateImageRequest)(m).GetImage())
}

// DeleteImageRequest from public import images/images.proto
type DeleteImageRequest images.DeleteImageRequest

func (m *DeleteImageRequest) Reset()         { (*images.DeleteImageRequest)(m).Reset() }
func (m *DeleteImageRequest) String() string { return (*images.DeleteImageRequest)(m).String() }
func (*DeleteImageRequest) ProtoMessage()    {}
func (m *DeleteImageRequest) GetId() string  { return (*images.DeleteImageRequest)(m).GetId() }

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x07, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x96, 0xe4, 0xfc, 0xa2, 0x54, 0x29, 0xc5,
	0xfc, 0xa2, 0xf4, 0xc4, 0xbc, 0xcc, 0xaa, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x7d, 0x14, 0x1e,
	0x44, 0xa1, 0x94, 0x38, 0x90, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x87, 0x31, 0xa0, 0x12, 0xc2,
	0xc5, 0xc9, 0xa9, 0x79, 0xa9, 0xc5, 0xfa, 0x10, 0x0a, 0x2a, 0x28, 0x99, 0x5f, 0x90, 0x5a, 0x94,
	0x58, 0x92, 0x5f, 0x04, 0x34, 0x0c, 0xc6, 0x82, 0x4a, 0xa9, 0x66, 0xe6, 0x26, 0xa6, 0xa7, 0xc6,
	0x17, 0xa7, 0x16, 0x96, 0xa6, 0xe6, 0x25, 0x03, 0x35, 0xa2, 0xf1, 0x61, 0xc6, 0x82, 0x85, 0xa1,
	0xb2, 0x50, 0xc1, 0x00, 0x86, 0x00, 0xc6, 0x00, 0xa6, 0x00, 0xe6, 0x00, 0x96, 0x00, 0xd6, 0x24,
	0x36, 0xb0, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x23, 0x40, 0x5a, 0xd1, 0x00, 0x00,
	0x00,
}
