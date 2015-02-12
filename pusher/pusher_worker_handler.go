package pusher

import (
	"log"
)

var WorkerQueue chan chan Message

func StartWorkers(nworkers int) {
	WorkerQueue = make(chan chan Message, nworkers)

	for i := 0; i < nworkers; i++ {
		log.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}
