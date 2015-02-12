package main

import (
	"github.com/VideofyMe/go-push-handler/api"
	"github.com/VideofyMe/go-push-handler/pusher"
	"os"
	"strconv"
)

func main() {
	worker_count, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		panic("Define worker count")
	}
	pusher.StartWorkers(worker_count)
	api.Start()
}
