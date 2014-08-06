// Code generated by protoc-gen-go.
// source: lifx.proto
// DO NOT EDIT!


/*
Package lifx is a generated protocol buffer package.

It is generated from these files:
	lifx.proto

It has these top-level messages:
	Packet
	Payload
	GetPanGateway
	PanGateway
	GetBulbLabel
	BulbLabel
	GetLightState
	LightStatus
	SetLightColor
*/
package lifx

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Payload_Type int32

const (
	Payload_getPanGateway Payload_Type = 1
	Payload_panGateway    Payload_Type = 2
	Payload_getBulbLabel  Payload_Type = 3
	Payload_bulbLabel     Payload_Type = 4
	Payload_getLightState Payload_Type = 5
	Payload_lightStatus   Payload_Type = 6
	Payload_setLightColor Payload_Type = 7
)

var Payload_Type_name = map[int32]string{
	1: "getPanGateway",
	2: "panGateway",
	3: "getBulbLabel",
	4: "bulbLabel",
	5: "getLightState",
	6: "lightStatus",
	7: "setLightColor",
}
var Payload_Type_value = map[string]int32{
	"getPanGateway": 1,
	"panGateway":    2,
	"getBulbLabel":  3,
	"bulbLabel":     4,
	"getLightState": 5,
	"lightStatus":   6,
	"setLightColor": 7,
}

func (x Payload_Type) Enum() *Payload_Type {
	p := new(Payload_Type)
	*p = x
	return p
}
func (x Payload_Type) String() string {
	return proto.EnumName(Payload_Type_name, int32(x))
}
func (x *Payload_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Payload_Type_value, data, "Payload_Type")
	if err != nil {
		return err
	}
	*x = Payload_Type(value)
	return nil
}

type Packet struct {
	Size             []byte   `protobuf:"bytes,1,req,name=size" json:"size,omitempty"`
	Protocol         []byte   `protobuf:"bytes,2,req,name=protocol" json:"protocol,omitempty"`
	Reserved1        *uint32  `protobuf:"varint,3,req,name=reserved1" json:"reserved1,omitempty"`
	TargetMacAddress []byte   `protobuf:"bytes,4,req,name=target_mac_address" json:"target_mac_address,omitempty"`
	Reserved2        []byte   `protobuf:"bytes,5,req,name=reserved2" json:"reserved2,omitempty"`
	Site             []byte   `protobuf:"bytes,6,req,name=site" json:"site,omitempty"`
	Reserved3        []byte   `protobuf:"bytes,7,req,name=reserved3" json:"reserved3,omitempty"`
	Timestamp        *uint64  `protobuf:"varint,8,req,name=timestamp" json:"timestamp,omitempty"`
	PacketType       []byte   `protobuf:"bytes,9,req,name=packet_type" json:"packet_type,omitempty"`
	Reserved4        []byte   `protobuf:"bytes,10,req,name=reserved4" json:"reserved4,omitempty"`
	Payload          *Payload `protobuf:"bytes,11,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}

func (m *Packet) GetSize() []byte {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *Packet) GetProtocol() []byte {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *Packet) GetReserved1() uint32 {
	if m != nil && m.Reserved1 != nil {
		return *m.Reserved1
	}
	return 0
}

func (m *Packet) GetTargetMacAddress() []byte {
	if m != nil {
		return m.TargetMacAddress
	}
	return nil
}

func (m *Packet) GetReserved2() []byte {
	if m != nil {
		return m.Reserved2
	}
	return nil
}

func (m *Packet) GetSite() []byte {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Packet) GetReserved3() []byte {
	if m != nil {
		return m.Reserved3
	}
	return nil
}

func (m *Packet) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Packet) GetPacketType() []byte {
	if m != nil {
		return m.PacketType
	}
	return nil
}

func (m *Packet) GetReserved4() []byte {
	if m != nil {
		return m.Reserved4
	}
	return nil
}

func (m *Packet) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Payload struct {
	Type             *Payload_Type             `protobuf:"varint,1,req,name=type,enum=lifx.Payload_Type" json:"type,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}

var extRange_Payload = []proto.ExtensionRange{
	{100, 536870911},
}

func (*Payload) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Payload
}
func (m *Payload) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Payload) GetType() Payload_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Payload_getPanGateway
}

type GetPanGateway struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPanGateway) Reset()         { *m = GetPanGateway{} }
func (m *GetPanGateway) String() string { return proto.CompactTextString(m) }
func (*GetPanGateway) ProtoMessage()    {}

var E_GetPanGateway_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*GetPanGateway)(nil),
	Field:         100,
	Name:          "lifx.getPanGateway.payload",
	Tag:           "bytes,100,req,name=payload",
}

