package main

import (
	"github.com/litterfeldt/go-sns-mobile-pusher/api"
	"github.com/litterfeldt/go-sns-mobile-pusher/pusher"
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
