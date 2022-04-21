package service

import (
	"math/rand"
	"time"

	"github.com/Bloodstein/pyramid-jobs-handler/pkg/repository"
	"github.com/google/uuid"
)

type Launcher struct {
	Store *repository.Store
	repo  repository.Repository
}

func (l Launcher) RunWorker(queue string) {
	guid := uuid.NewString()

	w := &repository.Worker{
		Guid:    guid,
		Queue:   queue,
		Abort:   make(chan bool),
		Process: PopJob,
	}

	l.Store.SaveWorker(w)

	go func() {
		for {
			select {
			case <-w.Abort:
				return
			default:
				w.Process(w.Guid, w.Queue)
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
}
