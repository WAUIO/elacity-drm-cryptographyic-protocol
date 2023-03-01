// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IERC2981EnhancedRoyaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type IERC2981EnhancedRoyaltyInfo struct {
	Receiver common.Address
	Amount   *big.Int
}

// IOperativeAccessLevel is an auto generated low-level Go binding around an user-defined struct.
type IOperativeAccessLevel struct {
	HaveAccess  bool
	Entitlement *big.Int
}

// OperativeMetaData contains all meta data concerning the Operative contract.
var OperativeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIERC2981Enhanced.RoyaltyInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentId\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_TYPE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACCESS_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROYALTY_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_RIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkAccess\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"haveAccess\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"entitlement\",\"type\":\"uint256\"}],\"internalType\":\"structIOperative.AccessLevel[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OperativeABI is the input ABI used to generate the binding from.
// Deprecated: Use OperativeMetaData.ABI instead.
var OperativeABI = OperativeMetaData.ABI

// Operative is an auto generated Go binding around an Ethereum contract.
type Operative struct {
	OperativeCaller     // Read-only binding to the contract
	OperativeTransactor // Write-only binding to the contract
	OperativeFilterer   // Log filterer for contract events
}

// OperativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperativeSession struct {
	Contract     *Operative        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperativeCallerSession struct {
	Contract *OperativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OperativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperativeTransactorSession struct {
	Contract     *OperativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OperativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperativeRaw struct {
	Contract *Operative // Generic contract binding to access the raw methods on
}

// OperativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperativeCallerRaw struct {
	Contract *OperativeCaller // Generic read-only contract binding to access the raw methods on
}

// OperativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperativeTransactorRaw struct {
	Contract *OperativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperative creates a new instance of Operative, bound to a specific deployed contract.
func NewOperative(address common.Address, backend bind.ContractBackend) (*Operative, error) {
	contract, err := bindOperative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operative{OperativeCaller: OperativeCaller{contract: contract}, OperativeTransactor: OperativeTransactor{contract: contract}, OperativeFilterer: OperativeFilterer{contract: contract}}, nil
}

// NewOperativeCaller creates a new read-only instance of Operative, bound to a specific deployed contract.
func NewOperativeCaller(address common.Address, caller bind.ContractCaller) (*OperativeCaller, error) {
	contract, err := bindOperative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperativeCaller{contract: contract}, nil
}

// NewOperativeTransactor creates a new write-only instance of Operative, bound to a specific deployed contract.
func NewOperativeTransactor(address common.Address, transactor bind.ContractTransactor) (*OperativeTransactor, error) {
	contract, err := bindOperative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperativeTransactor{contract: contract}, nil
}

// NewOperativeFilterer creates a new log filterer instance of Operative, bound to a specific deployed contract.
func NewOperativeFilterer(address common.Address, filterer bind.ContractFilterer) (*OperativeFilterer, error) {
	contract, err := bindOperative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperativeFilterer{contract: contract}, nil
}

// bindOperative binds a generic wrapper to an already deployed contract.
func bindOperative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperativeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operative *OperativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operative.Contract.OperativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operative *OperativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operative.Contract.OperativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operative *OperativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operative.Contract.OperativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operative *OperativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operative *OperativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operative *OperativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operative.Contract.contract.Transact(opts, method, params...)
}

// ACCESSTOKEN is a free data retrieval call binding the contract method 0x3eb5002a.
//
// Solidity: function ACCESS_TOKEN() view returns(uint256)
func (_Operative *OperativeCaller) ACCESSTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "ACCESS_TOKEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ACCESSTOKEN is a free data retrieval call binding the contract method 0x3eb5002a.
//
// Solidity: function ACCESS_TOKEN() view returns(uint256)
func (_Operative *OperativeSession) ACCESSTOKEN() (*big.Int, error) {
	return _Operative.Contract.ACCESSTOKEN(&_Operative.CallOpts)
}

// ACCESSTOKEN is a free data retrieval call binding the contract method 0x3eb5002a.
//
// Solidity: function ACCESS_TOKEN() view returns(uint256)
func (_Operative *OperativeCallerSession) ACCESSTOKEN() (*big.Int, error) {
	return _Operative.Contract.ACCESSTOKEN(&_Operative.CallOpts)
}

// DISTRIBUTIONRIGHT is a free data retrieval call binding the contract method 0x342a6e92.
//
// Solidity: function DISTRIBUTION_RIGHT() view returns(uint256)
func (_Operative *OperativeCaller) DISTRIBUTIONRIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "DISTRIBUTION_RIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONRIGHT is a free data retrieval call binding the contract method 0x342a6e92.
//
// Solidity: function DISTRIBUTION_RIGHT() view returns(uint256)
func (_Operative *OperativeSession) DISTRIBUTIONRIGHT() (*big.Int, error) {
	return _Operative.Contract.DISTRIBUTIONRIGHT(&_Operative.CallOpts)
}

