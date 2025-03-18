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

// * Spy	 sleeper that checks sleep and writting
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

// * Spy and Real sleeper at the same time. Its sleep time is customizable.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// * Spy struct to use inside ConfigurableSleeper{sleep: spyTime.Sleep}
// * Our spy struct does not sleep but changes its durationSlept to the time.Duration value
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	// * Here we call the real sleeper. In the tests, the fake one.
	// * The real sleeper Sleep() method implements the 1 sec delay. time.Sleep(1 * time.Second)
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
