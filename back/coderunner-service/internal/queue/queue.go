package queue

import (
	models "coderunner-service/pkg/models"
)

type (
	Queue struct {
		front  *node
		rear   *node
		length int
	}
	node struct {
		counter int
		value   *models.Block
		next    *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Return the number of items in the queue
func (q *Queue) Len() int {
	return q.length
}

// View the front item in the queue
func (q *Queue) Peek() (*models.Block, int) {
	if q.length == 0 {
		return nil, 0
	}
	return q.front.value, q.front.counter
}

// Dequeue removes and returns the front item from the queue
func (q *Queue) Dequeue() (*models.Block, int) {
	if q.length == 0 {
		return nil, 0
	}

	n := q.front
	if q.front == q.rear {
		q.rear = nil
	}
	q.front = n.next
	q.length--

	return n.value, n.counter
}

// Enqueue adds a value to the rear of the queue
func (q *Queue) Enqueue(value *models.Block, counter int) {
	n := &node{counter, value, nil}
	if q.length == 0 {
		q.front = n
		q.rear = n
	} else {
		q.rear.next = n
		q.rear = n
	}
	q.length++
}
