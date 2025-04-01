package sync

import "sync"

type Counter struct {
	count int
	lock sync.Mutex

}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc (){
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++ 
	
}

func (c *Counter) Value() int {
	return c.count
}