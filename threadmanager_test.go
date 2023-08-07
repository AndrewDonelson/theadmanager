package main

import (
	"testing"
	"time"
)

// Thread function
func doWork(params ...interface{}) {
	self := params[0].(*Thread) // First param is self

	// Can access self.Name etc
	self.Logger.Println("Doing work in thread:", self.Name, "Calls:", self.calls)

	// params[1:] will have original params
}

func TestTiming(t *testing.T) {
	tm := New()
	start := time.Now()

	//tm.Add("thread1", func(params ...interface{}) {}, 1500)
	tm.Add("thread1", doWork, 1500)
	tm.Start("thread1")

	time.Sleep(2 * time.Second)
	tm.Stop("thread1")

	elapsed := time.Since(start)
	if elapsed < 2*time.Second {
		t.Error("expected delay from timing")
	}
}

func TestNoTiming(t *testing.T) {
	tm := New()
	start := time.Now()

	tm.Add("thread1", doWork, 0)
	tm.Start("thread1")

	time.Sleep(1 * time.Second)
	tm.Stop("thread1")

	margin := 1 * time.Second / 1000000
	elapsed := time.Since(start) / 1000000
	if elapsed > margin {
		t.Error("no delay expected with 0 timing")
	}
}

// Existing tests...

func TestAddThreadTiming(t *testing.T) {
	tm := New()
	tm.Add("thread1", doWork, 5)
	tm.Start("thread1")

	time.Sleep(5 * time.Millisecond)

	if tm.threads["thread1"].calls != 13 {
		t.Error("calls not incremented correctly")
	}

	tm.Stop("thread1")

}

func TestMultipleThreadsTiming(t *testing.T) {

	tm := New()

	// Thread with 100 microsecond delay
	tm.Add("thread1", doWork, 100)

	// Thread with 250 microsecond delay
	tm.Add("thread2", doWork, 250)

	//  with 500 microsecond delay
	tm.Add("thread3", doWork, 500)

	tm.Start("thread1")
	tm.Start("thread2")
	tm.Start("thread3")

	// Sleep for 1 millisecond
	time.Sleep(1 * time.Millisecond)

	tm.Stop("thread1")
	tm.Stop("thread2")
	tm.Stop("thread3")

}
