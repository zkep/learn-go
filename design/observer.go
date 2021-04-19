package design

import (
	"fmt"
	"reflect"
	"sync"
)

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (self *Subject) Register(observer IObserver) {
	self.observers = append(self.observers, observer)
}

func (self *Subject) Remove(observer IObserver) {
	for i, obj := range self.observers {
		if obj == observer {
			self.observers = append(self.observers[:i], self.observers[i+1:]...)
		}
	}
}

func (self *Subject) Notify(msg string) {
	for _, o := range self.observers {
		o.Update(msg)
	}
}

type ObserverOne struct{}

func (obj *ObserverOne) Update(msg string) {
	fmt.Printf("ObserverOne:%s", msg)
}

type ObserverTwo struct{}

func (obj *ObserverTwo) Update(msg string) {
	fmt.Printf("ObserverTwo:%s", msg)
}

type Bus interface {
	Subscribe(topic string, handelr interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (self *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	self.lock.Lock()
	defer self.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := self.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	self.handlers[topic] = handler

	return nil
}

func (self *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := self.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}
	params := make([]reflect.Value, len(args))

	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}

	return
}
