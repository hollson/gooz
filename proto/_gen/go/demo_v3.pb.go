// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo_v3.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type PersonReply_Sex int32

const (
	PersonReply_FEMALE PersonReply_Sex = 0
	PersonReply_MALE   PersonReply_Sex = 1
)

var PersonReply_Sex_name = map[int32]string{
	0: "FEMALE",
	1: "MALE",
}

var PersonReply_Sex_value = map[string]int32{
	"FEMALE": 0,
	"MALE":   1,
}

func (x PersonReply_Sex) String() string {
	return proto.EnumName(PersonReply_Sex_name, int32(x))
}

func (PersonReply_Sex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ac5a0fb00e43f6a, []int{1, 0}
}

// 请求参数
type PersonRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonRequest) Reset()         { *m = PersonRequest{} }
func (m *PersonRequest) String() string { return proto.CompactTextString(m) }
func (*PersonRequest) ProtoMessage()    {}
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ac5a0fb00e43f6a, []int{0}
}

func (m *PersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonRequest.Unmarshal(m, b)
}
func (m *PersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonRequest.Marshal(b, m, deterministic)
}
func (m *PersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonRequest.Merge(m, src)
}
func (m *PersonRequest) XXX_Size() int {
	return xxx_messageInfo_PersonRequest.Size(m)
}
func (m *PersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonRequest proto.InternalMessageInfo

func (m *PersonRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 相应结果
type PersonReply struct {
	//姓名
	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IdCard string          `protobuf:"bytes,2,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	Age    int32           `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sex    PersonReply_Sex `protobuf:"varint,4,opt,name=sex,proto3,enum=PersonReply_Sex" json:"sex,omitempty"`
	//是否已婚
	Married              bool                `protobuf:"varint,5,opt,name=married,proto3" json:"married,omitempty"`
	Amount               float64             `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Address              []string            `protobuf:"bytes,7,rep,name=address,proto3" json:"address,omitempty"`
	Educational          *Append_Educational `protobuf:"bytes,8,opt,name=educational,proto3" json:"educational,omitempty"`
	Others               []*any.Any          `protobuf:"bytes,9,rep,name=Others,proto3" json:"Others,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PersonReply) Reset()         { *m = PersonReply{} }
func (m *PersonReply) String() string { return proto.CompactTextString(m) }
func (*PersonReply) ProtoMessage()    {}
func (*PersonReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ac5a0fb00e43f6a, []int{1}
}

func (m *PersonReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonReply.Unmarshal(m, b)
}
func (m *PersonReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonReply.Marshal(b, m, deterministic)
}
func (m *PersonReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonReply.Merge(m, src)
}
func (m *PersonReply) XXX_Size() int {
	return xxx_messageInfo_PersonReply.Size(m)
}
func (m *PersonReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonReply.DiscardUnknown(m)
}

var xxx_messageInfo_PersonReply proto.InternalMessageInfo

func (m *PersonReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonReply) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *PersonReply) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonReply) GetSex() PersonReply_Sex {
	if m != nil {
		return m.Sex
	}
	return PersonReply_FEMALE
}

func (m *PersonReply) GetMarried() bool {
	if m != nil {
		return m.Married
	}
	return false
}

func (m *PersonReply) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PersonReply) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PersonReply) GetEducational() *Append_Educational {
	if m != nil {
		return m.Educational
	}
	return nil
}

func (m *PersonReply) GetOthers() []*any.Any {
	if m != nil {
		return m.Others
	}
	return nil
}

// 附加信息
type Append struct {
	Educational          []*Append_Educational `protobuf:"bytes,1,rep,name=educational,proto3" json:"educational,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Append) Reset()         { *m = Append{} }
func (m *Append) String() string { return proto.CompactTextString(m) }
func (*Append) ProtoMessage()    {}
func (*Append) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ac5a0fb00e43f6a, []int{2}
}

func (m *Append) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Append.Unmarshal(m, b)
}
func (m *Append) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Append.Marshal(b, m, deterministic)
}
func (m *Append) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Append.Merge(m, src)
}
func (m *Append) XXX_Size() int {
	return xxx_messageInfo_Append.Size(m)
}
func (m *Append) XXX_DiscardUnknown() {
	xxx_messageInfo_Append.DiscardUnknown(m)
}

var xxx_messageInfo_Append proto.InternalMessageInfo

func (m *Append) GetEducational() []*Append_Educational {
	if m != nil {
		return m.Educational
	}
	return nil
}

type Append_Educational struct {
	University           string   `protobuf:"bytes,1,opt,name=university,proto3" json:"university,omitempty"`
	Specialty            []string `protobuf:"bytes,2,rep,name=specialty,proto3" json:"specialty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Append_Educational) Reset()         { *m = Append_Educational{} }
func (m *Append_Educational) String() string { return proto.CompactTextString(m) }
func (*Append_Educational) ProtoMessage()    {}
func (*Append_Educational) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ac5a0fb00e43f6a, []int{2, 0}
}

