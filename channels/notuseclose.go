package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "Warning"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	mesage   string
}

var logCh = make(chan logEntry, 50)
var donech = make(chan struct{})

// empty struct has no memory reserved for it
// acts like a boolean channel to indicate that task has been completed

func eg6() {
	go logger()

	// workaround 1 we can use a defer anonymous func and close the routine inside it

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting Down"}
	//noe the loger go routine will be terminated as the main function finishes execution
	time.Sleep(100 * time.Millisecond)
	donech <- struct{}{}

}
func logger() {
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.mesage)
	// }

	for {
		//slect statement does is that the entire statement is going to slot utill a message is received from one of it's channel
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.mesage)
		case <-donech:
			break
		}
	}
}
