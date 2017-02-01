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
var TaskIdDoesNotExist = errors.New("Task Id does not exist");

var IdNotFound = -1;

/**
JobManager is once existed . It should be called with this method everytime.
 */
func SingletonJobManager() *JobManager  {
	once.Do(func() {
		sharedJobManager = &JobManager{jobMap:make(map[string]*Job)};
	});
	return sharedJobManager;
}

/**
New job added with given string and worker count.
 */
func (jobManager * JobManager) NewJob(jobName string,workerCount int)  {
	jobManager.mutex.Lock();
	defer jobManager.mutex.Unlock();
	job := NewJob(jobName,workerCount);
	jobManager.jobMap[jobName] = job;
}

/**
if you already added input string, job pointer will be received with given string. If you don't it will return nil.
 */
func (jobManager * JobManager) GetJob(jobName string) *Job  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	return jobManager.jobMap[jobName];
}

/**
addTask allows you to add new task to given job name. The task will be assigned to available worker in job.
 WorkerInterface must be implemented for input @{task}
 If job name does not exist, it returns @{JobNameDoesNotExist} error
 */
func (jobManager * JobManager) AddTask(jobName string,executor Executor) (int,error)  {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	task := NewTask(executor);
	job := jobManager.jobMap[jobName];
	if job != nil{
		 job.NewTask(task)
		return task.Id,nil;
	}
	return IdNotFound,JobNameDoesNotExist;
}

/**
Starts job with given name. If job name does not exist it returns @{JobNameDoesNotExist} error.
If job is already started, it returns @{JobIsAlreadyStarted} error
 */
func (jobManager * JobManager) StartTasks(jobName string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[jobName] != nil {
		if jobManager.jobMap[jobName].areTasksRunning.getBoolValue() {
			return JobIsAlreadyStarted;
		}
		jobManager.jobMap[jobName].Start();
		return nil;
	}
	return JobNameDoesNotExist;
}

/**
Stops running jobs with given job name however if task is already in running progress, it can not be stopped. If it is already stopped, nothing is going to be happen.
If job name does not exist, it returns @{JobNameDoesNotExist} error.
 */
func (jobManager * JobManager) StopRunningTasks(jobName string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[jobName] != nil {
		jobManager.jobMap[jobName].Stop()
		return nil;
	}
	return JobNameDoesNotExist;
}

/**

 */
func (jobManager * JobManager) CancelAllTasks(jobName string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[jobName] != nil {
		jobManager.jobMap[jobName].CancelAllTasks()
		return nil;
	}
	return JobNameDoesNotExist;
}

/**

 */
func (jobManager * JobManager) CancelTask(taskId int,jobName string) error {
	jobManager.mutex.RLock();
	defer jobManager.mutex.RUnlock();
	if jobManager.jobMap[jobName] != nil {
		if jobManager.jobMap[jobName].IsTaskExistWithSync(taskId) {
			jobManager.jobMap[jobName].CancelTaskWith(taskId);
			logger.Debugf("Task with id is cancelled succesfully",taskId);
			return nil;
		}
		return TaskIdDoesNotExist;
	}
	return JobNameDoesNotExist;
}

func (jobManager * JobManager)  DebugConfiguration(isActive bool) {
	logger.DEBUG_ENABLED = isActive;
}

func (jobManager * JobManager) InfoConfiguration(isActive bool)  {
	logger.INFORMATION_ENABLED = isActive;
}

func (jobManager * JobManager) DisableLogging() {
	logger.INFORMATION_ENABLED = false;
	logger.DEBUG_ENABLED = false;
}

func (jobManager * JobManager) ActivateLogging() {
	logger.DEBUG_ENABLED = true;
	logger.INFORMATION_ENABLED = true;
}




