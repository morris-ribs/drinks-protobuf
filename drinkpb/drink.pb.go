// Code generated by protoc-gen-go.
// source: drink.proto
// DO NOT EDIT!

/*
Package drinkpb is a generated protocol buffer package.

It is generated from these files:
	drink.proto

It has these top-level messages:
	Drink
	Drinks
*/
package drinkpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Drink_DrinkType int32

const (
	Drink_BEER   Drink_DrinkType = 0
	Drink_WINE   Drink_DrinkType = 1
	Drink_SPIRIT Drink_DrinkType = 2
	Drink_OTHER  Drink_DrinkType = 3
)

var Drink_DrinkType_name = map[int32]string{
	0: "BEER",
	1: "WINE",
	2: "SPIRIT",
	3: "OTHER",
}
var Drink_DrinkType_value = map[string]int32{
	"BEER":   0,
	"WINE":   1,
	"SPIRIT": 2,
	"OTHER":  3,
}

func (x Drink_DrinkType) String() string {
	return proto.EnumName(Drink_DrinkType_name, int32(x))
}
func (Drink_DrinkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Drink_DrinkSubType int32

const (
	Drink_PALEALE Drink_DrinkSubType = 0
	Drink_DARK    Drink_DrinkSubType = 1
	Drink_RED     Drink_DrinkSubType = 2
	Drink_WHITE   Drink_DrinkSubType = 3
	Drink_ROSE    Drink_DrinkSubType = 4
	Drink_RHUM    Drink_DrinkSubType = 5
	Drink_VODKA   Drink_DrinkSubType = 6
	Drink_LIQUEUR Drink_DrinkSubType = 7
	Drink_JUICE   Drink_DrinkSubType = 8
	Drink_SODA    Drink_DrinkSubType = 9
)

var Drink_DrinkSubType_name = map[int32]string{
	0: "PALEALE",
	1: "DARK",
	2: "RED",
	3: "WHITE",
	4: "ROSE",
	5: "RHUM",
	6: "VODKA",
	7: "LIQUEUR",
	8: "JUICE",
	9: "SODA",
}
var Drink_DrinkSubType_value = map[string]int32{
	"PALEALE": 0,
	"DARK":    1,
	"RED":     2,
	"WHITE":   3,
	"ROSE":    4,
	"RHUM":    5,
	"VODKA":   6,
	"LIQUEUR": 7,
	"JUICE":   8,
	"SODA":    9,
}

func (x Drink_DrinkSubType) String() string {
	return proto.EnumName(Drink_DrinkSubType_name, int32(x))
}
func (Drink_DrinkSubType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Drink struct {
	Name  string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type1 Drink_DrinkType    `protobuf:"varint,2,opt,name=type1,enum=drinkpb.Drink_DrinkType" json:"type1,omitempty"`
	Type2 Drink_DrinkSubType `protobuf:"varint,3,opt,name=type2,enum=drinkpb.Drink_DrinkSubType" json:"type2,omitempty"`
	Price float32            `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty"`
	Stock int32              `protobuf:"varint,5,opt,name=stock" json:"stock,omitempty"`
}

func (m *Drink) Reset()                    { *m = Drink{} }
func (m *Drink) String() string            { return proto.CompactTextString(m) }
func (*Drink) ProtoMessage()               {}
func (*Drink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Drinks struct {
	Drinks []*Drink `protobuf:"bytes,1,rep,name=drinks" json:"drinks,omitempty"`
}

func (m *Drinks) Reset()                    { *m = Drinks{} }
func (m *Drinks) String() string            { return proto.CompactTextString(m) }
func (*Drinks) ProtoMessage()               {}
func (*Drinks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Drinks) GetDrinks() []*Drink {
	if m != nil {
		return m.Drinks
	}
	return nil
}

func init() {
	proto.RegisterType((*Drink)(nil), "drinkpb.Drink")
	proto.RegisterType((*Drinks)(nil), "drinkpb.Drinks")
	proto.RegisterEnum("drinkpb.Drink_DrinkType", Drink_DrinkType_name, Drink_DrinkType_value)
	proto.RegisterEnum("drinkpb.Drink_DrinkSubType", Drink_DrinkSubType_name, Drink_DrinkSubType_value)
}

func init() { proto.RegisterFile("drink.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0xfc, 0x6d, 0x92, 0x4d, 0x9a, 0xd7, 0x1f, 0x65, 0x59, 0x3c, 0x2c, 0x78, 0x09, 0x39, 0x48,
	0x4e, 0xc1, 0x46, 0xf0, 0x1e, 0xcd, 0x42, 0xd7, 0x56, 0x53, 0x5f, 0x12, 0x7b, 0xb6, 0x35, 0x87,
	0x52, 0x6c, 0x43, 0x1b, 0x0f, 0xf5, 0xc3, 0xfa, 0x59, 0x64, 0x77, 0x8b, 0x7f, 0xc0, 0xcb, 0x32,
	0xf3, 0x66, 0x66, 0x07, 0xde, 0x83, 0xe1, 0xcb, 0x7e, 0xbd, 0xdd, 0xa4, 0xdd, 0x7e, 0xd7, 0xef,
	0x78, 0x60, 0x48, 0xb7, 0x8c, 0x3f, 0x1c, 0xa0, 0x85, 0xc6, 0x9c, 0x83, 0xb7, 0x7d, 0x7e, 0x6d,
	0x05, 0x89, 0x48, 0x12, 0xa2, 0xc1, 0x3c, 0x05, 0xda, 0x1f, 0xbb, 0x76, 0x2c, 0x9c, 0x88, 0x24,
	0xa3, 0x4c, 0xa4, 0xa7, 0x58, 0x6a, 0x22, 0xf6, 0xad, 0x8f, 0x5d, 0x8b, 0xd6, 0xc6, 0xc7, 0xd6,
	0x9f, 0x09, 0xd7, 0xf8, 0xcf, 0xff, 0xf2, 0x57, 0x6f, 0xcb, 0xef, 0x48, 0xc6, 0xcf, 0x80, 0x76,
	0xfb, 0xf5, 0xaa, 0x15, 0x5e, 0x44, 0x12, 0x07, 0x2d, 0xd1, 0xd3, 0x43, 0xbf, 0x5b, 0x6d, 0x04,
	0x8d, 0x48, 0x42, 0xd1, 0x92, 0xf8, 0x1a, 0xc2, 0xaf, 0x4a, 0x3e, 0x00, 0xef, 0x46, 0x4a, 0x64,
	0xff, 0x34, 0x5a, 0xa8, 0x07, 0xc9, 0x08, 0x07, 0xf0, 0xab, 0xb9, 0x42, 0x55, 0x33, 0x87, 0x87,
	0x40, 0xcb, 0x7a, 0x22, 0x91, 0xb9, 0xf1, 0x3b, 0xfc, 0xff, 0x59, 0xcd, 0x87, 0x10, 0xcc, 0xf3,
	0x99, 0xcc, 0x67, 0xd2, 0xa6, 0x8b, 0x1c, 0xa7, 0x8c, 0xf0, 0x00, 0x5c, 0x94, 0x85, 0x8d, 0x2e,
	0x26, 0xaa, 0x96, 0xcc, 0xd5, 0x2a, 0x96, 0x95, 0x64, 0x9e, 0x41, 0x93, 0xe6, 0x9e, 0x51, 0x2d,
	0x3f, 0x95, 0xc5, 0x34, 0x67, 0xbe, 0xfe, 0x69, 0xa6, 0x1e, 0x1b, 0xd9, 0x20, 0x0b, 0xf4, 0xfc,
	0xae, 0x51, 0xb7, 0x92, 0x0d, 0xb4, 0xb9, 0x2a, 0x8b, 0x9c, 0x85, 0xf1, 0x25, 0xf8, 0xa6, 0xfb,
	0xc0, 0x2f, 0xc0, 0x37, 0xeb, 0x38, 0x08, 0x12, 0xb9, 0xc9, 0x30, 0x1b, 0xfd, 0xde, 0x0e, 0x9e,
	0xd4, 0xa5, 0x6f, 0x4e, 0x74, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x71, 0x7e, 0xc8, 0xb1,
	0x01, 0x00, 0x00,
}