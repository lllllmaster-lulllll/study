package queue

import (
	"container/list"
)

type linkedListDeque struct {
	data *list.List
}

func NewLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

func (l *linkedListDeque) pushFirst(val any) {
	l.data.PushFront(val)
}

func (l *linkedListDeque) pushLast(val any) {
	l.data.PushBack(val)
}

func (l *linkedListDeque) popFirst() any {
	if l.isEmpty() {
		return nil
	}
	val := l.data.Front()
	l.data.Remove(val)
	return val.Value
}

func (l *linkedListDeque) isEmpty() bool {
	return l.data.Len() == 0
}

func (l *linkedListDeque) popLast() any {
	if l.isEmpty() {
		return nil
	}
	val := l.data.Back()
	l.data.Remove(val)
	return val.Value
}

func (l *linkedListDeque) peekFirst() any {
	if l.isEmpty() {
		return nil
	}
	val := l.data.Front()
	return val.Value
}

func (l *linkedListDeque) peekLast() any {
	if l.isEmpty() {
		return nil
	}
	return l.data.Back().Value
}

func (l *linkedListDeque) size() int {
	return l.data.Len()
}

func (l *linkedListDeque) tolist() *list.List {
	return l.data
}
