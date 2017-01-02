package GoJob

import (
	"testing"
)

func TestSingletonJobManager(t *testing.T) {
	jobManager := SingletonJobManager();
	if sharedJobManager == nil || jobManager != sharedJobManager {
		 t.Error("SingletonJobManager function is failed");
	}
}

func TestNewJob(t *testing.T) {
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	job := SingletonJobManager().GetJob(name);
	if job == nil || job.WorkerQueue.count != 10{
		 t.Error("New Job not working properly");
	}
}

func TestAddTask(t *testing.T)  {
	testData := NewTestJobData(10);
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	job := SingletonJobManager().GetJob(name);
	SingletonJobManager().AddTask(name,testData);
	if job == nil || job.TaskQueue.count != 1{
		t.Error("Add Task is not working properly");
	}
	name = "notExistJob";
	_,err := SingletonJobManager().AddTask(name,testData);
	if err != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestStartTasks(t *testing.T) {
	testData := NewTestJobData(10)
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	SingletonJobManager().AddTask(name,testData);
	SingletonJobManager().StartTasks(name);
	job := SingletonJobManager().GetJob(name);
	if !job.areTasksRunning.getBoolValue() {
		t.Error("Job has to be running");
	}
	if SingletonJobManager().StartTasks(name) != JobIsAlreadyStarted {
		t.Error("Job should already be running");
	}
	SingletonJobManager().CancelAllTasks(name);
	name = "notExistJob"
	if SingletonJobManager().StartTasks(name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestStopRunningTasks(t *testing.T){
	testData := NewTestJobData(10);
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	SingletonJobManager().AddTask(name,testData);
	SingletonJobManager().StartTasks(name);
	SingletonJobManager().StopRunningTasks(name)
	job := SingletonJobManager().GetJob(name);
	if !job.shouldStop.getBoolValue() {
		t.Error("Job has to be stopped");
	}
	if !job.areTasksRunning.getBoolValue() {
		t.Error("Job has to be stopped running tasks");
	}
	name = "notExistJob"
	if SingletonJobManager().StopRunningTasks(name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestCancelAllTasks(t *testing.T) {
	name := "test.job";
	SingletonJobManager().NewJob(name,2);
	for i:= 0;i<1000;i++ {
		testData := NewTestJobData(i);
		SingletonJobManager().AddTask(name,testData);
	}
	SingletonJobManager().StartTasks(name);
	SingletonJobManager().CancelAllTasks(name);

	name = "notExistJob";
	if SingletonJobManager().CancelAllTasks(name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestCancelTaskWithId(t *testing.T) {
	name := "test.job";
	SingletonJobManager().NewJob(name,2);
	length := 1000;
	var taskId int = -1;
	for i:= 0;i<length;i++ {
		testData := NewTestJobData(i);
		taskId,_ = SingletonJobManager().AddTask(name,testData);
	}
	if taskId == -1 {
		 t.Error("Task must be assigned");
	}
	SingletonJobManager().StartTasks(name);
	if SingletonJobManager().CancelTask(taskId,name) != nil{
		t.Error("Job should be cancelled the task");
	}

	if SingletonJobManager().CancelTask(999,name) != TaskIdDoesNotExist{
		t.Error("Task id should not be existed");
	}

	name = "notExistJob";
	if SingletonJobManager().CancelTask(length - 1,name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}


