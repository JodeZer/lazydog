package inject // replacable
import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var gopaths []string

func init() {
	gopath := os.Getenv("GOPATH")
	gopaths = strings.Split(gopath, ";")
	for i := range gopaths {
		gopaths[i] += "/src/"
	}
}

// must only depend on standard
//  dont edit this file
func __traceStack() {
	caller, file, line := __caller()
	fmt.Printf("[lazydog] %s:%d caller= %s  \n", __prettyFile(file), line, __prettyCaller(caller))
}

// from https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang
func __caller() (string, string, int) {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "", "n/a", 0 // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "", "n/a", 0
	}

	file, line := fun.FileLine(0)
	// return its name
	return fun.Name(), file, line
}

func __prettyCaller(caller string) string {
	return string(string(caller[strings.LastIndex(caller, ".")+1:]))
}

func __prettyFile(file string) string {
	for _, gopath := range gopaths {
		if strings.Contains(file, gopath) {
			return strings.Replace(file, gopath, "", -1)
		}
	}
	return file
}
