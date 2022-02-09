package message

import (
	"encoding/hex"
	"reflect"
	"testing"
	"time"
)

const (
	testHeartBeatMsg = `adbccbda00000002000000000000000657534a542d580000000300000005322e342e3000000006633139643632`
	testStatus       = `adbccbda00000002000000010000000657534a542d580000000001142f2000000003465438000000055858585858000000022d3600000003465438000001000003750000037500000006495535504d50000000064a4e3533455200000004494f393100ffffffff0000ffffffffffffffff0000000744656661756c7400000025585858585820495535504d50204a4e35332020202020202020202020202020202020202020`
	testDecode       = `adbccbda00000002000000020000000657534a542d5801021b3ee0fffffff1bfb99999a000000000000581000000017e000000105858585858205959595959204c4f31310000`
	testClear        = `adbccbda00000002000000030000000657534a542d58`
	testQSOLogged    = `adbccbda00000002000000050000000657534a542d5800000000002587df024adb8f01000000055959595959000000044a4e383600000000006bf37c00000003465438000000032b3132000000032d323400000002323000000019465438202053656e743a202b31322020526376643a202d32340000000000000000002587df0249f275010000000000000006495535504d50000000064a4e353345520000000000000000ffffffff`
	testClose        = `adbccbda00000002000000060000000657534a542d58`
	testWSPRDecode   = `adbccbda000000020000000a0000000657534a542d580103cff3c0ffffffff3fb99999a00000000000000000d71ac4000000000000000645413755524300000004494d37370000001b00`
	testLoggedAdif   = `adbccbda000000020000000c0000000657534a542d58000001550a3c616469665f7665723a353e332e312e300a3c70726f6772616d69643a363e57534a542d580a3c454f483e0a3c63616c6c3a353e5959595959203c677269647371756172653a343e4a4e3836203c6d6f64653a333e465438203c7273745f73656e743a333e2b3132203c7273745f726376643a333e2d3234203c71736f5f646174653a383e3230323230323034203c74696d655f6f6e3a363e313034303030203c71736f5f646174655f6f66663a383e3230323230323034203c74696d655f6f66663a363e313034313030203c62616e643a333e34306d203c667265713a383e372e303734363834203c73746174696f6e5f63616c6c7369676e3a363e495535504d50203c6d795f677269647371756172653a363e4a4e35334552203c74785f7077723a323e3230203c636f6d6d656e743a32353e465438202053656e743a202b31322020526376643a202d3234203c454f523e`
)

type pArgs struct {
	buf []byte
}

type pResult struct {
	msg interface{}
	err error
}

