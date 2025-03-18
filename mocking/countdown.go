package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "GO!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

// * Real sleeper
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// * Fake sleeper that checks sleep and writting
// * To evaluate sleep it appends "sleep" every time it's Sleep() method is called
// * To evaluate writting it appends "write" every time it's customized Write() method it's called
const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	// * Here we call the real sleeper. In the tests, the fake one.
	// * The real sleeper Sleep() method implements the 1 sec delay.
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
