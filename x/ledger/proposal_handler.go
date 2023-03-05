package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/x/ledger/keeper"
	"github.com/sojahub/sojahub/x/ledger/types"
	rvotetypes "github.com/sojahub/sojahub/x/rvote/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewProposalHandler(k keeper.Keeper) rvotetypes.Handler {
	return func(ctx sdk.Context, content rvotetypes.Content) error {
		switch c := content.(type) {
		case *types.SetChainEraProposal:
			return k.ProcessSetChainEraProposal(ctx, c)
		case *types.BondReportProposal:
			return k.ProcessBondReportProposal(ctx, c)
		case *types.ActiveReportProposal:
			return k.ProcessActiveReportProposal(ctx, c)
		case *types.TransferReportProposal:
			return k.ProcessTransferReportProposal(ctx, c)
		case *types.ExecuteBondProposal:
			return k.ProcessExecuteBondProposal(ctx, c)
		case *types.InterchainTxProposal:
			return k.ProcessInterchainTxProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