type PanGateway struct {
	Service          []byte  `protobuf:"bytes,1,req,name=service" json:"service,omitempty"`
	Port             *uint32 `protobuf:"varint,2,req,name=port" json:"port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PanGateway) Reset()         { *m = PanGateway{} }
func (m *PanGateway) String() string { return proto.CompactTextString(m) }
func (*PanGateway) ProtoMessage()    {}

func (m *PanGateway) GetService() []byte {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *PanGateway) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

var E_PanGateway_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*PanGateway)(nil),
	Field:         101,
	Name:          "lifx.panGateway.payload",
	Tag:           "bytes,101,req,name=payload",
}

type GetBulbLabel struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetBulbLabel) Reset()         { *m = GetBulbLabel{} }
func (m *GetBulbLabel) String() string { return proto.CompactTextString(m) }
func (*GetBulbLabel) ProtoMessage()    {}

var E_GetBulbLabel_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*GetBulbLabel)(nil),
	Field:         102,
	Name:          "lifx.getBulbLabel.payload",
	Tag:           "bytes,102,req,name=payload",
}

type BulbLabel struct {
	Label            *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BulbLabel) Reset()         { *m = BulbLabel{} }
func (m *BulbLabel) String() string { return proto.CompactTextString(m) }
func (*BulbLabel) ProtoMessage()    {}

func (m *BulbLabel) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

var E_BulbLabel_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*BulbLabel)(nil),
	Field:         103,
	Name:          "lifx.bulbLabel.payload",
	Tag:           "bytes,103,req,name=payload",
}

type GetLightState struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetLightState) Reset()         { *m = GetLightState{} }
func (m *GetLightState) String() string { return proto.CompactTextString(m) }
func (*GetLightState) ProtoMessage()    {}

var E_GetLightState_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*GetLightState)(nil),
	Field:         104,
	Name:          "lifx.getLightState.payload",
	Tag:           "bytes,104,req,name=payload",
}

type LightStatus struct {
	Hue              []byte  `protobuf:"bytes,1,req,name=hue" json:"hue,omitempty"`
	Saturation       []byte  `protobuf:"bytes,2,req,name=saturation" json:"saturation,omitempty"`
	Brightness       []byte  `protobuf:"bytes,3,req,name=brightness" json:"brightness,omitempty"`
	Kelvin           []byte  `protobuf:"bytes,4,req,name=kelvin" json:"kelvin,omitempty"`
	Dim              []byte  `protobuf:"bytes,5,req,name=dim" json:"dim,omitempty"`
	Power            []byte  `protobuf:"bytes,6,req,name=power" json:"power,omitempty"`
	BulbLabel        *string `protobuf:"bytes,7,req,name=bulb_label" json:"bulb_label,omitempty"`
	Tags             *uint64 `protobuf:"varint,8,req,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LightStatus) Reset()         { *m = LightStatus{} }
func (m *LightStatus) String() string { return proto.CompactTextString(m) }
func (*LightStatus) ProtoMessage()    {}

func (m *LightStatus) GetHue() []byte {
	if m != nil {
		return m.Hue
	}
	return nil
}

func (m *LightStatus) GetSaturation() []byte {
	if m != nil {
		return m.Saturation
	}
	return nil
}

func (m *LightStatus) GetBrightness() []byte {
	if m != nil {
		return m.Brightness
	}
	return nil
}

func (m *LightStatus) GetKelvin() []byte {
	if m != nil {
		return m.Kelvin
	}
	return nil
}

func (m *LightStatus) GetDim() []byte {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *LightStatus) GetPower() []byte {
	if m != nil {
		return m.Power
	}
	return nil
}

func (m *LightStatus) GetBulbLabel() string {
	if m != nil && m.BulbLabel != nil {
		return *m.BulbLabel
	}
	return ""
}

func (m *LightStatus) GetTags() uint64 {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return 0
}

var E_LightStatus_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*LightStatus)(nil),
	Field:         105,
	Name:          "lifx.lightStatus.payload",
	Tag:           "bytes,105,req,name=payload",
}

type SetLightColor struct {
	Stream           []byte  `protobuf:"bytes,1,req,name=stream" json:"stream,omitempty"`
	Hue              []byte  `protobuf:"bytes,2,req,name=hue" json:"hue,omitempty"`
	Saturation       []byte  `protobuf:"bytes,3,req,name=saturation" json:"saturation,omitempty"`
	Brightness       []byte  `protobuf:"bytes,4,req,name=brightness" json:"brightness,omitempty"`
	Kelvin           []byte  `protobuf:"bytes,5,req,name=kelvin" json:"kelvin,omitempty"`
	FadeTime         *uint32 `protobuf:"varint,6,req,name=fade_time" json:"fade_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetLightColor) Reset()         { *m = SetLightColor{} }
func (m *SetLightColor) String() string { return proto.CompactTextString(m) }
func (*SetLightColor) ProtoMessage()    {}

func (m *SetLightColor) GetStream() []byte {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *SetLightColor) GetHue() []byte {
	if m != nil {
		return m.Hue
	}
	return nil
}

func (m *SetLightColor) GetSaturation() []byte {
	if m != nil {
		return m.Saturation
	}
	return nil
}

func (m *SetLightColor) GetBrightness() []byte {
	if m != nil {
		return m.Brightness
	}
	return nil
}

func (m *SetLightColor) GetKelvin() []byte {
	if m != nil {
		return m.Kelvin
	}
	return nil
}

func (m *SetLightColor) GetFadeTime() uint32 {
	if m != nil && m.FadeTime != nil {
		return *m.FadeTime
	}
	return 0
}

var E_SetLightColor_Payload = &proto.ExtensionDesc{
	ExtendedType:  (*Payload)(nil),
	ExtensionType: (*SetLightColor)(nil),
	Field:         106,
	Name:          "lifx.setLightColor.payload",
	Tag:           "bytes,106,req,name=payload",
}

func init() {
	proto.RegisterEnum("lifx.Payload_Type", Payload_Type_name, Payload_Type_value)
	proto.RegisterExtension(E_GetPanGateway_Payload)
	proto.RegisterExtension(E_PanGateway_Payload)
	proto.RegisterExtension(E_GetBulbLabel_Payload)
	proto.RegisterExtension(E_BulbLabel_Payload)
	proto.RegisterExtension(E_GetLightState_Payload)
	proto.RegisterExtension(E_LightStatus_Payload)
	proto.RegisterExtension(E_SetLightColor_Payload)
}
