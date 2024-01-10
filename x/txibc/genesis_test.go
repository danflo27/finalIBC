package txibc_test

import (
	"testing"

	keepertest "final/testutil/keeper"
	"final/testutil/nullify"
	"final/x/txibc"
	"final/x/txibc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TxibcKeeper(t)
	txibc.InitGenesis(ctx, *k, genesisState)
	got := txibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
