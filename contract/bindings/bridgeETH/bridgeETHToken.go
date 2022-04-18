// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeETH

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
)

// TokengoMetaData contains all meta data concerning the Tokengo contract.
var TokengoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mplgr_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cb_ddid\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_cb_rid\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DebugDepositMPLGR\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositMPLGR\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositMPLGRBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawMPLGR\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge_gas_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cb_ddid\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cb_rid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"handler_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mplgr_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mplgr_amounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cb_ddid\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_cb_rid\",\"type\":\"bytes32\"}],\"name\":\"admin_update_bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bridge_gas_fee\",\"type\":\"uint256\"}],\"name\":\"set_bridge_gas_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit_mplgr_bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"widthdraw_mplgr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit_mplgr\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TokengoABI is the input ABI used to generate the binding from.
// Deprecated: Use TokengoMetaData.ABI instead.
var TokengoABI = TokengoMetaData.ABI

// Tokengo is an auto generated Go binding around an Ethereum contract.
type Tokengo struct {
	TokengoCaller     // Read-only binding to the contract
	TokengoTransactor // Write-only binding to the contract
	TokengoFilterer   // Log filterer for contract events
}

// TokengoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokengoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokengoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokengoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokengoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokengoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokengoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokengoSession struct {
	Contract     *Tokengo          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokengoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokengoCallerSession struct {
	Contract *TokengoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TokengoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokengoTransactorSession struct {
	Contract     *TokengoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TokengoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokengoRaw struct {
	Contract *Tokengo // Generic contract binding to access the raw methods on
}

// TokengoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokengoCallerRaw struct {
	Contract *TokengoCaller // Generic read-only contract binding to access the raw methods on
}

// TokengoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokengoTransactorRaw struct {
	Contract *TokengoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokengo creates a new instance of Tokengo, bound to a specific deployed contract.
func NewTokengo(address common.Address, backend bind.ContractBackend) (*Tokengo, error) {
	contract, err := bindTokengo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokengo{TokengoCaller: TokengoCaller{contract: contract}, TokengoTransactor: TokengoTransactor{contract: contract}, TokengoFilterer: TokengoFilterer{contract: contract}}, nil
}

// NewTokengoCaller creates a new read-only instance of Tokengo, bound to a specific deployed contract.
func NewTokengoCaller(address common.Address, caller bind.ContractCaller) (*TokengoCaller, error) {
	contract, err := bindTokengo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokengoCaller{contract: contract}, nil
}

// NewTokengoTransactor creates a new write-only instance of Tokengo, bound to a specific deployed contract.
func NewTokengoTransactor(address common.Address, transactor bind.ContractTransactor) (*TokengoTransactor, error) {
	contract, err := bindTokengo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokengoTransactor{contract: contract}, nil
}

// NewTokengoFilterer creates a new log filterer instance of Tokengo, bound to a specific deployed contract.
func NewTokengoFilterer(address common.Address, filterer bind.ContractFilterer) (*TokengoFilterer, error) {
	contract, err := bindTokengo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokengoFilterer{contract: contract}, nil
}

// bindTokengo binds a generic wrapper to an already deployed contract.
func bindTokengo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokengoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokengo *TokengoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokengo.Contract.TokengoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokengo *TokengoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokengo.Contract.TokengoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokengo *TokengoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokengo.Contract.TokengoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokengo *TokengoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokengo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokengo *TokengoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokengo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokengo *TokengoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokengo.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokengo *TokengoCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokengo *TokengoSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.Balances(&_Tokengo.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokengo *TokengoCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.Balances(&_Tokengo.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x1255c189.
//
// Solidity: function bridge_address() view returns(address)
func (_Tokengo *TokengoCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "bridge_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x1255c189.
//
// Solidity: function bridge_address() view returns(address)
func (_Tokengo *TokengoSession) BridgeAddress() (common.Address, error) {
	return _Tokengo.Contract.BridgeAddress(&_Tokengo.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x1255c189.
//
// Solidity: function bridge_address() view returns(address)
func (_Tokengo *TokengoCallerSession) BridgeAddress() (common.Address, error) {
	return _Tokengo.Contract.BridgeAddress(&_Tokengo.CallOpts)
}

// BridgeGasFee is a free data retrieval call binding the contract method 0x176a78bb.
//
// Solidity: function bridge_gas_fee() view returns(uint256)
func (_Tokengo *TokengoCaller) BridgeGasFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "bridge_gas_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeGasFee is a free data retrieval call binding the contract method 0x176a78bb.
//
// Solidity: function bridge_gas_fee() view returns(uint256)
func (_Tokengo *TokengoSession) BridgeGasFee() (*big.Int, error) {
	return _Tokengo.Contract.BridgeGasFee(&_Tokengo.CallOpts)
}

// BridgeGasFee is a free data retrieval call binding the contract method 0x176a78bb.
//
// Solidity: function bridge_gas_fee() view returns(uint256)
func (_Tokengo *TokengoCallerSession) BridgeGasFee() (*big.Int, error) {
	return _Tokengo.Contract.BridgeGasFee(&_Tokengo.CallOpts)
}

// CbDdid is a free data retrieval call binding the contract method 0x04ecb844.
//
// Solidity: function cb_ddid() view returns(uint8)
func (_Tokengo *TokengoCaller) CbDdid(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "cb_ddid")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CbDdid is a free data retrieval call binding the contract method 0x04ecb844.
//
// Solidity: function cb_ddid() view returns(uint8)
func (_Tokengo *TokengoSession) CbDdid() (uint8, error) {
	return _Tokengo.Contract.CbDdid(&_Tokengo.CallOpts)
}

// CbDdid is a free data retrieval call binding the contract method 0x04ecb844.
//
// Solidity: function cb_ddid() view returns(uint8)
func (_Tokengo *TokengoCallerSession) CbDdid() (uint8, error) {
	return _Tokengo.Contract.CbDdid(&_Tokengo.CallOpts)
}

// CbRid is a free data retrieval call binding the contract method 0x2068a843.
//
// Solidity: function cb_rid() view returns(bytes32)
func (_Tokengo *TokengoCaller) CbRid(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "cb_rid")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CbRid is a free data retrieval call binding the contract method 0x2068a843.
//
// Solidity: function cb_rid() view returns(bytes32)
func (_Tokengo *TokengoSession) CbRid() ([32]byte, error) {
	return _Tokengo.Contract.CbRid(&_Tokengo.CallOpts)
}

// CbRid is a free data retrieval call binding the contract method 0x2068a843.
//
// Solidity: function cb_rid() view returns(bytes32)
func (_Tokengo *TokengoCallerSession) CbRid() ([32]byte, error) {
	return _Tokengo.Contract.CbRid(&_Tokengo.CallOpts)
}

// HandlerAddress is a free data retrieval call binding the contract method 0xab7272ea.
//
// Solidity: function handler_address() view returns(address)
func (_Tokengo *TokengoCaller) HandlerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "handler_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HandlerAddress is a free data retrieval call binding the contract method 0xab7272ea.
//
// Solidity: function handler_address() view returns(address)
func (_Tokengo *TokengoSession) HandlerAddress() (common.Address, error) {
	return _Tokengo.Contract.HandlerAddress(&_Tokengo.CallOpts)
}

// HandlerAddress is a free data retrieval call binding the contract method 0xab7272ea.
//
// Solidity: function handler_address() view returns(address)
func (_Tokengo *TokengoCallerSession) HandlerAddress() (common.Address, error) {
	return _Tokengo.Contract.HandlerAddress(&_Tokengo.CallOpts)
}

// MplgrAddress is a free data retrieval call binding the contract method 0xf8f1603e.
//
// Solidity: function mplgr_address() view returns(address)
func (_Tokengo *TokengoCaller) MplgrAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "mplgr_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MplgrAddress is a free data retrieval call binding the contract method 0xf8f1603e.
//
// Solidity: function mplgr_address() view returns(address)
func (_Tokengo *TokengoSession) MplgrAddress() (common.Address, error) {
	return _Tokengo.Contract.MplgrAddress(&_Tokengo.CallOpts)
}

// MplgrAddress is a free data retrieval call binding the contract method 0xf8f1603e.
//
// Solidity: function mplgr_address() view returns(address)
func (_Tokengo *TokengoCallerSession) MplgrAddress() (common.Address, error) {
	return _Tokengo.Contract.MplgrAddress(&_Tokengo.CallOpts)
}

// MplgrAmounts is a free data retrieval call binding the contract method 0x0016a071.
//
// Solidity: function mplgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoCaller) MplgrAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "mplgr_amounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MplgrAmounts is a free data retrieval call binding the contract method 0x0016a071.
//
// Solidity: function mplgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoSession) MplgrAmounts(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.MplgrAmounts(&_Tokengo.CallOpts, arg0)
}

// MplgrAmounts is a free data retrieval call binding the contract method 0x0016a071.
//
// Solidity: function mplgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoCallerSession) MplgrAmounts(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.MplgrAmounts(&_Tokengo.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokengo *TokengoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokengo *TokengoSession) Owner() (common.Address, error) {
	return _Tokengo.Contract.Owner(&_Tokengo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokengo *TokengoCallerSession) Owner() (common.Address, error) {
	return _Tokengo.Contract.Owner(&_Tokengo.CallOpts)
}

// AdminUpdateBridge is a paid mutator transaction binding the contract method 0xc7b4c5bf.
//
// Solidity: function admin_update_bridge(address _bridge_address, address _handler_address, uint8 _cb_ddid, bytes32 _cb_rid) returns()
func (_Tokengo *TokengoTransactor) AdminUpdateBridge(opts *bind.TransactOpts, _bridge_address common.Address, _handler_address common.Address, _cb_ddid uint8, _cb_rid [32]byte) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "admin_update_bridge", _bridge_address, _handler_address, _cb_ddid, _cb_rid)
}

// AdminUpdateBridge is a paid mutator transaction binding the contract method 0xc7b4c5bf.
//
// Solidity: function admin_update_bridge(address _bridge_address, address _handler_address, uint8 _cb_ddid, bytes32 _cb_rid) returns()
func (_Tokengo *TokengoSession) AdminUpdateBridge(_bridge_address common.Address, _handler_address common.Address, _cb_ddid uint8, _cb_rid [32]byte) (*types.Transaction, error) {
	return _Tokengo.Contract.AdminUpdateBridge(&_Tokengo.TransactOpts, _bridge_address, _handler_address, _cb_ddid, _cb_rid)
}

// AdminUpdateBridge is a paid mutator transaction binding the contract method 0xc7b4c5bf.
//
// Solidity: function admin_update_bridge(address _bridge_address, address _handler_address, uint8 _cb_ddid, bytes32 _cb_rid) returns()
func (_Tokengo *TokengoTransactorSession) AdminUpdateBridge(_bridge_address common.Address, _handler_address common.Address, _cb_ddid uint8, _cb_rid [32]byte) (*types.Transaction, error) {
	return _Tokengo.Contract.AdminUpdateBridge(&_Tokengo.TransactOpts, _bridge_address, _handler_address, _cb_ddid, _cb_rid)
}

// DepositMplgr is a paid mutator transaction binding the contract method 0xb285a9c1.
//
// Solidity: function deposit_mplgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoTransactor) DepositMplgr(opts *bind.TransactOpts, _owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "deposit_mplgr", _owner, amount)
}

// DepositMplgr is a paid mutator transaction binding the contract method 0xb285a9c1.
//
// Solidity: function deposit_mplgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoSession) DepositMplgr(_owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositMplgr(&_Tokengo.TransactOpts, _owner, amount)
}

// DepositMplgr is a paid mutator transaction binding the contract method 0xb285a9c1.
//
// Solidity: function deposit_mplgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoTransactorSession) DepositMplgr(_owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositMplgr(&_Tokengo.TransactOpts, _owner, amount)
}

