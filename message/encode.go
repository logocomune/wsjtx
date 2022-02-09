package message

func EncodeHearthBeat(h HeartbeatMessage) []byte {
	e := newEncoder(schemaNumber, heartbeatType)
	e.encodeUTF8(h.ID)
	e.encodeQUInt32(h.MaxSchemaNumber)
	e.encodeUTF8(h.Version)
	e.encodeUTF8(h.Revision)

	return e.bytes()
}

func EncodeClear(c ClearMessage) []byte {
	e := newEncoder(schemaNumber, clearType)
	e.encodeUTF8(c.ID)
	e.encodeQUInt8(c.Windows)

	return e.bytes()
}

func EncodeReply(r ReplyMessage) []byte {
	e := newEncoder(schemaNumber, replyType)
	e.encodeUTF8(r.ID)
	e.encodeQUInt32(r.MsSinceMN)
	e.encodeQInt32(r.SNR)
	e.encodeQFloat(r.DeltaTime)
	e.encodeQUInt32(r.DeltaFrequencyHZ)
	e.encodeUTF8(r.Mode)
	e.encodeUTF8(r.Message)
	e.encodeBoolean(r.LowConfidence)
	e.encodeQUInt8(r.Modifiers)

	return e.bytes()
}

func EncodeClose(c CloseMessage) []byte {
	e := newEncoder(schemaNumber, closeType)
	e.encodeUTF8(c.ID)

	return e.bytes()
}

func EncodeReplay(r ReplayMessage) []byte {
	e := newEncoder(schemaNumber, replayType)
	e.encodeUTF8(r.ID)

	return e.bytes()
}

func EncodeHaltTX(h HaltTXMessage) []byte {
	e := newEncoder(schemaNumber, haltTxType)
	e.encodeUTF8(h.ID)
	e.encodeBoolean(h.Auto)

	return e.bytes()
}

func EncodeFreeText(f FreeTextMessage) []byte {
	e := newEncoder(schemaNumber, freeTextType)
	e.encodeUTF8(f.ID)
	e.encodeUTF8(f.Text)
	e.encodeBoolean(f.Send)

	return e.bytes()
}

func EncodeLocation(l LocationMessage) []byte {
	e := newEncoder(schemaNumber, locationType)
	e.encodeUTF8(l.ID)
	e.encodeUTF8(l.Location)

	return e.bytes()
}

func EncodeHighlightCallsign(h HighlightCallsignMessage) []byte {
	e := newEncoder(schemaNumber, highlightCallsignType)
	e.encodeUTF8(h.ID)
	e.encodeUTF8(h.Callsign)
	e.encodeQColor(h.BackgroundColor)
	e.encodeQColor(h.ForegroundColor)
	e.encodeBoolean(h.HighlightLast)

	return e.bytes()
}

func EncodeSwitchConfiguration(s SwitchConfigurationMessage) []byte {
	e := newEncoder(schemaNumber, switchConfigurationType)
	e.encodeUTF8(s.ID)
	e.encodeUTF8(s.ConfigurationName)

	return e.bytes()
}

func EncodeConfigure(c ConfigurationMessage) []byte {
	e := newEncoder(schemaNumber, configureType)
	e.encodeUTF8(c.ID)
	e.encodeUTF8(c.Mode)
	e.encodeQUInt32(c.FrequencyTolerance)
	e.encodeUTF8(c.Submode)
	e.encodeBoolean(c.FastMode)
	e.encodeQUInt32(c.TRPeriod)
	e.encodeQUInt32(c.RXDF)
	e.encodeUTF8(c.DXCall)
	e.encodeUTF8(c.DXGrid)
	e.encodeBoolean(c.GenerateMessage)

	return e.bytes()
}
