package mq

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBroker_Send(t *testing.T) {
	b := NewBroker()
	topics := []string{"topic0", "topic1", "topic2"}
	// 发送消息
	for i := 0; i < 3; i++ {
		topic := topics[i]
		go func() {
			for {
				time.Sleep(time.Millisecond)
				err := b.Send(Msg{Content: topic}, topic)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	}

	// 订阅消息
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("消费者%v", i)
		topic := topics[i]
		go func() {
			defer wg.Done()
			msgs, err := b.Sub(100, topic)
			if err != nil {
				fmt.Println(err)
				return
			}
			for msg := range msgs {
				fmt.Println(name, msg.Content)
			}
		}()
	}
	wg.Wait()
}
