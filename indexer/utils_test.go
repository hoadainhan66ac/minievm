package indexer

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/initia-labs/minievm/x/evm/types"
)

func Test_UnpackData(t *testing.T) {
	resp := types.MsgCreateResponse{
		Result:       "ret",
		ContractAddr: types.StdAddress.Hex(),
		Logs: []types.Log{
			{
				Address: types.StdAddress.Hex(),
				Topics:  []string{"topic"},
				Data:    "data",
			},
		},
	}

	anyResp, err := codectypes.NewAnyWithValue(&resp)
	require.NoError(t, err)

	data, err := proto.Marshal(&sdk.TxMsgData{MsgResponses: []*codectypes.Any{anyResp}})
	require.NoError(t, err)

	var respOut types.MsgCreateResponse
	err = unpackData(data, &respOut)
	require.NoError(t, err)
	require.Equal(t, resp, respOut)
}
