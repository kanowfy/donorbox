package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainTransactor struct {
	EthClient *ethclient.Client
	Contract  *Contract
	AuthData  *bind.TransactOpts
}
