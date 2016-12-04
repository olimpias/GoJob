# GoJob

GoJob is background task manager. It allows to create different group of jobs and to run them in queue structure with given worker count concurrently. 

## Installation

```bash
go get github.com/olimpias/gojob
```

## Example

WorkerInterface help us to use goJobManager by implementing in needed struct as method.

```go
type WorkerInterface interface {
	Job();
}
```

In this example, Counter struct is incrementing atomicValue variable under job method. Counter uses increment method inside Job() interface which is bounded to GoJob

```go
type Counter struct {
	atomicValue uint64;
}

func (counter * Counter) increment(){
	atomic.AddUint64(&counter.atomicValue,1);
}

// WorkerInterface implementation for GoJobManager!
func (counter * Counter) Job(){
	counter.increment();
}
//Creates new Counter pointer struct
func newCounter() *Counter  {
	return &Counter{atomicValue:0};
}
```


```go
   //new Counter object
   newCounterPtr := newCounter();
   jobName := "incrementNumber";
   var goJobManager = GoJob.SingletonJobManager();
   taskName := "incrementNumber";
   //Creates new job with given job name and allows 4 workers to run concurrently
   goJobManager.NewJob(taskName,4);
   //Add new task to job
   goJobManager.AddTask(jobName,newCounterPtr); // newCounterPtr
   //Start task with given jobName
   goJobManager.StartTasks(jobName);
   
```

For more details, please check Counter example under Example folder.

## TODO

Cancel option specific tasks also for running tasks or queued tasks<br>
Debug Logging<br>
Priority Queue Structure Implementation for GoJob<br>
More Complex Example by using GoJob<br>





