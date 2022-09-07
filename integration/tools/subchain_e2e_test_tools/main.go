package main

import (
	"github.com/thetatoken/thetasubchain/integration/tools/subchain_e2e_test_tools/cmd"
)

const CHAIN_ID_OFFSET int64 = 360

func main() {

	// oneAccountRegister() //register subchain , its id is 360777
	// oneAccountStake(1)   // use the 10th account to deposit and become validator

	// //TFuel
	// mainchainTFuelLock(big.NewInt(10)) //lock some tfuel tokens on mainchain and transfer to subchain 360777
	// subchainTFuelBurn(big.NewInt(10))

	// //TNT20
	// mainchainTNT20Lock(big.NewInt(100)) //mainchainTNT20Lock(tokenAmount) last address printed in console is subchain voucher contract address
	// subchainTNT20Lock(big.NewInt(10))   //last address printed in console is mainchain voucher contract address
	// mainchainTNT20Burn(big.NewInt(10))
	// subchainTNT20Burn(big.NewInt(10))

	// //TNT721
	// mainchainTNT721Lock(big.NewInt(13)) //mainchainTNT721Lock(tokenid) last address printed in console is subchain voucher contract address
	// subchainTNT721Lock(big.NewInt(12))  //last address printed in console is mainchain voucher contract address
	// mainchainTNT721Burn(big.NewInt(12))
	// subchainTNT721Burn(big.NewInt(13))
	cmd.Execute()
}
