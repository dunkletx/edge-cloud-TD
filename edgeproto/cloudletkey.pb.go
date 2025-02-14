// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudletkey.proto

package edgeproto

import (
	"encoding/json"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud/log"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// CloudletKey uniquely identifies a Cloudlet.
type CloudletKey struct {
	// Organization of the cloudlet site
	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	// Name of the cloudlet
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Federated operator organization who shared this cloudlet
	FederatedOrganization string `protobuf:"bytes,3,opt,name=federated_organization,json=federatedOrganization,proto3" json:"federated_organization,omitempty"`
}

func (m *CloudletKey) Reset()         { *m = CloudletKey{} }
func (m *CloudletKey) String() string { return proto.CompactTextString(m) }
func (*CloudletKey) ProtoMessage()    {}
func (*CloudletKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a22a8ce50ccfe8b, []int{0}
}
func (m *CloudletKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloudletKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CloudletKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CloudletKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudletKey.Merge(m, src)
}
func (m *CloudletKey) XXX_Size() int {
	return m.Size()
}
func (m *CloudletKey) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudletKey.DiscardUnknown(m)
}

var xxx_messageInfo_CloudletKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CloudletKey)(nil), "edgeproto.CloudletKey")
}

func init() { proto.RegisterFile("cloudletkey.proto", fileDescriptor_7a22a8ce50ccfe8b) }

var fileDescriptor_7a22a8ce50ccfe8b = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xc9, 0x2f,
	0x4d, 0xc9, 0x49, 0x2d, 0xc9, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c,
	0x4d, 0x49, 0x4f, 0x05, 0x33, 0xa5, 0x2c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x73, 0xf3, 0x93, 0x32, 0x73, 0x40, 0x52, 0x15, 0xfa, 0x20, 0x52, 0x17, 0xac, 0x51,
	0x1f, 0xac, 0x2e, 0x3d, 0x35, 0x0f, 0xce, 0x80, 0x18, 0x22, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x66, 0xea, 0x83, 0x58, 0x10, 0x51, 0xa5, 0x9d, 0x4c, 0x5c, 0xdc, 0xce, 0x50, 0x0b, 0xbd, 0x53,
	0x2b, 0x85, 0x8c, 0xb9, 0x78, 0xf2, 0x8b, 0xd2, 0x13, 0xf3, 0x32, 0xab, 0x12, 0x4b, 0x32, 0xf3,
	0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xf8, 0x77, 0x7d, 0x93, 0xe0, 0x86, 0xb9, 0x2b,
	0xbf, 0x28, 0x3d, 0x08, 0x45, 0x91, 0x90, 0x02, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13,
	0x58, 0x31, 0xcf, 0xae, 0x6f, 0x12, 0x1c, 0x30, 0xc5, 0x41, 0x60, 0x19, 0x21, 0x77, 0x2e, 0xb1,
	0xb4, 0xd4, 0x94, 0xd4, 0xa2, 0xc4, 0x92, 0xd4, 0x94, 0x78, 0x14, 0x0b, 0x98, 0xc1, 0x7a, 0x04,
	0x76, 0x7d, 0x93, 0xe0, 0x81, 0xab, 0x00, 0xd9, 0x20, 0x0a, 0xe7, 0xf9, 0x23, 0x29, 0xb7, 0xea,
	0x64, 0x7c, 0xf1, 0x59, 0x82, 0xf1, 0xc7, 0x67, 0x09, 0xc6, 0x05, 0x5f, 0x24, 0x18, 0x67, 0x7d,
	0x95, 0xf0, 0x84, 0xd9, 0x64, 0xeb, 0x97, 0x98, 0x9b, 0xaa, 0x83, 0xe4, 0x48, 0x5b, 0x64, 0x8d,
	0x3a, 0x50, 0xe3, 0xf2, 0x8b, 0x40, 0x32, 0x6e, 0xd8, 0xcc, 0x3e, 0xf5, 0x55, 0x42, 0x10, 0x6c,
	0x06, 0xb2, 0xe0, 0xa7, 0xaf, 0x12, 0xc8, 0xe1, 0xb3, 0xe9, 0x9b, 0x04, 0x4b, 0x5e, 0x7e, 0x5e,
	0xea, 0x86, 0x05, 0xf2, 0x8c, 0x4e, 0x32, 0x27, 0x1e, 0xca, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0x38, 0x80, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0x20, 0x9e, 0x32, 0xd0, 0x01, 0x00, 0x00,
}

