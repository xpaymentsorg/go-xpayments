// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/accounts/abi"
	"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/event"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208ca3742a52f27c3b3feaa3db9ab5311e5753fe14de190fac8fa51ecaf431f73a0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// XPSValidatorABI is the input ABI used to generate the binding from.
const XPSValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVotedInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerToCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getLatestKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"invalidPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"KYCString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"invalidKYCCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getHashCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"voteInvalidKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kychash\",\"type\":\"string\"}],\"name\":\"uploadKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycHash\",\"type\":\"string\"}],\"name\":\"UploadedKYC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_masternodeOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_masternodes\",\"type\":\"address[]\"}],\"name\":\"InvalidatedNode\",\"type\":\"event\"}]"

// XPSValidatorBin is the compiled bytecode used for deploying new contracts.
const XPSValidatorBin = `0x608060405260006009556000600a553480156200001b57600080fd5b5060405162001e6438038062001e6483398101604090815281516020830151918301516060840151608085015160a086015160c087015160e0880151600b859055600c849055600d839055600e829055600f8190559588018051600955600780546001808201835560009283527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6889091018054600160a060020a031916600160a060020a038a16179055600a805490910190559098979097019694959394929391929091905b8851811015620003015760088982815181101515620000fc57fe5b602090810291909101810151825460018082018555600094855293839020018054600160a060020a031916600160a060020a0392831617905560408051606081018252918b168252918101929092528951908201908a90849081106200015e57fe5b90602001906020020151815250600160008b848151811015156200017e57fe5b602090810291909101810151600160a060020a03908116835282820193909352604091820160009081208551815493870151600160a060020a031990941695169490941760a060020a60ff02191674010000000000000000000000000000000000000000921515929092029190911783559201516001909101558951600291908b90849081106200020b57fe5b6020908102909101810151600160a060020a03908116835282820193909352604091820160009081208054600181018255908252828220018054948c16600160a060020a03199095168517905592835260069052902089518a90839081106200027057fe5b6020908102919091018101518254600180820185556000948552928420018054600160a060020a031916600160a060020a03909216919091179055600b548b519092908c9085908110620002c057fe5b6020908102909101810151600160a060020a0390811683528282019390935260409182016000908120938c16815260029093019052902055600101620000e1565b505050505050505050611b4a806200031a6000396000f3006080604052600436106101955763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301267951811461019a578063025e7c27146101b057806302aa9be2146101e457806306a49fce146102085780630db026221461026d5780630e3e4fb81461029457806315febd68146102cf5780632a3640b1146102e75780632d15cc041461030b5780632f9c4bba1461032c578063302b68721461034157806332658652146103685780633477ee2e146103fe578063441a3e701461041657806358e7525f146104315780635b860d27146104525780635b9cd8cc146104735780636dd7d8ea1461049757806372e44a38146104ab578063a9a981a3146104cc578063a9ff959e146104e1578063ae6e43f5146104f6578063b642facd14610517578063c45607df14610538578063d09f1ab414610559578063d161c7671461056e578063d51b9e9314610583578063d55b7dff146105a4578063ef18374a146105b9578063f2ee3c7d146105ce578063f5c95125146105ef578063f8ac9dd51461060f575b600080fd5b6101ae600160a060020a0360043516610624565b005b3480156101bc57600080fd5b506101c86004356108e5565b60408051600160a060020a039092168252519081900360200190f35b3480156101f057600080fd5b506101ae600160a060020a036004351660243561090d565b34801561021457600080fd5b5061021d610b03565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610259578181015183820152602001610241565b505050509050019250505060405180910390f35b34801561027957600080fd5b50610282610b66565b60408051918252519081900360200190f35b3480156102a057600080fd5b506102bb600160a060020a0360043581169060243516610b6c565b604080519115158252519081900360200190f35b3480156102db57600080fd5b50610282600435610b8c565b3480156102f357600080fd5b506101c8600160a060020a0360043516602435610bab565b34801561031757600080fd5b5061021d600160a060020a0360043516610be2565b34801561033857600080fd5b5061021d610c58565b34801561034d57600080fd5b50610282600160a060020a0360043581169060243516610cb9565b34801561037457600080fd5b50610389600160a060020a0360043516610ce8565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c35781810151838201526020016103ab565b50505050905090810190601f1680156103f05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040a57600080fd5b506101c8600435610e1e565b34801561042257600080fd5b506101ae600435602435610e2c565b34801561043d57600080fd5b50610282600160a060020a0360043516610f5c565b34801561045e57600080fd5b50610282600160a060020a0360043516610f7b565b34801561047f57600080fd5b50610389600160a060020a0360043516602435610fee565b6101ae600160a060020a03600435166110a4565b3480156104b757600080fd5b50610282600160a060020a0360043516611222565b3480156104d857600080fd5b50610282611234565b3480156104ed57600080fd5b5061028261123a565b34801561050257600080fd5b506101ae600160a060020a0360043516611240565b34801561052357600080fd5b506101c8600160a060020a036004351661149a565b34801561054457600080fd5b50610282600160a060020a03600435166114b8565b34801561056557600080fd5b506102826114d3565b34801561057a57600080fd5b506102826114d9565b34801561058f57600080fd5b506102bb600160a060020a03600435166114df565b3480156105b057600080fd5b50610282611504565b3480156105c557600080fd5b5061028261150a565b3480156105da57600080fd5b506101ae600160a060020a0360043516611510565b3480156105fb57600080fd5b506101ae6004803560248101910135611925565b34801561061b57600080fd5b506102826119b2565b600b5460009034101561063657600080fd5b33600090815260036020526040902054151580610660575033600090815260066020526040812054115b151561066b57600080fd5b600160a060020a038216600090815260016020526040902054829060a060020a900460ff161561069a57600080fd5b600160a060020a038316600090815260016020819052604090912001546106c7903463ffffffff6119b816565b6008805460018181019092557ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018054600160a060020a03808816600160a060020a03199283168117909355604080516060810182523380825260208281018881528385018a81526000988952898352858920945185549251151560a060020a0274ff000000000000000000000000000000000000000019919098169290981691909117969096169490941782559351958101959095559183526002909301909252205490925061079e903463ffffffff6119b816565b600160a060020a0384166000908152600160208181526040808420338552600201909152909120919091556009546107db9163ffffffff6119b816565b6009553360009081526006602052604090205415156108415760078054600181810183556000929092527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018054600160a060020a03191633179055600a805490910190555b336000818152600660209081526040808320805460018181018355918552838520018054600160a060020a038a16600160a060020a031991821681179092558186526002855283862080549384018155865294849020909101805490941685179093558051938452908301919091523482820152517f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c19181900360600190a1505050565b60078054829081106108f357fe5b600091825260209091200154600160a060020a0316905081565b600160a060020a03821660009081526001602090815260408083203384526002019091528120548390839081111561094457600080fd5b600160a060020a03828116600090815260016020526040902054163314156109ab57600b54600160a060020a03831660009081526001602090815260408083203384526002019091529020546109a0908363ffffffff6119ce16565b10156109ab57600080fd5b600160a060020a038516600090815260016020819052604090912001546109d8908563ffffffff6119ce16565b600160a060020a038616600090815260016020818152604080842092830194909455338352600290910190522054610a16908563ffffffff6119ce16565b600160a060020a0386166000908152600160209081526040808320338452600201909152902055600f54610a50904363ffffffff6119b816565b33600090815260208181526040808320848452909152902054909350610a7c908563ffffffff6119b816565b33600081815260208181526040808320888452808352818420959095558282526001948501805495860181558352918190209093018690558051918252600160a060020a0388169282019290925280820186905290517faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa29181900360600190a15050505050565b60606008805480602002602001604051908101604052809291908181526020018280548015610b5b57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b3d575b505050505090505b90565b600a5481565b600560209081526000928352604080842090915290825290205460ff1681565b336000908152602081815260408083208484529091529020545b919050565b600660205281600052604060002081815481101515610bc657fe5b600091825260209091200154600160a060020a03169150829050565b600160a060020a038116600090815260026020908152604091829020805483518184028101840190945280845260609392830182828015610c4c57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c2e575b50505050509050919050565b3360009081526020818152604091829020600101805483518184028101840190945280845260609392830182828015610b5b57602002820191906000526020600020905b815481526020019060010190808311610c9c575050505050905090565b600160a060020a0391821660009081526001602090815260408083209390941682526002909201909152205490565b6060610cf3826114df565b15610df65760036000610d058461149a565b600160a060020a0316600160a060020a03168152602001908152602001600020600160036000610d348661149a565b600160a060020a031681526020810191909152604001600020548254919003908110610d5c57fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815292830182828015610dea5780601f10610dbf57610100808354040283529160200191610dea565b820191906000526020600020905b815481529060010190602001808311610dcd57829003601f168201915b50505050509050610ba6565b600160a060020a038216600090815260036020526040902080546000198101908110610d5c57fe5b60088054829081106108f357fe5b60008282828211610e3c57600080fd5b43821115610e4957600080fd5b3360009081526020818152604080832085845290915281205411610e6c57600080fd5b336000908152602081905260409020600101805483919083908110610e8d57fe5b9060005260206000200154141515610ea457600080fd5b3360008181526020818152604080832089845280835290832080549084905593835291905260010180549194509085908110610edc57fe5b60009182526020822001819055604051339185156108fc02918691818181858888f19350505050158015610f14573d6000803e3d6000fd5b50604080513381526020810187905280820185905290517ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360600190a15050505050565b600160a060020a03166000908152600160208190526040909120015490565b600160a060020a0381166000908152600160205260408120548190839060a060020a900460ff161515610fad57600080fd5b610fb68461149a565b9150610fc061150a565b600160a060020a038316600090815260046020526040902054606402811515610fe557fe5b04949350505050565b60036020528160005260406000208181548110151561100957fe5b600091825260209182902001805460408051601f6002600019610100600187161502019094169390930492830185900485028101850190915281815294509092509083018282801561109c5780601f106110715761010080835404028352916020019161109c565b820191906000526020600020905b81548152906001019060200180831161107f57829003601f168201915b505050505081565b600c543410156110b357600080fd5b600160a060020a038116600090815260016020526040902054819060a060020a900460ff1615156110e357600080fd5b600160a060020a03821660009081526001602081905260409091200154611110903463ffffffff6119b816565b600160a060020a038316600090815260016020818152604080842092830194909455338352600290910190522054151561117d57600160a060020a0382166000908152600260209081526040822080546001810182559083529120018054600160a060020a031916331790555b600160a060020a03821660009081526001602090815260408083203384526002019091529020546111b4903463ffffffff6119b816565b600160a060020a0383166000818152600160209081526040808320338085526002909101835292819020949094558351918252810191909152348183015290517f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc9181900360600190a15050565b60046020526000908152604090205481565b60095481565b600f5481565b600160a060020a03808216600090815260016020526040812054909182918291859116331461126e57600080fd5b600160a060020a038516600090815260016020526040902054859060a060020a900460ff16151561129e57600080fd5b600160a060020a0386166000908152600160208190526040909120805474ff0000000000000000000000000000000000000000191690556009546112e79163ffffffff6119ce16565b600955600094505b6008548510156113645785600160a060020a031660088681548110151561131257fe5b600091825260209091200154600160a060020a0316141561135957600880548690811061133b57fe5b60009182526020909120018054600160a060020a0319169055611364565b6001909401936112ef565b600160a060020a0386166000818152600160208181526040808420338552600281018352908420549490935281905201549094506113a8908563ffffffff6119ce16565b600160a060020a0387166000908152600160208181526040808420928301949094553383526002909101905290812055600e546113eb904363ffffffff6119b816565b33600090815260208181526040808320848452909152902054909350611417908563ffffffff6119b816565b33600081815260208181526040808320888452808352818420959095558282526001948501805495860181558352918190209093018690558051918252600160a060020a0389169282019290925281517f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3929181900390910190a1505050505050565b600160a060020a039081166000908152600160205260409020541690565b600160a060020a031660009081526003602052604090205490565b600d5481565b600e5481565b600160a060020a031660009081526001602052604090205460a060020a900460ff1690565b600b5481565b600a5490565b336000818152600160205260408120549091829160609183918291829160a060020a900460ff16151561154257600080fd5b600160a060020a038816600090815260016020526040902054889060a060020a900460ff16151561157257600080fd5b61157b3361149a565b97506115868961149a565b600160a060020a03808a1660009081526005602090815260408083209385168352929052205490975060ff16156115bc57600080fd5b600160a060020a038089166000908152600560209081526040808320938b168352928152828220805460ff19166001908117909155600490915291902080549091019055604b61160a61150a565b600160a060020a03891660009081526004602052604090205460640281151561162f57fe5b041061191a57600160088054905003604051908082528060200260200182016040528015611667578160200160208202803883390190505b50955060009450600093505b6008548410156118065786600160a060020a03166116b360088681548110151561169957fe5b600091825260209091200154600160a060020a031661149a565b600160a060020a031614156117fb576009546116d690600163ffffffff6119ce16565b60095560088054859081106116e757fe5b60009182526020909120015486516001870196600160a060020a03909216918891811061171057fe5b600160a060020a03909216602092830290910190910152600880548590811061173557fe5b600091825260208220018054600160a060020a03191690556008805460019291908790811061176057fe5b6000918252602080832090910154600160a060020a0390811684528382019490945260409283018220805474ffffffffffffffffffffffffffffffffffffffffff19168155600101829055928a168152600390925281206117c0916119e0565b600160a060020a03871660009081526006602052604081206117e191611a01565b600160a060020a0387166000908152600460205260408120555b600190930192611673565b600092505b60075483101561188a5786600160a060020a031660078481548110151561182e57fe5b600091825260209091200154600160a060020a0316141561187f57600780548490811061185757fe5b60009182526020909120018054600160a060020a0319169055600a805460001901905561188a565b60019092019161180b565b7fe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e87876040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156119055781810151838201526020016118ed565b50505050905001935050505060405180910390a15b505050505050505050565b33600090815260036020908152604082208054600181018083559184529190922061195291018484611a1f565b50506040805133808252602082018381529282018490527f949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a9290918591859190606082018484808284376040519201829003965090945050505050a15050565b600c5481565b6000828201838110156119c757fe5b9392505050565b6000828211156119da57fe5b50900390565b50805460008255906000526020600020908101906119fe9190611a9d565b50565b50805460008255906000526020600020908101906119fe9190611ac0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a605782800160ff19823516178555611a8d565b82800160010185558215611a8d579182015b82811115611a8d578235825591602001919060010190611a72565b50611a99929150611ac0565b5090565b610b6391905b80821115611a99576000611ab78282611ada565b50600101611aa3565b610b6391905b80821115611a995760008155600101611ac6565b50805460018160011615610100020316600290046000825580601f10611b0057506119fe565b601f0160209004906000526020600020908101906119fe9190611ac05600a165627a7a7230582091f251d1839ba14c048ab89346724e97476645db9971322df1171fdc81b9bf6e0029`

