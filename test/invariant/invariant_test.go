package invariant_test

import (
	"os"
	"runtime"
	"testing"

	testCommon "github.com/allora-network/allora-chain/test/common"
)

func TestInvariantTestSuite(t *testing.T) {
	if _, isInvariant := os.LookupEnv("INVARIANT_TEST"); isInvariant == false {
		t.Skip("Skipping Invariant Test unless explicitly enabled")
	}

	numCPUs := runtime.NumCPU()
	gomaxprocs := runtime.GOMAXPROCS(0)
	t.Logf("Number of logical CPUs: %d, GOMAXPROCS %d \n", numCPUs, gomaxprocs)

	t.Log(">>> Setting up connection to local node <<<")

	seed := testCommon.LookupEnvInt(t, "SEED", 1)
	rpcMode := testCommon.LookupRpcMode(t, "RPC_MODE", testCommon.SingleRpc)
	rpcEndpoints := testCommon.LookupEnvStringArray("RPC_URLS", []string{"http://localhost:26657"})

	testConfig := testCommon.NewTestConfig(
		t,
		rpcMode,
		rpcEndpoints,
		"../devnet/genesis",
		seed,
	)

	// Read env vars with defaults
	maxIterations := testCommon.LookupEnvInt(t, "MAX_ITERATIONS", 1000)
	numActors := testCommon.LookupEnvInt(t, "NUM_ACTORS", 100)
	maxReputersPerTopic := testCommon.LookupEnvInt(t, "MAX_REPUTERS_PER_TOPIC", 20)
	maxWorkersPerTopic := testCommon.LookupEnvInt(t, "MAX_WORKERS_PER_TOPIC", 20)
	topicsMax := testCommon.LookupEnvInt(t, "TOPICS_MAX", 100)
	epochLength := testCommon.LookupEnvInt(t, "EPOCH_LENGTH", 12)

	t.Log("Max Actors: ", numActors)
	t.Log("Max Iterations: ", maxIterations)
	t.Log("Max Reputers per topic: ", maxReputersPerTopic)
	t.Log("Max Workers per topic: ", maxWorkersPerTopic)
	t.Log("Epoch Length: ", epochLength)

	t.Log(">>> Starting Test <<<")
	simulate(
		&testConfig,
		maxIterations,
		numActors,
		maxReputersPerTopic,
		maxWorkersPerTopic,
		topicsMax,
		epochLength,
	)
}
