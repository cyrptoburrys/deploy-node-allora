package invariant_test

import (
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
		maxTopics:           uint64(topicsMax),
		maxReputersPerTopic: maxReputersPerTopic,
		maxWorkersPerTopic:  maxWorkersPerTopic,
		epochLength:         int64(epochLength),
		actors:              actorsList,
		registeredWorkers:   testcommon.NewRandomKeyMap[Registration, struct{}](m.Client.Rand),
		registeredReputers:  testcommon.NewRandomKeyMap[Registration, struct{}](m.Client.Rand),
	}

	// iteration 0, always create a topic to start
	createTopic(m, faucet, 0, &simulationData, 0)

	// every iteration, pick an actor,
	// then pick a state transition function for that actor to do
	for i := 1; i < maxIterations; i++ {
		stateTransition := pickStateTransition(m, i, &simulationData)
		actor, topicId := pickActorAndTopicIdForStateTransition(m, stateTransition, &simulationData, numActors)
		stateTransition.f(m, actor, topicId, &simulationData, i)
		if err != nil {
			m.T.Fatal(err)
		}
	}

}
