// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"transfer_image_url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"transfer_note\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"name\":\"MilestoneReleaseStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"transfer_receipt_url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof_media_url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"name\":\"VerifiedProofStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"milestoneReleases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"transfer_image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transfer_note\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"transfer_image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transfer_note\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"name\":\"storeFundRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"transfer_receipt_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof_media_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"name\":\"storeVerifiedProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiedProofs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"project_id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"milestone_id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"transfer_receipt_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof_media_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"created_at\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// MilestoneReleases is a free data retrieval call binding the contract method 0xf8b7c122.
//
// Solidity: function milestoneReleases(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractCaller) MilestoneReleases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id               *big.Int
	ProjectId        uint64
	MilestoneId      uint64
	TransferImageUrl string
	TransferNote     string
	CreatedAt        string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "milestoneReleases", arg0)

	outstruct := new(struct {
		Id               *big.Int
		ProjectId        uint64
		MilestoneId      uint64
		TransferImageUrl string
		TransferNote     string
		CreatedAt        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProjectId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.MilestoneId = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.TransferImageUrl = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.TransferNote = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// MilestoneReleases is a free data retrieval call binding the contract method 0xf8b7c122.
//
// Solidity: function milestoneReleases(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractSession) MilestoneReleases(arg0 *big.Int) (struct {
	Id               *big.Int
	ProjectId        uint64
	MilestoneId      uint64
	TransferImageUrl string
	TransferNote     string
	CreatedAt        string
}, error) {
	return _Contract.Contract.MilestoneReleases(&_Contract.CallOpts, arg0)
}

// MilestoneReleases is a free data retrieval call binding the contract method 0xf8b7c122.
//
// Solidity: function milestoneReleases(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractCallerSession) MilestoneReleases(arg0 *big.Int) (struct {
	Id               *big.Int
	ProjectId        uint64
	MilestoneId      uint64
	TransferImageUrl string
	TransferNote     string
	CreatedAt        string
}, error) {
	return _Contract.Contract.MilestoneReleases(&_Contract.CallOpts, arg0)
}

// VerifiedProofs is a free data retrieval call binding the contract method 0x47d25c25.
//
// Solidity: function verifiedProofs(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractCaller) VerifiedProofs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                 *big.Int
	ProjectId          uint64
	MilestoneId        uint64
	TransferReceiptUrl string
	ProofMediaUrl      string
	CreatedAt          string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifiedProofs", arg0)

	outstruct := new(struct {
		Id                 *big.Int
		ProjectId          uint64
		MilestoneId        uint64
		TransferReceiptUrl string
		ProofMediaUrl      string
		CreatedAt          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProjectId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.MilestoneId = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.TransferReceiptUrl = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ProofMediaUrl = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// VerifiedProofs is a free data retrieval call binding the contract method 0x47d25c25.
//
// Solidity: function verifiedProofs(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractSession) VerifiedProofs(arg0 *big.Int) (struct {
	Id                 *big.Int
	ProjectId          uint64
	MilestoneId        uint64
	TransferReceiptUrl string
	ProofMediaUrl      string
	CreatedAt          string
}, error) {
	return _Contract.Contract.VerifiedProofs(&_Contract.CallOpts, arg0)
}

// VerifiedProofs is a free data retrieval call binding the contract method 0x47d25c25.
//
// Solidity: function verifiedProofs(uint256 ) view returns(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractCallerSession) VerifiedProofs(arg0 *big.Int) (struct {
	Id                 *big.Int
	ProjectId          uint64
	MilestoneId        uint64
	TransferReceiptUrl string
	ProofMediaUrl      string
	CreatedAt          string
}, error) {
	return _Contract.Contract.VerifiedProofs(&_Contract.CallOpts, arg0)
}

