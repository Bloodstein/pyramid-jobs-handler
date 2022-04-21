package service

import (
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/repository"
)

type Aborter struct {
	Store *repository.Store
}

func (a Aborter) StopWorker(workerGuid string) {
	w := a.Store.GetWorker(workerGuid)
	if w == nil {
		return
	}
	w.Abort <- true
	a.Store.DeleteWorker(workerGuid)
}
