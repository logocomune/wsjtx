package message

import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	rgbFormat        = 1
	AlphaTransparent = 0
	AlphaOpaque      = uint16(0xffff)
	padding          = uint16(0x0000)
)

type msgEncoder struct {
	buf bytes.Buffer
}

func newEncoder(schemaNumber uint32, messageType uint32) *msgEncoder {
	m := msgEncoder{
		buf: bytes.Buffer{},
	}

	m.encodeQUInt32(magicUint)
	m.encodeQUInt32(schemaNumber)
	m.encodeQUInt32(messageType)

	return &m
}

func (m *msgEncoder) bytes() []byte {
	return m.buf.Bytes()
}

func (m *msgEncoder) encodeUTF8(text string) {
	l := len(text)
	if l == 0 {
		m.encodeQUInt32(zeroLengthString)

		return
	}

	m.encodeQUInt32(uint32(l))
	m.buf.WriteString(text)
}

func (m *msgEncoder) encodeUint64(num uint64) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	binary.BigEndian.PutUint64(b, num)
	m.buf.Write(b)
}

func (m *msgEncoder) encodeQUInt32(i uint32) {
	b := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(b, i)
	m.buf.Write(b)
}

func (m *msgEncoder) encodeQUInt16(i uint16) {
	b := []byte{0, 0}
	binary.BigEndian.PutUint16(b, i)
	m.buf.Write(b)
}

func (m *msgEncoder) encodeQUInt8(i uint8) {
	m.buf.WriteByte(i)
}

func (m *msgEncoder) encodeQInt32(i int32) {
	m.encodeQUInt32(uint32(i))
}

func (m *msgEncoder) encodeQFloat(f float64) {
	bin := math.Float64bits(f)
	m.encodeUint64(bin)
}

func (m *msgEncoder) encodeBoolean(b bool) {
	if b {
		m.encodeQUInt8(1)
	} else {
		m.encodeQUInt8(0)
	}
}

func (m *msgEncoder) encodeQColor(q QColor) {
	m.encodeQUInt8(rgbFormat)
	m.encodeQUInt16(q.Alpha)
	m.encodeQUInt16(q.Red)
	m.encodeQUInt16(q.Green)
	m.encodeQUInt16(q.Blue)
	m.encodeQUInt16(padding)
}
