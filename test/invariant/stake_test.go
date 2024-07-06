package invariant_test

import (
	"context"

	testcommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// determine if this state transition is worth trying based on our knowledge of the state
func possibleStakeReputer(data *SimulationData) bool {
	return data.registeredReputers.Len() > 0
}

// stake as a reputer
func stakeAsReputer(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "staking as a reputer", actor, "in topic id", topicId)
	randomAmount, err := pickRandomBalanceLessThanHalf(m, actor)
	require.NoError(m.T, err)
	msg := emissionstypes.MsgAddStake{
		Sender:  actor.addr,
		TopicId: topicId,
		Amount:  randomAmount,
	}
	txResp, err := broadcastWithActor(m, actor, &msg)
	require.NoError(m.T, err)

	ctx := context.Background()
	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	registerWorkerResponse := &emissionstypes.MsgAddStakeResponse{}
	err = txResp.Decode(registerWorkerResponse)
	require.NoError(m.T, err)

	data.addReputerStake(topicId, actor, randomAmount)
	data.incrementStakeAsReputerCount()
	iterationLog(m.T, iteration, "staked ", actor, "as a reputer in topic id ", topicId, " in amount ", randomAmount.String())
}
