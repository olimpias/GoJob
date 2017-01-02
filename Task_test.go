package GoJob

import (
	"testing"
)


func TestNewTask(t *testing.T)  {

	task := NewTestJobData(10);

	taskS1 := NewTask(task);
	taskS2 := NewTask(task);
	if taskS1.Id == taskS2.Id {
		t.Error("Ids can not be equal");
	}
	if taskS1.IsCancelled() && taskS2.IsCancelled() {
		t.Error("must not start with cancel true");
	}
	if taskS1.executor != taskS2.executor {
		t.Error("must point same address");
	}
}

func TestCancel(t *testing.T)  {
	taskS1 := NewTask(nil);
	if taskS1.IsCancelled() {
		t.Error("cancelled property must be false");
	}
	taskS1.Cancel();
	if !taskS1.IsCancelled() {
		t.Error("cancelled property must be true");
	}
}
