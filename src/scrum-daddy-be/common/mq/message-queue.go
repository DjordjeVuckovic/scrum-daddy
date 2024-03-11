package mq

import "sync"

type MessageBroker interface {
	Publish(topic string, payload []byte) error
	Consume(topic string, handler func(payload []byte)) error
	Close()
}

type InMemoryMessageBroker[T any] struct {
	channel chan T
	close   chan bool
	wg      sync.WaitGroup
}

func NewMessageQueue[T any](capacity int) *InMemoryMessageBroker[T] {
	return &InMemoryMessageBroker[T]{
		channel: make(chan T, capacity),
		close:   make(chan bool),
	}
}

func (mq *InMemoryMessageBroker[T]) Publish(message T) {
	mq.wg.Add(1)
	mq.channel <- message
}
