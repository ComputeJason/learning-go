package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func Test_Countdown(t *testing.T){

	t.Run("test_print_behaviour_&_sleep_behaviour", func(t *testing.T){
		buf := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buf, spySleeper)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("got %d, want %d", spySleeper.Calls, 3)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
	
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
	
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}