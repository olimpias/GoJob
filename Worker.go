package GoJob

import (
	"fmt"
)

type WorkerInterface interface {
	Job();
}

type Worker struct {
	Id int;
	Task WorkerInterface;
	doneTaskCount int;
	workerChannel chan *Worker;
}

var workerStringInterfaceMessage string = "Worker Id:%d, Task Mem Address:%v, Total DONE TASK:%d";

func NewWorker(task WorkerInterface,Id int,workerChannel chan *Worker) *Worker  {
	return &Worker{Task:task,Id:Id,doneTaskCount:0,workerChannel:workerChannel};
}

func (worker * Worker) StartWorking() {
	if worker.Task != nil {
		go func() {
			worker.Task.Job();
			worker.doneTaskCount++;
			worker.workerChannel <- worker;
		}();
	}
}

func (worker * Worker) String() string  {
	return fmt.Sprintf(workerStringInterfaceMessage,worker.Id,worker.Task,worker.doneTaskCount);
}
