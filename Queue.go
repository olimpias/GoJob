package GoJob

import "sync"

/**
Thread safe generic queue implementation
 */

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

func (queue *  Queue) Count() int  {
	queue.mutex.RLock();
	defer queue.mutex.RUnlock();
	return queue.count;
}

func (queue Queue) isEmpty() bool  {
	return queue.head == nil && queue.tail == nil;
}

func (queue * Queue) IsEmpty() bool  {
	queue.mutex.RLock();
	defer queue.mutex.RUnlock();
	return queue.isEmpty();
}

func (queue * Queue) EnQueue(value interface{})  {
	queue.mutex.Lock();
	defer queue.mutex.Unlock();
	node := &Node{value:value};
	if queue.isEmpty() {
		queue.head = node;
	}else{
		queue.tail.next = node;
	}
	queue.tail = node;
	queue.count ++;
}

func (queue * Queue) DeQueue() interface{}  {
	queue.mutex.Lock();
	defer queue.mutex.Unlock();
	if queue.isEmpty() {
		return nil;
	}else{
		queue.count--;
		head := queue.head;
		if head == queue.tail{
			queue.head = nil;
			queue.tail = nil;
		}else{
			queue.head = head.next;
		}
		return head.value;
	}
}

func (queue * Queue) Clear(){
	for !queue.IsEmpty() {
		queue.DeQueue();
	}
}




