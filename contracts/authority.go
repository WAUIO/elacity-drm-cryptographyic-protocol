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

// EIP712Domain is an auto generated low-level Go binding around an user-defined struct.
type EIP712Domain struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}

// ILicenseProviderLicense is an auto generated low-level Go binding around an user-defined struct.
type ILicenseProviderLicense struct {
	Issuer    common.Address
	ToAddress common.Address
	Entity    IPEntity
	Key       [][16]byte
}

// IPEntity is an auto generated low-level Go binding around an user-defined struct.
type IPEntity struct {
	ContentId [16]byte
	Ledger    common.Address
	TokenId   *big.Int
}

// LicenseRequest is an auto generated low-level Go binding around an user-defined struct.
type LicenseRequest struct {
	Entitlement string
	Entity      IPEntity
}

// SignedLicenseRequest is an auto generated low-level Go binding around an user-defined struct.
type SignedLicenseRequest struct {
	Request   LicenseRequest
	Signature []byte
}

// AuthorityGatewayMetaData contains all meta data concerning the AuthorityGateway contract.
var AuthorityGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ipHash\",\"type\":\"bytes32\"}],\"name\":\"IPRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tkId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tkId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tkId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"ItemUnlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"PaymentLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"}],\"internalType\":\"structEIP712Domain\",\"name\":\"_input\",\"type\":\"tuple\"}],\"name\":\"GET_EIP712DOMAIN_PACKETHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"_input\",\"type\":\"tuple\"}],\"name\":\"GET_IPENTITY_PACKETHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"entitlement\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"entity\",\"type\":\"tuple\"}],\"internalType\":\"structLicenseRequest\",\"name\":\"_input\",\"type\":\"tuple\"}],\"name\":\"GET_LICENSEREQUEST_PACKETHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"entitlement\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"entity\",\"type\":\"tuple\"}],\"internalType\":\"structLicenseRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSignedLicenseRequest\",\"name\":\"_input\",\"type\":\"tuple\"}],\"name\":\"GET_SIGNEDLICENSEREQUEST_PACKETHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LICENSE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"}],\"name\":\"__displayKey\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorage\",\"outputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"bytes16\",\"name\":\"_contentId\",\"type\":\"bytes16\"}],\"name\":\"registerIP\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"bytes16\",\"name\":\"_contentId\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"_ipKey\",\"type\":\"bytes16\"}],\"name\":\"registerIPWithKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"entitlement\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"entity\",\"type\":\"tuple\"}],\"internalType\":\"structLicenseRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSignedLicenseRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"verifyLicenseRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"_dataStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"sellAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"sellAccessOnBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"buyAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"}],\"name\":\"buyAccess\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"withdrawListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"entitlement\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"entity\",\"type\":\"tuple\"}],\"internalType\":\"structLicenseRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSignedLicenseRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"acquireLicense\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"contentId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIPEntity\",\"name\":\"entity\",\"type\":\"tuple\"},{\"internalType\":\"bytes16[]\",\"name\":\"key\",\"type\":\"bytes16[]\"}],\"internalType\":\"structILicenseProvider.License\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ledger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"operative\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"sellersOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"op\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AuthorityGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthorityGatewayMetaData.ABI instead.
var AuthorityGatewayABI = AuthorityGatewayMetaData.ABI

// AuthorityGateway is an auto generated Go binding around an Ethereum contract.
type AuthorityGateway struct {
	AuthorityGatewayCaller     // Read-only binding to the contract
	AuthorityGatewayTransactor // Write-only binding to the contract
	AuthorityGatewayFilterer   // Log filterer for contract events
}

// AuthorityGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorityGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorityGatewaySession struct {
	Contract     *AuthorityGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityGatewayCallerSession struct {
	Contract *AuthorityGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AuthorityGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityGatewayTransactorSession struct {
	Contract     *AuthorityGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AuthorityGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityGatewayRaw struct {
	Contract *AuthorityGateway // Generic contract binding to access the raw methods on
}

// AuthorityGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityGatewayCallerRaw struct {
	Contract *AuthorityGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityGatewayTransactorRaw struct {
	Contract *AuthorityGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorityGateway creates a new instance of AuthorityGateway, bound to a specific deployed contract.
func NewAuthorityGateway(address common.Address, backend bind.ContractBackend) (*AuthorityGateway, error) {
	contract, err := bindAuthorityGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthorityGateway{AuthorityGatewayCaller: AuthorityGatewayCaller{contract: contract}, AuthorityGatewayTransactor: AuthorityGatewayTransactor{contract: contract}, AuthorityGatewayFilterer: AuthorityGatewayFilterer{contract: contract}}, nil
}

// NewAuthorityGatewayCaller creates a new read-only instance of AuthorityGateway, bound to a specific deployed contract.
func NewAuthorityGatewayCaller(address common.Address, caller bind.ContractCaller) (*AuthorityGatewayCaller, error) {
	contract, err := bindAuthorityGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayCaller{contract: contract}, nil
}

// NewAuthorityGatewayTransactor creates a new write-only instance of AuthorityGateway, bound to a specific deployed contract.
func NewAuthorityGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityGatewayTransactor, error) {
	contract, err := bindAuthorityGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayTransactor{contract: contract}, nil
}

// NewAuthorityGatewayFilterer creates a new log filterer instance of AuthorityGateway, bound to a specific deployed contract.
func NewAuthorityGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorityGatewayFilterer, error) {
	contract, err := bindAuthorityGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayFilterer{contract: contract}, nil
}

// bindAuthorityGateway binds a generic wrapper to an already deployed contract.
func bindAuthorityGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorityGateway *AuthorityGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthorityGateway.Contract.AuthorityGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorityGateway *AuthorityGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.AuthorityGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorityGateway *AuthorityGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.AuthorityGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthorityGateway *AuthorityGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthorityGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthorityGateway *AuthorityGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthorityGateway *AuthorityGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AuthorityGateway.Contract.DEFAULTADMINROLE(&_AuthorityGateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AuthorityGateway.Contract.DEFAULTADMINROLE(&_AuthorityGateway.CallOpts)
}

// GETEIP712DOMAINPACKETHASH is a free data retrieval call binding the contract method 0x98bc186a.
//
// Solidity: function GET_EIP712DOMAIN_PACKETHASH((string,string,uint256,address) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) GETEIP712DOMAINPACKETHASH(opts *bind.CallOpts, _input EIP712Domain) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "GET_EIP712DOMAIN_PACKETHASH", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GETEIP712DOMAINPACKETHASH is a free data retrieval call binding the contract method 0x98bc186a.
//
// Solidity: function GET_EIP712DOMAIN_PACKETHASH((string,string,uint256,address) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) GETEIP712DOMAINPACKETHASH(_input EIP712Domain) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETEIP712DOMAINPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETEIP712DOMAINPACKETHASH is a free data retrieval call binding the contract method 0x98bc186a.
//
// Solidity: function GET_EIP712DOMAIN_PACKETHASH((string,string,uint256,address) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) GETEIP712DOMAINPACKETHASH(_input EIP712Domain) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETEIP712DOMAINPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETIPENTITYPACKETHASH is a free data retrieval call binding the contract method 0xd4b84d8a.
//
// Solidity: function GET_IPENTITY_PACKETHASH((bytes16,address,uint256) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) GETIPENTITYPACKETHASH(opts *bind.CallOpts, _input IPEntity) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "GET_IPENTITY_PACKETHASH", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GETIPENTITYPACKETHASH is a free data retrieval call binding the contract method 0xd4b84d8a.
//
// Solidity: function GET_IPENTITY_PACKETHASH((bytes16,address,uint256) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) GETIPENTITYPACKETHASH(_input IPEntity) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETIPENTITYPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETIPENTITYPACKETHASH is a free data retrieval call binding the contract method 0xd4b84d8a.
//
// Solidity: function GET_IPENTITY_PACKETHASH((bytes16,address,uint256) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) GETIPENTITYPACKETHASH(_input IPEntity) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETIPENTITYPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xe8631f05.
//
// Solidity: function GET_LICENSEREQUEST_PACKETHASH((string,(bytes16,address,uint256)) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) GETLICENSEREQUESTPACKETHASH(opts *bind.CallOpts, _input LicenseRequest) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "GET_LICENSEREQUEST_PACKETHASH", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GETLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xe8631f05.
//
// Solidity: function GET_LICENSEREQUEST_PACKETHASH((string,(bytes16,address,uint256)) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) GETLICENSEREQUESTPACKETHASH(_input LicenseRequest) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETLICENSEREQUESTPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xe8631f05.
//
// Solidity: function GET_LICENSEREQUEST_PACKETHASH((string,(bytes16,address,uint256)) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) GETLICENSEREQUESTPACKETHASH(_input LicenseRequest) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETLICENSEREQUESTPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETSIGNEDLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xaad260fd.
//
// Solidity: function GET_SIGNEDLICENSEREQUEST_PACKETHASH(((string,(bytes16,address,uint256)),bytes) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) GETSIGNEDLICENSEREQUESTPACKETHASH(opts *bind.CallOpts, _input SignedLicenseRequest) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "GET_SIGNEDLICENSEREQUEST_PACKETHASH", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GETSIGNEDLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xaad260fd.
//
// Solidity: function GET_SIGNEDLICENSEREQUEST_PACKETHASH(((string,(bytes16,address,uint256)),bytes) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) GETSIGNEDLICENSEREQUESTPACKETHASH(_input SignedLicenseRequest) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETSIGNEDLICENSEREQUESTPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// GETSIGNEDLICENSEREQUESTPACKETHASH is a free data retrieval call binding the contract method 0xaad260fd.
//
// Solidity: function GET_SIGNEDLICENSEREQUEST_PACKETHASH(((string,(bytes16,address,uint256)),bytes) _input) pure returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) GETSIGNEDLICENSEREQUESTPACKETHASH(_input SignedLicenseRequest) ([32]byte, error) {
	return _AuthorityGateway.Contract.GETSIGNEDLICENSEREQUESTPACKETHASH(&_AuthorityGateway.CallOpts, _input)
}

// LICENSEMANAGERROLE is a free data retrieval call binding the contract method 0x23efa822.
//
// Solidity: function LICENSE_MANAGER_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) LICENSEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "LICENSE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LICENSEMANAGERROLE is a free data retrieval call binding the contract method 0x23efa822.
//
// Solidity: function LICENSE_MANAGER_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) LICENSEMANAGERROLE() ([32]byte, error) {
	return _AuthorityGateway.Contract.LICENSEMANAGERROLE(&_AuthorityGateway.CallOpts)
}

