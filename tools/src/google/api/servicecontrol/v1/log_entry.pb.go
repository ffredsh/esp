// Code generated by protoc-gen-go.
// source: google/api/servicecontrol/v1/log_entry.proto
// DO NOT EDIT!

package servicecontrol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google/api"
import _ "google/logging/type"
import google_logging_type1 "google/logging/type"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf2 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An individual log entry.
type LogEntry struct {
	// Required. The log to which this log entry belongs. Examples: `"syslog"`,
	// `"book_log"`.
	Name string `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
	// Optional. The time the event described by the log entry occurred. If
	// omitted, defaults to operation start time.
	Timestamp *google_protobuf3.Timestamp `protobuf:"bytes,11,opt,name=timestamp" json:"timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity google_logging_type1.LogSeverity `protobuf:"varint,12,opt,name=severity,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// A unique ID for the log entry used for deduplication. If omitted,
	// the implementation will generate one based on operation_id.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId" json:"insert_id,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_StructPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *google_protobuf1.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,oneof"`
}
type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,oneof"`
}
type LogEntry_StructPayload struct {
	StructPayload *google_protobuf2.Struct `protobuf:"bytes,6,opt,name=struct_payload,json=structPayload,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload()  {}
func (*LogEntry_TextPayload) isLogEntry_Payload()   {}
func (*LogEntry_StructPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *google_protobuf1.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetStructPayload() *google_protobuf2.Struct {
	if x, ok := m.GetPayload().(*LogEntry_StructPayload); ok {
		return x.StructPayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_StructPayload)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProtoPayload); err != nil {
			return err
		}
	case *LogEntry_TextPayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TextPayload)
	case *LogEntry_StructPayload:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructPayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 2: // payload.proto_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Any)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_ProtoPayload{msg}
		return true, err
	case 3: // payload.text_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogEntry_TextPayload{x}
		return true, err
	case 6: // payload.struct_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Struct)
		err := b.DecodeMessage(msg)
		m.Payload = &LogEntry_StructPayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// payload
	switch x := m.Payload.(type) {
	case *LogEntry_ProtoPayload:
		s := proto.Size(x.ProtoPayload)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_TextPayload:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TextPayload)))
		n += len(x.TextPayload)
	case *LogEntry_StructPayload:
		s := proto.Size(x.StructPayload)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.api.servicecontrol.v1.LogEntry")
}

func init() { proto.RegisterFile("google/api/servicecontrol/v1/log_entry.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x9b, 0xed, 0x52, 0x9a, 0x49, 0xbb, 0x42, 0xd6, 0x4a, 0x84, 0x50, 0x89, 0x08, 0x24,
	0xd4, 0x03, 0x72, 0xb4, 0xdd, 0xcb, 0xf2, 0xe7, 0x00, 0x95, 0x90, 0x0a, 0xea, 0x61, 0x95, 0xe5,
	0x5e, 0xb9, 0xed, 0x10, 0x2c, 0x52, 0x3b, 0x24, 0xd3, 0x88, 0x7c, 0x64, 0xbe, 0x05, 0x8a, 0x63,
	0x97, 0x65, 0x2b, 0xf5, 0x66, 0x7b, 0x7e, 0x6f, 0xe6, 0xbd, 0x31, 0xbc, 0xc9, 0xb4, 0xce, 0x72,
	0x4c, 0x44, 0x21, 0x93, 0x0a, 0xcb, 0x5a, 0x6e, 0x70, 0xa3, 0x15, 0x95, 0x3a, 0x4f, 0xea, 0xab,
	0x24, 0xd7, 0xd9, 0x0a, 0x15, 0x95, 0x0d, 0x2f, 0x4a, 0x4d, 0x9a, 0x4d, 0x3a, 0x9a, 0x8b, 0x42,
	0xf2, 0xff, 0x69, 0x5e, 0x5f, 0x45, 0x93, 0x7b, 0xbd, 0x84, 0x52, 0x9a, 0x04, 0x49, 0xad, 0xaa,
	0x4e, 0x1b, 0xbd, 0xb6, 0xd5, 0x5c, 0x67, 0x99, 0x54, 0x59, 0x42, 0x4d, 0x81, 0xc9, 0x0f, 0xa2,
	0x62, 0x55, 0xe2, 0xaf, 0x3d, 0x56, 0x74, 0x8a, 0x6b, 0x8d, 0x54, 0x58, 0x63, 0x29, 0xc9, 0x7a,
	0x89, 0x9e, 0x59, 0xce, 0xdc, 0xd6, 0xfb, 0xef, 0x89, 0x50, 0xae, 0x34, 0x79, 0x58, 0xaa, 0xa8,
	0xdc, 0x6f, 0xdc, 0x80, 0x17, 0x0f, 0xab, 0x24, 0x77, 0x58, 0x91, 0xd8, 0x15, 0x1d, 0xf0, 0xf2,
	0x4f, 0x1f, 0x86, 0x4b, 0x9d, 0x7d, 0x6e, 0x83, 0x33, 0x06, 0xe7, 0x4a, 0xec, 0x30, 0x84, 0xd8,
	0x9b, 0xfa, 0xa9, 0x39, 0xb3, 0x1b, 0xf0, 0x0f, 0x9a, 0x30, 0x88, 0xbd, 0x69, 0x30, 0x8b, 0xb8,
	0x5d, 0x8d, 0xeb, 0xca, 0xbf, 0x39, 0x22, 0xfd, 0x07, 0xb3, 0x0f, 0x30, 0x74, 0x31, 0xc2, 0x51,
	0xec, 0x4d, 0x2f, 0x66, 0xb1, 0x13, 0xda, 0xbc, 0xbc, 0xcd, 0xcb, 0x97, 0x3a, 0xbb, 0xb3, 0x5c,
	0x7a, 0x50, 0xb0, 0xe7, 0xe0, 0x4b, 0x55, 0x61, 0x49, 0x2b, 0xb9, 0x0d, 0xcf, 0x8d, 0xa1, 0x61,
	0xf7, 0xf0, 0x65, 0xcb, 0xbe, 0xc2, 0x20, 0x17, 0x6b, 0xcc, 0xab, 0x70, 0x1c, 0xf7, 0xa7, 0xc1,
	0x6c, 0xc6, 0x4f, 0x7d, 0x16, 0x77, 0x01, 0xf9, 0xd2, 0x88, 0xcc, 0x39, 0xb5, 0x1d, 0xd8, 0x7b,
	0x18, 0x9b, 0x1c, 0xab, 0x42, 0x34, 0xb9, 0x16, 0xdb, 0xf0, 0xcc, 0x84, 0xbc, 0x3c, 0x0a, 0xf9,
	0x49, 0x35, 0x8b, 0x5e, 0x3a, 0x32, 0xf7, 0xdb, 0x8e, 0x65, 0xaf, 0x60, 0x44, 0xf8, 0x9b, 0x0e,
	0xda, 0x7e, 0x6b, 0x74, 0xd1, 0x4b, 0x83, 0xf6, 0xd5, 0x41, 0x1f, 0xe1, 0xa2, 0xfb, 0x94, 0x03,
	0x36, 0x30, 0x23, 0x9e, 0x1e, 0x8d, 0xb8, 0x33, 0xd8, 0xa2, 0x97, 0x8e, 0x3b, 0x81, 0xed, 0x10,
	0xbd, 0x85, 0xe0, 0x9e, 0x75, 0xf6, 0x04, 0xfa, 0x3f, 0xb1, 0x09, 0x3d, 0xb3, 0x95, 0xf6, 0xc8,
	0x2e, 0xe1, 0x51, 0x2d, 0xf2, 0x3d, 0x1a, 0xf3, 0x7e, 0xda, 0x5d, 0xde, 0x9d, 0xdd, 0x78, 0x73,
	0x1f, 0x1e, 0xdb, 0xa9, 0xf3, 0x6b, 0x88, 0x37, 0x7a, 0x77, 0x72, 0x55, 0xf3, 0xb1, 0xdb, 0xd5,
	0xad, 0x89, 0xe9, 0xad, 0x07, 0xc6, 0xdc, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x33,
	0x47, 0x11, 0x3d, 0x03, 0x00, 0x00,
}
