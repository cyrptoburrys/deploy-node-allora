package invariant_test

import (
	"strconv"

	testcommon "github.com/allora-network/allora-chain/test/common"
	"github.com/stretchr/testify/require"
)

// Every function responsible for doing a state transition
// should adhere to this function signature
type StateTransitionFunc func(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int)

// keep track of the name of the state transition as well as the function
type StateTransition struct {
	name string
	f    StateTransitionFunc
}

// The list of possible state transitions we can take are:
//
// create a new topic,
// fund a topic some more,
// register as a reputer,
// register as a worker,
// unregister as a reputer,
// unregister as a worker,
// stake as a reputer,
// stake in a reputer (delegate),
// unstake as a reputer,
// unstake from a reputer (undelegate),
// cancel the removal of stake (as a reputer),
// cancel the removal of delegated stake (delegator),
// collect delegator rewards,
// produce an inference (insert a bulk worker payload),
// produce reputation scores (insert a bulk reputer payload)
func allTransitions() []StateTransition {
	return []StateTransition{
		{"createTopic", createTopic},
		{"fundTopic", fundTopic},
		{"registerWorker", registerWorker},
		{"registerReputer", registerReputer},
		{"unregisterWorker", unregisterWorker},
		{"unregisterReputer", unregisterReputer},
		{"stakeAsReputer", stakeAsReputer},
		{"delegateStake", delegateStake},
		{"unstakeAsReputer", unstakeAsReputer},
	}
}

// state transition counts, keep fields sync with allTransitions above
type StateTransitionCounts struct {
	createTopic       int
	fundTopic         int
	registerWorker    int
	registerReputer   int
	unregisterWorker  int
	unregisterReputer int
	stakeAsReputer    int
	delegateStake     int
	unstakeAsReputer  int
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
		"\ndelegateStake: " + strconv.Itoa(s.delegateStake) +
		"\nunstakeAsReputer: " + strconv.Itoa(s.unstakeAsReputer) +
		"\n}"
}

// state machine dependencies for valid transitions
//
// fundTopic: CreateTopic
// RegisterWorkerForTopic: CreateTopic
// RegisterReputerForTopic: CreateTopic
// unRegisterReputer: RegisterReputerForTopic
// unRegisterWorker: RegisterWorkerForTopic
// stakeReputer: RegisterReputerForTopic, CreateTopic
// delegateStake: CreateTopic, RegisterReputerForTopic
// unstakeReputer: stakeReputer
// unstakeDelegator: delegateStake
// cancelStakeRemoval: unstakeReputer
// cancelDelegateStakeRemoval: unstakeDelegator
// collectDelegatorRewards: delegateStake, fundTopic, InsertBulkWorkerPayload, InsertBulkReputerPayload
// InsertBulkWorkerPayload: RegisterWorkerForTopic, FundTopic
// InsertBulkReputerPayload: RegisterReputerForTopic, InsertBulkWorkerPayload
func isPossibleTransition(data *SimulationData, transition StateTransition) bool {
	switch transition.name {
	case "unregisterWorker":
		return anyWorkersRegistered(data)
	case "unregisterReputer":
		return anyReputersRegistered(data)
	case "stakeAsReputer":
		return anyReputersRegistered(data)
	case "delegateStake":
		return anyReputersRegistered(data)
	case "unstakeAsReputer":
		return anyReputersStaked(data)
	default:
		return true
	}
}

// pickStateTransition picks a random state transition to take and returns which one it picked.
func pickStateTransition(
	m *testcommon.TestConfig,
	iteration int,
	data *SimulationData,
) StateTransition {
	transitions := allTransitions()
	for {
		randIndex := m.Client.Rand.Intn(len(transitions))
		selectedTransition := transitions[randIndex]
		if isPossibleTransition(data, selectedTransition) {
			return selectedTransition
		} else {
			iterationLog(m.T, iteration, "Transition not possible: ", selectedTransition.name)
		}
	}
}

// pickRandomActor picks a random actor from the list of actors in the simulation data
func pickRandomActor(m *testcommon.TestConfig, data *SimulationData) Actor {
	return data.actors[m.Client.Rand.Intn(len(data.actors))]
}

// pickActorAndTopicIdForStateTransition picks a random actor
// who is able to take the state transition and returns which one it picked.
func pickActorAndTopicIdForStateTransition(
	m *testcommon.TestConfig,
	transition StateTransition,
	data *SimulationData,
) (Actor, uint64) {
	switch transition.name {
	case "unregisterWorker":
		return data.pickRandomRegisteredWorker()
	case "unregisterReputer":
		return data.pickRandomRegisteredReputer()
	case "stakeAsReputer":
		return data.pickRandomRegisteredReputer()
	case "delegateStake":
		return data.pickRandomRegisteredReputer()
	case "unstakeAsReputer":
		return data.pickRandomStakedReputer()
	default:
		randomTopicId, err := pickRandomTopicId(m)
		require.NoError(m.T, err)
		randomActor := pickRandomActor(m, data)
		return randomActor, randomTopicId
	}
}
