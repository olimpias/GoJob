package GoJob

import (
	"fmt"
)
/**
WorkerInterface allows struct to use GoJob efficiently. Tasks should be implemented in Job().
 */
type WorkerInterface interface {
	Job();
}

/**
Worker is used for how many Task is done by worker, how efficient is worker working and maintains the tasks.
Worker Id is unique for every Job --> @{Id}.
Task are passed with WorkerInterface --> @{Task}.
When task is done, Job queue will be unblocked with @{workerChannel}
@{doneTaskCount} keeps finished task count.
 */
type Worker struct {
	Id int;
	Task WorkerInterface;
	doneTaskCount int;
	workerChannel chan *Worker;
}

var workerStringInterfaceMessage string = "Worker Id:%d, Task Mem Address:%v, Total DONE TASK:%d";

/**
Creates New worker with task, id and workerChannel
Id and worker channel should be nil and id should be unique between job's workers.
 */
func NewWorker(task WorkerInterface,Id int,workerChannel chan *Worker) *Worker  {
	return &Worker{Task:task,Id:Id,doneTaskCount:0,workerChannel:workerChannel};
}

/**
Worker starts the task on different goroutine. When it is done, increases doneTaskCount and send itself to workerChannel.
 */
func (worker * Worker) StartWorking() {
	if worker.Task != nil {
		go func() {
			worker.Task.Job();
			worker.doneTaskCount++;
			worker.Task = nil;
			worker.workerChannel <- worker;
		}();
	}
}

/**
Returns current task pointer, Worker id and done Task Count values
 */
func (worker * Worker) String() string  {
	return fmt.Sprintf(workerStringInterfaceMessage,worker.Id,worker.Task,worker.doneTaskCount);
}
