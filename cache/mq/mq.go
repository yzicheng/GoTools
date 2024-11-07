package mq

import (
	"errors"
	"fmt"
	"sync"
)

var broker *Broker
var once sync.Once

type Broker struct {
	lock  sync.Mutex
	chans map[string][]chan Msg
}

func NewBroker() *Broker {
	once.Do(func() {
		chans := make(map[string][]chan Msg)
		broker = &Broker{
			chans: chans,
		}
		fmt.Println("create chans")
	})
	return broker
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
