// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/traits.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EXPERIMENTAL.
// NOTE: Please refer from using MutabilityMode for the time being. It will be replaced in the future (ROX-14276).
// MutabilityMode specifies whether and how an object can be modified. Default
// is ALLOW_MUTATE and means there are no modification restrictions; this is equivalent
// to the absence of MutabilityMode specification. ALLOW_MUTATE_FORCED forbids all
// modifying operations except object removal with force bit on.
//
// Be careful when changing the state of this field. For example, modifying an
// object from ALLOW_MUTATE to ALLOW_MUTATE_FORCED is allowed but will prohibit any further
// changes to it, including modifying it back to ALLOW_MUTATE.
type Traits_MutabilityMode int32

const (
	Traits_ALLOW_MUTATE        Traits_MutabilityMode = 0
	Traits_ALLOW_MUTATE_FORCED Traits_MutabilityMode = 1
)

var Traits_MutabilityMode_name = map[int32]string{
	0: "ALLOW_MUTATE",
	1: "ALLOW_MUTATE_FORCED",
}

var Traits_MutabilityMode_value = map[string]int32{
	"ALLOW_MUTATE":        0,
	"ALLOW_MUTATE_FORCED": 1,
}

func (x Traits_MutabilityMode) String() string {
	return proto.EnumName(Traits_MutabilityMode_name, int32(x))
}

func (Traits_MutabilityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0, 0}
}

// EXPERIMENTAL.
// visibility allows to specify whether the object should be visible for certain APIs.
type Traits_Visibility int32

const (
	Traits_VISIBLE Traits_Visibility = 0
	Traits_HIDDEN  Traits_Visibility = 1
)

var Traits_Visibility_name = map[int32]string{
	0: "VISIBLE",
	1: "HIDDEN",
}

var Traits_Visibility_value = map[string]int32{
	"VISIBLE": 0,
	"HIDDEN":  1,
}

func (x Traits_Visibility) String() string {
	return proto.EnumName(Traits_Visibility_name, int32(x))
}

func (Traits_Visibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0, 1}
}

// Origin specifies the origin of an object.
// Objects can have three different origins:
// - IMPERATIVE: the object was created via the API. This is assumed by default.
// - DEFAULT: the object is a default object, such as default roles, access scopes etc.
// - DECLARATIVE: the object is created via declarative configuration.
// Based on the origin, different rules apply to the objects.
// Objects with the DECLARATIVE origin are not allowed to be modified via API, only via declarative configuration.
// Additionally, they may not reference objects with the IMPERATIVE origin.
// Objects with the DEFAULT origin are not allowed to be modified via either API or declarative configuration.
// They may be referenced by all other objects.
// Objects with the IMPERATIVE origin are allowed to be modified via API, not via declarative configuration.
// They may reference all other objects.
type Traits_Origin int32

const (
	Traits_IMPERATIVE  Traits_Origin = 0
	Traits_DEFAULT     Traits_Origin = 1
	Traits_DECLARATIVE Traits_Origin = 2
)

var Traits_Origin_name = map[int32]string{
	0: "IMPERATIVE",
	1: "DEFAULT",
	2: "DECLARATIVE",
}

var Traits_Origin_value = map[string]int32{
	"IMPERATIVE":  0,
	"DEFAULT":     1,
	"DECLARATIVE": 2,
}

func (x Traits_Origin) String() string {
	return proto.EnumName(Traits_Origin_name, int32(x))
}

func (Traits_Origin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0, 2}
}

