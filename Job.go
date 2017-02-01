package GoJob

import (
	"sync"
)

/**
Job allow us to process our task with workers. It is thread safe. Task will be added to Job struct by using FIFO approach.
Moreover worker are stored in queue. When workers are idle, they will be kept in queue.
Every added task is stored in @{TaskMap}
 */
type Job struct {
	TaskQueue *Queue;
	WorkerQueue *Queue;
	workers chan *Worker;
	WorkerCounter int;
	Name string;
	shouldStop AtomicBool;
	areTasksRunning AtomicBool;
	TaskMap map[int]*Task;
	TaskMapMutex sync.RWMutex;
}

var descriptionForAddingTaskMessage = "%d task is added to %s job.";

/**
It creates new Job. @{Name} should be unique and it distinguish from other key pairs. @{WorkerCount} determines number of workers.
Workers will be created with NewJob method and is kept in queue.
 */
func NewJob(Name string,WorkerCount int) *Job  {
	job := &Job{Name:Name,WorkerCounter:WorkerCount,TaskQueue:NewQueue(),WorkerQueue:NewQueue(),workers:make(chan *Worker,WorkerCount), TaskMap:make(map[int]*Task)};
	for i := 0; i<WorkerCount;i++  {
		job.WorkerQueue.EnQueue(NewWorker(nil,i,job.workers));
	}
	return job;
}

/**
Starts worker to do their given task. If there is no task goroutine will be in infinite loop
 */
func (job *Job) Start() bool {
	if !job.areTasksRunning.getBoolValue() {
		job.areTasksRunning.setBoolValue(true);
		job.shouldStop.setBoolValue(false);
		go func() {
			for  {
				if !job.shouldStop.getBoolValue() {
					job.continueToWork();
				}else{
					break;
				}
			}
			job.areTasksRunning.setBoolValue(false);
		}();
		return true;
	}
	return false;
}


func (job * Job) continueToWork()  {
	select {
	case worker:= <-job.workers:
		job.WorkerQueue.EnQueue(worker);
	default:
		if !job.TaskQueue.IsEmpty() && !job.WorkerQueue.IsEmpty()  {
			task,isTask := job.TaskQueue.DeQueue().(*Task);
			if isTask {
				if task.IsCancelled() {
					job.removeTaskFromMapWithSync(task);
				}else{
					worker,isWorker := job.WorkerQueue.DeQueue().(*Worker);
					if isWorker {
						worker.Task = task;
						worker.StartWorking()
					}

				}
			}
		}
	}
}

/**
Stops queued tasks process, if workers are in process it wont be stopped until it finishes the task.
 */
func (job * Job) Stop()  {
	job.shouldStop.setBoolValue(true);
}

/**
New Task adds task to queue. It will enqueue and added to task map the given @{task}
 */
func (job * Job) NewTask(task *Task)  {
	job.addTaskToMap(task);
	job.TaskQueue.EnQueue(task);
	logger.Infof(descriptionForAddingTaskMessage,task.Id,job.Name);
}

/**
Adds given @{task} to @{job.TaskMap}
This method is thread safe.
 */
func (job * Job) addTaskToMap(task * Task) {
	job.TaskMapMutex.Lock();
	defer job.TaskMapMutex.Unlock();
	job.TaskMap[task.Id] = task;

}

/**
Removes given @{task} from @{job.TaskMap}
 */
func (job * Job) removeTaskFromMap(task * Task)  {
	delete(job.TaskMap,task.Id);
}

/**
Removes given @{task} from @{job.TaskMap}
This method is thread safe.
 */
func (job * Job) removeTaskFromMapWithSync(task * Task) {
	job.TaskMapMutex.Lock();
	defer job.TaskMapMutex.Unlock();
	job.removeTaskFromMap(task);
}

/**
This method cancel all tasks. Running Task's @{Cancelled} property becomes true. Queued tasks will be removed.
When this method is called, Start method needs to be called to run new incoming tasks.
 */
func (job * Job) CancelAllTasks()  {
	job.Stop();
	job.TaskQueue.Clear();
	job.TaskMapMutex.Lock()
	defer job.TaskMapMutex.Unlock();
	for _,value := range job.TaskMap{
		value.Cancel();
		job.removeTaskFromMap(value);
	}
}

/**
Checks, is task with given id exists on @{job.TaskMap}
 */
func (job * Job) isTaskExist(id int) bool  {
	return job.TaskMap[id] != nil;
}

/**
Checks, is task with given id exists on @{job.TaskMap}
This method is thread safe
 */
func (job * Job) IsTaskExistWithSync(id int) bool  {
	job.TaskMapMutex.Lock();
	defer job.TaskMapMutex.Unlock();
	return job.isTaskExist(id);
}

/**
Cancels task with given @{id}
 */
func (job * Job) CancelTaskWith(id int)  {
	job.TaskMapMutex.Lock();
	defer job.TaskMapMutex.Unlock();
	if job.isTaskExist(id) {
		task := job.TaskMap[id];
		task.Cancel();
		job.removeTaskFromMap(task);
	}
}


