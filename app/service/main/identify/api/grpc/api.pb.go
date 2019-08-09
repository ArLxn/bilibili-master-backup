// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/main/identify/api/grpc/api.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CookieReq request param for rpc CookieInfo
type GetCookieInfoReq struct {
	// all user cookie of HTTP request
	// example: 'SESSDATA=DEE4597D836A5A9DA29DFC1AB1EFFDEB;sid=exampleSID'
	Cookie               string   `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCookieInfoReq) Reset()      { *m = GetCookieInfoReq{} }
func (*GetCookieInfoReq) ProtoMessage() {}
func (*GetCookieInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_6c6ed1fc039864fa, []int{0}
}
func (m *GetCookieInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCookieInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCookieInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetCookieInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCookieInfoReq.Merge(dst, src)
}
func (m *GetCookieInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetCookieInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCookieInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetCookieInfoReq proto.InternalMessageInfo

func (m *GetCookieInfoReq) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

// CookieReply reply val for rpc CookieInfo
type GetCookieInfoReply struct {
	// 用户是否登录
	IsLogin bool `protobuf:"varint,1,opt,name=is_login,json=isLogin,proto3" json:"is_login"`
	// user mid
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid"`
	// cookie csrf
	// when token reqest this field is empty
	Csrf string `protobuf:"bytes,3,opt,name=csrf,proto3" json:"csrfToken"`
	// expire time(unix timestamp)
	Expires              int32    `protobuf:"varint,4,opt,name=expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCookieInfoReply) Reset()      { *m = GetCookieInfoReply{} }
func (*GetCookieInfoReply) ProtoMessage() {}
func (*GetCookieInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_6c6ed1fc039864fa, []int{1}
}
func (m *GetCookieInfoReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCookieInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCookieInfoReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetCookieInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCookieInfoReply.Merge(dst, src)
}
func (m *GetCookieInfoReply) XXX_Size() int {
	return m.Size()
}
func (m *GetCookieInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCookieInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetCookieInfoReply proto.InternalMessageInfo

func (m *GetCookieInfoReply) GetIsLogin() bool {
	if m != nil {
		return m.IsLogin
	}
	return false
}

func (m *GetCookieInfoReply) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *GetCookieInfoReply) GetCsrf() string {
	if m != nil {
		return m.Csrf
	}
	return ""
}

func (m *GetCookieInfoReply) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

// TokenReq request param for rpc TokenInfo
type GetTokenInfoReq struct {
	// user access token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" form:"access_key" validate:"required"`
	// buvid
	Buvid                string   `protobuf:"bytes,2,opt,name=buvid,proto3" json:"buvid,omitempty" form:"buvid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenInfoReq) Reset()      { *m = GetTokenInfoReq{} }
func (*GetTokenInfoReq) ProtoMessage() {}
func (*GetTokenInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_6c6ed1fc039864fa, []int{2}
}
func (m *GetTokenInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTokenInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTokenInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetTokenInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenInfoReq.Merge(dst, src)
}
func (m *GetTokenInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetTokenInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenInfoReq proto.InternalMessageInfo

func (m *GetTokenInfoReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetTokenInfoReq) GetBuvid() string {
	if m != nil {
		return m.Buvid
	}
	return ""
}

// TokenReply reply val for rpc TokenInfo
type GetTokenInfoReply struct {
	// 用户是否登录
	IsLogin bool `protobuf:"varint,1,opt,name=is_login,json=isLogin,proto3" json:"is_login"`
	// user mid
	Mid int64 `protobuf:"varint,2,opt,name=mid,proto3" json:"mid"`
	// cookie csrf
	// when token reqest this field is empty
	Csrf string `protobuf:"bytes,3,opt,name=csrf,proto3" json:"csrfToken"`
	// expire time(unix timestamp)
	Expires              int32    `protobuf:"varint,4,opt,name=expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenInfoReply) Reset()      { *m = GetTokenInfoReply{} }
func (*GetTokenInfoReply) ProtoMessage() {}
func (*GetTokenInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_6c6ed1fc039864fa, []int{3}
}
func (m *GetTokenInfoReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTokenInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTokenInfoReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetTokenInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenInfoReply.Merge(dst, src)
}
func (m *GetTokenInfoReply) XXX_Size() int {
	return m.Size()
}
func (m *GetTokenInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenInfoReply proto.InternalMessageInfo

func (m *GetTokenInfoReply) GetIsLogin() bool {
	if m != nil {
		return m.IsLogin
	}
	return false
}

func (m *GetTokenInfoReply) GetMid() int64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *GetTokenInfoReply) GetCsrf() string {
	if m != nil {
		return m.Csrf
	}
	return ""
}

func (m *GetTokenInfoReply) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*GetCookieInfoReq)(nil), "passport.service.identify.v1.GetCookieInfoReq")
	proto.RegisterType((*GetCookieInfoReply)(nil), "passport.service.identify.v1.GetCookieInfoReply")
	proto.RegisterType((*GetTokenInfoReq)(nil), "passport.service.identify.v1.GetTokenInfoReq")
	proto.RegisterType((*GetTokenInfoReply)(nil), "passport.service.identify.v1.GetTokenInfoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentifyClient is the client API for Identify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentifyClient interface {
	// CookieInfo identify info by cookie.
	GetCookieInfo(ctx context.Context, in *GetCookieInfoReq, opts ...grpc.CallOption) (*GetCookieInfoReply, error)
	// TokenInfo identify info by token.
	GetTokenInfo(ctx context.Context, in *GetTokenInfoReq, opts ...grpc.CallOption) (*GetTokenInfoReply, error)
}

