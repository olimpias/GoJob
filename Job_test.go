package GoJob

import (
	"testing"
	"time"
)





func TestJobIsTaskExistWith(t *testing.T) {
	job := NewJob("newtest",3);

	if job.WorkerQueue.count != 3 {
		t.Error("Error Occured on worker queue");
	}
	task := NewTask(NewTestJobData(10));
	job.NewTask(task);

	if !job.isTaskExist(task.Id) {
		t.Error("Task must exist with given id");
	}
	if job.IsTaskExistWithSync(100) {
		t.Error("Task should not exist with given id");
	}
}

func TestJobCancelTaskWith(t *testing.T) {
	job := NewJob("newtest",3);

	if job.WorkerQueue.count != 3 {
		t.Error("Error Occured on worker queue");
	}
	task := NewTask(NewTestJobData(10));
	job.NewTask(task);
	
	job.CancelTaskWith(task.Id);
	_,ok := job.TaskMap[task.Id]
	if ok {
		t.Error("Task should not exist");
	}
	if !task.IsCancelled() {
		t.Error("Task should be cancelled");
	}
}

func TestJobCancelAllTasks(t *testing.T) {
	job := NewJob("newtest",3);

	if job.WorkerQueue.count != 3 {
		t.Error("Error Occured on worker queue");
	}
	task1 := NewTask(NewTestJobData(10));
	job.NewTask(task1);
	task2 := NewTask(NewTestJobData(12));
	job.NewTask(task2);

	job.CancelAllTasks();
	if job.TaskQueue.Count() != 0 {
		t.Error("Task Queue must be empty");
	}

	if !job.shouldStop.getBoolValue() {
		t.Error("Job should be stopped");	
	}

	_,ok := job.TaskMap[task1.Id];
	if ok {
		t.Error("Task1 should not be existed on taskMap");
	}

	_,ok = job.TaskMap[task2.Id];
	if ok {
		t.Error("Task2 should not be existed on taskMap");
	}
}

func TestJobStart(t *testing.T)  {

	totalValueCount:= 0;
	job := NewJob("newtest",3);

	if job.WorkerQueue.count != 3 {
		t.Error("Error Occured on worker queue");
	}
	counter = 0;
	for i := 0; i < 10; i++ {
		totalValueCount += i;
		newTestData := NewTestJobData(i);
		newTestData.shouldIncrement = true;
		task := NewTask(newTestData);
		job.NewTask(task);
	}
	task := NewTask(NewTestJobData(11));
	job.NewTask(task);
	job.Start();
	task.Cancel();
	if job.Start() {
		t.Error("Workers should be running");
	}
	time.Sleep(time.Second * 1)
	mutex.Lock();
	if counter != totalValueCount {
		t.Error("Workers are not working properly");
	}
	mutex.Unlock();
	job.WorkerQueue.mutex.RLock();
	tmp := job.WorkerQueue.head;
	doneTaskCount := 0;
	for tmp != nil {
		worker,ok := tmp.value.(*Worker);
		if ok {
			doneTaskCount += worker.doneTaskCount;
		}else {
			t.Error("Node value is not a worker!");
		}
		tmp = tmp.next;
	}
	job.WorkerQueue.mutex.RUnlock();
	if doneTaskCount != 10 {
		t.Error("All task Are not done!");
	}
	job.Stop();
	if !job.shouldStop.getBoolValue() {
		t.Error("Should be false");
	}
	time.Sleep(time.Second * 1)
}