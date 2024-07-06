package invariant_test

import (
	"strconv"

	cosmossdk_io_math "cosmossdk.io/math"
	testcommon "github.com/allora-network/allora-chain/test/common"
)

type StateTransitionCounts struct {
	createTopic       int
	fundTopic         int
	registerWorker    int
	registerReputer   int
	unregisterWorker  int
	unregisterReputer int
	stakeAsReputer    int
}

// stringer for state transition counts
func (s StateTransitionCounts) String() string {
	return "{\ncreateTopic: " + strconv.Itoa(s.createTopic) + ", " +
		"\nfundTopic: " + strconv.Itoa(s.fundTopic) + ", " +
		"\nregisterWorker: " + strconv.Itoa(s.registerWorker) + ", " +
		"\nregisterReputer: " + strconv.Itoa(s.registerReputer) + ", " +
		"\nunregisterWorker: " + strconv.Itoa(s.unregisterWorker) + ", " +
		"\nunregisterReputer: " + strconv.Itoa(s.unregisterReputer) + ", " +
		"\nstakeAsReputer: " + strconv.Itoa(s.stakeAsReputer) +
		"\n}"
}

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
	counts              StateTransitionCounts
	registeredWorkers   *testcommon.RandomKeyMap[Registration, struct{}]
	registeredReputers  *testcommon.RandomKeyMap[Registration, struct{}]
	reputerStakes       *testcommon.RandomKeyMap[Registration, cosmossdk_io_math.Int]
}

type Registration struct {
	TopicId uint64
	Actor   Actor
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

// pickRandomReputerToStake picks a random reputer to stake
func (s *SimulationData) pickRandomReputerToStake() (Actor, uint64) {
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
