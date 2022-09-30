package events

// Event is just a type alias used as a generic type constraint
type Event any

// EventType is just a type alias used as a generic type constraint
type EventType comparable

// ListenerFunc represents a event listener for an event (e) of type eventType (ET)
type ListenerFunc[ET EventType, E Event] func(eventType ET, event E)

// EventEmitter is a generic event emitter
type EventEmitter[ET EventType, E Event] struct {
	listeners map[ET][]ListenerFunc[ET, E]
}

// NewEventEmitter creates a new event emitter
func NewEventEmitter[ET EventType, E Event]() *EventEmitter[ET, E] {
	return &EventEmitter[ET, E]{
		listeners: make(map[ET][]ListenerFunc[ET, E]),
	}
}

// On adds a new listener for the given event type
func (emitter *EventEmitter[ET, E]) On(eventType ET, listener ListenerFunc[ET, E]) {
	emitter.listeners[eventType] = append(emitter.listeners[eventType], listener)
}

// Emit emits an event to all listeners of type eventType
func (emitter *EventEmitter[ET, E]) Emit(eventType ET, event E) {
	for _, listener := range emitter.listeners[eventType] {
		listener(eventType, event)
	}
}
