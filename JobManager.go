package GoJob

import "sync"

type JobManager struct {
	jobMap map[string]*Job;
	mutex sync.RWMutex;
}


var sharedJobManager *JobManager;
var once sync.Once;

func SingletonJobManager() *JobManager  {
	once.Do(func() {
		sharedJobManager = &JobManager{jobMap:make(map[string]*Job)};
	});
	return sharedJobManager;
}

func (jobManager * JobManager) NewJob(Name string,WorkerCount int)  {
	jobManager.mutex.Lock();
	defer jobManager.mutex.Unlock();
	job := NewJob(Name,WorkerCount);
	jobManager.jobMap[Name] = job;
}

func (jobManager * JobManager) GetJob(Name string) *Job  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	return jobManager.jobMap[Name];
}

func (jobManager * JobManager) addTask(Name string,task WorkerInterface)  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	job := jobManager.jobMap[Name];
	if job != nil{
		 job.NewTask(task);
	}
}

