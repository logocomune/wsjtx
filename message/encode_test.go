package message

import (
	"encoding/hex"
	"reflect"
	"testing"
)

const (
	testHeartbeatGenerated  = `adbccbda00000002000000000000000657534a542d580000000300000005322e342e3000000006633139643632`
	testClearGenerated      = `adbccbda00000002000000030000000657534a542d5814`
	testReplyGenerated      = `adbccbda00000002000000040000000657534a542d58000003e8fffffff43ff4cccccccccccd001b9e500000000346543800000004544553540000`
	testCloseGenerated      = `adbccbda00000002000000060000000657534a542d58`
	testReplayGenerated     = `adbccbda00000002000000070000000657534a542d58`
	testHaltTxGenerated     = `adbccbda00000002000000080000000657534a542d5801`
	testFreeTextGenerated   = `adbccbda00000002000000090000000657534a542d58000000074351205445535400`
	testLocationGenerated   = `adbccbda000000020000000b0000000657534a542d58000000064a4e35336572`
	testHighlightCallsign   = `adbccbda000000020000000d0000000657534a542d5800000005787878787801ffffffff00000000000001ffff0000ffff0000000000`
	testSwitchConfiguration = `adbccbda000000020000000e0000000657534a542d580000000661206e616d65`
	testConfiguration       = `adbccbda000000020000000f0000000657534a542d580000000346743800000000000000044e6f6e65000000006400000016000000055858585858000000064a4e3534657200`
)

func TestEncodeHearthBeat(t *testing.T) {
	type args struct {
		h HeartbeatMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "HearthBeat",
			args: args{
				h: HeartbeatMessage{
					ID:              "WSJT-X",
					MaxSchemaNumber: 3,
					Version:         "2.4.0",
					Revision:        "c19d62",
				},
			},
			want: hexToBytes(testHeartbeatGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHearthBeat(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeHearthBeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func hexToBytes(msg string) []byte {
	decodeString, _ := hex.DecodeString(msg)
	return decodeString
}

func TestEncodeClear(t *testing.T) {
	type args struct {
		c ClearMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Clear",
			args: args{
				c: ClearMessage{
					ID:      "WSJT-X",
					Windows: ClearRXFrequency,
				},
			},
			want: hexToBytes(testClearGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeClear(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeClear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeReply(t *testing.T) {
	type args struct {
		r ReplyMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Reply",
			args: args{
				r: ReplyMessage{
					ID:               "WSJT-X",
					MsSinceMN:        1000,
					SNR:              -12,
					DeltaTime:        1.30,
					DeltaFrequencyHZ: 1810000,
					Mode:             "FT8",
					Message:          "TEST",
					LowConfidence:    false,
					Modifiers:        ReplyNoModifier,
				},
			},
			want: hexToBytes(testReplyGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeReply(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeReply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeClose(t *testing.T) {
	type args struct {
		c CloseMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Close",
			args: args{
				c: CloseMessage{
					ID: "WSJT-X",
				},
			},
			want: hexToBytes(testCloseGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeClose(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeClose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeReplay(t *testing.T) {
	type args struct {
		r ReplayMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Replay",
			args: args{r: ReplayMessage{
				ID: "WSJT-X",
			}},
			want: hexToBytes(testReplayGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeReplay(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeReplay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeHaltTX(t *testing.T) {
	type args struct {
		h HaltTXMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Halt Tx",
			args: args{h: HaltTXMessage{
				ID:   "WSJT-X",
				Auto: HaltAtTheEnd,
			}},
			want: hexToBytes(testHaltTxGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHaltTX(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeHaltTX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeFreeText(t *testing.T) {
	type args struct {
		f FreeTextMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Free Text",
			args: args{f: FreeTextMessage{
				ID:   "WSJT-X",
				Text: "CQ TEST",
				Send: false,
			}},
			want: hexToBytes(testFreeTextGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeFreeText(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeFreeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeLocation(t *testing.T) {
	type args struct {
		l LocationMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Location",
			args: args{
				l: LocationMessage{
					ID:       "WSJT-X",
					Location: "JN53er",
				},
			},
			want: hexToBytes(testLocationGenerated),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeLocation(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeHighlightCallsign(t *testing.T) {
	type args struct {
		h HighlightCallsignMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Highlight Callsign",
			args: args{
				h: HighlightCallsignMessage{
					ID:       "WSJT-X",
					Callsign: "xxxxx",
					BackgroundColor: QColor{
						Alpha: AlphaOpaque,
						Red:   uint16(0xffff),
						Green: 0,
						Blue:  0,
					},
					ForegroundColor: QColor{
						Alpha: AlphaOpaque,
						Red:   0,
						Green: uint16(0xffff),
						Blue:  0,
					},
					HighlightLast: false,
				},
			},
			want: hexToBytes(testHighlightCallsign),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeHighlightCallsign(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeHighlightCallsign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeSwitchConfiguration(t *testing.T) {
	type args struct {
		s SwitchConfigurationMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Switch Configuration",
			args: args{
				s: SwitchConfigurationMessage{
					ID:                "WSJT-X",
					ConfigurationName: "a name",
				},
			},
			want: hexToBytes(testSwitchConfiguration),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeSwitchConfiguration(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeSwitchConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeConfigure(t *testing.T) {
	type args struct {
		c ConfigurationMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Configure",
			args: args{c: ConfigurationMessage{
				ID:                 "WSJT-X",
				Mode:               "Ft8",
				FrequencyTolerance: 0,
				Submode:            "None",
				FastMode:           false,
				TRPeriod:           100,
				RXDF:               22,
				DXCall:             "XXXXX",
				DXGrid:             "JN54er",
				GenerateMessage:    false,
			}},
			want: hexToBytes(testConfiguration),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeConfigure(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeConfigure() = %v, want %v", got, tt.want)
			}
		})
	}
}
