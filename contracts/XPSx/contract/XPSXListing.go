// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/xpaymentsorg/go-xpayments/accounts/abi"
	"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/core/types"
)

// XPSXListingABI is the input ABI used to generate the binding from.
const XPSXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// XPSXListingBin is the compiled bytecode used for deploying new contracts.
const XPSXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployXPSXListing deploys a new Ethereum contract, binding an instance of XPSXListing to it.
func DeployXPSXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XPSXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(XPSXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XPSXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XPSXListing{XPSXListingCaller: XPSXListingCaller{contract: contract}, XPSXListingTransactor: XPSXListingTransactor{contract: contract}, XPSXListingFilterer: XPSXListingFilterer{contract: contract}}, nil
}

// XPSXListing is an auto generated Go binding around an Ethereum contract.
type XPSXListing struct {
	XPSXListingCaller     // Read-only binding to the contract
	XPSXListingTransactor // Write-only binding to the contract
	XPSXListingFilterer   // Log filterer for contract events
}

// XPSXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type XPSXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XPSXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XPSXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XPSXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XPSXListingSession struct {
	Contract     *XPSXListing      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XPSXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XPSXListingCallerSession struct {
	Contract *XPSXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// XPSXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XPSXListingTransactorSession struct {
	Contract     *XPSXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// XPSXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type XPSXListingRaw struct {
	Contract *XPSXListing // Generic contract binding to access the raw methods on
}

// XPSXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XPSXListingCallerRaw struct {
	Contract *XPSXListingCaller // Generic read-only contract binding to access the raw methods on
}

// XPSXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XPSXListingTransactorRaw struct {
	Contract *XPSXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXPSXListing creates a new instance of XPSXListing, bound to a specific deployed contract.
func NewXPSXListing(address common.Address, backend bind.ContractBackend) (*XPSXListing, error) {
	contract, err := bindXPSXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XPSXListing{XPSXListingCaller: XPSXListingCaller{contract: contract}, XPSXListingTransactor: XPSXListingTransactor{contract: contract}, XPSXListingFilterer: XPSXListingFilterer{contract: contract}}, nil
}

// NewXPSXListingCaller creates a new read-only instance of XPSXListing, bound to a specific deployed contract.
func NewXPSXListingCaller(address common.Address, caller bind.ContractCaller) (*XPSXListingCaller, error) {
	contract, err := bindXPSXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XPSXListingCaller{contract: contract}, nil
}

// NewXPSXListingTransactor creates a new write-only instance of XPSXListing, bound to a specific deployed contract.
func NewXPSXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*XPSXListingTransactor, error) {
	contract, err := bindXPSXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XPSXListingTransactor{contract: contract}, nil
}

// NewXPSXListingFilterer creates a new log filterer instance of XPSXListing, bound to a specific deployed contract.
func NewXPSXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*XPSXListingFilterer, error) {
	contract, err := bindXPSXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XPSXListingFilterer{contract: contract}, nil
}

// bindXPSXListing binds a generic wrapper to an already deployed contract.
func bindXPSXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XPSXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSXListing *XPSXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XPSXListing.Contract.XPSXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSXListing *XPSXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XPSXListing.Contract.XPSXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSXListing *XPSXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XPSXListing.Contract.XPSXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XPSXListing *XPSXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XPSXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XPSXListing *XPSXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XPSXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XPSXListing *XPSXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XPSXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_XPSXListing *XPSXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _XPSXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_XPSXListing *XPSXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _XPSXListing.Contract.GetTokenStatus(&_XPSXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_XPSXListing *XPSXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _XPSXListing.Contract.GetTokenStatus(&_XPSXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XPSXListing *XPSXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XPSXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XPSXListing *XPSXListingSession) Tokens() ([]common.Address, error) {
	return _XPSXListing.Contract.Tokens(&_XPSXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_XPSXListing *XPSXListingCallerSession) Tokens() ([]common.Address, error) {
	return _XPSXListing.Contract.Tokens(&_XPSXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XPSXListing *XPSXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _XPSXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XPSXListing *XPSXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XPSXListing.Contract.Apply(&_XPSXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_XPSXListing *XPSXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _XPSXListing.Contract.Apply(&_XPSXListing.TransactOpts, token)
}
