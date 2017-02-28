// Code generated by protoc-gen-go.
// source: pb/profile/profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	pb/profile/profile.proto

It has these top-level messages:
	Request
	Result
	Hotel
	Address
	Image
*/
package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	Locale   string   `protobuf:"bytes,2,opt,name=locale" json:"locale,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber string  `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string  `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string  `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string  `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
	Lat          float32 `protobuf:"fixed32,7,opt,name=lat" json:"lat,omitempty"`
	Lon          float32 `protobuf:"fixed32,8,opt,name=lon" json:"lon,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Profile service

type ProfileClient interface {
	GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/profile.Profile/GetProfiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/profile/profile.proto",
}

func init() { proto.RegisterFile("pb/profile/profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x74, 0x93, 0xb4, 0x13, 0xb4, 0xac, 0x46, 0x08, 0x59, 0x7b, 0x40, 0x51, 0x0e,
	0x28, 0xe2, 0xb0, 0xac, 0xda, 0x23, 0xe2, 0x80, 0x38, 0x40, 0x2f, 0x08, 0xf9, 0x0d, 0xdc, 0x78,
	0x4a, 0x23, 0xb9, 0x71, 0xb0, 0x9d, 0x43, 0x9f, 0x8f, 0x67, 0xe0, 0x7d, 0x90, 0x1d, 0xa7, 0x0d,
	0x7b, 0xca, 0xfc, 0xdf, 0x8c, 0xf3, 0xfb, 0xb7, 0x06, 0xd8, 0x70, 0xf8, 0x38, 0x18, 0x7d, 0xec,
	0x14, 0xcd, 0xdf, 0xa7, 0xc1, 0x68, 0xa7, 0xb1, 0x88, 0xb2, 0xfe, 0x0c, 0x05, 0xa7, 0xdf, 0x23,
	0x59, 0x87, 0x8f, 0xb0, 0x3e, 0x69, 0x47, 0x6a, 0x2f, 0x2d, 0x4b, 0xaa, 0x55, 0xb3, 0xe1, 0x57,
	0x8d, 0x6f, 0x21, 0x57, 0xba, 0x15, 0x8a, 0x58, 0x5a, 0x25, 0xcd, 0x86, 0x47, 0x55, 0x3f, 0x43,
	0xce, 0xc9, 0x8e, 0xca, 0xe1, 0x7b, 0xc8, 0xc3, 0xf4, 0x74, 0xb6, 0xdc, 0xde, 0x3f, 0xcd, 0x8e,
	0xdf, 0x3d, 0xe6, 0xb1, 0x5b, 0xff, 0x49, 0x20, 0x0b, 0x04, 0xef, 0x21, 0xed, 0x24, 0x4b, 0xc2,
	0xff, 0xd2, 0x4e, 0x22, 0xc2, 0x5d, 0x2f, 0xce, 0xb3, 0x43, 0xa8, 0xb1, 0x82, 0x72, 0x38, 0xe9,
	0x9e, 0x7e, 0x8c, 0xe7, 0x03, 0x19, 0xb6, 0x0a, 0xad, 0x25, 0xf2, 0x13, 0x92, 0x6c, 0x6b, 0xba,
	0xc1, 0x75, 0xba, 0x67, 0x77, 0xd3, 0xc4, 0x02, 0xe1, 0x07, 0x28, 0x84, 0x94, 0x86, 0xac, 0x65,
	0x59, 0x95, 0x34, 0xe5, 0xf6, 0xe1, 0x7a, 0xb5, 0x2f, 0x13, 0xe7, 0xf3, 0x80, 0x4f, 0xd1, 0x9d,
	0xc5, 0x2f, 0xb2, 0x2c, 0x7f, 0x91, 0x62, 0xef, 0x31, 0x8f, 0xdd, 0xfa, 0x6f, 0x02, 0x45, 0x3c,
	0x8c, 0x35, 0xbc, 0xb2, 0xce, 0x10, 0xb9, 0x78, 0xc9, 0x29, 0xd1, 0x7f, 0x0c, 0xdf, 0x01, 0x44,
	0x7d, 0x4b, 0xb8, 0x20, 0x3e, 0x7b, 0xdb, 0xb9, 0x4b, 0x0c, 0x18, 0x6a, 0x7c, 0x03, 0x99, 0x75,
	0xc2, 0x51, 0xcc, 0x34, 0x09, 0x64, 0x50, 0xb4, 0x7a, 0xec, 0x9d, 0xb9, 0x84, 0x34, 0x1b, 0x3e,
	0x4b, 0xef, 0x31, 0x68, 0xeb, 0x84, 0xfa, 0xaa, 0x25, 0xb1, 0x7c, 0xf2, 0xb8, 0x11, 0x7c, 0x80,
	0x95, 0x12, 0x8e, 0x15, 0x55, 0xd2, 0xa4, 0xdc, 0x97, 0x81, 0xe8, 0x9e, 0xad, 0x23, 0xd1, 0x7d,
	0xbd, 0x83, 0x2c, 0x04, 0xf5, 0xad, 0xd1, 0xa8, 0x98, 0xc5, 0x97, 0xde, 0x58, 0xd2, 0x51, 0x8c,
	0xca, 0x85, 0xfb, 0xaf, 0xf9, 0x2c, 0xb7, 0x9f, 0xa0, 0xf8, 0x39, 0xbd, 0x12, 0x3e, 0x43, 0xf9,
	0x8d, 0x5c, 0x54, 0x16, 0x6f, 0x2f, 0x1d, 0x97, 0xec, 0xf1, 0xf5, 0x82, 0xf8, 0xbd, 0x39, 0xe4,
	0x61, 0x21, 0x77, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xca, 0xda, 0x00, 0x3f, 0xac, 0x02, 0x00,
	0x00,
}