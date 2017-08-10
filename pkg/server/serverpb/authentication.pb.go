// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/server/serverpb/authentication.proto
// DO NOT EDIT!

package serverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// UserLoginRequest contains credentials a user must provide to log in.
type UserLoginRequest struct {
	// A username which must correspond to a database user on the cluster.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// A password for the provided username.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptorAuthentication, []int{0} }

// UserLoginResponse is currently empty. If a login is successful, an HTTP
// Set-Cookie header will be added to the response with a session
// cookie identifying the created session.
type UserLoginResponse struct {
}

func (m *UserLoginResponse) Reset()                    { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()               {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptorAuthentication, []int{1} }

// UserLogoutRequest will terminate the current session in use. The request
// is empty because the current session is identified by an HTTP cookie on the
// incoming request.
type UserLogoutRequest struct {
}

func (m *UserLogoutRequest) Reset()                    { *m = UserLogoutRequest{} }
func (m *UserLogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLogoutRequest) ProtoMessage()               {}
func (*UserLogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptorAuthentication, []int{2} }

type UserLogoutResponse struct {
}

func (m *UserLogoutResponse) Reset()                    { *m = UserLogoutResponse{} }
func (m *UserLogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLogoutResponse) ProtoMessage()               {}
func (*UserLogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptorAuthentication, []int{3} }

// SessionCookie is a message used to encode the authentication cookie returned
// from successful login requests.
type SessionCookie struct {
	// The unique ID of the session.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The secret needed to verify ownership of a session.
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *SessionCookie) Reset()                    { *m = SessionCookie{} }
func (m *SessionCookie) String() string            { return proto.CompactTextString(m) }
func (*SessionCookie) ProtoMessage()               {}
func (*SessionCookie) Descriptor() ([]byte, []int) { return fileDescriptorAuthentication, []int{4} }

func init() {
	proto.RegisterType((*UserLoginRequest)(nil), "cockroach.server.serverpb.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "cockroach.server.serverpb.UserLoginResponse")
	proto.RegisterType((*UserLogoutRequest)(nil), "cockroach.server.serverpb.UserLogoutRequest")
	proto.RegisterType((*UserLogoutResponse)(nil), "cockroach.server.serverpb.UserLogoutResponse")
	proto.RegisterType((*SessionCookie)(nil), "cockroach.server.serverpb.SessionCookie")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authentication service

type AuthenticationClient interface {
	// UserLogin is used to create a web authentication session.
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// UserLogout terminates an active authentication session.
	UserLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := grpc.Invoke(ctx, "/cockroach.server.serverpb.Authentication/UserLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) UserLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error) {
	out := new(UserLogoutResponse)
	err := grpc.Invoke(ctx, "/cockroach.server.serverpb.Authentication/UserLogout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// UserLogin is used to create a web authentication session.
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// UserLogout terminates an active authentication session.
	UserLogout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Authentication/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Authentication/UserLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).UserLogout(ctx, req.(*UserLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _Authentication_UserLogin_Handler,
		},
		{
			MethodName: "UserLogout",
			Handler:    _Authentication_UserLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cockroach/pkg/server/serverpb/authentication.proto",
}

func (m *UserLoginRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *UserLoginResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserLogoutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLogoutRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserLogoutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLogoutResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SessionCookie) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionCookie) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(m.Id))
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	return i, nil
}

func encodeFixed64Authentication(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Authentication(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAuthentication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserLoginRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func (m *UserLoginResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UserLogoutRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UserLogoutResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SessionCookie) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuthentication(uint64(m.Id))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func sovAuthentication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuthentication(x uint64) (n int) {
	return sovAuthentication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserLoginRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLoginRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
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
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
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
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLoginResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLoginResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLogoutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLogoutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLogoutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLogoutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLogoutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLogoutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *SessionCookie) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: SessionCookie: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionCookie: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = append(m.Secret[:0], dAtA[iNdEx:postIndex]...)
			if m.Secret == nil {
				m.Secret = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func skipAuthentication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthentication
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
					return 0, ErrIntOverflowAuthentication
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
					return 0, ErrIntOverflowAuthentication
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
				return 0, ErrInvalidLengthAuthentication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuthentication
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
				next, err := skipAuthentication(dAtA[start:])
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
	ErrInvalidLengthAuthentication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthentication   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("cockroach/pkg/server/serverpb/authentication.proto", fileDescriptorAuthentication)
}

var fileDescriptorAuthentication = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0xbb, 0x79, 0x50, 0xda, 0xe5, 0xbd, 0xbe, 0xbe, 0x7d, 0x45, 0x6a, 0x90, 0x20, 0x39,
	0x89, 0x3f, 0xb2, 0x58, 0x0f, 0x82, 0x37, 0xf5, 0x26, 0x9e, 0x22, 0x5e, 0xbc, 0xc8, 0x36, 0x1d,
	0xd2, 0xa5, 0x75, 0x27, 0xee, 0x6e, 0xea, 0xbd, 0xfe, 0x03, 0x82, 0x47, 0xff, 0xa1, 0x1e, 0x05,
	0x2f, 0x1e, 0x35, 0xfa, 0x87, 0x48, 0xd3, 0xf4, 0x97, 0x20, 0xf4, 0x14, 0x66, 0xe6, 0x3b, 0x33,
	0x9f, 0xf9, 0x66, 0x69, 0x2b, 0xc2, 0xa8, 0xa7, 0x51, 0x44, 0x5d, 0x9e, 0xf4, 0x62, 0x6e, 0x40,
	0x0f, 0x40, 0x17, 0x9f, 0xa4, 0xcd, 0x45, 0x6a, 0xbb, 0xa0, 0xac, 0x8c, 0x84, 0x95, 0xa8, 0x82,
	0x44, 0xa3, 0x45, 0xb6, 0x3e, 0xeb, 0x09, 0x26, 0xc2, 0x60, 0xaa, 0x77, 0x37, 0x62, 0xc4, 0xb8,
	0x0f, 0x5c, 0x24, 0x92, 0x0b, 0xa5, 0xd0, 0xe6, 0x7d, 0x66, 0xd2, 0xe8, 0x9f, 0xd1, 0xfa, 0xa5,
	0x01, 0x7d, 0x8e, 0xb1, 0x54, 0x21, 0xdc, 0xa6, 0x60, 0x2c, 0x73, 0x69, 0x25, 0x35, 0xa0, 0x95,
	0xb8, 0x81, 0x26, 0xd9, 0x24, 0x5b, 0xd5, 0x70, 0x16, 0x8f, 0x6b, 0x89, 0x30, 0xe6, 0x0e, 0x75,
	0xa7, 0xe9, 0x4c, 0x6a, 0xd3, 0xd8, 0xff, 0x4f, 0xff, 0x2d, 0xcc, 0x32, 0x09, 0x2a, 0x03, 0x0b,
	0x49, 0x4c, 0x6d, 0xb1, 0xc1, 0x6f, 0x50, 0xb6, 0x98, 0x2c, 0xa4, 0x87, 0xf4, 0xcf, 0x05, 0x18,
	0x23, 0x51, 0x9d, 0x22, 0xf6, 0x24, 0xb0, 0x1a, 0x75, 0x64, 0x27, 0x47, 0xf8, 0x15, 0x3a, 0xb2,
	0xc3, 0xd6, 0x68, 0xd9, 0x40, 0xa4, 0xc1, 0xe6, 0xab, 0x7f, 0x87, 0x45, 0xd4, 0x7a, 0x72, 0x68,
	0xed, 0x78, 0xc9, 0x16, 0x36, 0x24, 0xb4, 0x3a, 0x83, 0x61, 0x3b, 0xc1, 0x8f, 0xfe, 0x04, 0xdf,
	0xcf, 0x77, 0x77, 0x57, 0x13, 0x17, 0xd0, 0xee, 0xf0, 0xe5, 0xf3, 0xd1, 0x69, 0xf8, 0x7f, 0xf9,
	0xf5, 0xf8, 0xcf, 0xf0, 0xc1, 0x3e, 0xef, 0x8f, 0x05, 0x47, 0x64, 0x9b, 0xdd, 0x13, 0x4a, 0xe7,
	0x77, 0xb2, 0x15, 0x06, 0xcf, 0x3d, 0x72, 0xf7, 0x56, 0x54, 0x17, 0x1c, 0xcd, 0x9c, 0x83, 0xb1,
	0xfa, 0x12, 0x07, 0xa6, 0xf6, 0xc4, 0x1f, 0xbd, 0x7b, 0xa5, 0x51, 0xe6, 0x91, 0xe7, 0xcc, 0x23,
	0xaf, 0x99, 0x47, 0xde, 0x32, 0x8f, 0x3c, 0x7c, 0x78, 0xa5, 0xab, 0xca, 0x74, 0x60, 0xbb, 0x9c,
	0xbf, 0x86, 0x83, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xad, 0x02, 0x5a, 0x7c, 0x02, 0x00,
	0x00,
}
