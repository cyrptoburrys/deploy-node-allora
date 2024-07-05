package invariant_test

import (
	"os"
	"runtime"
	"testing"

	testcommon "github.com/allora-network/allora-chain/test/common"
)

func TestInvariantTestSuite(t *testing.T) {
	if _, isInvariant := os.LookupEnv("INVARIANT_TEST"); isInvariant == false {
		t.Skip("Skipping Invariant Test unless explicitly enabled")
	}

	numCPUs := runtime.NumCPU()
	gomaxprocs := runtime.GOMAXPROCS(0)
	t.Logf("Number of logical CPUs: %d, GOMAXPROCS %d \n", numCPUs, gomaxprocs)

	t.Log(">>> Setting up connection to local node <<<")

	seed := testcommon.LookupEnvInt(t, "SEED", 1)
	rpcMode := testcommon.LookupRpcMode(t, "RPC_MODE", testcommon.SingleRpc)
	rpcEndpoints := testcommon.LookupEnvStringArray("RPC_URLS", []string{"http://localhost:26657"})

	testConfig := testcommon.NewTestConfig(
		t,
		rpcMode,
		rpcEndpoints,
		"../devnet/genesis",
		seed,
	)

	// Read env vars with defaults
	maxIterations := testcommon.LookupEnvInt(t, "MAX_ITERATIONS", 1000)
	numActors := testcommon.LookupEnvInt(t, "NUM_ACTORS", 100)
	maxReputersPerTopic := testcommon.LookupEnvInt(t, "MAX_REPUTERS_PER_TOPIC", 20)
	maxWorkersPerTopic := testcommon.LookupEnvInt(t, "MAX_WORKERS_PER_TOPIC", 20)
	topicsMax := testcommon.LookupEnvInt(t, "TOPICS_MAX", 100)
	epochLength := testcommon.LookupEnvInt(t, "EPOCH_LENGTH", 12) // in blocks

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
