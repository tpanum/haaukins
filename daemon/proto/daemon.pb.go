// Copyright (c) 2018-2019 Aalborg University
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package daemon

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LoginUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserRequest) Reset()         { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()    {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{1}
}

func (m *LoginUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserRequest.Unmarshal(m, b)
}
func (m *LoginUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserRequest.Marshal(b, m, deterministic)
}
func (m *LoginUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserRequest.Merge(m, src)
}
func (m *LoginUserRequest) XXX_Size() int {
	return xxx_messageInfo_LoginUserRequest.Size(m)
}
func (m *LoginUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserRequest proto.InternalMessageInfo

func (m *LoginUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginUserResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUserResponse) Reset()         { *m = LoginUserResponse{} }
func (m *LoginUserResponse) String() string { return proto.CompactTextString(m) }
func (*LoginUserResponse) ProtoMessage()    {}
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{2}
}

func (m *LoginUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUserResponse.Unmarshal(m, b)
}
func (m *LoginUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUserResponse.Marshal(b, m, deterministic)
}
func (m *LoginUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUserResponse.Merge(m, src)
}
func (m *LoginUserResponse) XXX_Size() int {
	return xxx_messageInfo_LoginUserResponse.Size(m)
}
func (m *LoginUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUserResponse proto.InternalMessageInfo

func (m *LoginUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SignupUserRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupUserRequest) Reset()         { *m = SignupUserRequest{} }
func (m *SignupUserRequest) String() string { return proto.CompactTextString(m) }
func (*SignupUserRequest) ProtoMessage()    {}
func (*SignupUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{3}
}

func (m *SignupUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupUserRequest.Unmarshal(m, b)
}
func (m *SignupUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupUserRequest.Marshal(b, m, deterministic)
}
func (m *SignupUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupUserRequest.Merge(m, src)
}
func (m *SignupUserRequest) XXX_Size() int {
	return xxx_messageInfo_SignupUserRequest.Size(m)
}
func (m *SignupUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupUserRequest proto.InternalMessageInfo

func (m *SignupUserRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SignupUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignupUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type InviteUserRequest struct {
	SuperUser            bool     `protobuf:"varint,1,opt,name=super_user,json=superUser,proto3" json:"super_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteUserRequest) Reset()         { *m = InviteUserRequest{} }
func (m *InviteUserRequest) String() string { return proto.CompactTextString(m) }
func (*InviteUserRequest) ProtoMessage()    {}
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{4}
}

func (m *InviteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteUserRequest.Unmarshal(m, b)
}
func (m *InviteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteUserRequest.Marshal(b, m, deterministic)
}
func (m *InviteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteUserRequest.Merge(m, src)
}
func (m *InviteUserRequest) XXX_Size() int {
	return xxx_messageInfo_InviteUserRequest.Size(m)
}
func (m *InviteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InviteUserRequest proto.InternalMessageInfo

func (m *InviteUserRequest) GetSuperUser() bool {
	if m != nil {
		return m.SuperUser
	}
	return false
}

type InviteUserResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InviteUserResponse) Reset()         { *m = InviteUserResponse{} }
func (m *InviteUserResponse) String() string { return proto.CompactTextString(m) }
func (*InviteUserResponse) ProtoMessage()    {}
func (*InviteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{5}
}

func (m *InviteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InviteUserResponse.Unmarshal(m, b)
}
func (m *InviteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InviteUserResponse.Marshal(b, m, deterministic)
}
func (m *InviteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InviteUserResponse.Merge(m, src)
}
func (m *InviteUserResponse) XXX_Size() int {
	return xxx_messageInfo_InviteUserResponse.Size(m)
}
func (m *InviteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InviteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InviteUserResponse proto.InternalMessageInfo

func (m *InviteUserResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InviteUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateEventRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Frontends            []string `protobuf:"bytes,3,rep,name=frontends,proto3" json:"frontends,omitempty"`
	Exercises            []string `protobuf:"bytes,4,rep,name=exercises,proto3" json:"exercises,omitempty"`
	Available            int32    `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	Capacity             int32    `protobuf:"varint,6,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{6}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEventRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *CreateEventRequest) GetFrontends() []string {
	if m != nil {
		return m.Frontends
	}
	return nil
}

func (m *CreateEventRequest) GetExercises() []string {
	if m != nil {
		return m.Exercises
	}
	return nil
}

func (m *CreateEventRequest) GetAvailable() int32 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *CreateEventRequest) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type ListEventsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventsRequest) Reset()         { *m = ListEventsRequest{} }
func (m *ListEventsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEventsRequest) ProtoMessage()    {}
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{7}
}

func (m *ListEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventsRequest.Unmarshal(m, b)
}
func (m *ListEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventsRequest.Marshal(b, m, deterministic)
}
func (m *ListEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventsRequest.Merge(m, src)
}
func (m *ListEventsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEventsRequest.Size(m)
}
func (m *ListEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventsRequest proto.InternalMessageInfo

type ListEventsResponse struct {
	Events               []*ListEventsResponse_Events `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListEventsResponse) Reset()         { *m = ListEventsResponse{} }
func (m *ListEventsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEventsResponse) ProtoMessage()    {}
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{8}
}

func (m *ListEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventsResponse.Unmarshal(m, b)
}
func (m *ListEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventsResponse.Marshal(b, m, deterministic)
}
func (m *ListEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventsResponse.Merge(m, src)
}
func (m *ListEventsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEventsResponse.Size(m)
}
func (m *ListEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventsResponse proto.InternalMessageInfo

func (m *ListEventsResponse) GetEvents() []*ListEventsResponse_Events {
	if m != nil {
		return m.Events
	}
	return nil
}

type ListEventsResponse_Events struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TeamCount            int32    `protobuf:"varint,3,opt,name=teamCount,proto3" json:"teamCount,omitempty"`
	ExerciseCount        int32    `protobuf:"varint,4,opt,name=exerciseCount,proto3" json:"exerciseCount,omitempty"`
	Capacity             int32    `protobuf:"varint,5,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventsResponse_Events) Reset()         { *m = ListEventsResponse_Events{} }
func (m *ListEventsResponse_Events) String() string { return proto.CompactTextString(m) }
func (*ListEventsResponse_Events) ProtoMessage()    {}
func (*ListEventsResponse_Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{8, 0}
}

func (m *ListEventsResponse_Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventsResponse_Events.Unmarshal(m, b)
}
func (m *ListEventsResponse_Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventsResponse_Events.Marshal(b, m, deterministic)
}
func (m *ListEventsResponse_Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventsResponse_Events.Merge(m, src)
}
func (m *ListEventsResponse_Events) XXX_Size() int {
	return xxx_messageInfo_ListEventsResponse_Events.Size(m)
}
func (m *ListEventsResponse_Events) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventsResponse_Events.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventsResponse_Events proto.InternalMessageInfo

func (m *ListEventsResponse_Events) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ListEventsResponse_Events) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListEventsResponse_Events) GetTeamCount() int32 {
	if m != nil {
		return m.TeamCount
	}
	return 0
}

func (m *ListEventsResponse_Events) GetExerciseCount() int32 {
	if m != nil {
		return m.ExerciseCount
	}
	return 0
}

func (m *ListEventsResponse_Events) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type ListEventTeamsRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventTeamsRequest) Reset()         { *m = ListEventTeamsRequest{} }
func (m *ListEventTeamsRequest) String() string { return proto.CompactTextString(m) }
func (*ListEventTeamsRequest) ProtoMessage()    {}
func (*ListEventTeamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{9}
}

func (m *ListEventTeamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventTeamsRequest.Unmarshal(m, b)
}
func (m *ListEventTeamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventTeamsRequest.Marshal(b, m, deterministic)
}
func (m *ListEventTeamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventTeamsRequest.Merge(m, src)
}
func (m *ListEventTeamsRequest) XXX_Size() int {
	return xxx_messageInfo_ListEventTeamsRequest.Size(m)
}
func (m *ListEventTeamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventTeamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventTeamsRequest proto.InternalMessageInfo

func (m *ListEventTeamsRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ListEventTeamsResponse struct {
	Teams                []*ListEventTeamsResponse_Teams `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListEventTeamsResponse) Reset()         { *m = ListEventTeamsResponse{} }
func (m *ListEventTeamsResponse) String() string { return proto.CompactTextString(m) }
func (*ListEventTeamsResponse) ProtoMessage()    {}
func (*ListEventTeamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{10}
}

func (m *ListEventTeamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventTeamsResponse.Unmarshal(m, b)
}
func (m *ListEventTeamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventTeamsResponse.Marshal(b, m, deterministic)
}
func (m *ListEventTeamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventTeamsResponse.Merge(m, src)
}
func (m *ListEventTeamsResponse) XXX_Size() int {
	return xxx_messageInfo_ListEventTeamsResponse.Size(m)
}
func (m *ListEventTeamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventTeamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventTeamsResponse proto.InternalMessageInfo

func (m *ListEventTeamsResponse) GetTeams() []*ListEventTeamsResponse_Teams {
	if m != nil {
		return m.Teams
	}
	return nil
}

type ListEventTeamsResponse_Teams struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEventTeamsResponse_Teams) Reset()         { *m = ListEventTeamsResponse_Teams{} }
func (m *ListEventTeamsResponse_Teams) String() string { return proto.CompactTextString(m) }
func (*ListEventTeamsResponse_Teams) ProtoMessage()    {}
func (*ListEventTeamsResponse_Teams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{10, 0}
}

func (m *ListEventTeamsResponse_Teams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEventTeamsResponse_Teams.Unmarshal(m, b)
}
func (m *ListEventTeamsResponse_Teams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEventTeamsResponse_Teams.Marshal(b, m, deterministic)
}
func (m *ListEventTeamsResponse_Teams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEventTeamsResponse_Teams.Merge(m, src)
}
func (m *ListEventTeamsResponse_Teams) XXX_Size() int {
	return xxx_messageInfo_ListEventTeamsResponse_Teams.Size(m)
}
func (m *ListEventTeamsResponse_Teams) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEventTeamsResponse_Teams.DiscardUnknown(m)
}

var xxx_messageInfo_ListEventTeamsResponse_Teams proto.InternalMessageInfo

func (m *ListEventTeamsResponse_Teams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListEventTeamsResponse_Teams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListEventTeamsResponse_Teams) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RestartTeamLabRequest struct {
	EventTag             string   `protobuf:"bytes,1,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	LabTag               string   `protobuf:"bytes,2,opt,name=labTag,proto3" json:"labTag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartTeamLabRequest) Reset()         { *m = RestartTeamLabRequest{} }
func (m *RestartTeamLabRequest) String() string { return proto.CompactTextString(m) }
func (*RestartTeamLabRequest) ProtoMessage()    {}
func (*RestartTeamLabRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{11}
}

func (m *RestartTeamLabRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartTeamLabRequest.Unmarshal(m, b)
}
func (m *RestartTeamLabRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartTeamLabRequest.Marshal(b, m, deterministic)
}
func (m *RestartTeamLabRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartTeamLabRequest.Merge(m, src)
}
func (m *RestartTeamLabRequest) XXX_Size() int {
	return xxx_messageInfo_RestartTeamLabRequest.Size(m)
}
func (m *RestartTeamLabRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartTeamLabRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestartTeamLabRequest proto.InternalMessageInfo

func (m *RestartTeamLabRequest) GetEventTag() string {
	if m != nil {
		return m.EventTag
	}
	return ""
}

func (m *RestartTeamLabRequest) GetLabTag() string {
	if m != nil {
		return m.LabTag
	}
	return ""
}

type ResetExerciseRequest struct {
	ExerciseTag          string   `protobuf:"bytes,1,opt,name=exerciseTag,proto3" json:"exerciseTag,omitempty"`
	EventTag             string   `protobuf:"bytes,2,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	Teams                []*Team  `protobuf:"bytes,3,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetExerciseRequest) Reset()         { *m = ResetExerciseRequest{} }
func (m *ResetExerciseRequest) String() string { return proto.CompactTextString(m) }
func (*ResetExerciseRequest) ProtoMessage()    {}
func (*ResetExerciseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{12}
}

func (m *ResetExerciseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetExerciseRequest.Unmarshal(m, b)
}
func (m *ResetExerciseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetExerciseRequest.Marshal(b, m, deterministic)
}
func (m *ResetExerciseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetExerciseRequest.Merge(m, src)
}
func (m *ResetExerciseRequest) XXX_Size() int {
	return xxx_messageInfo_ResetExerciseRequest.Size(m)
}
func (m *ResetExerciseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetExerciseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetExerciseRequest proto.InternalMessageInfo

func (m *ResetExerciseRequest) GetExerciseTag() string {
	if m != nil {
		return m.ExerciseTag
	}
	return ""
}

func (m *ResetExerciseRequest) GetEventTag() string {
	if m != nil {
		return m.EventTag
	}
	return ""
}

func (m *ResetExerciseRequest) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type ListExercisesResponse struct {
	Exercises            []*ListExercisesResponse_Exercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListExercisesResponse) Reset()         { *m = ListExercisesResponse{} }
func (m *ListExercisesResponse) String() string { return proto.CompactTextString(m) }
func (*ListExercisesResponse) ProtoMessage()    {}
func (*ListExercisesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{13}
}

func (m *ListExercisesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExercisesResponse.Unmarshal(m, b)
}
func (m *ListExercisesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExercisesResponse.Marshal(b, m, deterministic)
}
func (m *ListExercisesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExercisesResponse.Merge(m, src)
}
func (m *ListExercisesResponse) XXX_Size() int {
	return xxx_messageInfo_ListExercisesResponse.Size(m)
}
func (m *ListExercisesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExercisesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListExercisesResponse proto.InternalMessageInfo

func (m *ListExercisesResponse) GetExercises() []*ListExercisesResponse_Exercise {
	if m != nil {
		return m.Exercises
	}
	return nil
}

type ListExercisesResponse_Exercise struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DockerImageCount     int32    `protobuf:"varint,3,opt,name=dockerImageCount,proto3" json:"dockerImageCount,omitempty"`
	VboxImageCount       int32    `protobuf:"varint,4,opt,name=vboxImageCount,proto3" json:"vboxImageCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListExercisesResponse_Exercise) Reset()         { *m = ListExercisesResponse_Exercise{} }
func (m *ListExercisesResponse_Exercise) String() string { return proto.CompactTextString(m) }
func (*ListExercisesResponse_Exercise) ProtoMessage()    {}
func (*ListExercisesResponse_Exercise) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{13, 0}
}

func (m *ListExercisesResponse_Exercise) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListExercisesResponse_Exercise.Unmarshal(m, b)
}
func (m *ListExercisesResponse_Exercise) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListExercisesResponse_Exercise.Marshal(b, m, deterministic)
}
func (m *ListExercisesResponse_Exercise) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListExercisesResponse_Exercise.Merge(m, src)
}
func (m *ListExercisesResponse_Exercise) XXX_Size() int {
	return xxx_messageInfo_ListExercisesResponse_Exercise.Size(m)
}
func (m *ListExercisesResponse_Exercise) XXX_DiscardUnknown() {
	xxx_messageInfo_ListExercisesResponse_Exercise.DiscardUnknown(m)
}

var xxx_messageInfo_ListExercisesResponse_Exercise proto.InternalMessageInfo

func (m *ListExercisesResponse_Exercise) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ListExercisesResponse_Exercise) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListExercisesResponse_Exercise) GetDockerImageCount() int32 {
	if m != nil {
		return m.DockerImageCount
	}
	return 0
}

func (m *ListExercisesResponse_Exercise) GetVboxImageCount() int32 {
	if m != nil {
		return m.VboxImageCount
	}
	return 0
}

type ResetTeamStatus struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetTeamStatus) Reset()         { *m = ResetTeamStatus{} }
func (m *ResetTeamStatus) String() string { return proto.CompactTextString(m) }
func (*ResetTeamStatus) ProtoMessage()    {}
func (*ResetTeamStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{14}
}

func (m *ResetTeamStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetTeamStatus.Unmarshal(m, b)
}
func (m *ResetTeamStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetTeamStatus.Marshal(b, m, deterministic)
}
func (m *ResetTeamStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetTeamStatus.Merge(m, src)
}
func (m *ResetTeamStatus) XXX_Size() int {
	return xxx_messageInfo_ResetTeamStatus.Size(m)
}
func (m *ResetTeamStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetTeamStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResetTeamStatus proto.InternalMessageInfo

func (m *ResetTeamStatus) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *ResetTeamStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type StopEventRequest struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopEventRequest) Reset()         { *m = StopEventRequest{} }
func (m *StopEventRequest) String() string { return proto.CompactTextString(m) }
func (*StopEventRequest) ProtoMessage()    {}
func (*StopEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{15}
}

func (m *StopEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopEventRequest.Unmarshal(m, b)
}
func (m *StopEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopEventRequest.Marshal(b, m, deterministic)
}
func (m *StopEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopEventRequest.Merge(m, src)
}
func (m *StopEventRequest) XXX_Size() int {
	return xxx_messageInfo_StopEventRequest.Size(m)
}
func (m *StopEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopEventRequest proto.InternalMessageInfo

func (m *StopEventRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type EventStatus struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventStatus) Reset()         { *m = EventStatus{} }
func (m *EventStatus) String() string { return proto.CompactTextString(m) }
func (*EventStatus) ProtoMessage()    {}
func (*EventStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{16}
}

func (m *EventStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStatus.Unmarshal(m, b)
}
func (m *EventStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStatus.Marshal(b, m, deterministic)
}
func (m *EventStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStatus.Merge(m, src)
}
func (m *EventStatus) XXX_Size() int {
	return xxx_messageInfo_EventStatus.Size(m)
}
func (m *EventStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventStatus proto.InternalMessageInfo

func (m *EventStatus) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *EventStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type MonitorHostResponse struct {
	MemoryPercent        float32  `protobuf:"fixed32,1,opt,name=MemoryPercent,proto3" json:"MemoryPercent,omitempty"`
	MemoryReadError      string   `protobuf:"bytes,2,opt,name=MemoryReadError,proto3" json:"MemoryReadError,omitempty"`
	CPUPercent           float32  `protobuf:"fixed32,3,opt,name=CPUPercent,proto3" json:"CPUPercent,omitempty"`
	CPUReadError         string   `protobuf:"bytes,4,opt,name=CPUReadError,proto3" json:"CPUReadError,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonitorHostResponse) Reset()         { *m = MonitorHostResponse{} }
func (m *MonitorHostResponse) String() string { return proto.CompactTextString(m) }
func (*MonitorHostResponse) ProtoMessage()    {}
func (*MonitorHostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{17}
}

func (m *MonitorHostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorHostResponse.Unmarshal(m, b)
}
func (m *MonitorHostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorHostResponse.Marshal(b, m, deterministic)
}
func (m *MonitorHostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorHostResponse.Merge(m, src)
}
func (m *MonitorHostResponse) XXX_Size() int {
	return xxx_messageInfo_MonitorHostResponse.Size(m)
}
func (m *MonitorHostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorHostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorHostResponse proto.InternalMessageInfo

func (m *MonitorHostResponse) GetMemoryPercent() float32 {
	if m != nil {
		return m.MemoryPercent
	}
	return 0
}

func (m *MonitorHostResponse) GetMemoryReadError() string {
	if m != nil {
		return m.MemoryReadError
	}
	return ""
}

func (m *MonitorHostResponse) GetCPUPercent() float32 {
	if m != nil {
		return m.CPUPercent
	}
	return 0
}

func (m *MonitorHostResponse) GetCPUReadError() string {
	if m != nil {
		return m.CPUReadError
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{18}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type VersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{19}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ListFrontendsResponse struct {
	Frontends            []*ListFrontendsResponse_Frontend `protobuf:"bytes,1,rep,name=frontends,proto3" json:"frontends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ListFrontendsResponse) Reset()         { *m = ListFrontendsResponse{} }
func (m *ListFrontendsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFrontendsResponse) ProtoMessage()    {}
func (*ListFrontendsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{20}
}

func (m *ListFrontendsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFrontendsResponse.Unmarshal(m, b)
}
func (m *ListFrontendsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFrontendsResponse.Marshal(b, m, deterministic)
}
func (m *ListFrontendsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFrontendsResponse.Merge(m, src)
}
func (m *ListFrontendsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFrontendsResponse.Size(m)
}
func (m *ListFrontendsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFrontendsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFrontendsResponse proto.InternalMessageInfo

func (m *ListFrontendsResponse) GetFrontends() []*ListFrontendsResponse_Frontend {
	if m != nil {
		return m.Frontends
	}
	return nil
}

type ListFrontendsResponse_Frontend struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	MemoryMB             int64    `protobuf:"varint,3,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	Cpu                  float32  `protobuf:"fixed32,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFrontendsResponse_Frontend) Reset()         { *m = ListFrontendsResponse_Frontend{} }
func (m *ListFrontendsResponse_Frontend) String() string { return proto.CompactTextString(m) }
func (*ListFrontendsResponse_Frontend) ProtoMessage()    {}
func (*ListFrontendsResponse_Frontend) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{20, 0}
}

func (m *ListFrontendsResponse_Frontend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFrontendsResponse_Frontend.Unmarshal(m, b)
}
func (m *ListFrontendsResponse_Frontend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFrontendsResponse_Frontend.Marshal(b, m, deterministic)
}
func (m *ListFrontendsResponse_Frontend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFrontendsResponse_Frontend.Merge(m, src)
}
func (m *ListFrontendsResponse_Frontend) XXX_Size() int {
	return xxx_messageInfo_ListFrontendsResponse_Frontend.Size(m)
}
func (m *ListFrontendsResponse_Frontend) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFrontendsResponse_Frontend.DiscardUnknown(m)
}

var xxx_messageInfo_ListFrontendsResponse_Frontend proto.InternalMessageInfo

func (m *ListFrontendsResponse_Frontend) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ListFrontendsResponse_Frontend) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListFrontendsResponse_Frontend) GetMemoryMB() int64 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

func (m *ListFrontendsResponse_Frontend) GetCpu() float32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

type ResetFrontendsRequest struct {
	EventTag             string   `protobuf:"bytes,2,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	Teams                []*Team  `protobuf:"bytes,3,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetFrontendsRequest) Reset()         { *m = ResetFrontendsRequest{} }
func (m *ResetFrontendsRequest) String() string { return proto.CompactTextString(m) }
func (*ResetFrontendsRequest) ProtoMessage()    {}
func (*ResetFrontendsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{21}
}

func (m *ResetFrontendsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetFrontendsRequest.Unmarshal(m, b)
}
func (m *ResetFrontendsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetFrontendsRequest.Marshal(b, m, deterministic)
}
func (m *ResetFrontendsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetFrontendsRequest.Merge(m, src)
}
func (m *ResetFrontendsRequest) XXX_Size() int {
	return xxx_messageInfo_ResetFrontendsRequest.Size(m)
}
func (m *ResetFrontendsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetFrontendsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResetFrontendsRequest proto.InternalMessageInfo

func (m *ResetFrontendsRequest) GetEventTag() string {
	if m != nil {
		return m.EventTag
	}
	return ""
}

func (m *ResetFrontendsRequest) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type SetFrontendMemoryRequest struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	MemoryMB             int64    `protobuf:"varint,2,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFrontendMemoryRequest) Reset()         { *m = SetFrontendMemoryRequest{} }
func (m *SetFrontendMemoryRequest) String() string { return proto.CompactTextString(m) }
func (*SetFrontendMemoryRequest) ProtoMessage()    {}
func (*SetFrontendMemoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{22}
}

func (m *SetFrontendMemoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFrontendMemoryRequest.Unmarshal(m, b)
}
func (m *SetFrontendMemoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFrontendMemoryRequest.Marshal(b, m, deterministic)
}
func (m *SetFrontendMemoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFrontendMemoryRequest.Merge(m, src)
}
func (m *SetFrontendMemoryRequest) XXX_Size() int {
	return xxx_messageInfo_SetFrontendMemoryRequest.Size(m)
}
func (m *SetFrontendMemoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFrontendMemoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFrontendMemoryRequest proto.InternalMessageInfo

func (m *SetFrontendMemoryRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SetFrontendMemoryRequest) GetMemoryMB() int64 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

type SetFrontendCpuRequest struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Cpu                  float32  `protobuf:"fixed32,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFrontendCpuRequest) Reset()         { *m = SetFrontendCpuRequest{} }
func (m *SetFrontendCpuRequest) String() string { return proto.CompactTextString(m) }
func (*SetFrontendCpuRequest) ProtoMessage()    {}
func (*SetFrontendCpuRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{23}
}

func (m *SetFrontendCpuRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFrontendCpuRequest.Unmarshal(m, b)
}
func (m *SetFrontendCpuRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFrontendCpuRequest.Marshal(b, m, deterministic)
}
func (m *SetFrontendCpuRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFrontendCpuRequest.Merge(m, src)
}
func (m *SetFrontendCpuRequest) XXX_Size() int {
	return xxx_messageInfo_SetFrontendCpuRequest.Size(m)
}
func (m *SetFrontendCpuRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFrontendCpuRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFrontendCpuRequest proto.InternalMessageInfo

func (m *SetFrontendCpuRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *SetFrontendCpuRequest) GetCpu() float32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

type GetTeamInfoRequest struct {
	TeamId               string   `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	EventTag             string   `protobuf:"bytes,2,opt,name=eventTag,proto3" json:"eventTag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamInfoRequest) Reset()         { *m = GetTeamInfoRequest{} }
func (m *GetTeamInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTeamInfoRequest) ProtoMessage()    {}
func (*GetTeamInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{24}
}

func (m *GetTeamInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamInfoRequest.Unmarshal(m, b)
}
func (m *GetTeamInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetTeamInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamInfoRequest.Merge(m, src)
}
func (m *GetTeamInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTeamInfoRequest.Size(m)
}
func (m *GetTeamInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamInfoRequest proto.InternalMessageInfo

func (m *GetTeamInfoRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *GetTeamInfoRequest) GetEventTag() string {
	if m != nil {
		return m.EventTag
	}
	return ""
}

type GetTeamInfoResponse struct {
	Instances            []*GetTeamInfoResponse_Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetTeamInfoResponse) Reset()         { *m = GetTeamInfoResponse{} }
func (m *GetTeamInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetTeamInfoResponse) ProtoMessage()    {}
func (*GetTeamInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{25}
}

func (m *GetTeamInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamInfoResponse.Unmarshal(m, b)
}
func (m *GetTeamInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetTeamInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamInfoResponse.Merge(m, src)
}
func (m *GetTeamInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetTeamInfoResponse.Size(m)
}
func (m *GetTeamInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamInfoResponse proto.InternalMessageInfo

func (m *GetTeamInfoResponse) GetInstances() []*GetTeamInfoResponse_Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type GetTeamInfoResponse_Instance struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	State                int32    `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamInfoResponse_Instance) Reset()         { *m = GetTeamInfoResponse_Instance{} }
func (m *GetTeamInfoResponse_Instance) String() string { return proto.CompactTextString(m) }
func (*GetTeamInfoResponse_Instance) ProtoMessage()    {}
func (*GetTeamInfoResponse_Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec90cbc4aa12fc6, []int{25, 0}
}

func (m *GetTeamInfoResponse_Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamInfoResponse_Instance.Unmarshal(m, b)
}
func (m *GetTeamInfoResponse_Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamInfoResponse_Instance.Marshal(b, m, deterministic)
}
func (m *GetTeamInfoResponse_Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamInfoResponse_Instance.Merge(m, src)
}
func (m *GetTeamInfoResponse_Instance) XXX_Size() int {
	return xxx_messageInfo_GetTeamInfoResponse_Instance.Size(m)
}
func (m *GetTeamInfoResponse_Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamInfoResponse_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamInfoResponse_Instance proto.InternalMessageInfo

func (m *GetTeamInfoResponse_Instance) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *GetTeamInfoResponse_Instance) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetTeamInfoResponse_Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetTeamInfoResponse_Instance) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*Team)(nil), "Team")
	proto.RegisterType((*LoginUserRequest)(nil), "LoginUserRequest")
	proto.RegisterType((*LoginUserResponse)(nil), "LoginUserResponse")
	proto.RegisterType((*SignupUserRequest)(nil), "SignupUserRequest")
	proto.RegisterType((*InviteUserRequest)(nil), "InviteUserRequest")
	proto.RegisterType((*InviteUserResponse)(nil), "InviteUserResponse")
	proto.RegisterType((*CreateEventRequest)(nil), "CreateEventRequest")
	proto.RegisterType((*ListEventsRequest)(nil), "ListEventsRequest")
	proto.RegisterType((*ListEventsResponse)(nil), "ListEventsResponse")
	proto.RegisterType((*ListEventsResponse_Events)(nil), "ListEventsResponse.Events")
	proto.RegisterType((*ListEventTeamsRequest)(nil), "ListEventTeamsRequest")
	proto.RegisterType((*ListEventTeamsResponse)(nil), "ListEventTeamsResponse")
	proto.RegisterType((*ListEventTeamsResponse_Teams)(nil), "ListEventTeamsResponse.Teams")
	proto.RegisterType((*RestartTeamLabRequest)(nil), "RestartTeamLabRequest")
	proto.RegisterType((*ResetExerciseRequest)(nil), "ResetExerciseRequest")
	proto.RegisterType((*ListExercisesResponse)(nil), "ListExercisesResponse")
	proto.RegisterType((*ListExercisesResponse_Exercise)(nil), "ListExercisesResponse.Exercise")
	proto.RegisterType((*ResetTeamStatus)(nil), "ResetTeamStatus")
	proto.RegisterType((*StopEventRequest)(nil), "StopEventRequest")
	proto.RegisterType((*EventStatus)(nil), "EventStatus")
	proto.RegisterType((*MonitorHostResponse)(nil), "MonitorHostResponse")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*VersionResponse)(nil), "VersionResponse")
	proto.RegisterType((*ListFrontendsResponse)(nil), "ListFrontendsResponse")
	proto.RegisterType((*ListFrontendsResponse_Frontend)(nil), "ListFrontendsResponse.Frontend")
	proto.RegisterType((*ResetFrontendsRequest)(nil), "ResetFrontendsRequest")
	proto.RegisterType((*SetFrontendMemoryRequest)(nil), "SetFrontendMemoryRequest")
	proto.RegisterType((*SetFrontendCpuRequest)(nil), "SetFrontendCpuRequest")
	proto.RegisterType((*GetTeamInfoRequest)(nil), "GetTeamInfoRequest")
	proto.RegisterType((*GetTeamInfoResponse)(nil), "GetTeamInfoResponse")
	proto.RegisterType((*GetTeamInfoResponse_Instance)(nil), "GetTeamInfoResponse.Instance")
}

func init() { proto.RegisterFile("daemon.proto", fileDescriptor_3ec90cbc4aa12fc6) }

var fileDescriptor_3ec90cbc4aa12fc6 = []byte{
	// 1207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0xb1, 0x13, 0x1f, 0xe7, 0x62, 0x8f, 0x13, 0x63, 0x96, 0x16, 0xa2, 0x51, 0x40,
	0x01, 0xa4, 0xa1, 0xa4, 0x15, 0x45, 0xa5, 0xa5, 0x2a, 0x26, 0x50, 0x43, 0x82, 0xa2, 0x4d, 0xc2,
	0x03, 0x12, 0x42, 0x6b, 0x7b, 0x1a, 0xad, 0x12, 0xef, 0x6e, 0x77, 0xc6, 0xa1, 0xe1, 0x07, 0xf0,
	0x88, 0xf8, 0x1b, 0xbc, 0x20, 0xde, 0x90, 0xf8, 0x27, 0x3c, 0xf1, 0x57, 0xd0, 0xdc, 0x76, 0x67,
	0x2f, 0x8e, 0xd4, 0xb7, 0x39, 0x67, 0xcf, 0x9c, 0xcb, 0x37, 0xe7, 0xb6, 0xb0, 0x3e, 0xf3, 0xe9,
	0x3c, 0x0a, 0x49, 0x9c, 0x44, 0x3c, 0xc2, 0x03, 0x58, 0x39, 0xa3, 0xfe, 0x1c, 0x6d, 0x42, 0x7d,
	0x3c, 0x1b, 0x3a, 0xbb, 0xce, 0x7e, 0xdb, 0xab, 0x8f, 0x67, 0xf8, 0x1b, 0xe8, 0x1e, 0x45, 0x17,
	0x41, 0x78, 0xce, 0x68, 0xe2, 0xd1, 0x97, 0x0b, 0xca, 0x38, 0x72, 0x61, 0x6d, 0xc1, 0x68, 0x12,
	0xfa, 0x73, 0xaa, 0x25, 0x53, 0x5a, 0x7c, 0x8b, 0x7d, 0xc6, 0x7e, 0x8e, 0x92, 0xd9, 0xb0, 0xae,
	0xbe, 0x19, 0x1a, 0x3f, 0x85, 0x9e, 0xa5, 0x8b, 0xc5, 0x51, 0xc8, 0x28, 0xda, 0x86, 0x26, 0x8f,
	0x2e, 0x69, 0xa8, 0x35, 0x29, 0x42, 0x70, 0x69, 0x92, 0x44, 0x89, 0xd6, 0xa1, 0x08, 0xfc, 0x23,
	0xf4, 0x4e, 0x83, 0x8b, 0x70, 0x11, 0xdb, 0xde, 0x74, 0xa1, 0x71, 0x49, 0x6f, 0xf4, 0x75, 0x71,
	0xcc, 0xf9, 0x57, 0xbf, 0xc5, 0xbf, 0x46, 0xc1, 0xbf, 0x03, 0xe8, 0x8d, 0xc3, 0xeb, 0x80, 0x53,
	0x5b, 0xfd, 0x5d, 0x00, 0xb6, 0x88, 0x69, 0xf2, 0x93, 0x50, 0x21, 0xad, 0xac, 0x79, 0x6d, 0xc9,
	0x11, 0x52, 0xf8, 0x31, 0x20, 0xfb, 0x8e, 0x0e, 0xaa, 0xec, 0x53, 0x75, 0x40, 0x7f, 0x39, 0x80,
	0x46, 0x09, 0xf5, 0x39, 0x3d, 0xbc, 0xa6, 0x21, 0x37, 0x36, 0x11, 0xac, 0x58, 0xe0, 0xca, 0xb3,
	0x50, 0xc9, 0xfd, 0x0b, 0x7d, 0x5d, 0x1c, 0xd1, 0x1d, 0x68, 0xbf, 0x48, 0xa2, 0x90, 0xd3, 0x70,
	0xc6, 0x86, 0x8d, 0xdd, 0xc6, 0x7e, 0xdb, 0xcb, 0x18, 0xe2, 0x2b, 0x7d, 0x45, 0x93, 0x69, 0xc0,
	0x28, 0x1b, 0xae, 0xa8, 0xaf, 0x29, 0x43, 0x7c, 0xf5, 0xaf, 0xfd, 0xe0, 0xca, 0x9f, 0x5c, 0xd1,
	0x61, 0x73, 0xd7, 0xd9, 0x6f, 0x7a, 0x19, 0x43, 0x80, 0x34, 0xf5, 0x63, 0x7f, 0x1a, 0xf0, 0x9b,
	0x61, 0x4b, 0x7e, 0x4c, 0x69, 0xdc, 0x87, 0xde, 0x51, 0xc0, 0xb8, 0xf4, 0x97, 0x69, 0x87, 0xf1,
	0xbf, 0x0e, 0x20, 0x9b, 0xab, 0x61, 0x38, 0x80, 0x16, 0x95, 0x9c, 0xa1, 0xb3, 0xdb, 0xd8, 0xef,
	0x1c, 0xb8, 0xa4, 0x2c, 0x44, 0x34, 0xa9, 0x25, 0xdd, 0xdf, 0x1c, 0x68, 0x29, 0x96, 0x09, 0xd9,
	0xc9, 0x42, 0x36, 0xc0, 0xd4, 0x2d, 0x60, 0xee, 0x40, 0x9b, 0x53, 0x7f, 0x3e, 0x8a, 0x16, 0x21,
	0x97, 0x4f, 0xda, 0xf4, 0x32, 0x06, 0xda, 0x83, 0x0d, 0x13, 0xb5, 0x92, 0x58, 0x91, 0x12, 0x79,
	0x66, 0x2e, 0xe0, 0x66, 0x21, 0xe0, 0xf7, 0x61, 0x27, 0xf5, 0x5a, 0x94, 0x08, 0xb3, 0x12, 0x2f,
	0xef, 0x1e, 0xfe, 0xdd, 0x81, 0x41, 0x51, 0x56, 0x43, 0x71, 0x1f, 0x9a, 0xc2, 0x29, 0x83, 0xc4,
	0x5d, 0x52, 0x2d, 0x47, 0x14, 0xa5, 0x64, 0xdd, 0x67, 0xd0, 0x94, 0x74, 0xb1, 0x2a, 0x05, 0x0e,
	0xdf, 0x59, 0x38, 0x88, 0xb3, 0xc8, 0xb0, 0xc3, 0xb9, 0x1f, 0x5c, 0xe9, 0xb4, 0x56, 0x04, 0xfe,
	0x16, 0x76, 0x3c, 0xca, 0xb8, 0x9f, 0x48, 0x3b, 0x47, 0xfe, 0xc4, 0x2a, 0x62, 0x89, 0xf8, 0x59,
	0x1a, 0x42, 0x4a, 0xa3, 0x01, 0xb4, 0xae, 0xfc, 0xc9, 0x59, 0x9a, 0x6e, 0x9a, 0xc2, 0x2f, 0x61,
	0xdb, 0xa3, 0x8c, 0xf2, 0x43, 0x0d, 0x9e, 0xd1, 0xb5, 0x0b, 0x1d, 0x83, 0x67, 0xa6, 0xce, 0x66,
	0xe5, 0xac, 0xd5, 0x0b, 0xd6, 0xde, 0x32, 0xd0, 0x34, 0x24, 0x34, 0x4d, 0x89, 0x81, 0x86, 0x00,
	0xff, 0xe7, 0x68, 0xf8, 0x4d, 0xea, 0xa6, 0x88, 0x3e, 0xb1, 0x13, 0x5c, 0xa1, 0xfa, 0x0e, 0xa9,
	0x14, 0x25, 0xa9, 0xbf, 0xd9, 0x0d, 0xf7, 0x57, 0x07, 0xd6, 0x0c, 0x5f, 0xe0, 0xc9, 0xfd, 0x0b,
	0xa5, 0xa6, 0xed, 0xc9, 0x73, 0x65, 0xae, 0x7d, 0x00, 0xdd, 0x59, 0x34, 0xbd, 0xa4, 0xc9, 0x78,
	0xee, 0x5f, 0x50, 0x3b, 0xe5, 0x4a, 0x7c, 0xf4, 0x1e, 0x6c, 0x5e, 0x4f, 0xa2, 0x57, 0x96, 0xa4,
	0x4a, 0xbd, 0x02, 0x17, 0x3f, 0x83, 0x2d, 0x09, 0xaa, 0x88, 0xfa, 0x94, 0xfb, 0x7c, 0xc1, 0x04,
	0xfe, 0x22, 0xfa, 0xf4, 0xc9, 0x35, 0x25, 0xf8, 0x4c, 0x4a, 0x98, 0x77, 0x51, 0x14, 0xde, 0x83,
	0xee, 0x29, 0x8f, 0xe2, 0x5c, 0x0f, 0x29, 0x67, 0xe7, 0x13, 0xe8, 0x48, 0x89, 0xcc, 0x08, 0x0d,
	0xb9, 0xc8, 0x78, 0x6d, 0x44, 0x51, 0x4b, 0x8d, 0xfc, 0xe1, 0x40, 0xff, 0x38, 0x0a, 0x03, 0x1e,
	0x25, 0xcf, 0x23, 0xc6, 0xd3, 0x77, 0xd8, 0x83, 0x8d, 0x63, 0x3a, 0x8f, 0x92, 0x9b, 0x13, 0x9a,
	0x4c, 0x69, 0xc8, 0xa5, 0xba, 0xba, 0x97, 0x67, 0xa2, 0x7d, 0xd8, 0x52, 0x0c, 0x8f, 0xfa, 0xb3,
	0x43, 0xab, 0x13, 0x16, 0xd9, 0xe8, 0x6d, 0x80, 0xd1, 0xc9, 0xb9, 0x51, 0xd6, 0x90, 0xca, 0x2c,
	0x0e, 0xc2, 0xb0, 0x3e, 0x3a, 0x39, 0xcf, 0xd4, 0xac, 0x48, 0x35, 0x39, 0x1e, 0x5e, 0x15, 0xb5,
	0x10, 0xf3, 0x1b, 0xfc, 0x21, 0x6c, 0x7d, 0x4f, 0x13, 0x16, 0x44, 0x61, 0xea, 0xef, 0x10, 0x56,
	0xaf, 0x15, 0x4b, 0x07, 0x6e, 0x48, 0xfc, 0x8f, 0xce, 0xb5, 0xaf, 0x4c, 0x13, 0xb5, 0x73, 0x2d,
	0x6b, 0xb5, 0x76, 0xae, 0x95, 0x44, 0x89, 0xe1, 0x58, 0xbd, 0xd8, 0x9d, 0xc0, 0x9a, 0x61, 0x8b,
	0x32, 0x0d, 0xc4, 0xe3, 0x9b, 0x79, 0x27, 0x09, 0x91, 0x6c, 0x2c, 0xf8, 0x45, 0x25, 0x5b, 0xc3,
	0x93, 0x67, 0x51, 0x33, 0x73, 0x89, 0xcd, 0xf1, 0x17, 0x12, 0x86, 0x86, 0x97, 0xd2, 0xe2, 0x75,
	0xa7, 0xf1, 0x42, 0xc6, 0x5e, 0xf7, 0xc4, 0x11, 0x9f, 0xc8, 0x42, 0xa7, 0xb6, 0x47, 0xe5, 0x42,
	0x7f, 0xad, 0xd2, 0x3b, 0x82, 0xe1, 0x69, 0xa6, 0xcf, 0x3c, 0x93, 0x52, 0x5a, 0x1d, 0x85, 0xed,
	0x71, 0x3d, 0xef, 0x31, 0x7e, 0x0a, 0x3b, 0x96, 0xb6, 0x51, 0xbc, 0xb8, 0x5d, 0x95, 0x0e, 0xb0,
	0x9e, 0x05, 0xf8, 0x1c, 0xd0, 0xd7, 0xaa, 0x4a, 0xc6, 0xe1, 0x8b, 0xc8, 0xdc, 0x5e, 0x56, 0x2a,
	0xb7, 0x44, 0x8d, 0xff, 0x74, 0xa0, 0x9f, 0x53, 0xa5, 0x5f, 0xf9, 0x33, 0x68, 0x07, 0x21, 0xe3,
	0x7e, 0x38, 0xa5, 0x59, 0x9f, 0xae, 0x10, 0x24, 0x63, 0x2d, 0xe5, 0x65, 0xf2, 0xee, 0x0f, 0xb0,
	0x66, 0xd8, 0xcb, 0xdf, 0x98, 0xdf, 0xc4, 0x69, 0x43, 0x11, 0x67, 0xd1, 0xd8, 0x03, 0xb3, 0x88,
	0xd4, 0x03, 0x99, 0x1d, 0xa2, 0xdc, 0xa8, 0xee, 0x15, 0x8a, 0x38, 0xf8, 0x7b, 0x15, 0x5a, 0x5f,
	0xca, 0x6d, 0x0d, 0x3d, 0x80, 0x76, 0xba, 0x43, 0xa1, 0x1e, 0x29, 0xee, 0x66, 0x2e, 0x22, 0xa5,
	0x15, 0x0b, 0xd7, 0xd0, 0x27, 0x00, 0xd9, 0xe2, 0x84, 0x10, 0x29, 0x6d, 0x51, 0x4b, 0xee, 0x3d,
	0x04, 0xc8, 0xb6, 0x1b, 0x84, 0x48, 0x69, 0x3d, 0x72, 0xfb, 0xa4, 0xbc, 0xfe, 0xe0, 0x1a, 0x7a,
	0x00, 0x1d, 0x6b, 0xaf, 0x41, 0x7d, 0x52, 0xde, 0x72, 0xdc, 0x75, 0x62, 0xb5, 0x23, 0x5c, 0xbb,
	0xe7, 0xa0, 0x7b, 0xd0, 0x4e, 0xfb, 0x18, 0xea, 0x91, 0x62, 0x4f, 0xab, 0xb8, 0xf1, 0x10, 0x20,
	0x5b, 0x29, 0x10, 0x22, 0xa5, 0xd5, 0xc4, 0xed, 0x57, 0xec, 0x1c, 0xb8, 0x86, 0x46, 0xb0, 0x99,
	0x9f, 0xc0, 0x68, 0x40, 0x2a, 0xc7, 0xbc, 0xfb, 0xc6, 0x92, 0x51, 0x8d, 0x6b, 0xe8, 0x11, 0x6c,
	0xe6, 0x87, 0x2b, 0x1a, 0x90, 0xca, 0x69, 0x5b, 0xe1, 0xf9, 0xc7, 0xb0, 0x91, 0x1b, 0x56, 0xa8,
	0x45, 0x64, 0xcb, 0x72, 0x07, 0xd5, 0x43, 0x0c, 0xd7, 0xd0, 0x63, 0xd8, 0xc8, 0x8d, 0x5f, 0xb4,
	0x43, 0xaa, 0xc6, 0xb1, 0xdb, 0x25, 0x85, 0x81, 0x62, 0x1b, 0x4c, 0xfb, 0x43, 0xc1, 0x60, 0xa9,
	0x93, 0xe1, 0x1a, 0xfa, 0x5c, 0xc6, 0x67, 0xf5, 0x14, 0x15, 0x5f, 0xb9, 0xc9, 0x2c, 0x31, 0xf9,
	0x29, 0xf4, 0x4a, 0x1d, 0x04, 0xbd, 0x49, 0x96, 0x75, 0x15, 0x57, 0x7b, 0x84, 0x6b, 0xe8, 0x00,
	0x36, 0xf3, 0xdd, 0x02, 0x0d, 0x48, 0x65, 0xfb, 0xb0, 0xee, 0x3c, 0x82, 0x8e, 0x55, 0xac, 0xa8,
	0x4f, 0xca, 0xed, 0xc2, 0xdd, 0xae, 0xaa, 0x67, 0x5c, 0x43, 0x1f, 0x41, 0xc7, 0x9a, 0x6d, 0x29,
	0x34, 0xdb, 0xa4, 0x62, 0xe2, 0xc9, 0xd0, 0xde, 0x85, 0x55, 0x3d, 0x58, 0x52, 0xe1, 0x2e, 0x29,
	0x8c, 0x1a, 0x5c, 0x9b, 0xb4, 0xe4, 0xdf, 0xd5, 0xfd, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xf5, 0x14, 0x5e, 0x6d, 0x0d, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaemonClient interface {
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error)
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (Daemon_CreateEventClient, error)
	StopEvent(ctx context.Context, in *StopEventRequest, opts ...grpc.CallOption) (Daemon_StopEventClient, error)
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
	ListEventTeams(ctx context.Context, in *ListEventTeamsRequest, opts ...grpc.CallOption) (*ListEventTeamsResponse, error)
	RestartTeamLab(ctx context.Context, in *RestartTeamLabRequest, opts ...grpc.CallOption) (Daemon_RestartTeamLabClient, error)
	ListExercises(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListExercisesResponse, error)
	ResetExercise(ctx context.Context, in *ResetExerciseRequest, opts ...grpc.CallOption) (Daemon_ResetExerciseClient, error)
	ListFrontends(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFrontendsResponse, error)
	ResetFrontends(ctx context.Context, in *ResetFrontendsRequest, opts ...grpc.CallOption) (Daemon_ResetFrontendsClient, error)
	SetFrontendMemory(ctx context.Context, in *SetFrontendMemoryRequest, opts ...grpc.CallOption) (*Empty, error)
	SetFrontendCpu(ctx context.Context, in *SetFrontendCpuRequest, opts ...grpc.CallOption) (*Empty, error)
	GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error)
	MonitorHost(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Daemon_MonitorHostClient, error)
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/Daemon/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SignupUser(ctx context.Context, in *SignupUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/Daemon/SignupUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error) {
	out := new(InviteUserResponse)
	err := c.cc.Invoke(ctx, "/Daemon/InviteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (Daemon_CreateEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[0], "/Daemon/CreateEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonCreateEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_CreateEventClient interface {
	Recv() (*EventStatus, error)
	grpc.ClientStream
}

type daemonCreateEventClient struct {
	grpc.ClientStream
}

func (x *daemonCreateEventClient) Recv() (*EventStatus, error) {
	m := new(EventStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) StopEvent(ctx context.Context, in *StopEventRequest, opts ...grpc.CallOption) (Daemon_StopEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[1], "/Daemon/StopEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonStopEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_StopEventClient interface {
	Recv() (*EventStatus, error)
	grpc.ClientStream
}

type daemonStopEventClient struct {
	grpc.ClientStream
}

func (x *daemonStopEventClient) Recv() (*EventStatus, error) {
	m := new(EventStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/Daemon/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ListEventTeams(ctx context.Context, in *ListEventTeamsRequest, opts ...grpc.CallOption) (*ListEventTeamsResponse, error) {
	out := new(ListEventTeamsResponse)
	err := c.cc.Invoke(ctx, "/Daemon/ListEventTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) RestartTeamLab(ctx context.Context, in *RestartTeamLabRequest, opts ...grpc.CallOption) (Daemon_RestartTeamLabClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[2], "/Daemon/RestartTeamLab", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonRestartTeamLabClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_RestartTeamLabClient interface {
	Recv() (*EventStatus, error)
	grpc.ClientStream
}

type daemonRestartTeamLabClient struct {
	grpc.ClientStream
}

func (x *daemonRestartTeamLabClient) Recv() (*EventStatus, error) {
	m := new(EventStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) ListExercises(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListExercisesResponse, error) {
	out := new(ListExercisesResponse)
	err := c.cc.Invoke(ctx, "/Daemon/ListExercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ResetExercise(ctx context.Context, in *ResetExerciseRequest, opts ...grpc.CallOption) (Daemon_ResetExerciseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[3], "/Daemon/ResetExercise", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonResetExerciseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_ResetExerciseClient interface {
	Recv() (*ResetTeamStatus, error)
	grpc.ClientStream
}

type daemonResetExerciseClient struct {
	grpc.ClientStream
}

func (x *daemonResetExerciseClient) Recv() (*ResetTeamStatus, error) {
	m := new(ResetTeamStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) ListFrontends(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListFrontendsResponse, error) {
	out := new(ListFrontendsResponse)
	err := c.cc.Invoke(ctx, "/Daemon/ListFrontends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ResetFrontends(ctx context.Context, in *ResetFrontendsRequest, opts ...grpc.CallOption) (Daemon_ResetFrontendsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[4], "/Daemon/ResetFrontends", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonResetFrontendsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_ResetFrontendsClient interface {
	Recv() (*ResetTeamStatus, error)
	grpc.ClientStream
}

type daemonResetFrontendsClient struct {
	grpc.ClientStream
}

func (x *daemonResetFrontendsClient) Recv() (*ResetTeamStatus, error) {
	m := new(ResetTeamStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) SetFrontendMemory(ctx context.Context, in *SetFrontendMemoryRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Daemon/SetFrontendMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SetFrontendCpu(ctx context.Context, in *SetFrontendCpuRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Daemon/SetFrontendCpu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error) {
	out := new(GetTeamInfoResponse)
	err := c.cc.Invoke(ctx, "/Daemon/GetTeamInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) MonitorHost(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Daemon_MonitorHostClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[5], "/Daemon/MonitorHost", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonMonitorHostClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_MonitorHostClient interface {
	Recv() (*MonitorHostResponse, error)
	grpc.ClientStream
}

type daemonMonitorHostClient struct {
	grpc.ClientStream
}

func (x *daemonMonitorHostClient) Recv() (*MonitorHostResponse, error) {
	m := new(MonitorHostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/Daemon/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
type DaemonServer interface {
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	SignupUser(context.Context, *SignupUserRequest) (*LoginUserResponse, error)
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error)
	CreateEvent(*CreateEventRequest, Daemon_CreateEventServer) error
	StopEvent(*StopEventRequest, Daemon_StopEventServer) error
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	ListEventTeams(context.Context, *ListEventTeamsRequest) (*ListEventTeamsResponse, error)
	RestartTeamLab(*RestartTeamLabRequest, Daemon_RestartTeamLabServer) error
	ListExercises(context.Context, *Empty) (*ListExercisesResponse, error)
	ResetExercise(*ResetExerciseRequest, Daemon_ResetExerciseServer) error
	ListFrontends(context.Context, *Empty) (*ListFrontendsResponse, error)
	ResetFrontends(*ResetFrontendsRequest, Daemon_ResetFrontendsServer) error
	SetFrontendMemory(context.Context, *SetFrontendMemoryRequest) (*Empty, error)
	SetFrontendCpu(context.Context, *SetFrontendCpuRequest) (*Empty, error)
	GetTeamInfo(context.Context, *GetTeamInfoRequest) (*GetTeamInfoResponse, error)
	MonitorHost(*Empty, Daemon_MonitorHostServer) error
	Version(context.Context, *Empty) (*VersionResponse, error)
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SignupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SignupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/SignupUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SignupUser(ctx, req.(*SignupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/InviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_CreateEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).CreateEvent(m, &daemonCreateEventServer{stream})
}

type Daemon_CreateEventServer interface {
	Send(*EventStatus) error
	grpc.ServerStream
}

type daemonCreateEventServer struct {
	grpc.ServerStream
}

func (x *daemonCreateEventServer) Send(m *EventStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_StopEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StopEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).StopEvent(m, &daemonStopEventServer{stream})
}

type Daemon_StopEventServer interface {
	Send(*EventStatus) error
	grpc.ServerStream
}

type daemonStopEventServer struct {
	grpc.ServerStream
}

func (x *daemonStopEventServer) Send(m *EventStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ListEventTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ListEventTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/ListEventTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ListEventTeams(ctx, req.(*ListEventTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_RestartTeamLab_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RestartTeamLabRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).RestartTeamLab(m, &daemonRestartTeamLabServer{stream})
}

type Daemon_RestartTeamLabServer interface {
	Send(*EventStatus) error
	grpc.ServerStream
}

type daemonRestartTeamLabServer struct {
	grpc.ServerStream
}

func (x *daemonRestartTeamLabServer) Send(m *EventStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_ListExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ListExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/ListExercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ListExercises(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ResetExercise_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResetExerciseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).ResetExercise(m, &daemonResetExerciseServer{stream})
}

type Daemon_ResetExerciseServer interface {
	Send(*ResetTeamStatus) error
	grpc.ServerStream
}

type daemonResetExerciseServer struct {
	grpc.ServerStream
}

func (x *daemonResetExerciseServer) Send(m *ResetTeamStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_ListFrontends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ListFrontends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/ListFrontends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ListFrontends(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ResetFrontends_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResetFrontendsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).ResetFrontends(m, &daemonResetFrontendsServer{stream})
}

type Daemon_ResetFrontendsServer interface {
	Send(*ResetTeamStatus) error
	grpc.ServerStream
}

type daemonResetFrontendsServer struct {
	grpc.ServerStream
}

func (x *daemonResetFrontendsServer) Send(m *ResetTeamStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_SetFrontendMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFrontendMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SetFrontendMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/SetFrontendMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SetFrontendMemory(ctx, req.(*SetFrontendMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SetFrontendCpu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFrontendCpuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SetFrontendCpu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/SetFrontendCpu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SetFrontendCpu(ctx, req.(*SetFrontendCpuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_GetTeamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetTeamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/GetTeamInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetTeamInfo(ctx, req.(*GetTeamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_MonitorHost_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).MonitorHost(m, &daemonMonitorHostServer{stream})
}

type Daemon_MonitorHostServer interface {
	Send(*MonitorHostResponse) error
	grpc.ServerStream
}

type daemonMonitorHostServer struct {
	grpc.ServerStream
}

func (x *daemonMonitorHostServer) Send(m *MonitorHostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _Daemon_LoginUser_Handler,
		},
		{
			MethodName: "SignupUser",
			Handler:    _Daemon_SignupUser_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _Daemon_InviteUser_Handler,
		},
		{
			MethodName: "ListEvents",
			Handler:    _Daemon_ListEvents_Handler,
		},
		{
			MethodName: "ListEventTeams",
			Handler:    _Daemon_ListEventTeams_Handler,
		},
		{
			MethodName: "ListExercises",
			Handler:    _Daemon_ListExercises_Handler,
		},
		{
			MethodName: "ListFrontends",
			Handler:    _Daemon_ListFrontends_Handler,
		},
		{
			MethodName: "SetFrontendMemory",
			Handler:    _Daemon_SetFrontendMemory_Handler,
		},
		{
			MethodName: "SetFrontendCpu",
			Handler:    _Daemon_SetFrontendCpu_Handler,
		},
		{
			MethodName: "GetTeamInfo",
			Handler:    _Daemon_GetTeamInfo_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Daemon_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateEvent",
			Handler:       _Daemon_CreateEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StopEvent",
			Handler:       _Daemon_StopEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RestartTeamLab",
			Handler:       _Daemon_RestartTeamLab_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ResetExercise",
			Handler:       _Daemon_ResetExercise_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ResetFrontends",
			Handler:       _Daemon_ResetFrontends_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MonitorHost",
			Handler:       _Daemon_MonitorHost_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "daemon.proto",
}
