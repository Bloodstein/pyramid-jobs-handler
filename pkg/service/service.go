package service

import (
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/repository"
)

type LauncherService interface {
	RunWorker(queue string)
}

type AborterService interface {
	StopWorker(guid string)
}

type MonitorService interface {
	Status() map[string][]string
}

type Service struct {
	Launcher LauncherService
	Aborter  AborterService
	Monitor  MonitorService
}

func NewLauncher(store *repository.Store, repo repository.Repository) Launcher {
	return Launcher{
		Store: store,
		repo:  repo,
	}
}

func NewAborter(store *repository.Store) Aborter {
	return Aborter{
		Store: store,
	}
}

func NewMonitor(store *repository.Store, repo repository.Repository) Monitor {
	return Monitor{
		Store: store,
		repo:  repo,
	}
}

func NewService(store *repository.Store, repo repository.Repository) Service {
	return Service{
		Launcher: NewLauncher(store, repo),
		Aborter:  NewAborter(store),
		Monitor:  NewMonitor(store, repo),
	}
}
