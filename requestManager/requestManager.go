// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package requestManager

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

// RequestManagerMetaData contains all meta data concerning the RequestManager contract.
var RequestManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_serviceCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountToLockForEachRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_requestTimeout\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_defaultMaxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_nodeManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmountToLockForEachRequest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmountToLockForEachRequest\",\"type\":\"uint256\"}],\"name\":\"AmountToLockForEachRequestChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldDefaultMaxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newDefaultMaxGasLimit\",\"type\":\"uint32\"}],\"name\":\"DefaultMaxGasLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositedForRequestTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldNodeManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newNodeManager\",\"type\":\"address\"}],\"name\":\"NodeManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"}],\"name\":\"RequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestTotalCost\",\"type\":\"uint256\"}],\"name\":\"RequestFulFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldRequestTimeout\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRequestTimeout\",\"type\":\"uint32\"}],\"name\":\"RequestTimeoutChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"RequestTopicCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldServiceCost\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newServiceCost\",\"type\":\"uint256\"}],\"name\":\"ServiceCostChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMaxGasLimit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMaxGasLimit\",\"type\":\"address\"}],\"name\":\"TopicRequestMaxGasLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequestTopicFunds\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amountToLockForEachRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"name\":\"createRequestTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultMaxGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"}],\"name\":\"depositForRequestTopic\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"signersPubKey\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"fullFillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"getMessageForRequestFulFillmentSig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestNum\",\"type\":\"uint256\"}],\"name\":\"getRequestData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumRequestManager.RequestStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"}],\"name\":\"getRequestTopicData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"hashToPoint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isUsedData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"publicKey\",\"type\":\"uint256[4]\"}],\"name\":\"isValidPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeManagerIns\",\"outputs\":[{\"internalType\":\"contractINodeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestTimeout\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"}],\"name\":\"requestTopicAvailableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newRequestTimeout\",\"type\":\"uint32\"}],\"name\":\"seNewtRequestTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmountToLockForEachRequest\",\"type\":\"uint256\"}],\"name\":\"setNewAmountToLockForEachRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newDefaultMaxGasLimit\",\"type\":\"uint32\"}],\"name\":\"setNewDefaultMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newNodeManager\",\"type\":\"address\"}],\"name\":\"setNewNodeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newServiceCost\",\"type\":\"uint256\"}],\"name\":\"setNewServiceCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newMaxGasLimit\",\"type\":\"uint32\"}],\"name\":\"setTopicRequestNewMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"pubkeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkeyG2\",\"type\":\"uint256[4]\"}],\"name\":\"verifyHelpedAggregationPublicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"aggPubkeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysG2\",\"type\":\"uint256[4][]\"}],\"name\":\"verifyHelpedAggregationPublicKeysMultiple\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRequestTopicFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RequestManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestManagerMetaData.ABI instead.
var RequestManagerABI = RequestManagerMetaData.ABI

// RequestManager is an auto generated Go binding around an Ethereum contract.
type RequestManager struct {
	RequestManagerCaller     // Read-only binding to the contract
	RequestManagerTransactor // Write-only binding to the contract
	RequestManagerFilterer   // Log filterer for contract events
}

// RequestManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestManagerSession struct {
	Contract     *RequestManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestManagerCallerSession struct {
	Contract *RequestManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RequestManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestManagerTransactorSession struct {
	Contract     *RequestManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RequestManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestManagerRaw struct {
	Contract *RequestManager // Generic contract binding to access the raw methods on
}

// RequestManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestManagerCallerRaw struct {
	Contract *RequestManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RequestManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestManagerTransactorRaw struct {
	Contract *RequestManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestManager creates a new instance of RequestManager, bound to a specific deployed contract.
func NewRequestManager(address common.Address, backend bind.ContractBackend) (*RequestManager, error) {
	contract, err := bindRequestManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestManager{RequestManagerCaller: RequestManagerCaller{contract: contract}, RequestManagerTransactor: RequestManagerTransactor{contract: contract}, RequestManagerFilterer: RequestManagerFilterer{contract: contract}}, nil
}

// NewRequestManagerCaller creates a new read-only instance of RequestManager, bound to a specific deployed contract.
func NewRequestManagerCaller(address common.Address, caller bind.ContractCaller) (*RequestManagerCaller, error) {
	contract, err := bindRequestManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestManagerCaller{contract: contract}, nil
}

// NewRequestManagerTransactor creates a new write-only instance of RequestManager, bound to a specific deployed contract.
func NewRequestManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestManagerTransactor, error) {
	contract, err := bindRequestManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestManagerTransactor{contract: contract}, nil
}

// NewRequestManagerFilterer creates a new log filterer instance of RequestManager, bound to a specific deployed contract.
func NewRequestManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestManagerFilterer, error) {
	contract, err := bindRequestManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestManagerFilterer{contract: contract}, nil
}

// bindRequestManager binds a generic wrapper to an already deployed contract.
func bindRequestManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestManager *RequestManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestManager.Contract.RequestManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestManager *RequestManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestManager.Contract.RequestManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestManager *RequestManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestManager.Contract.RequestManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestManager *RequestManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestManager *RequestManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestManager *RequestManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestManager.Contract.contract.Transact(opts, method, params...)
}

// AmountToLockForEachRequest is a free data retrieval call binding the contract method 0xb035880f.
//
// Solidity: function amountToLockForEachRequest() view returns(uint256)
func (_RequestManager *RequestManagerCaller) AmountToLockForEachRequest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "amountToLockForEachRequest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToLockForEachRequest is a free data retrieval call binding the contract method 0xb035880f.
//
// Solidity: function amountToLockForEachRequest() view returns(uint256)
func (_RequestManager *RequestManagerSession) AmountToLockForEachRequest() (*big.Int, error) {
	return _RequestManager.Contract.AmountToLockForEachRequest(&_RequestManager.CallOpts)
}

// AmountToLockForEachRequest is a free data retrieval call binding the contract method 0xb035880f.
//
// Solidity: function amountToLockForEachRequest() view returns(uint256)
func (_RequestManager *RequestManagerCallerSession) AmountToLockForEachRequest() (*big.Int, error) {
	return _RequestManager.Contract.AmountToLockForEachRequest(&_RequestManager.CallOpts)
}

// DefaultMaxGasLimit is a free data retrieval call binding the contract method 0x91a9ccb3.
//
// Solidity: function defaultMaxGasLimit() view returns(uint32)
func (_RequestManager *RequestManagerCaller) DefaultMaxGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "defaultMaxGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DefaultMaxGasLimit is a free data retrieval call binding the contract method 0x91a9ccb3.
//
// Solidity: function defaultMaxGasLimit() view returns(uint32)
func (_RequestManager *RequestManagerSession) DefaultMaxGasLimit() (uint32, error) {
	return _RequestManager.Contract.DefaultMaxGasLimit(&_RequestManager.CallOpts)
}

// DefaultMaxGasLimit is a free data retrieval call binding the contract method 0x91a9ccb3.
//
// Solidity: function defaultMaxGasLimit() view returns(uint32)
func (_RequestManager *RequestManagerCallerSession) DefaultMaxGasLimit() (uint32, error) {
	return _RequestManager.Contract.DefaultMaxGasLimit(&_RequestManager.CallOpts)
}

// GetMessageForRequestFulFillmentSig is a free data retrieval call binding the contract method 0x93c57d15.
//
// Solidity: function getMessageForRequestFulFillmentSig(address recipient, uint256 requestNum, uint256 data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerCaller) GetMessageForRequestFulFillmentSig(opts *bind.CallOpts, recipient common.Address, requestNum *big.Int, data *big.Int) ([]byte, []byte, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "getMessageForRequestFulFillmentSig", recipient, requestNum, data)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetMessageForRequestFulFillmentSig is a free data retrieval call binding the contract method 0x93c57d15.
//
// Solidity: function getMessageForRequestFulFillmentSig(address recipient, uint256 requestNum, uint256 data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerSession) GetMessageForRequestFulFillmentSig(recipient common.Address, requestNum *big.Int, data *big.Int) ([]byte, []byte, error) {
	return _RequestManager.Contract.GetMessageForRequestFulFillmentSig(&_RequestManager.CallOpts, recipient, requestNum, data)
}

// GetMessageForRequestFulFillmentSig is a free data retrieval call binding the contract method 0x93c57d15.
//
// Solidity: function getMessageForRequestFulFillmentSig(address recipient, uint256 requestNum, uint256 data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerCallerSession) GetMessageForRequestFulFillmentSig(recipient common.Address, requestNum *big.Int, data *big.Int) ([]byte, []byte, error) {
	return _RequestManager.Contract.GetMessageForRequestFulFillmentSig(&_RequestManager.CallOpts, recipient, requestNum, data)
}