// DISTRIBUTIONRIGHT is a free data retrieval call binding the contract method 0x342a6e92.
//
// Solidity: function DISTRIBUTION_RIGHT() view returns(uint256)
func (_Operative *OperativeCallerSession) DISTRIBUTIONRIGHT() (*big.Int, error) {
	return _Operative.Contract.DISTRIBUTIONRIGHT(&_Operative.CallOpts)
}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(uint16)
func (_Operative *OperativeCaller) OPTYPE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "OP_TYPE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(uint16)
func (_Operative *OperativeSession) OPTYPE() (uint16, error) {
	return _Operative.Contract.OPTYPE(&_Operative.CallOpts)
}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(uint16)
func (_Operative *OperativeCallerSession) OPTYPE() (uint16, error) {
	return _Operative.Contract.OPTYPE(&_Operative.CallOpts)
}

// ROYALTYSHARE is a free data retrieval call binding the contract method 0x916ff717.
//
// Solidity: function ROYALTY_SHARE() view returns(uint256)
func (_Operative *OperativeCaller) ROYALTYSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "ROYALTY_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROYALTYSHARE is a free data retrieval call binding the contract method 0x916ff717.
//
// Solidity: function ROYALTY_SHARE() view returns(uint256)
func (_Operative *OperativeSession) ROYALTYSHARE() (*big.Int, error) {
	return _Operative.Contract.ROYALTYSHARE(&_Operative.CallOpts)
}

// ROYALTYSHARE is a free data retrieval call binding the contract method 0x916ff717.
//
// Solidity: function ROYALTY_SHARE() view returns(uint256)
func (_Operative *OperativeCallerSession) ROYALTYSHARE() (*big.Int, error) {
	return _Operative.Contract.ROYALTYSHARE(&_Operative.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Operative *OperativeCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Operative *OperativeSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Operative.Contract.BalanceOf(&_Operative.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Operative *OperativeCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Operative.Contract.BalanceOf(&_Operative.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Operative *OperativeCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Operative *OperativeSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Operative.Contract.BalanceOfBatch(&_Operative.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Operative *OperativeCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Operative.Contract.BalanceOfBatch(&_Operative.CallOpts, accounts, ids)
}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address account) view returns((bool,uint256)[])
func (_Operative *OperativeCaller) CheckAccess(opts *bind.CallOpts, account common.Address) ([]IOperativeAccessLevel, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "checkAccess", account)

	if err != nil {
		return *new([]IOperativeAccessLevel), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOperativeAccessLevel)).(*[]IOperativeAccessLevel)

	return out0, err

}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address account) view returns((bool,uint256)[])
func (_Operative *OperativeSession) CheckAccess(account common.Address) ([]IOperativeAccessLevel, error) {
	return _Operative.Contract.CheckAccess(&_Operative.CallOpts, account)
}

// CheckAccess is a free data retrieval call binding the contract method 0x466a0146.
//
// Solidity: function checkAccess(address account) view returns((bool,uint256)[])
func (_Operative *OperativeCallerSession) CheckAccess(account common.Address) ([]IOperativeAccessLevel, error) {
	return _Operative.Contract.CheckAccess(&_Operative.CallOpts, account)
}

// ContentId is a free data retrieval call binding the contract method 0xd97aa977.
//
// Solidity: function contentId() view returns(bytes16)
func (_Operative *OperativeCaller) ContentId(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "contentId")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// ContentId is a free data retrieval call binding the contract method 0xd97aa977.
//
// Solidity: function contentId() view returns(bytes16)
func (_Operative *OperativeSession) ContentId() ([16]byte, error) {
	return _Operative.Contract.ContentId(&_Operative.CallOpts)
}

