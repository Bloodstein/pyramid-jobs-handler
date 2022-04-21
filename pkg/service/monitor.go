package service

import (
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/repository"
)

type Monitor struct {
	Store *repository.Store
	repo  repository.Repository
}

func (m Monitor) Status() map[string][]string {
	result := make(map[string][]string)

	for _, w := range m.Store.GetWorkers() {
		if _, ok := result[w.Queue]; !ok {
			result[w.Queue] = []string{}
		}
		result[w.Queue] = append(result[w.Queue], w.Guid)
	}

	return result
}
