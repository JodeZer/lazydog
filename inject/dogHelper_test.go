package inject

import "testing"

func TestTrace(t *testing.T) {
	__traceStack()
}

func TestWriteDogHelper(t *testing.T) {
	WriteDogHelper("", "heisshei")
}
