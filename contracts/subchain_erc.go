// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"

	"github.com/thetatoken/theta/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SubchainERCMetaData contains all meta data concerning the SubchainERC contract.
var SubchainERCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_PERSONAL_COLLATERAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getTotalCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorCandidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegalValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"getLegalValidatorsWithBlockHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
	Bin: "0x60806040526040518060400160405280600e81526020017f53696465636861696e546f6b656e0000000000000000000000000000000000008152506009908051906020019062000051929190620000d5565b506040518060400160405280600381526020017f5343540000000000000000000000000000000000000000000000000000000000815250600a90805190602001906200009f929190620000d5565b506012600b60006101000a81548160ff021916908360ff1602179055506064601155348015620000ce57600080fd5b50620001ea565b828054620000e390620001b4565b90600052602060002090601f01602090048101928262000107576000855562000153565b82601f106200012257805160ff191683800117855562000153565b8280016001018555821562000153579182015b828111156200015257825182559160200191906001019062000135565b5b50905062000162919062000166565b5090565b5b808211156200018157600081600090555060010162000167565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620001cd57607f821691505b60208210811415620001e457620001e362000185565b5b50919050565b6124f680620001fa6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806395d89b41116100a2578063a9059cbb11610071578063a9059cbb146102ba578063b309bb3b146102ea578063d6eb591014610308578063dd62ed3e14610326578063f8b2cb4f146103565761010b565b806395d89b41146102345780639864acc214610252578063a0712d6814610282578063a5d5db0c1461029e5761010b565b8063313ce567116100de578063313ce567146101ac57806342966c68146101ca578063473d12ba146101e657806370a08231146102045761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610386565b6040516101259190611b3e565b60405180910390f35b61014860048036038101906101439190611c08565b610414565b6040516101559190611c63565b60405180910390f35b610166610506565b6040516101739190611c8d565b60405180910390f35b61019660048036038101906101919190611ca8565b61050c565b6040516101a39190611c63565b60405180910390f35b6101b46106bd565b6040516101c19190611d17565b60405180910390f35b6101e460048036038101906101df9190611d32565b6106d0565b005b6101ee6107a8565b6040516101fb9190611e1d565b60405180910390f35b61021e60048036038101906102199190611e3f565b61082f565b60405161022b9190611c8d565b60405180910390f35b61023c610847565b6040516102499190611b3e565b60405180910390f35b61026c60048036038101906102679190611d32565b6108d5565b6040516102799190611e1d565b60405180910390f35b61029c60048036038101906102979190611d32565b610c22565b005b6102b860048036038101906102b39190611c08565b610f29565b005b6102d460048036038101906102cf9190611c08565b611805565b6040516102e19190611c63565b60405180910390f35b6102f2611922565b6040516102ff9190611c8d565b60405180910390f35b610310611928565b60405161031d9190611c8d565b60405180910390f35b610340600480360381019061033b9190611e6c565b611932565b60405161034d9190611c8d565b60405180910390f35b610370600480360381019061036b9190611e3f565b611957565b60405161037d9190611c8d565b60405180910390f35b6009805461039390611edb565b80601f01602080910402602001604051908101604052809291908181526020018280546103bf90611edb565b801561040c5780601f106103e15761010080835404028352916020019161040c565b820191906000526020600020905b8154815290600101906020018083116103ef57829003601f168201915b505050505081565b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104f49190611c8d565b60405180910390a36001905092915050565b60065481565b600081600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461059a9190611f3c565b9250508190555081600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105f09190611f3c565b9250508190555081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106469190611f70565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106aa9190611c8d565b60405180910390a3600190509392505050565b600b60009054906101000a900460ff1681565b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461071f9190611f3c565b9250508190555080600660008282546107389190611f3c565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161079d9190611c8d565b60405180910390a350565b6060600d73__IterableMapping_______________________63b3d15a3e90916040518263ffffffff1660e01b81526004016107e49190611fcd565b600060405180830381865af4158015610801573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061082a9190612145565b905090565b60076020528060005260406000206000915090505481565b600a805461085490611edb565b80601f016020809104026020016040519081016040528092919081815260200182805461088090611edb565b80156108cd5780601f106108a2576101008083540402835291602001916108cd565b820191906000526020600020905b8154815290600101906020018083116108b057829003601f168201915b505050505081565b6060438210610919576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610910906121da565b60405180910390fd5b6000600554905060606000821415610935578092505050610c1d565b83600460006001856109479190611f3c565b81526020019081526020016000206000015411610a0e576004600060018461096f9190611f3c565b8152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015610a0057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116109b6575b505050505092505050610c1d565b8360046000808152602001908152602001600020600001541115610a36578092505050610c1d565b600080600184610a469190611f3c565b90505b81811115610b7b57600060028383610a619190611f3c565b610a6b9190612229565b82610a769190611f3c565b90506000600460008381526020019081526020016000206040518060400160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020018280548015610b2757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610add575b50505050508152505090508781600001511415610b505780602001519650505050505050610c1d565b8781600001511015610b6457819350610b74565b600182610b719190611f3c565b92505b5050610a49565b60046000838152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015610c1157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610bc7575b50505050509450505050505b919050565b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c719190611f70565b925050819055508060066000828254610c8a9190611f70565b925050819055506000600d73__IterableMapping_______________________63732a2ccf9091336040518363ffffffff1660e01b8152600401610ccf929190612269565b602060405180830381865af4158015610cec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1091906122a7565b148015610d5e5750601154600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b15610ec057600d73__IterableMapping_______________________63bc2b405c90913360016040518463ffffffff1660e01b8152600401610da293929190612319565b60006040518083038186803b158015610dba57600080fd5b505af4158015610dce573d6000803e3d6000fd5b505050506040518060400160405280438152602001600d73__IterableMapping_______________________63b3d15a3e90916040518263ffffffff1660e01b8152600401610e1d9190611fcd565b600060405180830381865af4158015610e3a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610e639190612145565b815250600460006005548152602001908152602001600020600082015181600001556020820151816001019080519060200190610ea19291906119fe565b50905050600160056000828254610eb89190611f70565b925050819055505b3373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f1e9190611c8d565b60405180910390a350565b80600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610fab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa2906123c2565b60405180910390fd5b6000339050600082905082600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110049190611f70565b9250508190555082600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461105a9190611f3c565b925050819055506000600d73__IterableMapping_______________________63732a2ccf9091876040518363ffffffff1660e01b815260040161109f929190612269565b602060405180830381865af41580156110bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e091906122a7565b14801561112e5750601154600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b1561129057600d73__IterableMapping_______________________63bc2b405c90918660016040518463ffffffff1660e01b815260040161117293929190612319565b60006040518083038186803b15801561118a57600080fd5b505af415801561119e573d6000803e3d6000fd5b505050506040518060400160405280438152602001600d73__IterableMapping_______________________63b3d15a3e90916040518263ffffffff1660e01b81526004016111ed9190611fcd565b600060405180830381865af415801561120a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906112339190612145565b8152506004600060055481526020019081526020016000206000820151816000015560208201518160010190805190602001906112719291906119fe565b509050506001600560008282546112889190611f70565b925050819055505b6001600d73__IterableMapping_______________________63732a2ccf9091336040518363ffffffff1660e01b81526004016112ce929190612269565b602060405180830381865af41580156112eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130f91906122a7565b14801561135c5750601154600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054105b156113d157600d73__IterableMapping_______________________63bc2b405c90913360006040518463ffffffff1660e01b81526004016113a09392919061241d565b60006040518083038186803b1580156113b857600080fd5b505af41580156113cc573d6000803e3d6000fd5b505050505b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060405180604001604052908160008201548152602001600182015481525050905060006115118260200151856119a0565b905060006001846115229190611f70565b90506040518060400160405280438152602001838152506000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020600082015181600001556020820151816001015590505080600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600360008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060405180604001604052908160008201548152602001600182015481525050905060006117248260200151896119a0565b905060006001846117359190611f70565b9050604051806040016040528043815260200183815250600260008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020600082015181600001556020820151816001015590505080600360008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050505050505050505050565b600081600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546118569190611f3c565b9250508190555081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546118ac9190611f70565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516119109190611c8d565b60405180910390a36001905092915050565b60115481565b6000600c54905090565b6008602052816000526040600020602052806000526040600020600091509150505481565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008082846119af9190611f70565b9050838110156119f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119eb906124a0565b60405180910390fd5b8091505092915050565b828054828255906000526020600020908101928215611a77579160200282015b82811115611a765782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611a1e565b5b509050611a849190611a88565b5090565b5b80821115611aa1576000816000905550600101611a89565b5090565b600081519050919050565b600082825260208201905092915050565b60005b83811015611adf578082015181840152602081019050611ac4565b83811115611aee576000848401525b50505050565b6000601f19601f8301169050919050565b6000611b1082611aa5565b611b1a8185611ab0565b9350611b2a818560208601611ac1565b611b3381611af4565b840191505092915050565b60006020820190508181036000830152611b588184611b05565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611b9f82611b74565b9050919050565b611baf81611b94565b8114611bba57600080fd5b50565b600081359050611bcc81611ba6565b92915050565b6000819050919050565b611be581611bd2565b8114611bf057600080fd5b50565b600081359050611c0281611bdc565b92915050565b60008060408385031215611c1f57611c1e611b6a565b5b6000611c2d85828601611bbd565b9250506020611c3e85828601611bf3565b9150509250929050565b60008115159050919050565b611c5d81611c48565b82525050565b6000602082019050611c786000830184611c54565b92915050565b611c8781611bd2565b82525050565b6000602082019050611ca26000830184611c7e565b92915050565b600080600060608486031215611cc157611cc0611b6a565b5b6000611ccf86828701611bbd565b9350506020611ce086828701611bbd565b9250506040611cf186828701611bf3565b9150509250925092565b600060ff82169050919050565b611d1181611cfb565b82525050565b6000602082019050611d2c6000830184611d08565b92915050565b600060208284031215611d4857611d47611b6a565b5b6000611d5684828501611bf3565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611d9481611b94565b82525050565b6000611da68383611d8b565b60208301905092915050565b6000602082019050919050565b6000611dca82611d5f565b611dd48185611d6a565b9350611ddf83611d7b565b8060005b83811015611e10578151611df78882611d9a565b9750611e0283611db2565b925050600181019050611de3565b5085935050505092915050565b60006020820190508181036000830152611e378184611dbf565b905092915050565b600060208284031215611e5557611e54611b6a565b5b6000611e6384828501611bbd565b91505092915050565b60008060408385031215611e8357611e82611b6a565b5b6000611e9185828601611bbd565b9250506020611ea285828601611bbd565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611ef357607f821691505b60208210811415611f0757611f06611eac565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611f4782611bd2565b9150611f5283611bd2565b925082821015611f6557611f64611f0d565b5b828203905092915050565b6000611f7b82611bd2565b9150611f8683611bd2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611fbb57611fba611f0d565b5b828201905092915050565b8082525050565b6000602082019050611fe26000830184611fc6565b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61202582611af4565b810181811067ffffffffffffffff8211171561204457612043611fed565b5b80604052505050565b6000612057611b60565b9050612063828261201c565b919050565b600067ffffffffffffffff82111561208357612082611fed565b5b602082029050602081019050919050565b600080fd5b6000815190506120a881611ba6565b92915050565b60006120c16120bc84612068565b61204d565b905080838252602082019050602084028301858111156120e4576120e3612094565b5b835b8181101561210d57806120f98882612099565b8452602084019350506020810190506120e6565b5050509392505050565b600082601f83011261212c5761212b611fe8565b5b815161213c8482602086016120ae565b91505092915050565b60006020828403121561215b5761215a611b6a565b5b600082015167ffffffffffffffff81111561217957612178611b6f565b5b61218584828501612117565b91505092915050565b7f7468697320626c6f636b206973206e6f74207965742064657465726d696e6564600082015250565b60006121c4602083611ab0565b91506121cf8261218e565b602082019050919050565b600060208201905081810360008301526121f3816121b7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061223482611bd2565b915061223f83611bd2565b92508261224f5761224e6121fa565b5b828204905092915050565b61226381611b94565b82525050565b600060408201905061227e6000830185611fc6565b61228b602083018461225a565b9392505050565b6000815190506122a181611bdc565b92915050565b6000602082840312156122bd576122bc611b6a565b5b60006122cb84828501612292565b91505092915050565b6000819050919050565b6000819050919050565b60006123036122fe6122f9846122d4565b6122de565b611bd2565b9050919050565b612313816122e8565b82525050565b600060608201905061232e6000830186611fc6565b61233b602083018561225a565b612348604083018461230a565b949350505050565b7f596f752068617665206e6f7420656e6f7567682073696465636861696e20746f60008201527f6b656e2100000000000000000000000000000000000000000000000000000000602082015250565b60006123ac602483611ab0565b91506123b782612350565b604082019050919050565b600060208201905081810360008301526123db8161239f565b9050919050565b6000819050919050565b60006124076124026123fd846123e2565b6122de565b611bd2565b9050919050565b612417816123ec565b82525050565b60006060820190506124326000830186611fc6565b61243f602083018561225a565b61244c604083018461240e565b949350505050565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000600082015250565b600061248a601b83611ab0565b915061249582612454565b602082019050919050565b600060208201905081810360008301526124b98161247d565b905091905056fea264697066735822122064d07f30d3849b5f3d26fe19e4b747f08b5b4b81a7b0d4537d8e5d48474a667664736f6c634300080c0033",
}

