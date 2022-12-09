package widgets

import (
	"fmt"
	"time"
)

const (
	LogRecTypeInfo = "info"
	LogRecTypeWarn = "warn"
	LogRecTypeErr  = "err"
)

type LogRecord struct {
	moment  time.Time
	recType string // info, warn, err
	message string
}

func NewLogRecordPtr(mes, t string) *LogRecord {
	lr := NewLogRecord(mes, t)

	return &lr
}

func NewLogRecord(mes, t string) LogRecord {
	return LogRecord{
		moment:  time.Now(),
		recType: t,
		message: mes,
	}
}

func (lr *LogRecord) GetMomentFormat() string {
	moment := fmt.Sprintf(
		"%02d.%02d.%02d %02d:%02d",
		lr.moment.Day(),
		lr.moment.Month(),
		lr.moment.Year(),
		lr.moment.Hour(),
		lr.moment.Minute())

	return fmt.Sprintf("[%s](fg:white,mod:bold)", moment)
}

func (lr *LogRecord) GetRecType() string {
	switch lr.recType {
	case LogRecTypeInfo:
		return fmt.Sprintf("[%s](fg:green,mod:bold)", lr.recType)
	case LogRecTypeWarn:
		return fmt.Sprintf("[%s](fg:yellow,mod:bold)", lr.recType)
	case LogRecTypeErr:
		return fmt.Sprintf("[%s](fg:red,mod:bold)", lr.recType)
	default:
		return ""
	}
}

func (lr *LogRecord) GetMsg() string {
	return lr.message
}

func (lr *LogRecord) Copy() LogRecord {
	nLR := NewLogRecord(lr.message, lr.recType)
	nLR.moment = lr.moment

	return nLR
}
