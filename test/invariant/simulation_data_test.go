package invariant_test

import (
	"sync"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
type SimulationData struct {
	lock                sync.Mutex
	numTopics           uint64
	maxTopics           uint64
	maxReputersPerTopic int
	maxWorkersPerTopic  int
	epochLength         int64
	actors              []Actor
}

// incrementCountTopics increments the number of topics in the simulation data
// must be concurrency safe if we are creating topics in parallel which we are not, yet
func incrementCountTopics(s *SimulationData) {
	s.lock.Lock()
	s.numTopics++
	s.lock.Unlock()
}