// LICENSEMANAGERROLE is a free data retrieval call binding the contract method 0x23efa822.
//
// Solidity: function LICENSE_MANAGER_ROLE() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) LICENSEMANAGERROLE() ([32]byte, error) {
	return _AuthorityGateway.Contract.LICENSEMANAGERROLE(&_AuthorityGateway.CallOpts)
}

// DisplayKey is a free data retrieval call binding the contract method 0x892334dd.
//
// Solidity: function __displayKey(address ledger, bytes16 contentId) view returns(bytes16)
func (_AuthorityGateway *AuthorityGatewayCaller) DisplayKey(opts *bind.CallOpts, ledger common.Address, contentId [16]byte) ([16]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "__displayKey", ledger, contentId)

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// DisplayKey is a free data retrieval call binding the contract method 0x892334dd.
//
// Solidity: function __displayKey(address ledger, bytes16 contentId) view returns(bytes16)
func (_AuthorityGateway *AuthorityGatewaySession) DisplayKey(ledger common.Address, contentId [16]byte) ([16]byte, error) {
	return _AuthorityGateway.Contract.DisplayKey(&_AuthorityGateway.CallOpts, ledger, contentId)
}

// DisplayKey is a free data retrieval call binding the contract method 0x892334dd.
//
// Solidity: function __displayKey(address ledger, bytes16 contentId) view returns(bytes16)
func (_AuthorityGateway *AuthorityGatewayCallerSession) DisplayKey(ledger common.Address, contentId [16]byte) ([16]byte, error) {
	return _AuthorityGateway.Contract.DisplayKey(&_AuthorityGateway.CallOpts, ledger, contentId)
}

// AcquireLicense is a free data retrieval call binding the contract method 0x5923143c.
//
// Solidity: function acquireLicense(((string,(bytes16,address,uint256)),bytes) req) view returns((address,address,(bytes16,address,uint256),bytes16[]))
func (_AuthorityGateway *AuthorityGatewayCaller) AcquireLicense(opts *bind.CallOpts, req SignedLicenseRequest) (ILicenseProviderLicense, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "acquireLicense", req)

	if err != nil {
		return *new(ILicenseProviderLicense), err
	}

	out0 := *abi.ConvertType(out[0], new(ILicenseProviderLicense)).(*ILicenseProviderLicense)

	return out0, err

}

// AcquireLicense is a free data retrieval call binding the contract method 0x5923143c.
//
// Solidity: function acquireLicense(((string,(bytes16,address,uint256)),bytes) req) view returns((address,address,(bytes16,address,uint256),bytes16[]))
func (_AuthorityGateway *AuthorityGatewaySession) AcquireLicense(req SignedLicenseRequest) (ILicenseProviderLicense, error) {
	return _AuthorityGateway.Contract.AcquireLicense(&_AuthorityGateway.CallOpts, req)
}

// AcquireLicense is a free data retrieval call binding the contract method 0x5923143c.
//
// Solidity: function acquireLicense(((string,(bytes16,address,uint256)),bytes) req) view returns((address,address,(bytes16,address,uint256),bytes16[]))
func (_AuthorityGateway *AuthorityGatewayCallerSession) AcquireLicense(req SignedLicenseRequest) (ILicenseProviderLicense, error) {
	return _AuthorityGateway.Contract.AcquireLicense(&_AuthorityGateway.CallOpts, req)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_AuthorityGateway *AuthorityGatewayCaller) DataStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "dataStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_AuthorityGateway *AuthorityGatewaySession) DataStorage() (common.Address, error) {
	return _AuthorityGateway.Contract.DataStorage(&_AuthorityGateway.CallOpts)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_AuthorityGateway *AuthorityGatewayCallerSession) DataStorage() (common.Address, error) {
	return _AuthorityGateway.Contract.DataStorage(&_AuthorityGateway.CallOpts)
}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) DomainHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "domainHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) DomainHash() ([32]byte, error) {
	return _AuthorityGateway.Contract.DomainHash(&_AuthorityGateway.CallOpts)
}

