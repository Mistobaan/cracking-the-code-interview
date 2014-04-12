package chapter3

import (
	"errors"
)

var EmptyQueueError error = errors.New("Empty Queue")

// Queue using a Circular Buffer
type Queue struct {
	head int
	tail int
	m    []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		head: 0,
		tail: 0,
		m:    make([]interface{}, 1000),
	}
}

func (q *Queue) expand() {
	// Expand strategy ...
	new_m := make([]interface{}, len(q.m)*2)
	copy(new_m, q.m)
	q.m = new_m
}

func (q *Queue) expand2() {
	new_m := make([]interface{}, len(q.m)*2)
	copy(new_m, q.m[q.head:])
	copy(new_m[len(q.m)-q.head:], q.m[:q.head])
	q.head = 0
	q.tail = (len(q.m) - 1) // exclusive
	q.m = new_m
}

func (q *Queue) Push(item interface{}) {
	if q.tail+1 == len(q.m) && q.head == 0 { // Array full
		q.expand()
	}

	if q.tail+1 == q.head {
		q.expand2()
	}

	if (q.tail+1) == len(q.m) && q.head > 0 {
		q.tail = 0
	} else {
		q.tail += 1
	}
	q.m[q.tail] = item
}

func (q *Queue) Pop() (interface{}, error) {
	if q.head == q.tail {
		return nil, EmptyQueueError
	}

	item := q.m[q.head]

	q.head += 1

	if q.head == len(q.m) {
		q.head = 0
	}

	return item, nil
}
