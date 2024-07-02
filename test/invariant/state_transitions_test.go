package invariant_test

import (
	"sync"

	testCommon "github.com/allora-network/allora-chain/test/common"
)

// Every function responsible for doing a state transition
// should adhere to this function signature
type StateTransitionFunc func(wg *sync.WaitGroup, m *testCommon.TestConfig, actor Actor, data *SimulationData, iteration int)

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
func isPossibleTransition(actor Actor, data *SimulationData, transition StateTransition) bool {
	return true
}

// pickActorStateTransition picks a random state transition to take and returns which one it picked.
func pickActorStateTransition(
	m *testCommon.TestConfig,
	iteration int,
	actor Actor,
	data *SimulationData,
) StateTransitionFunc {
	transitions := allTransitions()
	for {
		randIndex := m.Client.Rand.Intn(len(transitions))
		selectedTransition := transitions[randIndex]
		if isPossibleTransition(actor, data, selectedTransition) {
			return selectedTransition.f
		} else {
			iterationLog(m.T, iteration, "Transition not possible: ", actor, " ", selectedTransition.name)
		}
	}
}
