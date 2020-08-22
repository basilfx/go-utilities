package observable

import (
	"sync"

	"github.com/twinj/uuid"
)

// Listener is the type of a reference to a registered listener.
type Listener uuid.UUID

// Observable is a variable that can be observed for changes.
type Observable struct {
	value     interface{}
	listeners map[Listener]chan interface{}
	lock      sync.Mutex
}

// New creates a new observable.
func New() *Observable {
	return &Observable{
		listeners: map[Listener]chan interface{}{},
	}
}

// NewWithValue creates a new observable with an initial value.
func NewWithValue(value interface{}) *Observable {
	o := New()

	o.value = value

	return o
}

// SetValue sets the value of the observable. All listeners will be informed.
func (o *Observable) SetValue(value interface{}) {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.value = value

	for _, v := range o.listeners {
		v <- value
	}
}

// GetValue gets the value of the observable.
func (o *Observable) GetValue() interface{} {
	return o.value
}

// Register interest in value changes.
func (o *Observable) Register() (Listener, chan interface{}) {
	c := make(chan interface{})
	id := Listener(uuid.NewV4())

	o.lock.Lock()
	defer o.lock.Unlock()

	o.listeners[id] = c

	return id, c
}

// Unregister interest in value changes.
func (o *Observable) Unregister(listener Listener) {
	o.lock.Lock()
	defer o.lock.Unlock()

	c, ok := o.listeners[listener]

	if !ok {
		return
	}

	delete(o.listeners, listener)

	close(c)
}
