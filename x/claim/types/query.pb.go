// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/claim/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClaimableRequest struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *ClaimableRequest) Reset()         { *m = ClaimableRequest{} }
func (m *ClaimableRequest) String() string { return proto.CompactTextString(m) }
func (*ClaimableRequest) ProtoMessage()    {}
func (*ClaimableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bba73508cdd8c1d, []int{0}
}
func (m *ClaimableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimableRequest.Merge(m, src)
}
func (m *ClaimableRequest) XXX_Size() int {
	return m.Size()
}
func (m *ClaimableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimableRequest proto.InternalMessageInfo

func (m *ClaimableRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type ClaimableResponse struct {
	Coins []types.Coin `protobuf:"bytes,1,rep,name=coins,proto3" json:"coins" yaml:"coins"`
}

func (m *ClaimableResponse) Reset()         { *m = ClaimableResponse{} }
func (m *ClaimableResponse) String() string { return proto.CompactTextString(m) }
func (*ClaimableResponse) ProtoMessage()    {}
func (*ClaimableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bba73508cdd8c1d, []int{1}
}
func (m *ClaimableResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimableResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimableResponse.Merge(m, src)
}
func (m *ClaimableResponse) XXX_Size() int {
	return m.Size()
}
func (m *ClaimableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimableResponse proto.InternalMessageInfo

func (m *ClaimableResponse) GetCoins() []types.Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

type ActivitiesRequest struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *ActivitiesRequest) Reset()         { *m = ActivitiesRequest{} }
func (m *ActivitiesRequest) String() string { return proto.CompactTextString(m) }
func (*ActivitiesRequest) ProtoMessage()    {}
func (*ActivitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bba73508cdd8c1d, []int{2}
}
func (m *ActivitiesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivitiesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivitiesRequest.Merge(m, src)
}
func (m *ActivitiesRequest) XXX_Size() int {
	return m.Size()
}
func (m *ActivitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivitiesRequest proto.InternalMessageInfo

func (m *ActivitiesRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type ActivitiesResponse struct {
	Sender    string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	All       []string `protobuf:"bytes,2,rep,name=all,proto3" json:"all,omitempty"`
	Completed []string `protobuf:"bytes,3,rep,name=completed,proto3" json:"completed,omitempty"`
}

func (m *ActivitiesResponse) Reset()         { *m = ActivitiesResponse{} }
func (m *ActivitiesResponse) String() string { return proto.CompactTextString(m) }
func (*ActivitiesResponse) ProtoMessage()    {}
func (*ActivitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1bba73508cdd8c1d, []int{3}
}
func (m *ActivitiesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivitiesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivitiesResponse.Merge(m, src)
}
func (m *ActivitiesResponse) XXX_Size() int {
	return m.Size()
}
func (m *ActivitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivitiesResponse proto.InternalMessageInfo

func (m *ActivitiesResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ActivitiesResponse) GetAll() []string {
	if m != nil {
		return m.All
	}
	return nil
}

func (m *ActivitiesResponse) GetCompleted() []string {
	if m != nil {
		return m.Completed
	}
	return nil
}

func init() {
	proto.RegisterType((*ClaimableRequest)(nil), "osmosis.claim.v1beta1.ClaimableRequest")
	proto.RegisterType((*ClaimableResponse)(nil), "osmosis.claim.v1beta1.ClaimableResponse")
	proto.RegisterType((*ActivitiesRequest)(nil), "osmosis.claim.v1beta1.ActivitiesRequest")
	proto.RegisterType((*ActivitiesResponse)(nil), "osmosis.claim.v1beta1.ActivitiesResponse")
}

func init() { proto.RegisterFile("osmosis/claim/v1beta1/query.proto", fileDescriptor_1bba73508cdd8c1d) }

var fileDescriptor_1bba73508cdd8c1d = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x1b, 0x5a, 0xc8, 0xa8, 0xd0, 0x1d, 0x2a, 0xc4, 0x50, 0xb2, 0x6b, 0x44, 0x4c,
	0x91, 0xce, 0xd0, 0x7a, 0x13, 0x14, 0xdc, 0x22, 0x9e, 0xcd, 0xb1, 0xb7, 0x49, 0x3a, 0xc4, 0x81,
	0x49, 0x5e, 0xba, 0x33, 0x5b, 0x5c, 0xc4, 0x8b, 0x57, 0x41, 0x04, 0xbf, 0x83, 0x9f, 0xa5, 0xc7,
	0x82, 0x17, 0x4f, 0x8b, 0xec, 0xfa, 0x09, 0xfa, 0x09, 0x24, 0x99, 0x34, 0x5d, 0x5d, 0x17, 0x17,
	0x6f, 0xc3, 0x9b, 0xdf, 0xfb, 0xbf, 0xff, 0x7b, 0x6f, 0x06, 0xdf, 0x07, 0x5d, 0x80, 0x96, 0x9a,
	0x65, 0x8a, 0xcb, 0x82, 0x9d, 0x1f, 0xa6, 0xc2, 0xf0, 0x43, 0x76, 0x36, 0x11, 0xe3, 0x29, 0xad,
	0xc6, 0x60, 0x80, 0xdc, 0x6d, 0x11, 0xda, 0x20, 0xb4, 0x45, 0x82, 0xdd, 0x1c, 0x72, 0x68, 0x08,
	0x56, 0x9f, 0x2c, 0x1c, 0xec, 0xe5, 0x00, 0xb9, 0x12, 0x8c, 0x57, 0x92, 0xf1, 0xb2, 0x04, 0xc3,
	0x8d, 0x84, 0x52, 0xb7, 0xb7, 0x61, 0xd6, 0x68, 0xb1, 0x94, 0x6b, 0xd1, 0xd5, 0xca, 0x40, 0x96,
	0xf6, 0x3e, 0x7a, 0x86, 0x77, 0x8e, 0xeb, 0x22, 0x3c, 0x55, 0x22, 0x11, 0x67, 0x13, 0xa1, 0x0d,
	0xd9, 0xc7, 0xdb, 0x5a, 0x94, 0xa7, 0x62, 0xec, 0xa3, 0x21, 0x8a, 0xbd, 0x51, 0xff, 0x6a, 0x36,
	0xb8, 0x33, 0xe5, 0x85, 0x7a, 0x1a, 0xd9, 0x78, 0x94, 0xb4, 0x40, 0x74, 0x82, 0xfb, 0x4b, 0xe9,
	0xba, 0x82, 0x52, 0x0b, 0xf2, 0x12, 0x6f, 0xd5, 0x15, 0xb4, 0x8f, 0x86, 0x6e, 0x7c, 0xeb, 0xe8,
	0x1e, 0xb5, 0x1e, 0x68, 0xed, 0xe1, 0xba, 0x19, 0x7a, 0x0c, 0xb2, 0x1c, 0xed, 0x5e, 0xcc, 0x06,
	0xce, 0xd5, 0x6c, 0x70, 0xdb, 0xaa, 0x37, 0x59, 0x51, 0x62, 0xb3, 0xa3, 0xe7, 0xb8, 0xff, 0x22,
	0x33, 0xf2, 0x5c, 0x1a, 0x29, 0xf4, 0x7f, 0x78, 0x03, 0x4c, 0x96, 0xf3, 0x5b, 0x73, 0x9b, 0x0b,
	0x90, 0x1d, 0xec, 0x72, 0xa5, 0xfc, 0xde, 0xd0, 0x8d, 0xbd, 0xa4, 0x3e, 0x92, 0x3d, 0xec, 0x65,
	0x50, 0x54, 0x4a, 0x18, 0x71, 0xea, 0xbb, 0x4d, 0xfc, 0x26, 0x70, 0xf4, 0xb5, 0x87, 0xb7, 0x5e,
	0xd7, 0x6b, 0x24, 0x1f, 0x11, 0xf6, 0xba, 0xb9, 0x90, 0x47, 0xf4, 0xaf, 0xfb, 0xa4, 0x7f, 0x0e,
	0x3e, 0x88, 0xff, 0x0d, 0xda, 0x2e, 0xa2, 0xc7, 0x1f, 0xbe, 0xfd, 0xfc, 0xd2, 0x7b, 0x48, 0x1e,
	0xb0, 0xdf, 0x5e, 0x53, 0xcd, 0x74, 0x5b, 0x7e, 0x67, 0xdb, 0x78, 0x4f, 0x3e, 0x21, 0x8c, 0x6f,
	0x26, 0x41, 0xd6, 0x55, 0x59, 0x19, 0x76, 0xb0, 0xbf, 0x01, 0xb9, 0xd6, 0x10, 0xb7, 0x90, 0xd0,
	0x2b, 0x86, 0x46, 0xaf, 0x2e, 0xe6, 0x21, 0xba, 0x9c, 0x87, 0xe8, 0xc7, 0x3c, 0x44, 0x9f, 0x17,
	0xa1, 0x73, 0xb9, 0x08, 0x9d, 0xef, 0x8b, 0xd0, 0x39, 0x39, 0xc8, 0xa5, 0x79, 0x33, 0x49, 0x69,
	0x06, 0xc5, 0xb5, 0xd0, 0x81, 0xe2, 0xa9, 0xee, 0x54, 0xdf, 0xb6, 0xdf, 0xc6, 0x4c, 0x2b, 0xa1,
	0xd3, 0xed, 0xe6, 0x11, 0x3f, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x89, 0xab, 0x6e, 0x86, 0x54,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Claimable(ctx context.Context, in *ClaimableRequest, opts ...grpc.CallOption) (*ClaimableResponse, error)
	Activities(ctx context.Context, in *ActivitiesRequest, opts ...grpc.CallOption) (*ActivitiesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Claimable(ctx context.Context, in *ClaimableRequest, opts ...grpc.CallOption) (*ClaimableResponse, error) {
	out := new(ClaimableResponse)
	err := c.cc.Invoke(ctx, "/osmosis.claim.v1beta1.Query/Claimable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Activities(ctx context.Context, in *ActivitiesRequest, opts ...grpc.CallOption) (*ActivitiesResponse, error) {
	out := new(ActivitiesResponse)
	err := c.cc.Invoke(ctx, "/osmosis.claim.v1beta1.Query/Activities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Claimable(context.Context, *ClaimableRequest) (*ClaimableResponse, error)
	Activities(context.Context, *ActivitiesRequest) (*ActivitiesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Claimable(ctx context.Context, req *ClaimableRequest) (*ClaimableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claimable not implemented")
}
func (*UnimplementedQueryServer) Activities(ctx context.Context, req *ActivitiesRequest) (*ActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activities not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Claimable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Claimable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.claim.v1beta1.Query/Claimable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Claimable(ctx, req.(*ClaimableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Activities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Activities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.claim.v1beta1.Query/Activities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Activities(ctx, req.(*ActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.claim.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claimable",
			Handler:    _Query_Claimable_Handler,
		},
		{
			MethodName: "Activities",
			Handler:    _Query_Activities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/claim/v1beta1/query.proto",
}

func (m *ClaimableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimableResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimableResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimableResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ActivitiesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivitiesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivitiesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActivitiesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivitiesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivitiesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Completed) > 0 {
		for iNdEx := len(m.Completed) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Completed[iNdEx])
			copy(dAtA[i:], m.Completed[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Completed[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.All) > 0 {
		for iNdEx := len(m.All) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.All[iNdEx])
			copy(dAtA[i:], m.All[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.All[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimableRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ClaimableResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *ActivitiesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ActivitiesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.All) > 0 {
		for _, s := range m.All {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Completed) > 0 {
		for _, s := range m.Completed {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimableRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimableResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ActivitiesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivitiesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivitiesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ActivitiesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivitiesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivitiesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field All", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.All = append(m.All, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Completed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Completed = append(m.Completed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)