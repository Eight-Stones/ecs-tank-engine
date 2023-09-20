package engine

import "sync"

// syncInfo inner app objects.
type syncInfo struct {
	mutex *sync.Mutex
	jobWG *sync.WaitGroup
}

// init initialize sync info object
func (si *syncInfo) init() {
	si.mutex = &sync.Mutex{}
	si.jobWG = &sync.WaitGroup{}
}