// StoreFundRelease is a paid mutator transaction binding the contract method 0x3fb66a2c.
//
// Solidity: function storeFundRelease(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at) returns()
func (_Contract *ContractTransactor) StoreFundRelease(opts *bind.TransactOpts, id *big.Int, project_id uint64, milestone_id uint64, transfer_image_url string, transfer_note string, created_at string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "storeFundRelease", id, project_id, milestone_id, transfer_image_url, transfer_note, created_at)
}

// StoreFundRelease is a paid mutator transaction binding the contract method 0x3fb66a2c.
//
// Solidity: function storeFundRelease(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at) returns()
func (_Contract *ContractSession) StoreFundRelease(id *big.Int, project_id uint64, milestone_id uint64, transfer_image_url string, transfer_note string, created_at string) (*types.Transaction, error) {
	return _Contract.Contract.StoreFundRelease(&_Contract.TransactOpts, id, project_id, milestone_id, transfer_image_url, transfer_note, created_at)
}

// StoreFundRelease is a paid mutator transaction binding the contract method 0x3fb66a2c.
//
// Solidity: function storeFundRelease(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at) returns()
func (_Contract *ContractTransactorSession) StoreFundRelease(id *big.Int, project_id uint64, milestone_id uint64, transfer_image_url string, transfer_note string, created_at string) (*types.Transaction, error) {
	return _Contract.Contract.StoreFundRelease(&_Contract.TransactOpts, id, project_id, milestone_id, transfer_image_url, transfer_note, created_at)
}

// StoreVerifiedProof is a paid mutator transaction binding the contract method 0x8981d422.
//
// Solidity: function storeVerifiedProof(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at) returns()
func (_Contract *ContractTransactor) StoreVerifiedProof(opts *bind.TransactOpts, id *big.Int, project_id uint64, milestone_id uint64, transfer_receipt_url string, proof_media_url string, created_at string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "storeVerifiedProof", id, project_id, milestone_id, transfer_receipt_url, proof_media_url, created_at)
}

// StoreVerifiedProof is a paid mutator transaction binding the contract method 0x8981d422.
//
// Solidity: function storeVerifiedProof(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at) returns()
func (_Contract *ContractSession) StoreVerifiedProof(id *big.Int, project_id uint64, milestone_id uint64, transfer_receipt_url string, proof_media_url string, created_at string) (*types.Transaction, error) {
	return _Contract.Contract.StoreVerifiedProof(&_Contract.TransactOpts, id, project_id, milestone_id, transfer_receipt_url, proof_media_url, created_at)
}

// StoreVerifiedProof is a paid mutator transaction binding the contract method 0x8981d422.
//
// Solidity: function storeVerifiedProof(uint256 id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at) returns()
func (_Contract *ContractTransactorSession) StoreVerifiedProof(id *big.Int, project_id uint64, milestone_id uint64, transfer_receipt_url string, proof_media_url string, created_at string) (*types.Transaction, error) {
	return _Contract.Contract.StoreVerifiedProof(&_Contract.TransactOpts, id, project_id, milestone_id, transfer_receipt_url, proof_media_url, created_at)
}

