// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block/asset.proto

package types

import (
	fmt "fmt"
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

type Asset struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Denom     string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	Decimal   string `protobuf:"bytes,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
	Price     string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	AppId     string `protobuf:"bytes,6,opt,name=appId,proto3" json:"appId,omitempty"`
	IbcStatus string `protobuf:"bytes,7,opt,name=ibcStatus,proto3" json:"ibcStatus,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a4cda20d7ab9845, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Asset) GetDecimal() string {
	if m != nil {
		return m.Decimal
	}
	return ""
}

func (m *Asset) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Asset) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Asset) GetIbcStatus() string {
	if m != nil {
		return m.IbcStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "block.block.Asset")
}

func init() { proto.RegisterFile("block/asset.proto", fileDescriptor_2a4cda20d7ab9845) }

var fileDescriptor_2a4cda20d7ab9845 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xca, 0xc9, 0x4f,
	0xce, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06,
	0x0b, 0xe9, 0x81, 0x49, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb8, 0x3e, 0x88, 0x05, 0x51,
	0x22, 0x25, 0x04, 0xd1, 0x55, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x0c, 0x11, 0x53, 0x5a, 0xca, 0xc8,
	0xc5, 0xea, 0x08, 0x32, 0x46, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x25, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0x49, 0xcd, 0xcb, 0xcf, 0x95, 0x60,
	0x06, 0x0b, 0x42, 0x38, 0x42, 0x12, 0x5c, 0xec, 0x29, 0xa9, 0xc9, 0x99, 0xb9, 0x89, 0x39, 0x12,
	0x2c, 0x60, 0x71, 0x18, 0x17, 0xa4, 0xbe, 0xa0, 0x28, 0x33, 0x39, 0x55, 0x82, 0x15, 0xa2, 0x1e,
	0xcc, 0x01, 0x89, 0x26, 0x16, 0x14, 0x78, 0xa6, 0x48, 0xb0, 0x41, 0x44, 0xc1, 0x1c, 0x21, 0x19,
	0x2e, 0xce, 0xcc, 0xa4, 0xe4, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x76, 0xb0, 0x0c, 0x42,
	0xc0, 0x49, 0xf7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x84, 0x21, 0xbe,
	0xaa, 0xd0, 0x87, 0xd0, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xdf, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xca, 0xfc, 0x71, 0x29, 0x01, 0x00, 0x00,
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcStatus) > 0 {
		i -= len(m.IbcStatus)
		copy(dAtA[i:], m.IbcStatus)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.IbcStatus)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.AppId) > 0 {
		i -= len(m.AppId)
		copy(dAtA[i:], m.AppId)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.AppId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Decimal) > 0 {
		i -= len(m.Decimal)
		copy(dAtA[i:], m.Decimal)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Decimal)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAsset(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Decimal)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.AppId)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.IbcStatus)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	return n
}

func sovAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Decimal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsset = fmt.Errorf("proto: unexpected end of group")
)
