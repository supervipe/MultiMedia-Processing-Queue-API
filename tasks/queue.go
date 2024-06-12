package tasks

import (
	"github.com/hibiken/asynq"
	"sync"
)

var (
	client *asynq.Client
	once   sync.Once
)

func Init(redisAddress string) {
	once.Do(func() {
		client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddress})
	})
}

func Close() {
	if client != nil {
		err := client.Close()
		if err != nil {
			return
		}
	}
}

func GetClient() *asynq.Client {
	return client
}
