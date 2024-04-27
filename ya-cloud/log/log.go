package log

type message struct {
	Message struct {
		Msg string `json:"msg"`
	} `json:"message"`
	Level messageType
}

func newMessage(msgType messageType, msg string) *message {
	return &message{
		Message: struct {
			Msg string `json:"msg"`
		}{Msg: msg},
		Level: msgType,
	}
}

type messageType string

const (
	TRACE messageType = "TRACE"
	DEBUG             = "DEBUG"
	INFO              = "INFO"
	WARN              = "WARN"
	ERROR             = "ERROR"
	FATAL             = "FATAL"
)

var loggerKey = "logger"
