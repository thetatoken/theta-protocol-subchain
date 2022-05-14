package predeployed

import (
	"github.com/thetatoken/theta/crypto"
)

var mintTNT721VouchersFuncSelector = crypto.Keccak256([]byte("mintVouchers(string,address,bytes)"))[:4]