func (m *Append_Educational) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Append_Educational.Unmarshal(m, b)
}
func (m *Append_Educational) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Append_Educational.Marshal(b, m, deterministic)
}
func (m *Append_Educational) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Append_Educational.Merge(m, src)
}
func (m *Append_Educational) XXX_Size() int {
	return xxx_messageInfo_Append_Educational.Size(m)
}
func (m *Append_Educational) XXX_DiscardUnknown() {
	xxx_messageInfo_Append_Educational.DiscardUnknown(m)
}

var xxx_messageInfo_Append_Educational proto.InternalMessageInfo

func (m *Append_Educational) GetUniversity() string {
	if m != nil {
		return m.University
	}
	return ""
}

func (m *Append_Educational) GetSpecialty() []string {
	if m != nil {
		return m.Specialty
	}
	return nil
}

func init() {
	proto.RegisterEnum("PersonReply_Sex", PersonReply_Sex_name, PersonReply_Sex_value)
	proto.RegisterType((*PersonRequest)(nil), "PersonRequest")
	proto.RegisterType((*PersonReply)(nil), "PersonReply")
	proto.RegisterType((*Append)(nil), "Append")
	proto.RegisterType((*Append_Educational)(nil), "Append.Educational")
}

func init() { proto.RegisterFile("demo_v3.proto", fileDescriptor_5ac5a0fb00e43f6a) }