// GetRequestData is a free data retrieval call binding the contract method 0x8af1af96.
//
// Solidity: function getRequestData(address requestBeneficiary, uint256 requestNum) view returns(uint256, uint8)
func (_RequestManager *RequestManagerCaller) GetRequestData(opts *bind.CallOpts, requestBeneficiary common.Address, requestNum *big.Int) (*big.Int, uint8, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "getRequestData", requestBeneficiary, requestNum)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetRequestData is a free data retrieval call binding the contract method 0x8af1af96.
//
// Solidity: function getRequestData(address requestBeneficiary, uint256 requestNum) view returns(uint256, uint8)
func (_RequestManager *RequestManagerSession) GetRequestData(requestBeneficiary common.Address, requestNum *big.Int) (*big.Int, uint8, error) {
	return _RequestManager.Contract.GetRequestData(&_RequestManager.CallOpts, requestBeneficiary, requestNum)
}

// GetRequestData is a free data retrieval call binding the contract method 0x8af1af96.
//
// Solidity: function getRequestData(address requestBeneficiary, uint256 requestNum) view returns(uint256, uint8)
func (_RequestManager *RequestManagerCallerSession) GetRequestData(requestBeneficiary common.Address, requestNum *big.Int) (*big.Int, uint8, error) {
	return _RequestManager.Contract.GetRequestData(&_RequestManager.CallOpts, requestBeneficiary, requestNum)
}

// GetRequestTopicData is a free data retrieval call binding the contract method 0xc791bff0.
//
// Solidity: function getRequestTopicData(address requestBeneficiary) view returns(address, uint256, uint256, uint256, uint32)
func (_RequestManager *RequestManagerCaller) GetRequestTopicData(opts *bind.CallOpts, requestBeneficiary common.Address) (common.Address, *big.Int, *big.Int, *big.Int, uint32, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "getRequestTopicData", requestBeneficiary)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return out0, out1, out2, out3, out4, err

}

// GetRequestTopicData is a free data retrieval call binding the contract method 0xc791bff0.
//
// Solidity: function getRequestTopicData(address requestBeneficiary) view returns(address, uint256, uint256, uint256, uint32)
func (_RequestManager *RequestManagerSession) GetRequestTopicData(requestBeneficiary common.Address) (common.Address, *big.Int, *big.Int, *big.Int, uint32, error) {
	return _RequestManager.Contract.GetRequestTopicData(&_RequestManager.CallOpts, requestBeneficiary)
}

// GetRequestTopicData is a free data retrieval call binding the contract method 0xc791bff0.
//
// Solidity: function getRequestTopicData(address requestBeneficiary) view returns(address, uint256, uint256, uint256, uint32)
func (_RequestManager *RequestManagerCallerSession) GetRequestTopicData(requestBeneficiary common.Address) (common.Address, *big.Int, *big.Int, *big.Int, uint32, error) {
	return _RequestManager.Contract.GetRequestTopicData(&_RequestManager.CallOpts, requestBeneficiary)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerCaller) HashToPoint(opts *bind.CallOpts, data []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "hashToPoint", data)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerSession) HashToPoint(data []byte) ([]byte, []byte, error) {
	return _RequestManager.Contract.HashToPoint(&_RequestManager.CallOpts, data)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestManager *RequestManagerCallerSession) HashToPoint(data []byte) ([]byte, []byte, error) {
	return _RequestManager.Contract.HashToPoint(&_RequestManager.CallOpts, data)
}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestManager *RequestManagerCaller) IsUsedData(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "isUsedData", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestManager *RequestManagerSession) IsUsedData(arg0 *big.Int) (bool, error) {
	return _RequestManager.Contract.IsUsedData(&_RequestManager.CallOpts, arg0)
}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestManager *RequestManagerCallerSession) IsUsedData(arg0 *big.Int) (bool, error) {
	return _RequestManager.Contract.IsUsedData(&_RequestManager.CallOpts, arg0)
}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestManager *RequestManagerCaller) IsValidPublicKey(opts *bind.CallOpts, publicKey [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "isValidPublicKey", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestManager *RequestManagerSession) IsValidPublicKey(publicKey [4]*big.Int) (bool, error) {
	return _RequestManager.Contract.IsValidPublicKey(&_RequestManager.CallOpts, publicKey)
}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestManager *RequestManagerCallerSession) IsValidPublicKey(publicKey [4]*big.Int) (bool, error) {
	return _RequestManager.Contract.IsValidPublicKey(&_RequestManager.CallOpts, publicKey)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestManager *RequestManagerCaller) IsValidSignature(opts *bind.CallOpts, signature [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "isValidSignature", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestManager *RequestManagerSession) IsValidSignature(signature [2]*big.Int) (bool, error) {
	return _RequestManager.Contract.IsValidSignature(&_RequestManager.CallOpts, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestManager *RequestManagerCallerSession) IsValidSignature(signature [2]*big.Int) (bool, error) {
	return _RequestManager.Contract.IsValidSignature(&_RequestManager.CallOpts, signature)
}

// NodeManagerIns is a free data retrieval call binding the contract method 0xef8e493a.
//
// Solidity: function nodeManagerIns() view returns(address)
func (_RequestManager *RequestManagerCaller) NodeManagerIns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "nodeManagerIns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeManagerIns is a free data retrieval call binding the contract method 0xef8e493a.
//
// Solidity: function nodeManagerIns() view returns(address)
func (_RequestManager *RequestManagerSession) NodeManagerIns() (common.Address, error) {
	return _RequestManager.Contract.NodeManagerIns(&_RequestManager.CallOpts)
}

// NodeManagerIns is a free data retrieval call binding the contract method 0xef8e493a.
//
// Solidity: function nodeManagerIns() view returns(address)
func (_RequestManager *RequestManagerCallerSession) NodeManagerIns() (common.Address, error) {
	return _RequestManager.Contract.NodeManagerIns(&_RequestManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestManager *RequestManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestManager *RequestManagerSession) Owner() (common.Address, error) {
	return _RequestManager.Contract.Owner(&_RequestManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestManager *RequestManagerCallerSession) Owner() (common.Address, error) {
	return _RequestManager.Contract.Owner(&_RequestManager.CallOpts)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint32)
func (_RequestManager *RequestManagerCaller) RequestTimeout(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "requestTimeout")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint32)
func (_RequestManager *RequestManagerSession) RequestTimeout() (uint32, error) {
	return _RequestManager.Contract.RequestTimeout(&_RequestManager.CallOpts)
}

// RequestTimeout is a free data retrieval call binding the contract method 0x3f20b4c9.
//
// Solidity: function requestTimeout() view returns(uint32)
func (_RequestManager *RequestManagerCallerSession) RequestTimeout() (uint32, error) {
	return _RequestManager.Contract.RequestTimeout(&_RequestManager.CallOpts)
}

// RequestTopicAvailableFunds is a free data retrieval call binding the contract method 0x917ec80f.
//
// Solidity: function requestTopicAvailableFunds(address requestBeneficiary) view returns(uint256)
func (_RequestManager *RequestManagerCaller) RequestTopicAvailableFunds(opts *bind.CallOpts, requestBeneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "requestTopicAvailableFunds", requestBeneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestTopicAvailableFunds is a free data retrieval call binding the contract method 0x917ec80f.
//
// Solidity: function requestTopicAvailableFunds(address requestBeneficiary) view returns(uint256)
func (_RequestManager *RequestManagerSession) RequestTopicAvailableFunds(requestBeneficiary common.Address) (*big.Int, error) {
	return _RequestManager.Contract.RequestTopicAvailableFunds(&_RequestManager.CallOpts, requestBeneficiary)
}

// RequestTopicAvailableFunds is a free data retrieval call binding the contract method 0x917ec80f.
//
// Solidity: function requestTopicAvailableFunds(address requestBeneficiary) view returns(uint256)
func (_RequestManager *RequestManagerCallerSession) RequestTopicAvailableFunds(requestBeneficiary common.Address) (*big.Int, error) {
	return _RequestManager.Contract.RequestTopicAvailableFunds(&_RequestManager.CallOpts, requestBeneficiary)
}

// ServiceCost is a free data retrieval call binding the contract method 0x1c47b5c6.
//
// Solidity: function serviceCost() view returns(uint256)
func (_RequestManager *RequestManagerCaller) ServiceCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "serviceCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceCost is a free data retrieval call binding the contract method 0x1c47b5c6.
//
// Solidity: function serviceCost() view returns(uint256)
func (_RequestManager *RequestManagerSession) ServiceCost() (*big.Int, error) {
	return _RequestManager.Contract.ServiceCost(&_RequestManager.CallOpts)
}

// ServiceCost is a free data retrieval call binding the contract method 0x1c47b5c6.
//
// Solidity: function serviceCost() view returns(uint256)
func (_RequestManager *RequestManagerCallerSession) ServiceCost() (*big.Int, error) {
	return _RequestManager.Contract.ServiceCost(&_RequestManager.CallOpts)
}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestManager *RequestManagerCaller) VerifyHelpedAggregationPublicKeys(opts *bind.CallOpts, pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "verifyHelpedAggregationPublicKeys", pubkeyG1, pubkeyG2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestManager *RequestManagerSession) VerifyHelpedAggregationPublicKeys(pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	return _RequestManager.Contract.VerifyHelpedAggregationPublicKeys(&_RequestManager.CallOpts, pubkeyG1, pubkeyG2)
}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestManager *RequestManagerCallerSession) VerifyHelpedAggregationPublicKeys(pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	return _RequestManager.Contract.VerifyHelpedAggregationPublicKeys(&_RequestManager.CallOpts, pubkeyG1, pubkeyG2)
}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestManager *RequestManagerCaller) VerifyHelpedAggregationPublicKeysMultiple(opts *bind.CallOpts, aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestManager.contract.Call(opts, &out, "verifyHelpedAggregationPublicKeysMultiple", aggPubkeyG1, pubkeysG2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestManager *RequestManagerSession) VerifyHelpedAggregationPublicKeysMultiple(aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	return _RequestManager.Contract.VerifyHelpedAggregationPublicKeysMultiple(&_RequestManager.CallOpts, aggPubkeyG1, pubkeysG2)
}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestManager *RequestManagerCallerSession) VerifyHelpedAggregationPublicKeysMultiple(aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	return _RequestManager.Contract.VerifyHelpedAggregationPublicKeysMultiple(&_RequestManager.CallOpts, aggPubkeyG1, pubkeysG2)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x887344ff.
//
// Solidity: function cancelRequest(address requestBeneficiary, uint256 requestNum) returns()
func (_RequestManager *RequestManagerTransactor) CancelRequest(opts *bind.TransactOpts, requestBeneficiary common.Address, requestNum *big.Int) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "cancelRequest", requestBeneficiary, requestNum)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x887344ff.
//
// Solidity: function cancelRequest(address requestBeneficiary, uint256 requestNum) returns()
func (_RequestManager *RequestManagerSession) CancelRequest(requestBeneficiary common.Address, requestNum *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.CancelRequest(&_RequestManager.TransactOpts, requestBeneficiary, requestNum)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x887344ff.
//
// Solidity: function cancelRequest(address requestBeneficiary, uint256 requestNum) returns()
func (_RequestManager *RequestManagerTransactorSession) CancelRequest(requestBeneficiary common.Address, requestNum *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.CancelRequest(&_RequestManager.TransactOpts, requestBeneficiary, requestNum)
}

// CreateRequest is a paid mutator transaction binding the contract method 0xc9579078.
//
// Solidity: function createRequest() returns()
func (_RequestManager *RequestManagerTransactor) CreateRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "createRequest")
}

