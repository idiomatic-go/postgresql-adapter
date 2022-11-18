package pgxsql

import (
	"github.com/idiomatic-go/common-lib/eventing"
	"github.com/idiomatic-go/common-lib/vhost"
)

var c = make(chan eventing.Message, 10)

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c)
	go receive()
}

var messageHandler eventing.MessageHandler = func(msg eventing.Message) {
	switch msg.Event {
	case eventing.StartupEvent:
		if !isClientStarted() {
			clientStartup(msg)
		}
	case eventing.ShutdownEvent:
		clientShutdown()
	}
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				return
			}
			messageHandler(msg)
		}
	}
}
