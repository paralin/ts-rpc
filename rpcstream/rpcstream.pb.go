// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.6.0
// source: github.com/aperturerobotics/starpc/rpcstream/rpcstream.proto

package rpcstream

import (
	base64 "encoding/base64"
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

// RpcStreamPacket is a packet encapsulating data for a RPC stream.
type RpcStreamPacket struct {
	unknownFields []byte
	// Types that are assignable to Body:
	//
	//	*RpcStreamPacket_Init
	//	*RpcStreamPacket_Ack
	//	*RpcStreamPacket_Data
	Body isRpcStreamPacket_Body `protobuf_oneof:"body"`
}

func (x *RpcStreamPacket) Reset() {
	*x = RpcStreamPacket{}
}

func (*RpcStreamPacket) ProtoMessage() {}

func (m *RpcStreamPacket) GetBody() isRpcStreamPacket_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *RpcStreamPacket) GetInit() *RpcStreamInit {
	if x, ok := x.GetBody().(*RpcStreamPacket_Init); ok {
		return x.Init
	}
	return nil
}

func (x *RpcStreamPacket) GetAck() *RpcAck {
	if x, ok := x.GetBody().(*RpcStreamPacket_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *RpcStreamPacket) GetData() []byte {
	if x, ok := x.GetBody().(*RpcStreamPacket_Data); ok {
		return x.Data
	}
	return nil
}

type isRpcStreamPacket_Body interface {
	isRpcStreamPacket_Body()
}

type RpcStreamPacket_Init struct {
	// Init is the first packet in the stream.
	// Sent by the initiator.
	Init *RpcStreamInit `protobuf:"bytes,1,opt,name=init,proto3,oneof"`
}

type RpcStreamPacket_Ack struct {
	// Ack is sent in response to Init.
	// Sent by the server.
	Ack *RpcAck `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

type RpcStreamPacket_Data struct {
	// Data is the encapsulated data packet.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3,oneof"`
}

func (*RpcStreamPacket_Init) isRpcStreamPacket_Body() {}

func (*RpcStreamPacket_Ack) isRpcStreamPacket_Body() {}

func (*RpcStreamPacket_Data) isRpcStreamPacket_Body() {}

// RpcStreamInit is the first message in a RPC stream.
type RpcStreamInit struct {
	unknownFields []byte
	// ComponentId is the identifier of the component making the request.
	ComponentId string `protobuf:"bytes,1,opt,name=component_id,json=componentId,proto3" json:"componentId,omitempty"`
}

func (x *RpcStreamInit) Reset() {
	*x = RpcStreamInit{}
}

func (*RpcStreamInit) ProtoMessage() {}

func (x *RpcStreamInit) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

// RpcAck is the ack message in a RPC stream.
type RpcAck struct {
	unknownFields []byte
	// Error indicates there was some error setting up the stream.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RpcAck) Reset() {
	*x = RpcAck{}
}

func (*RpcAck) ProtoMessage() {}

func (x *RpcAck) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (m *RpcStreamPacket) CloneVT() *RpcStreamPacket {
	if m == nil {
		return (*RpcStreamPacket)(nil)
	}
	r := new(RpcStreamPacket)
	if m.Body != nil {
		r.Body = m.Body.(interface{ CloneVT() isRpcStreamPacket_Body }).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RpcStreamPacket) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *RpcStreamPacket_Init) CloneVT() *RpcStreamPacket_Init {
	if m == nil {
		return (*RpcStreamPacket_Init)(nil)
	}
	r := new(RpcStreamPacket_Init)
	r.Init = m.Init.CloneVT()
	return r
}

func (m *RpcStreamPacket_Ack) CloneVT() *RpcStreamPacket_Ack {
	if m == nil {
		return (*RpcStreamPacket_Ack)(nil)
	}
	r := new(RpcStreamPacket_Ack)
	r.Ack = m.Ack.CloneVT()
	return r
}

func (m *RpcStreamPacket_Data) CloneVT() *RpcStreamPacket_Data {
	if m == nil {
		return (*RpcStreamPacket_Data)(nil)
	}
	r := new(RpcStreamPacket_Data)
	if rhs := m.Data; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.Data = tmpBytes
	}
	return r
}

