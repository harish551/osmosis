// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/poolmodels/balancer/v1beta1/tx.proto

package balancer

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// ===================== MsgCreatePool
type MsgCreateBalancerPool struct {
	Sender             string      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolParams         *PoolParams `protobuf:"bytes,2,opt,name=pool_params,json=poolParams,proto3" json:"pool_params,omitempty" yaml:"pool_params"`
	PoolAssets         []PoolAsset `protobuf:"bytes,3,rep,name=pool_assets,json=poolAssets,proto3" json:"pool_assets"`
	FuturePoolGovernor string      `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
}

func (m *MsgCreateBalancerPool) Reset()         { *m = MsgCreateBalancerPool{} }
func (m *MsgCreateBalancerPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBalancerPool) ProtoMessage()    {}
func (*MsgCreateBalancerPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d22c5192b37962a, []int{0}
}
func (m *MsgCreateBalancerPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBalancerPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBalancerPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBalancerPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBalancerPool.Merge(m, src)
}
func (m *MsgCreateBalancerPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBalancerPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBalancerPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBalancerPool proto.InternalMessageInfo

func (m *MsgCreateBalancerPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateBalancerPool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreateBalancerPool) GetPoolAssets() []PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func (m *MsgCreateBalancerPool) GetFuturePoolGovernor() string {
	if m != nil {
		return m.FuturePoolGovernor
	}
	return ""
}

// Returns the poolID
type MsgCreateBalancerPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateBalancerPoolResponse) Reset()         { *m = MsgCreateBalancerPoolResponse{} }
func (m *MsgCreateBalancerPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBalancerPoolResponse) ProtoMessage()    {}
func (*MsgCreateBalancerPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d22c5192b37962a, []int{1}
}
func (m *MsgCreateBalancerPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBalancerPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBalancerPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBalancerPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBalancerPoolResponse.Merge(m, src)
}
func (m *MsgCreateBalancerPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBalancerPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBalancerPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBalancerPoolResponse proto.InternalMessageInfo

func (m *MsgCreateBalancerPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateBalancerPool)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPool")
	proto.RegisterType((*MsgCreateBalancerPoolResponse)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPoolResponse")
}

func init() {
	proto.RegisterFile("osmosis/gamm/poolmodels/balancer/v1beta1/tx.proto", fileDescriptor_4d22c5192b37962a)
}

var fileDescriptor_4d22c5192b37962a = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0xdb, 0x28, 0x88, 0x8b, 0x18, 0x7a, 0x2a, 0x10, 0x05, 0x61, 0x07, 0x33, 0x10, 0x2a,
	0xe5, 0x4e, 0x09, 0x12, 0x43, 0x97, 0x0a, 0x53, 0x51, 0x75, 0xa8, 0x54, 0xb2, 0x95, 0xa5, 0x3a,
	0x3b, 0xd7, 0xc3, 0x92, 0xcf, 0xcf, 0xf2, 0x39, 0x51, 0x59, 0x19, 0x99, 0xf8, 0x05, 0x8c, 0xcc,
	0xfd, 0x19, 0x1d, 0x3b, 0x32, 0x59, 0x28, 0x19, 0x3a, 0xb0, 0xe5, 0x17, 0xa0, 0x3b, 0x9f, 0xdb,
	0x06, 0x05, 0x09, 0x89, 0xc5, 0xbe, 0x7b, 0xef, 0xfb, 0xbe, 0xf7, 0xee, 0x7b, 0x77, 0x68, 0x08,
	0x4a, 0x82, 0x8a, 0x15, 0x15, 0x4c, 0x4a, 0x9a, 0x01, 0x24, 0x12, 0x26, 0x3c, 0x51, 0x34, 0x64,
	0x09, 0x4b, 0x23, 0x9e, 0xd3, 0xd9, 0x30, 0xe4, 0x05, 0x1b, 0xd2, 0xe2, 0x9c, 0x64, 0x39, 0x14,
	0x80, 0xfb, 0x96, 0x42, 0x34, 0x85, 0xdc, 0x52, 0x48, 0x4d, 0x21, 0x96, 0xd2, 0xdd, 0x16, 0x20,
	0xc0, 0x90, 0xa8, 0x5e, 0x55, 0xfc, 0xee, 0x16, 0x93, 0x71, 0x0a, 0xd4, 0x7c, 0x6d, 0xe8, 0xc5,
	0x4a, 0x17, 0x75, 0xc5, 0x5a, 0xef, 0x18, 0x20, 0xb1, 0x40, 0x37, 0x32, 0x48, 0x1a, 0x32, 0xc5,
	0x6f, 0x70, 0x11, 0xc4, 0xa9, 0xcd, 0x7b, 0x02, 0x40, 0x24, 0x9c, 0x9a, 0x5d, 0x38, 0x3d, 0xa3,
	0x45, 0x2c, 0xb9, 0x2a, 0x98, 0xcc, 0x2c, 0xe0, 0xb1, 0x15, 0x90, 0x4a, 0xd0, 0xd9, 0x50, 0xff,
	0xaa, 0x84, 0xff, 0x6b, 0x03, 0x3d, 0x3c, 0x52, 0xe2, 0x6d, 0xce, 0x59, 0xc1, 0x83, 0x3b, 0x95,
	0xf1, 0x4b, 0xd4, 0x52, 0x3c, 0x9d, 0xf0, 0xbc, 0xe3, 0xf4, 0x9c, 0xfe, 0xfd, 0x60, 0x6b, 0x59,
	0x7a, 0x0f, 0x3e, 0x31, 0x99, 0xec, 0xfa, 0x55, 0xdc, 0x1f, 0x5b, 0x00, 0x3e, 0x41, 0x6d, 0xed,
	0xc7, 0x69, 0xc6, 0x72, 0x26, 0x55, 0x67, 0xa3, 0xe7, 0xf4, 0xdb, 0xa3, 0x1e, 0x59, 0x31, 0xcc,
	0x76, 0x4d, 0xb4, 0xf6, 0xb1, 0xc1, 0x05, 0x8f, 0x96, 0xa5, 0x87, 0x2b, 0xc5, 0x3b, 0x74, 0x7f,
	0x8c, 0xb2, 0x1b, 0x0c, 0x7e, 0x67, 0xa5, 0x99, 0x52, 0xbc, 0x50, 0x9d, 0xcd, 0xde, 0x66, 0xbf,
	0x3d, 0xf2, 0xfe, 0x2e, 0xfd, 0x46, 0xe3, 0x82, 0xe6, 0x65, 0xe9, 0x35, 0x2a, 0x1d, 0x13, 0x50,
	0xf8, 0x3d, 0xda, 0x3e, 0x9b, 0x16, 0xd3, 0x9c, 0x9f, 0x1a, 0x39, 0x01, 0x33, 0x9e, 0xa7, 0x90,
	0x77, 0x9a, 0xe6, 0x6c, 0xde, 0xb2, 0xf4, 0x9e, 0x54, 0x9d, 0xac, 0x43, 0xf9, 0x63, 0x5c, 0x85,
	0x75, 0x85, 0x03, 0x1b, 0xdc, 0x1d, 0x7e, 0xbe, 0xbe, 0xd8, 0xb1, 0x16, 0x7c, 0xb9, 0xbe, 0xd8,
	0x79, 0xb6, 0x32, 0xcd, 0xc8, 0x38, 0x3a, 0xa8, 0x87, 0x39, 0xd0, 0x82, 0xfe, 0x3e, 0x7a, 0xba,
	0xd6, 0xec, 0x31, 0x57, 0x19, 0xa4, 0x8a, 0xe3, 0xe7, 0xe8, 0x9e, 0xa9, 0x1c, 0x4f, 0x8c, 0xeb,
	0xcd, 0x00, 0xcd, 0x4b, 0xaf, 0xa5, 0x21, 0x87, 0xfb, 0xe3, 0x96, 0x4e, 0x1d, 0x4e, 0x46, 0xdf,
	0x1d, 0xb4, 0x79, 0xa4, 0x04, 0xfe, 0xe6, 0x20, 0xbc, 0x66, 0x70, 0x7b, 0xe4, 0x5f, 0x6f, 0x2a,
	0x59, 0xdb, 0x4c, 0xf7, 0xe0, 0x3f, 0x05, 0xea, 0xd3, 0x04, 0x27, 0x97, 0x73, 0xd7, 0xb9, 0x9a,
	0xbb, 0xce, 0xcf, 0xb9, 0xeb, 0x7c, 0x5d, 0xb8, 0x8d, 0xab, 0x85, 0xdb, 0xf8, 0xb1, 0x70, 0x1b,
	0x1f, 0xf6, 0x44, 0x5c, 0x7c, 0x9c, 0x86, 0x24, 0x02, 0x49, 0x6d, 0xb1, 0x41, 0xc2, 0x42, 0x55,
	0x6f, 0xe8, 0x6c, 0xf4, 0x9a, 0x9e, 0xdf, 0xbe, 0xce, 0xc1, 0x1f, 0xcf, 0x33, 0x6c, 0x99, 0xeb,
	0xfb, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x7b, 0x27, 0x3b, 0xc9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateBalancerPool(ctx context.Context, in *MsgCreateBalancerPool, opts ...grpc.CallOption) (*MsgCreateBalancerPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBalancerPool(ctx context.Context, in *MsgCreateBalancerPool, opts ...grpc.CallOption) (*MsgCreateBalancerPoolResponse, error) {
	out := new(MsgCreateBalancerPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/CreateBalancerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateBalancerPool(context.Context, *MsgCreateBalancerPool) (*MsgCreateBalancerPoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBalancerPool(ctx context.Context, req *MsgCreateBalancerPool) (*MsgCreateBalancerPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBalancerPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBalancerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBalancerPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBalancerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/CreateBalancerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBalancerPool(ctx, req.(*MsgCreateBalancerPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.gamm.poolmodels.balancer.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBalancerPool",
			Handler:    _Msg_CreateBalancerPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/poolmodels/balancer/v1beta1/tx.proto",
}

func (m *MsgCreateBalancerPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBalancerPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBalancerPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBalancerPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBalancerPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBalancerPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateBalancerPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBalancerPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateBalancerPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateBalancerPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBalancerPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateBalancerPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateBalancerPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBalancerPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
