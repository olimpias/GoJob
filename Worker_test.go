package GoJob

import (
	"testing"
	"fmt"
)

func TestNewWorker(t *testing.T) {
	testData := NewTestJobData(10);
	task := NewTask(testData);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(task,10,workerChannel);
	if worker.Id != 10 || worker.workerChannel != workerChannel || worker.Task.executor != testData {
		t.Error("NewWorker couldn't be created properly");
	}
}

func TestWorkerStartWorking(t *testing.T) {
	testData := NewTestJobData(10);
	task := NewTask(testData);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(task,10,workerChannel);
	worker.StartWorking();
	tmpWorker := <-workerChannel;
	if tmpWorker != worker {
		t.Error("StartWorking not working properly");
	}
}

func TestWorkerString(t *testing.T) {
	testData := NewTestJobData(10);
	task := NewTask(testData);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(task,10,workerChannel);
	worker.StartWorking();
	tmpWorker := <-workerChannel;
	if tmpWorker.String() != fmt.Sprintf(workerStringInterfaceMessage,worker.Id,worker.doneTaskCount) {
		t.Error("String implemented incorrectly");
	}

}
