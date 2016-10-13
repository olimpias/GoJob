package GoJob

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue();
	if queue.head != nil || queue.tail != nil {
		t.Error("head and tail must be nil");
	}
	if queue.count != 0 {
		t.Error("Count value must be 0");
	}
}

func TestQueueEnQueue(t *testing.T) {
	queue := NewQueue();
	value1 := 10;
	value2 := 20;
	queue.EnQueue(value1);
	if queue.head.value != value1 {
		t.Error("Enqeue not working");
	}
	queue.EnQueue(value2);
	if queue.head.value != value2 {
		t.Error("Enqeue not working");
	}
}

func TestQueueCount(t *testing.T) {
	queue := NewQueue();
	value1 := 10;
	value2 := 20;
	queue.EnQueue(value1);
	queue.EnQueue(value2);
	if queue.Count() != 2 {
		t.Error("Count is not matching with enqueue");
	}
	queue.DeQueue();
	if queue.Count() != 1 {
		t.Error("Count is not matching with enqueue and dequeue");
	}
}

func TestQueueDeQueue(t *testing.T) {
	queue := NewQueue();
	value1 := 10;
	value2 := 20;
	queue.EnQueue(value1);
	queue.EnQueue(value2);
	value11 := queue.DeQueue();
	value22 := queue.DeQueue();
	if value1 != value11 {
		t.Error("Dequeu is not matching with first enqueued value");
	}
	if value2 != value22 {
		t.Error("Dequeu is not matching with second enqueued value");
	}
	value33 := queue.DeQueue();
	if value33 != nil {
		t.Error("Dequeue is third value must be nil");
	}
}
