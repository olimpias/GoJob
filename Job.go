package GoJob

import (
	"sync"
)

/**
Job allow us to process our task with workers. It is thread safe. Task will be added to Job struct by using FIFO approach.
Moreover worker will be stored in queue. When workers are idle, will be kept in queue.
 */
type Job struct {
	TaskQueue *Queue;
	WorkerQueue *Queue;
	workers chan *Worker;
	WorkerCounter int;
	Name string;
	mutexWorkers sync.Mutex;
	shouldStop AtomicBool;
}


/**
It creates new Job. @{Name} should be unique and it distinguish from other key pairs. @{WorkerCount} determines number of workers.
Workers will be created with NewJob method and is kept in queue.
 */
func NewJob(Name string,WorkerCount int) *Job  {
	job := &Job{Name:Name,WorkerCounter:WorkerCount,TaskQueue:NewQueue(),WorkerQueue:NewQueue(),workers:make(chan *Worker,WorkerCount)};
	for i := 0; i<WorkerCount;i++  {
		job.WorkerQueue.EnQueue(NewWorker(nil,i,job.workers));
	}
	return job;
}

/**
Starts worker to do their given task. If there is no task goroutine will be in infinite loop

 */
func (job *Job) Start()  {
	job.shouldStop.setBoolValue(false);
	go func() {
		for  {
			if !job.shouldStop.getBoolValue() {
				job.continueToWork();
			}else{
				break;
			}
		}
	}();
}

func (job * Job) continueToWork()  {
	select {
	case worker:= <-job.workers:
		job.WorkerQueue.EnQueue(worker);
	default:
		if !job.TaskQueue.IsEmpty() && !job.WorkerQueue.IsEmpty()  {
			worker,isWorker := job.WorkerQueue.DeQueue().(*Worker);
			task,isTask := job.TaskQueue.DeQueue().(WorkerInterface);
			if isTask && isWorker {
				worker.Task = task;
				worker.StartWorking()
			}
		}
	}
}

/**
Stops jobs, if worker is in process it wont be stopped until it finishes the task.
 */
func (job * Job) Stop()  {
	job.shouldStop.setBoolValue(true);
}

/**
New Task adds task to queue. It will enqueue the given @{task}
 */
func (job * Job) NewTask(task WorkerInterface)  {
	job.TaskQueue.EnQueue(task);
}


