package main

func main() {

	__traceStack()
	foo3()
}

func foo() {
	__traceStack()
}

func foo2() {
	__traceStack()

	foo()
}

func foo3() {
	__traceStack()

	foo2()
}
