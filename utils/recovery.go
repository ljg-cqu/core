package utils

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
)

// recoverPanic tries to handle panics.
func recoverPanic(logger *log.Logger) {
	if r := recover(); r != nil {
		// record an error on the job with panic message and stacktrace
		stackBuf := make([]byte, 1024)
		n := runtime.Stack(stackBuf, false)

		buf := &bytes.Buffer{}
		fmt.Fprintf(buf, "%v\n", r)
		fmt.Fprintln(buf, string(stackBuf[:n]))
		fmt.Fprintln(buf, "[...]")
		stacktrace := buf.String()

		logger.Printf("Job panicked, stacktrace:%+v", stacktrace)
	}
}
