package invariant_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	cosmossdk_io_math "cosmossdk.io/math"
	testcommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
)

// log wrapper for consistent logging style
func iterationLog(t *testing.T, iteration int, a ...any) {
	t.Log(fmt.Sprint("[ITER ", iteration, "]: ", a))
}

// an actor in the simulation has a
// human readable name,
// string bech32 address,
// and an account with private key etc
// add a lock to this if you need to broadcast transactions in parallel
// from actors
type Actor struct {
	name string
	addr string
	acc  cosmosaccount.Account
}

// stringer for actor
func (a Actor) String() string {
	return a.name
}

// get the faucet name based on the seed for this test run
func getFaucetName(seed int) string {
	return "run" + strconv.Itoa(seed) + "_faucet"
}

// generates an actors name from seed and index
func getActorName(seed int, actorIndex int) string {
	return "run" + strconv.Itoa(seed) + "_actor" + strconv.Itoa(actorIndex)
}

// generate a libp2p key name for the actor
func getLibP2pKeyName(actor Actor) string {
	return "libp2p_key" + actor.name
}

// generate a multiaddress for an actor
func getMultiAddressName(actor Actor) string {
	return "multiaddress" + actor.name
}

// pick a random topic id that is between 1 and the number of topics
func pickRandomTopicId(m *testcommon.TestConfig) (uint64, error) {
	ctx := context.Background()
	numTopicsResponse, err := m.Client.QueryEmissions().GetNextTopicId(ctx, &emissionstypes.QueryNextTopicIdRequest{})
	if err != nil {
		return 0, err
	}
	ret := m.Client.Rand.Uint64() % numTopicsResponse.NextTopicId
	if ret == 0 {
		ret = 1
	}
	return ret, nil
}

// pick a random balance that is less than half of the actors balance
func pickRandomBalanceLessThanHalf(m *testcommon.TestConfig, actor Actor) (cosmossdk_io_math.Int, error) {
	balOfActor, err := actor.GetBalance(m)
	if err != nil {
		return cosmossdk_io_math.ZeroInt(), err
	}
	randomBalance := balOfActor.QuoRaw(2).QuoRaw(m.Client.Rand.Int63() % 1000)
	return randomBalance, nil
}