// DepositMplgrBridge is a paid mutator transaction binding the contract method 0x02c0ab3b.
//
// Solidity: function deposit_mplgr_bridge(bytes data) returns()
func (_Tokengo *TokengoTransactor) DepositMplgrBridge(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "deposit_mplgr_bridge", data)
}

// DepositMplgrBridge is a paid mutator transaction binding the contract method 0x02c0ab3b.
//
// Solidity: function deposit_mplgr_bridge(bytes data) returns()
func (_Tokengo *TokengoSession) DepositMplgrBridge(data []byte) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositMplgrBridge(&_Tokengo.TransactOpts, data)
}

// DepositMplgrBridge is a paid mutator transaction binding the contract method 0x02c0ab3b.
//
// Solidity: function deposit_mplgr_bridge(bytes data) returns()
func (_Tokengo *TokengoTransactorSession) DepositMplgrBridge(data []byte) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositMplgrBridge(&_Tokengo.TransactOpts, data)
}

// SetBridgeGasFee is a paid mutator transaction binding the contract method 0x44434d94.
//
// Solidity: function set_bridge_gas_fee(uint256 _bridge_gas_fee) returns()
func (_Tokengo *TokengoTransactor) SetBridgeGasFee(opts *bind.TransactOpts, _bridge_gas_fee *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "set_bridge_gas_fee", _bridge_gas_fee)
}

