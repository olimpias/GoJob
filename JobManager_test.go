package GoJob

import "testing"

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
	SingletonJobManager().addTask(name,testData);
	if job == nil || job.TaskQueue.count != 1{
		t.Error("Add Task is not working properly");
	}
}


