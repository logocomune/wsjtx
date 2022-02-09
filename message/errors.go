package message

import "errors"

var ErrMsgZeroSize = errors.New("parse error: message has 0 bytes")

var ErrMsgTooShort = errors.New("parse error: message too short")

var ErrInvalidMagic = errors.New("parse error: invalid magic")

var ErrUnknownSchema = errors.New("parse error: unknown schema")

var ErrDateTimeFormat = errors.New("parse error: invalid date/time format")
