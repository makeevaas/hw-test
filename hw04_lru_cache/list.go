package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	FrontState *ListItem
	BackState  *ListItem
	counter    int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.counter
}

func (l *list) PushFront(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.FrontState == nil {
		l.FrontState = lItem
		l.BackState = lItem
	} else {
		lItem.Next = l.FrontState
		l.FrontState.Prev = lItem
		l.FrontState = lItem
	}
	l.counter++
	return lItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.FrontState == nil {
		l.FrontState = lItem
		l.BackState = lItem
	} else {
		l.BackState.Next = lItem
		lItem.Prev = l.BackState
		l.BackState = lItem
	}
	l.counter++
	return lItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	l.counter--
}

func (l *list) Front() *ListItem {
	return l.FrontState
}

func (l *list) Back() *ListItem {
	return l.BackState
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.FrontState {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.BackState {
		l.BackState = i.Prev
	}

	i.Prev = nil
	i.Next = l.FrontState
	l.FrontState.Prev = i
	l.FrontState = i
}