// SetBridgeGasFee is a paid mutator transaction binding the contract method 0x44434d94.
//
// Solidity: function set_bridge_gas_fee(uint256 _bridge_gas_fee) returns()
func (_Tokengo *TokengoSession) SetBridgeGasFee(_bridge_gas_fee *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.SetBridgeGasFee(&_Tokengo.TransactOpts, _bridge_gas_fee)
}

// SetBridgeGasFee is a paid mutator transaction binding the contract method 0x44434d94.
//
// Solidity: function set_bridge_gas_fee(uint256 _bridge_gas_fee) returns()
func (_Tokengo *TokengoTransactorSession) SetBridgeGasFee(_bridge_gas_fee *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.SetBridgeGasFee(&_Tokengo.TransactOpts, _bridge_gas_fee)
}

// WidthdrawMplgr is a paid mutator transaction binding the contract method 0x5b1cf3d2.
//
// Solidity: function widthdraw_mplgr(uint256 amount) returns()
func (_Tokengo *TokengoTransactor) WidthdrawMplgr(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "widthdraw_mplgr", amount)
}

// WidthdrawMplgr is a paid mutator transaction binding the contract method 0x5b1cf3d2.
//
// Solidity: function widthdraw_mplgr(uint256 amount) returns()
func (_Tokengo *TokengoSession) WidthdrawMplgr(amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.WidthdrawMplgr(&_Tokengo.TransactOpts, amount)
}

