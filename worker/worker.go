package worker

import (
	"POC-Load-Videos/tasks"
	"github.com/hibiken/asynq"
	"log"
)

const redisAddress = "127.0.0.1:6379"

func RunWorkers() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddress},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeResizeImage, tasks.HandleResizeImageTask)
	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run asynq server: %v", err)
	}
}