// DeployXPSValidator deploys a new Ethereum contract, binding an instance of XPSValidator to it.
func DeployXPSValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *XPSValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(XPSValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XPSValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XPSValidator{XPSValidatorCaller: XPSValidatorCaller{contract: contract}, XPSValidatorTransactor: XPSValidatorTransactor{contract: contract}, XPSValidatorFilterer: XPSValidatorFilterer{contract: contract}}, nil
}

// XPSValidator is an auto generated Go binding around an Ethereum contract.
type XPSValidator struct {
	XPSValidatorCaller     // Read-only binding to the contract
	XPSValidatorTransactor // Write-only binding to the contract
	XPSValidatorFilterer   // Log filterer for contract events
}

// XPSValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type XPSValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XPSValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XPSValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XPSValidatorSession struct {
	Contract     *XPSValidator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XPSValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XPSValidatorCallerSession struct {
	Contract *XPSValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// XPSValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XPSValidatorTransactorSession struct {
	Contract     *XPSValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// XPSValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type XPSValidatorRaw struct {
	Contract *XPSValidator // Generic contract binding to access the raw methods on
}

// XPSValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XPSValidatorCallerRaw struct {
	Contract *XPSValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// XPSValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XPSValidatorTransactorRaw struct {
	Contract *XPSValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXPSValidator creates a new instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidator(address common.Address, backend bind.ContractBackend) (*XPSValidator, error) {
	contract, err := bindXPSValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XPSValidator{XPSValidatorCaller: XPSValidatorCaller{contract: contract}, XPSValidatorTransactor: XPSValidatorTransactor{contract: contract}, XPSValidatorFilterer: XPSValidatorFilterer{contract: contract}}, nil
}

