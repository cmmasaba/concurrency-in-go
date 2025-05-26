# Concurrency Patterns in Go

Possible definitions of concurrency:
- "Execution happens in some non-deterministic order"
- "Non-sequential execution"
- "Parts of a program execute out of order or in partial order"

The main thing to note is that concurrency is the composition of **independently** executing computations.<br>
<br>Concurrency is not the same as parallelism.<br>
- Parallelism is a type of computing in which many operations are carried out simultaneously. This typically involves using multiple processors.
- Concurrency is the ability of multiple tasks to appear to run simultaneously on a single processor, through rapid switching
- Concurrency doesn't make a program faster, parallelism does.

## Goroutines
A `goroutine` is an independently executing function managed by the Go runtime. It has its own stack which grows or shrinks as required.<br>
<br>
A goroutine **is not** a thread, instead a goroutine is multiplexed dynamically onto threads as needed to keep all goroutines running.<br>
<br>A goroutine is started using the command below:

```go
go f(x, y)
```
starts a new goroutine running function `f(x, y)`.<br>
<br>The evaluation of f(x, y) happens in the current goroutine while the execution of f(x, y) happens in a new goroutine.<br>
<br>Goroutines run in the same address space so access to shared memory must be synchronized. This is done using the `sync` package.

## Communication Between Goroutines
Since goroutines use the same address space and access to shared memory, they need a way communicate with each other. In Go this is made possible through `channels`.<br>
Channels are a typed medium through which you can send or receive data using the `<-` operator. The data flows in the direction of the <- operator.
```go
// Declare and initialize a channel of ints
var c chan int
c = make(chan int)

// Alternative and brief declaration and initialization
// c := make(chan int)

c <- 100            // Send a value to the channel
fmt.Println(<- c)   // Receivea value from the channel
```
By default, send and receive operations on a channel block until the other side is ready. Therefore, apart from communication, channels enable synchronization between goroutines.<br>
*[See example](/channels/main.go)*

#### Buffered Channels
Channels can be created with a buffer, but this removes the synchronization bit. The buffer length is provided as the second argument to the make command.<br>
```go
    ch := make(chan int, 100)  // A buffered channel that can hold 100 ints
```
Sends to a buffered channel block only when the buffer is full. Receives block only when the buffer is empty. Buffered channels should be used sparingly.<br>
*[See example](/channels/buffered.go)*


## Patterns
*Don't communicate by sharing memory, share memory by communicating.*<br>

### 1. Range-and-Close
A sender can `close` a channel to indicate no more values will be sent. A receiver can test whether a channel has been closed using the syntax below:
```go
if v, ok := <-ch; !ok{
    // ok is false if the channel is closed
}
```
- Only the sender should close the channel, never the receiver.
- Channels don't need to be closed. You only close them if the receiver must be told no more values are incoming.

<br>The range-and-close pattern allows you to loop over a channel until it is closed.<br>
*[See example](/channels/fibonacci.go)*

### 2. Select
The `select` statement lets a goroutine wait on multiple communication operations.<br>
<br>A select blocks until one of its cases can run, then it executes that case. It chooses one at randome if more than one cases are ready.<br>
The default case is run if no other case is ready.<br>
*[See example](/channels/select.go)*

### 3. Generator
This can be a function that returns a channel. This is possible because channels are first-class objects, just like integers etc.<br>
The channel returned can be used as a handle to a service.
*[See example](/generators-handles/main.go)*

### 4. Fan-in
Fan-in measures the number of components or dependencies that point towards a specific component.<br>
<br>Multiple goroutines send information to the same channel which aggregates and sends to a unified stream.<br>
*[See example](/fan-in/main.go)*

### 5. Mutual Exclusion
When goroutines are accessing the same underlying storage location we have to make sure only one goroutine can<br>
access the location at a time to avoid conflicts. This particular kind of conflict is called a `race condition`.<br>
```
A race condition or race hazard is the condition of an electronics, software, or other system where the system's 
substantive behavior is dependent on the sequence or timing of other uncontrollable events, leading to unexpected
or inconsistent results

A race condition can arise in software when a computer program has multiple code paths that are executing at the same
time. If the multiple code paths take a different amount of time than expected, they can finish in a different order
than expected, which can cause software bugs due to unanticipated behavior.
```
<br>This concept of controlling access to a shared variable is called mutual exclusion. In Go it is provided by the<br>
`sync.Mutex` data structure which has the following 2 methods: `Lock` and `Unlock`.<br>

<br>We can define a block of code to be executed in mutual exclusion by sorrounding it with a call to Lock and Unlock.<br>
Using locks can lead to another kind of problem called `deadlock`.


## Further Reading
1. Go Concurrency Patterns. [link](https://go.dev/talks/2012/concurrency.slide#1)
2. Go Concurrency Patterns: Pipelines and Cancellation. [link](https://go.dev/blog/pipelines)
3. Advanced Go Concurrency Patterns. [link](https://go.dev/talks/2013/advconc.slide#1)
4. Share Memory By Communicating. [link](https://go.dev/doc/codewalk/sharemem/)
4. Race Condition. [link](https://en.wikipedia.org/wiki/Race_condition#)
5. Mutual Exclusion. [link](https://en.wikipedia.org/wiki/Mutual_exclusion)
6. Deadlock. [link]()