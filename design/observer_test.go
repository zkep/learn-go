package design

import (
	"fmt"
	"testing"
	"time"
)

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	sub.Register(&ObserverOne{})
	sub.Register(&ObserverTwo{})
	sub.Notify("hi")
}

func TestAsyncEventBus(t *testing.T) {
	bus := NewAsyncEventBus()
	bus.Subscribe("topic:1", subOne)
	bus.Subscribe("topic:2", subTwo)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}

func subOne(msgs ...string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("subOne,%+v", msgs)
}

func subTwo(msgs ...string) {
	fmt.Printf("subTwo,%+v", msgs)
}