// NewXPSValidatorCaller creates a new read-only instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorCaller(address common.Address, caller bind.ContractCaller) (*XPSValidatorCaller, error) {
	contract, err := bindXPSValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XPSValidatorCaller{contract: contract}, nil
}

// NewXPSValidatorTransactor creates a new write-only instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*XPSValidatorTransactor, error) {
	contract, err := bindXPSValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XPSValidatorTransactor{contract: contract}, nil
}

// NewXPSValidatorFilterer creates a new log filterer instance of XPSValidator, bound to a specific deployed contract.
func NewXPSValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*XPSValidatorFilterer, error) {
	contract, err := bindXPSValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XPSValidatorFilterer{contract: contract}, nil
}

// bindXPSValidator binds a generic wrapper to an already deployed contract.
func bindXPSValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XPSValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSValidator *XPSValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XPSValidator.Contract.XPSValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSValidator *XPSValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XPSValidator.Contract.XPSValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSValidator *XPSValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XPSValidator.Contract.XPSValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSValidator *XPSValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XPSValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSValidator *XPSValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XPSValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSValidator *XPSValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XPSValidator.Contract.contract.Transact(opts, method, params...)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorCaller) KYCString(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "KYCString", arg0, arg1)
	return *ret0, err
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _XPSValidator.Contract.KYCString(&_XPSValidator.CallOpts, arg0, arg1)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_XPSValidator *XPSValidatorCallerSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _XPSValidator.Contract.KYCString(&_XPSValidator.CallOpts, arg0, arg1)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "candidateCount")
	return *ret0, err
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) CandidateCount() (*big.Int, error) {
	return _XPSValidator.Contract.CandidateCount(&_XPSValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) CandidateCount() (*big.Int, error) {
	return _XPSValidator.Contract.CandidateCount(&_XPSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "candidateWithdrawDelay")
	return *ret0, err
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _XPSValidator.Contract.CandidateWithdrawDelay(&_XPSValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _XPSValidator.Contract.CandidateWithdrawDelay(&_XPSValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.Candidates(&_XPSValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.Candidates(&_XPSValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
	return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetCandidateCap(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetCandidateCap(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getCandidateOwner", _candidate)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _XPSValidator.Contract.GetCandidateOwner(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _XPSValidator.Contract.GetCandidateOwner(&_XPSValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorSession) GetCandidates() ([]common.Address, error) {
	return _XPSValidator.Contract.GetCandidates(&_XPSValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_XPSValidator *XPSValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _XPSValidator.Contract.GetCandidates(&_XPSValidator.CallOpts)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetHashCount(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getHashCount", _address)
	return *ret0, err
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetHashCount(_address common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetHashCount(&_XPSValidator.CallOpts, _address)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetHashCount(_address common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetHashCount(&_XPSValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorCaller) GetLatestKYC(opts *bind.CallOpts, _address common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getLatestKYC", _address)
	return *ret0, err
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorSession) GetLatestKYC(_address common.Address) (string, error) {
	return _XPSValidator.Contract.GetLatestKYC(&_XPSValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_XPSValidator *XPSValidatorCallerSession) GetLatestKYC(_address common.Address) (string, error) {
	return _XPSValidator.Contract.GetLatestKYC(&_XPSValidator.CallOpts, _address)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetOwnerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getOwnerCount")
	return *ret0, err
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetOwnerCount() (*big.Int, error) {
	return _XPSValidator.Contract.GetOwnerCount(&_XPSValidator.CallOpts)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetOwnerCount() (*big.Int, error) {
	return _XPSValidator.Contract.GetOwnerCount(&_XPSValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
	return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetVoterCap(&_XPSValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.GetVoterCap(&_XPSValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getVoters", _candidate)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _XPSValidator.Contract.GetVoters(&_XPSValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_XPSValidator *XPSValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _XPSValidator.Contract.GetVoters(&_XPSValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getWithdrawBlockNumbers")
	return *ret0, err
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _XPSValidator.Contract.GetWithdrawBlockNumbers(&_XPSValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_XPSValidator *XPSValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _XPSValidator.Contract.GetWithdrawBlockNumbers(&_XPSValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "getWithdrawCap", _blockNumber)
	return *ret0, err
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _XPSValidator.Contract.GetWithdrawCap(&_XPSValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _XPSValidator.Contract.GetWithdrawCap(&_XPSValidator.CallOpts, _blockNumber)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorCaller) HasVotedInvalid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "hasVotedInvalid", arg0, arg1)
	return *ret0, err
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _XPSValidator.Contract.HasVotedInvalid(&_XPSValidator.CallOpts, arg0, arg1)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_XPSValidator *XPSValidatorCallerSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _XPSValidator.Contract.HasVotedInvalid(&_XPSValidator.CallOpts, arg0, arg1)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) InvalidKYCCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "invalidKYCCount", arg0)
	return *ret0, err
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.InvalidKYCCount(&_XPSValidator.CallOpts, arg0)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.InvalidKYCCount(&_XPSValidator.CallOpts, arg0)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) InvalidPercent(opts *bind.CallOpts, _invalidCandidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "invalidPercent", _invalidCandidate)
	return *ret0, err
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.InvalidPercent(&_XPSValidator.CallOpts, _invalidCandidate)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
	return _XPSValidator.Contract.InvalidPercent(&_XPSValidator.CallOpts, _invalidCandidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "isCandidate", _candidate)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _XPSValidator.Contract.IsCandidate(&_XPSValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_XPSValidator *XPSValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _XPSValidator.Contract.IsCandidate(&_XPSValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "maxValidatorNumber")
	return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _XPSValidator.Contract.MaxValidatorNumber(&_XPSValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _XPSValidator.Contract.MaxValidatorNumber(&_XPSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "minCandidateCap")
	return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _XPSValidator.Contract.MinCandidateCap(&_XPSValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _XPSValidator.Contract.MinCandidateCap(&_XPSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "minVoterCap")
	return *ret0, err
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) MinVoterCap() (*big.Int, error) {
	return _XPSValidator.Contract.MinVoterCap(&_XPSValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) MinVoterCap() (*big.Int, error) {
	return _XPSValidator.Contract.MinVoterCap(&_XPSValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) OwnerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "ownerCount")
	return *ret0, err
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) OwnerCount() (*big.Int, error) {
	return _XPSValidator.Contract.OwnerCount(&_XPSValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) OwnerCount() (*big.Int, error) {
	return _XPSValidator.Contract.OwnerCount(&_XPSValidator.CallOpts)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) OwnerToCandidate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "ownerToCandidate", arg0, arg1)
	return *ret0, err
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.OwnerToCandidate(&_XPSValidator.CallOpts, arg0, arg1)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.OwnerToCandidate(&_XPSValidator.CallOpts, arg0, arg1)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.Owners(&_XPSValidator.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_XPSValidator *XPSValidatorCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _XPSValidator.Contract.Owners(&_XPSValidator.CallOpts, arg0)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XPSValidator.contract.Call(opts, out, "voterWithdrawDelay")
	return *ret0, err
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
	return _XPSValidator.Contract.VoterWithdrawDelay(&_XPSValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_XPSValidator *XPSValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
	return _XPSValidator.Contract.VoterWithdrawDelay(&_XPSValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Propose(&_XPSValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Propose(&_XPSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Resign(&_XPSValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Resign(&_XPSValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _XPSValidator.Contract.Unvote(&_XPSValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _XPSValidator.Contract.Unvote(&_XPSValidator.TransactOpts, _candidate, _cap)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorTransactor) UploadKYC(opts *bind.TransactOpts, kychash string) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "uploadKYC", kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorSession) UploadKYC(kychash string) (*types.Transaction, error) {
	return _XPSValidator.Contract.UploadKYC(&_XPSValidator.TransactOpts, kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_XPSValidator *XPSValidatorTransactorSession) UploadKYC(kychash string) (*types.Transaction, error) {
	return _XPSValidator.Contract.UploadKYC(&_XPSValidator.TransactOpts, kychash)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Vote(&_XPSValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.Vote(&_XPSValidator.TransactOpts, _candidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorTransactor) VoteInvalidKYC(opts *bind.TransactOpts, _invalidCandidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "voteInvalidKYC", _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.VoteInvalidKYC(&_XPSValidator.TransactOpts, _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_XPSValidator *XPSValidatorTransactorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
	return _XPSValidator.Contract.VoteInvalidKYC(&_XPSValidator.TransactOpts, _invalidCandidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _XPSValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _XPSValidator.Contract.Withdraw(&_XPSValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_XPSValidator *XPSValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _XPSValidator.Contract.Withdraw(&_XPSValidator.TransactOpts, _blockNumber, _index)
}

// XPSValidatorInvalidatedNodeIterator is returned from FilterInvalidatedNode and is used to iterate over the raw logs and unpacked data for InvalidatedNode events raised by the XPSValidator contract.
type XPSValidatorInvalidatedNodeIterator struct {
	Event *XPSValidatorInvalidatedNode // Event containing the contract specifics and raw log

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
func (it *XPSValidatorInvalidatedNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorInvalidatedNode)
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
		it.Event = new(XPSValidatorInvalidatedNode)
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
func (it *XPSValidatorInvalidatedNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorInvalidatedNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorInvalidatedNode represents a InvalidatedNode event raised by the XPSValidator contract.
type XPSValidatorInvalidatedNode struct {
	MasternodeOwner common.Address
	Masternodes     []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInvalidatedNode is a free log retrieval operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_XPSValidator *XPSValidatorFilterer) FilterInvalidatedNode(opts *bind.FilterOpts) (*XPSValidatorInvalidatedNodeIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "InvalidatedNode")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorInvalidatedNodeIterator{contract: _XPSValidator.contract, event: "InvalidatedNode", logs: logs, sub: sub}, nil
}

// WatchInvalidatedNode is a free log subscription operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_XPSValidator *XPSValidatorFilterer) WatchInvalidatedNode(opts *bind.WatchOpts, sink chan<- *XPSValidatorInvalidatedNode) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "InvalidatedNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorInvalidatedNode)
				if err := _XPSValidator.contract.UnpackLog(event, "InvalidatedNode", log); err != nil {
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

// XPSValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the XPSValidator contract.
type XPSValidatorProposeIterator struct {
	Event *XPSValidatorPropose // Event containing the contract specifics and raw log

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
func (it *XPSValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorPropose)
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
		it.Event = new(XPSValidatorPropose)
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
func (it *XPSValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorPropose represents a Propose event raised by the XPSValidator contract.
type XPSValidatorPropose struct {
	Owner     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*XPSValidatorProposeIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorProposeIterator{contract: _XPSValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *XPSValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorPropose)
				if err := _XPSValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// XPSValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the XPSValidator contract.
type XPSValidatorResignIterator struct {
	Event *XPSValidatorResign // Event containing the contract specifics and raw log

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
func (it *XPSValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorResign)
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
		it.Event = new(XPSValidatorResign)
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
func (it *XPSValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorResign represents a Resign event raised by the XPSValidator contract.
type XPSValidatorResign struct {
	Owner     common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_XPSValidator *XPSValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*XPSValidatorResignIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorResignIterator{contract: _XPSValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_XPSValidator *XPSValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *XPSValidatorResign) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorResign)
				if err := _XPSValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// XPSValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the XPSValidator contract.
type XPSValidatorUnvoteIterator struct {
	Event *XPSValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *XPSValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorUnvote)
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
		it.Event = new(XPSValidatorUnvote)
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
func (it *XPSValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorUnvote represents a Unvote event raised by the XPSValidator contract.
type XPSValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*XPSValidatorUnvoteIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorUnvoteIterator{contract: _XPSValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *XPSValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorUnvote)
				if err := _XPSValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// XPSValidatorUploadedKYCIterator is returned from FilterUploadedKYC and is used to iterate over the raw logs and unpacked data for UploadedKYC events raised by the XPSValidator contract.
type XPSValidatorUploadedKYCIterator struct {
	Event *XPSValidatorUploadedKYC // Event containing the contract specifics and raw log

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
func (it *XPSValidatorUploadedKYCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorUploadedKYC)
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
		it.Event = new(XPSValidatorUploadedKYC)
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
func (it *XPSValidatorUploadedKYCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorUploadedKYCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorUploadedKYC represents a UploadedKYC event raised by the XPSValidator contract.
type XPSValidatorUploadedKYC struct {
	Owner   common.Address
	KycHash string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUploadedKYC is a free log retrieval operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_XPSValidator *XPSValidatorFilterer) FilterUploadedKYC(opts *bind.FilterOpts) (*XPSValidatorUploadedKYCIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "UploadedKYC")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorUploadedKYCIterator{contract: _XPSValidator.contract, event: "UploadedKYC", logs: logs, sub: sub}, nil
}

// WatchUploadedKYC is a free log subscription operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_XPSValidator *XPSValidatorFilterer) WatchUploadedKYC(opts *bind.WatchOpts, sink chan<- *XPSValidatorUploadedKYC) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "UploadedKYC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorUploadedKYC)
				if err := _XPSValidator.contract.UnpackLog(event, "UploadedKYC", log); err != nil {
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

// XPSValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the XPSValidator contract.
type XPSValidatorVoteIterator struct {
	Event *XPSValidatorVote // Event containing the contract specifics and raw log

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
func (it *XPSValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorVote)
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
		it.Event = new(XPSValidatorVote)
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
func (it *XPSValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorVote represents a Vote event raised by the XPSValidator contract.
type XPSValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*XPSValidatorVoteIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorVoteIterator{contract: _XPSValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *XPSValidatorVote) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorVote)
				if err := _XPSValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// XPSValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the XPSValidator contract.
type XPSValidatorWithdrawIterator struct {
	Event *XPSValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *XPSValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XPSValidatorWithdraw)
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
		it.Event = new(XPSValidatorWithdraw)
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
func (it *XPSValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XPSValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XPSValidatorWithdraw represents a Withdraw event raised by the XPSValidator contract.
type XPSValidatorWithdraw struct {
	Owner       common.Address
	BlockNumber *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*XPSValidatorWithdrawIterator, error) {

	logs, sub, err := _XPSValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &XPSValidatorWithdrawIterator{contract: _XPSValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_XPSValidator *XPSValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *XPSValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _XPSValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XPSValidatorWithdraw)
				if err := _XPSValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
