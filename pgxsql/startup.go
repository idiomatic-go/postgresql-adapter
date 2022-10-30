package pgxsql

import (
	start "github.com/idiomatic-go/common-lib/vhost/startup"
)

var c = make(chan start.Message, 10)
var started = false

func IsStarted() bool {
	return started
}

// init - registers package with a channel
func init() {
	start.RegisterPackage(Uri, c, nil)
	go receive()
}

func startup() {
	clientStartup()
}

func shutdown() {
	start.UnregisterPackage(Uri)
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
			case start.StartupEvent:
				if !started {
					started = true
					credentials = start.AccessCredentials(&msg)
					startup()
				}
			case start.ShutdownEvent:
				shutdown()
			}
		}
	}
}
