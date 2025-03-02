package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestValidateMsgCancelOrder(t *testing.T) {
	msg := &types.MsgCancelOrders{
		Creator:      "sei1yezq49upxhunjjhudql2fnj5dgvcwjj87pn2wx",
		ContractAddr: "sei1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqladqwc",
		Cancellations: []*types.Cancellation{{
			Price: sdk.OneDec().Neg(),
		}},
	}
	require.Error(t, msg.ValidateBasic())
}
