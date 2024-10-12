package net

import "sync"

type NetRes int

type Window struct {
	ch       chan NetRes
	size     int
	winQueue []NetRes
	lock     sync.RWMutex
}

func NewWindow(size int, queue chan NetRes, winQueue []NetRes) *Window {
	win := &Window{
		ch:       queue,
		size:     size,
		winQueue: winQueue,
		lock:     sync.RWMutex{},
	}
	return win
}

func (win *Window) Start() error {
	for {
		select {
		case <-win.ch:
			for i := win.size; i < len(win.winQueue); i++ {

			}
		}
	}
}

func (win *Window) Send(res NetRes) error {
	win.lock.Lock()
	win.lock.Unlock()
	win.winQueue = append(win.winQueue, res)
	go func() {
		win.ch <- res
		return
	}()
	return nil
}
