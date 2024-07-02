package invariant_test

import (
	"fmt"
	"sync"

	alloraMath "github.com/allora-network/allora-chain/math"
	testCommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// Use actor to create a new topic
func createTopic(wg *sync.WaitGroup, m *testCommon.TestConfig, actor Actor, data *SimulationData, iteration int) {
	defer wg.Done()
	iterationLog(m.T, iteration, actor, "creating new topic")
	createTopicRequest := &emissionstypes.MsgCreateNewTopic{
		Creator:         actor.addr,
		Metadata:        fmt.Sprintf("Created topic iteration %d", iteration),
		LossLogic:       "bafybeid7mmrv5qr4w5un6c64a6kt2y4vce2vylsmfvnjt7z2wodngknway",
		LossMethod:      "loss-calculation-eth.wasm",
		InferenceLogic:  "bafybeigx43n7kho3gslauwtsenaxehki6ndjo3s63ahif3yc5pltno3pyq",
		InferenceMethod: "allora-inference-function.wasm",
		EpochLength:     data.epochLength,
		GroundTruthLag:  0,
		DefaultArg:      "ETH",
		PNorm:           alloraMath.NewDecFromInt64(3),
		AlphaRegret:     alloraMath.NewDecFromInt64(1),
		AllowNegative:   true,
		Tolerance:       alloraMath.MustNewDecFromString("0.01"),
	}

	txResp, err := broadcastWithActor(m, actor, createTopicRequest)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(m.Ctx, txResp.TxHash)
	require.NoError(m.T, err)

	createTopicResponse := &emissionstypes.MsgCreateNewTopicResponse{}
	err = txResp.Decode(createTopicResponse)
	require.NoError(m.T, err)

	incrementCountTopics(data)

	iterationLog(m.T, iteration, actor, " created topic ", createTopicResponse.TopicId)
}