func TestParser_ParseMessage(t *testing.T) {
	y, m, d := time.Now().Date()
	type args struct {
		buf  []byte
		size int
	}
	tests := []struct {
		name    string
		args    pArgs
		want    interface{}
		wantErr bool
	}{
		{
			name: "Heartbeat",
			args: argsBuilder(testHeartBeatMsg),
			want: Response{
				ResponseType: HeartbeatType,
				Message: HeartbeatResponse{
					ID:              "WSJT-X",
					MaxSchemaNumber: 3,
					Version:         "2.4.0",
					Revision:        "c19d62",
				},
			},

			wantErr: false,
		},
		{
			name: "Status",
			args: argsBuilder(testStatus),
			want: Response{
				ResponseType: StatusType,
				Message: StatusResponse{
					ID:                   "WSJT-X",
					Dial:                 18100000,
					Mode:                 "FT8",
					DXCall:               "XXXXX",
					Report:               "-6",
					TXMode:               "FT8",
					TXEnabled:            false,
					Transmitting:         false,
					Decoding:             true,
					RXDF:                 885,
					TXDF:                 885,
					DECall:               "IU5PMP",
					DEGrid:               "JN53ER",
					DXGrid:               "IO91",
					TXWatchdog:           false,
					SUBMode:              "",
					FastMode:             false,
					SpecialOperationMode: "NONE",
					FrequencyTolerance:   4294967295,
					TRPeriod:             4294967295,
					ConfigurationName:    "Default",
					TXMessage:            "XXXXX IU5PMP JN53                    ",
				},
			},

			wantErr: false,
		},
		{
			name: "Decode",
			args: argsBuilder(testDecode),
			want: Response{
				ResponseType: DecodeType,
				Message: DecodeResponse{
					ID:               "WSJT-X",
					New:              true,
					Time:             35340000,
					FullTime:         time.Now().Truncate(24 * time.Hour).Add(time.Duration(35340000/1000) * time.Second).UTC(),
					SNR:              -15,
					DeltaTime:        -0.10000000149011612,
					DeltaFrequencyHz: 1409,
					Mode:             "~",
					Message:          "XXXXX YYYYY LO11",
					LowConfidence:    false,
					OffAir:           false,
				},
			},
			wantErr: false,
		},
		{
			name: "Clear",
			args: argsBuilder(testClear),
			want: Response{
				ResponseType: ClearType,
				Message: ClearResponse{
					ID:      "WSJT-X",
					Windows: 0,
				},
			},
			wantErr: false,
		},
		{
			name: "QSO Logged",
			args: argsBuilder(testQSOLogged),
			want: Response{
				ResponseType: QSOLoggedType,
				Message: QSOLoggedResponse{
					ID:                  "WSJT-X",
					DateAndTimeOff:      time.Date(2022, 0o2, 0o4, 10, 41, 0o0, 0o0, time.UTC).UTC(),
					DXCall:              "YYYYY",
					DXGrid:              "JN86",
					TXFrequencyHz:       7074684,
					Mode:                "FT8",
					ReportSent:          "+12",
					ReportReceived:      "-24",
					TXPower:             "20",
					Comments:            "FT8  Sent: +12  Rcvd: -24",
					Name:                "",
					DateAndTimeOn:       time.Date(2022, 0o2, 0o4, 10, 40, 0o0, 0o0, time.UTC).UTC(),
					OperatorCall:        "",
					MyCall:              "IU5PMP",
					MyGrid:              "JN53ER",
					ExchangeSent:        "",
					ExchangeReceived:    "",
					ADIFPropagationMode: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Close",
			args: argsBuilder(testClose),
			want: Response{
				ResponseType: CloseType,
				Message: CloseResponse{
					ID: "WSJT-X",
				},
			},
			wantErr: false,
		},
		{
			name: "WSPRDecode",
			args: argsBuilder(testWSPRDecode),
			want: Response{
				ResponseType: WSPRDecodeType,
				Message: WSPRDecodeResponse{
					ID:          "WSJT-X",
					New:         true,
					Time:        63960000,
					FullTime:    time.Date(y, m, d, 17, 46, 0o0, 0o0, time.UTC).UTC(),
					SNR:         -1,
					DeltaTime:   0.10000000149011612,
					FrequencyHz: 14097092,
					DriftHz:     0,
					Callsign:    "EA7URC",
					Grid:        "IM77",
					PowerdBm:    27,
					PowerWatt:   0.5011872336272725,
					OffAir:      false,
				},
			},
			wantErr: false,
		},
		{
			name: "LoggedADIF",
			args: argsBuilder(testLoggedAdif),
			want: Response{
				ResponseType: LoggedADIFType,
				Message: LoggedADIFResponse{
					ID: "WSJT-X",
					ADIF: `
<adif_ver:5>3.1.0
<programid:6>WSJT-X
<EOH>
<call:5>YYYYY <gridsquare:4>JN86 <mode:3>FT8 <rst_sent:3>+12 <rst_rcvd:3>-24 <qso_date:8>20220204 <time_on:6>104000 <qso_date_off:8>20220204 <time_off:6>104100 <band:3>40m <freq:8>7.074684 <station_callsign:6>IU5PMP <my_gridsquare:6>JN53ER <tx_pwr:2>20 <comment:25>FT8  Sent: +12  Rcvd: -24 <EOR>`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func argsBuilder(msg string) pArgs {
	decodeString, _ := hex.DecodeString(msg)
	return pArgs{
		buf: decodeString,
	}
}

func Test_msgParser_parseStatus(t *testing.T) {
	type fields struct {
		buf []byte
		len int
		pos int
	}
	tests := []struct {
		name    string
		fields  fields
		want    StatusMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &msgDecoder{
				buf: tt.fields.buf,
				len: tt.fields.len,
				pos: tt.fields.pos,
			}
			got, err := m.parseStatus()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
