# Advanced Golang Concurrency – Patterns and Best Practices

A comprehensive guide to advanced concurrency patterns and best practices in Go with practical examples.

<!-- TOC -->
* [Advanced Golang Concurrency – Patterns and Best Practices](#advanced-golang-concurrency--patterns-and-best-practices)
  * [Basic Sync Primitives](#basic-sync-primitives)
    * [Mutex vs RWMutex](#mutex-vs-rwmutex)
    * [Semaphore vs Worker Pool](#semaphore-vs-worker-pool)
  * [Advanced Sync Primitives](#advanced-sync-primitives)
    * [Futex vs Mutex](#futex-vs-mutex)
    * [When to Use Each Sync Primitive](#when-to-use-each-sync-primitive)
    * [Spin Lock vs Ticket Lock](#spin-lock-vs-ticket-lock)
  * [Practical Concurrency Constructs](#practical-concurrency-constructs)
    * [Context Package](#context-package)
    * [Channels and Select](#channels-and-select)
  * [Concurrency Patterns](#concurrency-patterns)
    * [Core Patterns](#core-patterns)
  * [Avoiding Common Pitfalls](#avoiding-common-pitfalls)
    * [Deadlocks](#deadlocks)
    * [Starvation and Livelock](#starvation-and-livelock)
    * [Race Conditions](#race-conditions)
    * [Goroutine Leaks](#goroutine-leaks)
  * [Data Structures for Concurrency](#data-structures-for-concurrency)
  * [What if not Goroutines? – Other practices to consider](#what-if-not-goroutines--other-practices-to-consider)
    * [Lock Free Data Structures](#lock-free-data-structures)
    * [False Sharing and Cache Coherence](#false-sharing-and-cache-coherence)
  * [Practical Examples](#practical-examples)
<!-- TOC -->

## Basic Sync Primitives

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

## Advanced Sync Primitives

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

## Practical Concurrency Constructs

### Context Package

1. **Context**: Carry deadlines, cancellation signals, and request-scoped values.
   - **Cancel**: Propagate cancellation signals to goroutines. 

### Channels and Select

1. **Channels**: Typed conduits for communication between goroutines.
2. **Select**: Wait on multiple channel operations.
    - **Select Timeout**: Implement timeouts in `select` statements to avoid blocking indefinitely. `time.After` can be used to create a timeout channel.

## Concurrency Patterns

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


## Avoiding Common Pitfalls

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

## Data Structures for Concurrency

1. **Ring Buffer**: Circular buffer for fixed-size data.
2. **Concurrent Map**: Map with concurrent access support.

## What if not Goroutines? – Other practices to consider

### Lock Free Data Structures


### False Sharing and Cache Coherence

## Practical Examples

1. **Rate Limiting**
2. **Debouncing**
3. **Throttling**
4. **Resource Pooling**

