package mq

import (
	"errors"
	"sync"
)

type Broker struct {
	lock  sync.Mutex
	chans map[string][]chan Msg
}

func NewBroker() *Broker {
	chans := make(map[string][]chan Msg)
	return &Broker{chans: chans}
}

func (b *Broker) Send(m Msg, topic string) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	if chs, ok := b.chans[topic]; ok {
		for _, ch := range chs {
			select {
			case ch <- m:
			default:
				return errors.New("send err")
			}
		}
	}
	return nil
}

func (b *Broker) Sub(maxCap int, topic string) (<-chan Msg, error) {
	channel := make(chan Msg, maxCap)
	b.lock.Lock()
	defer b.lock.Unlock()
	b.chans[topic] = append(b.chans[topic], channel)
	return channel, nil
}

type Msg struct {
	Content string
}
