package broadcast

import (
	"errors"
	"sync"
)

type Broadcast[T any] struct {
	sync.Mutex
	channels map[string]*Channel[T]
}

func NewBroadcast[T any]() *Broadcast[T] {
	return &Broadcast[T]{
		channels: map[string]*Channel[T]{},
	}
}

func (b *Broadcast[T]) AddChannel(channelID string) (ch *Channel[T], err error) {
	b.Lock()
	defer b.Unlock()

	var ok bool
	if ch, ok = b.channels[channelID]; ok && ch != nil {
		err = errors.New("channel_exists")
		return
	}

	ch = newChannel[T]()
	b.channels[channelID] = ch

	return
}
