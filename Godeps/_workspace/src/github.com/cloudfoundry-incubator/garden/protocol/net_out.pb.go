// Code generated by protoc-gen-gogo.
// source: net_out.proto
// DO NOT EDIT!

package garden

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type NetOutRequest_Protocol int32

const (
	NetOutRequest_TCP  NetOutRequest_Protocol = 0
	NetOutRequest_ALL  NetOutRequest_Protocol = 1
	NetOutRequest_ICMP NetOutRequest_Protocol = 2
	NetOutRequest_UDP  NetOutRequest_Protocol = 3
)

var NetOutRequest_Protocol_name = map[int32]string{
	0: "TCP",
	1: "ALL",
	2: "ICMP",
	3: "UDP",
}
var NetOutRequest_Protocol_value = map[string]int32{
	"TCP":  0,
	"ALL":  1,
	"ICMP": 2,
	"UDP":  3,
}

func (x NetOutRequest_Protocol) Enum() *NetOutRequest_Protocol {
	p := new(NetOutRequest_Protocol)
	*p = x
	return p
}
func (x NetOutRequest_Protocol) String() string {
	return proto.EnumName(NetOutRequest_Protocol_name, int32(x))
}
func (x *NetOutRequest_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NetOutRequest_Protocol_value, data, "NetOutRequest_Protocol")
	if err != nil {
		return err
	}
	*x = NetOutRequest_Protocol(value)
	return nil
}

type NetOutRequest struct {
	Handle           *string                 `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	Network          *string                 `protobuf:"bytes,2,opt,name=network" json:"network,omitempty"`
	Port             *uint32                 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	PortRange        *string                 `protobuf:"bytes,4,opt,name=port_range" json:"port_range,omitempty"`
	Protocol         *NetOutRequest_Protocol `protobuf:"varint,5,opt,name=protocol,enum=garden.NetOutRequest_Protocol" json:"protocol,omitempty"`
	IcmpType         *int32                  `protobuf:"varint,6,opt,name=icmp_type,def=-1" json:"icmp_type,omitempty"`
	IcmpCode         *int32                  `protobuf:"varint,7,opt,name=icmp_code,def=-1" json:"icmp_code,omitempty"`
	Log              *bool                   `protobuf:"varint,8,opt,name=log" json:"log,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *NetOutRequest) Reset()         { *m = NetOutRequest{} }
func (m *NetOutRequest) String() string { return proto.CompactTextString(m) }
func (*NetOutRequest) ProtoMessage()    {}

const Default_NetOutRequest_IcmpType int32 = -1
const Default_NetOutRequest_IcmpCode int32 = -1

func (m *NetOutRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func (m *NetOutRequest) GetNetwork() string {
	if m != nil && m.Network != nil {
		return *m.Network
	}
	return ""
}

func (m *NetOutRequest) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *NetOutRequest) GetPortRange() string {
	if m != nil && m.PortRange != nil {
		return *m.PortRange
	}
	return ""
}

func (m *NetOutRequest) GetProtocol() NetOutRequest_Protocol {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return NetOutRequest_TCP
}

func (m *NetOutRequest) GetIcmpType() int32 {
	if m != nil && m.IcmpType != nil {
		return *m.IcmpType
	}
	return Default_NetOutRequest_IcmpType
}

func (m *NetOutRequest) GetIcmpCode() int32 {
	if m != nil && m.IcmpCode != nil {
		return *m.IcmpCode
	}
	return Default_NetOutRequest_IcmpCode
}

func (m *NetOutRequest) GetLog() bool {
	if m != nil && m.Log != nil {
		return *m.Log
	}
	return false
}

type NetOutResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NetOutResponse) Reset()         { *m = NetOutResponse{} }
func (m *NetOutResponse) String() string { return proto.CompactTextString(m) }
func (*NetOutResponse) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("garden.NetOutRequest_Protocol", NetOutRequest_Protocol_name, NetOutRequest_Protocol_value)
}
