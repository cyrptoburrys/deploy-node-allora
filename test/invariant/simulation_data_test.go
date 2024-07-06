package invariant_test

import (
	cosmossdk_io_math "cosmossdk.io/math"
	testcommon "github.com/allora-network/allora-chain/test/common"
)

// SimulationData stores the active set of states we think we're in
// so that we can choose to take a transition that is valid
// right now it doesn't need mutexes, if we parallelize this test ever it will
// to read and write out of the simulation data
type SimulationData struct {
	epochLength        int64
	actors             []Actor
	counts             StateTransitionCounts
	registeredWorkers  *testcommon.RandomKeyMap[Registration, struct{}]
	registeredReputers *testcommon.RandomKeyMap[Registration, struct{}]
	reputerStakes      *testcommon.RandomKeyMap[Registration, cosmossdk_io_math.Int]
	delegatorStakes    *testcommon.RandomKeyMap[Delegation, cosmossdk_io_math.Int]
}

type Registration struct {
	TopicId uint64
	Actor   Actor
}

type Delegation struct {
	TopicId   uint64
	Delegator Actor
	Reputer   Actor
}

// how many times have we created topics?
func (s *SimulationData) incrementCreateTopicCount() {
	s.counts.createTopic++
}

// how many times have we funded topics?
func (s *SimulationData) incrementFundTopicCount() {
	s.counts.fundTopic++
}

// how many times have we registered workers?
func (s *SimulationData) incrementRegisterWorkerCount() {
	s.counts.registerWorker++
}

// how many times have we registered reputers?
func (s *SimulationData) incrementRegisterReputerCount() {
	s.counts.registerReputer++
}

// how many times have we unregistered workers?
func (s *SimulationData) incrementUnregisterWorkerCount() {
	s.counts.unregisterWorker++
}

// how many times have we unregistered reputers?
func (s *SimulationData) incrementUnregisterReputerCount() {
	s.counts.unregisterReputer++
}

// how many times have we staked as a reputer?
func (s *SimulationData) incrementStakeAsReputerCount() {
	s.counts.stakeAsReputer++
}

// how many times have we delegated stake?
func (s *SimulationData) incrementDelegateStakeCount() {
	s.counts.delegateStake++
}

// return the counts of state transitions
func (s *SimulationData) getCounts() StateTransitionCounts {
	return s.counts
}

// addWorkerRegistration adds a worker registration to the simulation data
func (s *SimulationData) addWorkerRegistration(topicId uint64, actor Actor) {
	s.registeredWorkers.Upsert(Registration{
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
	s.registeredReputers.Upsert(Registration{
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

// pickRandomRegisteredWorker picks a random worker that is currently registered
func (s *SimulationData) pickRandomRegisteredWorker() (Actor, uint64) {
	ret := s.registeredWorkers.RandomKey()
	return ret.Actor, ret.TopicId
}

// pickRandomRegisteredReputer picks a random reputer that is currently registered
func (s *SimulationData) pickRandomRegisteredReputer() (Actor, uint64) {
	ret := s.registeredReputers.RandomKey()
	return ret.Actor, ret.TopicId
}

// addReputerStake adds a reputer stake to the simulation data
func (s *SimulationData) addReputerStake(topicId uint64, actor Actor, amount cosmossdk_io_math.Int) {
	reg := Registration{
		TopicId: topicId,
		Actor:   actor,
	}
	prevStake, exists := s.reputerStakes.Get(reg)
	if !exists {
		prevStake = cosmossdk_io_math.ZeroInt()
	}
	newValue := prevStake.Add(amount)
	s.reputerStakes.Upsert(reg, newValue)
}

// addDelegatorStake adds a delegator stake to the simulation data
func (s *SimulationData) addDelegatorStake(topicId uint64, delegator Actor, reputer Actor, amount cosmossdk_io_math.Int) {
	delegation := Delegation{
		TopicId:   topicId,
		Delegator: delegator,
		Reputer:   reputer,
	}
	prevStake, exists := s.delegatorStakes.Get(delegation)
	if !exists {
		prevStake = cosmossdk_io_math.ZeroInt()
	}
	newValue := prevStake.Add(amount)
	s.delegatorStakes.Upsert(delegation, newValue)
}
