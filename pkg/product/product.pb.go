// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package product

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

type ProductDetail struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductDetail) Reset()         { *m = ProductDetail{} }
func (m *ProductDetail) String() string { return proto.CompactTextString(m) }
func (*ProductDetail) ProtoMessage()    {}
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *ProductDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductDetail.Unmarshal(m, b)
}
func (m *ProductDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductDetail.Marshal(b, m, deterministic)
}
func (m *ProductDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductDetail.Merge(m, src)
}
func (m *ProductDetail) XXX_Size() int {
	return xxx_messageInfo_ProductDetail.Size(m)
}
func (m *ProductDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ProductDetail proto.InternalMessageInfo

func (m *ProductDetail) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductDetail) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type PurchaseRequest struct {
	CustomerId           int32            `protobuf:"varint,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	Products             []*ProductDetail `protobuf:"bytes,2,rep,name=Products,proto3" json:"Products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PurchaseRequest) Reset()         { *m = PurchaseRequest{} }
func (m *PurchaseRequest) String() string { return proto.CompactTextString(m) }
func (*PurchaseRequest) ProtoMessage()    {}
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *PurchaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseRequest.Unmarshal(m, b)
}
func (m *PurchaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseRequest.Marshal(b, m, deterministic)
}
func (m *PurchaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseRequest.Merge(m, src)
}
func (m *PurchaseRequest) XXX_Size() int {
	return xxx_messageInfo_PurchaseRequest.Size(m)
}
func (m *PurchaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseRequest proto.InternalMessageInfo

func (m *PurchaseRequest) GetCustomerId() int32 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *PurchaseRequest) GetProducts() []*ProductDetail {
	if m != nil {
		return m.Products
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PurchaseResult struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseResult) Reset()         { *m = PurchaseResult{} }
func (m *PurchaseResult) String() string { return proto.CompactTextString(m) }
func (*PurchaseResult) ProtoMessage()    {}
func (*PurchaseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{3}
}

func (m *PurchaseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseResult.Unmarshal(m, b)
}
func (m *PurchaseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseResult.Marshal(b, m, deterministic)
}
func (m *PurchaseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseResult.Merge(m, src)
}
func (m *PurchaseResult) XXX_Size() int {
	return xxx_messageInfo_PurchaseResult.Size(m)
}
func (m *PurchaseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseResult.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseResult proto.InternalMessageInfo

func (m *PurchaseResult) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PurchaseResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PurchaseResponse struct {
	Response             *Response         `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Result               []*PurchaseResult `protobuf:"bytes,2,rep,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PurchaseResponse) Reset()         { *m = PurchaseResponse{} }
func (m *PurchaseResponse) String() string { return proto.CompactTextString(m) }
func (*PurchaseResponse) ProtoMessage()    {}
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{4}
}

func (m *PurchaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseResponse.Unmarshal(m, b)
}
func (m *PurchaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseResponse.Marshal(b, m, deterministic)
}
func (m *PurchaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseResponse.Merge(m, src)
}
func (m *PurchaseResponse) XXX_Size() int {
	return xxx_messageInfo_PurchaseResponse.Size(m)
}
func (m *PurchaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseResponse proto.InternalMessageInfo

func (m *PurchaseResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PurchaseResponse) GetResult() []*PurchaseResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductDetail)(nil), "ProductDetail")
	proto.RegisterType((*PurchaseRequest)(nil), "PurchaseRequest")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*PurchaseResult)(nil), "PurchaseResult")
	proto.RegisterType((*PurchaseResponse)(nil), "PurchaseResponse")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0xa5, 0x68, 0x2d, 0x8f, 0x50, 0xea, 0x1c, 0x4c, 0xc3, 0xc1, 0x90, 0x4d, 0x8c, 0xc4,
	0xc3, 0xc6, 0x94, 0x8b, 0x89, 0x89, 0x17, 0xb8, 0x70, 0x30, 0xc1, 0xbd, 0x7b, 0x28, 0x74, 0xa3,
	0x24, 0xc8, 0xd6, 0xee, 0xee, 0xc1, 0x7f, 0x6f, 0x58, 0xb6, 0x2d, 0xe0, 0x6d, 0xe6, 0x4d, 0xe7,
	0xcd, 0xf7, 0xba, 0x18, 0x94, 0x95, 0x2a, 0xec, 0xda, 0xf0, 0xb2, 0x52, 0x46, 0xb1, 0x17, 0x0c,
	0x96, 0x07, 0x61, 0x2e, 0x4d, 0xbe, 0xd9, 0x52, 0x8c, 0x60, 0x51, 0xa4, 0x9d, 0x71, 0x67, 0x72,
	0x25, 0x82, 0x45, 0x41, 0x23, 0x44, 0xef, 0x36, 0xdf, 0x99, 0x8d, 0xf9, 0x4d, 0x03, 0xa7, 0x36,
	0x3d, 0xfb, 0xc0, 0x70, 0x69, 0xab, 0xf5, 0x57, 0xae, 0xa5, 0x90, 0x3f, 0x56, 0x6a, 0x43, 0x77,
	0xc0, 0xcc, 0x6a, 0xa3, 0xbe, 0x65, 0xd5, 0xd8, 0x1c, 0x29, 0xf4, 0x88, 0xc8, 0xdf, 0xd3, 0x69,
	0x30, 0xee, 0x4e, 0xfa, 0x59, 0xcc, 0x4f, 0x00, 0x44, 0x33, 0x67, 0x4f, 0x88, 0x84, 0xd4, 0xa5,
	0xda, 0x69, 0x49, 0x84, 0xcb, 0x99, 0x2a, 0xa4, 0x77, 0x74, 0x35, 0x25, 0xe8, 0xbe, 0xe9, 0x4f,
	0x47, 0xd5, 0x13, 0xfb, 0x92, 0x3d, 0x23, 0x6e, 0x81, 0xb4, 0xdd, 0x9a, 0x7f, 0x71, 0x6e, 0x11,
	0x56, 0x6e, 0xe2, 0xd7, 0x7c, 0xc7, 0x56, 0x48, 0x8e, 0x36, 0x0f, 0x37, 0xef, 0xdb, 0xfb, 0xce,
	0xa1, 0x9f, 0xf5, 0x78, 0x2d, 0x88, 0x16, 0xed, 0x01, 0xa1, 0xa8, 0x2d, 0xf7, 0x81, 0x86, 0xfc,
	0x94, 0x41, 0xf8, 0x71, 0xf6, 0x8a, 0x6b, 0x9f, 0x8d, 0xa6, 0xc0, 0x5c, 0xd5, 0x9f, 0x51, 0xc2,
	0xcf, 0x7e, 0xe3, 0xe8, 0x86, 0x9f, 0xd3, 0xb0, 0x8b, 0x55, 0xe8, 0x9e, 0x6c, 0xfa, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x40, 0xe0, 0xd0, 0xab, 0xc3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	DoPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) DoPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseResponse, error) {
	out := new(PurchaseResponse)
	err := c.cc.Invoke(ctx, "/Product/DoPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	DoPurchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error)
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) DoPurchase(ctx context.Context, req *PurchaseRequest) (*PurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoPurchase not implemented")
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_DoPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DoPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Product/DoPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DoPurchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoPurchase",
			Handler:    _Product_DoPurchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