var fileDescriptor_5ac5a0fb00e43f6a = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x71, 0xea, 0x24, 0x63, 0x5a, 0x45, 0x4b, 0x04, 0xab, 0x80, 0x90, 0x31, 0x17,
	0x1f, 0xd0, 0x46, 0x72, 0x05, 0x17, 0x2e, 0xa4, 0x34, 0xfc, 0x11, 0x20, 0x2a, 0xf7, 0xc6, 0xa5,
	0xda, 0x78, 0xa7, 0x61, 0x25, 0x7b, 0xd7, 0xac, 0xd7, 0x55, 0xfc, 0x0e, 0x3c, 0x09, 0x0f, 0xc1,
	0xb3, 0xa1, 0xd8, 0x89, 0x6a, 0x10, 0x12, 0xb7, 0xf9, 0xbe, 0xf9, 0x79, 0x3c, 0xfa, 0x66, 0xe1,
	0x44, 0x62, 0x61, 0xae, 0x6f, 0xcf, 0x78, 0x69, 0x8d, 0x33, 0xf3, 0xa7, 0x4a, 0x67, 0x79, 0x2d,
	0x71, 0xb1, 0x31, 0x66, 0x93, 0xe3, 0xa2, 0x75, 0xd7, 0xf5, 0xcd, 0x42, 0xe8, 0xa6, 0x43, 0xa2,
	0x67, 0x70, 0x72, 0x89, 0xb6, 0x32, 0x3a, 0xc5, 0xef, 0x35, 0x56, 0x8e, 0x52, 0x18, 0x6a, 0x51,
	0x20, 0x23, 0x21, 0x89, 0x27, 0x69, 0x5b, 0x47, 0xbf, 0x06, 0x10, 0x1c, 0xa8, 0x32, 0x6f, 0xfe,
	0xc5, 0xd0, 0x87, 0x30, 0x52, 0xf2, 0x3a, 0x13, 0x56, 0xb2, 0x41, 0x6b, 0xfb, 0x4a, 0xbe, 0x11,
	0x56, 0xd2, 0x29, 0x78, 0x62, 0x83, 0xcc, 0x0b, 0x49, 0x7c, 0x9c, 0xee, 0x4a, 0x1a, 0x81, 0x57,
	0xe1, 0x96, 0x0d, 0x43, 0x12, 0x9f, 0x26, 0x53, 0xde, 0x9b, 0xcc, 0xaf, 0x70, 0x9b, 0xee, 0x9a,
	0x94, 0xc1, 0xa8, 0x10, 0xd6, 0x2a, 0x94, 0xec, 0x38, 0x24, 0xf1, 0x38, 0x3d, 0x48, 0xfa, 0x00,
	0x7c, 0x51, 0x98, 0x5a, 0x3b, 0xe6, 0x87, 0x24, 0x26, 0xe9, 0x5e, 0xed, 0xbe, 0x10, 0x52, 0x5a,
	0xac, 0x2a, 0x36, 0x0a, 0xbd, 0x78, 0x92, 0x1e, 0x24, 0x7d, 0x01, 0x01, 0xca, 0x3a, 0x13, 0x4e,
	0x19, 0x2d, 0x72, 0x36, 0x0e, 0x49, 0x1c, 0x24, 0xf7, 0xf9, 0xb2, 0x2c, 0x51, 0x4b, 0xbe, 0xba,
	0x6b, 0xa5, 0x7d, 0x8e, 0x3e, 0x07, 0xff, 0x8b, 0xfb, 0x86, 0xb6, 0x62, 0x93, 0xd0, 0x8b, 0x83,
	0x64, 0xc6, 0xbb, 0x18, 0xf9, 0x21, 0x46, 0xbe, 0xd4, 0x4d, 0xba, 0x67, 0xa2, 0x47, 0xe0, 0x5d,
	0xe1, 0x96, 0x02, 0xf8, 0x6f, 0x57, 0x9f, 0x97, 0x9f, 0x56, 0xd3, 0x23, 0x3a, 0x86, 0x61, 0x5b,
	0x91, 0xe8, 0x07, 0x01, 0xbf, 0xfb, 0xdd, 0xdf, 0xcb, 0x90, 0x76, 0xf4, 0x7f, 0x97, 0x99, 0x7f,
	0x84, 0xa0, 0xd7, 0xa3, 0x4f, 0x00, 0x6a, 0xad, 0x6e, 0xd1, 0x56, 0xca, 0x35, 0xfb, 0x3b, 0xf4,
	0x1c, 0xfa, 0x18, 0x26, 0x55, 0x89, 0x99, 0x12, 0xb9, 0x6b, 0xd8, 0xa0, 0x8d, 0xe3, 0xce, 0x48,
	0x5e, 0x83, 0xdf, 0x85, 0x4e, 0x5f, 0xc2, 0xec, 0x1d, 0xba, 0x4e, 0x7c, 0xd0, 0x37, 0xe6, 0xbd,
	0xd0, 0x32, 0x47, 0x4b, 0x4f, 0xf9, 0x1f, 0xaf, 0x62, 0x7e, 0xaf, 0x7f, 0xa5, 0xe8, 0xe8, 0x3c,
	0x81, 0x59, 0x66, 0x0a, 0x8e, 0x5b, 0x51, 0x94, 0x39, 0x72, 0x57, 0x3b, 0x63, 0x95, 0xc8, 0xcf,
	0x27, 0x1d, 0x76, 0xe1, 0xcc, 0xd7, 0x21, 0x7f, 0x55, 0xae, 0x7f, 0x0e, 0xc6, 0x97, 0xd6, 0xf0,
	0x0b, 0x2c, 0xcc, 0xda, 0x6f, 0x73, 0x3b, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x84, 0x1d, 0x07,
	0xe2, 0xa5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonClient is the client API for Person service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonClient interface {
	// 传入一个姓名，返回相应的个人信息
	GetPersonInfoHandler(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error)
}

type personClient struct {
	cc *grpc.ClientConn
}

func NewPersonClient(cc *grpc.ClientConn) PersonClient {
	return &personClient{cc}
}

func (c *personClient) GetPersonInfoHandler(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonReply, error) {
	out := new(PersonReply)
	err := c.cc.Invoke(ctx, "/Person/GetPersonInfoHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonServer is the server API for Person service.
type PersonServer interface {
	// 传入一个姓名，返回相应的个人信息
	GetPersonInfoHandler(context.Context, *PersonRequest) (*PersonReply, error)
}

// UnimplementedPersonServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServer struct {
}

func (*UnimplementedPersonServer) GetPersonInfoHandler(ctx context.Context, req *PersonRequest) (*PersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonInfoHandler not implemented")
}

func RegisterPersonServer(s *grpc.Server, srv PersonServer) {
	s.RegisterService(&_Person_serviceDesc, srv)
}

func _Person_GetPersonInfoHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServer).GetPersonInfoHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Person/GetPersonInfoHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServer).GetPersonInfoHandler(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Person_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Person",
	HandlerType: (*PersonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonInfoHandler",
			Handler:    _Person_GetPersonInfoHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo_v3.proto",
}
