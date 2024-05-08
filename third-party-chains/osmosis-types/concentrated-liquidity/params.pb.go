// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentratedliquidity/params.proto

package concentrated_liquidity

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Params struct {
	// authorized_tick_spacing is an array of uint64s that represents the tick
	// spacing values concentrated-liquidity pools can be created with. For
	// example, an authorized_tick_spacing of [1, 10, 30] allows for pools
	// to be created with tick spacing of 1, 10, or 30.
	AuthorizedTickSpacing   []uint64                      `protobuf:"varint,1,rep,packed,name=authorized_tick_spacing,json=authorizedTickSpacing,proto3" json:"authorized_tick_spacing,omitempty" yaml:"authorized_tick_spacing"`
	AuthorizedSpreadFactors []cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,rep,name=authorized_spread_factors,json=authorizedSpreadFactors,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"authorized_spread_factors" yaml:"authorized_spread_factors"`
	// balancer_shares_reward_discount is the rate by which incentives flowing
	// from CL to Balancer pools will be discounted to encourage LPs to migrate.
	// e.g. a rate of 0.05 means Balancer LPs get 5% less incentives than full
	// range CL LPs.
	// This field can range from (0,1]. If set to 1, it indicates that all
	// incentives stay at cl pool.
	BalancerSharesRewardDiscount cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=balancer_shares_reward_discount,json=balancerSharesRewardDiscount,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"balancer_shares_reward_discount" yaml:"balancer_shares_reward_discount"`
	// authorized_quote_denoms is a list of quote denoms that can be used as
	// token1 when creating a pool. We limit the quote assets to a small set for
	// the purposes of having convenient price increments stemming from tick to
	// price conversion. These increments are in a human readable magnitude only
	// for token1 as a quote. For limit orders in the future, this will be a
	// desirable property in terms of UX as to allow users to set limit orders at
	// prices in terms of token1 (quote asset) that are easy to reason about.
	AuthorizedQuoteDenoms []string        `protobuf:"bytes,4,rep,name=authorized_quote_denoms,json=authorizedQuoteDenoms,proto3" json:"authorized_quote_denoms,omitempty" yaml:"authorized_quote_denoms"`
	AuthorizedUptimes     []time.Duration `protobuf:"bytes,5,rep,name=authorized_uptimes,json=authorizedUptimes,proto3,stdduration" json:"duration,omitempty" yaml:"authorized_uptimes"`
	// is_permissionless_pool_creation_enabled is a boolean that determines if
	// concentrated liquidity pools can be created via message. At launch,
	// we consider allowing only governance to create pools, and then later
	// allowing permissionless pool creation by switching this flag to true
	// with a governance proposal.
	IsPermissionlessPoolCreationEnabled bool `protobuf:"varint,6,opt,name=is_permissionless_pool_creation_enabled,json=isPermissionlessPoolCreationEnabled,proto3" json:"is_permissionless_pool_creation_enabled,omitempty" yaml:"is_permissionless_pool_creation_enabled"`
	// unrestricted_pool_creator_whitelist is a list of addresses that are
	// allowed to bypass restrictions on permissionless supercharged pool
	// creation, like pool_creation_enabled, restricted quote assets, no
	// double creation of pools, etc.
	UnrestrictedPoolCreatorWhitelist []string `protobuf:"bytes,7,rep,name=unrestricted_pool_creator_whitelist,json=unrestrictedPoolCreatorWhitelist,proto3" json:"unrestricted_pool_creator_whitelist,omitempty" yaml:"unrestricted_pool_creator_whitelist"`
	HookGasLimit                     uint64   `protobuf:"varint,8,opt,name=hook_gas_limit,json=hookGasLimit,proto3" json:"hook_gas_limit,omitempty" yaml:"hook_gas_limit"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a3f6981164624c, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuthorizedTickSpacing() []uint64 {
	if m != nil {
		return m.AuthorizedTickSpacing
	}
	return nil
}

func (m *Params) GetAuthorizedQuoteDenoms() []string {
	if m != nil {
		return m.AuthorizedQuoteDenoms
	}
	return nil
}

func (m *Params) GetAuthorizedUptimes() []time.Duration {
	if m != nil {
		return m.AuthorizedUptimes
	}
	return nil
}

func (m *Params) GetIsPermissionlessPoolCreationEnabled() bool {
	if m != nil {
		return m.IsPermissionlessPoolCreationEnabled
	}
	return false
}

func (m *Params) GetUnrestrictedPoolCreatorWhitelist() []string {
	if m != nil {
		return m.UnrestrictedPoolCreatorWhitelist
	}
	return nil
}

func (m *Params) GetHookGasLimit() uint64 {
	if m != nil {
		return m.HookGasLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.concentratedliquidity.Params")
}

func init() {
	proto.RegisterFile("osmosis/concentratedliquidity/params.proto", fileDescriptor_42a3f6981164624c)
}

var fileDescriptor_42a3f6981164624c = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0x6e, 0xed, 0x8f, 0x28, 0x82, 0xc1, 0x62, 0xb6, 0x6a, 0x12, 0x52, 0xd0, 0xa5,
	0xb8, 0x09, 0xd4, 0x9b, 0x1e, 0x84, 0x75, 0xb5, 0x97, 0x0a, 0x35, 0x55, 0x84, 0x5e, 0x86, 0xd9,
	0xc9, 0x34, 0x19, 0x36, 0xc9, 0x4b, 0x67, 0x26, 0x96, 0x2d, 0x78, 0x12, 0xc1, 0xa3, 0x07, 0x0f,
	0xfe, 0x19, 0xfe, 0x19, 0x3d, 0xf6, 0x28, 0x1e, 0xa2, 0xb4, 0x37, 0x8f, 0xfb, 0x17, 0xc8, 0x66,
	0x76, 0xdb, 0xac, 0xad, 0xd8, 0x5b, 0xe6, 0xfb, 0xfd, 0xbc, 0x1f, 0x79, 0x3c, 0x9e, 0xbe, 0x06,
	0x22, 0x05, 0xc1, 0x84, 0x4f, 0x20, 0x23, 0x34, 0x93, 0x1c, 0x4b, 0x1a, 0x26, 0x6c, 0xaf, 0x60,
	0x21, 0x93, 0x43, 0x3f, 0xc7, 0x1c, 0xa7, 0xc2, 0xcb, 0x39, 0x48, 0x30, 0xee, 0x4d, 0x58, 0xef,
	0x42, 0x76, 0xa5, 0x45, 0x2a, 0x1f, 0x55, 0xb0, 0xaf, 0x1e, 0x2a, 0x72, 0xe5, 0x56, 0x04, 0x11,
	0x28, 0x7d, 0xfc, 0x35, 0x51, 0xad, 0x08, 0x20, 0x4a, 0xa8, 0x5f, 0xbd, 0xfa, 0xc5, 0xae, 0x1f,
	0x16, 0x1c, 0x4b, 0x06, 0x99, 0xf2, 0xdd, 0x6f, 0x0b, 0xfa, 0xfc, 0x56, 0xd5, 0x80, 0xb1, 0xa3,
	0xdf, 0xc6, 0x85, 0x8c, 0x81, 0xb3, 0x03, 0x1a, 0x22, 0xc9, 0xc8, 0x00, 0x89, 0x1c, 0x13, 0x96,
	0x45, 0xa6, 0xe6, 0x34, 0xdb, 0x73, 0x5d, 0x77, 0x54, 0xda, 0xd6, 0x10, 0xa7, 0xc9, 0x63, 0xf7,
	0x1f, 0xa0, 0x1b, 0x2c, 0x9f, 0x39, 0xaf, 0x19, 0x19, 0x6c, 0x2b, 0xdd, 0xf8, 0xa0, 0xe9, 0xad,
	0x5a, 0x8c, 0xc8, 0x39, 0xc5, 0x21, 0xda, 0xc5, 0x44, 0x02, 0x17, 0xe6, 0x15, 0xa7, 0xd9, 0x5e,
	0xea, 0x6e, 0x1c, 0x96, 0x76, 0xe3, 0x47, 0x69, 0xdf, 0x51, 0xbf, 0x25, 0xc2, 0x81, 0xc7, 0xc0,
	0x4f, 0xb1, 0x8c, 0xbd, 0x4d, 0x1a, 0x61, 0x32, 0xec, 0x51, 0x32, 0x2a, 0x6d, 0xe7, 0x5c, 0x07,
	0xb3, 0xd9, 0xdc, 0xa0, 0xf6, 0x1b, 0xdb, 0x95, 0xf5, 0x42, 0x39, 0xc6, 0x17, 0x4d, 0xb7, 0xfb,
	0x38, 0xc1, 0x19, 0xa1, 0x1c, 0x89, 0x18, 0x73, 0x2a, 0x10, 0xa7, 0xfb, 0x98, 0x87, 0x28, 0x64,
	0x82, 0x40, 0x91, 0x49, 0xb3, 0xe9, 0x68, 0xed, 0xa5, 0xee, 0xcb, 0xcb, 0xf5, 0x72, 0x5f, 0xf5,
	0xf2, 0x9f, 0x9c, 0x6e, 0x70, 0x77, 0x4a, 0x6c, 0x57, 0x40, 0x50, 0xf9, 0xbd, 0x89, 0xfd, 0xd7,
	0xe0, 0xf7, 0x0a, 0x90, 0x14, 0x85, 0x34, 0x83, 0x54, 0x98, 0x73, 0xd5, 0x64, 0x2e, 0x1e, 0x7c,
	0x1d, 0x9c, 0x19, 0xfc, 0xab, 0xb1, 0xd1, 0xab, 0x74, 0xe3, 0xa3, 0xa6, 0x1b, 0xb5, 0x98, 0x22,
	0x97, 0x2c, 0xa5, 0xc2, 0xbc, 0xea, 0x34, 0xdb, 0xd7, 0xd6, 0x5b, 0x9e, 0xda, 0x0e, 0x6f, 0xba,
	0x1d, 0x5e, 0x6f, 0xb2, 0x1d, 0xdd, 0x27, 0xe3, 0x01, 0xfc, 0x2e, 0x6d, 0x63, 0xba, 0x2f, 0x0f,
	0x21, 0x65, 0x92, 0xa6, 0xb9, 0x1c, 0x8e, 0x4a, 0xbb, 0x75, 0xae, 0x99, 0x49, 0x62, 0xf7, 0xeb,
	0x4f, 0x5b, 0x0b, 0x6e, 0x9e, 0x19, 0x6f, 0x94, 0x6e, 0x7c, 0xd2, 0xf4, 0x07, 0x4c, 0xa0, 0x9c,
	0xf2, 0x94, 0x09, 0xc1, 0x20, 0x4b, 0xa8, 0x10, 0x28, 0x07, 0x48, 0x10, 0xe1, 0xb4, 0xaa, 0x80,
	0x68, 0x86, 0xfb, 0x09, 0x0d, 0xcd, 0x79, 0x47, 0x6b, 0x2f, 0x76, 0xd7, 0x47, 0xa5, 0xed, 0xa9,
	0x3a, 0x97, 0x0c, 0x74, 0x83, 0x55, 0x26, 0xb6, 0x66, 0xc0, 0x2d, 0x80, 0xe4, 0xd9, 0x04, 0x7b,
	0xae, 0x28, 0xe3, 0xbd, 0xbe, 0x5a, 0x64, 0x9c, 0x0a, 0xc9, 0x19, 0x91, 0x34, 0xac, 0xe5, 0x02,
	0x8e, 0xf6, 0x63, 0x26, 0x69, 0xc2, 0x84, 0x34, 0x17, 0xaa, 0xd1, 0x7b, 0xa3, 0xd2, 0x5e, 0x53,
	0x5d, 0x5c, 0x22, 0xc8, 0x0d, 0x9c, 0x3a, 0x75, 0x5a, 0x1d, 0xf8, 0xdb, 0x29, 0x62, 0x3c, 0xd5,
	0x6f, 0xc4, 0x00, 0x03, 0x14, 0x61, 0x81, 0x12, 0x96, 0x32, 0x69, 0x2e, 0x3a, 0x5a, 0x7b, 0xae,
	0xdb, 0x1a, 0x95, 0xf6, 0xb2, 0xaa, 0x34, 0xeb, 0xbb, 0xc1, 0xf5, 0xb1, 0xb0, 0x81, 0xc5, 0xe6,
	0xf8, 0xd9, 0x1d, 0x1e, 0x1e, 0x5b, 0xda, 0xd1, 0xb1, 0xa5, 0xfd, 0x3a, 0xb6, 0xb4, 0xcf, 0x27,
	0x56, 0xe3, 0xe8, 0xc4, 0x6a, 0x7c, 0x3f, 0xb1, 0x1a, 0x3b, 0x28, 0x62, 0x32, 0x2e, 0xfa, 0x1e,
	0x81, 0xd4, 0xdf, 0x2b, 0x18, 0x19, 0x08, 0x96, 0xbc, 0xa3, 0xbc, 0x73, 0x00, 0x19, 0xad, 0x0b,
	0xbe, 0x8c, 0x19, 0x0f, 0x3b, 0x39, 0xe6, 0x72, 0xd8, 0x21, 0x31, 0x66, 0x99, 0xf0, 0x27, 0x77,
	0xa7, 0x23, 0x87, 0x39, 0x9d, 0xbd, 0x54, 0x9d, 0xd3, 0xf3, 0xd3, 0x9f, 0xaf, 0x16, 0xe5, 0xd1,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb6, 0x63, 0x63, 0xd2, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HookGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HookGasLimit))
		i--
		dAtA[i] = 0x40
	}
	if len(m.UnrestrictedPoolCreatorWhitelist) > 0 {
		for iNdEx := len(m.UnrestrictedPoolCreatorWhitelist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UnrestrictedPoolCreatorWhitelist[iNdEx])
			copy(dAtA[i:], m.UnrestrictedPoolCreatorWhitelist[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.UnrestrictedPoolCreatorWhitelist[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.IsPermissionlessPoolCreationEnabled {
		i--
		if m.IsPermissionlessPoolCreationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.AuthorizedUptimes) > 0 {
		for iNdEx := len(m.AuthorizedUptimes) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.AuthorizedUptimes[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.AuthorizedUptimes[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintParams(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.AuthorizedQuoteDenoms) > 0 {
		for iNdEx := len(m.AuthorizedQuoteDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedQuoteDenoms[iNdEx])
			copy(dAtA[i:], m.AuthorizedQuoteDenoms[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AuthorizedQuoteDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.BalancerSharesRewardDiscount.Size()
		i -= size
		if _, err := m.BalancerSharesRewardDiscount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.AuthorizedSpreadFactors) > 0 {
		for iNdEx := len(m.AuthorizedSpreadFactors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.AuthorizedSpreadFactors[iNdEx].Size()
				i -= size
				if _, err := m.AuthorizedSpreadFactors[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AuthorizedTickSpacing) > 0 {
		dAtA2 := make([]byte, len(m.AuthorizedTickSpacing)*10)
		var j1 int
		for _, num := range m.AuthorizedTickSpacing {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintParams(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthorizedTickSpacing) > 0 {
		l = 0
		for _, e := range m.AuthorizedTickSpacing {
			l += sovParams(uint64(e))
		}
		n += 1 + sovParams(uint64(l)) + l
	}
	if len(m.AuthorizedSpreadFactors) > 0 {
		for _, e := range m.AuthorizedSpreadFactors {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.BalancerSharesRewardDiscount.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.AuthorizedQuoteDenoms) > 0 {
		for _, s := range m.AuthorizedQuoteDenoms {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.AuthorizedUptimes) > 0 {
		for _, e := range m.AuthorizedUptimes {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.IsPermissionlessPoolCreationEnabled {
		n += 2
	}
	if len(m.UnrestrictedPoolCreatorWhitelist) > 0 {
		for _, s := range m.UnrestrictedPoolCreatorWhitelist {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.HookGasLimit != 0 {
		n += 1 + sovParams(uint64(m.HookGasLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.AuthorizedTickSpacing = append(m.AuthorizedTickSpacing, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthParams
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthParams
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.AuthorizedTickSpacing) == 0 {
					m.AuthorizedTickSpacing = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.AuthorizedTickSpacing = append(m.AuthorizedTickSpacing, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedTickSpacing", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedSpreadFactors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.AuthorizedSpreadFactors = append(m.AuthorizedSpreadFactors, v)
			if err := m.AuthorizedSpreadFactors[len(m.AuthorizedSpreadFactors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerSharesRewardDiscount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BalancerSharesRewardDiscount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedQuoteDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedQuoteDenoms = append(m.AuthorizedQuoteDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedUptimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedUptimes = append(m.AuthorizedUptimes, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.AuthorizedUptimes[len(m.AuthorizedUptimes)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPermissionlessPoolCreationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.IsPermissionlessPoolCreationEnabled = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnrestrictedPoolCreatorWhitelist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnrestrictedPoolCreatorWhitelist = append(m.UnrestrictedPoolCreatorWhitelist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HookGasLimit", wireType)
			}
			m.HookGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HookGasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)