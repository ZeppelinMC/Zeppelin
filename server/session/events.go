package session

func NewEvent[T any](handlers ...func(T)) Event[T] {
	return Event[T]{handlers}
}

type Event[T any] struct {
	handlers []func(T)
}

// Append adds handlers for the event. They will be handled last
func (e Event[T]) Append(handlers ...func(T)) {
	e.handlers = append(e.handlers, handlers...)
}

// Prepend adds handlers for the event. They will be handled first
func (e Event[T]) Prepend(handlers ...func(T)) {
	e.handlers = append(handlers, e.handlers...)
}

// Shift removes the first i handlers from the event
func (e Event[T]) Shift(i int) {
	e.handlers = e.handlers[i:]
}

// Pop removes the last i handlers
func (e Event[T]) Pop(i int) {
	e.handlers = e.handlers[i:]
}

// Override replaces all of the handlers of the event with the new handlers
func (e Event[T]) Override(handlers ...func(T)) {
	e.handlers = handlers
}

// Chan creates a new channel handler
func (e Event[T]) Chan() <-chan T {
	c := make(chan T)
	e.Append(func(t T) {
		c <- t
	})

	return c
}

// Await creates a new channel handler and waits for the event to be emitted
func (e Event[T]) Await() {
	<-e.Chan()
}

func (e Event[T]) call(v T) {
	for _, handler := range e.handlers {
		handler(v)
	}
}

type EventManager struct {
	OnSessionAdd    Event[Session]
	OnSessionRemove Event[Session]
}

// Default is the default event manager
var Default = EventManager{
	OnSessionAdd: NewEvent(onSessionAdd),
}
