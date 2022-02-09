package message

import (
	"encoding/binary"
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/julian"
)

const (
	boolSize         = 1
	quint8Size       = 1
	quint32Size      = 4
	quint64Size      = 8
	floatSize        = 8
	zeroLengthString = uint32(0xffffffff)

	aDay = 24 * time.Hour
)

type msgDecoder struct {
	buf []byte
	len int
	pos int
}

func (m *msgDecoder) decodeBoolean() (bool, error) {
	end := m.pos
	if m.len < end {
		return false, ErrMsgTooShort
	}

	m.pos = end + boolSize

	return m.buf[end] != 0, nil
}

func (m *msgDecoder) decodeQUINT8() (uint8, error) {
	end := m.pos
	if m.len < end+quint8Size {
		return 0, ErrMsgTooShort
	}

	m.pos = end + quint8Size

	return m.buf[end], nil
}

func (m *msgDecoder) decodeQUINT32() (uint32, error) {
	end := m.pos + quint32Size
	if m.len < end {
		return 0, ErrMsgTooShort
	}

	u := binary.BigEndian.Uint32(m.buf[m.pos:end])
	m.pos = end

	return u, nil
}

func (m *msgDecoder) decodeQUINT64() (uint64, error) {
	end := m.pos + quint64Size
	if m.len < end {
		return 0, ErrMsgTooShort
	}

	u := binary.BigEndian.Uint64(m.buf[m.pos:end])
	m.pos = end

	return u, nil
}

func (m *msgDecoder) decodeQINT32() (int32, error) {
	quint32, err := m.decodeQUINT32()

	return int32(quint32), err
}

func (m *msgDecoder) decodeFloat() (float64, error) {
	end := m.pos + floatSize
	if m.len < end {
		return 0, ErrMsgTooShort
	}

	bits := binary.BigEndian.Uint64(m.buf[m.pos:end])
	u := math.Float64frombits(bits)
	m.pos = end

	return u, nil
}

func (m *msgDecoder) decodeUTF8() (string, error) {
	u, err := m.decodeQUINT32()
	if err != nil {
		return "", err
	}

	if u == zeroLengthString {
		return "", nil
	}

	end := m.pos + int(u)

	if m.len < end {
		return "", ErrMsgTooShort
	}

	s := string(m.buf[m.pos:end])
	m.pos = end

	return s, nil
}

func (m *msgDecoder) decodeQDateTime() (time.Time, error) {
	jd, err := m.decodeQUINT64()
	if err != nil {
		return time.Time{}, err
	}

	year, month, dayF := julian.JDToCalendar(float64(jd))
	day := int(dayF)

	msFromMD, err := m.decodeQUINT32()
	if err != nil {
		return time.Time{}, err
	}

	timespec, err := m.decodeQUINT8()
	if err != nil {
		return time.Time{}, err
	}

	if timespec > 1 {
		return time.Time{}, ErrDateTimeFormat
	}

	timeLocation := time.UTC

	if timespec == 0 {
		timeLocation = time.Local
	}

	dt := time.Date(year, time.Month(month), day, 0, 0, 0, 0, timeLocation)

	epoch := dt.Unix() + int64(msFromMD/1000)

	return time.Unix(epoch, 0).UTC(), nil
}

func (m *msgDecoder) decodeQTime() (uint32, time.Time, error) {
	now := time.Now().Truncate(aDay)
	msFromMD, err := m.decodeQUINT32()

	if err != nil {
		return 0, time.Time{}, err
	}

	return msFromMD, time.Unix(now.Unix()+int64(msFromMD/1000), 0).UTC(), nil
}
