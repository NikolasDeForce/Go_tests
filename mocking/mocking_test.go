package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyTime struct {
	durationSleep time.Duration
}

type SpySleepper struct {
	Calls int
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleepper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepCounter := &SpyCountOperations{}
		Countdown(spySleepCounter, spySleepCounter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepCounter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepCounter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSleep != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSleep)
	}
}

func (s *SpySleepper) Sleep() {
	s.Calls++
}

type SpyCountOperations struct {
	Calls []string
}

func (c *SpyCountOperations) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *SpyCountOperations) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSleep = duration
}

const write = "write"
const sleep = "sleep"
