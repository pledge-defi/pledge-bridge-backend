// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeBSC

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_plgr_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridge_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cb_ddid\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_cb_rid\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"DepositPLGR\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawPLGR\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge_gas_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can_release\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cb_ddid\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cb_rid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"handler_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"locked_infos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked_plgr_tx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"plgr_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"plgr_amounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_mplgr_release\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"total_plgr_locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_cb_ddid\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_cb_rid\",\"type\":\"bytes32\"}],\"name\":\"admin_update_bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bridge_gas_fee\",\"type\":\"uint256\"}],\"name\":\"set_bridge_gas_fee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_total_mplgr_release\",\"type\":\"uint256\"}],\"name\":\"set_total_release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit_plgr\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"widthdraw_plgr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit_mplgr_bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"is_release_all_locked_plgr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"execute_upkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// CanRelease is a free data retrieval call binding the contract method 0xb9e77873.
//
// Solidity: function can_release(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCaller) CanRelease(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "can_release", arg0)

	outstruct := new(struct {
		Owner  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CanRelease is a free data retrieval call binding the contract method 0xb9e77873.
//
// Solidity: function can_release(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoSession) CanRelease(arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.CanRelease(&_Tokengo.CallOpts, arg0)
}

// CanRelease is a free data retrieval call binding the contract method 0xb9e77873.
//
// Solidity: function can_release(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCallerSession) CanRelease(arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.CanRelease(&_Tokengo.CallOpts, arg0)
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

// Factor is a free data retrieval call binding the contract method 0x54f703f8.
//
// Solidity: function factor() view returns(uint256)
func (_Tokengo *TokengoCaller) Factor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "factor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Factor is a free data retrieval call binding the contract method 0x54f703f8.
//
// Solidity: function factor() view returns(uint256)
func (_Tokengo *TokengoSession) Factor() (*big.Int, error) {
	return _Tokengo.Contract.Factor(&_Tokengo.CallOpts)
}

// Factor is a free data retrieval call binding the contract method 0x54f703f8.
//
// Solidity: function factor() view returns(uint256)
func (_Tokengo *TokengoCallerSession) Factor() (*big.Int, error) {
	return _Tokengo.Contract.Factor(&_Tokengo.CallOpts)
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

// IsReleaseAllLockedPlgr is a free data retrieval call binding the contract method 0x7742cffa.
//
// Solidity: function is_release_all_locked_plgr() view returns(bool)
func (_Tokengo *TokengoCaller) IsReleaseAllLockedPlgr(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "is_release_all_locked_plgr")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleaseAllLockedPlgr is a free data retrieval call binding the contract method 0x7742cffa.
//
// Solidity: function is_release_all_locked_plgr() view returns(bool)
func (_Tokengo *TokengoSession) IsReleaseAllLockedPlgr() (bool, error) {
	return _Tokengo.Contract.IsReleaseAllLockedPlgr(&_Tokengo.CallOpts)
}

// IsReleaseAllLockedPlgr is a free data retrieval call binding the contract method 0x7742cffa.
//
// Solidity: function is_release_all_locked_plgr() view returns(bool)
func (_Tokengo *TokengoCallerSession) IsReleaseAllLockedPlgr() (bool, error) {
	return _Tokengo.Contract.IsReleaseAllLockedPlgr(&_Tokengo.CallOpts)
}

// LockedInfos is a free data retrieval call binding the contract method 0xc970bdf4.
//
// Solidity: function locked_infos(uint256 ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCaller) LockedInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "locked_infos", arg0)

	outstruct := new(struct {
		Owner  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedInfos is a free data retrieval call binding the contract method 0xc970bdf4.
//
// Solidity: function locked_infos(uint256 ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoSession) LockedInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.LockedInfos(&_Tokengo.CallOpts, arg0)
}

// LockedInfos is a free data retrieval call binding the contract method 0xc970bdf4.
//
// Solidity: function locked_infos(uint256 ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCallerSession) LockedInfos(arg0 *big.Int) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.LockedInfos(&_Tokengo.CallOpts, arg0)
}

// LockedPlgrTx is a free data retrieval call binding the contract method 0x6da538ec.
//
// Solidity: function locked_plgr_tx(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCaller) LockedPlgrTx(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "locked_plgr_tx", arg0)

	outstruct := new(struct {
		Owner  common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedPlgrTx is a free data retrieval call binding the contract method 0x6da538ec.
//
// Solidity: function locked_plgr_tx(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoSession) LockedPlgrTx(arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.LockedPlgrTx(&_Tokengo.CallOpts, arg0)
}

// LockedPlgrTx is a free data retrieval call binding the contract method 0x6da538ec.
//
// Solidity: function locked_plgr_tx(address ) view returns(address owner, uint256 amount)
func (_Tokengo *TokengoCallerSession) LockedPlgrTx(arg0 common.Address) (struct {
	Owner  common.Address
	Amount *big.Int
}, error) {
	return _Tokengo.Contract.LockedPlgrTx(&_Tokengo.CallOpts, arg0)
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

// PlgrAddress is a free data retrieval call binding the contract method 0x3219c4df.
//
// Solidity: function plgr_address() view returns(address)
func (_Tokengo *TokengoCaller) PlgrAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "plgr_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlgrAddress is a free data retrieval call binding the contract method 0x3219c4df.
//
// Solidity: function plgr_address() view returns(address)
func (_Tokengo *TokengoSession) PlgrAddress() (common.Address, error) {
	return _Tokengo.Contract.PlgrAddress(&_Tokengo.CallOpts)
}

// PlgrAddress is a free data retrieval call binding the contract method 0x3219c4df.
//
// Solidity: function plgr_address() view returns(address)
func (_Tokengo *TokengoCallerSession) PlgrAddress() (common.Address, error) {
	return _Tokengo.Contract.PlgrAddress(&_Tokengo.CallOpts)
}

// PlgrAmounts is a free data retrieval call binding the contract method 0xdfad049e.
//
// Solidity: function plgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoCaller) PlgrAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "plgr_amounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlgrAmounts is a free data retrieval call binding the contract method 0xdfad049e.
//
// Solidity: function plgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoSession) PlgrAmounts(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.PlgrAmounts(&_Tokengo.CallOpts, arg0)
}

// PlgrAmounts is a free data retrieval call binding the contract method 0xdfad049e.
//
// Solidity: function plgr_amounts(address ) view returns(uint256)
func (_Tokengo *TokengoCallerSession) PlgrAmounts(arg0 common.Address) (*big.Int, error) {
	return _Tokengo.Contract.PlgrAmounts(&_Tokengo.CallOpts, arg0)
}

// TotalMplgrRelease is a free data retrieval call binding the contract method 0x585fb29d.
//
// Solidity: function total_mplgr_release() view returns(uint256)
func (_Tokengo *TokengoCaller) TotalMplgrRelease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "total_mplgr_release")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMplgrRelease is a free data retrieval call binding the contract method 0x585fb29d.
//
// Solidity: function total_mplgr_release() view returns(uint256)
func (_Tokengo *TokengoSession) TotalMplgrRelease() (*big.Int, error) {
	return _Tokengo.Contract.TotalMplgrRelease(&_Tokengo.CallOpts)
}

// TotalMplgrRelease is a free data retrieval call binding the contract method 0x585fb29d.
//
// Solidity: function total_mplgr_release() view returns(uint256)
func (_Tokengo *TokengoCallerSession) TotalMplgrRelease() (*big.Int, error) {
	return _Tokengo.Contract.TotalMplgrRelease(&_Tokengo.CallOpts)
}

// TotalPlgrLocked is a free data retrieval call binding the contract method 0xbb837979.
//
// Solidity: function total_plgr_locked() view returns(uint256)
func (_Tokengo *TokengoCaller) TotalPlgrLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokengo.contract.Call(opts, &out, "total_plgr_locked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPlgrLocked is a free data retrieval call binding the contract method 0xbb837979.
//
// Solidity: function total_plgr_locked() view returns(uint256)
func (_Tokengo *TokengoSession) TotalPlgrLocked() (*big.Int, error) {
	return _Tokengo.Contract.TotalPlgrLocked(&_Tokengo.CallOpts)
}

// TotalPlgrLocked is a free data retrieval call binding the contract method 0xbb837979.
//
// Solidity: function total_plgr_locked() view returns(uint256)
func (_Tokengo *TokengoCallerSession) TotalPlgrLocked() (*big.Int, error) {
	return _Tokengo.Contract.TotalPlgrLocked(&_Tokengo.CallOpts)
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

// DepositPlgr is a paid mutator transaction binding the contract method 0x2e7e0c29.
//
// Solidity: function deposit_plgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoTransactor) DepositPlgr(opts *bind.TransactOpts, _owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "deposit_plgr", _owner, amount)
}

// DepositPlgr is a paid mutator transaction binding the contract method 0x2e7e0c29.
//
// Solidity: function deposit_plgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoSession) DepositPlgr(_owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositPlgr(&_Tokengo.TransactOpts, _owner, amount)
}

// DepositPlgr is a paid mutator transaction binding the contract method 0x2e7e0c29.
//
// Solidity: function deposit_plgr(address _owner, uint256 amount) payable returns()
func (_Tokengo *TokengoTransactorSession) DepositPlgr(_owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.DepositPlgr(&_Tokengo.TransactOpts, _owner, amount)
}

// ExecuteUpkeep is a paid mutator transaction binding the contract method 0xadd8bb96.
//
// Solidity: function execute_upkeep() returns()
func (_Tokengo *TokengoTransactor) ExecuteUpkeep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "execute_upkeep")
}

