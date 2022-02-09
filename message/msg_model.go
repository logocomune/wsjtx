package message

import (
	"time"
)

const (
	magic     = 0xadbccbda
	magicUint = uint32(magic)

	schemaNumber = 2

	heartbeatType           = 0  // Heartbeat Value
	statusType              = 1  // Status Value
	decodeType              = 2  // Decode Value
	clearType               = 3  // Clear Value
	replyType               = 4  // Reply Value
	qsoLoggedType           = 5  // Qso Logged Value
	closeType               = 6  // Close Value
	replayType              = 7  // Replay Value
	haltTxType              = 8  // Halt Tx Value
	freeTextType            = 9  // Free Text Value
	wsprDecodeType          = 10 // Wspr Decode Value
	locationType            = 11 // Locator Value
	loggedADIFType          = 12 // Logged ADIF Value
	highlightCallsignType   = 13 // Highlight Value
	switchConfigurationType = 14 // Switch ConfigurationMessage Value
	configureType           = 15 // Configure Value

	ClearBandActivity = iota
	ClearRXFrequency
	ClearBandAndRXFrequency

	ReplyNoModifier      = uint8(0x00)
	ReplyShiftModifier   = uint8(0x02)
	ReplyControlModifier = uint8(0x04)
	ReplyAltModifier     = uint8(0x08)
	ReplyMetaModifier    = uint8(0x10)
	ReplyKeypadModifier  = uint8(0x20)
	ReplyGroupModifier   = uint8(0x40)

	HaltTxImmediately = false
	HaltAtTheEnd      = true
)

type HeartbeatMessage struct {
	ID              string
	MaxSchemaNumber uint32
	Version         string
	Revision        string
}

type StatusMessage struct {
	ID                   string
	Dial                 uint64
	Mode                 string
	DXCall               string
	Report               string
	TXMode               string
	TXEnabled            bool
	Transmitting         bool
	Decoding             bool
	RXDF                 uint32
	TXDF                 uint32
	DECall               string
	DEGrid               string
	DXGrid               string
	TXWatchdog           bool
	SUBMode              string
	FastMode             bool
	SpecialOperationMode string
	FrequencyTolerance   uint32
	TRPeriod             uint32
	ConfigurationName    string
	TXMessage            string
}

type DecodeMessage struct {
	ID               string
	New              bool
	Time             uint32
	FullTime         time.Time
	SNR              int32
	DeltaTime        float64
	DeltaFrequencyHZ uint32
	Mode             string
	Message          string
	LowConfidence    bool
	OffAir           bool
}

type ClearMessage struct {
	ID      string
	Windows uint8
}

type ReplyMessage struct {
	ID               string
	MsSinceMN        uint32
	SNR              int32
	DeltaTime        float64
	DeltaFrequencyHZ uint32
	Mode             string
	Message          string
	LowConfidence    bool
	Modifiers        uint8
}

type QSOLoggedMessage struct {
	ID                  string
	DateAndTimeOff      time.Time
	DXCall              string
	DXGrid              string
	TXFrequencyHZ       uint64
	Mode                string
	ReportSent          string
	ReportReceived      string
	TXPower             string
	Comments            string
	Name                string
	DateAndTimeOn       time.Time
	OperatorCall        string
	MyCall              string
	MyGrid              string
	ExchangeSent        string
	ExchangeReceived    string
	ADIFPropagationMode string
}

type CloseMessage struct {
	ID string
}

type ReplayMessage struct {
	ID string
}

type HaltTXMessage struct {
	ID   string
	Auto bool
}

type FreeTextMessage struct {
	ID   string
	Text string
	Send bool
}

type WSPRDecodeMessage struct {
	ID          string
	New         bool
	Time        uint32
	FullTime    time.Time
	SNR         int32
	DeltaTime   float64
	FrequencyHZ uint64
	DriftHz     int32
	Callsign    string
	Grid        string
	PowerdBm    int32
	PowerWatts  float64
	OffAir      bool
}

type LocationMessage struct {
	ID       string
	Location string
}

type HighlightCallsignMessage struct {
	ID              string
	Callsign        string
	BackgroundColor QColor
	ForegroundColor QColor
	HighlightLast   bool
}

type QColor struct {
	Alpha uint16
	Red   uint16
	Green uint16
	Blue  uint16
}

type SwitchConfigurationMessage struct {
	ID                string
	ConfigurationName string
}

type ConfigurationMessage struct {
	ID                 string
	Mode               string
	FrequencyTolerance uint32
	Submode            string
	FastMode           bool
	TRPeriod           uint32
	RXDF               uint32
	DXCall             string
	DXGrid             string
	GenerateMessage    bool
}
