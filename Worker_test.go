package GoJob

import (
	"testing"
	"fmt"
)

type TestData struct {
	Value int;
}

func (testData * TestData)  Job(){
	fmt.Printf("Job value is %d\n",testData.Value);
}
func NewTestData(value int) *TestData  {
	return &TestData{Value:value};
}

func TestNewWorker(t *testing.T) {
	testData := NewTestData(10);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(testData,10,workerChannel);
	if worker.Id != 10 || worker.workerChannel != workerChannel || worker.Task != testData {
		t.Error("NewWorker couldn't be created properly");
	}
}

func TestWorkerStartWorking(t *testing.T) {
	testData := NewTestData(10);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(testData,10,workerChannel);
	worker.StartWorking();
	tmpWorker := <-workerChannel;
	if tmpWorker != worker {
		t.Error("StartWorking not working properly");
	}
}

func TestWorkerString(t *testing.T) {
	testData := NewTestData(10);
	workerChannel := make(chan *Worker,1);
	worker := NewWorker(testData,10,workerChannel);
	worker.StartWorking();
	tmpWorker := <-workerChannel;
	if tmpWorker.String() != fmt.Sprintf(workerStringInterfaceMessage,worker.Id,worker.Task,worker.doneTaskCount) {
		t.Error("String implemented in correctly");
	}
}
