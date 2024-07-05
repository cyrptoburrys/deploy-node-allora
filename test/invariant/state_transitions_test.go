package invariant_test

import (
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
// stake in a reputer (delegate),
// stake as a reputer,
// register as a reputer,
// register as a worker,
// unregister as a reputer,
// unregister as a worker,
// unstake from a reputer (undelegate),
// cancel the removal of delegated stake (delegator),
// collect delegator rewards,
// unstake as a reputer,
// cancel the removal of stake (as a reputer),
// produce an inference (insert a bulk worker payload),
// produce reputation scores (insert a bulk reputer payload)
//
// IMPORTANT: if you change getTransitionNameFromIndex function, you must
// also change this function to match it!!
func allTransitions() []StateTransition {
	return []StateTransition{
		{"createTopic", createTopic},
		{"fundTopic", fundTopic},
		{"registerWorker", registerWorker},
		{"registerReputer", registerReputer},
		{"unregisterWorker", unregisterWorker},
		{"unregisterReputer", unregisterReputer},
	}
}

// state machine dependencies for valid transitions
//
// fundTopic: CreateTopic
// RegisterWorkerForTopic: CreateTopic
// RegisterReputerForTopic: CreateTopic
// stakeReputer: RegisterReputerForTopic, CreateTopic
// delegateStake: CreateTopic, RegisterReputerForTopic
// unRegisterReputer: RegisterReputerForTopic
// unRegisterWorker: RegisterWorkerForTopic
// unstakeReputer: stakeReputer
// cancelStakeRemoval: unstakeReputer
// unstakeDelegator: delegateStake
// cancelDelegateStakeRemoval: unstakeDelegator
// collectDelegatorRewards: delegateStake, fundTopic, InsertBulkWorkerPayload, InsertBulkReputerPayload
// InsertBulkWorkerPayload: RegisterWorkerForTopic, FundTopic
// InsertBulkReputerPayload: RegisterReputerForTopic, InsertBulkWorkerPayload
func isPossibleTransition(data *SimulationData, transition StateTransition) bool {
	switch transition.name {
	case "unregisterWorker":
		return possibleUnregisterWorker(data)
	case "unregisterReputer":
		return possibleUnregisterReputer(data)
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

// pickActorAndTopicIdForStateTransition picks a random actor
// who is able to take the state transition and returns which one it picked.
func pickActorAndTopicIdForStateTransition(
	m *testcommon.TestConfig,
	transition StateTransition,
	data *SimulationData,
	numActors int,
) (Actor, uint64) {
	switch transition.name {
	case "unregisterWorker":
		return data.pickRandomWorkerToUnregister()
	case "unregisterReputer":
		return data.pickRandomReputerToUnregister()
	default:
		randomTopicId, err := pickRandomTopicId(m)
		require.NoError(m.T, err)
		randomActor := data.actors[m.Client.Rand.Intn(numActors)]
		return randomActor, randomTopicId
	}
}
