// Copyright 2016 The go-ethereum Authors
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

// Package xpaymentsclient provides a client for the xPayments RPC API.
package xpaymentsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	xpayments "github.com/xpaymentsorg/go-xpayments"
	"github.com/xpaymentsorg/go-xpayments/common"
	"github.com/xpaymentsorg/go-xpayments/common/hexutil"
	"github.com/xpaymentsorg/go-xpayments/consensus/clique"
	"github.com/xpaymentsorg/go-xpayments/core/types"
	"github.com/xpaymentsorg/go-xpayments/rlp"
	"github.com/xpaymentsorg/go-xpayments/rpc"
)

// Client defines typed wrappers for the xPayments RPC API.
type ClientEth struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func DialEth(rawurl string) (*Client, error) {
	c, err := rpc.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClientEth(c *rpc.Client) *Client {
	return &Client{c}
}

// Blockchain Access

// BlockByHash returns the given full block.
//
// Note that loading full blocks requires two requests. Use HeaderByHash
// if you don't need all transactions or uncle headers.
func (ec *Client) BlockByHashEth(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return ec.getBlock(ctx, "eth_getBlockByHash", hash, true)
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (ec *Client) BlockByNumberEth(ctx context.Context, number *big.Int) (*types.Block, error) {
	return ec.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
}

// LatestBlockNumber gets latest block number
func (ec *Client) LatestBlockNumberEth(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

type rpcBlockEth struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}

func (ec *Client) getBlockEth(ctx context.Context, method string, args ...interface{}) (*types.Block, error) {
	var raw json.RawMessage
	err := ec.c.CallContext(ctx, &raw, method, args...)
	if err != nil {
		return nil, err
	} else if len(raw) == 0 {
		return nil, xpayments.NotFound
	}
	// Decode header and transactions.
	var head *types.Header
	var body rpcBlock
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	// Quick-verify transaction and uncle lists. This mostly helps with debugging the server.
	if head.UncleHash == types.EmptyUncleHash && len(body.UncleHashes) > 0 {
		return nil, fmt.Errorf("server returned non-empty uncle list but block header indicates no uncles")
	}
	if head.UncleHash != types.EmptyUncleHash && len(body.UncleHashes) == 0 {
		return nil, fmt.Errorf("server returned empty uncle list but block header indicates uncles")
	}
	if head.TxHash == types.EmptyRootHash && len(body.Transactions) > 0 {
		return nil, fmt.Errorf("server returned non-empty transaction list but block header indicates no transactions")
	}
	if head.TxHash != types.EmptyRootHash && len(body.Transactions) == 0 {
		return nil, fmt.Errorf("server returned empty transaction list but block header indicates transactions")
	}
	// Load uncles because they are not included in the block response.
	var uncles []*types.Header
	if len(body.UncleHashes) > 0 {
		uncles = make([]*types.Header, len(body.UncleHashes))
		reqs := make([]rpc.BatchElem, len(body.UncleHashes))
		for i := range reqs {
			reqs[i] = rpc.BatchElem{
				Method: "eth_getUncleByBlockHashAndIndex",
				Args:   []interface{}{body.Hash, hexutil.EncodeUint64(uint64(i))},
				Result: &uncles[i],
			}
		}
		if err := ec.c.BatchCallContext(ctx, reqs); err != nil {
			return nil, err
		}
		for i := range reqs {
			if reqs[i].Error != nil {
				return nil, reqs[i].Error
			}
			if uncles[i] == nil {
				return nil, fmt.Errorf("got null header for uncle %d of block %x", i, body.Hash[:])
			}
		}
	}
	// Fill the sender cache of transactions in the block.
	txs := make([]*types.Transaction, len(body.Transactions))
	for i, tx := range body.Transactions {
		if tx.From != nil {
			setSenderFromServer(ctx, tx.tx, *tx.From, body.Hash)
		}
		txs[i] = tx.tx
	}
	return types.NewBlockWithHeader(head).WithBody(txs, uncles), nil
}

// HeaderByHash returns the block header with the given hash.
func (ec *Client) HeaderByHashEth(ctx context.Context, hash common.Hash) (*types.Header, error) {
	var head *types.Header
	err := ec.c.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
	if err == nil && head == nil {
		err = xpayments.NotFound
	}
	return head, err
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (ec *Client) HeaderByNumberEth(ctx context.Context, number *big.Int) (*types.Header, error) {
	var head *types.Header
	err := ec.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err == nil && head == nil {
		err = xpayments.NotFound
	}
	return head, err
}

type rpcTransactionEth struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfoEth struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSONEth(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

// TransactionByHash returns the transaction with the given hash.
func (ec *Client) TransactionByHashEth(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	var json *rpcTransaction
	err = ec.c.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, err
	} else if json == nil {
		return nil, false, xpayments.NotFound
	} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
		return nil, false, fmt.Errorf("server returned transaction without signature")
	}
	if json.From != nil && json.BlockHash != nil {
		setSenderFromServer(ctx, json.tx, *json.From, *json.BlockHash)
	}
	return json.tx, json.BlockNumber == nil, nil
}

type TransactionFullEth struct {
	types.Transaction
	txExtraInfo
}

// TransactionByHashFull returns the transaction with the given hash with extra data included such as From
func (ec *Client) TransactionByHashFullEth(ctx context.Context, hash common.Hash) (tx *TransactionFull, isPending bool, err error) {
	var json *rpcTransaction
	err = ec.c.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, err
	} else if json == nil {
		return nil, false, xpayments.NotFound
	} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
		return nil, false, fmt.Errorf("server returned transaction without signature")
	}
	if json.From != nil && json.BlockHash != nil {
		setSenderFromServer(ctx, json.tx, *json.From, *json.BlockHash)
	}

	return &TransactionFull{Transaction: *json.tx, txExtraInfo: txExtraInfo{From: json.From, BlockHash: json.BlockHash, BlockNumber: json.BlockNumber}}, json.BlockNumber == nil, nil
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (ec *Client) TransactionSenderEth(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	// Try to load the address from the cache.
	sender, err := types.Sender(&senderFromServer{blockhash: block}, tx)
	if err == nil {
		return sender, nil
	}
	var meta struct {
		Hash common.Hash
		From common.Address
	}
	if err = ec.c.CallContext(ctx, &meta, "eth_getTransactionByBlockHashAndIndex", block, hexutil.Uint64(index)); err != nil {
		return common.Address{}, err
	}
	if meta.Hash == (common.Hash{}) || meta.Hash != tx.Hash() {
		return common.Address{}, errors.New("wrong inclusion block/index")
	}
	return meta.From, nil
}

// TransactionCount returns the total number of transactions in the given block.
func (ec *Client) TransactionCountEth(ctx context.Context, blockHash common.Hash) (uint, error) {
	var num hexutil.Uint
	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByHash", blockHash)
	return uint(num), err
}

// TransactionInBlock returns a single transaction at index in the given block.
func (ec *Client) TransactionInBlockEth(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	var json *rpcTransaction
	err := ec.c.CallContext(ctx, &json, "eth_getTransactionByBlockHashAndIndex", blockHash, hexutil.Uint64(index))
	if err == nil {
		if json == nil {
			return nil, xpayments.NotFound
		} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
			return nil, fmt.Errorf("server returned transaction without signature")
		}
	}
	if json.From != nil && json.BlockHash != nil {
		setSenderFromServer(ctx, json.tx, *json.From, *json.BlockHash)
	}
	return json.tx, err
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec *Client) TransactionReceiptEth(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := ec.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, xpayments.NotFound
		}
	}
	return r, err
}