func (m *RpcStreamInit) CloneVT() *RpcStreamInit {
	if m == nil {
		return (*RpcStreamInit)(nil)
	}
	r := new(RpcStreamInit)
	r.ComponentId = m.ComponentId
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RpcStreamInit) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *RpcAck) CloneVT() *RpcAck {
	if m == nil {
		return (*RpcAck)(nil)
	}
	r := new(RpcAck)
	r.Error = m.Error
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *RpcAck) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (this *RpcStreamPacket) EqualVT(that *RpcStreamPacket) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Body == nil && that.Body != nil {
		return false
	} else if this.Body != nil {
		if that.Body == nil {
			return false
		}
		if !this.Body.(interface {
			EqualVT(isRpcStreamPacket_Body) bool
		}).EqualVT(that.Body) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RpcStreamPacket) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*RpcStreamPacket)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *RpcStreamPacket_Init) EqualVT(thatIface isRpcStreamPacket_Body) bool {
	that, ok := thatIface.(*RpcStreamPacket_Init)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Init, that.Init; p != q {
		if p == nil {
			p = &RpcStreamInit{}
		}
		if q == nil {
			q = &RpcStreamInit{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *RpcStreamPacket_Ack) EqualVT(thatIface isRpcStreamPacket_Body) bool {
	that, ok := thatIface.(*RpcStreamPacket_Ack)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if p, q := this.Ack, that.Ack; p != q {
		if p == nil {
			p = &RpcAck{}
		}
		if q == nil {
			q = &RpcAck{}
		}
		if !p.EqualVT(q) {
			return false
		}
	}
	return true
}

func (this *RpcStreamPacket_Data) EqualVT(thatIface isRpcStreamPacket_Body) bool {
	that, ok := thatIface.(*RpcStreamPacket_Data)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if string(this.Data) != string(that.Data) {
		return false
	}
	return true
}

func (this *RpcStreamInit) EqualVT(that *RpcStreamInit) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.ComponentId != that.ComponentId {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RpcStreamInit) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*RpcStreamInit)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *RpcAck) EqualVT(that *RpcAck) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Error != that.Error {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *RpcAck) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*RpcAck)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// MarshalProtoJSON marshals the RpcStreamPacket message to JSON.
func (x *RpcStreamPacket) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Body != nil {
		switch ov := x.Body.(type) {
		case *RpcStreamPacket_Init:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("init")
			ov.Init.MarshalProtoJSON(s.WithField("init"))
		case *RpcStreamPacket_Ack:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("ack")
			ov.Ack.MarshalProtoJSON(s.WithField("ack"))
		case *RpcStreamPacket_Data:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("data")
			s.WriteBytes(ov.Data)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the RpcStreamPacket to JSON.
func (x *RpcStreamPacket) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the RpcStreamPacket message from JSON.
func (x *RpcStreamPacket) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "init":
			ov := &RpcStreamPacket_Init{}
			x.Body = ov
			if s.ReadNil() {
				ov.Init = nil
				return
			}
			ov.Init = &RpcStreamInit{}
			ov.Init.UnmarshalProtoJSON(s.WithField("init", true))
		case "ack":
			ov := &RpcStreamPacket_Ack{}
			x.Body = ov
			if s.ReadNil() {
				ov.Ack = nil
				return
			}
			ov.Ack = &RpcAck{}
			ov.Ack.UnmarshalProtoJSON(s.WithField("ack", true))
		case "data":
			s.AddField("data")
			ov := &RpcStreamPacket_Data{}
			x.Body = ov
			ov.Data = s.ReadBytes()
		}
	})
}

