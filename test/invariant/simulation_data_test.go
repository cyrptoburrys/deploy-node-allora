package invariant_test

import (
	testcommon "github.com/allora-network/allora-chain/test/common"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
// right now it doesn't need mutexes, if we parallelize this test ever it will
// to read and write out of the simulation data
type SimulationData struct {
	maxTopics           uint64
	maxReputersPerTopic int
	maxWorkersPerTopic  int
	epochLength         int64
	actors              []Actor
	registeredWorkers   *testcommon.RandomKeyMap[Registration, struct{}]
	registeredReputers  *testcommon.RandomKeyMap[Registration, struct{}]
}

type Registration struct {
	TopicId uint64
	Actor   Actor
}

// addWorkerRegistration adds a worker registration to the simulation data
func (s *SimulationData) addWorkerRegistration(topicId uint64, actor Actor) {
	s.registeredWorkers.Insert(Registration{
		TopicId: topicId,
		Actor:   actor,
	}, struct{}{})
}

func (s *SimulationData) removeWorkerRegistration(topicId uint64, actor Actor) {
	s.registeredWorkers.Delete(Registration{
		TopicId: topicId,
		Actor:   actor,
	})
}

// addReputerRegistration adds a reputer registration to the simulation data
func (s *SimulationData) addReputerRegistration(topicId uint64, actor Actor) {
	s.registeredReputers.Insert(Registration{
		TopicId: topicId,
		Actor:   actor,
	}, struct{}{})
}

func (s *SimulationData) removeReputerRegistration(topicId uint64, actor Actor) {
	s.registeredReputers.Delete(Registration{
		TopicId: topicId,
		Actor:   actor,
	})
}

// pickRandomWorkerToUnregister picks a random worker to unregister
func (s *SimulationData) pickRandomWorkerToUnregister() (Actor, uint64) {
	ret := s.registeredWorkers.RandomKey()
	return ret.Actor, ret.TopicId
}

// pickRandomReputerToUnregister picks a random reputer to unregister
func (s *SimulationData) pickRandomReputerToUnregister() (Actor, uint64) {
	ret := s.registeredReputers.RandomKey()
	return ret.Actor, ret.TopicId
}
