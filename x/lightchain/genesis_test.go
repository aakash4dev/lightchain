package lightchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lightchain/testutil/keeper"
	"lightchain/testutil/nullify"
	"lightchain/x/lightchain"
	"lightchain/x/lightchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LightchainKeeper(t)
	lightchain.InitGenesis(ctx, *k, genesisState)
	got := lightchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
