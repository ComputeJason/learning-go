package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// this chapter is pretty good at explaining something comlicated simply: MOCKING
// please relook it , & also try it yourself
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

// essentially we are mocking this for 2 reasons
// the dependencies are hard observe eg. time.Sleep & printing to Stdout, so we mock these
// in real life, mock when real operations are expensive/difficult/slow or have real side effects

const finalWord = "Go!"
const countdownStart = 3

func main(){
	Countdown(os.Stdout, &ConfigurableSleeper{2 * time.Second, time.Sleep})
}

func Countdown(output io.Writer, sleeper Sleeper){

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(output,i)
		sleeper.Sleep()
	} 
	fmt.Fprint(output, finalWord)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const sleep,write string= "sleep","write"

func (s *SpyCountdownOperations) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}



type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}