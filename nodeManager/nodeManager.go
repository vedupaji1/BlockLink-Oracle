// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodeManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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
	_ = abi.ConvertType
)

// NodeManagerMetaData contains all meta data concerning the NodeManager contract.
var NodeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_curTickNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tickDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxJoinsInTick\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requiredNodeBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeIdHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExtraDepositForNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMaxJoinsInTick\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMaxJoinsInTick\",\"type\":\"uint256\"}],\"name\":\"MaxJoinsInTickChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeIdHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tickNum\",\"type\":\"uint256\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeIdHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOwner\",\"type\":\"address\"}],\"name\":\"NodeLeft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeIdHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldRequiredNodeBalance\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newRequiredNodeBalance\",\"type\":\"uint256\"}],\"name\":\"RequiredNodeBalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldTickDuration\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newTickDuration\",\"type\":\"uint256\"}],\"name\":\"TickDurationChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"curTickNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"nodeId\",\"type\":\"uint256[4]\"}],\"name\":\"depositForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getMessageForNodeParticipationSig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nodeToRemoveHashId\",\"type\":\"bytes32\"}],\"name\":\"getMessageForNodeRemovalSig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"nodeId\",\"type\":\"uint256[4]\"}],\"name\":\"getNodeIdHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"nodeId\",\"type\":\"uint256[4]\"}],\"name\":\"isNodeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"nodeId\",\"type\":\"uint256[4]\"}],\"name\":\"isValidNodeId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxJoinsInTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodesJoinedInCurTick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nonceForNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"signersPubKey\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes32\",\"name\":\"nodeToRemoveHashId\",\"type\":\"bytes32\"}],\"name\":\"removeNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredNodeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxJoinsInTick\",\"type\":\"uint256\"}],\"name\":\"setNewMaxJoinsInTick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRequiredNodeBalance\",\"type\":\"uint256\"}],\"name\":\"setNewRequiredNodeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTickDuration\",\"type\":\"uint256\"}],\"name\":\"setNewTickDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"nodeId\",\"type\":\"uint256[4]\"}],\"name\":\"withdrawNodeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NodeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeManagerMetaData.ABI instead.
var NodeManagerABI = NodeManagerMetaData.ABI

// NodeManager is an auto generated Go binding around an Ethereum contract.
type NodeManager struct {
	NodeManagerCaller     // Read-only binding to the contract
	NodeManagerTransactor // Write-only binding to the contract
	NodeManagerFilterer   // Log filterer for contract events
}

// NodeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerSession struct {
	Contract     *NodeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerCallerSession struct {
	Contract *NodeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerTransactorSession struct {
	Contract     *NodeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerRaw struct {
	Contract *NodeManager // Generic contract binding to access the raw methods on
}

// NodeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerCallerRaw struct {
	Contract *NodeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerTransactorRaw struct {
	Contract *NodeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManager creates a new instance of NodeManager, bound to a specific deployed contract.
func NewNodeManager(address common.Address, backend bind.ContractBackend) (*NodeManager, error) {
	contract, err := bindNodeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NewNodeManagerCaller creates a new read-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerCaller, error) {
	contract, err := bindNodeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerCaller{contract: contract}, nil
}

// NewNodeManagerTransactor creates a new write-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerTransactor, error) {
	contract, err := bindNodeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTransactor{contract: contract}, nil
}

// NewNodeManagerFilterer creates a new log filterer instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerFilterer, error) {
	contract, err := bindNodeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerFilterer{contract: contract}, nil
}

// bindNodeManager binds a generic wrapper to an already deployed contract.
func bindNodeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.NodeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transact(opts, method, params...)
}

// CurTickNum is a free data retrieval call binding the contract method 0x712368b1.
//
// Solidity: function curTickNum() view returns(uint256)
func (_NodeManager *NodeManagerCaller) CurTickNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "curTickNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurTickNum is a free data retrieval call binding the contract method 0x712368b1.
//
// Solidity: function curTickNum() view returns(uint256)
func (_NodeManager *NodeManagerSession) CurTickNum() (*big.Int, error) {
	return _NodeManager.Contract.CurTickNum(&_NodeManager.CallOpts)
}

// CurTickNum is a free data retrieval call binding the contract method 0x712368b1.
//
// Solidity: function curTickNum() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) CurTickNum() (*big.Int, error) {
	return _NodeManager.Contract.CurTickNum(&_NodeManager.CallOpts)
}

