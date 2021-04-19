package design

import "fmt"

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
