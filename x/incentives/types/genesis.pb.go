// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the incentives module's various parameters when first
// initialized
type GenesisState struct {
	// params are all the parameters of the module
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// gauges are all gauges (not including group gauges) that should exist at
	// genesis
	Gauges []Gauge `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges"`
	// lockable_durations are all lockup durations that gauges can be locked for
	// in order to receive incentives
	LockableDurations []time.Duration `protobuf:"bytes,3,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
	// last_gauge_id is what the gauge number will increment from when creating
	// the next gauge after genesis
	LastGaugeId uint64 `protobuf:"varint,4,opt,name=last_gauge_id,json=lastGaugeId,proto3" json:"last_gauge_id,omitempty"`
	// gauges are all group gauges that should exist at genesis
	GroupGauges []Gauge `protobuf:"bytes,5,rep,name=group_gauges,json=groupGauges,proto3" json:"group_gauges"`
	// groups are all the groups that should exist at genesis
	Groups []Group `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a288ccc95d977d2d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGauges() []Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func (m *GenesisState) GetLastGaugeId() uint64 {
	if m != nil {
		return m.LastGaugeId
	}
	return 0
}

func (m *GenesisState) GetGroupGauges() []Gauge {
	if m != nil {
		return m.GroupGauges
	}
	return nil
}

func (m *GenesisState) GetGroups() []Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.incentives.GenesisState")
}

func init() { proto.RegisterFile("osmosis/incentives/genesis.proto", fileDescriptor_a288ccc95d977d2d) }

var fileDescriptor_a288ccc95d977d2d = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0xe1, 0xdf, 0xfe, 0x39, 0x40, 0x3d, 0xb8, 0xf1, 0x40, 0x7b, 0x00, 0x42, 0x62, 0xd2,
	0x8b, 0x6c, 0xac, 0x89, 0x1a, 0x8f, 0xc4, 0xa4, 0xf1, 0xd6, 0xd4, 0x9b, 0x17, 0xb2, 0xb4, 0x2b,
	0x12, 0x81, 0x25, 0xcc, 0xd2, 0xd8, 0xb7, 0x30, 0xf1, 0xe2, 0x23, 0xf5, 0xd8, 0xa3, 0xa7, 0x6a,
	0xda, 0x37, 0xf0, 0x09, 0x0c, 0x0b, 0xab, 0x26, 0x6d, 0x1a, 0x6f, 0x0c, 0xf3, 0x9b, 0x6f, 0xe6,
	0xfb, 0x56, 0x77, 0x18, 0xa4, 0x0c, 0x62, 0xc0, 0x71, 0x36, 0xa1, 0x19, 0x8f, 0x67, 0x14, 0x70,
	0x44, 0x33, 0x0a, 0x31, 0x78, 0x79, 0xc1, 0x38, 0x43, 0xa8, 0x21, 0xbc, 0x1f, 0xa2, 0x77, 0x14,
	0xb1, 0x88, 0x89, 0x36, 0xae, 0xbe, 0x6a, 0xb2, 0x67, 0x45, 0x8c, 0x45, 0x09, 0xc5, 0xa2, 0x0a,
	0xcb, 0x7b, 0x3c, 0x2d, 0x0b, 0xc2, 0x63, 0x96, 0x35, 0x7d, 0x7b, 0xc7, 0xae, 0x9c, 0x14, 0x24,
	0x05, 0x29, 0xb0, 0xeb, 0x18, 0x52, 0x46, 0x74, 0x5f, 0xbf, 0x60, 0x65, 0x5e, 0xf7, 0xdd, 0x97,
	0x96, 0xde, 0x19, 0xd6, 0xc7, 0xdf, 0x72, 0xc2, 0x29, 0xba, 0xd4, 0xb5, 0x7a, 0x81, 0xa9, 0x3a,
	0x6a, 0xdf, 0x18, 0xf4, 0xbc, 0x6d, 0x33, 0xde, 0x48, 0x10, 0x7e, 0x7b, 0xb1, 0xb2, 0x95, 0x71,
	0xc3, 0xa3, 0x0b, 0x5d, 0x13, 0x9b, 0xc1, 0xfc, 0xe7, 0xb4, 0xfa, 0xc6, 0xa0, 0xbb, 0x6b, 0x72,
	0x58, 0x11, 0x72, 0xb0, 0xc6, 0x11, 0xd3, 0x51, 0xc2, 0x26, 0x8f, 0x24, 0x4c, 0x68, 0x20, 0xfd,
	0x83, 0xd9, 0x6a, 0x44, 0xea, 0x84, 0x3c, 0x99, 0x90, 0x77, 0xdd, 0x10, 0xfe, 0x71, 0x25, 0xf2,
	0xb9, 0xb2, 0xbb, 0x73, 0x92, 0x26, 0x57, 0xee, 0xb6, 0x84, 0xfb, 0xfa, 0x6e, 0xab, 0xe3, 0x43,
	0xd9, 0x90, 0x83, 0x80, 0x5c, 0xfd, 0x20, 0x21, 0xc0, 0x03, 0xb1, 0x3f, 0x88, 0xa7, 0x66, 0xdb,
	0x51, 0xfb, 0xed, 0xb1, 0x51, 0xfd, 0x14, 0x07, 0xde, 0x4c, 0x91, 0xaf, 0x77, 0x44, 0x4e, 0x41,
	0xe3, 0xe9, 0xff, 0xdf, 0x3c, 0x19, 0x62, 0x68, 0x58, 0x1b, 0xab, 0x12, 0xa9, 0x4a, 0x30, 0xb5,
	0x3d, 0xd3, 0x15, 0xf1, 0x9d, 0x88, 0xc0, 0xfd, 0xd1, 0x62, 0x6d, 0xa9, 0xcb, 0xb5, 0xa5, 0x7e,
	0xac, 0x2d, 0xf5, 0x79, 0x63, 0x29, 0xcb, 0x8d, 0xa5, 0xbc, 0x6d, 0x2c, 0xe5, 0xee, 0x3c, 0x8a,
	0xf9, 0x43, 0x19, 0x7a, 0x13, 0x96, 0xe2, 0x46, 0xec, 0x24, 0x21, 0x21, 0xc8, 0x02, 0xcf, 0x06,
	0xa7, 0xf8, 0xe9, 0xf7, 0x6b, 0xf3, 0x79, 0x4e, 0x21, 0xd4, 0x44, 0x7e, 0x67, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x8a, 0xe3, 0xc3, 0xbd, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.GroupGauges) > 0 {
		for iNdEx := len(m.GroupGauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GroupGauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.LastGaugeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastGaugeId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGenesis(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastGaugeId != 0 {
		n += 1 + sovGenesis(uint64(m.LastGaugeId))
	}
	if len(m.GroupGauges) > 0 {
		for _, e := range m.GroupGauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gauges = append(m.Gauges, Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastGaugeId", wireType)
			}
			m.LastGaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastGaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupGauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupGauges = append(m.GroupGauges, Gauge{})
			if err := m.GroupGauges[len(m.GroupGauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, Group{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