func toBlockNumArgEth(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

type rpcProgressEth struct {
	StartingBlock hexutil.Uint64
	CurrentBlock  hexutil.Uint64
	HighestBlock  hexutil.Uint64
	PulledStates  hexutil.Uint64
	KnownStates   hexutil.Uint64
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (ec *Client) SyncProgressEth(ctx context.Context) (*xpayments.SyncProgress, error) {
	var raw json.RawMessage
	if err := ec.c.CallContext(ctx, &raw, "eth_syncing"); err != nil {
		return nil, err
	}
	// Handle the possible response types
	var syncing bool
	if err := json.Unmarshal(raw, &syncing); err == nil {
		return nil, nil // Not syncing (always false)
	}
	var progress *rpcProgress
	if err := json.Unmarshal(raw, &progress); err != nil {
		return nil, err
	}
	return &xpayments.SyncProgress{
		StartingBlock: uint64(progress.StartingBlock),
		CurrentBlock:  uint64(progress.CurrentBlock),
		HighestBlock:  uint64(progress.HighestBlock),
		PulledStates:  uint64(progress.PulledStates),
		KnownStates:   uint64(progress.KnownStates),
	}, nil
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (ec *Client) SubscribeNewHeadEth(ctx context.Context, ch chan<- *types.Header) (xpayments.Subscription, error) {
	return ec.c.XpsSubscribe(ctx, ch, "newHeads", map[string]struct{}{})
}

// State Access

// NetworkID returns the network ID for this chain.
func (ec *Client) NetworkIDEth(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := ec.c.CallContext(ctx, &ver, "net_version"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}

// ChainID returns the chain ID for this chain.
func (ec *Client) ChainIDEth(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_chainId")
	return (*big.Int)(&result), err
}

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (ec *Client) BalanceAtEth(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (ec *Client) StorageAtEth(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, toBlockNumArg(blockNumber))
	return result, err
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (ec *Client) CodeAtEth(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, toBlockNumArg(blockNumber))
	return result, err
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (ec *Client) NonceAtEth(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, toBlockNumArg(blockNumber))
	return uint64(result), err
}

// Filters

// FilterLogs executes a filter query.
func (ec *Client) FilterLogsEth(ctx context.Context, q xpayments.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = ec.c.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (ec *Client) SubscribeFilterLogsEth(ctx context.Context, q xpayments.FilterQuery, ch chan<- types.Log) (xpayments.Subscription, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	return ec.c.XpsSubscribe(ctx, ch, "logs", arg)
}

func toFilterArgEth(q xpayments.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

// Pending State

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (ec *Client) PendingBalanceAtEth(ctx context.Context, account common.Address) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, "pending")
	return (*big.Int)(&result), err
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (ec *Client) PendingStorageAtEth(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getStorageAt", account, key, "pending")
	return result, err
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (ec *Client) PendingCodeAtEth(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
	return result, err
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (ec *Client) PendingNonceAtEth(ctx context.Context, account common.Address) (uint64, error) {
	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")
	return uint64(result), err
}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (ec *Client) PendingTransactionCountEth(ctx context.Context) (uint, error) {
	var num hexutil.Uint
	err := ec.c.CallContext(ctx, &num, "eth_getBlockTransactionCountByNumber", "pending")
	return uint(num), err
}

// TODO: SubscribePendingTransactions (needs server side)

// Contract Calling

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (ec *Client) CallContractEth(ctx context.Context, msg xpayments.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (ec *Client) PendingCallContractEth(ctx context.Context, msg xpayments.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), "pending")
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (ec *Client) SuggestGasPriceEth(ctx context.Context) (*big.Int, error) {
	var hex hexutil.Big
	if err := ec.c.CallContext(ctx, &hex, "eth_gasPrice"); err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ec *Client) EstimateGasEth(ctx context.Context, msg xpayments.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := ec.c.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(msg))
	if err != nil {
		return 0, err
	}
	return uint64(hex), nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransactionEth(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	return ec.c.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
}

// SignersAt returns the set of clique signers at the given block.
func (ec *Client) SignersAtEth(ctx context.Context, blockNumber *big.Int) ([]common.Address, error) {
	var signers []common.Address
	err := ec.c.CallContext(ctx, &signers, "clique_getSigners", toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return signers, nil
}

// SignersAtHash returns the set of clique signers at the given block.
func (ec *Client) SignersAtHashEth(ctx context.Context, blockHash common.Hash) ([]common.Address, error) {
	var signers []common.Address
	err := ec.c.CallContext(ctx, &signers, "clique_getSignersAtHash", blockHash)
	if err != nil {
		return nil, err
	}
	return signers, nil
}

// VotersAt returns the set of clique voters at the given block.
func (ec *Client) VotersAtEth(ctx context.Context, blockNumber *big.Int) ([]common.Address, error) {
	var voters []common.Address
	err := ec.c.CallContext(ctx, &voters, "clique_getVoters", toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return voters, nil
}

// VotersAtHash returns the set of clique voters at the given block.
func (ec *Client) VotersAtHashEth(ctx context.Context, blockHash common.Hash) ([]common.Address, error) {
	var voters []common.Address
	err := ec.c.CallContext(ctx, &voters, "clique_getVotersAtHash", blockHash)
	if err != nil {
		return nil, err
	}
	return voters, nil
}

// SnapshotAt returns the clique snapshot at the given block.
func (ec *Client) SnapshotAtEth(ctx context.Context, blockNumber *big.Int) (*clique.Snapshot, error) {
	var s clique.Snapshot
	err := ec.c.CallContext(ctx, &s, "clique_getSnapshot", toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// SnapshotAt returns the clique snapshot at the given block.
func (ec *Client) SnapshotAtHashEth(ctx context.Context, blockHash common.Hash) (*clique.Snapshot, error) {
	var s clique.Snapshot
	err := ec.c.CallContext(ctx, &s, "clique_getSnapshotAtHash", blockHash)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func toCallArgEth(msg xpayments.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
