package stack

import (
	models "coderunner-service/pkg/models"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		counter int
		value   *models.Workflow
		prev    *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() (*models.Workflow, int) {
	if this.length == 0 {
		return nil, 0
	}
	return this.top.value, this.top.counter
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() (*models.Workflow, int) {
	if this.length == 0 {
		return nil, 0
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value, n.counter
}

// Push a value onto the top of the stack
func (this *Stack) Push(value *models.Workflow, counter int) {
	n := &node{counter, value, this.top}
	this.top = n
	this.length++
}
