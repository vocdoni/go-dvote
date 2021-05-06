// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EnsRegistryWithFallbackABI is the input ABI used to generate the binding from.
const EnsRegistryWithFallbackABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EnsRegistryWithFallbackBin is the compiled bytecode used for deploying new contracts.
var EnsRegistryWithFallbackBin = "0x608060405234801561001057600080fd5b5060008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b03191633179055610a97806100596000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80635b0fc9c3116100715780635b0fc9c31461015d5780635ef2c7f014610170578063a22cb46514610183578063cf40882314610196578063e985e9c5146101a9578063f79fe538146101c9576100b4565b80630178b8bf146100b957806302571be3146100e257806306ab5923146100f557806314ab90381461011557806316a25cbd1461012a5780631896f70a1461014a575b600080fd5b6100cc6100c736600461088a565b6101dc565b6040516100d991906109e6565b60405180910390f35b6100cc6100f036600461088a565b6101fd565b610108610103366004610924565b61022d565b6040516100d99190610a05565b6101286101233660046109b5565b610317565b005b61013d61013836600461088a565b6103f6565b6040516100d99190610a34565b6101286101583660046108a2565b61041c565b61012861016b3660046108a2565b6104ee565b61012861017e36600461095c565b61059d565b61012861019136600461084f565b6105bf565b6101286101a43660046108c6565b61062e565b6101bc6101b736600461081b565b610649565b6040516100d991906109fa565b6101bc6101d736600461088a565b610677565b6000818152602081905260409020600101546001600160a01b03165b919050565b6000818152602081905260408120546001600160a01b0316308114156102275760009150506101f8565b92915050565b60008381526020819052604081205484906001600160a01b03163381148061027857506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61029d5760405162461bcd60e51b815260040161029490610a0e565b60405180910390fd5b600086866040516020016102b29291906109d8565b6040516020818303038152906040528051906020012090506102d48186610694565b85877fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e828760405161030591906109e6565b60405180910390a39695505050505050565b60008281526020819052604090205482906001600160a01b03163381148061036257506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b61037e5760405162461bcd60e51b815260040161029490610a0e565b837f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68846040516103ae9190610a34565b60405180910390a25050600091825260208290526040909120600101805467ffffffffffffffff909216600160a01b0267ffffffffffffffff60a01b19909216919091179055565b600090815260208190526040902060010154600160a01b900467ffffffffffffffff1690565b60008281526020819052604090205482906001600160a01b03163381148061046757506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6104835760405162461bcd60e51b815260040161029490610a0e565b837f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0846040516104b391906109e6565b60405180910390a2505060009182526020829052604090912060010180546001600160a01b0319166001600160a01b03909216919091179055565b60008281526020819052604090205482906001600160a01b03163381148061053957506001600160a01b038116600090815260016020908152604080832033845290915290205460ff165b6105555760405162461bcd60e51b815260040161029490610a0e565b61055f8484610694565b837fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d2668460405161058f91906109e6565b60405180910390a250505050565b60006105aa86868661022d565b90506105b78184846106c2565b505050505050565b3360008181526001602090815260408083206001600160a01b038716808552925291829020805460ff191685151517905590519091907f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31906106229085906109fa565b60405180910390a35050565b61063884846104ee565b6106438483836106c2565b50505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b6000908152602081905260409020546001600160a01b0316151590565b60009182526020829052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6000838152602081905260409020600101546001600160a01b0383811691161461074b576000838152602081905260409081902060010180546001600160a01b0319166001600160a01b0385161790555183907f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0906107429085906109e6565b60405180910390a25b60008381526020819052604090206001015467ffffffffffffffff828116600160a01b90920416146107e75760008381526020819052604090819020600101805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff8516021790555183907f1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68906107de908490610a34565b60405180910390a25b505050565b80356001600160a01b038116811461022757600080fd5b803567ffffffffffffffff8116811461022757600080fd5b6000806040838503121561082d578182fd5b61083784846107ec565b915061084684602085016107ec565b90509250929050565b60008060408385031215610861578182fd5b61086b84846107ec565b91506020830135801515811461087f578182fd5b809150509250929050565b60006020828403121561089b578081fd5b5035919050565b600080604083850312156108b4578182fd5b82359150602083013561087f81610a49565b600080600080608085870312156108db578182fd5b8435935060208501356108ed81610a49565b925060408501356108fd81610a49565b9150606085013567ffffffffffffffff81168114610919578182fd5b939692955090935050565b600080600060608486031215610938578283fd5b8335925060208401359150604084013561095181610a49565b809150509250925092565b600080600080600060a08688031215610973578081fd5b853594506020860135935061098b87604088016107ec565b925061099a87606088016107ec565b91506109a98760808801610803565b90509295509295909350565b600080604083850312156109c7578182fd5b823591506108468460208501610803565b918252602082015260400190565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b67ffffffffffffffff91909116815260200190565b6001600160a01b0381168114610a5e57600080fd5b5056fea264697066735822122000a4777da7ded13dba7a494b31d08ba62a15329a81330f2a911b3ab400df88fd64736f6c634300060c0033"