// CreateRequest is a paid mutator transaction binding the contract method 0xc9579078.
//
// Solidity: function createRequest() returns()
func (_RequestManager *RequestManagerSession) CreateRequest() (*types.Transaction, error) {
	return _RequestManager.Contract.CreateRequest(&_RequestManager.TransactOpts)
}

// CreateRequest is a paid mutator transaction binding the contract method 0xc9579078.
//
// Solidity: function createRequest() returns()
func (_RequestManager *RequestManagerTransactorSession) CreateRequest() (*types.Transaction, error) {
	return _RequestManager.Contract.CreateRequest(&_RequestManager.TransactOpts)
}

// CreateRequestTopic is a paid mutator transaction binding the contract method 0x46e0f6d7.
//
// Solidity: function createRequestTopic(address requestBeneficiary, uint32 maxGasLimit) payable returns()
func (_RequestManager *RequestManagerTransactor) CreateRequestTopic(opts *bind.TransactOpts, requestBeneficiary common.Address, maxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "createRequestTopic", requestBeneficiary, maxGasLimit)
}

// CreateRequestTopic is a paid mutator transaction binding the contract method 0x46e0f6d7.
//
// Solidity: function createRequestTopic(address requestBeneficiary, uint32 maxGasLimit) payable returns()
func (_RequestManager *RequestManagerSession) CreateRequestTopic(requestBeneficiary common.Address, maxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.CreateRequestTopic(&_RequestManager.TransactOpts, requestBeneficiary, maxGasLimit)
}

// CreateRequestTopic is a paid mutator transaction binding the contract method 0x46e0f6d7.
//
// Solidity: function createRequestTopic(address requestBeneficiary, uint32 maxGasLimit) payable returns()
func (_RequestManager *RequestManagerTransactorSession) CreateRequestTopic(requestBeneficiary common.Address, maxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.CreateRequestTopic(&_RequestManager.TransactOpts, requestBeneficiary, maxGasLimit)
}

// DepositForRequestTopic is a paid mutator transaction binding the contract method 0x9edfeec1.
//
// Solidity: function depositForRequestTopic(address requestBeneficiary) payable returns()
func (_RequestManager *RequestManagerTransactor) DepositForRequestTopic(opts *bind.TransactOpts, requestBeneficiary common.Address) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "depositForRequestTopic", requestBeneficiary)
}

// DepositForRequestTopic is a paid mutator transaction binding the contract method 0x9edfeec1.
//
// Solidity: function depositForRequestTopic(address requestBeneficiary) payable returns()
func (_RequestManager *RequestManagerSession) DepositForRequestTopic(requestBeneficiary common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.DepositForRequestTopic(&_RequestManager.TransactOpts, requestBeneficiary)
}

// DepositForRequestTopic is a paid mutator transaction binding the contract method 0x9edfeec1.
//
// Solidity: function depositForRequestTopic(address requestBeneficiary) payable returns()
func (_RequestManager *RequestManagerTransactorSession) DepositForRequestTopic(requestBeneficiary common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.DepositForRequestTopic(&_RequestManager.TransactOpts, requestBeneficiary)
}

// FullFillRequest is a paid mutator transaction binding the contract method 0xa2cb753e.
//
// Solidity: function fullFillRequest(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address requestBeneficiary, uint256 requestNum, uint256 data) returns()
func (_RequestManager *RequestManagerTransactor) FullFillRequest(opts *bind.TransactOpts, signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, requestBeneficiary common.Address, requestNum *big.Int, data *big.Int) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "fullFillRequest", signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, requestBeneficiary, requestNum, data)
}

// FullFillRequest is a paid mutator transaction binding the contract method 0xa2cb753e.
//
// Solidity: function fullFillRequest(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address requestBeneficiary, uint256 requestNum, uint256 data) returns()
func (_RequestManager *RequestManagerSession) FullFillRequest(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, requestBeneficiary common.Address, requestNum *big.Int, data *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.FullFillRequest(&_RequestManager.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, requestBeneficiary, requestNum, data)
}

// FullFillRequest is a paid mutator transaction binding the contract method 0xa2cb753e.
//
// Solidity: function fullFillRequest(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address requestBeneficiary, uint256 requestNum, uint256 data) returns()
func (_RequestManager *RequestManagerTransactorSession) FullFillRequest(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, requestBeneficiary common.Address, requestNum *big.Int, data *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.FullFillRequest(&_RequestManager.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, requestBeneficiary, requestNum, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestManager *RequestManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestManager *RequestManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestManager.Contract.RenounceOwnership(&_RequestManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestManager *RequestManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestManager.Contract.RenounceOwnership(&_RequestManager.TransactOpts)
}

// SeNewtRequestTimeout is a paid mutator transaction binding the contract method 0xf854df3a.
//
// Solidity: function seNewtRequestTimeout(uint32 newRequestTimeout) returns()
func (_RequestManager *RequestManagerTransactor) SeNewtRequestTimeout(opts *bind.TransactOpts, newRequestTimeout uint32) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "seNewtRequestTimeout", newRequestTimeout)
}

// SeNewtRequestTimeout is a paid mutator transaction binding the contract method 0xf854df3a.
//
// Solidity: function seNewtRequestTimeout(uint32 newRequestTimeout) returns()
func (_RequestManager *RequestManagerSession) SeNewtRequestTimeout(newRequestTimeout uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SeNewtRequestTimeout(&_RequestManager.TransactOpts, newRequestTimeout)
}

// SeNewtRequestTimeout is a paid mutator transaction binding the contract method 0xf854df3a.
//
// Solidity: function seNewtRequestTimeout(uint32 newRequestTimeout) returns()
func (_RequestManager *RequestManagerTransactorSession) SeNewtRequestTimeout(newRequestTimeout uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SeNewtRequestTimeout(&_RequestManager.TransactOpts, newRequestTimeout)
}

