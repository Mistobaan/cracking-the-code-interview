package chapter3

import (
	"errors"
)

var EmptyQueueError error = errors.New("Empty Queue")

// Queue using a Circular Buffer
type Queue struct {
	first int
	last  int
	m     []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		first: 0,
		last:  0,
		m:     make([]interface{}, 1000),
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
	copy(new_m, q.m[q.first:])
	copy(new_m[len(q.m)-q.first:], q.m[:q.first])
	q.first = 0
	q.last = (len(q.m) - 1) // exclusive
	q.m = new_m
}

func (q *Queue) Push(item interface{}) {
	if q.last+1 == len(q.m) && q.first == 0 { // Array full
		q.expand()
	}

	if q.last+1 == q.first {
		q.expand2()
	}

	if (q.last+1) == len(q.m) && q.first > 0 {
		q.last = 0
	} else {
		q.last += 1
	}
	q.m[q.last] = item
}

func (q *Queue) Pop() (interface{}, error) {
	if q.first == q.last {
		return nil, EmptyQueueError
	}

	item := q.m[q.first]

	q.first += 1

	if q.first == len(q.m) {
		q.first = 0
	}

	return item, nil
}