// DeployEnsRegistryWithFallback deploys a new Ethereum contract, binding an instance of EnsRegistryWithFallback to it.
func DeployEnsRegistryWithFallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnsRegistryWithFallback, error) {
	parsed, err := abi.JSON(strings.NewReader(EnsRegistryWithFallbackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnsRegistryWithFallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnsRegistryWithFallback{EnsRegistryWithFallbackCaller: EnsRegistryWithFallbackCaller{contract: contract}, EnsRegistryWithFallbackTransactor: EnsRegistryWithFallbackTransactor{contract: contract}, EnsRegistryWithFallbackFilterer: EnsRegistryWithFallbackFilterer{contract: contract}}, nil
}

// EnsRegistryWithFallback is an auto generated Go binding around an Ethereum contract.
type EnsRegistryWithFallback struct {
	EnsRegistryWithFallbackCaller     // Read-only binding to the contract
	EnsRegistryWithFallbackTransactor // Write-only binding to the contract
	EnsRegistryWithFallbackFilterer   // Log filterer for contract events
}

// EnsRegistryWithFallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnsRegistryWithFallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistryWithFallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnsRegistryWithFallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistryWithFallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnsRegistryWithFallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistryWithFallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnsRegistryWithFallbackSession struct {
	Contract     *EnsRegistryWithFallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnsRegistryWithFallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnsRegistryWithFallbackCallerSession struct {
	Contract *EnsRegistryWithFallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// EnsRegistryWithFallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnsRegistryWithFallbackTransactorSession struct {
	Contract     *EnsRegistryWithFallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// EnsRegistryWithFallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnsRegistryWithFallbackRaw struct {
	Contract *EnsRegistryWithFallback // Generic contract binding to access the raw methods on
}

// EnsRegistryWithFallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnsRegistryWithFallbackCallerRaw struct {
	Contract *EnsRegistryWithFallbackCaller // Generic read-only contract binding to access the raw methods on
}

// EnsRegistryWithFallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnsRegistryWithFallbackTransactorRaw struct {
	Contract *EnsRegistryWithFallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnsRegistryWithFallback creates a new instance of EnsRegistryWithFallback, bound to a specific deployed contract.
func NewEnsRegistryWithFallback(address common.Address, backend bind.ContractBackend) (*EnsRegistryWithFallback, error) {
	contract, err := bindEnsRegistryWithFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallback{EnsRegistryWithFallbackCaller: EnsRegistryWithFallbackCaller{contract: contract}, EnsRegistryWithFallbackTransactor: EnsRegistryWithFallbackTransactor{contract: contract}, EnsRegistryWithFallbackFilterer: EnsRegistryWithFallbackFilterer{contract: contract}}, nil
}

// NewEnsRegistryWithFallbackCaller creates a new read-only instance of EnsRegistryWithFallback, bound to a specific deployed contract.
func NewEnsRegistryWithFallbackCaller(address common.Address, caller bind.ContractCaller) (*EnsRegistryWithFallbackCaller, error) {
	contract, err := bindEnsRegistryWithFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackCaller{contract: contract}, nil
}

// NewEnsRegistryWithFallbackTransactor creates a new write-only instance of EnsRegistryWithFallback, bound to a specific deployed contract.
func NewEnsRegistryWithFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*EnsRegistryWithFallbackTransactor, error) {
	contract, err := bindEnsRegistryWithFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackTransactor{contract: contract}, nil
}

