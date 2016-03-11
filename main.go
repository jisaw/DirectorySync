package main

import (
	"github.com/howeyc/fsnotify"
	. "github.com/jisaw/goUtils"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	CheckErr(err)

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}

	}()

	err = watcher.Watch("target")
	CheckErr(err)

	<-done

	watcher.Close()

}
