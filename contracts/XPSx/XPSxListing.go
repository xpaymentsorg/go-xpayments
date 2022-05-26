package XPSx

import (
	"github.com/xpaymentsorg/go-xpayments/accounts/abi/bind"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/contracts/XPSx/contract"
)

type XPSXListing struct {
	*contract.XPSXListingSession
	contractBackend bind.ContractBackend
}

func NewMyXPSXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*XPSXListing, error) {
	smartContract, err := contract.NewXPSXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &XPSXListing{
		&contract.XPSXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployXPSXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *XPSXListing, error) {
	contractAddr, _, _, err := contract.DeployXPSXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyXPSXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
