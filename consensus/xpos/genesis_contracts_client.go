// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package xpos

import (
	"context"
	"math"
	"math/big"
	"strings"

	"github.com/xpaymentsorg/go-xpayments/accounts/abi"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/hexutil"
	"github.com/xpaymentsorg/go-xpayments/core/state"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/internal/ethapi"
	"github.com/xpaymentsorg/go-xpayments/log"
	"github.com/xpaymentsorg/go-xpayments/params"
	"github.com/xpaymentsorg/go-xpayments/rlp"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

type GenesisContractsClient struct {
	validatorSetABI       abi.ABI
	stateReceiverABI      abi.ABI
	ValidatorContract     string
	StateReceiverContract string
	chainConfig           *params.ChainConfig
	ethAPI                *ethapi.PublicBlockChainAPI
}

const validatorsetABI = `[{"constant":true,"inputs":[],"name":"SPRINT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"SYSTEM_ADDRESS","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"CHAIN","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"FIRST_END_BLOCK","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"producers","outputs":[{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"power","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ROUND_TYPE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"XPOS_ID","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"spanNumbers","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"VOTE_TYPE","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"name":"validators","outputs":[{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"power","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"spans","outputs":[{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint256","name":"startBlock","type":"uint256"},{"internalType":"uint256","name":"endBlock","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"startBlock","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"endBlock","type":"uint256"}],"name":"NewSpan","type":"event"},{"constant":true,"inputs":[],"name":"currentSprint","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"}],"name":"getSpan","outputs":[{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint256","name":"startBlock","type":"uint256"},{"internalType":"uint256","name":"endBlock","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getCurrentSpan","outputs":[{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint256","name":"startBlock","type":"uint256"},{"internalType":"uint256","name":"endBlock","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getNextSpan","outputs":[{"internalType":"uint256","name":"number","type":"uint256"},{"internalType":"uint256","name":"startBlock","type":"uint256"},{"internalType":"uint256","name":"endBlock","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"number","type":"uint256"}],"name":"getSpanByBlock","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"currentSpanNumber","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"}],"name":"getValidatorsTotalStakeBySpan","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"}],"name":"getProducersTotalStakeBySpan","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"name":"getValidatorBySigner","outputs":[{"components":[{"internalType":"uint256","name":"id","type":"uint256"},{"internalType":"uint256","name":"power","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"internalType":"struct XPoSValidatorSet.Validator","name":"result","type":"tuple"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"name":"isValidator","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"},{"internalType":"address","name":"signer","type":"address"}],"name":"isProducer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"signer","type":"address"}],"name":"isCurrentValidator","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"signer","type":"address"}],"name":"isCurrentProducer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"number","type":"uint256"}],"name":"getXPoSValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"},{"internalType":"uint256[]","name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getInitialValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"},{"internalType":"uint256[]","name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"},{"internalType":"uint256[]","name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"newSpan","type":"uint256"},{"internalType":"uint256","name":"startBlock","type":"uint256"},{"internalType":"uint256","name":"endBlock","type":"uint256"},{"internalType":"bytes","name":"validatorBytes","type":"bytes"},{"internalType":"bytes","name":"producerBytes","type":"bytes"}],"name":"commitSpan","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"span","type":"uint256"},{"internalType":"bytes32","name":"dataHash","type":"bytes32"},{"internalType":"bytes","name":"sigs","type":"bytes"}],"name":"getStakePowerBySigs","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"rootHash","type":"bytes32"},{"internalType":"bytes32","name":"leaf","type":"bytes32"},{"internalType":"bytes","name":"proof","type":"bytes"}],"name":"checkMembership","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"d","type":"bytes32"}],"name":"leafNode","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes32","name":"left","type":"bytes32"},{"internalType":"bytes32","name":"right","type":"bytes32"}],"name":"innerNode","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"pure","type":"function"}]`
const stateReceiverABI = `[{"constant":true,"inputs":[],"name":"SYSTEM_ADDRESS","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"lastStateId","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"syncTime","type":"uint256"},{"internalType":"bytes","name":"recordBytes","type":"bytes"}],"name":"commitState","outputs":[{"internalType":"bool","name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

func NewGenesisContractsClient(
	chainConfig *params.ChainConfig,
	validatorContract,
	stateReceiverContract string,
	ethAPI *ethapi.PublicBlockChainAPI,
) *GenesisContractsClient {
	vABI, _ := abi.JSON(strings.NewReader(validatorsetABI))
	sABI, _ := abi.JSON(strings.NewReader(stateReceiverABI))
	return &GenesisContractsClient{
		validatorSetABI:       vABI,
		stateReceiverABI:      sABI,
		ValidatorContract:     validatorContract,
		StateReceiverContract: stateReceiverContract,
		chainConfig:           chainConfig,
		ethAPI:                ethAPI,
	}
}

func (gc *GenesisContractsClient) CommitState(
	event *EventRecordWithTime,
	state *state.StateDB,
	header *types.Header,
	chCtx chainContext,
) error {
	eventRecord := event.BuildEventRecord()
	recordBytes, err := rlp.EncodeToBytes(eventRecord)
	if err != nil {
		return err
	}
	method := "commitState"
	t := event.Time.Unix()
	data, err := gc.stateReceiverABI.Pack(method, big.NewInt(0).SetInt64(t), recordBytes)
	if err != nil {
		log.Error("Unable to pack tx for commitState", "error", err)
		return err
	}
	log.Info("→ committing new state", "eventRecord", event.String())
	msg := getSystemMessage(common.HexToAddress(gc.StateReceiverContract), data)
	if err := applyMessage(msg, state, header, gc.chainConfig, chCtx); err != nil {
		return err
	}
	return nil
}

func (gc *GenesisContractsClient) LastStateId(snapshotNumber uint64) (*big.Int, error) {
	blockNr := rpc.BlockNumber(snapshotNumber)
	method := "lastStateId"
	data, err := gc.stateReceiverABI.Pack(method)
	if err != nil {
		log.Error("Unable to pack tx for LastStateId", "error", err)
		return nil, err
	}

	msgData := (hexutil.Bytes)(data)
	toAddress := common.HexToAddress(gc.StateReceiverContract)
	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
	result, err := gc.ethAPI.Call(context.Background(), ethapi.TransactionArgs{
		Gas:  &gas,
		To:   &toAddress,
		Data: &msgData,
	}, rpc.BlockNumberOrHash{BlockNumber: &blockNr}, nil)
	if err != nil {
		return nil, err
	}

	var ret = new(*big.Int)
	if err := gc.stateReceiverABI.UnpackIntoInterface(ret, method, result); err != nil {
		return nil, err
	}
	return *ret, nil
}
