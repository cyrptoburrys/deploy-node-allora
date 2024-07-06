package invariant_test

import (
	"context"

	testcommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// register actor as a new worker
func registerWorker(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "registering ", actor, "as worker in topic id", topicId)
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &emissionstypes.MsgRegister{
		Sender:       actor.addr,
		Owner:        actor.addr, // todo pick random other actor
		LibP2PKey:    getLibP2pKeyName(actor),
		MultiAddress: getMultiAddressName(actor),
		IsReputer:    false,
		TopicId:      topicId,
	})
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	registerWorkerResponse := &emissionstypes.MsgRegisterResponse{}
	err = txResp.Decode(registerWorkerResponse)
	require.NoError(m.T, err)
	require.True(m.T, registerWorkerResponse.Success)

	data.addWorkerRegistration(topicId, actor)
	data.incrementRegisterWorkerCount()
	iterationLog(m.T, iteration, "registered ", actor, "as worker in topic id ", topicId)
}

// determine if this state transition is worth trying based on our knowledge of the state
func anyWorkersRegistered(data *SimulationData) bool {
	return data.registeredWorkers.Len() > 0
}

// unregister actor from being a worker
func unregisterWorker(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "unregistering ", actor, "as worker in topic id", topicId)
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &emissionstypes.MsgRemoveRegistration{
		Sender:    actor.addr,
		TopicId:   topicId,
		IsReputer: false,
	})
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	removeRegistrationResponse := &emissionstypes.MsgRemoveRegistrationResponse{}
	err = txResp.Decode(removeRegistrationResponse)
	require.NoError(m.T, err)
	require.True(m.T, removeRegistrationResponse.Success)

	data.removeWorkerRegistration(topicId, actor)
	data.incrementUnregisterWorkerCount()
	iterationLog(m.T, iteration, "unregistered ", actor, "as worker in topic id ", topicId)
}

// register actor as a new actor
func registerReputer(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "registering ", actor, "as reputer in topic id", topicId)
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &emissionstypes.MsgRegister{
		Sender:       actor.addr,
		Owner:        actor.addr, // todo pick random other actor
		LibP2PKey:    getLibP2pKeyName(actor),
		MultiAddress: getMultiAddressName(actor),
		IsReputer:    true,
		TopicId:      topicId,
	})
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	registerWorkerResponse := &emissionstypes.MsgRegisterResponse{}
	err = txResp.Decode(registerWorkerResponse)
	require.NoError(m.T, err)
	require.True(m.T, registerWorkerResponse.Success)

	data.addReputerRegistration(topicId, actor)
	data.incrementRegisterReputerCount()
	iterationLog(m.T, iteration, "registered ", actor, "as reputer in topic id ", topicId)
}

// determine if this state transition is worth trying based on our knowledge of the state
func anyReputersRegistered(data *SimulationData) bool {
	return data.registeredReputers.Len() > 0
}

// unregister reputer
func unregisterReputer(m *testcommon.TestConfig, actor Actor, topicId uint64, data *SimulationData, iteration int) {
	iterationLog(m.T, iteration, "unregistering ", actor, "as reputer in topic id", topicId)
	ctx := context.Background()
	txResp, err := m.Client.BroadcastTx(ctx, actor.acc, &emissionstypes.MsgRemoveRegistration{
		Sender:    actor.addr,
		TopicId:   topicId,
		IsReputer: true,
	})
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(ctx, txResp.TxHash)
	require.NoError(m.T, err)

	removeRegistrationResponseMsg := &emissionstypes.MsgRemoveRegistrationResponse{}
	err = txResp.Decode(removeRegistrationResponseMsg)
	require.NoError(m.T, err)
	require.True(m.T, removeRegistrationResponseMsg.Success)

	data.removeReputerRegistration(topicId, actor)
	data.incrementUnregisterReputerCount()
	iterationLog(m.T, iteration, "unregistered ", actor, "as reputer in topic id ", topicId)
}
