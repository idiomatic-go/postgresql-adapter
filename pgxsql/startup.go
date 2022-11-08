package pgxsql

import (
	"github.com/idiomatic-go/common-lib/vhost"
)

var c = make(chan vhost.Message, 10)

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c, nil)
	go receive()
}

func startup(msg vhost.Message) {
	clientStartup(msg)
}

func shutdown() {
	clientShutdown()
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				return
			}
			switch msg.Event {
			case vhost.StartupEvent:
				if !isClientStarted() {
					startup(msg)
				}
			case vhost.ShutdownEvent:
				shutdown()
			}
		}
	}
}
