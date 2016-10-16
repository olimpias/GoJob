package GoJob

import (
	"sync"
	"testing"
	"time"
)

type TestJobData struct {
	Value int;
}
var mutex sync.Mutex;
var counter = 0;

func (testData * TestJobData)  Job(){
	mutex.Lock();
	defer mutex.Unlock();
	counter += testData.Value;
}
func NewTestJobData(value int) *TestJobData  {
	return &TestJobData{Value:value};
}



func TestJobStart(t *testing.T)  {
	totalValueCount:= 0;
	job := NewJob("newtest",3);

	if job.WorkerQueue.count != 3 {
		t.Error("Error Occured on worker queue");
	}
	for i := 0; i < 10; i++ {
		totalValueCount += i;
		job.NewTask(NewTestJobData(i));
	}
	job.Start();
	time.Sleep(time.Second * 3)

	if counter != totalValueCount {
		t.Error("Workers are not working properly");
	}
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
	if doneTaskCount != 10 {
		t.Error("All task Are not done!");
	}
	job.Stop();
	if !job.shouldStop.getBoolValue() {
		t.Error("Should be false");
	}
}