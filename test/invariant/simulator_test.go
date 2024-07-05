package invariant_test

import (
	"sync"

	testcommon "github.com/allora-network/allora-chain/test/common"
)

// run the outer loop of the simulator
func simulate(
	m *testcommon.TestConfig,
	maxIterations int,
	numActors int,
	maxReputersPerTopic int,
	maxWorkersPerTopic int,
	topicsMax int,
	epochLength int,
) {
	// fund all actors from the faucet with some amount
	// give everybody the same amount of money to start with
	actorsList := createActors(m, numActors)
	faucet := Actor{
		name: getFaucetName(m.Seed),
		addr: m.FaucetAddr,
		acc:  m.FaucetAcc,
		lock: &sync.Mutex{},
	}
	preFundAmount, err := getPreFundAmount(m, faucet, numActors)
	if err != nil {
		m.T.Fatal(err)
	}
	err = fundActors(
		m,
		faucet,
		actorsList,
		preFundAmount,
	)
	if err != nil {
		m.T.Fatal(err)
	}
	simulationData := SimulationData{
		lock:                sync.RWMutex{},
		maxTopics:           uint64(topicsMax),
		maxReputersPerTopic: maxReputersPerTopic,
		maxWorkersPerTopic:  maxWorkersPerTopic,
		epochLength:         int64(epochLength),
		actors:              actorsList,
	}

	// iteration 0, always create a topic to start
	createTopic(m, faucet, &simulationData, 0)

	// every iteration, pick an actor,
	// then pick a state transition function for that actor to do
	for i := 1; i < maxIterations; i++ {
		actorNum := m.Client.Rand.Intn(numActors)
		stateTransitionFunc := pickActorStateTransition(m, i, actorsList[actorNum], &simulationData)
		stateTransitionFunc(m, actorsList[actorNum], &simulationData, i)
		if err != nil {
			m.T.Fatal(err)
		}
	}

}
