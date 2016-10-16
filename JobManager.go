package GoJob

import "sync"

/**
JobManager saves jobs with given string values. It seperates jobs with given string.
Careful while adding new job with string. Remember, if you already added given string, past created job will be lost on map.
 */
type JobManager struct {
	jobMap map[string]*Job;
	mutex sync.RWMutex;
}


var sharedJobManager *JobManager;
var once sync.Once;

/**
JobManager is exists in once. It should be called with this method.
 */
func SingletonJobManager() *JobManager  {
	once.Do(func() {
		sharedJobManager = &JobManager{jobMap:make(map[string]*Job)};
	});
	return sharedJobManager;
}

/**
New job added with given string and worker count value.
 */
func (jobManager * JobManager) NewJob(Name string,WorkerCount int)  {
	jobManager.mutex.Lock();
	defer jobManager.mutex.Unlock();
	job := NewJob(Name,WorkerCount);
	jobManager.jobMap[Name] = job;
}

/**
if you already added input string, job pointer will be received with given string. If you don't it will return nil.
 */
func (jobManager * JobManager) GetJob(Name string) *Job  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	return jobManager.jobMap[Name];
}

/**
addTask allows you to add new task to given job name. The task will be assigned to available worker in job.
 WorkerInterface must be implemented for input @{task}
 */
func (jobManager * JobManager) addTask(Name string,task WorkerInterface)  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	job := jobManager.jobMap[Name];
	if job != nil{
		 job.NewTask(task);
	}
}

