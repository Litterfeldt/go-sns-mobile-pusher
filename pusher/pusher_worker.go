package pusher

import (
	"log"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan Message, 100000)

func NewWorker(id int, workerQueue chan chan Message) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan Message),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan Message
	WorkerQueue chan chan Message
	QuitChan    chan bool
}

func (w Worker) Start() {
	go func() {
		pushr := New()
		for {
			w.WorkerQueue <- w.Work

			select {
			case message := <-w.Work:
				_, err := pushr.Push(message)
				if err != nil {
					w.Log(err)
				} else {
					w.Log("Sent a push")
				}
			case <-w.QuitChan:
				w.Log("Stopped")
				return
			}
		}
	}()
}

func (w Worker) Log(this interface{}) {
	log.Println("WorkerLogger", w.ID, ": ", this)
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