type identifyClient struct {
	cc *grpc.ClientConn
}

func NewIdentifyClient(cc *grpc.ClientConn) IdentifyClient {
	return &identifyClient{cc}
}

func (c *identifyClient) GetCookieInfo(ctx context.Context, in *GetCookieInfoReq, opts ...grpc.CallOption) (*GetCookieInfoReply, error) {
	out := new(GetCookieInfoReply)
	err := c.cc.Invoke(ctx, "/passport.service.identify.v1.Identify/GetCookieInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identifyClient) GetTokenInfo(ctx context.Context, in *GetTokenInfoReq, opts ...grpc.CallOption) (*GetTokenInfoReply, error) {
	out := new(GetTokenInfoReply)
	err := c.cc.Invoke(ctx, "/passport.service.identify.v1.Identify/GetTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentifyServer is the server API for Identify service.
type IdentifyServer interface {
	// CookieInfo identify info by cookie.
	GetCookieInfo(context.Context, *GetCookieInfoReq) (*GetCookieInfoReply, error)
	// TokenInfo identify info by token.
	GetTokenInfo(context.Context, *GetTokenInfoReq) (*GetTokenInfoReply, error)
}

func RegisterIdentifyServer(s *grpc.Server, srv IdentifyServer) {
	s.RegisterService(&_Identify_serviceDesc, srv)
}

func _Identify_GetCookieInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCookieInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentifyServer).GetCookieInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passport.service.identify.v1.Identify/GetCookieInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentifyServer).GetCookieInfo(ctx, req.(*GetCookieInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identify_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentifyServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passport.service.identify.v1.Identify/GetTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentifyServer).GetTokenInfo(ctx, req.(*GetTokenInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passport.service.identify.v1.Identify",
	HandlerType: (*IdentifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCookieInfo",
			Handler:    _Identify_GetCookieInfo_Handler,
		},
		{
			MethodName: "GetTokenInfo",
			Handler:    _Identify_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/main/identify/api/grpc/api.proto",
}

func (m *GetCookieInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCookieInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cookie) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Cookie)))
		i += copy(dAtA[i:], m.Cookie)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetCookieInfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCookieInfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsLogin {
		dAtA[i] = 0x8
		i++
		if m.IsLogin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Mid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
	}
	if len(m.Csrf) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Csrf)))
		i += copy(dAtA[i:], m.Csrf)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetTokenInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTokenInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Buvid) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Buvid)))
		i += copy(dAtA[i:], m.Buvid)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetTokenInfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTokenInfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IsLogin {
		dAtA[i] = 0x8
		i++
		if m.IsLogin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Mid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Mid))
	}
	if len(m.Csrf) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Csrf)))
		i += copy(dAtA[i:], m.Csrf)
	}
	if m.Expires != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetCookieInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cookie)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetCookieInfoReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsLogin {
		n += 2
	}
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	l = len(m.Csrf)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovApi(uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetTokenInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Buvid)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetTokenInfoReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsLogin {
		n += 2
	}
	if m.Mid != 0 {
		n += 1 + sovApi(uint64(m.Mid))
	}
	l = len(m.Csrf)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Expires != 0 {
		n += 1 + sovApi(uint64(m.Expires))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetCookieInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCookieInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCookieInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cookie", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cookie = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetCookieInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetCookieInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCookieInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLogin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLogin = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTokenInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTokenInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTokenInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buvid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buvid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetTokenInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetTokenInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTokenInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLogin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLogin = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mid", wireType)
			}
			m.Mid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expires", wireType)
			}
			m.Expires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expires |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/main/identify/api/grpc/api.proto", fileDescriptor_api_6c6ed1fc039864fa)
}

