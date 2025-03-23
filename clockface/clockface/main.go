package main

import (
	"os"
	"time" 
    "github.com/jdra000/learn-go-with-tests/clockface/svg"
)
func main() {
    t := time.Now()
    svg.SVGWriter(os.Stdout, t)
}
