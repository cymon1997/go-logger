package utils

import (
	"path/filepath"
	"runtime"
)

func FuncCaller(skip int) (fileName, fnName string, lineNum int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		file = filepath.Base(file)
		fName := filepath.Base(runtime.FuncForPC(pc).Name())
		return file, fName, line
	}
	return "", "", 0
}
