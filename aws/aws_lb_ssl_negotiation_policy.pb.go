// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aws_lb_ssl_negotiation_policy.proto

package terraform_aws

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AwsLbSslNegotiationPolicy struct {
	Attribute    []*AwsLbSslNegotiationPolicy_Attribute `protobuf:"bytes,1,rep,name=attribute,proto3" json:"attribute,omitempty"`
	LbPort       int64                                  `protobuf:"varint,2,opt,name=lb_port,json=lbPort,proto3" json:"lb_port,omitempty"`
	LoadBalancer string                                 `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer,proto3" json:"load_balancer,omitempty"`
	Name         string                                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *AwsLbSslNegotiationPolicy) Reset()      { *m = AwsLbSslNegotiationPolicy{} }
func (*AwsLbSslNegotiationPolicy) ProtoMessage() {}
func (*AwsLbSslNegotiationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e2a9bc55ad20ed, []int{0}
}
func (m *AwsLbSslNegotiationPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AwsLbSslNegotiationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AwsLbSslNegotiationPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AwsLbSslNegotiationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsLbSslNegotiationPolicy.Merge(m, src)
}
func (m *AwsLbSslNegotiationPolicy) XXX_Size() int {
	return m.Size()
}
func (m *AwsLbSslNegotiationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsLbSslNegotiationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AwsLbSslNegotiationPolicy proto.InternalMessageInfo

func (m *AwsLbSslNegotiationPolicy) GetAttribute() []*AwsLbSslNegotiationPolicy_Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *AwsLbSslNegotiationPolicy) GetLbPort() int64 {
	if m != nil {
		return m.LbPort
	}
	return 0
}

func (m *AwsLbSslNegotiationPolicy) GetLoadBalancer() string {
	if m != nil {
		return m.LoadBalancer
	}
	return ""
}

func (m *AwsLbSslNegotiationPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AwsLbSslNegotiationPolicy_Attribute struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AwsLbSslNegotiationPolicy_Attribute) Reset()      { *m = AwsLbSslNegotiationPolicy_Attribute{} }
func (*AwsLbSslNegotiationPolicy_Attribute) ProtoMessage() {}
func (*AwsLbSslNegotiationPolicy_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e2a9bc55ad20ed, []int{0, 0}
}
func (m *AwsLbSslNegotiationPolicy_Attribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AwsLbSslNegotiationPolicy_Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AwsLbSslNegotiationPolicy_Attribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AwsLbSslNegotiationPolicy_Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsLbSslNegotiationPolicy_Attribute.Merge(m, src)
}
func (m *AwsLbSslNegotiationPolicy_Attribute) XXX_Size() int {
	return m.Size()
}
func (m *AwsLbSslNegotiationPolicy_Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsLbSslNegotiationPolicy_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_AwsLbSslNegotiationPolicy_Attribute proto.InternalMessageInfo

func (m *AwsLbSslNegotiationPolicy_Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AwsLbSslNegotiationPolicy_Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsLbSslNegotiationPolicy)(nil), "terraform.aws.AwsLbSslNegotiationPolicy")
	proto.RegisterType((*AwsLbSslNegotiationPolicy_Attribute)(nil), "terraform.aws.AwsLbSslNegotiationPolicy.Attribute")
}

func init() {
	proto.RegisterFile("aws_lb_ssl_negotiation_policy.proto", fileDescriptor_47e2a9bc55ad20ed)
}