var fileDescriptor_api_6c6ed1fc039864fa = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0x31, 0x6f, 0x13, 0x31,
	0x18, 0xc5, 0x4d, 0xd3, 0x24, 0x26, 0x15, 0xc5, 0x42, 0x28, 0x44, 0x28, 0x17, 0x2c, 0x15, 0x02,
	0x52, 0x6d, 0x02, 0x5b, 0x17, 0xa4, 0x63, 0xa8, 0x2a, 0x31, 0x59, 0x4c, 0x2c, 0xd5, 0xe5, 0xe2,
	0x3b, 0xac, 0x24, 0x67, 0xc7, 0x76, 0x4e, 0x84, 0x7f, 0x03, 0xbf, 0x86, 0x91, 0x91, 0xe9, 0x04,
	0x19, 0x23, 0xb1, 0xf4, 0x17, 0x20, 0xfb, 0xee, 0x50, 0xdb, 0x01, 0x94, 0xad, 0xcb, 0xe9, 0xbd,
	0xe7, 0xef, 0x7d, 0x3e, 0x7f, 0xdf, 0x83, 0x34, 0x52, 0x8a, 0x1a, 0xae, 0x73, 0x11, 0x73, 0xba,
	0x88, 0x44, 0x46, 0xc5, 0x94, 0x67, 0x56, 0x24, 0x6b, 0x1a, 0x29, 0x41, 0x53, 0xad, 0x62, 0x9a,
	0x8f, 0x1d, 0x26, 0x4a, 0x4b, 0x2b, 0xd1, 0x63, 0x15, 0x19, 0xa3, 0xa4, 0xb6, 0xa4, 0x72, 0x91,
	0xda, 0x40, 0xf2, 0x71, 0xff, 0x24, 0x15, 0xf6, 0xe3, 0x6a, 0x42, 0x62, 0xb9, 0xa0, 0xa9, 0x4c,
	0x25, 0xf5, 0xa6, 0xc9, 0x2a, 0xf1, 0xcc, 0x13, 0x8f, 0xca, 0x66, 0xf8, 0x05, 0x3c, 0x3a, 0xe3,
	0xf6, 0xad, 0x94, 0x33, 0xc1, 0xcf, 0xb3, 0x44, 0x32, 0xbe, 0x44, 0x0f, 0xe1, 0x41, 0xec, 0x85,
	0x1e, 0x18, 0x82, 0x51, 0x87, 0x55, 0x0c, 0x7f, 0x05, 0x10, 0xdd, 0x28, 0x56, 0xf3, 0x35, 0x7a,
	0x06, 0xdb, 0xc2, 0x5c, 0xcc, 0x65, 0x2a, 0x32, 0x6f, 0x68, 0x87, 0xdd, 0x6d, 0x11, 0xfc, 0xd5,
	0x58, 0x4b, 0x98, 0x77, 0x0e, 0xa0, 0x47, 0xb0, 0xb1, 0x10, 0xd3, 0xde, 0xde, 0x10, 0x8c, 0x1a,
	0x61, 0x6b, 0x5b, 0x04, 0x8e, 0x32, 0xf7, 0x41, 0x4f, 0xe0, 0x7e, 0x6c, 0x74, 0xd2, 0x6b, 0xb8,
	0x0b, 0xc3, 0xc3, 0x6d, 0x11, 0x74, 0x1c, 0x7f, 0x2f, 0x67, 0x3c, 0x63, 0xfe, 0x08, 0x1d, 0xc3,
	0x16, 0xff, 0xa4, 0x84, 0xe6, 0xa6, 0xb7, 0x3f, 0x04, 0xa3, 0x66, 0x78, 0x77, 0x5b, 0x04, 0xb5,
	0xc4, 0x6a, 0x80, 0x3f, 0xc3, 0x7b, 0x67, 0xdc, 0x7a, 0x63, 0xfd, 0x9e, 0x37, 0xb0, 0x69, 0x1d,
	0x2f, 0x9f, 0x13, 0x3e, 0xbf, 0x2c, 0x82, 0xe3, 0x44, 0xea, 0xc5, 0x29, 0x8e, 0xe2, 0x98, 0x1b,
	0x73, 0x31, 0xe3, 0x6b, 0x3c, 0xcc, 0xa3, 0xb9, 0x98, 0x46, 0x96, 0x9f, 0x62, 0xcd, 0x97, 0x2b,
	0xa1, 0xf9, 0x14, 0xb3, 0xd2, 0x87, 0x9e, 0xc2, 0xe6, 0x64, 0x95, 0x57, 0xbf, 0xde, 0x09, 0x8f,
	0x2e, 0x8b, 0xa0, 0x5b, 0x36, 0xf0, 0x32, 0x66, 0xe5, 0x31, 0xfe, 0x02, 0xe0, 0xfd, 0xeb, 0x97,
	0xdf, 0xba, 0xf9, 0xbc, 0xfa, 0x0d, 0x60, 0xfb, 0xbc, 0xca, 0x0b, 0x5a, 0xc2, 0xc3, 0x6b, 0x0b,
	0x45, 0x84, 0xfc, 0x2b, 0x5c, 0xe4, 0x66, 0x54, 0xfa, 0x2f, 0x77, 0xaa, 0x77, 0xd3, 0xc8, 0x60,
	0xf7, 0xea, 0x88, 0xd0, 0xc9, 0x7f, 0x3b, 0x5c, 0xdd, 0x65, 0x9f, 0xee, 0x52, 0xae, 0xe6, 0xeb,
	0xf0, 0xc1, 0x8f, 0x5f, 0x83, 0x3b, 0xdf, 0x36, 0x03, 0xf0, 0x7d, 0x33, 0x00, 0x3f, 0x37, 0x03,
	0xf0, 0x61, 0x2f, 0x1f, 0x4f, 0x0e, 0x7c, 0xfa, 0x5f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x68,
	0xd8, 0x37, 0xd6, 0x7d, 0x03, 0x00, 0x00,
}
