package observer

import (
	"fmt"
)

type IEventManager interface {
	Subscribe(eventType string,  subscriber IEventListener)
	Unsubscribe(eventType string,  subscriber IEventListener)
	Notify(eventType string, data any)
	ListListeners()
}

type EventManager struct {
	listeners map[string][] IEventListener
}
func(eventManager *EventManager) Subscribe(eventType string,  subscriber IEventListener) {
	fmt.Printf("Subscribing subscriber %s to event: %s\n",subscriber.Name(), eventType)
	
	if _, ok := eventManager.listeners[eventType]; !ok {
		eventManager.listeners[eventType] = [] IEventListener{}
	}
	eventManager.listeners[eventType] = append(eventManager.listeners[eventType], subscriber)
}

func(eventManager *EventManager) Unsubscribe(eventType string,  subscriber IEventListener) {
	fmt.Printf("Unsubscribing subscriber %s to event: %s\n",subscriber.Name(), eventType)
	
	if _, ok := eventManager.listeners[eventType]; ok {
		for i, listener := range eventManager.listeners[eventType] {
			if listener == subscriber {
				eventManager.listeners[eventType] = append(eventManager.listeners[eventType][:i], eventManager.listeners[eventType][i+1:]...)
				break
			}
		}
	}
}

func(eventManager *EventManager) Notify(eventType string, data string) {
	fmt.Printf("> Notifying event: %s\n", eventType)
	if _, ok := eventManager.listeners[eventType]; ok {
		for _, listener := range eventManager.listeners[eventType] {
			(listener).Update(data)
		}
	}
}

func(eventManager *EventManager) ListListeners() {
	fmt.Println("\n----------------------------------------")
	fmt.Println("\tListing listeners:")
	fmt.Println("----------------------------------------")
	for eventType, listeners := range eventManager.listeners {
		fmt.Printf("\tEvent: %s\n", eventType)
		for _, listener := range listeners {
			fmt.Printf("\t* %s at address %p\n",listener.Name(), listener)
		}
	}
	fmt.Println("")
}

func NewEventManager() *EventManager {
	return &EventManager{listeners: make(map[string][] IEventListener)}
}