// ContractMilestoneReleaseStoredIterator is returned from FilterMilestoneReleaseStored and is used to iterate over the raw logs and unpacked data for MilestoneReleaseStored events raised by the Contract contract.
type ContractMilestoneReleaseStoredIterator struct {
	Event *ContractMilestoneReleaseStored // Event containing the contract specifics and raw log

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
func (it *ContractMilestoneReleaseStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMilestoneReleaseStored)
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
		it.Event = new(ContractMilestoneReleaseStored)
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
func (it *ContractMilestoneReleaseStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMilestoneReleaseStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMilestoneReleaseStored represents a MilestoneReleaseStored event raised by the Contract contract.
type ContractMilestoneReleaseStored struct {
	Id               *big.Int
	ProjectId        uint64
	MilestoneId      uint64
	TransferImageUrl string
	TransferNote     string
	CreatedAt        string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMilestoneReleaseStored is a free log retrieval operation binding the contract event 0x5417b1846464590ed2346103bd0690e69d2712bcbcfabe3409e64cc88af1aff1.
//
// Solidity: event MilestoneReleaseStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractFilterer) FilterMilestoneReleaseStored(opts *bind.FilterOpts, id []*big.Int) (*ContractMilestoneReleaseStoredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MilestoneReleaseStored", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractMilestoneReleaseStoredIterator{contract: _Contract.contract, event: "MilestoneReleaseStored", logs: logs, sub: sub}, nil
}

// WatchMilestoneReleaseStored is a free log subscription operation binding the contract event 0x5417b1846464590ed2346103bd0690e69d2712bcbcfabe3409e64cc88af1aff1.
//
// Solidity: event MilestoneReleaseStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractFilterer) WatchMilestoneReleaseStored(opts *bind.WatchOpts, sink chan<- *ContractMilestoneReleaseStored, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MilestoneReleaseStored", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMilestoneReleaseStored)
				if err := _Contract.contract.UnpackLog(event, "MilestoneReleaseStored", log); err != nil {
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

// ParseMilestoneReleaseStored is a log parse operation binding the contract event 0x5417b1846464590ed2346103bd0690e69d2712bcbcfabe3409e64cc88af1aff1.
//
// Solidity: event MilestoneReleaseStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_image_url, string transfer_note, string created_at)
func (_Contract *ContractFilterer) ParseMilestoneReleaseStored(log types.Log) (*ContractMilestoneReleaseStored, error) {
	event := new(ContractMilestoneReleaseStored)
	if err := _Contract.contract.UnpackLog(event, "MilestoneReleaseStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractVerifiedProofStoredIterator is returned from FilterVerifiedProofStored and is used to iterate over the raw logs and unpacked data for VerifiedProofStored events raised by the Contract contract.
type ContractVerifiedProofStoredIterator struct {
	Event *ContractVerifiedProofStored // Event containing the contract specifics and raw log

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
func (it *ContractVerifiedProofStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractVerifiedProofStored)
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
		it.Event = new(ContractVerifiedProofStored)
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
func (it *ContractVerifiedProofStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractVerifiedProofStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractVerifiedProofStored represents a VerifiedProofStored event raised by the Contract contract.
type ContractVerifiedProofStored struct {
	Id                 *big.Int
	ProjectId          uint64
	MilestoneId        uint64
	TransferReceiptUrl string
	ProofMediaUrl      string
	CreatedAt          string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerifiedProofStored is a free log retrieval operation binding the contract event 0x0f1681c253fafd2ff469b1085b3d475af453854c8e1374d643c356a356368916.
//
// Solidity: event VerifiedProofStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractFilterer) FilterVerifiedProofStored(opts *bind.FilterOpts, id []*big.Int) (*ContractVerifiedProofStoredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "VerifiedProofStored", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractVerifiedProofStoredIterator{contract: _Contract.contract, event: "VerifiedProofStored", logs: logs, sub: sub}, nil
}

// WatchVerifiedProofStored is a free log subscription operation binding the contract event 0x0f1681c253fafd2ff469b1085b3d475af453854c8e1374d643c356a356368916.
//
// Solidity: event VerifiedProofStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractFilterer) WatchVerifiedProofStored(opts *bind.WatchOpts, sink chan<- *ContractVerifiedProofStored, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "VerifiedProofStored", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractVerifiedProofStored)
				if err := _Contract.contract.UnpackLog(event, "VerifiedProofStored", log); err != nil {
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

// ParseVerifiedProofStored is a log parse operation binding the contract event 0x0f1681c253fafd2ff469b1085b3d475af453854c8e1374d643c356a356368916.
//
// Solidity: event VerifiedProofStored(uint256 indexed id, uint64 project_id, uint64 milestone_id, string transfer_receipt_url, string proof_media_url, string created_at)
func (_Contract *ContractFilterer) ParseVerifiedProofStored(log types.Log) (*ContractVerifiedProofStored, error) {
	event := new(ContractVerifiedProofStored)
	if err := _Contract.contract.UnpackLog(event, "VerifiedProofStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