// ExecuteUpkeep is a paid mutator transaction binding the contract method 0xadd8bb96.
//
// Solidity: function execute_upkeep() returns()
func (_Tokengo *TokengoSession) ExecuteUpkeep() (*types.Transaction, error) {
	return _Tokengo.Contract.ExecuteUpkeep(&_Tokengo.TransactOpts)
}

// ExecuteUpkeep is a paid mutator transaction binding the contract method 0xadd8bb96.
//
// Solidity: function execute_upkeep() returns()
func (_Tokengo *TokengoTransactorSession) ExecuteUpkeep() (*types.Transaction, error) {
	return _Tokengo.Contract.ExecuteUpkeep(&_Tokengo.TransactOpts)
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

// SetTotalRelease is a paid mutator transaction binding the contract method 0x5ca65060.
//
// Solidity: function set_total_release(uint256 _total_mplgr_release) returns()
func (_Tokengo *TokengoTransactor) SetTotalRelease(opts *bind.TransactOpts, _total_mplgr_release *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "set_total_release", _total_mplgr_release)
}

// SetTotalRelease is a paid mutator transaction binding the contract method 0x5ca65060.
//
// Solidity: function set_total_release(uint256 _total_mplgr_release) returns()
func (_Tokengo *TokengoSession) SetTotalRelease(_total_mplgr_release *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.SetTotalRelease(&_Tokengo.TransactOpts, _total_mplgr_release)
}

