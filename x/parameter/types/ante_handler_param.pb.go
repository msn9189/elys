// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/parameter/ante_handler_param.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AnteHandlerParam struct {
	MinCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=minCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"minCommissionRate"`
	MaxVotingPower    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=maxVotingPower,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"maxVotingPower"`
	MinSelfDelegation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=minSelfDelegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minSelfDelegation"`
}

func (m *AnteHandlerParam) Reset()         { *m = AnteHandlerParam{} }
func (m *AnteHandlerParam) String() string { return proto.CompactTextString(m) }
func (*AnteHandlerParam) ProtoMessage()    {}
func (*AnteHandlerParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ae17167f40c52b, []int{0}
}
func (m *AnteHandlerParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnteHandlerParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnteHandlerParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnteHandlerParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnteHandlerParam.Merge(m, src)
}
func (m *AnteHandlerParam) XXX_Size() int {
	return m.Size()
}
func (m *AnteHandlerParam) XXX_DiscardUnknown() {
	xxx_messageInfo_AnteHandlerParam.DiscardUnknown(m)
}

var xxx_messageInfo_AnteHandlerParam proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AnteHandlerParam)(nil), "elysnetwork.elys.parameter.AnteHandlerParam")
}

func init() {
	proto.RegisterFile("elys/parameter/ante_handler_param.proto", fileDescriptor_61ae17167f40c52b)
}

var fileDescriptor_61ae17167f40c52b = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x67, 0x2a, 0x08, 0xce, 0x42, 0xb4, 0xb8, 0x28, 0x5d, 0xa4, 0xe2, 0x42, 0xdd, 0x34,
	0xa1, 0xf8, 0x04, 0xd6, 0x2e, 0xea, 0xae, 0x54, 0xe8, 0x42, 0x84, 0x92, 0x99, 0x5e, 0xa7, 0xa1,
	0x93, 0xdc, 0x92, 0x5c, 0x6d, 0xfb, 0x16, 0x3e, 0x85, 0xcf, 0xd2, 0x65, 0x97, 0xe2, 0xa2, 0xc8,
	0xcc, 0x8b, 0xc8, 0xfc, 0x20, 0xc5, 0xae, 0x74, 0x95, 0x9b, 0xe4, 0x9c, 0xef, 0x84, 0x9c, 0xe0,
	0x0a, 0x92, 0x95, 0x13, 0x73, 0x69, 0xa5, 0x06, 0x02, 0x2b, 0xa4, 0x21, 0x18, 0x4f, 0xa5, 0x99,
	0x24, 0x60, 0xc7, 0xc5, 0x31, 0x9f, 0x5b, 0x24, 0xac, 0x37, 0x73, 0xa1, 0x01, 0x5a, 0xa0, 0x9d,
	0xf1, 0x7c, 0xe6, 0x3f, 0xa6, 0xe6, 0x59, 0x8c, 0x31, 0x16, 0x32, 0x91, 0x4f, 0xa5, 0xa3, 0xc9,
	0x22, 0x74, 0x1a, 0x9d, 0x08, 0xa5, 0x03, 0xf1, 0xda, 0x09, 0x81, 0x64, 0x47, 0x44, 0xa8, 0x4c,
	0x79, 0x7f, 0xf1, 0x5e, 0x0b, 0x4e, 0x6e, 0x0d, 0x41, 0xbf, 0x4c, 0x1b, 0xe4, 0xb8, 0xfa, 0x53,
	0x70, 0xaa, 0x95, 0xb9, 0x43, 0xad, 0x95, 0x73, 0x0a, 0xcd, 0x50, 0x12, 0x34, 0xfc, 0x73, 0xff,
	0xfa, 0xa8, 0xcb, 0xd7, 0xdb, 0x96, 0xf7, 0xb9, 0x6d, 0x5d, 0xc6, 0x8a, 0xa6, 0x2f, 0x21, 0x8f,
	0x50, 0x8b, 0x2a, 0xa2, 0x5c, 0xda, 0x6e, 0x32, 0x13, 0xb4, 0x9a, 0x83, 0xe3, 0x3d, 0x88, 0x86,
	0xfb, 0xa0, 0xfa, 0x28, 0x38, 0xd6, 0x72, 0x39, 0x42, 0x52, 0x26, 0x1e, 0xe0, 0x02, 0x6c, 0xa3,
	0xf6, 0x2f, 0xf4, 0x2f, 0x4a, 0xf5, 0xea, 0x07, 0x48, 0x9e, 0x7b, 0x90, 0x40, 0x2c, 0x49, 0xa1,
	0x69, 0x1c, 0xfc, 0x19, 0x7d, 0x6f, 0x68, 0xb8, 0x0f, 0xea, 0xf6, 0xd7, 0x29, 0xf3, 0x37, 0x29,
	0xf3, 0xbf, 0x52, 0xe6, 0xbf, 0x65, 0xcc, 0xdb, 0x64, 0xcc, 0xfb, 0xc8, 0x98, 0xf7, 0xc8, 0x77,
	0xa0, 0x79, 0x27, 0xed, 0xaa, 0xa0, 0x62, 0x23, 0x96, 0x3b, 0xbd, 0x16, 0x01, 0xe1, 0x61, 0xf1,
	0xf3, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x48, 0x22, 0x58, 0xf6, 0x01, 0x00, 0x00,
}

func (m *AnteHandlerParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnteHandlerParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnteHandlerParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinSelfDelegation.Size()
		i -= size
		if _, err := m.MinSelfDelegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxVotingPower.Size()
		i -= size
		if _, err := m.MaxVotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAnteHandlerParam(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAnteHandlerParam(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnteHandlerParam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnteHandlerParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MaxVotingPower.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	l = m.MinSelfDelegation.Size()
	n += 1 + l + sovAnteHandlerParam(uint64(l))
	return n
}

func sovAnteHandlerParam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnteHandlerParam(x uint64) (n int) {
	return sovAnteHandlerParam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnteHandlerParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnteHandlerParam
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
			return fmt.Errorf("proto: AnteHandlerParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnteHandlerParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxVotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnteHandlerParam
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
				return ErrInvalidLengthAnteHandlerParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnteHandlerParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnteHandlerParam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnteHandlerParam
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
func skipAnteHandlerParam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
					return 0, ErrIntOverflowAnteHandlerParam
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
				return 0, ErrInvalidLengthAnteHandlerParam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnteHandlerParam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnteHandlerParam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnteHandlerParam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnteHandlerParam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnteHandlerParam = fmt.Errorf("proto: unexpected end of group")
)
