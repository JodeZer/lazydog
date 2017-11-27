package inject

import "testing"

func TestTrace(t *testing.T) {
	__traceStack()
}

// func TestWriteDogHelper(t *testing.T) {
// 	WriteDogHelper("", "heisshei")
// }

func TestGetGid(t *testing.T) {
	t.Log(__curGid())
}
