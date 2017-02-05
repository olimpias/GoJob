package GoJob

import (
	"sync"
)

/**
Executor is an interface which hosts @{Job} method. Job methods allows user to run their codes on different goroutines concurrently.
JobManager can cancel tasks. If user wants to stop running goroutine, when task is cancelled, I suggest that task.Cancelled should be used by user.

Example:
func (test * Test) Job(task * Task) {
	.....
	task.IsCancelled(){
		return;
	}
	....
}
 */
type Executor interface {
	Job(task * Task);
}

/**
Task is a container for Executor
Id unifies tasks. Every task has unique @{Id}.
@{executor} is interface which needs to be implemented by user.
@{Cancelled} checks -> Is Given @{executor} cancelled.
 */
type Task struct {
	Id int
	executor Executor;
	cancelled bool;
	cancelMutex sync.RWMutex;
}

var IdCounter = 0;
var counterMutex sync.Mutex;

var taskCreationInfoMessage = "Task is created with %d id. Executor mem: %p ";
/**
Creates New Task pointer struct with given @{executor}. New struct will be assigned with unique Id.
 */
func NewTask(executor Executor) *Task  {
	counterMutex.Lock();
	defer counterMutex.Unlock();
	task := &Task{Id:IdCounter,executor:executor};
	IdCounter++;
	logger.Infof(taskCreationInfoMessage,task.Id,executor);
	return task;
}
/**
Checks is task cancelled
 */
func (task * Task) IsCancelled() bool  {
	task.cancelMutex.RLock();
	defer task.cancelMutex.RUnlock();
	return task.cancelled;
}

/**
Changes @{Cancelled} value to false. Cancelled will be used as argument in Job method. @{Cancelled} needs to be checked inside Job method.
Also it will effect queued task. When this method used, cancelled task will be bypassed.
 */
func (task * Task) Cancel()   {
	task.cancelMutex.Lock();
	defer  task.cancelMutex.Unlock();
	task.cancelled = true;
}
