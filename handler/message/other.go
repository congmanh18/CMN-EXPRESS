package message

import (
	"express_be/entity"
	"express_be/utils/chat/logger"
	"time"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger("server")
}

var (
	subscribe   = make(chan (chan<- Subscription), 10)
	unsubscribe = make(chan (<-chan Event), 10)
	publish     = make(chan Event, 10)
)

// Event is to define the event
type Event struct {
	EvtType   string
	User      string
	Timestamp int
	Text      string
}

// Subscription is to manage subscribe events
type Subscription struct {
	Archive []Event
	New     <-chan Event
}

// Message is the data structure of messages
type Message struct {
	User      string
	Timestamp int
	Message   string
}

// Subscribe adds a subscriber
func Subscribe() Subscription {
	c := make(chan Subscription)
	subscribe <- c
	return <-c
}

// Join creates a join event
func Join(user string) {
	timestamp := time.Now().Unix()
	publish <- NewEvent("join", user, int(timestamp), "")
}

// Leave creates a leave event
func Leave(user string) {
	timestamp := time.Now().Unix()
	publish <- NewEvent("leave", user, int(timestamp), "")
}

// Say publishes a message event
func Say(msg entity.Message) {
	timestamp := time.Now().Unix()
	publish <- NewEvent("message", *msg.UserID, int(timestamp), *msg.Content)
}

// NewEvent creates a new chat event
func NewEvent(evtType, user string, timestamp int, text string) Event {
	return Event{EvtType: evtType, User: user, Timestamp: timestamp, Text: text}
}

func (s Subscription) Cancel() {
	unsubscribe <- s.New

	for { // infinite loop
		select {
		case _, ok := <-s.New:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
