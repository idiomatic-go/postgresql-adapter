package pgxsql

import (
	"github.com/idiomatic-go/common-lib/vhost"
)

var c = make(chan vhost.Message, 10)
var started = false

func IsStarted() bool {
	return started
}

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c, nil)
	go receive()
}

func startup() {
	clientStartup()
}

func shutdown() {
	vhost.UnregisterPackage(Uri)
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
				if !started {
					started = true
					credentials = vhost.AccessCredentials(&msg)
					startup()
				}
			case vhost.ShutdownEvent:
				shutdown()
			}
		}
	}
}
