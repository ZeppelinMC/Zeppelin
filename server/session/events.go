package session

import (
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// A handler function. Should return false to break the event
type Handler[T any] func(T) bool

func NewEvent[T any](handlers ...Handler[T]) Event[T] {
	return Event[T]{handlers}
}

type Event[T any] struct {
	handlers []Handler[T]
}

// Append adds handlers for the event. They will be handled last
func (e Event[T]) Append(handlers ...Handler[T]) {
	e.handlers = append(e.handlers, handlers...)
}

// Prepend adds handlers for the event. They will be handled first
func (e Event[T]) Prepend(handlers ...Handler[T]) {
	e.handlers = append(handlers, e.handlers...)
}

// Override replaces all of the handlers of the event with the new handlers
func (e Event[T]) Override(handlers ...Handler[T]) {
	e.handlers = e.handlers[:copy(e.handlers, handlers)]
}

// Chan creates a new channel handler
func (e Event[T]) Chan() <-chan T {
	c := make(chan T)
	e.Append(func(t T) bool {
		c <- t
		return true
	})

	return c
}

// Await creates a new channel handler and waits for the event to be triggered
func (e Event[T]) Await() {
	<-e.Chan()
}

func (e Event[T]) Trigger(v T) {
	for _, handler := range e.handlers {
		if !handler(v) {
			break
		}
	}
}

type EventManager struct {
	OnSessionAdd    Event[Session]
	OnSessionRemove Event[Session]
	OnChatMessage   Event[ChatMessageEvent]
}

// Default is the default event manager
var Default = EventManager{
	OnSessionAdd:    NewEvent(onSessionAdd),
	OnSessionRemove: NewEvent(onSessionRemove),
	OnChatMessage:   NewEvent(onChatMessage),
}

func onSessionAdd(s Session) bool {
	s.Broadcast().SystemChatMessage(text.Unmarshalf('&', "&e%s joined the game", s.Username()))
	return true
}

func onSessionRemove(s Session) bool {
	s.Broadcast().SystemChatMessage(text.Unmarshalf('&', "&e%s left the game", s.Username()))
	return true
}
