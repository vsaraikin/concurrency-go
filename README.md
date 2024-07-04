# Advanced Golang Concurrency – Patterns and Best Practices

A comprehensive guide to advanced concurrency patterns and best practices in Go with practical examples.

## Table of contents

[//]: # (TBD:)

[//]: # ()
[//]: # (1. Restarter)

[//]: # (2. Pause, Graceful Shutdown, Cancellation)

[//]: # (3. Future)

[//]: # (4. Flushing consumers before shutdown)

[//]: # (5. )

## 1. Basic Sync Primitives

1. **Mutex (sync.Mutex)**: Mutual exclusion lock to prevent race conditions.
2. **RWMutex (sync.RWMutex)**: A reader/writer mutual exclusion lock.
3. **WaitGroup (sync.WaitGroup)**: Wait for a collection of goroutines to finish.
4. **Cond (sync.Cond)**: Conditional variable for signaling between goroutines.
5. **Once (sync.Once)**: Ensures a function is only called once.
6. **Atomic Operations (sync/atomic)**: Atomic operations for low-level synchronization.
7. **Semaphore**: A signaling mechanism to control access to resources.

### Mutex vs RWMutex

RWMutex is thus preferable for data that is mostly read.

### Semaphore vs Worker Pool

Both patterns are used to control access to resources, but they serve different purposes.

- **Semaphore**: runs each **_new task_** in a **_new separate goroutine_**.
- **Worker Pool**: runs each **_new task_** in a **_pre-allocated goroutine_**.


## 2. Advanced Sync Primitives

1. **Futex & Mutex**: Lightweight user-space locks.
   - **Futex**: Fast user-space mutex.
   - **Mutex**: Kernel-assisted mutex.
2. **Spin Lock**: Busy-wait loop for synchronization.
3. **Ticket Lock**: Fair locking mechanism using ticket numbers.

### Futex vs Mutex

- Futex is a user-space lock that avoids system calls when possible. In Golang, there is no direct access to futexes, but the runtime uses them internally. 
- Mutex is a kernel-assisted lock that uses system calls to block/unblock goroutines.

### When to Use Each Sync Primitive

### Spin Lock vs Ticket Lock


## 3. Concurrency Control

1. **Context**: Carry deadlines, cancellation signals, and request-scoped values.
   - **Cancel**: Propagate cancellation signals to goroutines.
2. **Channels**: Typed conduits for communication between goroutines.
3. **Select**: Wait on multiple channel operations.
    - **Select Timeout**: Implement timeouts in `select` statements to avoid blocking indefinitely. `time.After` can be used to create a timeout channel.

## 4. Concurrency Patterns

### Core Patterns

1. **Fan-in**: Multiple goroutines sending data to a single channel.
2. **Fan-out**: A single goroutine sending data to multiple goroutines.
3. **Pipeline**: Chain of stages where each stage passes data to the next via channels.
4. **Generator**: Function that returns a channel to produce a stream of values.
5. **Queuing**: Managing tasks using a queue to balance workload.
6. **Semaphore**: Control access to a resource by multiple goroutines.
7. **Worker Pool**: Pool of worker goroutines to handle tasks concurrently.
8. **Bridge**: Connect multiple channels together.
9. **Tee Channel**: Split data from one channel into multiple channels.


## 5. Avoiding Common Pitfalls

### Deadlocks

1. **Deadlock**: Situation where goroutines are stuck waiting for each other.
    - **Deadlock Example**: Practical example illustrating a deadlock scenario and how to avoid it.

### Starvation and Livelock

1. **Starvation**: A goroutine is perpetually denied access to resources.
2. **Livelock**: Goroutines are active but unable to make progress.

### Race Conditions

1. **Race Conditions**: Concurrent access to shared resources leading to inconsistent results.
    - **Tools**: Use tools like `go run -race` to detect race conditions.

### Goroutine Leaks

## 6. Data Structures for Concurrency

1. **Ring Buffer**: Circular buffer for fixed-size data.

## 7. Misc Topics

### Lock Free Data Structures

### False Sharing and Cache Coherence

## 8. Practical Examples

1. **Rate Limiting**
2. **Debouncing**
3. **Throttling**
4. **Resource Pooling**

## References

- [Lock Internals - Ravikumar](https://blog.stackademic.com/mutex-internals-in-golang-1624749f35a6)
- [What’s false sharing and how to solve it - Genchi Lu](https://medium.com/@genchilu/whats-false-sharing-and-how-to-solve-it-using-golang-as-example-ef978a305e10)