var fileDescriptor_47e2a9bc55ad20ed = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7d, 0xff, 0xf6, 0x2f, 0x8a, 0xa1, 0x8b, 0x85, 0x44, 0x60, 0xb8, 0xaa, 0xe8, 0xd2,
	0x29, 0x43, 0x81, 0x07, 0x68, 0x67, 0x84, 0xaa, 0xf0, 0x00, 0x96, 0x5d, 0x0c, 0x8a, 0xe4, 0xc6,
	0x91, 0xed, 0x12, 0xb1, 0xf1, 0x08, 0x3c, 0x06, 0x8f, 0xc2, 0x98, 0xb1, 0x23, 0x71, 0x16, 0xc6,
	0x4e, 0xcc, 0x88, 0x00, 0x8d, 0x18, 0xd8, 0xee, 0x3d, 0xf7, 0x7c, 0xe7, 0x48, 0x97, 0x8e, 0x45,
	0xe9, 0xb8, 0x96, 0xdc, 0x39, 0xcd, 0x73, 0x75, 0x67, 0x7c, 0x26, 0x7c, 0x66, 0x72, 0x5e, 0x18,
	0x9d, 0x2d, 0x1f, 0x92, 0xc2, 0x1a, 0x6f, 0xd8, 0xd0, 0x2b, 0x6b, 0xc5, 0xad, 0xb1, 0xab, 0x44,
	0x94, 0xee, 0xf4, 0x1d, 0xe8, 0xf1, 0xac, 0x74, 0x97, 0xf2, 0xda, 0xe9, 0xab, 0x8e, 0x59, 0xb4,
	0x08, 0x5b, 0xd0, 0x48, 0x78, 0x6f, 0x33, 0xb9, 0xf6, 0x2a, 0x86, 0x51, 0x6f, 0xb2, 0x3f, 0x9d,
	0x26, 0xbf, 0x02, 0x92, 0x3f, 0xe1, 0x64, 0xf6, 0x43, 0xa6, 0x5d, 0x08, 0x3b, 0xa2, 0x7b, 0x5a,
	0xf2, 0xc2, 0x58, 0x1f, 0xff, 0x1b, 0xc1, 0xa4, 0x97, 0x0e, 0xb4, 0x5c, 0x18, 0xeb, 0xd9, 0x98,
	0x0e, 0xb5, 0x11, 0x37, 0x5c, 0x0a, 0x2d, 0xf2, 0xa5, 0xb2, 0x71, 0x6f, 0x04, 0x93, 0x28, 0x3d,
	0xf8, 0x14, 0xe7, 0xdf, 0x1a, 0x63, 0xb4, 0x9f, 0x8b, 0x95, 0x8a, 0xfb, 0xed, 0xad, 0x9d, 0x4f,
	0x2e, 0x68, 0xb4, 0x6b, 0xda, 0x19, 0xa0, 0x33, 0xb0, 0x43, 0xfa, 0xff, 0x5e, 0xe8, 0xb5, 0x6a,
	0x0b, 0xa3, 0xf4, 0x6b, 0x99, 0x9f, 0x57, 0x35, 0x92, 0x4d, 0x8d, 0x64, 0x5b, 0x23, 0x3c, 0x06,
	0x84, 0xe7, 0x80, 0xf0, 0x12, 0x10, 0xaa, 0x80, 0xf0, 0x1a, 0x10, 0xde, 0x02, 0x92, 0x6d, 0x40,
	0x78, 0x6a, 0x90, 0x54, 0x0d, 0x92, 0x4d, 0x83, 0x44, 0x0e, 0xda, 0x27, 0x9e, 0x7d, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0x2d, 0x4c, 0xf7, 0x6b, 0x01, 0x00, 0x00,
}

func (this *AwsLbSslNegotiationPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsLbSslNegotiationPolicy)
	if !ok {
		that2, ok := that.(AwsLbSslNegotiationPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Attribute) != len(that1.Attribute) {
		return false
	}
	for i := range this.Attribute {
		if !this.Attribute[i].Equal(that1.Attribute[i]) {
			return false
		}
	}
	if this.LbPort != that1.LbPort {
		return false
	}
	if this.LoadBalancer != that1.LoadBalancer {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *AwsLbSslNegotiationPolicy_Attribute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsLbSslNegotiationPolicy_Attribute)
	if !ok {
		that2, ok := that.(AwsLbSslNegotiationPolicy_Attribute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *AwsLbSslNegotiationPolicy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&terraform_aws.AwsLbSslNegotiationPolicy{")
	if this.Attribute != nil {
		s = append(s, "Attribute: "+fmt.Sprintf("%#v", this.Attribute)+",\n")
	}
	s = append(s, "LbPort: "+fmt.Sprintf("%#v", this.LbPort)+",\n")
	s = append(s, "LoadBalancer: "+fmt.Sprintf("%#v", this.LoadBalancer)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AwsLbSslNegotiationPolicy_Attribute) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&terraform_aws.AwsLbSslNegotiationPolicy_Attribute{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAwsLbSslNegotiationPolicy(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AwsLbSslNegotiationPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsLbSslNegotiationPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Attribute) > 0 {
		for _, msg := range m.Attribute {
			dAtA[i] = 0xa
			i++
			i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LbPort != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(m.LbPort))
	}
	if len(m.LoadBalancer) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(len(m.LoadBalancer)))
		i += copy(dAtA[i:], m.LoadBalancer)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *AwsLbSslNegotiationPolicy_Attribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsLbSslNegotiationPolicy_Attribute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAwsLbSslNegotiationPolicy(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeVarintAwsLbSslNegotiationPolicy(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AwsLbSslNegotiationPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Attribute) > 0 {
		for _, e := range m.Attribute {
			l = e.Size()
			n += 1 + l + sovAwsLbSslNegotiationPolicy(uint64(l))
		}
	}
	if m.LbPort != 0 {
		n += 1 + sovAwsLbSslNegotiationPolicy(uint64(m.LbPort))
	}
	l = len(m.LoadBalancer)
	if l > 0 {
		n += 1 + l + sovAwsLbSslNegotiationPolicy(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAwsLbSslNegotiationPolicy(uint64(l))
	}
	return n
}

func (m *AwsLbSslNegotiationPolicy_Attribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAwsLbSslNegotiationPolicy(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovAwsLbSslNegotiationPolicy(uint64(l))
	}
	return n
}

func sovAwsLbSslNegotiationPolicy(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAwsLbSslNegotiationPolicy(x uint64) (n int) {
	return sovAwsLbSslNegotiationPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AwsLbSslNegotiationPolicy) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AwsLbSslNegotiationPolicy{`,
		`Attribute:` + strings.Replace(fmt.Sprintf("%v", this.Attribute), "AwsLbSslNegotiationPolicy_Attribute", "AwsLbSslNegotiationPolicy_Attribute", 1) + `,`,
		`LbPort:` + fmt.Sprintf("%v", this.LbPort) + `,`,
		`LoadBalancer:` + fmt.Sprintf("%v", this.LoadBalancer) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AwsLbSslNegotiationPolicy_Attribute) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AwsLbSslNegotiationPolicy_Attribute{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAwsLbSslNegotiationPolicy(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AwsLbSslNegotiationPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAwsLbSslNegotiationPolicy
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
			return fmt.Errorf("proto: AwsLbSslNegotiationPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AwsLbSslNegotiationPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attribute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attribute = append(m.Attribute, &AwsLbSslNegotiationPolicy_Attribute{})
			if err := m.Attribute[len(m.Attribute)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LbPort", wireType)
			}
			m.LbPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LbPort |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadBalancer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LoadBalancer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAwsLbSslNegotiationPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
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
func (m *AwsLbSslNegotiationPolicy_Attribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAwsLbSslNegotiationPolicy
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
			return fmt.Errorf("proto: Attribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAwsLbSslNegotiationPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAwsLbSslNegotiationPolicy
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
func skipAwsLbSslNegotiationPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAwsLbSslNegotiationPolicy
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
					return 0, ErrIntOverflowAwsLbSslNegotiationPolicy
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
					return 0, ErrIntOverflowAwsLbSslNegotiationPolicy
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
				return 0, ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAwsLbSslNegotiationPolicy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAwsLbSslNegotiationPolicy
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
				next, err := skipAwsLbSslNegotiationPolicy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAwsLbSslNegotiationPolicy
				}
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
	ErrInvalidLengthAwsLbSslNegotiationPolicy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAwsLbSslNegotiationPolicy   = fmt.Errorf("proto: integer overflow")
)