package mq

import (
	"errors"
	"sync"
)

type Broker struct {
	lock  sync.Mutex
	chans []chan Msg
}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) Send(m Msg) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	for _, ch := range b.chans {
		select {
		case ch <- m:
		default:
			return errors.New("send err")
		}
	}
	return nil
}

func (b *Broker) Sub(maxCap int) (<-chan Msg, error) {
	channel := make(chan Msg, maxCap)
	b.lock.Lock()
	defer b.lock.Unlock()
	b.chans = append(b.chans, channel)
	return channel, nil
}

type Msg struct {
	Content string
}
