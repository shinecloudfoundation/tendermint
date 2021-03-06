// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/evidence/msgs.proto

package evidence

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
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

type List struct {
	Evidence []*types.Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence,omitempty"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_241b0b668c736d2f, []int{0}
}
func (m *List) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_List.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return m.Size()
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

func (m *List) GetEvidence() []*types.Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

type Info struct {
	Committed bool           `protobuf:"varint,1,opt,name=committed,proto3" json:"committed,omitempty"`
	Priority  int64          `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Evidence  types.Evidence `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_241b0b668c736d2f, []int{1}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Info.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return m.Size()
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetCommitted() bool {
	if m != nil {
		return m.Committed
	}
	return false
}

func (m *Info) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Info) GetEvidence() types.Evidence {
	if m != nil {
		return m.Evidence
	}
	return types.Evidence{}
}

func init() {
	proto.RegisterType((*List)(nil), "tendermint.evidence.List")
	proto.RegisterType((*Info)(nil), "tendermint.evidence.Info")
}

func init() { proto.RegisterFile("tendermint/evidence/msgs.proto", fileDescriptor_241b0b668c736d2f) }

var fileDescriptor_241b0b668c736d2f = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x4f, 0x2d, 0xcb, 0x4c, 0x49, 0xcd, 0x4b, 0x4e, 0xd5,
	0xcf, 0x2d, 0x4e, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x46, 0xc8, 0xeb, 0xc1,
	0xe4, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xf2, 0xfa, 0x20, 0x16, 0x44, 0xa9, 0x94, 0x3c,
	0x92, 0x51, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x70, 0x03, 0x21, 0x0a, 0x94, 0xec, 0xb8, 0x58, 0x7c,
	0x32, 0x8b, 0x4b, 0x84, 0xcc, 0xb8, 0x38, 0x60, 0x32, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0x52, 0x7a, 0x48, 0xd6, 0x80, 0xf5, 0xea, 0xb9, 0x42, 0x55, 0x04, 0xc1, 0xd5, 0x2a, 0xd5, 0x71,
	0xb1, 0x78, 0xe6, 0xa5, 0xe5, 0x0b, 0xc9, 0x70, 0x71, 0x26, 0xe7, 0xe7, 0xe6, 0x66, 0x96, 0x94,
	0xa4, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x0a,
	0x8a, 0x32, 0xf3, 0x8b, 0x32, 0x4b, 0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0xe0, 0x7c,
	0x21, 0x1b, 0x24, 0x9b, 0x99, 0x15, 0x18, 0xf1, 0xdb, 0xec, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x03,
	0xc2, 0x7e, 0xa7, 0x90, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x4a,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x0e, 0x05, 0x04, 0x13, 0x12,
	0x5a, 0x58, 0x02, 0x3b, 0x89, 0x0d, 0x2c, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xec, 0x27,
	0xa7, 0x19, 0x8a, 0x01, 0x00, 0x00,
}

func (m *List) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *List) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *List) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Priority != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x10
	}
	if m.Committed {
		i--
		if m.Committed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *List) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
			l = e.Size()
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	return n
}

func (m *Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Committed {
		n += 2
	}
	if m.Priority != 0 {
		n += 1 + sovMsgs(uint64(m.Priority))
	}
	l = m.Evidence.Size()
	n += 1 + l + sovMsgs(uint64(l))
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *List) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: List: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: List: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidence = append(m.Evidence, &types.Evidence{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Committed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Committed = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
