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
	FirstState *ListItem
	LastState  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	var counter int
	firstEl := l.FirstState
	for firstEl != nil {
		firstEl = firstEl.Next
		counter++
	}
	return counter
}

func (l *list) PushFront(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.FirstState == nil {
		l.FirstState = lItem
		l.LastState = lItem
	} else {
		lItem.Next = l.FirstState
		l.FirstState.Prev = lItem
		l.FirstState = lItem
	}
	return lItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.FirstState == nil {
		l.FirstState = lItem
		l.LastState = lItem
	} else {
		l.LastState.Next = lItem
		lItem.Prev = l.LastState
		l.LastState = lItem
	}
	return lItem
}

func (l *list) Remove(i *ListItem) {
	firstEl := l.FirstState
	for firstEl != nil {
		firstEl = firstEl.Next
		if firstEl == i {
			if firstEl.Prev != nil {
				firstEl.Prev.Next = firstEl.Next
			} else {
				l.FirstState = firstEl.Next
			}
			if firstEl.Next != nil {
				firstEl.Next.Prev = firstEl.Prev
			} else {
				l.LastState = firstEl.Prev
			}
			break
		}
	}
}

func (l *list) Front() *ListItem {
	return l.FirstState
}

func (l *list) Back() *ListItem {
	return l.LastState
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.FirstState == nil || i == l.FirstState {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.LastState {
		l.LastState = i.Prev
	}

	i.Prev = nil
	i.Next = l.FirstState
	l.FirstState.Prev = i
	l.FirstState = i
}
