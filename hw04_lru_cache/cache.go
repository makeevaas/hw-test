package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if el, found := l.items[key]; found {
		el.Value = value
		l.queue.MoveToFront(el)
		return true
	}

	if l.queue.Len() >= l.capacity {
		el := l.queue.Back()
		if el != nil {
			l.queue.Remove(el)
			delete(l.items, Key(el.Key))
		}
	}

	el := l.queue.PushFront(value)
	el.Key = string(key)
	l.items[key] = el
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if el, found := l.items[key]; found {
		l.queue.MoveToFront(el)
		return el.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, 0)
}