func (this *CloudletKey) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&edgeproto.CloudletKey{")
	s = append(s, "Organization: "+fmt.Sprintf("%#v", this.Organization)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "FederatedOrganization: "+fmt.Sprintf("%#v", this.FederatedOrganization)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCloudletkey(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CloudletKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudletKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudletKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FederatedOrganization) > 0 {
		i -= len(m.FederatedOrganization)
		copy(dAtA[i:], m.FederatedOrganization)
		i = encodeVarintCloudletkey(dAtA, i, uint64(len(m.FederatedOrganization)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCloudletkey(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Organization) > 0 {
		i -= len(m.Organization)
		copy(dAtA[i:], m.Organization)
		i = encodeVarintCloudletkey(dAtA, i, uint64(len(m.Organization)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCloudletkey(dAtA []byte, offset int, v uint64) int {
	offset -= sovCloudletkey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CloudletKey) Matches(o *CloudletKey, fopts ...MatchOpt) bool {
	opts := MatchOptions{}
	applyMatchOptions(&opts, fopts...)
	if o == nil {
		if opts.Filter {
			return true
		}
		return false
	}
	if !opts.Filter || o.Organization != "" {
		if o.Organization != m.Organization {
			return false
		}
	}
	if !opts.Filter || o.Name != "" {
		if o.Name != m.Name {
			return false
		}
	}
	if !opts.Filter || o.FederatedOrganization != "" {
		if o.FederatedOrganization != m.FederatedOrganization {
			return false
		}
	}
	return true
}

func (m *CloudletKey) CopyInFields(src *CloudletKey) int {
	changed := 0
	if m.Organization != src.Organization {
		m.Organization = src.Organization
		changed++
	}
	if m.Name != src.Name {
		m.Name = src.Name
		changed++
	}
	if m.FederatedOrganization != src.FederatedOrganization {
		m.FederatedOrganization = src.FederatedOrganization
		changed++
	}
	return changed
}

func (m *CloudletKey) DeepCopyIn(src *CloudletKey) {
	m.Organization = src.Organization
	m.Name = src.Name
	m.FederatedOrganization = src.FederatedOrganization
}

func (m *CloudletKey) GetKeyString() string {
	key, err := json.Marshal(m)
	if err != nil {
		log.FatalLog("Failed to marshal CloudletKey key string", "obj", m)
	}
	return string(key)
}

func CloudletKeyStringParse(str string, key *CloudletKey) {
	err := json.Unmarshal([]byte(str), key)
	if err != nil {
		log.FatalLog("Failed to unmarshal CloudletKey key string", "str", str)
	}
}

func (m *CloudletKey) NotFoundError() error {
	return fmt.Errorf("Cloudlet key %s not found", m.GetKeyString())
}

func (m *CloudletKey) ExistsError() error {
	return fmt.Errorf("Cloudlet key %s already exists", m.GetKeyString())
}

func (m *CloudletKey) BeingDeletedError() error {
	return fmt.Errorf("Cloudlet %s is being deleted", m.GetKeyString())
}

var CloudletKeyTagOrganization = "cloudletorg"
var CloudletKeyTagName = "cloudlet"
var CloudletKeyTagFederatedOrganization = "federatedorg"

func (m *CloudletKey) GetTags() map[string]string {
	tags := make(map[string]string)
	tags["cloudletorg"] = m.Organization
	tags["cloudlet"] = m.Name
	tags["federatedorg"] = m.FederatedOrganization
	return tags
}

// Helper method to check that enums have valid values
func (m *CloudletKey) ValidateEnums() error {
	return nil
}

func (s *CloudletKey) ClearTagged(tags map[string]struct{}) {
}

func (m *CloudletKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovCloudletkey(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCloudletkey(uint64(l))
	}
	l = len(m.FederatedOrganization)
	if l > 0 {
		n += 1 + l + sovCloudletkey(uint64(l))
	}
	return n
}

func sovCloudletkey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCloudletkey(x uint64) (n int) {
	return sovCloudletkey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CloudletKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloudletkey
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
			return fmt.Errorf("proto: CloudletKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudletKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudletkey
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
				return ErrInvalidLengthCloudletkey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudletkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudletkey
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
				return ErrInvalidLengthCloudletkey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudletkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FederatedOrganization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloudletkey
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
				return ErrInvalidLengthCloudletkey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloudletkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FederatedOrganization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloudletkey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCloudletkey
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCloudletkey
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
func skipCloudletkey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCloudletkey
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
					return 0, ErrIntOverflowCloudletkey
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
					return 0, ErrIntOverflowCloudletkey
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
				return 0, ErrInvalidLengthCloudletkey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCloudletkey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCloudletkey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCloudletkey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCloudletkey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCloudletkey = fmt.Errorf("proto: unexpected end of group")
)
