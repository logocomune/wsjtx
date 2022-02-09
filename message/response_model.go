package message

import "time"

const (
	HeartbeatType  = "HEARTBEAT"
	StatusType     = "STATUS"
	DecodeType     = "DECODE"
	ClearType      = "CLEAR"
	QSOLoggedType  = "QSOLogged"
	CloseType      = "CLOSE"
	WSPRDecodeType = "WSPRDecode"
	LoggedADIFType = "LoggedADIF"
)

type Response struct {
	ResponseType string
	Message      interface{}
}

type HeartbeatResponse struct {
	ID              string `json:"id"`
	MaxSchemaNumber uint32 `json:"maxSchemaNumber"`
	Version         string `json:"version"`
	Revision        string `json:"revision"`
}

type StatusResponse struct {
	ID                   string `json:"id"`
	Dial                 uint64 `json:"dial"`
	Mode                 string `json:"mode"`
	DXCall               string `json:"dxCall"`
	Report               string `json:"report"`
	TXMode               string `json:"txMode"`
	TXEnabled            bool   `json:"txEnabled"`
	Transmitting         bool   `json:"transmitting"`
	Decoding             bool   `json:"decoding"`
	RXDF                 uint32 `json:"rxDf"`
	TXDF                 uint32 `json:"txDf"`
	DECall               string `json:"deCall"`
	DEGrid               string `json:"deGrid"`
	DXGrid               string `json:"dxGrid"`
	TXWatchdog           bool   `json:"tXWatchdog"`
	SUBMode              string `json:"subMode"`
	FastMode             bool   `json:"fastMode"`
	SpecialOperationMode string `json:"specialOperationMode"`
	FrequencyTolerance   uint32 `json:"frequencyTolerance"`
	TRPeriod             uint32 `json:"trPeriod"`
	ConfigurationName    string `json:"configurationName"`
	TXMessage            string `json:"txMessage"`
}

type DecodeResponse struct {
	ID               string    `json:"id"`
	New              bool      `json:"new"`
	Time             uint32    `json:"time"`
	FullTime         time.Time `json:"fullTime"`
	SNR              int32     `json:"snr"`
	DeltaTime        float64   `json:"deltaTime"`
	DeltaFrequencyHz uint32    `json:"deltaFrequencyHz"`
	Mode             string    `json:"mode"`
	Message          string    `json:"message"`
	LowConfidence    bool      `json:"lowConfidence"`
	OffAir           bool      `json:"offAir"`
}

type ClearResponse struct {
	ID      string `json:"id"`
	Windows uint8  `json:"windows"`
}

type QSOLoggedResponse struct {
	ID                  string    `json:"id"`
	DateAndTimeOff      time.Time `json:"dateAndTimeOff"`
	DXCall              string    `json:"dxCall"`
	DXGrid              string    `json:"dxGrid"`
	TXFrequencyHz       uint64    `json:"txFrequencyHz"`
	Mode                string    `json:"mode"`
	ReportSent          string    `json:"reportSent"`
	ReportReceived      string    `json:"reportReceived"`
	TXPower             string    `json:"txPower"`
	Comments            string    `json:"comments"`
	Name                string    `json:"name"`
	DateAndTimeOn       time.Time `json:"dateAndTimeOn"`
	OperatorCall        string    `json:"operatorCall"`
	MyCall              string    `json:"myCall"`
	MyGrid              string    `json:"myGrid"`
	ExchangeSent        string    `json:"exchangeSent"`
	ExchangeReceived    string    `json:"exchangeReceived"`
	ADIFPropagationMode string    `json:"adifPropagationMode"`
}

type CloseResponse struct {
	ID string `json:"id"`
}

type WSPRDecodeResponse struct {
	ID          string    `json:"id"`
	New         bool      `json:"new"`
	Time        uint32    `json:"timeMs"`
	FullTime    time.Time `json:"fullTime"`
	SNR         int32     `json:"snr"`
	DeltaTime   float64   `json:"deltaTime"`
	FrequencyHz uint64    `json:"frequencyHz"`
	DriftHz     int32     `json:"driftHz"`
	Callsign    string    `json:"callsign"`
	Grid        string    `json:"grid"`
	PowerdBm    int32     `json:"powerdBm"`
	PowerWatt   float64   `json:"powerWatt"`
	OffAir      bool      `json:"offAir"`
}

type LoggedADIFResponse struct {
	ID   string `json:"id"`
	ADIF string `json:"adif"`
}
