package invariant_test

import (
	"fmt"
	"math/rand"

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

// how many times have we unstaked as a reputer?
func (s *SimulationData) incrementUnstakeAsReputerCount() {
	s.counts.unstakeAsReputer++
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

// pickRandomStakedReputer picks a random reputer that is currently staked
func (s *SimulationData) pickRandomStakedReputer() (Actor, uint64) {
	ret := s.reputerStakes.RandomKey()
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

// markStakeRemovalReputerStake marks a reputer stake for removal in the simulation data
// rather than try to keep copy of such complex state, we just pretend removals are instant
func (s *SimulationData) markStakeRemovalReputerStake(topicId uint64, actor Actor, amount cosmossdk_io_math.Int) {
	reg := Registration{
		TopicId: topicId,
		Actor:   actor,
	}
	prevStake, exists := s.reputerStakes.Get(reg)
	if !exists {
		prevStake = cosmossdk_io_math.ZeroInt()
	}
	newValue := prevStake.Sub(amount)
	if newValue.IsNegative() {
		panic(fmt.Sprintf("negative stake disallowed, topic id %d actor %s amount %s", topicId, actor, amount))
	}
	if !newValue.IsZero() {
		s.reputerStakes.Upsert(reg, newValue)
	} else {
		s.reputerStakes.Delete(reg)
	}
}

// pickRandomPercentOfStakeByReputer picks a random percent (1/10, 1/3, 1/2, 6/7, or full amount) of the stake by a reputer
func (s *SimulationData) pickRandomPercentOfStakeByReputer(rand *rand.Rand, topicId uint64, actor Actor) cosmossdk_io_math.Int {
	reg := Registration{
		TopicId: topicId,
		Actor:   actor,
	}
	stake, exists := s.reputerStakes.Get(reg)
	if !exists {
		return cosmossdk_io_math.ZeroInt()
	}
	percent := rand.Intn(5)
	switch percent {
	case 0:
		return stake.QuoRaw(10)
	case 1:
		return stake.QuoRaw(3)
	case 2:
		return stake.QuoRaw(2)
	case 3:
		return stake.MulRaw(6).QuoRaw(7)
	default:
		return stake
	}
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
