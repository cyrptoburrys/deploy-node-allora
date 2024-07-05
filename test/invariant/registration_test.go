package invariant_test

import (
	"context"

	testcommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// register actor as a new worker
func registerWorker(m *testcommon.TestConfig, actor Actor, data *SimulationData, iteration int) {
	randomTopicId, err := pickRandomTopicId(m)
	require.NoError(m.T, err)
	txResp, err := broadcastWithActor(m, actor, &emissionstypes.MsgRegister{
		Sender:       actor.addr,
		Owner:        actor.addr, // todo pick random other actor
		LibP2PKey:    getLibP2pKeyName(actor),
		MultiAddress: getMultiAddressName(actor),
		IsReputer:    false,
		TopicId:      randomTopicId,
	})
	require.NoError(m.T, err)

	ctx := context.Background()
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	registerWorkerResponse := &emissionstypes.MsgRegisterResponse{}
	err = txResp.Decode(registerWorkerResponse)
	require.NoError(m.T, err)

	// TODO write back to the simulation data that we registered a worker in topic id
	// incrementCountWorkers(data)
	iterationLog(m.T, iteration, "registered ", actor, "as worker in topic id ", randomTopicId)
}
