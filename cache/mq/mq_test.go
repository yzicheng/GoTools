package mq

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBroker_Send(t *testing.T) {
	b := NewBroker()
	// 发送消息
	for i := 0; i < 3; i++ {
		go func() {
			for {
				time.Sleep(time.Millisecond)
				err := b.Send(Msg{Content: "xxxx"})
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
		go func() {
			defer wg.Done()
			msgs, err := b.Sub(100)
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
