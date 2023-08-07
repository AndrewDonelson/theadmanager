# ThreadManager

Simple concurrency manager for Go.

## Concurrency in Go

Go is designed with concurrency in mind. Concurrency refers to executing multiple tasks simultaneously by dividing them into separate processes that can run independently. Here's an overview of concurrency in Go and how to use it effectively.

## Goroutines 

The basic unit of concurrency in Go is the goroutine. A goroutine is a lightweight thread that runs concurrently with other goroutines in the same address space. To start a goroutine, simply put the `go` keyword in front of a function call:

```go
go doSomething()
```

This will execute doSomething() concurrently in a new goroutine. The main code will continue executing without waiting for the goroutine to finish.

## Channels
Goroutines need a way to communicate with each other. This is done through channels. A channel is a pipe that passes a value from one goroutine to another. For example:

```go
ch := make(chan int)

go func() {
  ch <- 42 // Send 42 to the channel
}()

i := <-ch // Receive from the channel
```

This passes the integer 42 from one goroutine to the main goroutine through the channel ch.

Channels are typed, so you can only pass a certain type of value through a particular channel. By default sends and receives block until both sides of the channel are ready. This allows goroutines to synchronize without explicit locks or condition variables.

## WaitGroups

The sync package provides a WaitGroup type that allows you to wait for a collection of goroutines to finish executing. This is useful when you spawn several concurrent tasks and want to wait for them all to complete before moving on.

For example:
```go
var wg sync.WaitGroup

for _, item := range items {
  wg.Add(1) // Increment wait group counter

  go func(item) {
    defer wg.Done() // Decrement counter when goroutine exits

    // Process item

  }(item)
}

wg.Wait() // Block until counter reaches zero
```

The WaitGroup allows you to wait for concurrent goroutines in a simple way.

## Mutexes

To synchronize access to shared memory in Go (such as a global variable), you can use a mutex from the sync package. This will lock access to the shared variable so only one goroutine at a time can read or write to it.

For example:
```go
var count int
var mutex sync.Mutex

func Increment() {
  mutex.Lock() 
  count++
  mutex.Unlock()
}
```

This ensures count is incremented atomically without conflicts between goroutines.

## When to Use Concurrency

Here are some good opportunities to introduce concurrency in your Go programs:

- Network/IO-bound tasks - launching a goroutine to handle slow IO in the background
- Parallelizing computation - divide work across CPU cores
- Handling requests concurrently - launch a goroutine per request

Overusing concurrency can add unnecessary complexity so only reach for it when you have an appropriate use case. Overall it's a powerful tool for building efficient and scalable applications in Go.
