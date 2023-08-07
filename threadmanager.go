package main

import (
	"log"
	"sync"
	"time"
)

// Thread struct
type Thread struct {
	Name   string
	Fn     func(params ...interface{})
	Logger *log.Logger
	Timing int
	quit   chan bool
	calls  int64
}

// ThreadManager
type ThreadManager struct {
	threads map[string]*Thread
	wg      sync.WaitGroup
	log     *log.Logger
}

// New
func New() *ThreadManager {
	return &ThreadManager{
		threads: make(map[string]*Thread),
		log:     log.Default(),
	}
}

// AddThread
func (tm *ThreadManager) Add(name string, fn func(params ...interface{}), timing int) {

	t := &Thread{
		Name:   name,
		Fn:     fn,
		Timing: timing, // Set timing first
		Logger: tm.log,
	}

	tm.threads[name] = t

	// Prepend self param
	wrapperFn := func(params ...interface{}) {
		allParams := append([]interface{}{t}, params...)
		fn(allParams...)
	}

	t.Fn = wrapperFn

	tm.log.Println("Added thread:", name)
}

// StartThread
func (tm *ThreadManager) Start(name string, params ...interface{}) {

	t := tm.threads[name] // Timing already set

	t.quit = make(chan bool)
	tm.wg.Add(1)
	go tm.run(t, params...) // Start goroutine
	tm.log.Println("Started thread:", name)
}

// StopThread
func (tm *ThreadManager) Stop(name string) {
	t := tm.threads[name]
	quit := make(chan bool)
	go func() {
		<-quit
		t.quit <- true
	}()
	quit <- true
	tm.wg.Wait()
	delete(tm.threads, name)
	tm.log.Println("Stopped thread:", name, "with calls:", t.calls)
}

// runThread
func (tm *ThreadManager) run(t *Thread, params ...interface{}) {

	defer tm.wg.Done()

	for {

		select {
		case <-t.quit:
			return

		default:
			t.calls++

			if t.Timing == 0 {
				// Immediately invoke function if timing is 0
				t.Fn(params...)
			} else {
				t.Fn(params...)
				time.Sleep(time.Duration(t.Timing) * time.Microsecond)
			}
		}
	}
}

// Log
func (t *Thread) Log(msg string) {
	t.Logger.Println(msg)
}
