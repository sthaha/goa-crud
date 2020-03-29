// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create.proto

package createpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateRequest struct {
	// ID of a person
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of person
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Comments
	Comments             []string `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

type CreateResponse struct {
	// ID of a person
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Name of person
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Comments
	Comments             []string `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateResponse) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Field                []*Storedblog `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetField() []*Storedblog {
	if m != nil {
		return m.Field
	}
	return nil
}

// A Storedblog describes a blog retrieved by the storage service.
type Storedblog struct {
	// ID is the unique id of the blog.
	Id uint32 `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name of person
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Comments
	Comments             []string `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Storedblog) Reset()         { *m = Storedblog{} }
func (m *Storedblog) String() string { return proto.CompactTextString(m) }
func (*Storedblog) ProtoMessage()    {}
func (*Storedblog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{4}
}

func (m *Storedblog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storedblog.Unmarshal(m, b)
}
func (m *Storedblog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storedblog.Marshal(b, m, deterministic)
}
func (m *Storedblog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storedblog.Merge(m, src)
}
func (m *Storedblog) XXX_Size() int {
	return xxx_messageInfo_Storedblog.Size(m)
}
func (m *Storedblog) XXX_DiscardUnknown() {
	xxx_messageInfo_Storedblog.DiscardUnknown(m)
}

var xxx_messageInfo_Storedblog proto.InternalMessageInfo

func (m *Storedblog) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Storedblog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Storedblog) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

type RemoveRequest struct {
	// ID of blog to remove
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{5}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{6}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

type UpdateRequest struct {
	// ID of blog to be updated
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Details of blog to be updated
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Comments to be updated
	Comments             []string `protobuf:"bytes,3,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d26d5dcda09a78, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "create.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "create.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "create.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "create.ListResponse")
	proto.RegisterType((*Storedblog)(nil), "create.Storedblog")
	proto.RegisterType((*RemoveRequest)(nil), "create.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "create.RemoveResponse")
	proto.RegisterType((*UpdateRequest)(nil), "create.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "create.UpdateResponse")
}

func init() {
	proto.RegisterFile("create.proto", fileDescriptor_a4d26d5dcda09a78)
}

var fileDescriptor_a4d26d5dcda09a78 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4f, 0xb3, 0x40,
	0x10, 0x86, 0xb3, 0xc0, 0x47, 0xe8, 0xb4, 0x90, 0x66, 0x3f, 0x35, 0x84, 0x8b, 0x84, 0x13, 0xa7,
	0x26, 0xd6, 0x43, 0x3d, 0xeb, 0xb5, 0x89, 0x66, 0x8d, 0x17, 0x6f, 0xa5, 0x8c, 0x86, 0xa4, 0xb0,
	0x08, 0xab, 0x3f, 0xd9, 0xdf, 0x61, 0x60, 0x77, 0xbb, 0x2e, 0x8d, 0x17, 0xbd, 0xcd, 0xcc, 0xce,
	0xbc, 0xf3, 0xce, 0x03, 0xb0, 0xd8, 0x77, 0xb8, 0x13, 0xb8, 0x6a, 0x3b, 0x2e, 0x38, 0xf5, 0x65,
	0x96, 0xdd, 0x43, 0x78, 0x37, 0x46, 0x0c, 0xdf, 0xde, 0xb1, 0x17, 0x34, 0x02, 0xa7, 0x2a, 0x63,
	0x27, 0x25, 0x79, 0xc8, 0x9c, 0xaa, 0xa4, 0x14, 0xbc, 0x66, 0x57, 0x63, 0x4c, 0x52, 0x92, 0xcf,
	0xd8, 0x18, 0xd3, 0x04, 0x82, 0x3d, 0xaf, 0x6b, 0x6c, 0x44, 0x1f, 0x7b, 0xa9, 0x9b, 0xcf, 0xd8,
	0x31, 0xcf, 0x1e, 0x20, 0xd2, 0x82, 0x7d, 0xcb, 0x9b, 0x1e, 0xff, 0xac, 0x18, 0xc2, 0x7c, 0x5b,
	0xf5, 0x42, 0x19, 0xcc, 0x6e, 0x60, 0x21, 0x53, 0x25, 0x9f, 0xc3, 0xbf, 0x97, 0x0a, 0x0f, 0x65,
	0x4c, 0x52, 0x37, 0x9f, 0xaf, 0xe9, 0x4a, 0xdd, 0xf9, 0x28, 0x78, 0x87, 0x65, 0x71, 0xe0, 0xaf,
	0x4c, 0x36, 0x64, 0x5b, 0x00, 0x53, 0x54, 0xb6, 0x82, 0x5f, 0xdb, 0xba, 0x84, 0x90, 0x61, 0xcd,
	0x3f, 0x26, 0xe4, 0x88, 0x16, 0xcc, 0x96, 0x10, 0xe9, 0x06, 0x69, 0x75, 0x80, 0xfd, 0xd4, 0x96,
	0x27, 0xb0, 0xc9, 0x89, 0x07, 0xe7, 0x07, 0x0f, 0xee, 0xc4, 0xc3, 0x12, 0x22, 0x2d, 0x28, 0x57,
	0xac, 0x3f, 0x09, 0xf8, 0x92, 0x3f, 0xdd, 0x1c, 0xa3, 0x73, 0xcd, 0xc4, 0xfa, 0xd4, 0xc9, 0xc5,
	0xb4, 0xac, 0x88, 0x5e, 0x81, 0x37, 0x10, 0xa6, 0xff, 0xf5, 0xfb, 0x37, 0xfc, 0xc9, 0x99, 0x5d,
	0x54, 0x23, 0x1b, 0xf0, 0xe5, 0xad, 0x66, 0x97, 0x05, 0xc7, 0xec, 0xb2, 0x91, 0x0c, 0x83, 0xf2,
	0x02, 0x33, 0x68, 0x21, 0x32, 0x83, 0xf6, 0xa1, 0xb7, 0xf0, 0x1c, 0xc8, 0x87, 0xb6, 0x28, 0xfc,
	0xf1, 0x9f, 0xbe, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x81, 0x1e, 0x78, 0x50, 0xe3, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreateClient is the client API for Create service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateClient interface {
	// Add new blog and return its ID.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// List all entries
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Remove blog from storage
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	// Updating the existing blog
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type createClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateClient(cc grpc.ClientConnInterface) CreateClient {
	return &createClient{cc}
}

func (c *createClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/create.Create/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/create.Create/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/create.Create/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/create.Create/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateServer is the server API for Create service.
type CreateServer interface {
	// Add new blog and return its ID.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// List all entries
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Remove blog from storage
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// Updating the existing blog
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

// UnimplementedCreateServer can be embedded to have forward compatible implementations.
type UnimplementedCreateServer struct {
}

func (*UnimplementedCreateServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCreateServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCreateServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedCreateServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterCreateServer(s *grpc.Server, srv CreateServer) {
	s.RegisterService(&_Create_serviceDesc, srv)
}

func _Create_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/create.Create/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Create_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/create.Create/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Create_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/create.Create/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Create_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/create.Create/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Create_serviceDesc = grpc.ServiceDesc{
	ServiceName: "create.Create",
	HandlerType: (*CreateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Create_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Create_List_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Create_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Create_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "create.proto",
}