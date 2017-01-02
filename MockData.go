package GoJob

import "sync"

type TestJobData struct {
	Value int;
	shouldIncrement bool
}
var mutex sync.Mutex;
var counter = 0;

func (testData * TestJobData)  Job(task * Task){
	if task.IsCancelled() {
		return;
	}
	mutex.Lock();
	defer mutex.Unlock();
	if testData.shouldIncrement {
		counter += testData.Value;
	}

}
func NewTestJobData(value int) *TestJobData  {
	return &TestJobData{Value:value, shouldIncrement:false};
}
