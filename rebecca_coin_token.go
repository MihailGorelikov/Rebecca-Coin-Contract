package rebecca_coin_contract

import (
	"context"
	"fmt"
	"math/big"

	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type (
	// ERC20Token is the interface for ERC20 token
	ERC20Token interface {
		// Name returns the name of the token.
		// function name() external view returns (string memory);
		Name(ctx context.Context) (string, error)

		// Symbol returns the symbol of the token.
		// function symbol() external view returns (string memory);
		Symbol(ctx context.Context) (string, error)

		// Decimals returns the number of decimals the token uses.
		// function decimals() external view returns(uint8);
		Decimals(ctx context.Context) (uint8, error)

		// TotalSupply returns the total token supply.
		// function totalSupply() external view returns (uint256);
		TotalSupply(ctx context.Context) (*big.Int, error)

		// BalanceOf returns the account balance of another account with address _owner.
		// function balanceOf(address _owner) external view returns (uint256);
		BalanceOf(ctx context.Context, address string) (*big.Int, error)

		// Transfer transfers _value amount of tokens to address _to, and MUST fire the Transfer event.
		// function transfer(address _to, uint256 _value) external returns(bool);
		Transfer(ctx context.Context, to string, amount *big.Int) (bool, error)

		// TransferFrom transfers _value amount of tokens from address _from to address _to, and MUST fire the Transfer event.
		// function transferFrom(address _from, address _to, uint256 _value) external returns (bool success);
		TransferFrom(ctx context.Context, from, to string, amount *big.Int) (bool, error)

		// Approve allows _spender to withdraw from your account multiple times, up to the _value amount.
		// function approve(address _spender, uint256 _value) external returns (bool success);
		Approve(ctx context.Context, spender string, amount *big.Int) (bool, error)

		// Allowance returns the amount which _spender is still allowed to withdraw from _owner.
		// function allowance(address _owner, address _spender) external view returns (uint256 remaining);
		Allowance(ctx context.Context, owner, spender string) (*big.Int, error)
	}

	// RebeccaCoinToken is the implementaiont of the ERC20 token
	RebeccaCoinToken struct {
		client                *ethclient.Client
		contractAddress       common.Address
		contractABIJSONSource string
	}
)

// NewRebeccaCoinToken creates a new RebeccaCoinToken instance.
func NewRebeccaCoinToken(client *ethclient.Client, contractAddress string) *RebeccaCoinToken {
	return &RebeccaCoinToken{
		client:          client,
		contractAddress: common.HexToAddress(contractAddress),
		contractABIJSONSource: `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "initialAuthority",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "authority",
				"type": "address"
			}
		],
		"name": "AccessManagedInvalidAuthority",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "caller",
				"type": "address"
			},
			{
				"internalType": "uint32",
				"name": "delay",
				"type": "uint32"
			}
		],
		"name": "AccessManagedRequiredDelay",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "caller",
				"type": "address"
			}
		],
		"name": "AccessManagedUnauthorized",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "ECDSAInvalidSignature",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "length",
				"type": "uint256"
			}
		],
		"name": "ECDSAInvalidSignatureLength",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "ECDSAInvalidSignatureS",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "allowance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientAllowance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "balance",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "needed",
				"type": "uint256"
			}
		],
		"name": "ERC20InsufficientBalance",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "approver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidApprover",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			}
		],
		"name": "ERC20InvalidReceiver",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSender",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "ERC20InvalidSpender",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "ERC2612ExpiredSignature",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "signer",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "ERC2612InvalidSigner",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "currentNonce",
				"type": "uint256"
			}
		],
		"name": "InvalidAccountNonce",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "InvalidShortString",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "str",
				"type": "string"
			}
		],
		"name": "StringTooLong",
		"type": "error"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "authority",
				"type": "address"
			}
		],
		"name": "AuthorityUpdated",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "EIP712DomainChanged",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "DOMAIN_SEPARATOR",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			}
		],
		"name": "allowance",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "approve",
		"stateMutability": "nonpayable",
		"type": "function",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		]
	},
	{
		"inputs": [],
		"name": "authority",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		]
	},
	{
		"inputs": [],
		"name": "decimals",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		]
	},
	{
		"inputs": [],
		"name": "eip712Domain",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "bytes1",
				"name": "fields",
				"type": "bytes1"
			},
			{
				"internalType": "string",
				"name": "name",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "version",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "chainId",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "verifyingContract",
				"type": "address"
			},
			{
				"internalType": "bytes32",
				"name": "salt",
				"type": "bytes32"
			},
			{
				"internalType": "uint256[]",
				"name": "extensions",
				"type": "uint256[]"
			}
		]
	},
	{
		"inputs": [],
		"name": "isConsumingScheduledOp",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "bytes4",
				"name": "",
				"type": "bytes4"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "mint",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "nonces",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "permit",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newAuthority",
				"type": "address"
			}
		],
		"name": "setAuthority",
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		]
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"stateMutability": "view",
		"type": "function",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"stateMutability": "nonpayable",
		"type": "function",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		]
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"stateMutability": "nonpayable",
		"type": "function",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		]
	}
]`,
	}
}

// Allowance returns the amount which _spender is still allowed to withdraw from _owner.
func (token *RebeccaCoinToken) Allowance(ctx context.Context, owner string, spender string) (*big.Int, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	_owner := common.HexToAddress(owner)
	_spender := common.HexToAddress(spender)

	message, err := abi.Pack("allowance", _owner, _spender)
	if err != nil {
		return nil, fmt.Errorf("failed to pack allowance message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var allowance *big.Int
	err = abi.UnpackIntoInterface(&allowance, "allowance", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack allowance: %w", err)
	}

	return allowance, nil
}

// Approve allows _spender to withdraw from your account multiple times, up to the _value amount.
func (token *RebeccaCoinToken) Approve(ctx context.Context, spender string, amount *big.Int) (bool, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return false, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	_spender := common.HexToAddress(spender)

	message, err := abi.Pack("approve", _spender, amount)
	if err != nil {
		return false, fmt.Errorf("failed to pack approve message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return false, fmt.Errorf("failed to call contract: %w", err)
	}

	var success bool
	err = abi.UnpackIntoInterface(&success, "approve", output)
	if err != nil {
		return false, fmt.Errorf("failed to unpack approve: %w", err)
	}

	return success, nil
}

// BalanceOf returns the account balance of another account with address _owner.
func (token *RebeccaCoinToken) BalanceOf(ctx context.Context, address string) (*big.Int, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	_address := common.HexToAddress(address)

	message, err := abi.Pack("balanceOf", _address)
	if err != nil {
		return nil, fmt.Errorf("failed to pack balanceOf message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var balance *big.Int
	err = abi.UnpackIntoInterface(&balance, "balanceOf", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack balanceOf: %w", err)
	}

	return balance, nil
}

// Decimals returns the number of decimals the token uses.
func (token *RebeccaCoinToken) Decimals(ctx context.Context) (uint8, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return 0, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	message, err := abi.Pack("decimals")
	if err != nil {
		return 0, fmt.Errorf("failed to pack decimals message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return 0, fmt.Errorf("failed to call contract: %w", err)
	}

	var decimals uint8
	err = abi.UnpackIntoInterface(&decimals, "decimals", output)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack decimals: %w", err)
	}

	return decimals, nil
}

// Name returns the name of the token.
func (token *RebeccaCoinToken) Name(ctx context.Context) (string, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return "", fmt.Errorf("failed to get contract ABI: %w", err)
	}

	message, err := abi.Pack("name")
	if err != nil {
		return "", fmt.Errorf("failed to pack name message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var name string
	err = abi.UnpackIntoInterface(&name, "name", output)
	if err != nil {
		return "", fmt.Errorf("failed to unpack name: %w", err)
	}

	return name, nil
}

// Symbol returns the symbol of the token.
func (token *RebeccaCoinToken) Symbol(ctx context.Context) (string, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return "", fmt.Errorf("failed to get contract ABI: %w", err)
	}

	message, err := abi.Pack("symbol")
	if err != nil {
		return "", fmt.Errorf("failed to pack symbol message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	var symbol string
	err = abi.UnpackIntoInterface(&symbol, "symbol", output)
	if err != nil {
		return "", fmt.Errorf("failed to unpack symbol: %w", err)
	}

	return symbol, nil
}

// TotalSupply returns the total token supply.
func (token *RebeccaCoinToken) TotalSupply(ctx context.Context) (*big.Int, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	message, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("failed to pack totalSupply message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	var totalSupply *big.Int
	err = abi.UnpackIntoInterface(&totalSupply, "totalSupply", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack totalSupply: %w", err)
	}

	return totalSupply, nil
}

// Transfer transfers _value amount of tokens to address _to, and MUST fire the Transfer event.
func (token *RebeccaCoinToken) Transfer(ctx context.Context, to string, amount *big.Int) (bool, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return false, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	_to := common.HexToAddress(to)

	message, err := abi.Pack("transfer", _to, amount)
	if err != nil {
		return false, fmt.Errorf("failed to pack transfer message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return false, fmt.Errorf("failed to call contract: %w", err)
	}

	var success bool
	err = abi.UnpackIntoInterface(&success, "transfer", output)
	if err != nil {
		return false, fmt.Errorf("failed to unpack transfer: %w", err)
	}

	return success, nil
}

// TransferFrom transfers _value amount of tokens from address _from to address _to, and MUST fire the Transfer event.
func (token *RebeccaCoinToken) TransferFrom(ctx context.Context, from string, to string, amount *big.Int) (bool, error) {
	abi, err := token.getContractABI()
	if err != nil {
		return false, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	_from := common.HexToAddress(from)
	_to := common.HexToAddress(to)

	message, err := abi.Pack("transferFrom", _from, _to, amount)
	if err != nil {
		return false, fmt.Errorf("failed to pack transferFrom message: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: token.contractAddress,
		To:   &token.contractAddress,
		Data: message,
	}

	blockNumber, err := token.client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get block number: %w", err)
	}

	output, err := token.client.CallContract(ctx, callMsg, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return false, fmt.Errorf("failed to call contract: %w", err)
	}

	var success bool
	err = abi.UnpackIntoInterface(&success, "transferFrom", output)
	if err != nil {
		return false, fmt.Errorf("failed to unpack transferFrom: %w", err)
	}

	return success, nil
}

func (token *RebeccaCoinToken) getContractABI() (abi.ABI, error) {
	contractABIReader := strings.NewReader(token.contractABIJSONSource)

	contractABI, err := abi.JSON(contractABIReader)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return contractABI, nil
}
