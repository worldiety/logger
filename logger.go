package logger //TODO change me

import (
	"fmt"
	"time"
)

var logger Logger = &defaultLogger{}

// SetLogger sets the module/package level logger
func SetLogger(log Logger) {
	logger = log
}

// Fields is just a type alias to avoid verbosity while using the logger
type Fields = map[string]interface{}

// A Logger is just a simple interface which can be easily satisfied by any implementor
type Logger interface {
	Debug(fields Fields)
	Info(fields Fields)
	Warn(fields Fields)
	Error(fields Fields)
}

// the default logger just prints as json into stdout
type defaultLogger struct {
}

func (l *defaultLogger) Info(fields map[string]interface{}) {
	fields["level"] = "INFO"
	l.log(fields)
}

func (l *defaultLogger) Debug(fields map[string]interface{}) {
	fields["level"] = "DEBUG"
	l.log(fields)
}

func (l *defaultLogger) Warn(fields map[string]interface{}) {
	fields["level"] = "WARN"
	l.log(fields)
}

func (l *defaultLogger) Error(fields map[string]interface{}) {
	fields["level"] = "ERROR"
	l.log(fields)
}

func (l *defaultLogger) log(fields map[string]interface{}) {
	fields["ts"] = time.Now().String()
	fields["pkg"] = "mypackage" //TODO change me
	tmp, _ := json.Marshal(fields)
	fmt.Println(string(tmp))
}
