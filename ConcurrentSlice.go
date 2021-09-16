package main

import (
	"fmt"
	"sync"
)


type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
}


type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}


func NewConcurrentSlice() *ConcurrentSlice {
	cs := &ConcurrentSlice{
		items: make([]interface{}, 0),
	}

	return cs
}


func (cs *ConcurrentSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()
	cs.items = append(cs.items, item)
}


func (cs *ConcurrentSlice) Get(index int) (item interface{}) {
	cs.RLock()
	defer cs.RUnlock()
	if isset(cs.items, index) {
		return cs.items[index]
	}
	return nil
}

func isset(arr []interface{}, index int) bool {
	return (len(arr) > index)
}


func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)
	f := func() {
		cs.RLock()
		defer cs.RUnlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

func main() {
	s := NewConcurrentSlice()

	s.Append(1)
	s.Append(2)
	s.Append(3)

	ch := s.Iter()

	for item := range ch {
		fmt.Printf("Item: %v\n", item.Value)
	}
}

