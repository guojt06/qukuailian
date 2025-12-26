package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel = 1
	InfoLevel  = 2
	WarnLevel  = 3
	ErrorLevel = 4
)

func (l Level) String() string {
	var str string
	switch l {
	case DebugLevel:
		str = "dubug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warn"
	case ErrorLevel:
		str = "error"
	default:
		str = "其他"
	}
	return str
}

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}
