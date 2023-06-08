// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmiccredit/credit/credit.proto

package types

import (
	fmt "fmt"
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

type Credit struct {
	Owner      string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Credited   uint64 `protobuf:"varint,2,opt,name=credited,proto3" json:"credited,omitempty"`
	Rewarding  string `protobuf:"bytes,3,opt,name=rewarding,proto3" json:"rewarding,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Bond       uint64 `protobuf:"varint,5,opt,name=bond,proto3" json:"bond,omitempty"`
}

func (m *Credit) Reset()         { *m = Credit{} }
func (m *Credit) String() string { return proto.CompactTextString(m) }
func (*Credit) ProtoMessage()    {}
func (*Credit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a29330c50cf2b1a3, []int{0}
}
func (m *Credit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credit.Merge(m, src)
}
func (m *Credit) XXX_Size() int {
	return m.Size()
}
func (m *Credit) XXX_DiscardUnknown() {
	xxx_messageInfo_Credit.DiscardUnknown(m)
}

var xxx_messageInfo_Credit proto.InternalMessageInfo

func (m *Credit) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Credit) GetCredited() uint64 {
	if m != nil {
		return m.Credited
	}
	return 0
}

func (m *Credit) GetRewarding() string {
	if m != nil {
		return m.Rewarding
	}
	return ""
}

func (m *Credit) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *Credit) GetBond() uint64 {
	if m != nil {
		return m.Bond
	}
	return 0
}

func init() {
	proto.RegisterType((*Credit)(nil), "cosmiccredit.credit.Credit")
}

func init() { proto.RegisterFile("cosmiccredit/credit/credit.proto", fileDescriptor_a29330c50cf2b1a3) }

var fileDescriptor_a29330c50cf2b1a3 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x4c, 0x4e, 0x2e, 0x4a, 0x4d, 0xc9, 0x2c, 0xd1, 0x47, 0xa1, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x84, 0x91, 0x55, 0xe8, 0x41, 0x28, 0xa5, 0x1e, 0x46, 0x2e, 0x36, 0x67, 0x30, 0x53,
	0x48, 0x84, 0x8b, 0x35, 0xbf, 0x3c, 0x2f, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xc2, 0x11, 0x92, 0xe2, 0xe2, 0x80, 0x28, 0x4d, 0x4d, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09,
	0x82, 0xf3, 0x85, 0x64, 0xb8, 0x38, 0x8b, 0x52, 0xcb, 0x13, 0x8b, 0x52, 0x32, 0xf3, 0xd2, 0x25,
	0x98, 0xc1, 0xba, 0x10, 0x02, 0x42, 0x72, 0x5c, 0x5c, 0xc9, 0xf9, 0x39, 0x39, 0x89, 0x25, 0xa9,
	0x45, 0x89, 0x39, 0x12, 0x2c, 0x60, 0x69, 0x24, 0x11, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0xbc,
	0x14, 0x09, 0x56, 0xb0, 0xa9, 0x60, 0xb6, 0x93, 0xd9, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0xc9, 0x40, 0x5c, 0xaf, 0x0b, 0xf5, 0x59, 0x05, 0xcc, 0x8b, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x2f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x13, 0xf8, 0x1c,
	0x8c, 0x06, 0x01, 0x00, 0x00,
}

func (m *Credit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bond != 0 {
		i = encodeVarintCredit(dAtA, i, uint64(m.Bond))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintCredit(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Rewarding) > 0 {
		i -= len(m.Rewarding)
		copy(dAtA[i:], m.Rewarding)
		i = encodeVarintCredit(dAtA, i, uint64(len(m.Rewarding)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Credited != 0 {
		i = encodeVarintCredit(dAtA, i, uint64(m.Credited))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCredit(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCredit(dAtA []byte, offset int, v uint64) int {
	offset -= sovCredit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Credit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCredit(uint64(l))
	}
	if m.Credited != 0 {
		n += 1 + sovCredit(uint64(m.Credited))
	}
	l = len(m.Rewarding)
	if l > 0 {
		n += 1 + l + sovCredit(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovCredit(uint64(l))
	}
	if m.Bond != 0 {
		n += 1 + sovCredit(uint64(m.Bond))
	}
	return n
}

func sovCredit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCredit(x uint64) (n int) {
	return sovCredit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Credit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCredit
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
			return fmt.Errorf("proto: Credit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredit
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
				return ErrInvalidLengthCredit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credited", wireType)
			}
			m.Credited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Credited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewarding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredit
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
				return ErrInvalidLengthCredit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewarding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredit
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
				return ErrInvalidLengthCredit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCredit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bond", wireType)
			}
			m.Bond = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCredit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bond |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCredit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCredit
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
func skipCredit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCredit
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
					return 0, ErrIntOverflowCredit
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
					return 0, ErrIntOverflowCredit
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
				return 0, ErrInvalidLengthCredit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCredit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCredit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCredit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCredit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCredit = fmt.Errorf("proto: unexpected end of group")
)