// SetTotalRelease is a paid mutator transaction binding the contract method 0x5ca65060.
//
// Solidity: function set_total_release(uint256 _total_mplgr_release) returns()
func (_Tokengo *TokengoTransactorSession) SetTotalRelease(_total_mplgr_release *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.SetTotalRelease(&_Tokengo.TransactOpts, _total_mplgr_release)
}

// WidthdrawPlgr is a paid mutator transaction binding the contract method 0x83080704.
//
// Solidity: function widthdraw_plgr(uint256 amount) returns()
func (_Tokengo *TokengoTransactor) WidthdrawPlgr(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.contract.Transact(opts, "widthdraw_plgr", amount)
}

// WidthdrawPlgr is a paid mutator transaction binding the contract method 0x83080704.
//
// Solidity: function widthdraw_plgr(uint256 amount) returns()
func (_Tokengo *TokengoSession) WidthdrawPlgr(amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.WidthdrawPlgr(&_Tokengo.TransactOpts, amount)
}

// WidthdrawPlgr is a paid mutator transaction binding the contract method 0x83080704.
//
// Solidity: function widthdraw_plgr(uint256 amount) returns()
func (_Tokengo *TokengoTransactorSession) WidthdrawPlgr(amount *big.Int) (*types.Transaction, error) {
	return _Tokengo.Contract.WidthdrawPlgr(&_Tokengo.TransactOpts, amount)
}

