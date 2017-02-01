# GoJob

GoJob is background task manager. It allows you to create different group of jobs and to run them in queue structure with given worker count concurrently.

## Installation

```bash
go get github.com/olimpias/gojob
```

## Example

Executor interface is used for running requested tasks.

```go
type Executor interface {
	Job(task * Task);
}
```

In this example, Counter struct is incrementing atomicValue variable under job method. Counter struct implements  Job(task * Task) method which is included by Executor interface.

```go
type Counter struct {
	atomicValue uint64;
}

func (counter * Counter) increment(){
	atomic.AddUint64(&counter.atomicValue,1);
}

// Executor interface implementation for Counter struct
func (counter * Counter) Job(task * Task){
	//If you are going to use cancel operation for tasks. Don't forget to implement task.IsCancelled() check.
	//Also remember to use this in loops in checking is task Cancelled.
	if task.isCancelled() {
		return;
	}
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

### Version 0.3

You will able to cancel tasks with version 0.3.

```go
   taskId := goJobManager.AddTask(jobName,newCounterPtr); // will return taskID
```

By using taskID, you can cancel task.

```go
   goJobManager.CancelTask(taskId,jobName) // Cancels running or queued task.
   //Do not forget to use @{task.isCancelled()} inside Executor interface method.!!
```

```go
   goJobManager.CancelAllTasks(jobName) // Cancel all running and queued tasks.
   //Do not forget to use @{task.isCancelled()} inside Executor interface method.!!
```


## TODO

Priority Queue Structure Implementation for GoJob<br>
Complex Example by using GoJob<br>
