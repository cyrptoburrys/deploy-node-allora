package invariant_test

import (
	"context"

	testcommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

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
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &msg)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	response := &emissionstypes.MsgAddStakeResponse{}
	err = txResp.Decode(response)
	require.NoError(m.T, err)

	data.addReputerStake(topicId, actor, randomAmount)
	data.incrementStakeAsReputerCount()
	iterationLog(m.T, iteration, "staked ", actor, "as a reputer in topic id ", topicId, " in amount ", randomAmount.String())
}

// tell if any reputers are currently staked
func anyReputersStaked(data *SimulationData) bool {
	return data.reputerStakes.Len() > 0
}

// mark stake for removal as a reputer
func unstakeAsReputer(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "unstaking as a reputer", actor, "in topic id", topicId)
	randomAmount := data.pickRandomPercentOfStakeByReputer(m.Client.Rand, topicId, actor)
	msg := emissionstypes.MsgRemoveStake{
		Sender:  actor.addr,
		TopicId: topicId,
		Amount:  randomAmount,
	}
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &msg)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	response := &emissionstypes.MsgRemoveStakeResponse{}
	err = txResp.Decode(response)
	require.NoError(m.T, err)

	data.markStakeRemovalReputerStake(topicId, actor, randomAmount)
	data.incrementUnstakeAsReputerCount()
	iterationLog(m.T, iteration, "unstaked from ", actor, "as a reputer in topic id ", topicId, " in amount ", randomAmount.String())
}

// stake as a delegator upon a reputer
// NOTE: in this case, the param actor is the reputer being staked upon,
// rather than the actor doing the staking.
func delegateStake(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "delegating stake")
	delegator := pickRandomActor(m, data)
	randomAmount, err := pickRandomBalanceLessThanHalf(m, delegator)
	require.NoError(m.T, err)
	msg := emissionstypes.MsgDelegateStake{
		Sender:  delegator.addr,
		Reputer: actor.addr,
		TopicId: topicId,
		Amount:  randomAmount,
	}
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, delegator.acc, &msg)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	registerWorkerResponse := &emissionstypes.MsgDelegateStakeResponse{}
	err = txResp.Decode(registerWorkerResponse)
	require.NoError(m.T, err)

	data.addDelegatorStake(topicId, delegator, actor, randomAmount)
	data.incrementDelegateStakeCount()
	iterationLog(m.T, iteration, delegator, "delegated in topic id ", topicId, "upon reputer ", actor, " in amount ", randomAmount.String())
}
