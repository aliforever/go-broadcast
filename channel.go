package broadcast

import "sync"

type Channel[T any] struct {
	sync.Mutex
	listeners []chan T
}

func newChannel[T any]() *Channel[T] {
	return &Channel[T]{}
}

func (c *Channel[T]) AddListener() chan T {
	c.Lock()
	defer c.Unlock()

	ch := make(chan T)
	c.listeners = append(c.listeners, ch)
	return ch
}

func (c *Channel[T]) InformListeners(data T) {
	c.Lock()
	defer c.Unlock()

	for _, listener := range c.listeners {
		go func(listener chan T) {
			listener <- data
		}(listener)
	}
}
