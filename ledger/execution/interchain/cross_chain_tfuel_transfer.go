package interchain

import (
	"fmt"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/thetasubchain/contracts/predeployed"
	score "github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

func ConstructProxyMintTFuelVoucherSctx(blockProposer common.Address, view *slst.StoreView, icme *score.InterChainMessageEvent) (*types.SmartContractTx, error) {
	ccte, err := score.ParseToCrossChainTFuelTransferEvent(icme)
	if err != nil {
		return nil, err
	}

	err = score.ValidateDenom(ccte.Denom)
	if err != nil {
		return nil, err
	}

	tokenType, err := score.ExtractCrossChainTokenTypeFromDenom(ccte.Denom)
	if err != nil {
		return nil, fmt.Errorf("failed to get token type from denom %v: %v", ccte.Denom, err)
	}
	if tokenType != score.CrossChainTokenTypeTFuel {
		return nil, fmt.Errorf("denom %v is not a TFuel token", ccte.Denom)
	}

	tokenBank := predeployed.NewTFuelTokenBank()
	proxySctx, err := tokenBank.GetMintVouchersProxySctx(blockProposer, view, ccte)
	if err != nil {
		return nil, err
	}

	return proxySctx, nil
}
