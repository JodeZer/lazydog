package main

import "runtime"
import "fmt"

func __traceStack() {
	caller, file, line := __caller()
	fmt.Printf("%s %s %d\n", caller, file, line)
}

func __caller() (string, string, int) {

	fpcs := make([]uintptr, 1)

	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "", "n/a", 0
	}

	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "", "n/a", 0
	}

	file, line := fun.FileLine(0)

	return fun.Name(), file, line
}
