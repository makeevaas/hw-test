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
	Key   string
}

type list struct {
	front   *ListItem
	back    *ListItem
	counter int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.counter
}

func (l *list) PushFront(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.front == nil {
		l.front = lItem
		l.back = lItem
	} else {
		lItem.Next = l.front
		l.front.Prev = lItem
		l.front = lItem
	}
	l.counter++
	return lItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	lItem := &ListItem{Value: v}
	if l.front == nil {
		l.front = lItem
		l.back = lItem
	} else {
		l.back.Next = lItem
		lItem.Prev = l.back
		l.back = lItem
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
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.back {
		l.back = i.Prev
	}

	i.Prev = nil
	i.Next = l.front
	l.front.Prev = i
	l.front = i
}
