package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lightchain/testutil/keeper"
	"lightchain/x/lightchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LightchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
