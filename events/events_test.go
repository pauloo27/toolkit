package events_test

import (
	"testing"

	"github.com/pauloo27/toolkit/events"
	"github.com/stretchr/testify/assert"
)

type Event struct {
	Message string
}

type EventType int

const (
	EventTypeHello EventType = iota
	EventTypeBye
)

func TestEventEmitter(t *testing.T) {
	em := events.NewEventEmitter[EventType, Event]()

	ch := make(chan Event, 1)

	message := "Hola Mundo"
	eventType := EventTypeHello

	em.On(eventType, func(eventType EventType, event Event) {
		assert.Equal(t, eventType, EventTypeHello)
		assert.Equal(t, message, event.Message)
		ch <- event
	})

	em.Emit(eventType, Event{Message: message})

	event := <-ch
	assert.Equal(t, message, event.Message)
}

func TestMultipleListeners(t *testing.T) {
	em := events.NewEventEmitter[EventType, Event]()

	ch := make(chan Event, 2)

	message := "Hola Mundo"
	eventType := EventTypeHello

	em.On(eventType, func(eventType EventType, event Event) {
		assert.Equal(t, eventType, EventTypeHello)
		assert.Equal(t, message, event.Message)
		ch <- event
	})

	em.On(eventType, func(eventType EventType, event Event) {
		assert.Equal(t, eventType, EventTypeHello)
		assert.Equal(t, message, event.Message)
		ch <- event
	})

	em.Emit(eventType, Event{Message: message})

	event := <-ch
	assert.Equal(t, message, event.Message)

	event = <-ch
	assert.Equal(t, message, event.Message)
}

func TestMultipleTypes(t *testing.T) {
	em := events.NewEventEmitter[EventType, Event]()

	ch := make(chan Event, 1)
	messageHello := "Hola Mundo"
	messageBye := "Adios Mundo" // not emitted

	em.On(EventTypeHello, func(eventType EventType, event Event) {
		assert.Equal(t, eventType, EventTypeHello)
		assert.Equal(t, messageHello, event.Message)
		ch <- event
	})

	em.On(EventTypeBye, func(eventType EventType, event Event) {
		assert.Equal(t, eventType, EventTypeBye)
		assert.Equal(t, messageBye, event.Message)
		ch <- event
	})

	em.Emit(EventTypeHello, Event{Message: messageHello})

	event := <-ch
	assert.Equal(t, messageHello, event.Message)
}
