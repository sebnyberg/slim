// Code generated by protoc-gen-go. DO NOT EDIT.
// source: polyarray.proto

package array

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

// PolyArray uses one or more polynomial to compress and store an array of int32.
//
// The general idea
//
// We use a y = a + b x + c x^2 to describe trend of the numbers.
// And for every number we add a residual to fit the gap between y(i) and
// nums[i].
// E.g. If there are 4 numbers: 0, 15, 33, 50
// The polynomial and residuals could be:
//     y = 16 x
//     0, -1, 1, 2
//
// And the residuals require only 3 bits for any of them.
// To retrieve the numbers, we evaluate y(i) and add the residual to it:
//     get(0) = y(0) + 0 = 16 * 0 + 0 = 0
//     get(1) = y(1) - 1 = 16 * 1 - 1 = 15
//     get(2) = y(2) + 1 = 16 * 2 + 1 = 33
//     get(3) = y(3) + 2 = 16 * 3 + 2 = 50
//
// Data structure
//
// PolyArray splits the entire array into segment of 1024 elts.
// And then splits every segment into several spans.
// Every span has its own polynomial. A span is of length = 16*k
// Thus a segment has at most 64 spans.
//
// For every segment there is a uint64 bitmap to describe what spans a
// polynomial spans:
// A "1" at i and a "1" at j means a polynomial spans nums[(i+1)*16:(j+1)*16]
//
//     000101110000...(all 0)...1
//     <-- least significant bit
//
// In the above example, there are 5 polinomials, each of them spans:
// nums[0:64], nums[64:96], nums[96:112], nums[112:128], nums[128:1024]
//
// The "Bitmap" array allocates 2 uint64 for every segment:
// One of them is bitmap, another is bitmap rank index for speeding up locating
// a polynomial.
//
// Every polynomial needs 4 float64 in the "Polynomials" array:
// In our implementation polynomials are all of degree 2:
//    y = a + b x + c x^2
//
// Thus to store a polynomial we need 3 float64.
// And we use another float64 to store additional info:
// The width in bit of residuals, and where the first residual is in array
// "Residuals"
//
// Since 0.5.2
type PolyArray struct {
	N                    int32     `protobuf:"varint,10,opt,name=N,proto3" json:"N,omitempty"`
	Bitmap               []uint64  `protobuf:"varint,20,rep,packed,name=Bitmap,proto3" json:"Bitmap,omitempty"`
	Polynomials          []float64 `protobuf:"fixed64,21,rep,packed,name=Polynomials,proto3" json:"Polynomials,omitempty"`
	Residuals            []uint64  `protobuf:"varint,22,rep,packed,name=Residuals,proto3" json:"Residuals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PolyArray) Reset()         { *m = PolyArray{} }
func (m *PolyArray) String() string { return proto.CompactTextString(m) }
func (*PolyArray) ProtoMessage()    {}
func (*PolyArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_polyarray_ca9acf8057c5185f, []int{0}
}
func (m *PolyArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolyArray.Unmarshal(m, b)
}
func (m *PolyArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolyArray.Marshal(b, m, deterministic)
}
func (dst *PolyArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolyArray.Merge(dst, src)
}
func (m *PolyArray) XXX_Size() int {
	return xxx_messageInfo_PolyArray.Size(m)
}
func (m *PolyArray) XXX_DiscardUnknown() {
	xxx_messageInfo_PolyArray.DiscardUnknown(m)
}

var xxx_messageInfo_PolyArray proto.InternalMessageInfo

func (m *PolyArray) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *PolyArray) GetBitmap() []uint64 {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *PolyArray) GetPolynomials() []float64 {
	if m != nil {
		return m.Polynomials
	}
	return nil
}

func (m *PolyArray) GetResiduals() []uint64 {
	if m != nil {
		return m.Residuals
	}
	return nil
}

func init() {
	proto.RegisterType((*PolyArray)(nil), "PolyArray")
}

func init() { proto.RegisterFile("polyarray.proto", fileDescriptor_polyarray_ca9acf8057c5185f) }

var fileDescriptor_polyarray_ca9acf8057c5185f = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc8, 0xcf, 0xa9,
	0x4c, 0x2c, 0x2a, 0x4a, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x2a, 0xe4, 0xe2, 0x0c,
	0xc8, 0xcf, 0xa9, 0x74, 0x04, 0x09, 0x09, 0xf1, 0x70, 0x31, 0xfa, 0x49, 0x70, 0x29, 0x30, 0x6a,
	0xb0, 0x06, 0x31, 0xfa, 0x09, 0x89, 0x71, 0xb1, 0x39, 0x65, 0x96, 0xe4, 0x26, 0x16, 0x48, 0x88,
	0x28, 0x30, 0x6b, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x0a, 0x5c, 0xdc, 0x20, 0x2d, 0x79, 0xf9, 0xb9,
	0x99, 0x89, 0x39, 0xc5, 0x12, 0xa2, 0x0a, 0xcc, 0x1a, 0x8c, 0x41, 0xc8, 0x42, 0x42, 0x32, 0x5c,
	0x9c, 0x41, 0xa9, 0xc5, 0x99, 0x29, 0xa5, 0x20, 0x79, 0x31, 0xb0, 0x66, 0x84, 0x80, 0x13, 0x7b,
	0x14, 0x2b, 0xd8, 0x05, 0x49, 0x6c, 0x60, 0x27, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x47, 0x8d, 0xb0, 0x95, 0x00, 0x00, 0x00,
}