// TokengoDepositPLGRIterator is returned from FilterDepositPLGR and is used to iterate over the raw logs and unpacked data for DepositPLGR events raised by the Tokengo contract.
type TokengoDepositPLGRIterator struct {
	Event *TokengoDepositPLGR // Event containing the contract specifics and raw log

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
func (it *TokengoDepositPLGRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoDepositPLGR)
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
		it.Event = new(TokengoDepositPLGR)
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
func (it *TokengoDepositPLGRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoDepositPLGRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoDepositPLGR represents a DepositPLGR event raised by the Tokengo contract.
type TokengoDepositPLGR struct {
	Txid   [32]byte
	Owner  common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositPLGR is a free log retrieval operation binding the contract event 0x1008429730dc0e33a6e63ee42972c220ab8becaf21d29983a99a80f45efa060b.
//
// Solidity: event DepositPLGR(bytes32 txid, address owner, uint256 amount, uint256 time)
func (_Tokengo *TokengoFilterer) FilterDepositPLGR(opts *bind.FilterOpts) (*TokengoDepositPLGRIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "DepositPLGR")
	if err != nil {
		return nil, err
	}
	return &TokengoDepositPLGRIterator{contract: _Tokengo.contract, event: "DepositPLGR", logs: logs, sub: sub}, nil
}

// WatchDepositPLGR is a free log subscription operation binding the contract event 0x1008429730dc0e33a6e63ee42972c220ab8becaf21d29983a99a80f45efa060b.
//
// Solidity: event DepositPLGR(bytes32 txid, address owner, uint256 amount, uint256 time)
func (_Tokengo *TokengoFilterer) WatchDepositPLGR(opts *bind.WatchOpts, sink chan<- *TokengoDepositPLGR) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "DepositPLGR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoDepositPLGR)
				if err := _Tokengo.contract.UnpackLog(event, "DepositPLGR", log); err != nil {
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

// ParseDepositPLGR is a log parse operation binding the contract event 0x1008429730dc0e33a6e63ee42972c220ab8becaf21d29983a99a80f45efa060b.
//
// Solidity: event DepositPLGR(bytes32 txid, address owner, uint256 amount, uint256 time)
func (_Tokengo *TokengoFilterer) ParseDepositPLGR(log types.Log) (*TokengoDepositPLGR, error) {
	event := new(TokengoDepositPLGR)
	if err := _Tokengo.contract.UnpackLog(event, "DepositPLGR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokengoWithdrawPLGRIterator is returned from FilterWithdrawPLGR and is used to iterate over the raw logs and unpacked data for WithdrawPLGR events raised by the Tokengo contract.
type TokengoWithdrawPLGRIterator struct {
	Event *TokengoWithdrawPLGR // Event containing the contract specifics and raw log

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
func (it *TokengoWithdrawPLGRIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokengoWithdrawPLGR)
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
		it.Event = new(TokengoWithdrawPLGR)
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
func (it *TokengoWithdrawPLGRIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokengoWithdrawPLGRIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokengoWithdrawPLGR represents a WithdrawPLGR event raised by the Tokengo contract.
type TokengoWithdrawPLGR struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawPLGR is a free log retrieval operation binding the contract event 0x67081655db492a9224c5bc8d7d3ccadb8bb37d603b8697c1cb2081a682513854.
//
// Solidity: event WithdrawPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) FilterWithdrawPLGR(opts *bind.FilterOpts) (*TokengoWithdrawPLGRIterator, error) {

	logs, sub, err := _Tokengo.contract.FilterLogs(opts, "WithdrawPLGR")
	if err != nil {
		return nil, err
	}
	return &TokengoWithdrawPLGRIterator{contract: _Tokengo.contract, event: "WithdrawPLGR", logs: logs, sub: sub}, nil
}

// WatchWithdrawPLGR is a free log subscription operation binding the contract event 0x67081655db492a9224c5bc8d7d3ccadb8bb37d603b8697c1cb2081a682513854.
//
// Solidity: event WithdrawPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) WatchWithdrawPLGR(opts *bind.WatchOpts, sink chan<- *TokengoWithdrawPLGR) (event.Subscription, error) {

	logs, sub, err := _Tokengo.contract.WatchLogs(opts, "WithdrawPLGR")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokengoWithdrawPLGR)
				if err := _Tokengo.contract.UnpackLog(event, "WithdrawPLGR", log); err != nil {
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

// ParseWithdrawPLGR is a log parse operation binding the contract event 0x67081655db492a9224c5bc8d7d3ccadb8bb37d603b8697c1cb2081a682513854.
//
// Solidity: event WithdrawPLGR(address owner, uint256 amount)
func (_Tokengo *TokengoFilterer) ParseWithdrawPLGR(log types.Log) (*TokengoWithdrawPLGR, error) {
	event := new(TokengoWithdrawPLGR)
	if err := _Tokengo.contract.UnpackLog(event, "WithdrawPLGR", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
