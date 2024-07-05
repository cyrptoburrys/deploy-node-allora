package invariant_test

import (
	"sync"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
type SimulationData struct {
	lock                sync.RWMutex
	maxTopics           uint64
	maxReputersPerTopic int
	maxWorkersPerTopic  int
	epochLength         int64
	actors              []Actor
}
