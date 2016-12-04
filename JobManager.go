package GoJob

import (
	"sync"
	"errors"
)

/**
!!IMPORTANT INFORMATION!!
Every job must have unique identifier string.
Jobs are containing tasks. It keeps them in queue.
Every job has its own goroutine. Also it creates goroutine for each worker. Worker count can not be changed when it is set.

JobManager saves jobs with given string values. It seperates jobs with given string.
Careful while adding new job with string. Remember, if you already added given string, past created job will be lost on map.
 */
type JobManager struct {
	jobMap map[string]*Job;
	mutex sync.RWMutex;
}


var sharedJobManager *JobManager;
var once sync.Once;

//Error
var JobIsAlreadyStarted = errors.New("Job is already started");
var JobNameDoesNotExist = errors.New("Job name does not exist");

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
 If job name does not exist, it returns @{JobNameDoesNotExist} error
 */
func (jobManager * JobManager) AddTask(Name string,task WorkerInterface) error  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	job := jobManager.jobMap[Name];
	if job != nil{
		 job.NewTask(task);
		return nil;
	}
	return JobNameDoesNotExist;
}

/**
Starts job with given name. If job name does not exist it returns @{JobNameDoesNotExist} error.
If job is already started, it returns @{JobIsAlreadyStarted} error
 */
func (jobManager * JobManager) StartTasks(Name string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[Name] != nil {
		if jobManager.jobMap[Name].isTaskRunning.getBoolValue() {
			return JobIsAlreadyStarted;
		}
		jobManager.jobMap[Name].Start();
		return nil;
	}
	return JobNameDoesNotExist;
}

/**
TODO users should able to stop running tasks!
Stops tasks with given job name however if task is already in running progress, it can not be stopped. If it is already stopped, nothing is going to be happen.
If job name does not exist, it returns @{JobNameDoesNotExist} error.
 */
func (jobManager * JobManager) StopTasks(Name string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[Name] != nil {
		jobManager.jobMap[Name].Stop();
		return nil;
	}
	return JobNameDoesNotExist;

}

