package GoJob

import "sync"

type Node struct {
	value interface{};
	next *Node;
}

type Queue struct {
	head *Node;
	tail *Node;
	count int;
	mutex sync.RWMutex;
}

func NewQueue() *Queue  {
	return &Queue{count:0};
}

func (queue Queue) Count() int  {
	queue.mutex.RLock();
	defer queue.mutex.RUnlock();
	return queue.count;
}

func (queue Queue) IsEmpty() bool  {
	return queue.head == nil && queue.tail == nil;
}

func (queue * Queue) EnQueue(value interface{})  {
	queue.mutex.Lock();
	defer queue.mutex.Unlock();
	node := &Node{value:value};
	if queue.IsEmpty() {
		queue.tail = node;
	}else{
		queue.head.next = node;
	}
	queue.head = node;
	queue.count ++;
}

func (queue * Queue) DeQueue() interface{}  {
	queue.mutex.Lock();
	defer queue.mutex.Unlock();
	tail := queue.tail;
	if tail != nil {
		queue.tail = tail.next;
	}else{
		queue.head =  nil;
	}
	queue.count --;
	if tail != nil{
		return tail.value;
	}
	return nil;
}




