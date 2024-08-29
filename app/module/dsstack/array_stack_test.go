package dsstack

import (
	"testing"
)

func TestStack(t *testing.T) {
	d := NewArrayStack()
	d.push(1)
	d.push(2)
	d.push(3)
	d.push(4)
	d.push(5)
	d.push(6)

	d.peek()
	d.pop()
	d.peek()
	d.pop()
	d.peek()
	d.isEmpty()
	d.size()
	d.toSlice()
}