// GetMessageForNodeParticipationSig is a free data retrieval call binding the contract method 0xd2869e27.
//
// Solidity: function getMessageForNodeParticipationSig(string message, uint256 nonce) view returns(bytes, bytes)
func (_NodeManager *NodeManagerCaller) GetMessageForNodeParticipationSig(opts *bind.CallOpts, message string, nonce *big.Int) ([]byte, []byte, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getMessageForNodeParticipationSig", message, nonce)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetMessageForNodeParticipationSig is a free data retrieval call binding the contract method 0xd2869e27.
//
// Solidity: function getMessageForNodeParticipationSig(string message, uint256 nonce) view returns(bytes, bytes)
func (_NodeManager *NodeManagerSession) GetMessageForNodeParticipationSig(message string, nonce *big.Int) ([]byte, []byte, error) {
	return _NodeManager.Contract.GetMessageForNodeParticipationSig(&_NodeManager.CallOpts, message, nonce)
}

// GetMessageForNodeParticipationSig is a free data retrieval call binding the contract method 0xd2869e27.
//
// Solidity: function getMessageForNodeParticipationSig(string message, uint256 nonce) view returns(bytes, bytes)
func (_NodeManager *NodeManagerCallerSession) GetMessageForNodeParticipationSig(message string, nonce *big.Int) ([]byte, []byte, error) {
	return _NodeManager.Contract.GetMessageForNodeParticipationSig(&_NodeManager.CallOpts, message, nonce)
}

// GetMessageForNodeRemovalSig is a free data retrieval call binding the contract method 0x54414df8.
//
// Solidity: function getMessageForNodeRemovalSig(bytes32 nodeToRemoveHashId) view returns(bytes, bytes)
func (_NodeManager *NodeManagerCaller) GetMessageForNodeRemovalSig(opts *bind.CallOpts, nodeToRemoveHashId [32]byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getMessageForNodeRemovalSig", nodeToRemoveHashId)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetMessageForNodeRemovalSig is a free data retrieval call binding the contract method 0x54414df8.
//
// Solidity: function getMessageForNodeRemovalSig(bytes32 nodeToRemoveHashId) view returns(bytes, bytes)
func (_NodeManager *NodeManagerSession) GetMessageForNodeRemovalSig(nodeToRemoveHashId [32]byte) ([]byte, []byte, error) {
	return _NodeManager.Contract.GetMessageForNodeRemovalSig(&_NodeManager.CallOpts, nodeToRemoveHashId)
}

// GetMessageForNodeRemovalSig is a free data retrieval call binding the contract method 0x54414df8.
//
// Solidity: function getMessageForNodeRemovalSig(bytes32 nodeToRemoveHashId) view returns(bytes, bytes)
func (_NodeManager *NodeManagerCallerSession) GetMessageForNodeRemovalSig(nodeToRemoveHashId [32]byte) ([]byte, []byte, error) {
	return _NodeManager.Contract.GetMessageForNodeRemovalSig(&_NodeManager.CallOpts, nodeToRemoveHashId)
}

// GetNodeIdHash is a free data retrieval call binding the contract method 0xfa954f95.
//
// Solidity: function getNodeIdHash(uint256[4] nodeId) pure returns(bytes32)
func (_NodeManager *NodeManagerCaller) GetNodeIdHash(opts *bind.CallOpts, nodeId [4]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "getNodeIdHash", nodeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeIdHash is a free data retrieval call binding the contract method 0xfa954f95.
//
// Solidity: function getNodeIdHash(uint256[4] nodeId) pure returns(bytes32)
func (_NodeManager *NodeManagerSession) GetNodeIdHash(nodeId [4]*big.Int) ([32]byte, error) {
	return _NodeManager.Contract.GetNodeIdHash(&_NodeManager.CallOpts, nodeId)
}

// GetNodeIdHash is a free data retrieval call binding the contract method 0xfa954f95.
//
// Solidity: function getNodeIdHash(uint256[4] nodeId) pure returns(bytes32)
func (_NodeManager *NodeManagerCallerSession) GetNodeIdHash(nodeId [4]*big.Int) ([32]byte, error) {
	return _NodeManager.Contract.GetNodeIdHash(&_NodeManager.CallOpts, nodeId)
}

// IsNodeExists is a free data retrieval call binding the contract method 0xbcb8bb9d.
//
// Solidity: function isNodeExists(uint256[4] nodeId) view returns(bool)
func (_NodeManager *NodeManagerCaller) IsNodeExists(opts *bind.CallOpts, nodeId [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "isNodeExists", nodeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNodeExists is a free data retrieval call binding the contract method 0xbcb8bb9d.
//
// Solidity: function isNodeExists(uint256[4] nodeId) view returns(bool)
func (_NodeManager *NodeManagerSession) IsNodeExists(nodeId [4]*big.Int) (bool, error) {
	return _NodeManager.Contract.IsNodeExists(&_NodeManager.CallOpts, nodeId)
}

// IsNodeExists is a free data retrieval call binding the contract method 0xbcb8bb9d.
//
// Solidity: function isNodeExists(uint256[4] nodeId) view returns(bool)
func (_NodeManager *NodeManagerCallerSession) IsNodeExists(nodeId [4]*big.Int) (bool, error) {
	return _NodeManager.Contract.IsNodeExists(&_NodeManager.CallOpts, nodeId)
}

// IsValidNodeId is a free data retrieval call binding the contract method 0x5a4320e6.
//
// Solidity: function isValidNodeId(uint256[4] nodeId) pure returns(bool)
func (_NodeManager *NodeManagerCaller) IsValidNodeId(opts *bind.CallOpts, nodeId [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "isValidNodeId", nodeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNodeId is a free data retrieval call binding the contract method 0x5a4320e6.
//
// Solidity: function isValidNodeId(uint256[4] nodeId) pure returns(bool)
func (_NodeManager *NodeManagerSession) IsValidNodeId(nodeId [4]*big.Int) (bool, error) {
	return _NodeManager.Contract.IsValidNodeId(&_NodeManager.CallOpts, nodeId)
}

// IsValidNodeId is a free data retrieval call binding the contract method 0x5a4320e6.
//
// Solidity: function isValidNodeId(uint256[4] nodeId) pure returns(bool)
func (_NodeManager *NodeManagerCallerSession) IsValidNodeId(nodeId [4]*big.Int) (bool, error) {
	return _NodeManager.Contract.IsValidNodeId(&_NodeManager.CallOpts, nodeId)
}

// MaxJoinsInTick is a free data retrieval call binding the contract method 0x96dedbfa.
//
// Solidity: function maxJoinsInTick() view returns(uint256)
func (_NodeManager *NodeManagerCaller) MaxJoinsInTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "maxJoinsInTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxJoinsInTick is a free data retrieval call binding the contract method 0x96dedbfa.
//
// Solidity: function maxJoinsInTick() view returns(uint256)
func (_NodeManager *NodeManagerSession) MaxJoinsInTick() (*big.Int, error) {
	return _NodeManager.Contract.MaxJoinsInTick(&_NodeManager.CallOpts)
}

// MaxJoinsInTick is a free data retrieval call binding the contract method 0x96dedbfa.
//
// Solidity: function maxJoinsInTick() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) MaxJoinsInTick() (*big.Int, error) {
	return _NodeManager.Contract.MaxJoinsInTick(&_NodeManager.CallOpts)
}

// NodeData is a free data retrieval call binding the contract method 0xf63553b5.
//
// Solidity: function nodeData(bytes32 ) view returns(uint256 balance, address owner)
func (_NodeManager *NodeManagerCaller) NodeData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Balance *big.Int
	Owner   common.Address
}, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "nodeData", arg0)

	outstruct := new(struct {
		Balance *big.Int
		Owner   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// NodeData is a free data retrieval call binding the contract method 0xf63553b5.
//
// Solidity: function nodeData(bytes32 ) view returns(uint256 balance, address owner)
func (_NodeManager *NodeManagerSession) NodeData(arg0 [32]byte) (struct {
	Balance *big.Int
	Owner   common.Address
}, error) {
	return _NodeManager.Contract.NodeData(&_NodeManager.CallOpts, arg0)
}

// NodeData is a free data retrieval call binding the contract method 0xf63553b5.
//
// Solidity: function nodeData(bytes32 ) view returns(uint256 balance, address owner)
func (_NodeManager *NodeManagerCallerSession) NodeData(arg0 [32]byte) (struct {
	Balance *big.Int
	Owner   common.Address
}, error) {
	return _NodeManager.Contract.NodeData(&_NodeManager.CallOpts, arg0)
}

// NodesJoinedInCurTick is a free data retrieval call binding the contract method 0xde9b4477.
//
// Solidity: function nodesJoinedInCurTick() view returns(uint256)
func (_NodeManager *NodeManagerCaller) NodesJoinedInCurTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "nodesJoinedInCurTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesJoinedInCurTick is a free data retrieval call binding the contract method 0xde9b4477.
//
// Solidity: function nodesJoinedInCurTick() view returns(uint256)
func (_NodeManager *NodeManagerSession) NodesJoinedInCurTick() (*big.Int, error) {
	return _NodeManager.Contract.NodesJoinedInCurTick(&_NodeManager.CallOpts)
}

// NodesJoinedInCurTick is a free data retrieval call binding the contract method 0xde9b4477.
//
// Solidity: function nodesJoinedInCurTick() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) NodesJoinedInCurTick() (*big.Int, error) {
	return _NodeManager.Contract.NodesJoinedInCurTick(&_NodeManager.CallOpts)
}

// NonceForNode is a free data retrieval call binding the contract method 0xff0f98a8.
//
// Solidity: function nonceForNode(bytes32 ) view returns(uint256)
func (_NodeManager *NodeManagerCaller) NonceForNode(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "nonceForNode", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceForNode is a free data retrieval call binding the contract method 0xff0f98a8.
//
// Solidity: function nonceForNode(bytes32 ) view returns(uint256)
func (_NodeManager *NodeManagerSession) NonceForNode(arg0 [32]byte) (*big.Int, error) {
	return _NodeManager.Contract.NonceForNode(&_NodeManager.CallOpts, arg0)
}

// NonceForNode is a free data retrieval call binding the contract method 0xff0f98a8.
//
// Solidity: function nonceForNode(bytes32 ) view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) NonceForNode(arg0 [32]byte) (*big.Int, error) {
	return _NodeManager.Contract.NonceForNode(&_NodeManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManager *NodeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManager *NodeManagerSession) Owner() (common.Address, error) {
	return _NodeManager.Contract.Owner(&_NodeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeManager *NodeManagerCallerSession) Owner() (common.Address, error) {
	return _NodeManager.Contract.Owner(&_NodeManager.CallOpts)
}

// RequiredNodeBalance is a free data retrieval call binding the contract method 0xf6b9e4b6.
//
// Solidity: function requiredNodeBalance() view returns(uint256)
func (_NodeManager *NodeManagerCaller) RequiredNodeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "requiredNodeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredNodeBalance is a free data retrieval call binding the contract method 0xf6b9e4b6.
//
// Solidity: function requiredNodeBalance() view returns(uint256)
func (_NodeManager *NodeManagerSession) RequiredNodeBalance() (*big.Int, error) {
	return _NodeManager.Contract.RequiredNodeBalance(&_NodeManager.CallOpts)
}

// RequiredNodeBalance is a free data retrieval call binding the contract method 0xf6b9e4b6.
//
// Solidity: function requiredNodeBalance() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) RequiredNodeBalance() (*big.Int, error) {
	return _NodeManager.Contract.RequiredNodeBalance(&_NodeManager.CallOpts)
}

// TickDuration is a free data retrieval call binding the contract method 0x37dafa5b.
//
// Solidity: function tickDuration() view returns(uint256)
func (_NodeManager *NodeManagerCaller) TickDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "tickDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickDuration is a free data retrieval call binding the contract method 0x37dafa5b.
//
// Solidity: function tickDuration() view returns(uint256)
func (_NodeManager *NodeManagerSession) TickDuration() (*big.Int, error) {
	return _NodeManager.Contract.TickDuration(&_NodeManager.CallOpts)
}

// TickDuration is a free data retrieval call binding the contract method 0x37dafa5b.
//
// Solidity: function tickDuration() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) TickDuration() (*big.Int, error) {
	return _NodeManager.Contract.TickDuration(&_NodeManager.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NodeManager *NodeManagerCaller) TotalNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeManager.contract.Call(opts, &out, "totalNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NodeManager *NodeManagerSession) TotalNodes() (*big.Int, error) {
	return _NodeManager.Contract.TotalNodes(&_NodeManager.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_NodeManager *NodeManagerCallerSession) TotalNodes() (*big.Int, error) {
	return _NodeManager.Contract.TotalNodes(&_NodeManager.CallOpts)
}

// DepositForNode is a paid mutator transaction binding the contract method 0x6081425f.
//
// Solidity: function depositForNode(uint256[4] nodeId) payable returns()
func (_NodeManager *NodeManagerTransactor) DepositForNode(opts *bind.TransactOpts, nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "depositForNode", nodeId)
}

// DepositForNode is a paid mutator transaction binding the contract method 0x6081425f.
//
// Solidity: function depositForNode(uint256[4] nodeId) payable returns()
func (_NodeManager *NodeManagerSession) DepositForNode(nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.DepositForNode(&_NodeManager.TransactOpts, nodeId)
}

// DepositForNode is a paid mutator transaction binding the contract method 0x6081425f.
//
// Solidity: function depositForNode(uint256[4] nodeId) payable returns()
func (_NodeManager *NodeManagerTransactorSession) DepositForNode(nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.DepositForNode(&_NodeManager.TransactOpts, nodeId)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xde397534.
//
// Solidity: function removeNode(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, bytes32 nodeToRemoveHashId) returns()
func (_NodeManager *NodeManagerTransactor) RemoveNode(opts *bind.TransactOpts, signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, nodeToRemoveHashId [32]byte) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "removeNode", signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, nodeToRemoveHashId)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xde397534.
//
// Solidity: function removeNode(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, bytes32 nodeToRemoveHashId) returns()
func (_NodeManager *NodeManagerSession) RemoveNode(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, nodeToRemoveHashId [32]byte) (*types.Transaction, error) {
	return _NodeManager.Contract.RemoveNode(&_NodeManager.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, nodeToRemoveHashId)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xde397534.
//
// Solidity: function removeNode(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, bytes32 nodeToRemoveHashId) returns()
func (_NodeManager *NodeManagerTransactorSession) RemoveNode(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, nodeToRemoveHashId [32]byte) (*types.Transaction, error) {
	return _NodeManager.Contract.RemoveNode(&_NodeManager.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, nodeToRemoveHashId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeManager *NodeManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeManager *NodeManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeManager.Contract.RenounceOwnership(&_NodeManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeManager *NodeManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeManager.Contract.RenounceOwnership(&_NodeManager.TransactOpts)
}

// SetNewMaxJoinsInTick is a paid mutator transaction binding the contract method 0xf4e7a056.
//
// Solidity: function setNewMaxJoinsInTick(uint256 newMaxJoinsInTick) returns()
func (_NodeManager *NodeManagerTransactor) SetNewMaxJoinsInTick(opts *bind.TransactOpts, newMaxJoinsInTick *big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "setNewMaxJoinsInTick", newMaxJoinsInTick)
}

// SetNewMaxJoinsInTick is a paid mutator transaction binding the contract method 0xf4e7a056.
//
// Solidity: function setNewMaxJoinsInTick(uint256 newMaxJoinsInTick) returns()
func (_NodeManager *NodeManagerSession) SetNewMaxJoinsInTick(newMaxJoinsInTick *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewMaxJoinsInTick(&_NodeManager.TransactOpts, newMaxJoinsInTick)
}

// SetNewMaxJoinsInTick is a paid mutator transaction binding the contract method 0xf4e7a056.
//
// Solidity: function setNewMaxJoinsInTick(uint256 newMaxJoinsInTick) returns()
func (_NodeManager *NodeManagerTransactorSession) SetNewMaxJoinsInTick(newMaxJoinsInTick *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewMaxJoinsInTick(&_NodeManager.TransactOpts, newMaxJoinsInTick)
}

// SetNewRequiredNodeBalance is a paid mutator transaction binding the contract method 0xd48903d0.
//
// Solidity: function setNewRequiredNodeBalance(uint256 newRequiredNodeBalance) returns()
func (_NodeManager *NodeManagerTransactor) SetNewRequiredNodeBalance(opts *bind.TransactOpts, newRequiredNodeBalance *big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "setNewRequiredNodeBalance", newRequiredNodeBalance)
}

// SetNewRequiredNodeBalance is a paid mutator transaction binding the contract method 0xd48903d0.
//
// Solidity: function setNewRequiredNodeBalance(uint256 newRequiredNodeBalance) returns()
func (_NodeManager *NodeManagerSession) SetNewRequiredNodeBalance(newRequiredNodeBalance *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewRequiredNodeBalance(&_NodeManager.TransactOpts, newRequiredNodeBalance)
}

// SetNewRequiredNodeBalance is a paid mutator transaction binding the contract method 0xd48903d0.
//
// Solidity: function setNewRequiredNodeBalance(uint256 newRequiredNodeBalance) returns()
func (_NodeManager *NodeManagerTransactorSession) SetNewRequiredNodeBalance(newRequiredNodeBalance *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewRequiredNodeBalance(&_NodeManager.TransactOpts, newRequiredNodeBalance)
}

// SetNewTickDuration is a paid mutator transaction binding the contract method 0x9e21a658.
//
// Solidity: function setNewTickDuration(uint256 newTickDuration) returns()
func (_NodeManager *NodeManagerTransactor) SetNewTickDuration(opts *bind.TransactOpts, newTickDuration *big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "setNewTickDuration", newTickDuration)
}

// SetNewTickDuration is a paid mutator transaction binding the contract method 0x9e21a658.
//
// Solidity: function setNewTickDuration(uint256 newTickDuration) returns()
func (_NodeManager *NodeManagerSession) SetNewTickDuration(newTickDuration *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewTickDuration(&_NodeManager.TransactOpts, newTickDuration)
}

// SetNewTickDuration is a paid mutator transaction binding the contract method 0x9e21a658.
//
// Solidity: function setNewTickDuration(uint256 newTickDuration) returns()
func (_NodeManager *NodeManagerTransactorSession) SetNewTickDuration(newTickDuration *big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.SetNewTickDuration(&_NodeManager.TransactOpts, newTickDuration)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManager *NodeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManager *NodeManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeManager.Contract.TransferOwnership(&_NodeManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeManager *NodeManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeManager.Contract.TransferOwnership(&_NodeManager.TransactOpts, newOwner)
}

// WithdrawNodeFunds is a paid mutator transaction binding the contract method 0xf392b58b.
//
// Solidity: function withdrawNodeFunds(uint256[4] nodeId) returns()
func (_NodeManager *NodeManagerTransactor) WithdrawNodeFunds(opts *bind.TransactOpts, nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "withdrawNodeFunds", nodeId)
}

// WithdrawNodeFunds is a paid mutator transaction binding the contract method 0xf392b58b.
//
// Solidity: function withdrawNodeFunds(uint256[4] nodeId) returns()
func (_NodeManager *NodeManagerSession) WithdrawNodeFunds(nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.WithdrawNodeFunds(&_NodeManager.TransactOpts, nodeId)
}

// WithdrawNodeFunds is a paid mutator transaction binding the contract method 0xf392b58b.
//
// Solidity: function withdrawNodeFunds(uint256[4] nodeId) returns()
func (_NodeManager *NodeManagerTransactorSession) WithdrawNodeFunds(nodeId [4]*big.Int) (*types.Transaction, error) {
	return _NodeManager.Contract.WithdrawNodeFunds(&_NodeManager.TransactOpts, nodeId)
}

// NodeManagerExtraDepositForNodeIterator is returned from FilterExtraDepositForNode and is used to iterate over the raw logs and unpacked data for ExtraDepositForNode events raised by the NodeManager contract.
type NodeManagerExtraDepositForNodeIterator struct {
	Event *NodeManagerExtraDepositForNode // Event containing the contract specifics and raw log

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
func (it *NodeManagerExtraDepositForNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerExtraDepositForNode)
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
		it.Event = new(NodeManagerExtraDepositForNode)
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
func (it *NodeManagerExtraDepositForNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerExtraDepositForNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerExtraDepositForNode represents a ExtraDepositForNode event raised by the NodeManager contract.
type NodeManagerExtraDepositForNode struct {
	NodeIdHash [32]byte
	Depositor  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExtraDepositForNode is a free log retrieval operation binding the contract event 0xd2f5404eeaa8d6039ed9a0fe016750b2fbba2fe1014878a6344246304404a2ad.
//
// Solidity: event ExtraDepositForNode(bytes32 indexed nodeIdHash, address indexed depositor, uint256 amount)
func (_NodeManager *NodeManagerFilterer) FilterExtraDepositForNode(opts *bind.FilterOpts, nodeIdHash [][32]byte, depositor []common.Address) (*NodeManagerExtraDepositForNodeIterator, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "ExtraDepositForNode", nodeIdHashRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerExtraDepositForNodeIterator{contract: _NodeManager.contract, event: "ExtraDepositForNode", logs: logs, sub: sub}, nil
}

// WatchExtraDepositForNode is a free log subscription operation binding the contract event 0xd2f5404eeaa8d6039ed9a0fe016750b2fbba2fe1014878a6344246304404a2ad.
//
// Solidity: event ExtraDepositForNode(bytes32 indexed nodeIdHash, address indexed depositor, uint256 amount)
func (_NodeManager *NodeManagerFilterer) WatchExtraDepositForNode(opts *bind.WatchOpts, sink chan<- *NodeManagerExtraDepositForNode, nodeIdHash [][32]byte, depositor []common.Address) (event.Subscription, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "ExtraDepositForNode", nodeIdHashRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerExtraDepositForNode)
				if err := _NodeManager.contract.UnpackLog(event, "ExtraDepositForNode", log); err != nil {
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

// ParseExtraDepositForNode is a log parse operation binding the contract event 0xd2f5404eeaa8d6039ed9a0fe016750b2fbba2fe1014878a6344246304404a2ad.
//
// Solidity: event ExtraDepositForNode(bytes32 indexed nodeIdHash, address indexed depositor, uint256 amount)
func (_NodeManager *NodeManagerFilterer) ParseExtraDepositForNode(log types.Log) (*NodeManagerExtraDepositForNode, error) {
	event := new(NodeManagerExtraDepositForNode)
	if err := _NodeManager.contract.UnpackLog(event, "ExtraDepositForNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerMaxJoinsInTickChangedIterator is returned from FilterMaxJoinsInTickChanged and is used to iterate over the raw logs and unpacked data for MaxJoinsInTickChanged events raised by the NodeManager contract.
type NodeManagerMaxJoinsInTickChangedIterator struct {
	Event *NodeManagerMaxJoinsInTickChanged // Event containing the contract specifics and raw log

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
func (it *NodeManagerMaxJoinsInTickChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerMaxJoinsInTickChanged)
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
		it.Event = new(NodeManagerMaxJoinsInTickChanged)
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
func (it *NodeManagerMaxJoinsInTickChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerMaxJoinsInTickChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerMaxJoinsInTickChanged represents a MaxJoinsInTickChanged event raised by the NodeManager contract.
type NodeManagerMaxJoinsInTickChanged struct {
	OldMaxJoinsInTick *big.Int
	NewMaxJoinsInTick *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMaxJoinsInTickChanged is a free log retrieval operation binding the contract event 0xa81ff423c50ddd43fa6a8c14e2d35f6571ce5d3af846bd9a837cbaf369a5149f.
//
// Solidity: event MaxJoinsInTickChanged(uint256 indexed oldMaxJoinsInTick, uint256 indexed newMaxJoinsInTick)
func (_NodeManager *NodeManagerFilterer) FilterMaxJoinsInTickChanged(opts *bind.FilterOpts, oldMaxJoinsInTick []*big.Int, newMaxJoinsInTick []*big.Int) (*NodeManagerMaxJoinsInTickChangedIterator, error) {

	var oldMaxJoinsInTickRule []interface{}
	for _, oldMaxJoinsInTickItem := range oldMaxJoinsInTick {
		oldMaxJoinsInTickRule = append(oldMaxJoinsInTickRule, oldMaxJoinsInTickItem)
	}
	var newMaxJoinsInTickRule []interface{}
	for _, newMaxJoinsInTickItem := range newMaxJoinsInTick {
		newMaxJoinsInTickRule = append(newMaxJoinsInTickRule, newMaxJoinsInTickItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "MaxJoinsInTickChanged", oldMaxJoinsInTickRule, newMaxJoinsInTickRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerMaxJoinsInTickChangedIterator{contract: _NodeManager.contract, event: "MaxJoinsInTickChanged", logs: logs, sub: sub}, nil
}

// WatchMaxJoinsInTickChanged is a free log subscription operation binding the contract event 0xa81ff423c50ddd43fa6a8c14e2d35f6571ce5d3af846bd9a837cbaf369a5149f.
//
// Solidity: event MaxJoinsInTickChanged(uint256 indexed oldMaxJoinsInTick, uint256 indexed newMaxJoinsInTick)
func (_NodeManager *NodeManagerFilterer) WatchMaxJoinsInTickChanged(opts *bind.WatchOpts, sink chan<- *NodeManagerMaxJoinsInTickChanged, oldMaxJoinsInTick []*big.Int, newMaxJoinsInTick []*big.Int) (event.Subscription, error) {

	var oldMaxJoinsInTickRule []interface{}
	for _, oldMaxJoinsInTickItem := range oldMaxJoinsInTick {
		oldMaxJoinsInTickRule = append(oldMaxJoinsInTickRule, oldMaxJoinsInTickItem)
	}
	var newMaxJoinsInTickRule []interface{}
	for _, newMaxJoinsInTickItem := range newMaxJoinsInTick {
		newMaxJoinsInTickRule = append(newMaxJoinsInTickRule, newMaxJoinsInTickItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "MaxJoinsInTickChanged", oldMaxJoinsInTickRule, newMaxJoinsInTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerMaxJoinsInTickChanged)
				if err := _NodeManager.contract.UnpackLog(event, "MaxJoinsInTickChanged", log); err != nil {
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

// ParseMaxJoinsInTickChanged is a log parse operation binding the contract event 0xa81ff423c50ddd43fa6a8c14e2d35f6571ce5d3af846bd9a837cbaf369a5149f.
//
// Solidity: event MaxJoinsInTickChanged(uint256 indexed oldMaxJoinsInTick, uint256 indexed newMaxJoinsInTick)
func (_NodeManager *NodeManagerFilterer) ParseMaxJoinsInTickChanged(log types.Log) (*NodeManagerMaxJoinsInTickChanged, error) {
	event := new(NodeManagerMaxJoinsInTickChanged)
	if err := _NodeManager.contract.UnpackLog(event, "MaxJoinsInTickChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the NodeManager contract.
type NodeManagerNodeAddedIterator struct {
	Event *NodeManagerNodeAdded // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeAdded)
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
		it.Event = new(NodeManagerNodeAdded)
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
func (it *NodeManagerNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeAdded represents a NodeAdded event raised by the NodeManager contract.
type NodeManagerNodeAdded struct {
	NodeIdHash [32]byte
	NodeOwner  common.Address
	TickNum    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xe7aeb8c7da993ec76eb1440800ed15083dbf35b984ff73ec80e8d0cca20d86a6.
//
// Solidity: event NodeAdded(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 tickNum)
func (_NodeManager *NodeManagerFilterer) FilterNodeAdded(opts *bind.FilterOpts, nodeIdHash [][32]byte, nodeOwner []common.Address) (*NodeManagerNodeAddedIterator, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeAdded", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeAddedIterator{contract: _NodeManager.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xe7aeb8c7da993ec76eb1440800ed15083dbf35b984ff73ec80e8d0cca20d86a6.
//
// Solidity: event NodeAdded(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 tickNum)
func (_NodeManager *NodeManagerFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeAdded, nodeIdHash [][32]byte, nodeOwner []common.Address) (event.Subscription, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeAdded", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeAdded)
				if err := _NodeManager.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0xe7aeb8c7da993ec76eb1440800ed15083dbf35b984ff73ec80e8d0cca20d86a6.
//
// Solidity: event NodeAdded(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 tickNum)
func (_NodeManager *NodeManagerFilterer) ParseNodeAdded(log types.Log) (*NodeManagerNodeAdded, error) {
	event := new(NodeManagerNodeAdded)
	if err := _NodeManager.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerNodeLeftIterator is returned from FilterNodeLeft and is used to iterate over the raw logs and unpacked data for NodeLeft events raised by the NodeManager contract.
type NodeManagerNodeLeftIterator struct {
	Event *NodeManagerNodeLeft // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeLeft)
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
		it.Event = new(NodeManagerNodeLeft)
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
func (it *NodeManagerNodeLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeLeft represents a NodeLeft event raised by the NodeManager contract.
type NodeManagerNodeLeft struct {
	NodeIdHash [32]byte
	NodeOwner  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeLeft is a free log retrieval operation binding the contract event 0xb03075d77f987a47c111a292ed269586f40b157471734a6fa420ccdedd400a94.
//
// Solidity: event NodeLeft(bytes32 indexed nodeIdHash, address indexed nodeOwner)
func (_NodeManager *NodeManagerFilterer) FilterNodeLeft(opts *bind.FilterOpts, nodeIdHash [][32]byte, nodeOwner []common.Address) (*NodeManagerNodeLeftIterator, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeLeft", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeLeftIterator{contract: _NodeManager.contract, event: "NodeLeft", logs: logs, sub: sub}, nil
}

// WatchNodeLeft is a free log subscription operation binding the contract event 0xb03075d77f987a47c111a292ed269586f40b157471734a6fa420ccdedd400a94.
//
// Solidity: event NodeLeft(bytes32 indexed nodeIdHash, address indexed nodeOwner)
func (_NodeManager *NodeManagerFilterer) WatchNodeLeft(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeLeft, nodeIdHash [][32]byte, nodeOwner []common.Address) (event.Subscription, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeLeft", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeLeft)
				if err := _NodeManager.contract.UnpackLog(event, "NodeLeft", log); err != nil {
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

// ParseNodeLeft is a log parse operation binding the contract event 0xb03075d77f987a47c111a292ed269586f40b157471734a6fa420ccdedd400a94.
//
// Solidity: event NodeLeft(bytes32 indexed nodeIdHash, address indexed nodeOwner)
func (_NodeManager *NodeManagerFilterer) ParseNodeLeft(log types.Log) (*NodeManagerNodeLeft, error) {
	event := new(NodeManagerNodeLeft)
	if err := _NodeManager.contract.UnpackLog(event, "NodeLeft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the NodeManager contract.
type NodeManagerNodeRemovedIterator struct {
	Event *NodeManagerNodeRemoved // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeRemoved)
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
		it.Event = new(NodeManagerNodeRemoved)
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
func (it *NodeManagerNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeRemoved represents a NodeRemoved event raised by the NodeManager contract.
type NodeManagerNodeRemoved struct {
	NodeIdHash   [32]byte
	NodeOwner    common.Address
	AmountLocked *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xc8f6618824a7ceeff63db2209bda1261c5ff91d1b3b6225a8edaeb0b59d93bca.
//
// Solidity: event NodeRemoved(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 amountLocked)
func (_NodeManager *NodeManagerFilterer) FilterNodeRemoved(opts *bind.FilterOpts, nodeIdHash [][32]byte, nodeOwner []common.Address) (*NodeManagerNodeRemovedIterator, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeRemoved", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeRemovedIterator{contract: _NodeManager.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xc8f6618824a7ceeff63db2209bda1261c5ff91d1b3b6225a8edaeb0b59d93bca.
//
// Solidity: event NodeRemoved(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 amountLocked)
func (_NodeManager *NodeManagerFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeRemoved, nodeIdHash [][32]byte, nodeOwner []common.Address) (event.Subscription, error) {

	var nodeIdHashRule []interface{}
	for _, nodeIdHashItem := range nodeIdHash {
		nodeIdHashRule = append(nodeIdHashRule, nodeIdHashItem)
	}
	var nodeOwnerRule []interface{}
	for _, nodeOwnerItem := range nodeOwner {
		nodeOwnerRule = append(nodeOwnerRule, nodeOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeRemoved", nodeIdHashRule, nodeOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeRemoved)
				if err := _NodeManager.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// ParseNodeRemoved is a log parse operation binding the contract event 0xc8f6618824a7ceeff63db2209bda1261c5ff91d1b3b6225a8edaeb0b59d93bca.
//
// Solidity: event NodeRemoved(bytes32 indexed nodeIdHash, address indexed nodeOwner, uint256 amountLocked)
func (_NodeManager *NodeManagerFilterer) ParseNodeRemoved(log types.Log) (*NodeManagerNodeRemoved, error) {
	event := new(NodeManagerNodeRemoved)
	if err := _NodeManager.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeManager contract.
type NodeManagerOwnershipTransferredIterator struct {
	Event *NodeManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerOwnershipTransferred)
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
		it.Event = new(NodeManagerOwnershipTransferred)
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
func (it *NodeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerOwnershipTransferred represents a OwnershipTransferred event raised by the NodeManager contract.
type NodeManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeManager *NodeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerOwnershipTransferredIterator{contract: _NodeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeManager *NodeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerOwnershipTransferred)
				if err := _NodeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeManager *NodeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*NodeManagerOwnershipTransferred, error) {
	event := new(NodeManagerOwnershipTransferred)
	if err := _NodeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerRequiredNodeBalanceChangedIterator is returned from FilterRequiredNodeBalanceChanged and is used to iterate over the raw logs and unpacked data for RequiredNodeBalanceChanged events raised by the NodeManager contract.
type NodeManagerRequiredNodeBalanceChangedIterator struct {
	Event *NodeManagerRequiredNodeBalanceChanged // Event containing the contract specifics and raw log

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
func (it *NodeManagerRequiredNodeBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerRequiredNodeBalanceChanged)
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
		it.Event = new(NodeManagerRequiredNodeBalanceChanged)
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
func (it *NodeManagerRequiredNodeBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerRequiredNodeBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerRequiredNodeBalanceChanged represents a RequiredNodeBalanceChanged event raised by the NodeManager contract.
type NodeManagerRequiredNodeBalanceChanged struct {
	OldRequiredNodeBalance *big.Int
	NewRequiredNodeBalance *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRequiredNodeBalanceChanged is a free log retrieval operation binding the contract event 0x846d8a38dca81b94cb0a160b0fa12d0f013228d6c586b68e16dc9ca390ca16de.
//
// Solidity: event RequiredNodeBalanceChanged(uint256 indexed oldRequiredNodeBalance, uint256 indexed newRequiredNodeBalance)
func (_NodeManager *NodeManagerFilterer) FilterRequiredNodeBalanceChanged(opts *bind.FilterOpts, oldRequiredNodeBalance []*big.Int, newRequiredNodeBalance []*big.Int) (*NodeManagerRequiredNodeBalanceChangedIterator, error) {

	var oldRequiredNodeBalanceRule []interface{}
	for _, oldRequiredNodeBalanceItem := range oldRequiredNodeBalance {
		oldRequiredNodeBalanceRule = append(oldRequiredNodeBalanceRule, oldRequiredNodeBalanceItem)
	}
	var newRequiredNodeBalanceRule []interface{}
	for _, newRequiredNodeBalanceItem := range newRequiredNodeBalance {
		newRequiredNodeBalanceRule = append(newRequiredNodeBalanceRule, newRequiredNodeBalanceItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "RequiredNodeBalanceChanged", oldRequiredNodeBalanceRule, newRequiredNodeBalanceRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerRequiredNodeBalanceChangedIterator{contract: _NodeManager.contract, event: "RequiredNodeBalanceChanged", logs: logs, sub: sub}, nil
}

// WatchRequiredNodeBalanceChanged is a free log subscription operation binding the contract event 0x846d8a38dca81b94cb0a160b0fa12d0f013228d6c586b68e16dc9ca390ca16de.
//
// Solidity: event RequiredNodeBalanceChanged(uint256 indexed oldRequiredNodeBalance, uint256 indexed newRequiredNodeBalance)
func (_NodeManager *NodeManagerFilterer) WatchRequiredNodeBalanceChanged(opts *bind.WatchOpts, sink chan<- *NodeManagerRequiredNodeBalanceChanged, oldRequiredNodeBalance []*big.Int, newRequiredNodeBalance []*big.Int) (event.Subscription, error) {

	var oldRequiredNodeBalanceRule []interface{}
	for _, oldRequiredNodeBalanceItem := range oldRequiredNodeBalance {
		oldRequiredNodeBalanceRule = append(oldRequiredNodeBalanceRule, oldRequiredNodeBalanceItem)
	}
	var newRequiredNodeBalanceRule []interface{}
	for _, newRequiredNodeBalanceItem := range newRequiredNodeBalance {
		newRequiredNodeBalanceRule = append(newRequiredNodeBalanceRule, newRequiredNodeBalanceItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "RequiredNodeBalanceChanged", oldRequiredNodeBalanceRule, newRequiredNodeBalanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerRequiredNodeBalanceChanged)
				if err := _NodeManager.contract.UnpackLog(event, "RequiredNodeBalanceChanged", log); err != nil {
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

// ParseRequiredNodeBalanceChanged is a log parse operation binding the contract event 0x846d8a38dca81b94cb0a160b0fa12d0f013228d6c586b68e16dc9ca390ca16de.
//
// Solidity: event RequiredNodeBalanceChanged(uint256 indexed oldRequiredNodeBalance, uint256 indexed newRequiredNodeBalance)
func (_NodeManager *NodeManagerFilterer) ParseRequiredNodeBalanceChanged(log types.Log) (*NodeManagerRequiredNodeBalanceChanged, error) {
	event := new(NodeManagerRequiredNodeBalanceChanged)
	if err := _NodeManager.contract.UnpackLog(event, "RequiredNodeBalanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeManagerTickDurationChangedIterator is returned from FilterTickDurationChanged and is used to iterate over the raw logs and unpacked data for TickDurationChanged events raised by the NodeManager contract.
type NodeManagerTickDurationChangedIterator struct {
	Event *NodeManagerTickDurationChanged // Event containing the contract specifics and raw log

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
func (it *NodeManagerTickDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerTickDurationChanged)
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
		it.Event = new(NodeManagerTickDurationChanged)
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
func (it *NodeManagerTickDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerTickDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerTickDurationChanged represents a TickDurationChanged event raised by the NodeManager contract.
type NodeManagerTickDurationChanged struct {
	OldTickDuration *big.Int
	NewTickDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTickDurationChanged is a free log retrieval operation binding the contract event 0xe1eee662f666eb47b78bcc0c334a0d11d4bd1b5208b5a94c638e8dc117a6f5dc.
//
// Solidity: event TickDurationChanged(uint256 indexed oldTickDuration, uint256 indexed newTickDuration)
func (_NodeManager *NodeManagerFilterer) FilterTickDurationChanged(opts *bind.FilterOpts, oldTickDuration []*big.Int, newTickDuration []*big.Int) (*NodeManagerTickDurationChangedIterator, error) {

	var oldTickDurationRule []interface{}
	for _, oldTickDurationItem := range oldTickDuration {
		oldTickDurationRule = append(oldTickDurationRule, oldTickDurationItem)
	}
	var newTickDurationRule []interface{}
	for _, newTickDurationItem := range newTickDuration {
		newTickDurationRule = append(newTickDurationRule, newTickDurationItem)
	}

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "TickDurationChanged", oldTickDurationRule, newTickDurationRule)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTickDurationChangedIterator{contract: _NodeManager.contract, event: "TickDurationChanged", logs: logs, sub: sub}, nil
}

// WatchTickDurationChanged is a free log subscription operation binding the contract event 0xe1eee662f666eb47b78bcc0c334a0d11d4bd1b5208b5a94c638e8dc117a6f5dc.
//
// Solidity: event TickDurationChanged(uint256 indexed oldTickDuration, uint256 indexed newTickDuration)
func (_NodeManager *NodeManagerFilterer) WatchTickDurationChanged(opts *bind.WatchOpts, sink chan<- *NodeManagerTickDurationChanged, oldTickDuration []*big.Int, newTickDuration []*big.Int) (event.Subscription, error) {

	var oldTickDurationRule []interface{}
	for _, oldTickDurationItem := range oldTickDuration {
		oldTickDurationRule = append(oldTickDurationRule, oldTickDurationItem)
	}
	var newTickDurationRule []interface{}
	for _, newTickDurationItem := range newTickDuration {
		newTickDurationRule = append(newTickDurationRule, newTickDurationItem)
	}

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "TickDurationChanged", oldTickDurationRule, newTickDurationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerTickDurationChanged)
				if err := _NodeManager.contract.UnpackLog(event, "TickDurationChanged", log); err != nil {
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

// ParseTickDurationChanged is a log parse operation binding the contract event 0xe1eee662f666eb47b78bcc0c334a0d11d4bd1b5208b5a94c638e8dc117a6f5dc.
//
// Solidity: event TickDurationChanged(uint256 indexed oldTickDuration, uint256 indexed newTickDuration)
func (_NodeManager *NodeManagerFilterer) ParseTickDurationChanged(log types.Log) (*NodeManagerTickDurationChanged, error) {
	event := new(NodeManagerTickDurationChanged)
	if err := _NodeManager.contract.UnpackLog(event, "TickDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