// SubchainERCABI is the input ABI used to generate the binding from.
// Deprecated: Use SubchainERCMetaData.ABI instead.
var SubchainERCABI = SubchainERCMetaData.ABI

// SubchainERCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubchainERCMetaData.Bin instead.
var SubchainERCBin = SubchainERCMetaData.Bin

// DeploySubchainERC deploys a new Ethereum contract, binding an instance of SubchainERC to it.
func DeploySubchainERC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SubchainERC, error) {
	parsed, err := SubchainERCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubchainERCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubchainERC{SubchainERCCaller: SubchainERCCaller{contract: contract}, SubchainERCTransactor: SubchainERCTransactor{contract: contract}, SubchainERCFilterer: SubchainERCFilterer{contract: contract}}, nil
}

// SubchainERC is an auto generated Go binding around an Ethereum contract.
type SubchainERC struct {
	SubchainERCCaller     // Read-only binding to the contract
	SubchainERCTransactor // Write-only binding to the contract
	SubchainERCFilterer   // Log filterer for contract events
}

// SubchainERCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubchainERCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainERCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubchainERCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainERCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubchainERCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubchainERCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubchainERCSession struct {
	Contract     *SubchainERC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubchainERCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubchainERCCallerSession struct {
	Contract *SubchainERCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SubchainERCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubchainERCTransactorSession struct {
	Contract     *SubchainERCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SubchainERCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubchainERCRaw struct {
	Contract *SubchainERC // Generic contract binding to access the raw methods on
}

// SubchainERCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubchainERCCallerRaw struct {
	Contract *SubchainERCCaller // Generic read-only contract binding to access the raw methods on
}

// SubchainERCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubchainERCTransactorRaw struct {
	Contract *SubchainERCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubchainERC creates a new instance of SubchainERC, bound to a specific deployed contract.
func NewSubchainERC(address common.Address, backend bind.ContractBackend) (*SubchainERC, error) {
	contract, err := bindSubchainERC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubchainERC{SubchainERCCaller: SubchainERCCaller{contract: contract}, SubchainERCTransactor: SubchainERCTransactor{contract: contract}, SubchainERCFilterer: SubchainERCFilterer{contract: contract}}, nil
}

// NewSubchainERCCaller creates a new read-only instance of SubchainERC, bound to a specific deployed contract.
func NewSubchainERCCaller(address common.Address, caller bind.ContractCaller) (*SubchainERCCaller, error) {
	contract, err := bindSubchainERC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainERCCaller{contract: contract}, nil
}

// NewSubchainERCTransactor creates a new write-only instance of SubchainERC, bound to a specific deployed contract.
func NewSubchainERCTransactor(address common.Address, transactor bind.ContractTransactor) (*SubchainERCTransactor, error) {
	contract, err := bindSubchainERC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubchainERCTransactor{contract: contract}, nil
}

// NewSubchainERCFilterer creates a new log filterer instance of SubchainERC, bound to a specific deployed contract.
func NewSubchainERCFilterer(address common.Address, filterer bind.ContractFilterer) (*SubchainERCFilterer, error) {
	contract, err := bindSubchainERC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubchainERCFilterer{contract: contract}, nil
}

// bindSubchainERC binds a generic wrapper to an already deployed contract.
func bindSubchainERC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubchainERCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainERC *SubchainERCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainERC.Contract.SubchainERCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainERC *SubchainERCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainERC.Contract.SubchainERCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainERC *SubchainERCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainERC.Contract.SubchainERCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubchainERC *SubchainERCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubchainERC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubchainERC *SubchainERCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubchainERC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubchainERC *SubchainERCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubchainERC.Contract.contract.Transact(opts, method, params...)
}

