package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

func main_writeBatches() {
	eventDispatcher := NewEventDispatcher(5)
	go eventDispatcher.Listen(context.TODO())
	// go eventDispatcher.ListenWithBufferedChan(context.TODO())
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	eventDispatcher.AddConsumer("127.0.0.3:8081")
	// }()

	eventDispatcher.AddConsumer("127.0.0.1:8080")
	// eventDispatcher.AddConsumer("127.0.0.2:8080")

	idx := 0
	for {
		idx++
		time.Sleep(2 * time.Second)
		for counter := range 5 {
			str := fmt.Sprintf("Click-%d", idx*counter)
			event := UserEvent{"user-event", []byte(str)}
			fmt.Println("Event created", event)
			go eventDispatcher.SendEvent(event)
		}
	}
}

func sendDataSomewhere(ctx context.Context, addr string, events []UserEvent) error {
	for _, event := range events {
		fmt.Printf("[SLICE] %s sent to %s. Data: %s.\n", event.EventName, addr, string(event.Data))
	}
	return nil
}

func sendDataBuffChanSomewhere(ctx context.Context, addr string, events []UserEvent) error {
	for _, event := range events {
		fmt.Printf("[CHANNEL] %s sent to %s. Data: %s.\n", event.EventName, addr, string(event.Data))
	}
	return nil
}

type UserEvent struct {
	EventName string
	Data      []byte
}

type EventsDispatcher struct {
	writeChan              chan struct{}
	eventQueueChan         chan UserEvent
	eventQueue             []UserEvent
	batchSize              int
	destinations           []string
	simpleListenerActive   bool
	bufferedListenerActive bool
	mu                     *sync.Mutex
}

func NewEventDispatcher(batchSize int) *EventsDispatcher {
	return &EventsDispatcher{
		writeChan:              make(chan struct{}),
		eventQueueChan:         make(chan UserEvent, batchSize),
		eventQueue:             make([]UserEvent, 0),
		batchSize:              batchSize,
		destinations:           make([]string, 0),
		simpleListenerActive:   false,
		bufferedListenerActive: false,
		mu:                     &sync.Mutex{},
	}
}

func (ed *EventsDispatcher) AddConsumer(addr string) {
	ed.destinations = append(ed.destinations, addr)
}

func (ed *EventsDispatcher) SendEvent(event UserEvent) {
	if ed.bufferedListenerActive {
		ed.eventQueueChan <- event
		return
	}

	ed.mu.Lock()
	defer ed.mu.Unlock()
	if ed.simpleListenerActive {
		ed.eventQueue = append(ed.eventQueue, event)
		if len(ed.eventQueue) >= ed.batchSize {
			ed.writeChan <- struct{}{}
		}
	}
}

func (ed *EventsDispatcher) ListenWithBufferedChan(ctx context.Context) error {
	ed.bufferedListenerActive = true
	eventQueue := make([]UserEvent, 0, ed.batchSize)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-ed.eventQueueChan:
			fmt.Println("event__", event)
			eventQueue = append(eventQueue, event)
			if len(eventQueue) >= ed.batchSize {
				for _, addr := range ed.destinations {
					go sendDataBuffChanSomewhere(context.TODO(), addr, eventQueue[:ed.batchSize])
				}
				eventQueue = eventQueue[ed.batchSize:]
			}
		}
	}
}

func (ed *EventsDispatcher) Listen(ctx context.Context) error {
	ed.mu.Lock()
	ed.simpleListenerActive = true
	ed.mu.Unlock()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ed.writeChan:
			lastIdx := math.Min(float64(ed.batchSize), float64(len(ed.eventQueue)))
			batch := ed.eventQueue[:int(lastIdx)]
			for _, addr := range ed.destinations {
				go sendDataSomewhere(context.TODO(), addr, batch)
			}

			ed.mu.Lock()
			ed.eventQueue = ed.eventQueue[int(lastIdx):]
			ed.mu.Unlock()
		}
	}
}