// DomainHash is a free data retrieval call binding the contract method 0xdfe86ac5.
//
// Solidity: function domainHash() view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) DomainHash() ([32]byte, error) {
	return _AuthorityGateway.Contract.DomainHash(&_AuthorityGateway.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AuthorityGateway.Contract.GetRoleAdmin(&_AuthorityGateway.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuthorityGateway *AuthorityGatewayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AuthorityGateway.Contract.GetRoleAdmin(&_AuthorityGateway.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuthorityGateway *AuthorityGatewayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuthorityGateway *AuthorityGatewaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AuthorityGateway.Contract.HasRole(&_AuthorityGateway.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuthorityGateway *AuthorityGatewayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AuthorityGateway.Contract.HasRole(&_AuthorityGateway.CallOpts, role, account)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address op, uint256 tokenId, address seller) view returns(uint256, uint256, address)
func (_AuthorityGateway *AuthorityGatewayCaller) Listings(opts *bind.CallOpts, op common.Address, tokenId *big.Int, seller common.Address) (*big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "listings", op, tokenId, seller)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address op, uint256 tokenId, address seller) view returns(uint256, uint256, address)
func (_AuthorityGateway *AuthorityGatewaySession) Listings(op common.Address, tokenId *big.Int, seller common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _AuthorityGateway.Contract.Listings(&_AuthorityGateway.CallOpts, op, tokenId, seller)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address op, uint256 tokenId, address seller) view returns(uint256, uint256, address)
func (_AuthorityGateway *AuthorityGatewayCallerSession) Listings(op common.Address, tokenId *big.Int, seller common.Address) (*big.Int, *big.Int, common.Address, error) {
	return _AuthorityGateway.Contract.Listings(&_AuthorityGateway.CallOpts, op, tokenId, seller)
}

// Operative is a free data retrieval call binding the contract method 0x70611dd2.
//
// Solidity: function operative(address ledger, uint256 tokenId) view returns(address)
func (_AuthorityGateway *AuthorityGatewayCaller) Operative(opts *bind.CallOpts, ledger common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "operative", ledger, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operative is a free data retrieval call binding the contract method 0x70611dd2.
//
// Solidity: function operative(address ledger, uint256 tokenId) view returns(address)
func (_AuthorityGateway *AuthorityGatewaySession) Operative(ledger common.Address, tokenId *big.Int) (common.Address, error) {
	return _AuthorityGateway.Contract.Operative(&_AuthorityGateway.CallOpts, ledger, tokenId)
}

// Operative is a free data retrieval call binding the contract method 0x70611dd2.
//
// Solidity: function operative(address ledger, uint256 tokenId) view returns(address)
func (_AuthorityGateway *AuthorityGatewayCallerSession) Operative(ledger common.Address, tokenId *big.Int) (common.Address, error) {
	return _AuthorityGateway.Contract.Operative(&_AuthorityGateway.CallOpts, ledger, tokenId)
}

// SellersOf is a free data retrieval call binding the contract method 0x997eab2d.
//
// Solidity: function sellersOf(address op, uint256 tokenId) view returns(address[])
func (_AuthorityGateway *AuthorityGatewayCaller) SellersOf(opts *bind.CallOpts, op common.Address, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "sellersOf", op, tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SellersOf is a free data retrieval call binding the contract method 0x997eab2d.
//
// Solidity: function sellersOf(address op, uint256 tokenId) view returns(address[])
func (_AuthorityGateway *AuthorityGatewaySession) SellersOf(op common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _AuthorityGateway.Contract.SellersOf(&_AuthorityGateway.CallOpts, op, tokenId)
}

// SellersOf is a free data retrieval call binding the contract method 0x997eab2d.
//
// Solidity: function sellersOf(address op, uint256 tokenId) view returns(address[])
func (_AuthorityGateway *AuthorityGatewayCallerSession) SellersOf(op common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _AuthorityGateway.Contract.SellersOf(&_AuthorityGateway.CallOpts, op, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuthorityGateway *AuthorityGatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuthorityGateway *AuthorityGatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AuthorityGateway.Contract.SupportsInterface(&_AuthorityGateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuthorityGateway *AuthorityGatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AuthorityGateway.Contract.SupportsInterface(&_AuthorityGateway.CallOpts, interfaceId)
}

// VerifyLicenseRequest is a free data retrieval call binding the contract method 0xe2439183.
//
// Solidity: function verifyLicenseRequest(((string,(bytes16,address,uint256)),bytes) req) view returns(address)
func (_AuthorityGateway *AuthorityGatewayCaller) VerifyLicenseRequest(opts *bind.CallOpts, req SignedLicenseRequest) (common.Address, error) {
	var out []interface{}
	err := _AuthorityGateway.contract.Call(opts, &out, "verifyLicenseRequest", req)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyLicenseRequest is a free data retrieval call binding the contract method 0xe2439183.
//
// Solidity: function verifyLicenseRequest(((string,(bytes16,address,uint256)),bytes) req) view returns(address)
func (_AuthorityGateway *AuthorityGatewaySession) VerifyLicenseRequest(req SignedLicenseRequest) (common.Address, error) {
	return _AuthorityGateway.Contract.VerifyLicenseRequest(&_AuthorityGateway.CallOpts, req)
}

// VerifyLicenseRequest is a free data retrieval call binding the contract method 0xe2439183.
//
// Solidity: function verifyLicenseRequest(((string,(bytes16,address,uint256)),bytes) req) view returns(address)
func (_AuthorityGateway *AuthorityGatewayCallerSession) VerifyLicenseRequest(req SignedLicenseRequest) (common.Address, error) {
	return _AuthorityGateway.Contract.VerifyLicenseRequest(&_AuthorityGateway.CallOpts, req)
}

// BuyAccess is a paid mutator transaction binding the contract method 0x0ede2294.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) BuyAccess(opts *bind.TransactOpts, seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "buyAccess", seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// BuyAccess is a paid mutator transaction binding the contract method 0x0ede2294.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewaySession) BuyAccess(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.BuyAccess(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// BuyAccess is a paid mutator transaction binding the contract method 0x0ede2294.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) BuyAccess(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.BuyAccess(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// BuyAccess0 is a paid mutator transaction binding the contract method 0xf7580ad9.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken) payable returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) BuyAccess0(opts *bind.TransactOpts, seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "buyAccess0", seller, ledger, tokenId, _quantity, _pricePerToken)
}

// BuyAccess0 is a paid mutator transaction binding the contract method 0xf7580ad9.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken) payable returns()
func (_AuthorityGateway *AuthorityGatewaySession) BuyAccess0(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.BuyAccess0(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken)
}

// BuyAccess0 is a paid mutator transaction binding the contract method 0xf7580ad9.
//
// Solidity: function buyAccess(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken) payable returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) BuyAccess0(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.BuyAccess0(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.GrantRole(&_AuthorityGateway.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.GrantRole(&_AuthorityGateway.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorage) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) Initialize(opts *bind.TransactOpts, _dataStorage common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "initialize", _dataStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorage) returns()
func (_AuthorityGateway *AuthorityGatewaySession) Initialize(_dataStorage common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.Initialize(&_AuthorityGateway.TransactOpts, _dataStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorage) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) Initialize(_dataStorage common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.Initialize(&_AuthorityGateway.TransactOpts, _dataStorage)
}

// RegisterIP is a paid mutator transaction binding the contract method 0xca689cf0.
//
// Solidity: function registerIP(address ledger, bytes16 _contentId) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewayTransactor) RegisterIP(opts *bind.TransactOpts, ledger common.Address, _contentId [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "registerIP", ledger, _contentId)
}

// RegisterIP is a paid mutator transaction binding the contract method 0xca689cf0.
//
// Solidity: function registerIP(address ledger, bytes16 _contentId) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewaySession) RegisterIP(ledger common.Address, _contentId [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RegisterIP(&_AuthorityGateway.TransactOpts, ledger, _contentId)
}

// RegisterIP is a paid mutator transaction binding the contract method 0xca689cf0.
//
// Solidity: function registerIP(address ledger, bytes16 _contentId) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewayTransactorSession) RegisterIP(ledger common.Address, _contentId [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RegisterIP(&_AuthorityGateway.TransactOpts, ledger, _contentId)
}

// RegisterIPWithKey is a paid mutator transaction binding the contract method 0xfd59b1e4.
//
// Solidity: function registerIPWithKey(address ledger, bytes16 _contentId, bytes16 _ipKey) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewayTransactor) RegisterIPWithKey(opts *bind.TransactOpts, ledger common.Address, _contentId [16]byte, _ipKey [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "registerIPWithKey", ledger, _contentId, _ipKey)
}

// RegisterIPWithKey is a paid mutator transaction binding the contract method 0xfd59b1e4.
//
// Solidity: function registerIPWithKey(address ledger, bytes16 _contentId, bytes16 _ipKey) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewaySession) RegisterIPWithKey(ledger common.Address, _contentId [16]byte, _ipKey [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RegisterIPWithKey(&_AuthorityGateway.TransactOpts, ledger, _contentId, _ipKey)
}

// RegisterIPWithKey is a paid mutator transaction binding the contract method 0xfd59b1e4.
//
// Solidity: function registerIPWithKey(address ledger, bytes16 _contentId, bytes16 _ipKey) returns(bytes32, bytes16)
func (_AuthorityGateway *AuthorityGatewayTransactorSession) RegisterIPWithKey(ledger common.Address, _contentId [16]byte, _ipKey [16]byte) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RegisterIPWithKey(&_AuthorityGateway.TransactOpts, ledger, _contentId, _ipKey)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RenounceRole(&_AuthorityGateway.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RenounceRole(&_AuthorityGateway.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RevokeRole(&_AuthorityGateway.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.RevokeRole(&_AuthorityGateway.TransactOpts, role, account)
}

// SellAccess is a paid mutator transaction binding the contract method 0x9a3fa9f5.
//
// Solidity: function sellAccess(address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) SellAccess(opts *bind.TransactOpts, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "sellAccess", ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// SellAccess is a paid mutator transaction binding the contract method 0x9a3fa9f5.
//
// Solidity: function sellAccess(address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewaySession) SellAccess(ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.SellAccess(&_AuthorityGateway.TransactOpts, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// SellAccess is a paid mutator transaction binding the contract method 0x9a3fa9f5.
//
// Solidity: function sellAccess(address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) SellAccess(ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.SellAccess(&_AuthorityGateway.TransactOpts, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// SellAccessOnBehalf is a paid mutator transaction binding the contract method 0x3e0da94c.
//
// Solidity: function sellAccessOnBehalf(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) SellAccessOnBehalf(opts *bind.TransactOpts, seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "sellAccessOnBehalf", seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// SellAccessOnBehalf is a paid mutator transaction binding the contract method 0x3e0da94c.
//
// Solidity: function sellAccessOnBehalf(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewaySession) SellAccessOnBehalf(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.SellAccessOnBehalf(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// SellAccessOnBehalf is a paid mutator transaction binding the contract method 0x3e0da94c.
//
// Solidity: function sellAccessOnBehalf(address seller, address ledger, uint256 tokenId, uint256 _quantity, uint256 _pricePerToken, address _payToken) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) SellAccessOnBehalf(seller common.Address, ledger common.Address, tokenId *big.Int, _quantity *big.Int, _pricePerToken *big.Int, _payToken common.Address) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.SellAccessOnBehalf(&_AuthorityGateway.TransactOpts, seller, ledger, tokenId, _quantity, _pricePerToken, _payToken)
}

// WithdrawListing is a paid mutator transaction binding the contract method 0x3e65bbba.
//
// Solidity: function withdrawListing(address op, uint256 tokenId, uint256 quantity) returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) WithdrawListing(opts *bind.TransactOpts, op common.Address, tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.contract.Transact(opts, "withdrawListing", op, tokenId, quantity)
}

// WithdrawListing is a paid mutator transaction binding the contract method 0x3e65bbba.
//
// Solidity: function withdrawListing(address op, uint256 tokenId, uint256 quantity) returns()
func (_AuthorityGateway *AuthorityGatewaySession) WithdrawListing(op common.Address, tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.WithdrawListing(&_AuthorityGateway.TransactOpts, op, tokenId, quantity)
}

// WithdrawListing is a paid mutator transaction binding the contract method 0x3e65bbba.
//
// Solidity: function withdrawListing(address op, uint256 tokenId, uint256 quantity) returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) WithdrawListing(op common.Address, tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _AuthorityGateway.Contract.WithdrawListing(&_AuthorityGateway.TransactOpts, op, tokenId, quantity)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AuthorityGateway *AuthorityGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthorityGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AuthorityGateway *AuthorityGatewaySession) Receive() (*types.Transaction, error) {
	return _AuthorityGateway.Contract.Receive(&_AuthorityGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AuthorityGateway *AuthorityGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _AuthorityGateway.Contract.Receive(&_AuthorityGateway.TransactOpts)
}

// AuthorityGatewayIPRegisteredIterator is returned from FilterIPRegistered and is used to iterate over the raw logs and unpacked data for IPRegistered events raised by the AuthorityGateway contract.
type AuthorityGatewayIPRegisteredIterator struct {
	Event *AuthorityGatewayIPRegistered // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayIPRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayIPRegistered)
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
		it.Event = new(AuthorityGatewayIPRegistered)
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
func (it *AuthorityGatewayIPRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayIPRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayIPRegistered represents a IPRegistered event raised by the AuthorityGateway contract.
type AuthorityGatewayIPRegistered struct {
	Ledger    common.Address
	ContentId [16]byte
	IpHash    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIPRegistered is a free log retrieval operation binding the contract event 0x839cbad40c9e5c0919d31cb0168c7b9c9f67b47ff3bfb01c94649eea863c2317.
//
// Solidity: event IPRegistered(address indexed ledger, bytes16 contentId, bytes32 ipHash)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterIPRegistered(opts *bind.FilterOpts, ledger []common.Address) (*AuthorityGatewayIPRegisteredIterator, error) {

	var ledgerRule []interface{}
	for _, ledgerItem := range ledger {
		ledgerRule = append(ledgerRule, ledgerItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "IPRegistered", ledgerRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayIPRegisteredIterator{contract: _AuthorityGateway.contract, event: "IPRegistered", logs: logs, sub: sub}, nil
}

// WatchIPRegistered is a free log subscription operation binding the contract event 0x839cbad40c9e5c0919d31cb0168c7b9c9f67b47ff3bfb01c94649eea863c2317.
//
// Solidity: event IPRegistered(address indexed ledger, bytes16 contentId, bytes32 ipHash)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchIPRegistered(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayIPRegistered, ledger []common.Address) (event.Subscription, error) {

	var ledgerRule []interface{}
	for _, ledgerItem := range ledger {
		ledgerRule = append(ledgerRule, ledgerItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "IPRegistered", ledgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayIPRegistered)
				if err := _AuthorityGateway.contract.UnpackLog(event, "IPRegistered", log); err != nil {
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

// ParseIPRegistered is a log parse operation binding the contract event 0x839cbad40c9e5c0919d31cb0168c7b9c9f67b47ff3bfb01c94649eea863c2317.
//
// Solidity: event IPRegistered(address indexed ledger, bytes16 contentId, bytes32 ipHash)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseIPRegistered(log types.Log) (*AuthorityGatewayIPRegistered, error) {
	event := new(AuthorityGatewayIPRegistered)
	if err := _AuthorityGateway.contract.UnpackLog(event, "IPRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the AuthorityGateway contract.
type AuthorityGatewayItemListedIterator struct {
	Event *AuthorityGatewayItemListed // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayItemListed)
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
		it.Event = new(AuthorityGatewayItemListed)
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
func (it *AuthorityGatewayItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayItemListed represents a ItemListed event raised by the AuthorityGateway contract.
type AuthorityGatewayItemListed struct {
	Seller        common.Address
	Op            common.Address
	TkId          *big.Int
	Quantity      *big.Int
	PricePerToken *big.Int
	PayToken      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0x90aecdd7f5269ac7f11bea516b4768d0391e0a54aabc19aea64c7758104f66d2.
//
// Solidity: event ItemListed(address indexed seller, address indexed op, uint256 tkId, uint256 quantity, uint256 pricePerToken, address payToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterItemListed(opts *bind.FilterOpts, seller []common.Address, op []common.Address) (*AuthorityGatewayItemListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "ItemListed", sellerRule, opRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayItemListedIterator{contract: _AuthorityGateway.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0x90aecdd7f5269ac7f11bea516b4768d0391e0a54aabc19aea64c7758104f66d2.
//
// Solidity: event ItemListed(address indexed seller, address indexed op, uint256 tkId, uint256 quantity, uint256 pricePerToken, address payToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayItemListed, seller []common.Address, op []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "ItemListed", sellerRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayItemListed)
				if err := _AuthorityGateway.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0x90aecdd7f5269ac7f11bea516b4768d0391e0a54aabc19aea64c7758104f66d2.
//
// Solidity: event ItemListed(address indexed seller, address indexed op, uint256 tkId, uint256 quantity, uint256 pricePerToken, address payToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseItemListed(log types.Log) (*AuthorityGatewayItemListed, error) {
	event := new(AuthorityGatewayItemListed)
	if err := _AuthorityGateway.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the AuthorityGateway contract.
type AuthorityGatewayItemSoldIterator struct {
	Event *AuthorityGatewayItemSold // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayItemSold)
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
		it.Event = new(AuthorityGatewayItemSold)
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
func (it *AuthorityGatewayItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayItemSold represents a ItemSold event raised by the AuthorityGateway contract.
type AuthorityGatewayItemSold struct {
	Seller    common.Address
	Buyer     common.Address
	Op        common.Address
	TkId      *big.Int
	PayToken  common.Address
	UnitPrice *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x60cd9eee664e26e142eb54813d426c273cd85605b8bfb72f707e4f2927b6a955.
//
// Solidity: event ItemSold(address seller, address indexed buyer, address indexed op, uint256 indexed tkId, address payToken, uint256 unitPrice, uint256 price)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterItemSold(opts *bind.FilterOpts, buyer []common.Address, op []common.Address, tkId []*big.Int) (*AuthorityGatewayItemSoldIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}
	var tkIdRule []interface{}
	for _, tkIdItem := range tkId {
		tkIdRule = append(tkIdRule, tkIdItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "ItemSold", buyerRule, opRule, tkIdRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayItemSoldIterator{contract: _AuthorityGateway.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x60cd9eee664e26e142eb54813d426c273cd85605b8bfb72f707e4f2927b6a955.
//
// Solidity: event ItemSold(address seller, address indexed buyer, address indexed op, uint256 indexed tkId, address payToken, uint256 unitPrice, uint256 price)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayItemSold, buyer []common.Address, op []common.Address, tkId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}
	var tkIdRule []interface{}
	for _, tkIdItem := range tkId {
		tkIdRule = append(tkIdRule, tkIdItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "ItemSold", buyerRule, opRule, tkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayItemSold)
				if err := _AuthorityGateway.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x60cd9eee664e26e142eb54813d426c273cd85605b8bfb72f707e4f2927b6a955.
//
// Solidity: event ItemSold(address seller, address indexed buyer, address indexed op, uint256 indexed tkId, address payToken, uint256 unitPrice, uint256 price)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseItemSold(log types.Log) (*AuthorityGatewayItemSold, error) {
	event := new(AuthorityGatewayItemSold)
	if err := _AuthorityGateway.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayItemUnlistedIterator is returned from FilterItemUnlisted and is used to iterate over the raw logs and unpacked data for ItemUnlisted events raised by the AuthorityGateway contract.
type AuthorityGatewayItemUnlistedIterator struct {
	Event *AuthorityGatewayItemUnlisted // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayItemUnlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayItemUnlisted)
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
		it.Event = new(AuthorityGatewayItemUnlisted)
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
func (it *AuthorityGatewayItemUnlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayItemUnlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayItemUnlisted represents a ItemUnlisted event raised by the AuthorityGateway contract.
type AuthorityGatewayItemUnlisted struct {
	Seller   common.Address
	Op       common.Address
	TkId     *big.Int
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUnlisted is a free log retrieval operation binding the contract event 0xdb6bedce61ad043a5e9d9ac95f248702233e64e5818e58734aa38e7fd86db415.
//
// Solidity: event ItemUnlisted(address indexed seller, address indexed op, uint256 indexed tkId, uint256 quantity)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterItemUnlisted(opts *bind.FilterOpts, seller []common.Address, op []common.Address, tkId []*big.Int) (*AuthorityGatewayItemUnlistedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}
	var tkIdRule []interface{}
	for _, tkIdItem := range tkId {
		tkIdRule = append(tkIdRule, tkIdItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "ItemUnlisted", sellerRule, opRule, tkIdRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayItemUnlistedIterator{contract: _AuthorityGateway.contract, event: "ItemUnlisted", logs: logs, sub: sub}, nil
}

// WatchItemUnlisted is a free log subscription operation binding the contract event 0xdb6bedce61ad043a5e9d9ac95f248702233e64e5818e58734aa38e7fd86db415.
//
// Solidity: event ItemUnlisted(address indexed seller, address indexed op, uint256 indexed tkId, uint256 quantity)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchItemUnlisted(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayItemUnlisted, seller []common.Address, op []common.Address, tkId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}
	var tkIdRule []interface{}
	for _, tkIdItem := range tkId {
		tkIdRule = append(tkIdRule, tkIdItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "ItemUnlisted", sellerRule, opRule, tkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayItemUnlisted)
				if err := _AuthorityGateway.contract.UnpackLog(event, "ItemUnlisted", log); err != nil {
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

// ParseItemUnlisted is a log parse operation binding the contract event 0xdb6bedce61ad043a5e9d9ac95f248702233e64e5818e58734aa38e7fd86db415.
//
// Solidity: event ItemUnlisted(address indexed seller, address indexed op, uint256 indexed tkId, uint256 quantity)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseItemUnlisted(log types.Log) (*AuthorityGatewayItemUnlisted, error) {
	event := new(AuthorityGatewayItemUnlisted)
	if err := _AuthorityGateway.contract.UnpackLog(event, "ItemUnlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayPaymentLogIterator is returned from FilterPaymentLog and is used to iterate over the raw logs and unpacked data for PaymentLog events raised by the AuthorityGateway contract.
type AuthorityGatewayPaymentLogIterator struct {
	Event *AuthorityGatewayPaymentLog // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayPaymentLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayPaymentLog)
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
		it.Event = new(AuthorityGatewayPaymentLog)
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
func (it *AuthorityGatewayPaymentLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayPaymentLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayPaymentLog represents a PaymentLog event raised by the AuthorityGateway contract.
type AuthorityGatewayPaymentLog struct {
	From         common.Address
	To           common.Address
	Amount       *big.Int
	PaymentToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPaymentLog is a free log retrieval operation binding the contract event 0x69b13643ce291ff134c1f55603354c8a34bf63578d50af035122d94a48813338.
//
// Solidity: event PaymentLog(address indexed from, address indexed to, uint256 amount, address paymentToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterPaymentLog(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AuthorityGatewayPaymentLogIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "PaymentLog", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayPaymentLogIterator{contract: _AuthorityGateway.contract, event: "PaymentLog", logs: logs, sub: sub}, nil
}

// WatchPaymentLog is a free log subscription operation binding the contract event 0x69b13643ce291ff134c1f55603354c8a34bf63578d50af035122d94a48813338.
//
// Solidity: event PaymentLog(address indexed from, address indexed to, uint256 amount, address paymentToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchPaymentLog(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayPaymentLog, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "PaymentLog", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayPaymentLog)
				if err := _AuthorityGateway.contract.UnpackLog(event, "PaymentLog", log); err != nil {
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

// ParsePaymentLog is a log parse operation binding the contract event 0x69b13643ce291ff134c1f55603354c8a34bf63578d50af035122d94a48813338.
//
// Solidity: event PaymentLog(address indexed from, address indexed to, uint256 amount, address paymentToken)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParsePaymentLog(log types.Log) (*AuthorityGatewayPaymentLog, error) {
	event := new(AuthorityGatewayPaymentLog)
	if err := _AuthorityGateway.contract.UnpackLog(event, "PaymentLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the AuthorityGateway contract.
type AuthorityGatewayPaymentReceivedIterator struct {
	Event *AuthorityGatewayPaymentReceived // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayPaymentReceived)
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
		it.Event = new(AuthorityGatewayPaymentReceived)
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
func (it *AuthorityGatewayPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayPaymentReceived represents a PaymentReceived event raised by the AuthorityGateway contract.
type AuthorityGatewayPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*AuthorityGatewayPaymentReceivedIterator, error) {

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayPaymentReceivedIterator{contract: _AuthorityGateway.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayPaymentReceived)
				if err := _AuthorityGateway.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParsePaymentReceived(log types.Log) (*AuthorityGatewayPaymentReceived, error) {
	event := new(AuthorityGatewayPaymentReceived)
	if err := _AuthorityGateway.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AuthorityGateway contract.
type AuthorityGatewayRoleAdminChangedIterator struct {
	Event *AuthorityGatewayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayRoleAdminChanged)
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
		it.Event = new(AuthorityGatewayRoleAdminChanged)
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
func (it *AuthorityGatewayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayRoleAdminChanged represents a RoleAdminChanged event raised by the AuthorityGateway contract.
type AuthorityGatewayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AuthorityGatewayRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayRoleAdminChangedIterator{contract: _AuthorityGateway.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayRoleAdminChanged)
				if err := _AuthorityGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseRoleAdminChanged(log types.Log) (*AuthorityGatewayRoleAdminChanged, error) {
	event := new(AuthorityGatewayRoleAdminChanged)
	if err := _AuthorityGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AuthorityGateway contract.
type AuthorityGatewayRoleGrantedIterator struct {
	Event *AuthorityGatewayRoleGranted // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayRoleGranted)
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
		it.Event = new(AuthorityGatewayRoleGranted)
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
func (it *AuthorityGatewayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayRoleGranted represents a RoleGranted event raised by the AuthorityGateway contract.
type AuthorityGatewayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AuthorityGatewayRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayRoleGrantedIterator{contract: _AuthorityGateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayRoleGranted)
				if err := _AuthorityGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseRoleGranted(log types.Log) (*AuthorityGatewayRoleGranted, error) {
	event := new(AuthorityGatewayRoleGranted)
	if err := _AuthorityGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthorityGatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AuthorityGateway contract.
type AuthorityGatewayRoleRevokedIterator struct {
	Event *AuthorityGatewayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AuthorityGatewayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorityGatewayRoleRevoked)
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
		it.Event = new(AuthorityGatewayRoleRevoked)
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
func (it *AuthorityGatewayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorityGatewayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorityGatewayRoleRevoked represents a RoleRevoked event raised by the AuthorityGateway contract.
type AuthorityGatewayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AuthorityGatewayRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuthorityGateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AuthorityGatewayRoleRevokedIterator{contract: _AuthorityGateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AuthorityGatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuthorityGateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorityGatewayRoleRevoked)
				if err := _AuthorityGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuthorityGateway *AuthorityGatewayFilterer) ParseRoleRevoked(log types.Log) (*AuthorityGatewayRoleRevoked, error) {
	event := new(AuthorityGatewayRoleRevoked)
	if err := _AuthorityGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
