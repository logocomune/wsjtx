package message

import (
	"math"
	"time"
)

const (
	opModeNone     = 0
	opModeNAVhf    = 1
	opModeEUVhf    = 2
	opModeFieldDay = 3
	opModeRttyRU   = 4
	opModeWWDigi   = 5
	opModeFox      = 6
	opModeHound    = 7
)

// Parse messages send from WSJT-X on UDP
func Parse(buf []byte) (Response, error) {
	size := len(buf)
	if size == 0 || len(buf) == 0 {
		return Response{}, ErrMsgTooShort
	}

	mP := &msgDecoder{buf: buf[:size], len: size, pos: 0}
	magicElement, err := mP.decodeQUINT32()

	if err != nil {
		return Response{}, err
	}

	if magicElement != uint32(magic) {
		return Response{}, ErrInvalidMagic
	}

	_, err = mP.decodeQUINT32()
	if err != nil {
		return Response{}, err
	}

	messageType, _ := mP.decodeQUINT32()
	resp := Response{}

	switch messageType {
	case heartbeatType:
		resp.ResponseType = HeartbeatType
		resp.Message, err = mP.parseHeartbeat()

		if err != nil {
			return Response{}, err
		}
	case statusType:
		resp.ResponseType = StatusType
		resp.Message, err = mP.parseStatus()

		if err != nil {
			return Response{}, err
		}
	case decodeType:
		resp.ResponseType = DecodeType
		resp.Message, err = mP.parseDecode()

		if err != nil {
			return Response{}, err
		}
	case clearType:
		resp.ResponseType = ClearType
		resp.Message, err = mP.parseClear()

		if err != nil {
			return Response{}, err
		}
	case qsoLoggedType:
		resp.ResponseType = QSOLoggedType
		resp.Message, err = mP.parseQSOLoggedMessage()

		if err != nil {
			return Response{}, err
		}

	case closeType:
		resp.ResponseType = CloseType
		resp.Message, err = mP.parseCloseMessage()

		if err != nil {
			return Response{}, err
		}

	case wsprDecodeType:
		resp.ResponseType = WSPRDecodeType
		resp.Message, err = mP.parseWSPRDecodeMessage()

		if err != nil {
			return Response{}, err
		}

	case loggedADIFType:
		resp.ResponseType = LoggedADIFType
		resp.Message, err = mP.parseLoggedADIFMessage()

		if err != nil {
			return Response{}, err
		}
	default:
		return Response{}, ErrUnknownSchema
	}

	return resp, nil
}

func (m *msgDecoder) parseHeartbeat() (HeartbeatResponse, error) {
	r := HeartbeatResponse{}

	var err error
	r.ID, err = m.decodeUTF8()

	if err != nil {
		return r, err
	}

	r.MaxSchemaNumber, err = m.decodeQUINT32()

	if err != nil {
		return r, err
	}

	r.Version, err = m.decodeUTF8()

	if err != nil {
		return r, err
	}

	r.Revision, err = m.decodeUTF8()

	if err != nil {
		return r, err
	}

	return r, err
}

func (m *msgDecoder) parseStatus() (StatusResponse, error) {
	msg := StatusResponse{}

	var err error
	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Dial, err = m.decodeQUINT64(); err != nil {
		return msg, err
	}

	if msg.Mode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DXCall, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Report, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXMode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXEnabled, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.Transmitting, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.Decoding, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.RXDF, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.TXDF, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.DECall, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DEGrid, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DXGrid, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXWatchdog, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.SUBMode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.FastMode, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	special, err := m.decodeQUINT8()
	if err != nil {
		return msg, err
	}

	switch special {
	case opModeNone:
		msg.SpecialOperationMode = "NONE"
	case opModeNAVhf:
		msg.SpecialOperationMode = "NA VHF"
	case opModeEUVhf:
		msg.SpecialOperationMode = "EU VHF"
	case opModeFieldDay:
		msg.SpecialOperationMode = "FIELD DAY"
	case opModeRttyRU:
		msg.SpecialOperationMode = "RTTY RU"
	case opModeWWDigi:
		msg.SpecialOperationMode = "WW DIGI"
	case opModeFox:
		msg.SpecialOperationMode = "FOX"
	case opModeHound:
		msg.SpecialOperationMode = "HOUND"
	}

	if msg.FrequencyTolerance, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.TRPeriod, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.ConfigurationName, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXMessage, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseDecode() (DecodeResponse, error) {
	msg := DecodeResponse{}

	var err error

	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.New, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.Time, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.SNR, err = m.decodeQINT32(); err != nil {
		return msg, err
	}

	if msg.DeltaTime, err = m.decodeFloat(); err != nil {
		return msg, err
	}

	msg.FullTime = time.Unix(time.Now().Truncate(24*time.Hour).Add(time.Duration(msg.Time/1000)*time.Second).Unix(), 0).UTC()
	if msg.DeltaFrequencyHz, err = m.decodeQUINT32(); err != nil {
		return msg, err
	}

	if msg.Mode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Message, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.LowConfidence, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.OffAir, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseClear() (ClearResponse, error) {
	msg := ClearResponse{}

	var err error
	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseQSOLoggedMessage() (QSOLoggedResponse, error) {
	msg := QSOLoggedResponse{}

	var err error

	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DateAndTimeOff, err = m.decodeQDateTime(); err != nil {
		return msg, err
	}

	if msg.DXCall, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DXGrid, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXFrequencyHz, err = m.decodeQUINT64(); err != nil {
		return msg, err
	}

	if msg.Mode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ReportSent, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ReportReceived, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.TXPower, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Comments, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Name, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.DateAndTimeOn, err = m.decodeQDateTime(); err != nil {
		return msg, err
	}

	if msg.OperatorCall, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.MyCall, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.MyGrid, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ExchangeSent, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ExchangeReceived, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ADIFPropagationMode, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseCloseMessage() (CloseResponse, error) {
	msg := CloseResponse{}

	var err error

	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseWSPRDecodeMessage() (WSPRDecodeResponse, error) {
	msg := WSPRDecodeResponse{}

	var err error

	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.New, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	if msg.Time, msg.FullTime, err = m.decodeQTime(); err != nil {
		return msg, err
	}

	if msg.SNR, err = m.decodeQINT32(); err != nil {
		return msg, err
	}

	if msg.DeltaTime, err = m.decodeFloat(); err != nil {
		return msg, err
	}

	if msg.FrequencyHz, err = m.decodeQUINT64(); err != nil {
		return msg, err
	}

	if msg.DriftHz, err = m.decodeQINT32(); err != nil {
		return msg, err
	}

	if msg.Callsign, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.Grid, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.PowerdBm, err = m.decodeQINT32(); err != nil {
		return msg, err
	}

	msg.PowerWatt = dBmToWatt(msg.PowerdBm)

	if msg.OffAir, err = m.decodeBoolean(); err != nil {
		return msg, err
	}

	return msg, nil
}

func (m *msgDecoder) parseLoggedADIFMessage() (LoggedADIFResponse, error) {
	msg := LoggedADIFResponse{}

	var err error

	if msg.ID, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	if msg.ADIF, err = m.decodeUTF8(); err != nil {
		return msg, err
	}

	return msg, nil
}

func dBmToWatt(dBm int32) float64 {
	return math.Pow(10, float64(dBm)/10) / 1000
}
