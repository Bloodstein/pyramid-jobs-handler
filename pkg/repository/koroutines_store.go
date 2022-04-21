package repository

type Worker struct {
	Guid    string
	Queue   string
	Abort   chan bool
	Process func(guid string, queue string)
}

type Store struct {
	s map[string]*Worker
}

func (s *Store) Len() int {
	return len(s.s)
}

func NewStore() *Store {
	return &Store{
		s: make(map[string]*Worker),
	}
}

func (s *Store) SaveWorker(w *Worker) {
	s.s[w.Guid] = w
}

func (s *Store) GetWorker(workerGuid string) *Worker {
	if len(s.s) == 0 {
		return nil
	}

	exists := false

	for g := range s.s {
		if g == workerGuid {
			exists = true
		}
	}

	if !exists {
		return nil
	}

	return s.s[workerGuid]
}

func (s *Store) GetWorkers() map[string]*Worker {
	return s.s
}

func (s *Store) DeleteWorker(workerGuid string) {
	delete(s.s, workerGuid)
}