// ContentId is a free data retrieval call binding the contract method 0xd97aa977.
//
// Solidity: function contentId() view returns(bytes16)
func (_Operative *OperativeCallerSession) ContentId() ([16]byte, error) {
	return _Operative.Contract.ContentId(&_Operative.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Operative *OperativeCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Operative *OperativeSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Operative.Contract.IsApprovedForAll(&_Operative.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Operative *OperativeCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Operative.Contract.IsApprovedForAll(&_Operative.CallOpts, account, operator)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0xcef6d368.
//
// Solidity: function royaltyInfo(uint256 _salePrice) view returns((address,uint256)[])
func (_Operative *OperativeCaller) RoyaltyInfo(opts *bind.CallOpts, _salePrice *big.Int) ([]IERC2981EnhancedRoyaltyInfo, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "royaltyInfo", _salePrice)

	if err != nil {
		return *new([]IERC2981EnhancedRoyaltyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC2981EnhancedRoyaltyInfo)).(*[]IERC2981EnhancedRoyaltyInfo)

	return out0, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0xcef6d368.
//
// Solidity: function royaltyInfo(uint256 _salePrice) view returns((address,uint256)[])
func (_Operative *OperativeSession) RoyaltyInfo(_salePrice *big.Int) ([]IERC2981EnhancedRoyaltyInfo, error) {
	return _Operative.Contract.RoyaltyInfo(&_Operative.CallOpts, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0xcef6d368.
//
// Solidity: function royaltyInfo(uint256 _salePrice) view returns((address,uint256)[])
func (_Operative *OperativeCallerSession) RoyaltyInfo(_salePrice *big.Int) ([]IERC2981EnhancedRoyaltyInfo, error) {
	return _Operative.Contract.RoyaltyInfo(&_Operative.CallOpts, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Operative *OperativeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Operative.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Operative *OperativeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Operative.Contract.SupportsInterface(&_Operative.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Operative *OperativeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Operative.Contract.SupportsInterface(&_Operative.CallOpts, interfaceId)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Operative *OperativeTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Operative *OperativeSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.Contract.SafeBatchTransferFrom(&_Operative.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Operative *OperativeTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.Contract.SafeBatchTransferFrom(&_Operative.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Operative *OperativeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Operative *OperativeSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.Contract.SafeTransferFrom(&_Operative.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Operative *OperativeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Operative.Contract.SafeTransferFrom(&_Operative.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Operative *OperativeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Operative.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Operative *OperativeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Operative.Contract.SetApprovalForAll(&_Operative.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Operative *OperativeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Operative.Contract.SetApprovalForAll(&_Operative.TransactOpts, operator, approved)
}

// OperativeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Operative contract.
type OperativeApprovalForAllIterator struct {
	Event *OperativeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OperativeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperativeApprovalForAll)
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
		it.Event = new(OperativeApprovalForAll)
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
func (it *OperativeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperativeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperativeApprovalForAll represents a ApprovalForAll event raised by the Operative contract.
type OperativeApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Operative *OperativeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*OperativeApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Operative.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OperativeApprovalForAllIterator{contract: _Operative.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Operative *OperativeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OperativeApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Operative.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperativeApprovalForAll)
				if err := _Operative.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Operative *OperativeFilterer) ParseApprovalForAll(log types.Log) (*OperativeApprovalForAll, error) {
	event := new(OperativeApprovalForAll)
	if err := _Operative.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperativeTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Operative contract.
type OperativeTransferBatchIterator struct {
	Event *OperativeTransferBatch // Event containing the contract specifics and raw log

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
func (it *OperativeTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperativeTransferBatch)
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
		it.Event = new(OperativeTransferBatch)
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
func (it *OperativeTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperativeTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperativeTransferBatch represents a TransferBatch event raised by the Operative contract.
type OperativeTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Operative *OperativeFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*OperativeTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operative.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OperativeTransferBatchIterator{contract: _Operative.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Operative *OperativeFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *OperativeTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operative.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperativeTransferBatch)
				if err := _Operative.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Operative *OperativeFilterer) ParseTransferBatch(log types.Log) (*OperativeTransferBatch, error) {
	event := new(OperativeTransferBatch)
	if err := _Operative.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperativeTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Operative contract.
type OperativeTransferSingleIterator struct {
	Event *OperativeTransferSingle // Event containing the contract specifics and raw log

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
func (it *OperativeTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperativeTransferSingle)
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
		it.Event = new(OperativeTransferSingle)
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
func (it *OperativeTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperativeTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperativeTransferSingle represents a TransferSingle event raised by the Operative contract.
type OperativeTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Operative *OperativeFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*OperativeTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operative.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OperativeTransferSingleIterator{contract: _Operative.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Operative *OperativeFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *OperativeTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Operative.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperativeTransferSingle)
				if err := _Operative.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Operative *OperativeFilterer) ParseTransferSingle(log types.Log) (*OperativeTransferSingle, error) {
	event := new(OperativeTransferSingle)
	if err := _Operative.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperativeURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Operative contract.
type OperativeURIIterator struct {
	Event *OperativeURI // Event containing the contract specifics and raw log

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
func (it *OperativeURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperativeURI)
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
		it.Event = new(OperativeURI)
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
func (it *OperativeURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperativeURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperativeURI represents a URI event raised by the Operative contract.
type OperativeURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Operative *OperativeFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*OperativeURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Operative.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &OperativeURIIterator{contract: _Operative.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Operative *OperativeFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *OperativeURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Operative.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperativeURI)
				if err := _Operative.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Operative *OperativeFilterer) ParseURI(log types.Log) (*OperativeURI, error) {
	event := new(OperativeURI)
	if err := _Operative.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