// MINPERSONALCOLLATERALAMOUNT is a free data retrieval call binding the contract method 0xb309bb3b.
//
// Solidity: function MIN_PERSONAL_COLLATERAL_AMOUNT() view returns(uint256)
func (_SubchainERC *SubchainERCCaller) MINPERSONALCOLLATERALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "MIN_PERSONAL_COLLATERAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPERSONALCOLLATERALAMOUNT is a free data retrieval call binding the contract method 0xb309bb3b.
//
// Solidity: function MIN_PERSONAL_COLLATERAL_AMOUNT() view returns(uint256)
func (_SubchainERC *SubchainERCSession) MINPERSONALCOLLATERALAMOUNT() (*big.Int, error) {
	return _SubchainERC.Contract.MINPERSONALCOLLATERALAMOUNT(&_SubchainERC.CallOpts)
}

// MINPERSONALCOLLATERALAMOUNT is a free data retrieval call binding the contract method 0xb309bb3b.
//
// Solidity: function MIN_PERSONAL_COLLATERAL_AMOUNT() view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) MINPERSONALCOLLATERALAMOUNT() (*big.Int, error) {
	return _SubchainERC.Contract.MINPERSONALCOLLATERALAMOUNT(&_SubchainERC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SubchainERC *SubchainERCCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SubchainERC *SubchainERCSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.Allowance(&_SubchainERC.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.Allowance(&_SubchainERC.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SubchainERC *SubchainERCCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SubchainERC *SubchainERCSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.BalanceOf(&_SubchainERC.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.BalanceOf(&_SubchainERC.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainERC *SubchainERCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainERC *SubchainERCSession) Decimals() (uint8, error) {
	return _SubchainERC.Contract.Decimals(&_SubchainERC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SubchainERC *SubchainERCCallerSession) Decimals() (uint8, error) {
	return _SubchainERC.Contract.Decimals(&_SubchainERC.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SubchainERC *SubchainERCCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SubchainERC *SubchainERCSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.GetBalance(&_SubchainERC.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _SubchainERC.Contract.GetBalance(&_SubchainERC.CallOpts, addr)
}

// GetLegalValidators is a free data retrieval call binding the contract method 0x473d12ba.
//
// Solidity: function getLegalValidators() view returns(address[])
func (_SubchainERC *SubchainERCCaller) GetLegalValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "getLegalValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLegalValidators is a free data retrieval call binding the contract method 0x473d12ba.
//
// Solidity: function getLegalValidators() view returns(address[])
func (_SubchainERC *SubchainERCSession) GetLegalValidators() ([]common.Address, error) {
	return _SubchainERC.Contract.GetLegalValidators(&_SubchainERC.CallOpts)
}

// GetLegalValidators is a free data retrieval call binding the contract method 0x473d12ba.
//
// Solidity: function getLegalValidators() view returns(address[])
func (_SubchainERC *SubchainERCCallerSession) GetLegalValidators() ([]common.Address, error) {
	return _SubchainERC.Contract.GetLegalValidators(&_SubchainERC.CallOpts)
}

// GetLegalValidatorsWithBlockHeight is a free data retrieval call binding the contract method 0x9864acc2.
//
// Solidity: function getLegalValidatorsWithBlockHeight(uint256 blockHeight) view returns(address[])
func (_SubchainERC *SubchainERCCaller) GetLegalValidatorsWithBlockHeight(opts *bind.CallOpts, blockHeight *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "getLegalValidatorsWithBlockHeight", blockHeight)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLegalValidatorsWithBlockHeight is a free data retrieval call binding the contract method 0x9864acc2.
//
// Solidity: function getLegalValidatorsWithBlockHeight(uint256 blockHeight) view returns(address[])
func (_SubchainERC *SubchainERCSession) GetLegalValidatorsWithBlockHeight(blockHeight *big.Int) ([]common.Address, error) {
	return _SubchainERC.Contract.GetLegalValidatorsWithBlockHeight(&_SubchainERC.CallOpts, blockHeight)
}

// GetLegalValidatorsWithBlockHeight is a free data retrieval call binding the contract method 0x9864acc2.
//
// Solidity: function getLegalValidatorsWithBlockHeight(uint256 blockHeight) view returns(address[])
func (_SubchainERC *SubchainERCCallerSession) GetLegalValidatorsWithBlockHeight(blockHeight *big.Int) ([]common.Address, error) {
	return _SubchainERC.Contract.GetLegalValidatorsWithBlockHeight(&_SubchainERC.CallOpts, blockHeight)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_SubchainERC *SubchainERCCaller) GetTotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "getTotalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_SubchainERC *SubchainERCSession) GetTotalCollateral() (*big.Int, error) {
	return _SubchainERC.Contract.GetTotalCollateral(&_SubchainERC.CallOpts)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) GetTotalCollateral() (*big.Int, error) {
	return _SubchainERC.Contract.GetTotalCollateral(&_SubchainERC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainERC *SubchainERCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainERC *SubchainERCSession) Name() (string, error) {
	return _SubchainERC.Contract.Name(&_SubchainERC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SubchainERC *SubchainERCCallerSession) Name() (string, error) {
	return _SubchainERC.Contract.Name(&_SubchainERC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainERC *SubchainERCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainERC *SubchainERCSession) Symbol() (string, error) {
	return _SubchainERC.Contract.Symbol(&_SubchainERC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SubchainERC *SubchainERCCallerSession) Symbol() (string, error) {
	return _SubchainERC.Contract.Symbol(&_SubchainERC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainERC *SubchainERCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubchainERC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainERC *SubchainERCSession) TotalSupply() (*big.Int, error) {
	return _SubchainERC.Contract.TotalSupply(&_SubchainERC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SubchainERC *SubchainERCCallerSession) TotalSupply() (*big.Int, error) {
	return _SubchainERC.Contract.TotalSupply(&_SubchainERC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Approve(&_SubchainERC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Approve(&_SubchainERC.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SubchainERC *SubchainERCSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Burn(&_SubchainERC.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Burn(&_SubchainERC.TransactOpts, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address validatorCandidate, uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactor) DepositCollateral(opts *bind.TransactOpts, validatorCandidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "depositCollateral", validatorCandidate, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address validatorCandidate, uint256 amount) returns()
func (_SubchainERC *SubchainERCSession) DepositCollateral(validatorCandidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.DepositCollateral(&_SubchainERC.TransactOpts, validatorCandidate, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address validatorCandidate, uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactorSession) DepositCollateral(validatorCandidate common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.DepositCollateral(&_SubchainERC.TransactOpts, validatorCandidate, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_SubchainERC *SubchainERCSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Mint(&_SubchainERC.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_SubchainERC *SubchainERCTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Mint(&_SubchainERC.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Transfer(&_SubchainERC.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.Transfer(&_SubchainERC.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.TransferFrom(&_SubchainERC.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SubchainERC *SubchainERCTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SubchainERC.Contract.TransferFrom(&_SubchainERC.TransactOpts, sender, recipient, amount)
}

// SubchainERCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SubchainERC contract.
type SubchainERCApprovalIterator struct {
	Event *SubchainERCApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubchainERCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainERCApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubchainERCApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubchainERCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainERCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainERCApproval represents a Approval event raised by the SubchainERC contract.
type SubchainERCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainERC *SubchainERCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SubchainERCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainERC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SubchainERCApprovalIterator{contract: _SubchainERC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainERC *SubchainERCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SubchainERCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SubchainERC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainERCApproval)
				if err := _SubchainERC.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SubchainERC *SubchainERCFilterer) ParseApproval(log types.Log) (*SubchainERCApproval, error) {
	event := new(SubchainERCApproval)
	if err := _SubchainERC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubchainERCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SubchainERC contract.
type SubchainERCTransferIterator struct {
	Event *SubchainERCTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubchainERCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubchainERCTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubchainERCTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubchainERCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubchainERCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubchainERCTransfer represents a Transfer event raised by the SubchainERC contract.
type SubchainERCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainERC *SubchainERCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SubchainERCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainERC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubchainERCTransferIterator{contract: _SubchainERC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainERC *SubchainERCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SubchainERCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SubchainERC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubchainERCTransfer)
				if err := _SubchainERC.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SubchainERC *SubchainERCFilterer) ParseTransfer(log types.Log) (*SubchainERCTransfer, error) {
	event := new(SubchainERCTransfer)
	if err := _SubchainERC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}