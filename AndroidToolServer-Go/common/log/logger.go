package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

const (
	sFunctionName = "S-FunctionName"
	sFunctionLine = "S-FunctionLine"
)

var logger = logrus.New()

func WithFieldsMsg(fields map[string]interface{}, msg interface{}) {
	line, functionName := 0, "???"
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
	}

	fields[sFunctionName] = functionName
	fields[sFunctionLine] = line

	logger.WithFields(fields).Info(msg)
}