// NewEnsRegistryWithFallbackFilterer creates a new log filterer instance of EnsRegistryWithFallback, bound to a specific deployed contract.
func NewEnsRegistryWithFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*EnsRegistryWithFallbackFilterer, error) {
	contract, err := bindEnsRegistryWithFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackFilterer{contract: contract}, nil
}

// bindEnsRegistryWithFallback binds a generic wrapper to an already deployed contract.
func bindEnsRegistryWithFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnsRegistryWithFallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnsRegistryWithFallback.Contract.EnsRegistryWithFallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.EnsRegistryWithFallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.EnsRegistryWithFallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnsRegistryWithFallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.contract.Transact(opts, method, params...)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _EnsRegistryWithFallback.contract.Call(opts, &out, "isApprovedForAll", _owner, operator)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) IsApprovedForAll(_owner common.Address, operator common.Address) (bool, error) {
	return _EnsRegistryWithFallback.Contract.IsApprovedForAll(&_EnsRegistryWithFallback.CallOpts, _owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address operator) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerSession) IsApprovedForAll(_owner common.Address, operator common.Address) (bool, error) {
	return _EnsRegistryWithFallback.Contract.IsApprovedForAll(&_EnsRegistryWithFallback.CallOpts, _owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EnsRegistryWithFallback.contract.Call(opts, &out, "owner", node)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) Owner(node [32]byte) (common.Address, error) {
	return _EnsRegistryWithFallback.Contract.Owner(&_EnsRegistryWithFallback.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _EnsRegistryWithFallback.Contract.Owner(&_EnsRegistryWithFallback.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCaller) RecordExists(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _EnsRegistryWithFallback.contract.Call(opts, &out, "recordExists", node)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) RecordExists(node [32]byte) (bool, error) {
	return _EnsRegistryWithFallback.Contract.RecordExists(&_EnsRegistryWithFallback.CallOpts, node)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 node) view returns(bool)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerSession) RecordExists(node [32]byte) (bool, error) {
	return _EnsRegistryWithFallback.Contract.RecordExists(&_EnsRegistryWithFallback.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EnsRegistryWithFallback.contract.Call(opts, &out, "resolver", node)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) Resolver(node [32]byte) (common.Address, error) {
	return _EnsRegistryWithFallback.Contract.Resolver(&_EnsRegistryWithFallback.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) view returns(address)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _EnsRegistryWithFallback.Contract.Resolver(&_EnsRegistryWithFallback.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _EnsRegistryWithFallback.contract.Call(opts, &out, "ttl", node)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) Ttl(node [32]byte) (uint64, error) {
	return _EnsRegistryWithFallback.Contract.Ttl(&_EnsRegistryWithFallback.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) view returns(uint64)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _EnsRegistryWithFallback.Contract.Ttl(&_EnsRegistryWithFallback.CallOpts, node)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetApprovalForAll(&_EnsRegistryWithFallback.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetApprovalForAll(&_EnsRegistryWithFallback.TransactOpts, operator, approved)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetOwner(&_EnsRegistryWithFallback.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetOwner(&_EnsRegistryWithFallback.TransactOpts, node, owner)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetRecord(&_EnsRegistryWithFallback.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetRecord(&_EnsRegistryWithFallback.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetResolver(&_EnsRegistryWithFallback.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetResolver(&_EnsRegistryWithFallback.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetSubnodeOwner(&_EnsRegistryWithFallback.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns(bytes32)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetSubnodeOwner(&_EnsRegistryWithFallback.TransactOpts, node, label, owner)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetSubnodeRecord(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setSubnodeRecord", node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetSubnodeRecord(&_EnsRegistryWithFallback.TransactOpts, node, label, owner, resolver, ttl)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x5ef2c7f0.
//
// Solidity: function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetSubnodeRecord(node [32]byte, label [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetSubnodeRecord(&_EnsRegistryWithFallback.TransactOpts, node, label, owner, resolver, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetTTL(&_EnsRegistryWithFallback.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistryWithFallback.Contract.SetTTL(&_EnsRegistryWithFallback.TransactOpts, node, ttl)
}

// EnsRegistryWithFallbackApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackApprovalForAllIterator struct {
	Event *EnsRegistryWithFallbackApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EnsRegistryWithFallbackApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryWithFallbackApprovalForAll)
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
		it.Event = new(EnsRegistryWithFallbackApprovalForAll)
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
func (it *EnsRegistryWithFallbackApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryWithFallbackApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryWithFallbackApprovalForAll represents a ApprovalForAll event raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*EnsRegistryWithFallbackApprovalForAllIterator, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackApprovalForAllIterator{contract: _EnsRegistryWithFallback.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EnsRegistryWithFallbackApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryWithFallbackApprovalForAll)
				if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) ParseApprovalForAll(log types.Log) (*EnsRegistryWithFallbackApprovalForAll, error) {
	event := new(EnsRegistryWithFallbackApprovalForAll)
	if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryWithFallbackNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewOwnerIterator struct {
	Event *EnsRegistryWithFallbackNewOwner // Event containing the contract specifics and raw log

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
func (it *EnsRegistryWithFallbackNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryWithFallbackNewOwner)
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
		it.Event = new(EnsRegistryWithFallbackNewOwner)
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
func (it *EnsRegistryWithFallbackNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryWithFallbackNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryWithFallbackNewOwner represents a NewOwner event raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*EnsRegistryWithFallbackNewOwnerIterator, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackNewOwnerIterator{contract: _EnsRegistryWithFallback.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *EnsRegistryWithFallbackNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryWithFallbackNewOwner)
				if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) ParseNewOwner(log types.Log) (*EnsRegistryWithFallbackNewOwner, error) {
	event := new(EnsRegistryWithFallbackNewOwner)
	if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryWithFallbackNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewResolverIterator struct {
	Event *EnsRegistryWithFallbackNewResolver // Event containing the contract specifics and raw log

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
func (it *EnsRegistryWithFallbackNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryWithFallbackNewResolver)
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
		it.Event = new(EnsRegistryWithFallbackNewResolver)
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
func (it *EnsRegistryWithFallbackNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryWithFallbackNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryWithFallbackNewResolver represents a NewResolver event raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryWithFallbackNewResolverIterator, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackNewResolverIterator{contract: _EnsRegistryWithFallback.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *EnsRegistryWithFallbackNewResolver, node [][32]byte) (event.Subscription, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryWithFallbackNewResolver)
				if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) ParseNewResolver(log types.Log) (*EnsRegistryWithFallbackNewResolver, error) {
	event := new(EnsRegistryWithFallbackNewResolver)
	if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryWithFallbackNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewTTLIterator struct {
	Event *EnsRegistryWithFallbackNewTTL // Event containing the contract specifics and raw log

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
func (it *EnsRegistryWithFallbackNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryWithFallbackNewTTL)
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
		it.Event = new(EnsRegistryWithFallbackNewTTL)
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
func (it *EnsRegistryWithFallbackNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryWithFallbackNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryWithFallbackNewTTL represents a NewTTL event raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryWithFallbackNewTTLIterator, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackNewTTLIterator{contract: _EnsRegistryWithFallback.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *EnsRegistryWithFallbackNewTTL, node [][32]byte) (event.Subscription, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryWithFallbackNewTTL)
				if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) ParseNewTTL(log types.Log) (*EnsRegistryWithFallbackNewTTL, error) {
	event := new(EnsRegistryWithFallbackNewTTL)
	if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryWithFallbackTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackTransferIterator struct {
	Event *EnsRegistryWithFallbackTransfer // Event containing the contract specifics and raw log

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
func (it *EnsRegistryWithFallbackTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryWithFallbackTransfer)
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
		it.Event = new(EnsRegistryWithFallbackTransfer)
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
func (it *EnsRegistryWithFallbackTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryWithFallbackTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryWithFallbackTransfer represents a Transfer event raised by the EnsRegistryWithFallback contract.
type EnsRegistryWithFallbackTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryWithFallbackTransferIterator, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryWithFallbackTransferIterator{contract: _EnsRegistryWithFallback.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EnsRegistryWithFallbackTransfer, node [][32]byte) (event.Subscription, error) {
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistryWithFallback.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryWithFallbackTransfer)
				if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistryWithFallback *EnsRegistryWithFallbackFilterer) ParseTransfer(log types.Log) (*EnsRegistryWithFallbackTransfer, error) {
	event := new(EnsRegistryWithFallbackTransfer)
	if err := _EnsRegistryWithFallback.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
