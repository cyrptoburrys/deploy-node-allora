package invariant_test

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	testCommon "github.com/allora-network/allora-chain/test/common"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

// log wrapper for consistent logging style
func iterationLog(t *testing.T, iteration int, a ...any) {
	t.Log(fmt.Sprint("[ITER ", iteration, "]: ", a))
}

// an actor in the simulation has a
// human readable name,
// string bech32 address,
// and an account with private key etc
type Actor struct {
	name string
	addr string
	acc  cosmosaccount.Account
	lock sync.Mutex
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

// the actors can have nonce issues if you parallelize using them,
// so make sure to check the mutex before sending the tx
func broadcastWithActor(m *testCommon.TestConfig, actor Actor, msgs ...sdktypes.Msg) (cosmosclient.Response, error) {
	actor.lock.Lock()
	ret, err := m.Client.BroadcastTx(m.Ctx, actor.acc, msgs...)
	actor.lock.Unlock()
	return ret, err
}
