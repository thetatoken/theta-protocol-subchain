package predeployed

import (
	"strconv"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/rlp"

	"github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
)

// Bytecode of the smart contracts hardcoded in the genesis block through pre-deployment
const TNT20TokenBankContractBytecode = ""

// TNT20TokenBank implements the TokenBank interface.
type TNT20TokenBank struct {
}

func NewTNT20TokenBank() *TNT20TokenBank {
	return &TNT20TokenBank{}
}

// Mint vouchers for the token transferred cross-chain. If the voucher contract for the token does not yet exist, the
// TokenBank contract deploys the the vouncher contract first and then mints the vouchers in the same call.
func (tb *TNT20TokenBank) GenerateMintVouchersProxySctx(blockProposer common.Address, view *slst.StoreView, ccte *core.CrossChainTNT20TransferEvent) (*types.SmartContractTx, error) {
	voucherReceiver := ccte.Receiver
	denom := ccte.Denom
	name := ccte.Name
	symbol := ccte.Symbol
	decimals := ccte.Decimals
	amount := ccte.Amount
	callData, err := rlp.EncodeToBytes([]interface{}{
		mintVouchersFuncSelector,
		denom,
		name,
		symbol,
		strconv.FormatUint(uint64(decimals), 10),
		voucherReceiver.Bytes(),
		amount.Bytes(),
	})
	if err != nil {
		return nil, err
	}

	tnt20TokenBankContractAddr := view.GetTNT20TokenBankContractAddress()
	sctx, err := constructProxySmartContractTx(blockProposer, *tnt20TokenBankContractAddr, callData)
	if err != nil {
		return nil, err
	}

	return sctx, nil
}
