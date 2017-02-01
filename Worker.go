package GoJob

import (
	"fmt"
)

/**
Worker is used maintaining the tasks. It needs to done its task and notify channel.
Worker Id is unique for every worker in Job --> @{Id}.
Task are passed with Executor --> @{Task}.
When task is done, Job queue will be unblocked with @{workerChannel}
@{doneTaskCount} counts finished tasks.
 */
type Worker struct {
	Id int;
	Task *Task;
	doneTaskCount int;
	workerChannel chan *Worker;
}

var workerStringInterfaceMessage string = "Worker Id:%d, Task Mem Address:%p Task Id:%d, Total DONE TASK:%d";
var taskRunStatusMessage = "Task id:%d is %s";

/**
Creates New worker with task, id and workerChannel
Id and worker channel should be nil and id should be unique between job's workers.
 */
func NewWorker(task *Task,Id int,workerChannel chan *Worker) *Worker  {
	return &Worker{Task:task,Id:Id,doneTaskCount:0,workerChannel:workerChannel};
}

/**
Worker handles given task on new goroutine. When it is done, increases doneTaskCount and send itself to workerChannel.
 */
func (worker * Worker) StartWorking() {
	task := worker.Task;
	if task != nil {
		go func() {
			logger.Debug(worker.String());
			logger.Infof(taskRunStatusMessage,worker.Task.Id,"starting");
			worker.Task.executor.Job(task);
			logger.Infof(taskRunStatusMessage,worker.Task.Id,"ended");
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
	return fmt.Sprintf(workerStringInterfaceMessage,worker.Id,worker.Task.executor,worker.Task.Id,worker.doneTaskCount);
}
