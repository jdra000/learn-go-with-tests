package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	// * There is a difference between Fprintf and Printf
	// ! func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	// !	p := newPrinter()
	// !	p.doPrintf(format, a)
	// !	n, err = w.Write(p.buf)
	// !	p.free()
	// !	return

	// !  func Printf(format string, a ...any) (n int, err error) {
	// !	return Fprintf(os.Stdout, format, a...)
	// * See how Printf writes to standard output
	// * Printf passes os.Stdout to Fprintf which expects an io.writer
	// * We want to modify the buffer, so we use Fprintf instead
	fmt.Fprintf(writer, "Hello, %s", name)
}
func main() {
	Greet(os.Stdout, "Eloide")
}
