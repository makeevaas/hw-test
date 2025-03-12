package hw04lrucache

import (
	"encoding/json"
	"fmt"
)

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
		l.Clear()
	}

	if l.queue.Len() > l.capacity {
		return false
	}

	el := l.queue.PushFront(value)
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
	if l.queue.Len() == 0 {
		return
	}
	el := l.queue.Back()
	if el != nil {
		l.queue.Remove(el)
		b, err := json.Marshal(el.Value)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		delete(l.items, Key(string(b)))
	}
}