// WidthdrawMplgr is a paid mutator transaction binding the contract method 0x5b1cf3d2.
//
// Solidity: function widthdraw_mplgr(uint256 amount) returns()
func (_Tokengo *TokengoTransactorSession) WidthdrawMplgr(amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.WidthdrawMplgr(&_Tokengo.TransactOpts, amount)
}

// TokengoDebugDepositMPLGRIterator is returned from FilterDebugDepositMPLGR and is used to iterate over the raw logs and unpacked data for DebugDepositMPLGR events raised by the Tokengo contract.
type TokengoDebugDepositMPLGRIterator struct {
	Event *TokengoDebugDepositMPLGR // Event containing the contract specifics and raw log

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
func (it *TokengoDebugDepositMPLGRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoDebugDepositMPLGR)
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
		it.Event = new(TokengoDebugDepositMPLGR)
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
func (it *TokengoDebugDepositMPLGRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoDebugDepositMPLGRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoDebugDepositMPLGR represents a DebugDepositMPLGR event raised by the Tokengo contract.
type TokengoDebugDepositMPLGR struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDebugDepositMPLGR is a free log retrieval operation binding the contract event 0x28df32dd70ae590e6391db741d3f5890440105eb3645092993a488ee1ba573de.
//
// Solidity: event DebugDepositMPLGR(bytes data)
func (_Tokengo *TokengoFilterer) FilterDebugDepositMPLGR(opts *bind.FilterOpts) (*TokengoDebugDepositMPLGRIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "DebugDepositMPLGR")
	if err != nil {
		return nil, err
	}
	return &TokengoDebugDepositMPLGRIterator{contract: _Tokengo.contract, event: "DebugDepositMPLGR", logs: logs, sub: sub}, nil
}

// WatchDebugDepositMPLGR is a free log subscription operation binding the contract event 0x28df32dd70ae590e6391db741d3f5890440105eb3645092993a488ee1ba573de.
//
// Solidity: event DebugDepositMPLGR(bytes data)
func (_Tokengo *TokengoFilterer) WatchDebugDepositMPLGR(opts *bind.WatchOpts, sink chan<- *TokengoDebugDepositMPLGR) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "DebugDepositMPLGR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoDebugDepositMPLGR)
				if err := _Tokengo.contract.UnpackLog(event, "DebugDepositMPLGR", log); err != nil {
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

// ParseDebugDepositMPLGR is a log parse operation binding the contract event 0x28df32dd70ae590e6391db741d3f5890440105eb3645092993a488ee1ba573de.
//
// Solidity: event DebugDepositMPLGR(bytes data)
func (_Tokengo *TokengoFilterer) ParseDebugDepositMPLGR(log types.Log) (*TokengoDebugDepositMPLGR, error) {
	event := new(TokengoDebugDepositMPLGR)
	if err := _Tokengo.contract.UnpackLog(event, "DebugDepositMPLGR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokengoDepositMPLGRIterator is returned from FilterDepositMPLGR and is used to iterate over the raw logs and unpacked data for DepositMPLGR events raised by the Tokengo contract.
type TokengoDepositMPLGRIterator struct {
	Event *TokengoDepositMPLGR // Event containing the contract specifics and raw log

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
func (it *TokengoDepositMPLGRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoDepositMPLGR)
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
		it.Event = new(TokengoDepositMPLGR)
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
func (it *TokengoDepositMPLGRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoDepositMPLGRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoDepositMPLGR represents a DepositMPLGR event raised by the Tokengo contract.
type TokengoDepositMPLGR struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositMPLGR is a free log retrieval operation binding the contract event 0xd6f88f0ecba27de30dcf24e4923aeca8b1943c9cdda43f961beaa91c4ce465ac.
//
// Solidity: event DepositMPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) FilterDepositMPLGR(opts *bind.FilterOpts) (*TokengoDepositMPLGRIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "DepositMPLGR")
	if err != nil {
		return nil, err
	}
	return &TokengoDepositMPLGRIterator{contract: _Tokengo.contract, event: "DepositMPLGR", logs: logs, sub: sub}, nil
}

// WatchDepositMPLGR is a free log subscription operation binding the contract event 0xd6f88f0ecba27de30dcf24e4923aeca8b1943c9cdda43f961beaa91c4ce465ac.
//
// Solidity: event DepositMPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) WatchDepositMPLGR(opts *bind.WatchOpts, sink chan<- *TokengoDepositMPLGR) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "DepositMPLGR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoDepositMPLGR)
				if err := _Tokengo.contract.UnpackLog(event, "DepositMPLGR", log); err != nil {
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

// ParseDepositMPLGR is a log parse operation binding the contract event 0xd6f88f0ecba27de30dcf24e4923aeca8b1943c9cdda43f961beaa91c4ce465ac.
//
// Solidity: event DepositMPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) ParseDepositMPLGR(log types.Log) (*TokengoDepositMPLGR, error) {
	event := new(TokengoDepositMPLGR)
	if err := _Tokengo.contract.UnpackLog(event, "DepositMPLGR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokengoDepositMPLGRBridgeIterator is returned from FilterDepositMPLGRBridge and is used to iterate over the raw logs and unpacked data for DepositMPLGRBridge events raised by the Tokengo contract.
type TokengoDepositMPLGRBridgeIterator struct {
	Event *TokengoDepositMPLGRBridge // Event containing the contract specifics and raw log

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
func (it *TokengoDepositMPLGRBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoDepositMPLGRBridge)
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
		it.Event = new(TokengoDepositMPLGRBridge)
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
func (it *TokengoDepositMPLGRBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoDepositMPLGRBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoDepositMPLGRBridge represents a DepositMPLGRBridge event raised by the Tokengo contract.
type TokengoDepositMPLGRBridge struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositMPLGRBridge is a free log retrieval operation binding the contract event 0x7913b5c1ff9a95dd5f46c9ffdaec78cd8b8a88cad0a011de4f70383ac4155ec7.
//
// Solidity: event DepositMPLGRBridge(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) FilterDepositMPLGRBridge(opts *bind.FilterOpts) (*TokengoDepositMPLGRBridgeIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "DepositMPLGRBridge")
	if err != nil {
		return nil, err
	}
	return &TokengoDepositMPLGRBridgeIterator{contract: _Tokengo.contract, event: "DepositMPLGRBridge", logs: logs, sub: sub}, nil
}

// WatchDepositMPLGRBridge is a free log subscription operation binding the contract event 0x7913b5c1ff9a95dd5f46c9ffdaec78cd8b8a88cad0a011de4f70383ac4155ec7.
//
// Solidity: event DepositMPLGRBridge(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) WatchDepositMPLGRBridge(opts *bind.WatchOpts, sink chan<- *TokengoDepositMPLGRBridge) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "DepositMPLGRBridge")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoDepositMPLGRBridge)
				if err := _Tokengo.contract.UnpackLog(event, "DepositMPLGRBridge", log); err != nil {
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

// ParseDepositMPLGRBridge is a log parse operation binding the contract event 0x7913b5c1ff9a95dd5f46c9ffdaec78cd8b8a88cad0a011de4f70383ac4155ec7.
//
// Solidity: event DepositMPLGRBridge(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) ParseDepositMPLGRBridge(log types.Log) (*TokengoDepositMPLGRBridge, error) {
	event := new(TokengoDepositMPLGRBridge)
	if err := _Tokengo.contract.UnpackLog(event, "DepositMPLGRBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokengoWithdrawMPLGRIterator is returned from FilterWithdrawMPLGR and is used to iterate over the raw logs and unpacked data for WithdrawMPLGR events raised by the Tokengo contract.
type TokengoWithdrawMPLGRIterator struct {
	Event *TokengoWithdrawMPLGR // Event containing the contract specifics and raw log

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
func (it *TokengoWithdrawMPLGRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoWithdrawMPLGR)
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
		it.Event = new(TokengoWithdrawMPLGR)
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
func (it *TokengoWithdrawMPLGRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoWithdrawMPLGRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoWithdrawMPLGR represents a WithdrawMPLGR event raised by the Tokengo contract.
type TokengoWithdrawMPLGR struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMPLGR is a free log retrieval operation binding the contract event 0x2a6c9f7fc2e6de3721bffcd5fa4eb025b223d57905c95dbdfcadab0eb6310335.
//
// Solidity: event WithdrawMPLGR(address recipient, uint256 amount)
func (_Tokengo *TokengoFilterer) FilterWithdrawMPLGR(opts *bind.FilterOpts) (*TokengoWithdrawMPLGRIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "WithdrawMPLGR")
	if err != nil {
		return nil, err
	}
	return &TokengoWithdrawMPLGRIterator{contract: _Tokengo.contract, event: "WithdrawMPLGR", logs: logs, sub: sub}, nil
}

// WatchWithdrawMPLGR is a free log subscription operation binding the contract event 0x2a6c9f7fc2e6de3721bffcd5fa4eb025b223d57905c95dbdfcadab0eb6310335.
//
// Solidity: event WithdrawMPLGR(address recipient, uint256 amount)
func (_Tokengo *TokengoFilterer) WatchWithdrawMPLGR(opts *bind.WatchOpts, sink chan<- *TokengoWithdrawMPLGR) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "WithdrawMPLGR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoWithdrawMPLGR)
				if err := _Tokengo.contract.UnpackLog(event, "WithdrawMPLGR", log); err != nil {
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

// ParseWithdrawMPLGR is a log parse operation binding the contract event 0x2a6c9f7fc2e6de3721bffcd5fa4eb025b223d57905c95dbdfcadab0eb6310335.
//
// Solidity: event WithdrawMPLGR(address recipient, uint256 amount)
func (_Tokengo *TokengoFilterer) ParseWithdrawMPLGR(log types.Log) (*TokengoWithdrawMPLGR, error) {
	event := new(TokengoWithdrawMPLGR)
	if err := _Tokengo.contract.UnpackLog(event, "WithdrawMPLGR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