// SetNewAmountToLockForEachRequest is a paid mutator transaction binding the contract method 0xbd57b043.
//
// Solidity: function setNewAmountToLockForEachRequest(uint256 newAmountToLockForEachRequest) returns()
func (_RequestManager *RequestManagerTransactor) SetNewAmountToLockForEachRequest(opts *bind.TransactOpts, newAmountToLockForEachRequest *big.Int) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "setNewAmountToLockForEachRequest", newAmountToLockForEachRequest)
}

// SetNewAmountToLockForEachRequest is a paid mutator transaction binding the contract method 0xbd57b043.
//
// Solidity: function setNewAmountToLockForEachRequest(uint256 newAmountToLockForEachRequest) returns()
func (_RequestManager *RequestManagerSession) SetNewAmountToLockForEachRequest(newAmountToLockForEachRequest *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewAmountToLockForEachRequest(&_RequestManager.TransactOpts, newAmountToLockForEachRequest)
}

// SetNewAmountToLockForEachRequest is a paid mutator transaction binding the contract method 0xbd57b043.
//
// Solidity: function setNewAmountToLockForEachRequest(uint256 newAmountToLockForEachRequest) returns()
func (_RequestManager *RequestManagerTransactorSession) SetNewAmountToLockForEachRequest(newAmountToLockForEachRequest *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewAmountToLockForEachRequest(&_RequestManager.TransactOpts, newAmountToLockForEachRequest)
}

// SetNewDefaultMaxGasLimit is a paid mutator transaction binding the contract method 0x02385adf.
//
// Solidity: function setNewDefaultMaxGasLimit(uint32 newDefaultMaxGasLimit) returns()
func (_RequestManager *RequestManagerTransactor) SetNewDefaultMaxGasLimit(opts *bind.TransactOpts, newDefaultMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "setNewDefaultMaxGasLimit", newDefaultMaxGasLimit)
}

// SetNewDefaultMaxGasLimit is a paid mutator transaction binding the contract method 0x02385adf.
//
// Solidity: function setNewDefaultMaxGasLimit(uint32 newDefaultMaxGasLimit) returns()
func (_RequestManager *RequestManagerSession) SetNewDefaultMaxGasLimit(newDefaultMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewDefaultMaxGasLimit(&_RequestManager.TransactOpts, newDefaultMaxGasLimit)
}

// SetNewDefaultMaxGasLimit is a paid mutator transaction binding the contract method 0x02385adf.
//
// Solidity: function setNewDefaultMaxGasLimit(uint32 newDefaultMaxGasLimit) returns()
func (_RequestManager *RequestManagerTransactorSession) SetNewDefaultMaxGasLimit(newDefaultMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewDefaultMaxGasLimit(&_RequestManager.TransactOpts, newDefaultMaxGasLimit)
}

// SetNewNodeManager is a paid mutator transaction binding the contract method 0x31ffcb2a.
//
// Solidity: function setNewNodeManager(address newNodeManager) returns()
func (_RequestManager *RequestManagerTransactor) SetNewNodeManager(opts *bind.TransactOpts, newNodeManager common.Address) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "setNewNodeManager", newNodeManager)
}

// SetNewNodeManager is a paid mutator transaction binding the contract method 0x31ffcb2a.
//
// Solidity: function setNewNodeManager(address newNodeManager) returns()
func (_RequestManager *RequestManagerSession) SetNewNodeManager(newNodeManager common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewNodeManager(&_RequestManager.TransactOpts, newNodeManager)
}

// SetNewNodeManager is a paid mutator transaction binding the contract method 0x31ffcb2a.
//
// Solidity: function setNewNodeManager(address newNodeManager) returns()
func (_RequestManager *RequestManagerTransactorSession) SetNewNodeManager(newNodeManager common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewNodeManager(&_RequestManager.TransactOpts, newNodeManager)
}

// SetNewServiceCost is a paid mutator transaction binding the contract method 0xb0a142e5.
//
// Solidity: function setNewServiceCost(uint256 newServiceCost) returns()
func (_RequestManager *RequestManagerTransactor) SetNewServiceCost(opts *bind.TransactOpts, newServiceCost *big.Int) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "setNewServiceCost", newServiceCost)
}

// SetNewServiceCost is a paid mutator transaction binding the contract method 0xb0a142e5.
//
// Solidity: function setNewServiceCost(uint256 newServiceCost) returns()
func (_RequestManager *RequestManagerSession) SetNewServiceCost(newServiceCost *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewServiceCost(&_RequestManager.TransactOpts, newServiceCost)
}

// SetNewServiceCost is a paid mutator transaction binding the contract method 0xb0a142e5.
//
// Solidity: function setNewServiceCost(uint256 newServiceCost) returns()
func (_RequestManager *RequestManagerTransactorSession) SetNewServiceCost(newServiceCost *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.SetNewServiceCost(&_RequestManager.TransactOpts, newServiceCost)
}

// SetTopicRequestNewMaxGasLimit is a paid mutator transaction binding the contract method 0xe9691003.
//
// Solidity: function setTopicRequestNewMaxGasLimit(uint32 newMaxGasLimit) returns()
func (_RequestManager *RequestManagerTransactor) SetTopicRequestNewMaxGasLimit(opts *bind.TransactOpts, newMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "setTopicRequestNewMaxGasLimit", newMaxGasLimit)
}

// SetTopicRequestNewMaxGasLimit is a paid mutator transaction binding the contract method 0xe9691003.
//
// Solidity: function setTopicRequestNewMaxGasLimit(uint32 newMaxGasLimit) returns()
func (_RequestManager *RequestManagerSession) SetTopicRequestNewMaxGasLimit(newMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SetTopicRequestNewMaxGasLimit(&_RequestManager.TransactOpts, newMaxGasLimit)
}

