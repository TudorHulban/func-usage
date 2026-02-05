package funcusage

import (
	"fmt"
	"runtime"
)

func TraceExit() {
	pc, _, line, couldGetInfo := runtime.Caller(1)
	if couldGetInfo {
		fmt.Printf(
			"exiting function %s at line %d.\n",

			runtime.FuncForPC(pc).Name(),
			line,
		)
	}
}