type Traits struct {
	MutabilityMode       Traits_MutabilityMode `protobuf:"varint,1,opt,name=mutability_mode,json=mutabilityMode,proto3,enum=storage.Traits_MutabilityMode" json:"mutability_mode,omitempty"`
	Visibility           Traits_Visibility     `protobuf:"varint,2,opt,name=visibility,proto3,enum=storage.Traits_Visibility" json:"visibility,omitempty"`
	Origin               Traits_Origin         `protobuf:"varint,3,opt,name=origin,proto3,enum=storage.Traits_Origin" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Traits) Reset()         { *m = Traits{} }
func (m *Traits) String() string { return proto.CompactTextString(m) }
func (*Traits) ProtoMessage()    {}
func (*Traits) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0}
}
func (m *Traits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Traits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Traits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Traits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Traits.Merge(m, src)
}
func (m *Traits) XXX_Size() int {
	return m.Size()
}
func (m *Traits) XXX_DiscardUnknown() {
	xxx_messageInfo_Traits.DiscardUnknown(m)
}

var xxx_messageInfo_Traits proto.InternalMessageInfo

func (m *Traits) GetMutabilityMode() Traits_MutabilityMode {
	if m != nil {
		return m.MutabilityMode
	}
	return Traits_ALLOW_MUTATE
}

func (m *Traits) GetVisibility() Traits_Visibility {
	if m != nil {
		return m.Visibility
	}
	return Traits_VISIBLE
}

func (m *Traits) GetOrigin() Traits_Origin {
	if m != nil {
		return m.Origin
	}
	return Traits_IMPERATIVE
}

func (m *Traits) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Traits) Clone() *Traits {
	if m == nil {
		return nil
	}
	cloned := new(Traits)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.Traits_MutabilityMode", Traits_MutabilityMode_name, Traits_MutabilityMode_value)
	proto.RegisterEnum("storage.Traits_Visibility", Traits_Visibility_name, Traits_Visibility_value)
	proto.RegisterEnum("storage.Traits_Origin", Traits_Origin_name, Traits_Origin_value)
	proto.RegisterType((*Traits)(nil), "storage.Traits")
}

func init() { proto.RegisterFile("storage/traits.proto", fileDescriptor_ec31914177d462a1) }

var fileDescriptor_ec31914177d462a1 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4b, 0x84, 0x40,
	0x14, 0xc6, 0x77, 0x0c, 0x14, 0x9e, 0xe1, 0xca, 0x14, 0x65, 0x1d, 0x24, 0x84, 0xa0, 0xd3, 0x04,
	0x15, 0x1d, 0xea, 0xe4, 0xae, 0xb3, 0x25, 0x68, 0x86, 0xb9, 0x06, 0x5d, 0xc4, 0x6d, 0x65, 0x19,
	0xca, 0x26, 0x74, 0x8a, 0xfa, 0x4f, 0xfa, 0x93, 0x3a, 0x76, 0xed, 0x16, 0xf6, 0x8f, 0xc4, 0xaa,
	0xbb, 0xed, 0xee, 0x71, 0xbe, 0xef, 0xf7, 0xbd, 0x79, 0x8f, 0x0f, 0x36, 0x4b, 0xc1, 0x8b, 0x74,
	0x92, 0x1d, 0x8a, 0x22, 0x65, 0xa2, 0x24, 0xcf, 0x05, 0x17, 0x1c, 0x2b, 0xad, 0x6a, 0x7d, 0x4b,
	0x20, 0x47, 0xb5, 0x83, 0x2f, 0xa0, 0x9b, 0xbf, 0x88, 0x74, 0xc4, 0x1e, 0x99, 0x78, 0x4f, 0x72,
	0x3e, 0xce, 0x0c, 0xb4, 0x87, 0x0e, 0xb4, 0x23, 0x93, 0xb4, 0x34, 0x69, 0x48, 0xe2, 0xcf, 0x31,
	0x9f, 0x8f, 0xb3, 0x50, 0xcb, 0x97, 0xde, 0xf8, 0x0c, 0xe0, 0x95, 0x95, 0xac, 0x51, 0x0c, 0xa9,
	0x9e, 0xb1, 0xbb, 0x3a, 0x23, 0x9e, 0x13, 0xe1, 0x02, 0x8d, 0x09, 0xc8, 0xbc, 0x60, 0x13, 0xf6,
	0x64, 0xac, 0xd5, 0xb9, 0xad, 0xd5, 0x5c, 0x50, 0xbb, 0x61, 0x4b, 0x59, 0xe7, 0xa0, 0x2d, 0x6f,
	0x83, 0x75, 0x58, 0xb7, 0x3d, 0x2f, 0xb8, 0x4d, 0xfc, 0x61, 0x64, 0x47, 0x54, 0xef, 0xe0, 0x6d,
	0xd8, 0x58, 0x54, 0x92, 0x41, 0x10, 0xf6, 0xa9, 0xa3, 0x23, 0x6b, 0x1f, 0xe0, 0x7f, 0x0d, 0xac,
	0x82, 0x12, 0xbb, 0x37, 0x6e, 0xcf, 0x9b, 0x66, 0x00, 0xe4, 0x4b, 0xd7, 0x71, 0xe8, 0x95, 0x8e,
	0xac, 0x53, 0x90, 0x9b, 0x5f, 0xb1, 0x06, 0xe0, 0xfa, 0xd7, 0x34, 0xb4, 0x23, 0x37, 0x9e, 0x52,
	0x2a, 0x28, 0x0e, 0x1d, 0xd8, 0x43, 0x2f, 0xd2, 0x11, 0xee, 0x82, 0xea, 0xd0, 0xbe, 0x67, 0xb7,
	0xae, 0xd4, 0x3b, 0xf9, 0xac, 0x4c, 0xf4, 0x55, 0x99, 0xe8, 0xa7, 0x32, 0xd1, 0xc7, 0xaf, 0xd9,
	0x81, 0x1d, 0xc6, 0x49, 0x29, 0xd2, 0xfb, 0x87, 0x82, 0xbf, 0x35, 0x4d, 0xcc, 0xce, 0xbb, 0x9b,
	0x35, 0x32, 0x92, 0x6b, 0xfd, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x32, 0x93, 0xfd, 0xa1, 0xb9,
	0x01, 0x00, 0x00,
}

func (m *Traits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Traits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Traits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Origin != 0 {
		i = encodeVarintTraits(dAtA, i, uint64(m.Origin))
		i--
		dAtA[i] = 0x18
	}
	if m.Visibility != 0 {
		i = encodeVarintTraits(dAtA, i, uint64(m.Visibility))
		i--
		dAtA[i] = 0x10
	}
	if m.MutabilityMode != 0 {
		i = encodeVarintTraits(dAtA, i, uint64(m.MutabilityMode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTraits(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraits(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Traits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MutabilityMode != 0 {
		n += 1 + sovTraits(uint64(m.MutabilityMode))
	}
	if m.Visibility != 0 {
		n += 1 + sovTraits(uint64(m.Visibility))
	}
	if m.Origin != 0 {
		n += 1 + sovTraits(uint64(m.Origin))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTraits(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraits(x uint64) (n int) {
	return sovTraits(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Traits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraits
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
			return fmt.Errorf("proto: Traits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Traits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutabilityMode", wireType)
			}
			m.MutabilityMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraits
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MutabilityMode |= Traits_MutabilityMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			m.Visibility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraits
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Visibility |= Traits_Visibility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			m.Origin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraits
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Origin |= Traits_Origin(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraits(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraits
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
func skipTraits(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraits
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
					return 0, ErrIntOverflowTraits
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
					return 0, ErrIntOverflowTraits
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
				return 0, ErrInvalidLengthTraits
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTraits
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTraits
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTraits        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraits          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTraits = fmt.Errorf("proto: unexpected end of group")
)
