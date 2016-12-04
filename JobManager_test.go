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
	testData := NewTestData(10);
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	job := SingletonJobManager().GetJob(name);
	SingletonJobManager().AddTask(name,testData);
	if job == nil || job.TaskQueue.count != 1{
		t.Error("Add Task is not working properly");
	}
	name = "notExistJob";
	if SingletonJobManager().AddTask(name,testData) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestStartTasks(t *testing.T) {
	testData := NewTestData(10);
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	SingletonJobManager().AddTask(name,testData);
	SingletonJobManager().StartTasks(name);
	job := SingletonJobManager().GetJob(name);
	if !job.isTaskRunning.getBoolValue() {
		t.Error("Job has to be running");
	}
	if SingletonJobManager().StartTasks(name) != JobIsAlreadyStarted {
		t.Error("Job should already be running");
	}
	SingletonJobManager().StopTasks(name);
	name = "notExistJob"
	if SingletonJobManager().StartTasks(name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}

func TestStopTasks(t *testing.T){
	testData := NewTestData(10);
	name := "test.job";
	SingletonJobManager().NewJob(name,10);
	SingletonJobManager().AddTask(name,testData);
	SingletonJobManager().StopTasks(name);
	job := SingletonJobManager().GetJob(name);
	if job.isTaskRunning.getBoolValue() {
		t.Error("Job has to be stopped");
	}
	name = "notExistJob"
	if SingletonJobManager().StopTasks(name) != JobNameDoesNotExist {
		t.Error("Job shouldnt be exist");
	}
}


