# Advanced Golang Concurrency – Patterns and Best Practices

## Sync Primitives and Patterns

### Basic Sync Primitives

1. **Mutex (sync.Mutex)**: Mutual exclusion lock to prevent race conditions.
2. **RWMutex (sync.RWMutex)**: A reader/writer mutual exclusion lock.
3. **WaitGroup (sync.WaitGroup)**: Wait for a collection of goroutines to finish.
4. **Cond (sync.Cond)**: Conditional variable for signaling between goroutines.
5. **Once (sync.Once)**: Ensures a function is only called once.
6. **Atomic Operations (sync/atomic)**: Atomic operations for low-level synchronization.
7. **Semaphore**: A signaling mechanism to control access to resources.

#### Mutex vs RWMutex

RWMutex is thus preferable for data that is mostly read.

#### Semaphore vs Worker Pool

Both patterns are used to control access to resources, but they serve different purposes.

- **Semaphore**: runs each **_new task_** in a **_new separate goroutine_**.
- **Worker Pool**: runs each **_new task_** in a **_pre-allocated goroutine_**.

### Advanced Sync Primitives

1. **Futex & Mutex**: Lightweight user-space locks.
   - **Futex**: Fast user-space mutex.
   - **Mutex**: Kernel-assisted mutex.
2. **Spin Lock**: Busy-wait loop for synchronization.
3. **Ticket Lock**: Fair locking mechanism using ticket numbers.


#### When to Use Each Sync Primitive

#### Futex vs Mutex

#### Spin Lock vs Ticket Lock

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

## Practical Concurrency Constructs

### Channels and Select

1. **Channels**: Typed conduits for communication between goroutines.
2. **Select**: Wait on multiple channel operations.
    - **Select Timeout**: Implement timeouts in `select` statements to avoid blocking indefinitely. `time.After` can be used to create a timeout channel.

### Context Package

1. **Context**: Carry deadlines, cancellation signals, and request-scoped values.
    - **Cancel**: Propagate cancellation signals to goroutines.

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

## Best Practices and Performance Optimization

1. **Profiling**: Use Go's profiling tools to identify and optimize performance bottlenecks.
2. **Goroutine Leaks**: Ensure all started goroutines are properly terminated.
3. **Batching**: Combine multiple operations to reduce the overhead of context switching.
4. **Load Balancing**: Distribute tasks evenly across goroutines to avoid overloading any single goroutine.


## Practical Examples

1. **Rate Limiting**
2. **Debouncing**
3. **Throttling**
4. **Resource Pooling**

