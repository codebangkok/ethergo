// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bondcoin

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

// BondcoinABI is the input ABI used to generate the binding from.
const BondcoinABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BondcoinBin is the compiled bytecode used for deploying new contracts.
var BondcoinBin = "0x608060405234801561001057600080fd5b5061065e806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806340c10f191161005b57806340c10f191461014457806370a082311461019257806395d89b41146101ea578063a9059cbb1461026d5761007d565b806306fdde031461008257806318160ddd14610105578063313ce56714610123575b600080fd5b61008a6102d1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100ca5780820151818401526020810190506100af565b50505050905090810190601f1680156100f75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010d61030e565b6040518082815260200191505060405180910390f35b61012b610317565b604051808260ff16815260200191505060405180910390f35b6101906004803603604081101561015a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061031c565b005b6101d4600480360360208110156101a857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103e2565b6040518082815260200191505060405180910390f35b6101f261042b565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610232578082015181840152602081019050610217565b50505050905090810190601f16801561025f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102b96004803603604081101561028357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610468565b60405180821515815260200191505060405180910390f35b60606040518060400160405280600981526020017f426f6e6420436f696e0000000000000000000000000000000000000000000000815250905090565b60008054905090565b600090565b80600080828254019250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606040518060400160405280600381526020017f424f4e0000000000000000000000000000000000000000000000000000000000815250905090565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561051f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4e6f7420656e6f75676820636f696e000000000000000000000000000000000081525060200191505060405180910390fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509291505056fea264697066735822122067ed8e7578b90e97fcb792d51cb3c504297a699daae252a34486e7ae9c5d147164736f6c634300060c0033"

// DeployBondcoin deploys a new Ethereum contract, binding an instance of Bondcoin to it.
func DeployBondcoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bondcoin, error) {
	parsed, err := abi.JSON(strings.NewReader(BondcoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BondcoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bondcoin{BondcoinCaller: BondcoinCaller{contract: contract}, BondcoinTransactor: BondcoinTransactor{contract: contract}, BondcoinFilterer: BondcoinFilterer{contract: contract}}, nil
}

// Bondcoin is an auto generated Go binding around an Ethereum contract.
type Bondcoin struct {
	BondcoinCaller     // Read-only binding to the contract
	BondcoinTransactor // Write-only binding to the contract
	BondcoinFilterer   // Log filterer for contract events
}

// BondcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondcoinSession struct {
	Contract     *Bondcoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondcoinCallerSession struct {
	Contract *BondcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BondcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondcoinTransactorSession struct {
	Contract     *BondcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BondcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondcoinRaw struct {
	Contract *Bondcoin // Generic contract binding to access the raw methods on
}

// BondcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondcoinCallerRaw struct {
	Contract *BondcoinCaller // Generic read-only contract binding to access the raw methods on
}

// BondcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondcoinTransactorRaw struct {
	Contract *BondcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondcoin creates a new instance of Bondcoin, bound to a specific deployed contract.
func NewBondcoin(address common.Address, backend bind.ContractBackend) (*Bondcoin, error) {
	contract, err := bindBondcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bondcoin{BondcoinCaller: BondcoinCaller{contract: contract}, BondcoinTransactor: BondcoinTransactor{contract: contract}, BondcoinFilterer: BondcoinFilterer{contract: contract}}, nil
}

// NewBondcoinCaller creates a new read-only instance of Bondcoin, bound to a specific deployed contract.
func NewBondcoinCaller(address common.Address, caller bind.ContractCaller) (*BondcoinCaller, error) {
	contract, err := bindBondcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondcoinCaller{contract: contract}, nil
}

// NewBondcoinTransactor creates a new write-only instance of Bondcoin, bound to a specific deployed contract.
func NewBondcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*BondcoinTransactor, error) {
	contract, err := bindBondcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondcoinTransactor{contract: contract}, nil
}

// NewBondcoinFilterer creates a new log filterer instance of Bondcoin, bound to a specific deployed contract.
func NewBondcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*BondcoinFilterer, error) {
	contract, err := bindBondcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondcoinFilterer{contract: contract}, nil
}

// bindBondcoin binds a generic wrapper to an already deployed contract.
func bindBondcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondcoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bondcoin *BondcoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bondcoin.Contract.BondcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bondcoin *BondcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bondcoin.Contract.BondcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bondcoin *BondcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bondcoin.Contract.BondcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bondcoin *BondcoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bondcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bondcoin *BondcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bondcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bondcoin *BondcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bondcoin.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Bondcoin *BondcoinCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bondcoin.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Bondcoin *BondcoinSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Bondcoin.Contract.BalanceOf(&_Bondcoin.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Bondcoin *BondcoinCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Bondcoin.Contract.BalanceOf(&_Bondcoin.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Bondcoin *BondcoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Bondcoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Bondcoin *BondcoinSession) Decimals() (uint8, error) {
	return _Bondcoin.Contract.Decimals(&_Bondcoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Bondcoin *BondcoinCallerSession) Decimals() (uint8, error) {
	return _Bondcoin.Contract.Decimals(&_Bondcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Bondcoin *BondcoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bondcoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Bondcoin *BondcoinSession) Name() (string, error) {
	return _Bondcoin.Contract.Name(&_Bondcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Bondcoin *BondcoinCallerSession) Name() (string, error) {
	return _Bondcoin.Contract.Name(&_Bondcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Bondcoin *BondcoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bondcoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Bondcoin *BondcoinSession) Symbol() (string, error) {
	return _Bondcoin.Contract.Symbol(&_Bondcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Bondcoin *BondcoinCallerSession) Symbol() (string, error) {
	return _Bondcoin.Contract.Symbol(&_Bondcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bondcoin *BondcoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bondcoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bondcoin *BondcoinSession) TotalSupply() (*big.Int, error) {
	return _Bondcoin.Contract.TotalSupply(&_Bondcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bondcoin *BondcoinCallerSession) TotalSupply() (*big.Int, error) {
	return _Bondcoin.Contract.TotalSupply(&_Bondcoin.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _owner, uint256 _value) returns()
func (_Bondcoin *BondcoinTransactor) Mint(opts *bind.TransactOpts, _owner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.contract.Transact(opts, "mint", _owner, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _owner, uint256 _value) returns()
func (_Bondcoin *BondcoinSession) Mint(_owner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.Contract.Mint(&_Bondcoin.TransactOpts, _owner, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _owner, uint256 _value) returns()
func (_Bondcoin *BondcoinTransactorSession) Mint(_owner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.Contract.Mint(&_Bondcoin.TransactOpts, _owner, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bondcoin *BondcoinTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bondcoin *BondcoinSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.Contract.Transfer(&_Bondcoin.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Bondcoin *BondcoinTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Bondcoin.Contract.Transfer(&_Bondcoin.TransactOpts, _to, _value)
}

// BondcoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bondcoin contract.
type BondcoinTransferIterator struct {
	Event *BondcoinTransfer // Event containing the contract specifics and raw log

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
func (it *BondcoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondcoinTransfer)
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
		it.Event = new(BondcoinTransfer)
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
func (it *BondcoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondcoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondcoinTransfer represents a Transfer event raised by the Bondcoin contract.
type BondcoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Bondcoin *BondcoinFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*BondcoinTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Bondcoin.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BondcoinTransferIterator{contract: _Bondcoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Bondcoin *BondcoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BondcoinTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Bondcoin.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondcoinTransfer)
				if err := _Bondcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Bondcoin *BondcoinFilterer) ParseTransfer(log types.Log) (*BondcoinTransfer, error) {
	event := new(BondcoinTransfer)
	if err := _Bondcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