// UnmarshalJSON unmarshals the RpcStreamPacket from JSON.
func (x *RpcStreamPacket) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the RpcStreamInit message to JSON.
func (x *RpcStreamInit) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ComponentId != "" || s.HasField("componentId") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("componentId")
		s.WriteString(x.ComponentId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the RpcStreamInit to JSON.
func (x *RpcStreamInit) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the RpcStreamInit message from JSON.
func (x *RpcStreamInit) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "component_id", "componentId":
			s.AddField("component_id")
			x.ComponentId = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the RpcStreamInit from JSON.
func (x *RpcStreamInit) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the RpcAck message to JSON.
func (x *RpcAck) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Error != "" || s.HasField("error") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("error")
		s.WriteString(x.Error)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the RpcAck to JSON.
func (x *RpcAck) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the RpcAck message from JSON.
func (x *RpcAck) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "error":
			s.AddField("error")
			x.Error = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the RpcAck from JSON.
func (x *RpcAck) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

func (m *RpcStreamPacket) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcStreamPacket) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcStreamPacket) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Body.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RpcStreamPacket_Init) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcStreamPacket_Init) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Init != nil {
		size, err := m.Init.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	} else {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RpcStreamPacket_Ack) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcStreamPacket_Ack) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Ack != nil {
		size, err := m.Ack.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	} else {
		i = protobuf_go_lite.EncodeVarint(dAtA, i, 0)
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *RpcStreamPacket_Data) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcStreamPacket_Data) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Data)
	copy(dAtA[i:], m.Data)
	i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Data)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *RpcStreamInit) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcStreamInit) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcStreamInit) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.ComponentId) > 0 {
		i -= len(m.ComponentId)
		copy(dAtA[i:], m.ComponentId)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.ComponentId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcAck) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcAck) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *RpcAck) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcStreamPacket) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Body.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RpcStreamPacket_Init) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Init != nil {
		l = m.Init.SizeVT()
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *RpcStreamPacket_Ack) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ack != nil {
		l = m.Ack.SizeVT()
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	} else {
		n += 2
	}
	return n
}
func (m *RpcStreamPacket_Data) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	return n
}
func (m *RpcStreamInit) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ComponentId)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RpcAck) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (x *RpcStreamPacket) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("RpcStreamPacket { ")
	switch body := x.Body.(type) {
	case *RpcStreamPacket_Init:
		if body.Init != nil {
			sb.WriteString(" init: ")
			sb.WriteString(body.Init.MarshalProtoText())
		}
	case *RpcStreamPacket_Ack:
		if body.Ack != nil {
			sb.WriteString(" ack: ")
			sb.WriteString(body.Ack.MarshalProtoText())
		}
	case *RpcStreamPacket_Data:
		if len(body.Data) > 0 {
			sb.WriteString(" data: ")
			sb.WriteString("\"")
			sb.WriteString(base64.StdEncoding.EncodeToString(body.Data))
			sb.WriteString("\"")
		}
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *RpcStreamPacket) String() string {
	return x.MarshalProtoText()
}
func (x *RpcStreamInit) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("RpcStreamInit { ")
	if x.ComponentId != "" {
		sb.WriteString(" component_id: ")
		sb.WriteString(strconv.Quote(x.ComponentId))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *RpcStreamInit) String() string {
	return x.MarshalProtoText()
}
func (x *RpcAck) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("RpcAck { ")
	if x.Error != "" {
		sb.WriteString(" error: ")
		sb.WriteString(strconv.Quote(x.Error))
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *RpcAck) String() string {
	return x.MarshalProtoText()
}
func (m *RpcStreamPacket) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: RpcStreamPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcStreamPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Init", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Body.(*RpcStreamPacket_Init); ok {
				if err := oneof.Init.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &RpcStreamInit{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Body = &RpcStreamPacket_Init{Init: v}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ack", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.Body.(*RpcStreamPacket_Ack); ok {
				if err := oneof.Ack.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &RpcAck{}
				if err := v.UnmarshalVT(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.Body = &RpcStreamPacket_Ack{Ack: v}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Body = &RpcStreamPacket_Data{Data: v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcStreamInit) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: RpcStreamInit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcStreamInit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComponentId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComponentId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcAck) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
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
			return fmt.Errorf("proto: RpcAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
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
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