// SetTopicRequestNewMaxGasLimit is a paid mutator transaction binding the contract method 0xe9691003.
//
// Solidity: function setTopicRequestNewMaxGasLimit(uint32 newMaxGasLimit) returns()
func (_RequestManager *RequestManagerTransactorSession) SetTopicRequestNewMaxGasLimit(newMaxGasLimit uint32) (*types.Transaction, error) {
	return _RequestManager.Contract.SetTopicRequestNewMaxGasLimit(&_RequestManager.TransactOpts, newMaxGasLimit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestManager *RequestManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestManager *RequestManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.TransferOwnership(&_RequestManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestManager *RequestManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RequestManager.Contract.TransferOwnership(&_RequestManager.TransactOpts, newOwner)
}

// WithdrawRequestTopicFunds is a paid mutator transaction binding the contract method 0xd1a19979.
//
// Solidity: function withdrawRequestTopicFunds(address requestBeneficiary, uint256 amount) payable returns()
func (_RequestManager *RequestManagerTransactor) WithdrawRequestTopicFunds(opts *bind.TransactOpts, requestBeneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RequestManager.contract.Transact(opts, "withdrawRequestTopicFunds", requestBeneficiary, amount)
}

// WithdrawRequestTopicFunds is a paid mutator transaction binding the contract method 0xd1a19979.
//
// Solidity: function withdrawRequestTopicFunds(address requestBeneficiary, uint256 amount) payable returns()
func (_RequestManager *RequestManagerSession) WithdrawRequestTopicFunds(requestBeneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.WithdrawRequestTopicFunds(&_RequestManager.TransactOpts, requestBeneficiary, amount)
}

// WithdrawRequestTopicFunds is a paid mutator transaction binding the contract method 0xd1a19979.
//
// Solidity: function withdrawRequestTopicFunds(address requestBeneficiary, uint256 amount) payable returns()
func (_RequestManager *RequestManagerTransactorSession) WithdrawRequestTopicFunds(requestBeneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RequestManager.Contract.WithdrawRequestTopicFunds(&_RequestManager.TransactOpts, requestBeneficiary, amount)
}

// RequestManagerAmountToLockForEachRequestChangedIterator is returned from FilterAmountToLockForEachRequestChanged and is used to iterate over the raw logs and unpacked data for AmountToLockForEachRequestChanged events raised by the RequestManager contract.
type RequestManagerAmountToLockForEachRequestChangedIterator struct {
	Event *RequestManagerAmountToLockForEachRequestChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerAmountToLockForEachRequestChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerAmountToLockForEachRequestChanged)
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
		it.Event = new(RequestManagerAmountToLockForEachRequestChanged)
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
func (it *RequestManagerAmountToLockForEachRequestChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerAmountToLockForEachRequestChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerAmountToLockForEachRequestChanged represents a AmountToLockForEachRequestChanged event raised by the RequestManager contract.
type RequestManagerAmountToLockForEachRequestChanged struct {
	OldAmountToLockForEachRequest *big.Int
	NewAmountToLockForEachRequest *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterAmountToLockForEachRequestChanged is a free log retrieval operation binding the contract event 0x33bb42f83c8795adefaafbbee880ce0c393a6d8f595abb9f4ba4bf2fdd2a4b22.
//
// Solidity: event AmountToLockForEachRequestChanged(uint256 oldAmountToLockForEachRequest, uint256 newAmountToLockForEachRequest)
func (_RequestManager *RequestManagerFilterer) FilterAmountToLockForEachRequestChanged(opts *bind.FilterOpts) (*RequestManagerAmountToLockForEachRequestChangedIterator, error) {

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "AmountToLockForEachRequestChanged")
	if err != nil {
		return nil, err
	}
	return &RequestManagerAmountToLockForEachRequestChangedIterator{contract: _RequestManager.contract, event: "AmountToLockForEachRequestChanged", logs: logs, sub: sub}, nil
}

// WatchAmountToLockForEachRequestChanged is a free log subscription operation binding the contract event 0x33bb42f83c8795adefaafbbee880ce0c393a6d8f595abb9f4ba4bf2fdd2a4b22.
//
// Solidity: event AmountToLockForEachRequestChanged(uint256 oldAmountToLockForEachRequest, uint256 newAmountToLockForEachRequest)
func (_RequestManager *RequestManagerFilterer) WatchAmountToLockForEachRequestChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerAmountToLockForEachRequestChanged) (event.Subscription, error) {

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "AmountToLockForEachRequestChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerAmountToLockForEachRequestChanged)
				if err := _RequestManager.contract.UnpackLog(event, "AmountToLockForEachRequestChanged", log); err != nil {
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

// ParseAmountToLockForEachRequestChanged is a log parse operation binding the contract event 0x33bb42f83c8795adefaafbbee880ce0c393a6d8f595abb9f4ba4bf2fdd2a4b22.
//
// Solidity: event AmountToLockForEachRequestChanged(uint256 oldAmountToLockForEachRequest, uint256 newAmountToLockForEachRequest)
func (_RequestManager *RequestManagerFilterer) ParseAmountToLockForEachRequestChanged(log types.Log) (*RequestManagerAmountToLockForEachRequestChanged, error) {
	event := new(RequestManagerAmountToLockForEachRequestChanged)
	if err := _RequestManager.contract.UnpackLog(event, "AmountToLockForEachRequestChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerDefaultMaxGasLimitChangedIterator is returned from FilterDefaultMaxGasLimitChanged and is used to iterate over the raw logs and unpacked data for DefaultMaxGasLimitChanged events raised by the RequestManager contract.
type RequestManagerDefaultMaxGasLimitChangedIterator struct {
	Event *RequestManagerDefaultMaxGasLimitChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerDefaultMaxGasLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerDefaultMaxGasLimitChanged)
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
		it.Event = new(RequestManagerDefaultMaxGasLimitChanged)
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
func (it *RequestManagerDefaultMaxGasLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerDefaultMaxGasLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerDefaultMaxGasLimitChanged represents a DefaultMaxGasLimitChanged event raised by the RequestManager contract.
type RequestManagerDefaultMaxGasLimitChanged struct {
	OldDefaultMaxGasLimit uint32
	NewDefaultMaxGasLimit uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDefaultMaxGasLimitChanged is a free log retrieval operation binding the contract event 0x2ed2752067b53f0b11e2489a62ec48cbca6b4ede4be6a72be34bfc442d6bdd82.
//
// Solidity: event DefaultMaxGasLimitChanged(uint32 oldDefaultMaxGasLimit, uint32 newDefaultMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) FilterDefaultMaxGasLimitChanged(opts *bind.FilterOpts) (*RequestManagerDefaultMaxGasLimitChangedIterator, error) {

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "DefaultMaxGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return &RequestManagerDefaultMaxGasLimitChangedIterator{contract: _RequestManager.contract, event: "DefaultMaxGasLimitChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultMaxGasLimitChanged is a free log subscription operation binding the contract event 0x2ed2752067b53f0b11e2489a62ec48cbca6b4ede4be6a72be34bfc442d6bdd82.
//
// Solidity: event DefaultMaxGasLimitChanged(uint32 oldDefaultMaxGasLimit, uint32 newDefaultMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) WatchDefaultMaxGasLimitChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerDefaultMaxGasLimitChanged) (event.Subscription, error) {

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "DefaultMaxGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerDefaultMaxGasLimitChanged)
				if err := _RequestManager.contract.UnpackLog(event, "DefaultMaxGasLimitChanged", log); err != nil {
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

// ParseDefaultMaxGasLimitChanged is a log parse operation binding the contract event 0x2ed2752067b53f0b11e2489a62ec48cbca6b4ede4be6a72be34bfc442d6bdd82.
//
// Solidity: event DefaultMaxGasLimitChanged(uint32 oldDefaultMaxGasLimit, uint32 newDefaultMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) ParseDefaultMaxGasLimitChanged(log types.Log) (*RequestManagerDefaultMaxGasLimitChanged, error) {
	event := new(RequestManagerDefaultMaxGasLimitChanged)
	if err := _RequestManager.contract.UnpackLog(event, "DefaultMaxGasLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerDepositedForRequestTopicIterator is returned from FilterDepositedForRequestTopic and is used to iterate over the raw logs and unpacked data for DepositedForRequestTopic events raised by the RequestManager contract.
type RequestManagerDepositedForRequestTopicIterator struct {
	Event *RequestManagerDepositedForRequestTopic // Event containing the contract specifics and raw log

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
func (it *RequestManagerDepositedForRequestTopicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerDepositedForRequestTopic)
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
		it.Event = new(RequestManagerDepositedForRequestTopic)
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
func (it *RequestManagerDepositedForRequestTopicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerDepositedForRequestTopicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerDepositedForRequestTopic represents a DepositedForRequestTopic event raised by the RequestManager contract.
type RequestManagerDepositedForRequestTopic struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositedForRequestTopic is a free log retrieval operation binding the contract event 0x6a495c19bf6acafb6141071eca8adc31ef6e431325493affdbbe87a2d077a6ca.
//
// Solidity: event DepositedForRequestTopic(address indexed depositor, uint256 amount)
func (_RequestManager *RequestManagerFilterer) FilterDepositedForRequestTopic(opts *bind.FilterOpts, depositor []common.Address) (*RequestManagerDepositedForRequestTopicIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "DepositedForRequestTopic", depositorRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerDepositedForRequestTopicIterator{contract: _RequestManager.contract, event: "DepositedForRequestTopic", logs: logs, sub: sub}, nil
}

// WatchDepositedForRequestTopic is a free log subscription operation binding the contract event 0x6a495c19bf6acafb6141071eca8adc31ef6e431325493affdbbe87a2d077a6ca.
//
// Solidity: event DepositedForRequestTopic(address indexed depositor, uint256 amount)
func (_RequestManager *RequestManagerFilterer) WatchDepositedForRequestTopic(opts *bind.WatchOpts, sink chan<- *RequestManagerDepositedForRequestTopic, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "DepositedForRequestTopic", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerDepositedForRequestTopic)
				if err := _RequestManager.contract.UnpackLog(event, "DepositedForRequestTopic", log); err != nil {
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

// ParseDepositedForRequestTopic is a log parse operation binding the contract event 0x6a495c19bf6acafb6141071eca8adc31ef6e431325493affdbbe87a2d077a6ca.
//
// Solidity: event DepositedForRequestTopic(address indexed depositor, uint256 amount)
func (_RequestManager *RequestManagerFilterer) ParseDepositedForRequestTopic(log types.Log) (*RequestManagerDepositedForRequestTopic, error) {
	event := new(RequestManagerDepositedForRequestTopic)
	if err := _RequestManager.contract.UnpackLog(event, "DepositedForRequestTopic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerNodeManagerChangedIterator is returned from FilterNodeManagerChanged and is used to iterate over the raw logs and unpacked data for NodeManagerChanged events raised by the RequestManager contract.
type RequestManagerNodeManagerChangedIterator struct {
	Event *RequestManagerNodeManagerChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerNodeManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerNodeManagerChanged)
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
		it.Event = new(RequestManagerNodeManagerChanged)
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
func (it *RequestManagerNodeManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerNodeManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerNodeManagerChanged represents a NodeManagerChanged event raised by the RequestManager contract.
type RequestManagerNodeManagerChanged struct {
	OldNodeManager common.Address
	NewNodeManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeManagerChanged is a free log retrieval operation binding the contract event 0x61c078703028a9ecfa3f88e7d06a4c88a0952ddd43bbbb19c22a630292ba3f0e.
//
// Solidity: event NodeManagerChanged(address indexed oldNodeManager, address indexed newNodeManager)
func (_RequestManager *RequestManagerFilterer) FilterNodeManagerChanged(opts *bind.FilterOpts, oldNodeManager []common.Address, newNodeManager []common.Address) (*RequestManagerNodeManagerChangedIterator, error) {

	var oldNodeManagerRule []interface{}
	for _, oldNodeManagerItem := range oldNodeManager {
		oldNodeManagerRule = append(oldNodeManagerRule, oldNodeManagerItem)
	}
	var newNodeManagerRule []interface{}
	for _, newNodeManagerItem := range newNodeManager {
		newNodeManagerRule = append(newNodeManagerRule, newNodeManagerItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "NodeManagerChanged", oldNodeManagerRule, newNodeManagerRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerNodeManagerChangedIterator{contract: _RequestManager.contract, event: "NodeManagerChanged", logs: logs, sub: sub}, nil
}

// WatchNodeManagerChanged is a free log subscription operation binding the contract event 0x61c078703028a9ecfa3f88e7d06a4c88a0952ddd43bbbb19c22a630292ba3f0e.
//
// Solidity: event NodeManagerChanged(address indexed oldNodeManager, address indexed newNodeManager)
func (_RequestManager *RequestManagerFilterer) WatchNodeManagerChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerNodeManagerChanged, oldNodeManager []common.Address, newNodeManager []common.Address) (event.Subscription, error) {

	var oldNodeManagerRule []interface{}
	for _, oldNodeManagerItem := range oldNodeManager {
		oldNodeManagerRule = append(oldNodeManagerRule, oldNodeManagerItem)
	}
	var newNodeManagerRule []interface{}
	for _, newNodeManagerItem := range newNodeManager {
		newNodeManagerRule = append(newNodeManagerRule, newNodeManagerItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "NodeManagerChanged", oldNodeManagerRule, newNodeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerNodeManagerChanged)
				if err := _RequestManager.contract.UnpackLog(event, "NodeManagerChanged", log); err != nil {
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

// ParseNodeManagerChanged is a log parse operation binding the contract event 0x61c078703028a9ecfa3f88e7d06a4c88a0952ddd43bbbb19c22a630292ba3f0e.
//
// Solidity: event NodeManagerChanged(address indexed oldNodeManager, address indexed newNodeManager)
func (_RequestManager *RequestManagerFilterer) ParseNodeManagerChanged(log types.Log) (*RequestManagerNodeManagerChanged, error) {
	event := new(RequestManagerNodeManagerChanged)
	if err := _RequestManager.contract.UnpackLog(event, "NodeManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RequestManager contract.
type RequestManagerOwnershipTransferredIterator struct {
	Event *RequestManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RequestManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerOwnershipTransferred)
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
		it.Event = new(RequestManagerOwnershipTransferred)
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
func (it *RequestManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerOwnershipTransferred represents a OwnershipTransferred event raised by the RequestManager contract.
type RequestManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestManager *RequestManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RequestManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerOwnershipTransferredIterator{contract: _RequestManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestManager *RequestManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RequestManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerOwnershipTransferred)
				if err := _RequestManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RequestManager *RequestManagerFilterer) ParseOwnershipTransferred(log types.Log) (*RequestManagerOwnershipTransferred, error) {
	event := new(RequestManagerOwnershipTransferred)
	if err := _RequestManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerRequestCanceledIterator is returned from FilterRequestCanceled and is used to iterate over the raw logs and unpacked data for RequestCanceled events raised by the RequestManager contract.
type RequestManagerRequestCanceledIterator struct {
	Event *RequestManagerRequestCanceled // Event containing the contract specifics and raw log

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
func (it *RequestManagerRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerRequestCanceled)
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
		it.Event = new(RequestManagerRequestCanceled)
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
func (it *RequestManagerRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerRequestCanceled represents a RequestCanceled event raised by the RequestManager contract.
type RequestManagerRequestCanceled struct {
	Receiver   common.Address
	RequestNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestCanceled is a free log retrieval operation binding the contract event 0x21a680835504a0149cc7bf56bae8db68f89e6a3b8b87888390a4bfeef107f237.
//
// Solidity: event RequestCanceled(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) FilterRequestCanceled(opts *bind.FilterOpts, receiver []common.Address, requestNum []*big.Int) (*RequestManagerRequestCanceledIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "RequestCanceled", receiverRule, requestNumRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerRequestCanceledIterator{contract: _RequestManager.contract, event: "RequestCanceled", logs: logs, sub: sub}, nil
}

// WatchRequestCanceled is a free log subscription operation binding the contract event 0x21a680835504a0149cc7bf56bae8db68f89e6a3b8b87888390a4bfeef107f237.
//
// Solidity: event RequestCanceled(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) WatchRequestCanceled(opts *bind.WatchOpts, sink chan<- *RequestManagerRequestCanceled, receiver []common.Address, requestNum []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "RequestCanceled", receiverRule, requestNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerRequestCanceled)
				if err := _RequestManager.contract.UnpackLog(event, "RequestCanceled", log); err != nil {
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

// ParseRequestCanceled is a log parse operation binding the contract event 0x21a680835504a0149cc7bf56bae8db68f89e6a3b8b87888390a4bfeef107f237.
//
// Solidity: event RequestCanceled(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) ParseRequestCanceled(log types.Log) (*RequestManagerRequestCanceled, error) {
	event := new(RequestManagerRequestCanceled)
	if err := _RequestManager.contract.UnpackLog(event, "RequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RequestManager contract.
type RequestManagerRequestCreatedIterator struct {
	Event *RequestManagerRequestCreated // Event containing the contract specifics and raw log

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
func (it *RequestManagerRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerRequestCreated)
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
		it.Event = new(RequestManagerRequestCreated)
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
func (it *RequestManagerRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerRequestCreated represents a RequestCreated event raised by the RequestManager contract.
type RequestManagerRequestCreated struct {
	Receiver   common.Address
	RequestNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x85dcc3449f98075b485886e476487433e044f2db02a65b5bf5d29e92392b5c2c.
//
// Solidity: event RequestCreated(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) FilterRequestCreated(opts *bind.FilterOpts, receiver []common.Address, requestNum []*big.Int) (*RequestManagerRequestCreatedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "RequestCreated", receiverRule, requestNumRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerRequestCreatedIterator{contract: _RequestManager.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x85dcc3449f98075b485886e476487433e044f2db02a65b5bf5d29e92392b5c2c.
//
// Solidity: event RequestCreated(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RequestManagerRequestCreated, receiver []common.Address, requestNum []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "RequestCreated", receiverRule, requestNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerRequestCreated)
				if err := _RequestManager.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0x85dcc3449f98075b485886e476487433e044f2db02a65b5bf5d29e92392b5c2c.
//
// Solidity: event RequestCreated(address indexed receiver, uint256 indexed requestNum)
func (_RequestManager *RequestManagerFilterer) ParseRequestCreated(log types.Log) (*RequestManagerRequestCreated, error) {
	event := new(RequestManagerRequestCreated)
	if err := _RequestManager.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerRequestFulFilledIterator is returned from FilterRequestFulFilled and is used to iterate over the raw logs and unpacked data for RequestFulFilled events raised by the RequestManager contract.
type RequestManagerRequestFulFilledIterator struct {
	Event *RequestManagerRequestFulFilled // Event containing the contract specifics and raw log

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
func (it *RequestManagerRequestFulFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerRequestFulFilled)
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
		it.Event = new(RequestManagerRequestFulFilled)
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
func (it *RequestManagerRequestFulFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerRequestFulFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerRequestFulFilled represents a RequestFulFilled event raised by the RequestManager contract.
type RequestManagerRequestFulFilled struct {
	Receiver         common.Address
	RequestNum       *big.Int
	Data             *big.Int
	RequestTotalCost *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestFulFilled is a free log retrieval operation binding the contract event 0x67e969a0bb29dc1e551daf4d597a3ce553c08b462c4095214cd805cdc061f10b.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed requestNum, uint256 indexed data, uint256 requestTotalCost)
func (_RequestManager *RequestManagerFilterer) FilterRequestFulFilled(opts *bind.FilterOpts, receiver []common.Address, requestNum []*big.Int, data []*big.Int) (*RequestManagerRequestFulFilledIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "RequestFulFilled", receiverRule, requestNumRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerRequestFulFilledIterator{contract: _RequestManager.contract, event: "RequestFulFilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulFilled is a free log subscription operation binding the contract event 0x67e969a0bb29dc1e551daf4d597a3ce553c08b462c4095214cd805cdc061f10b.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed requestNum, uint256 indexed data, uint256 requestTotalCost)
func (_RequestManager *RequestManagerFilterer) WatchRequestFulFilled(opts *bind.WatchOpts, sink chan<- *RequestManagerRequestFulFilled, receiver []common.Address, requestNum []*big.Int, data []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var requestNumRule []interface{}
	for _, requestNumItem := range requestNum {
		requestNumRule = append(requestNumRule, requestNumItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "RequestFulFilled", receiverRule, requestNumRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerRequestFulFilled)
				if err := _RequestManager.contract.UnpackLog(event, "RequestFulFilled", log); err != nil {
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

// ParseRequestFulFilled is a log parse operation binding the contract event 0x67e969a0bb29dc1e551daf4d597a3ce553c08b462c4095214cd805cdc061f10b.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed requestNum, uint256 indexed data, uint256 requestTotalCost)
func (_RequestManager *RequestManagerFilterer) ParseRequestFulFilled(log types.Log) (*RequestManagerRequestFulFilled, error) {
	event := new(RequestManagerRequestFulFilled)
	if err := _RequestManager.contract.UnpackLog(event, "RequestFulFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerRequestTimeoutChangedIterator is returned from FilterRequestTimeoutChanged and is used to iterate over the raw logs and unpacked data for RequestTimeoutChanged events raised by the RequestManager contract.
type RequestManagerRequestTimeoutChangedIterator struct {
	Event *RequestManagerRequestTimeoutChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerRequestTimeoutChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerRequestTimeoutChanged)
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
		it.Event = new(RequestManagerRequestTimeoutChanged)
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
func (it *RequestManagerRequestTimeoutChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerRequestTimeoutChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerRequestTimeoutChanged represents a RequestTimeoutChanged event raised by the RequestManager contract.
type RequestManagerRequestTimeoutChanged struct {
	OldRequestTimeout uint32
	NewRequestTimeout uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRequestTimeoutChanged is a free log retrieval operation binding the contract event 0x1f05198eedff5b8adc4418ed51108975b6c6614dd7125a56a62509785caeb75e.
//
// Solidity: event RequestTimeoutChanged(uint32 oldRequestTimeout, uint32 newRequestTimeout)
func (_RequestManager *RequestManagerFilterer) FilterRequestTimeoutChanged(opts *bind.FilterOpts) (*RequestManagerRequestTimeoutChangedIterator, error) {

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "RequestTimeoutChanged")
	if err != nil {
		return nil, err
	}
	return &RequestManagerRequestTimeoutChangedIterator{contract: _RequestManager.contract, event: "RequestTimeoutChanged", logs: logs, sub: sub}, nil
}

// WatchRequestTimeoutChanged is a free log subscription operation binding the contract event 0x1f05198eedff5b8adc4418ed51108975b6c6614dd7125a56a62509785caeb75e.
//
// Solidity: event RequestTimeoutChanged(uint32 oldRequestTimeout, uint32 newRequestTimeout)
func (_RequestManager *RequestManagerFilterer) WatchRequestTimeoutChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerRequestTimeoutChanged) (event.Subscription, error) {

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "RequestTimeoutChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerRequestTimeoutChanged)
				if err := _RequestManager.contract.UnpackLog(event, "RequestTimeoutChanged", log); err != nil {
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

// ParseRequestTimeoutChanged is a log parse operation binding the contract event 0x1f05198eedff5b8adc4418ed51108975b6c6614dd7125a56a62509785caeb75e.
//
// Solidity: event RequestTimeoutChanged(uint32 oldRequestTimeout, uint32 newRequestTimeout)
func (_RequestManager *RequestManagerFilterer) ParseRequestTimeoutChanged(log types.Log) (*RequestManagerRequestTimeoutChanged, error) {
	event := new(RequestManagerRequestTimeoutChanged)
	if err := _RequestManager.contract.UnpackLog(event, "RequestTimeoutChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerRequestTopicCreatedIterator is returned from FilterRequestTopicCreated and is used to iterate over the raw logs and unpacked data for RequestTopicCreated events raised by the RequestManager contract.
type RequestManagerRequestTopicCreatedIterator struct {
	Event *RequestManagerRequestTopicCreated // Event containing the contract specifics and raw log

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
func (it *RequestManagerRequestTopicCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerRequestTopicCreated)
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
		it.Event = new(RequestManagerRequestTopicCreated)
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
func (it *RequestManagerRequestTopicCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerRequestTopicCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerRequestTopicCreated represents a RequestTopicCreated event raised by the RequestManager contract.
type RequestManagerRequestTopicCreated struct {
	RequestBeneficiary common.Address
	Creator            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRequestTopicCreated is a free log retrieval operation binding the contract event 0xe882f1672e5c5dbf926ba7866e593663620171ef677628b67fd64870083a4870.
//
// Solidity: event RequestTopicCreated(address indexed requestBeneficiary, address indexed creator)
func (_RequestManager *RequestManagerFilterer) FilterRequestTopicCreated(opts *bind.FilterOpts, requestBeneficiary []common.Address, creator []common.Address) (*RequestManagerRequestTopicCreatedIterator, error) {

	var requestBeneficiaryRule []interface{}
	for _, requestBeneficiaryItem := range requestBeneficiary {
		requestBeneficiaryRule = append(requestBeneficiaryRule, requestBeneficiaryItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "RequestTopicCreated", requestBeneficiaryRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerRequestTopicCreatedIterator{contract: _RequestManager.contract, event: "RequestTopicCreated", logs: logs, sub: sub}, nil
}

// WatchRequestTopicCreated is a free log subscription operation binding the contract event 0xe882f1672e5c5dbf926ba7866e593663620171ef677628b67fd64870083a4870.
//
// Solidity: event RequestTopicCreated(address indexed requestBeneficiary, address indexed creator)
func (_RequestManager *RequestManagerFilterer) WatchRequestTopicCreated(opts *bind.WatchOpts, sink chan<- *RequestManagerRequestTopicCreated, requestBeneficiary []common.Address, creator []common.Address) (event.Subscription, error) {

	var requestBeneficiaryRule []interface{}
	for _, requestBeneficiaryItem := range requestBeneficiary {
		requestBeneficiaryRule = append(requestBeneficiaryRule, requestBeneficiaryItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "RequestTopicCreated", requestBeneficiaryRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerRequestTopicCreated)
				if err := _RequestManager.contract.UnpackLog(event, "RequestTopicCreated", log); err != nil {
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

// ParseRequestTopicCreated is a log parse operation binding the contract event 0xe882f1672e5c5dbf926ba7866e593663620171ef677628b67fd64870083a4870.
//
// Solidity: event RequestTopicCreated(address indexed requestBeneficiary, address indexed creator)
func (_RequestManager *RequestManagerFilterer) ParseRequestTopicCreated(log types.Log) (*RequestManagerRequestTopicCreated, error) {
	event := new(RequestManagerRequestTopicCreated)
	if err := _RequestManager.contract.UnpackLog(event, "RequestTopicCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerServiceCostChangedIterator is returned from FilterServiceCostChanged and is used to iterate over the raw logs and unpacked data for ServiceCostChanged events raised by the RequestManager contract.
type RequestManagerServiceCostChangedIterator struct {
	Event *RequestManagerServiceCostChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerServiceCostChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerServiceCostChanged)
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
		it.Event = new(RequestManagerServiceCostChanged)
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
func (it *RequestManagerServiceCostChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerServiceCostChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerServiceCostChanged represents a ServiceCostChanged event raised by the RequestManager contract.
type RequestManagerServiceCostChanged struct {
	OldServiceCost *big.Int
	NewServiceCost *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterServiceCostChanged is a free log retrieval operation binding the contract event 0x2f16e3bdc4a9598d18ce012294fbb3af468bd1b6c4273b93b26a033d6ee8fe72.
//
// Solidity: event ServiceCostChanged(uint256 indexed oldServiceCost, uint256 indexed newServiceCost)
func (_RequestManager *RequestManagerFilterer) FilterServiceCostChanged(opts *bind.FilterOpts, oldServiceCost []*big.Int, newServiceCost []*big.Int) (*RequestManagerServiceCostChangedIterator, error) {

	var oldServiceCostRule []interface{}
	for _, oldServiceCostItem := range oldServiceCost {
		oldServiceCostRule = append(oldServiceCostRule, oldServiceCostItem)
	}
	var newServiceCostRule []interface{}
	for _, newServiceCostItem := range newServiceCost {
		newServiceCostRule = append(newServiceCostRule, newServiceCostItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "ServiceCostChanged", oldServiceCostRule, newServiceCostRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerServiceCostChangedIterator{contract: _RequestManager.contract, event: "ServiceCostChanged", logs: logs, sub: sub}, nil
}

// WatchServiceCostChanged is a free log subscription operation binding the contract event 0x2f16e3bdc4a9598d18ce012294fbb3af468bd1b6c4273b93b26a033d6ee8fe72.
//
// Solidity: event ServiceCostChanged(uint256 indexed oldServiceCost, uint256 indexed newServiceCost)
func (_RequestManager *RequestManagerFilterer) WatchServiceCostChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerServiceCostChanged, oldServiceCost []*big.Int, newServiceCost []*big.Int) (event.Subscription, error) {

	var oldServiceCostRule []interface{}
	for _, oldServiceCostItem := range oldServiceCost {
		oldServiceCostRule = append(oldServiceCostRule, oldServiceCostItem)
	}
	var newServiceCostRule []interface{}
	for _, newServiceCostItem := range newServiceCost {
		newServiceCostRule = append(newServiceCostRule, newServiceCostItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "ServiceCostChanged", oldServiceCostRule, newServiceCostRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerServiceCostChanged)
				if err := _RequestManager.contract.UnpackLog(event, "ServiceCostChanged", log); err != nil {
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

// ParseServiceCostChanged is a log parse operation binding the contract event 0x2f16e3bdc4a9598d18ce012294fbb3af468bd1b6c4273b93b26a033d6ee8fe72.
//
// Solidity: event ServiceCostChanged(uint256 indexed oldServiceCost, uint256 indexed newServiceCost)
func (_RequestManager *RequestManagerFilterer) ParseServiceCostChanged(log types.Log) (*RequestManagerServiceCostChanged, error) {
	event := new(RequestManagerServiceCostChanged)
	if err := _RequestManager.contract.UnpackLog(event, "ServiceCostChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerTopicRequestMaxGasLimitChangedIterator is returned from FilterTopicRequestMaxGasLimitChanged and is used to iterate over the raw logs and unpacked data for TopicRequestMaxGasLimitChanged events raised by the RequestManager contract.
type RequestManagerTopicRequestMaxGasLimitChangedIterator struct {
	Event *RequestManagerTopicRequestMaxGasLimitChanged // Event containing the contract specifics and raw log

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
func (it *RequestManagerTopicRequestMaxGasLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerTopicRequestMaxGasLimitChanged)
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
		it.Event = new(RequestManagerTopicRequestMaxGasLimitChanged)
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
func (it *RequestManagerTopicRequestMaxGasLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerTopicRequestMaxGasLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerTopicRequestMaxGasLimitChanged represents a TopicRequestMaxGasLimitChanged event raised by the RequestManager contract.
type RequestManagerTopicRequestMaxGasLimitChanged struct {
	OldMaxGasLimit common.Address
	NewMaxGasLimit common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTopicRequestMaxGasLimitChanged is a free log retrieval operation binding the contract event 0x76d445333b12308d4f2b8e4180f270fbca946f71fd591f7a511f0385f7e862d3.
//
// Solidity: event TopicRequestMaxGasLimitChanged(address indexed oldMaxGasLimit, address indexed newMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) FilterTopicRequestMaxGasLimitChanged(opts *bind.FilterOpts, oldMaxGasLimit []common.Address, newMaxGasLimit []common.Address) (*RequestManagerTopicRequestMaxGasLimitChangedIterator, error) {

	var oldMaxGasLimitRule []interface{}
	for _, oldMaxGasLimitItem := range oldMaxGasLimit {
		oldMaxGasLimitRule = append(oldMaxGasLimitRule, oldMaxGasLimitItem)
	}
	var newMaxGasLimitRule []interface{}
	for _, newMaxGasLimitItem := range newMaxGasLimit {
		newMaxGasLimitRule = append(newMaxGasLimitRule, newMaxGasLimitItem)
	}

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "TopicRequestMaxGasLimitChanged", oldMaxGasLimitRule, newMaxGasLimitRule)
	if err != nil {
		return nil, err
	}
	return &RequestManagerTopicRequestMaxGasLimitChangedIterator{contract: _RequestManager.contract, event: "TopicRequestMaxGasLimitChanged", logs: logs, sub: sub}, nil
}

// WatchTopicRequestMaxGasLimitChanged is a free log subscription operation binding the contract event 0x76d445333b12308d4f2b8e4180f270fbca946f71fd591f7a511f0385f7e862d3.
//
// Solidity: event TopicRequestMaxGasLimitChanged(address indexed oldMaxGasLimit, address indexed newMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) WatchTopicRequestMaxGasLimitChanged(opts *bind.WatchOpts, sink chan<- *RequestManagerTopicRequestMaxGasLimitChanged, oldMaxGasLimit []common.Address, newMaxGasLimit []common.Address) (event.Subscription, error) {

	var oldMaxGasLimitRule []interface{}
	for _, oldMaxGasLimitItem := range oldMaxGasLimit {
		oldMaxGasLimitRule = append(oldMaxGasLimitRule, oldMaxGasLimitItem)
	}
	var newMaxGasLimitRule []interface{}
	for _, newMaxGasLimitItem := range newMaxGasLimit {
		newMaxGasLimitRule = append(newMaxGasLimitRule, newMaxGasLimitItem)
	}

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "TopicRequestMaxGasLimitChanged", oldMaxGasLimitRule, newMaxGasLimitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerTopicRequestMaxGasLimitChanged)
				if err := _RequestManager.contract.UnpackLog(event, "TopicRequestMaxGasLimitChanged", log); err != nil {
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

// ParseTopicRequestMaxGasLimitChanged is a log parse operation binding the contract event 0x76d445333b12308d4f2b8e4180f270fbca946f71fd591f7a511f0385f7e862d3.
//
// Solidity: event TopicRequestMaxGasLimitChanged(address indexed oldMaxGasLimit, address indexed newMaxGasLimit)
func (_RequestManager *RequestManagerFilterer) ParseTopicRequestMaxGasLimitChanged(log types.Log) (*RequestManagerTopicRequestMaxGasLimitChanged, error) {
	event := new(RequestManagerTopicRequestMaxGasLimitChanged)
	if err := _RequestManager.contract.UnpackLog(event, "TopicRequestMaxGasLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestManagerWithdrawRequestTopicFundsIterator is returned from FilterWithdrawRequestTopicFunds and is used to iterate over the raw logs and unpacked data for WithdrawRequestTopicFunds events raised by the RequestManager contract.
type RequestManagerWithdrawRequestTopicFundsIterator struct {
	Event *RequestManagerWithdrawRequestTopicFunds // Event containing the contract specifics and raw log

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
func (it *RequestManagerWithdrawRequestTopicFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestManagerWithdrawRequestTopicFunds)
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
		it.Event = new(RequestManagerWithdrawRequestTopicFunds)
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
func (it *RequestManagerWithdrawRequestTopicFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestManagerWithdrawRequestTopicFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestManagerWithdrawRequestTopicFunds represents a WithdrawRequestTopicFunds event raised by the RequestManager contract.
type RequestManagerWithdrawRequestTopicFunds struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequestTopicFunds is a free log retrieval operation binding the contract event 0x67d7f21a3646788aa9d3c4cdf40f1663c29db753483b4d36e0621b23fdbe2954.
//
// Solidity: event WithdrawRequestTopicFunds(uint256 amount)
func (_RequestManager *RequestManagerFilterer) FilterWithdrawRequestTopicFunds(opts *bind.FilterOpts) (*RequestManagerWithdrawRequestTopicFundsIterator, error) {

	logs, sub, err := _RequestManager.contract.FilterLogs(opts, "WithdrawRequestTopicFunds")
	if err != nil {
		return nil, err
	}
	return &RequestManagerWithdrawRequestTopicFundsIterator{contract: _RequestManager.contract, event: "WithdrawRequestTopicFunds", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequestTopicFunds is a free log subscription operation binding the contract event 0x67d7f21a3646788aa9d3c4cdf40f1663c29db753483b4d36e0621b23fdbe2954.
//
// Solidity: event WithdrawRequestTopicFunds(uint256 amount)
func (_RequestManager *RequestManagerFilterer) WatchWithdrawRequestTopicFunds(opts *bind.WatchOpts, sink chan<- *RequestManagerWithdrawRequestTopicFunds) (event.Subscription, error) {

	logs, sub, err := _RequestManager.contract.WatchLogs(opts, "WithdrawRequestTopicFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestManagerWithdrawRequestTopicFunds)
				if err := _RequestManager.contract.UnpackLog(event, "WithdrawRequestTopicFunds", log); err != nil {
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

// ParseWithdrawRequestTopicFunds is a log parse operation binding the contract event 0x67d7f21a3646788aa9d3c4cdf40f1663c29db753483b4d36e0621b23fdbe2954.
//
// Solidity: event WithdrawRequestTopicFunds(uint256 amount)
func (_RequestManager *RequestManagerFilterer) ParseWithdrawRequestTopicFunds(log types.Log) (*RequestManagerWithdrawRequestTopicFunds, error) {
	event := new(RequestManagerWithdrawRequestTopicFunds)
	if err := _RequestManager.contract.UnpackLog(event, "WithdrawRequestTopicFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
