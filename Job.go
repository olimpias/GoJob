package GoJob

import "sync"

type Job struct {
	TaskQueue *Queue;
	WorkerQueue *Queue;
	workers chan *Worker;
	WorkerCounter int;
	Name string;
	mutexWorkers sync.Mutex;
}

func NewJob(Name string,WorkerCount int) *Job  {
	job := &Job{Name:Name,WorkerCounter:WorkerCount,TaskQueue:NewQueue(),WorkerQueue:NewQueue(),workers:make(chan *Worker,WorkerCount)};
	for i := 0; i<WorkerCount;i++  {
		job.WorkerQueue.EnQueue(NewWorker(nil,i,job.workers));
	}
	return job;
}

func (job *Job) Start()  {
	go func() {
		for  {
			select {
			case worker:= <-job.workers:
				job.WorkerQueue.EnQueue(worker);
			default:
				if !job.TaskQueue.IsEmpty() {
					worker,isWorker := job.WorkerQueue.DeQueue().(*Worker);
					task,isTask := job.TaskQueue.DeQueue().(WorkerInterface);
					if isTask && isWorker {
						worker.Task = task;
					}
				}
			}
		}
	}();
}

func (job * Job) NewTask(task WorkerInterface)  {
	job.TaskQueue.EnQueue(task);
}


