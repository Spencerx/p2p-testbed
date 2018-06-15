package main

import (
	"github.com/opentracing/opentracing-go"
	"math/rand"
	"time"
)

// MsgType indicates the type of message being sent
type MsgType string

// Message is a message to send from one peer to the other
type Message struct {
	ID        string
	Type      MsgType
	Created   time.Time
	Deadline  time.Time
	Initiator string
	Headers   map[string]string
	Tracing   opentracing.TextMapCarrier
	Payload   interface{}
}

// NewMessage creates a message
func NewMessage(initiator string, t MsgType, payload interface{}, span opentracing.Span) Message {
	msg := Message{
		ID:        NewMessageID(),
		Type:      t,
		Payload:   payload,
		Initiator: initiator,
		Created:   time.Now(),
		Deadline:  time.Now().Add(time.Minute * 2),
		Headers:   map[string]string{},
	}

	if span != nil {
		msg.Tracing = opentracing.TextMapCarrier{}
		span.Tracer().Inject(span.Context(), opentracing.TextMap, msg.Tracing)
	}

	return msg
}

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// NewMessageID generates a random message identifier
func NewMessageID() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
