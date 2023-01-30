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

// KuiperNFTMetaData contains all meta data concerning the KuiperNFT contract.
var KuiperNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoyaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoyaltyRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyFee\",\"type\":\"uint256\"}],\"name\":\"updateRoyaltyFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200201338038062002013833981016040819052620000349162000339565b8484600062000044838262000461565b50600162000053828262000461565b505050620000706200006a6200012460201b60201c565b62000128565b612710821115620000da5760405162461bcd60e51b815260206004820152602960248201527f726f79616c7479206665652063616e2774206265206d6f7265207468616e20316044820152680c081c195c98d95b9d60ba1b60648201526084015b60405180910390fd5b6001600160a01b038116620000ee57600080fd5b6009829055600a80546001600160a01b0319166001600160a01b03831617905562000119836200017a565b50505050506200052d565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b62000184620001f9565b6001600160a01b038116620001eb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000d1565b620001f68162000128565b50565b6007546001600160a01b03163314620002555760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000d1565b565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200027f57600080fd5b81516001600160401b03808211156200029c576200029c62000257565b604051601f8301601f19908116603f01168101908282118183101715620002c757620002c762000257565b81604052838152602092508683858801011115620002e457600080fd5b600091505b83821015620003085785820183015181830184015290820190620002e9565b600093810190920192909252949350505050565b80516001600160a01b03811681146200033457600080fd5b919050565b600080600080600060a086880312156200035257600080fd5b85516001600160401b03808211156200036a57600080fd5b6200037889838a016200026d565b965060208801519150808211156200038f57600080fd5b506200039e888289016200026d565b945050620003af604087016200031c565b925060608601519150620003c6608087016200031c565b90509295509295909350565b600181811c90821680620003e757607f821691505b6020821081036200040857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200045c57600081815260208120601f850160051c81016020861015620004375750805b601f850160051c820191505b81811015620004585782815560010162000443565b5050505b505050565b81516001600160401b038111156200047d576200047d62000257565b62000495816200048e8454620003d2565b846200040e565b602080601f831160018114620004cd5760008415620004b45750858301515b600019600386901b1c1916600185901b17855562000458565b600085815260208120601f198616915b82811015620004fe57888601518255948401946001909101908401620004dd565b50858210156200051d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611ad6806200053d6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063715018a6116100b8578063a22cb4651161007c578063a22cb46514610273578063b88d4fde14610286578063c87b56dd14610299578063d204c45e146102ac578063e985e9c5146102bf578063f2fde38b146102d257600080fd5b8063715018a614610239578063820bdcdc146102415780638da5cb5b1461024957806395d89b411461025a57806395edc18c1461026257600080fd5b806342842e0e116100ff57806342842e0e146101cc57806342966c68146101df5780634e83be47146101f25780636352211e1461020557806370a082311461021857600080fd5b806301ffc9a71461013c57806306fdde0314610164578063081812fc14610179578063095ea7b3146101a457806323b872dd146101b9575b600080fd5b61014f61014a3660046114a0565b6102e5565b60405190151581526020015b60405180910390f35b61016c610337565b60405161015b919061150d565b61018c610187366004611520565b6103c9565b6040516001600160a01b03909116815260200161015b565b6101b76101b2366004611555565b6103f0565b005b6101b76101c736600461157f565b61050a565b6101b76101da36600461157f565b61053c565b6101b76101ed366004611520565b610557565b6101b7610200366004611520565b610588565b61018c610213366004611520565b6105e7565b61022b6102263660046115bb565b610647565b60405190815260200161015b565b6101b76106cd565b60095461022b565b6007546001600160a01b031661018c565b61016c6106e1565b600a546001600160a01b031661018c565b6101b76102813660046115d6565b6106f0565b6101b761029436600461169e565b6106ff565b61016c6102a7366004611520565b610737565b6101b76102ba36600461171a565b610742565b61014f6102cd36600461177c565b610779565b6101b76102e03660046115bb565b6107a7565b60006001600160e01b031982166380ac58cd60e01b148061031657506001600160e01b03198216635b5e139f60e01b145b8061033157506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060008054610346906117af565b80601f0160208091040260200160405190810160405280929190818152602001828054610372906117af565b80156103bf5780601f10610394576101008083540402835291602001916103bf565b820191906000526020600020905b8154815290600101906020018083116103a257829003601f168201915b5050505050905090565b60006103d48261081d565b506000908152600460205260409020546001600160a01b031690565b60006103fb826105e7565b9050806001600160a01b0316836001600160a01b03160361046d5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061048957506104898133610779565b6104fb5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610464565b610505838361087c565b505050565b610515335b826108ea565b6105315760405162461bcd60e51b8152600401610464906117e9565b610505838383610949565b610505838383604051806020016040528060008152506106ff565b6105603361050f565b61057c5760405162461bcd60e51b8152600401610464906117e9565b61058581610aba565b50565b610590610ac3565b6127108111156105e25760405162461bcd60e51b815260206004820152601a60248201527f63616e2774206d6f7265207468616e2031302070657263656e740000000000006044820152606401610464565b600955565b6000818152600260205260408120546001600160a01b0316806103315760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610464565b60006001600160a01b0382166106b15760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610464565b506001600160a01b031660009081526003602052604090205490565b6106d5610ac3565b6106df6000610b1d565b565b606060018054610346906117af565b6106fb338383610b6f565b5050565b61070933836108ea565b6107255760405162461bcd60e51b8152600401610464906117e9565b61073184848484610c3d565b50505050565b606061033182610c70565b61074a610ac3565b600061075560085490565b9050610765600880546001019055565b61076f8382610d78565b6105058183610d92565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6107af610ac3565b6001600160a01b0381166108145760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610464565b61058581610b1d565b6000818152600260205260409020546001600160a01b03166105855760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610464565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108b1826105e7565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806108f6836105e7565b9050806001600160a01b0316846001600160a01b0316148061091d575061091d8185610779565b806109415750836001600160a01b0316610936846103c9565b6001600160a01b0316145b949350505050565b826001600160a01b031661095c826105e7565b6001600160a01b0316146109825760405162461bcd60e51b815260040161046490611836565b6001600160a01b0382166109e45760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610464565b6109f18383836001610e25565b826001600160a01b0316610a04826105e7565b6001600160a01b031614610a2a5760405162461bcd60e51b815260040161046490611836565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b61058581610ead565b6007546001600160a01b031633146106df5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610464565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610bd05760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610464565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610c48848484610949565b610c5484848484610eed565b6107315760405162461bcd60e51b81526004016104649061187b565b6060610c7b8261081d565b60008281526006602052604081208054610c94906117af565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc0906117af565b8015610d0d5780601f10610ce257610100808354040283529160200191610d0d565b820191906000526020600020905b815481529060010190602001808311610cf057829003601f168201915b505050505090506000610d2b60408051602081019091526000815290565b90508051600003610d3d575092915050565b815115610d6f578082604051602001610d579291906118cd565b60405160208183030381529060405292505050919050565b61094184610fee565b6106fb828260405180602001604052806000815250611062565b6000828152600260205260409020546001600160a01b0316610e0d5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610464565b6000828152600660205260409020610505828261194a565b6001811115610731576001600160a01b03841615610e6b576001600160a01b03841660009081526003602052604081208054839290610e65908490611a20565b90915550505b6001600160a01b03831615610731576001600160a01b03831660009081526003602052604081208054839290610ea2908490611a33565b909155505050505050565b610eb681611095565b60008181526006602052604090208054610ecf906117af565b1590506105855760008181526006602052604081206105859161143c565b60006001600160a01b0384163b15610fe357604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610f31903390899088908890600401611a46565b6020604051808303816000875af1925050508015610f6c575060408051601f3d908101601f19168201909252610f6991810190611a83565b60015b610fc9573d808015610f9a576040519150601f19603f3d011682016040523d82523d6000602084013e610f9f565b606091505b508051600003610fc15760405162461bcd60e51b81526004016104649061187b565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610941565b506001949350505050565b6060610ff98261081d565b600061101060408051602081019091526000815290565b90506000815111611030576040518060200160405280600081525061105b565b8061103a84611138565b60405160200161104b9291906118cd565b6040516020818303038152906040525b9392505050565b61106c83836111cb565b6110796000848484610eed565b6105055760405162461bcd60e51b81526004016104649061187b565b60006110a0826105e7565b90506110b0816000846001610e25565b6110b9826105e7565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6060600061114583611364565b600101905060008167ffffffffffffffff81111561116557611165611612565b6040519080825280601f01601f19166020018201604052801561118f576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461119957509392505050565b6001600160a01b0382166112215760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610464565b6000818152600260205260409020546001600160a01b0316156112865760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610464565b611294600083836001610e25565b6000818152600260205260409020546001600160a01b0316156112f95760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610464565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106113a35772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106113cf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106113ed57662386f26fc10000830492506010015b6305f5e1008310611405576305f5e100830492506008015b612710831061141957612710830492506004015b6064831061142b576064830492506002015b600a83106103315760010192915050565b508054611448906117af565b6000825580601f10611458575050565b601f01602090049060005260206000209081019061058591905b808211156114865760008155600101611472565b5090565b6001600160e01b03198116811461058557600080fd5b6000602082840312156114b257600080fd5b813561105b8161148a565b60005b838110156114d85781810151838201526020016114c0565b50506000910152565b600081518084526114f98160208601602086016114bd565b601f01601f19169290920160200192915050565b60208152600061105b60208301846114e1565b60006020828403121561153257600080fd5b5035919050565b80356001600160a01b038116811461155057600080fd5b919050565b6000806040838503121561156857600080fd5b61157183611539565b946020939093013593505050565b60008060006060848603121561159457600080fd5b61159d84611539565b92506115ab60208501611539565b9150604084013590509250925092565b6000602082840312156115cd57600080fd5b61105b82611539565b600080604083850312156115e957600080fd5b6115f283611539565b91506020830135801515811461160757600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561164357611643611612565b604051601f8501601f19908116603f0116810190828211818310171561166b5761166b611612565b8160405280935085815286868601111561168457600080fd5b858560208301376000602087830101525050509392505050565b600080600080608085870312156116b457600080fd5b6116bd85611539565b93506116cb60208601611539565b925060408501359150606085013567ffffffffffffffff8111156116ee57600080fd5b8501601f810187136116ff57600080fd5b61170e87823560208401611628565b91505092959194509250565b6000806040838503121561172d57600080fd5b61173683611539565b9150602083013567ffffffffffffffff81111561175257600080fd5b8301601f8101851361176357600080fd5b61177285823560208401611628565b9150509250929050565b6000806040838503121561178f57600080fd5b61179883611539565b91506117a660208401611539565b90509250929050565b600181811c908216806117c357607f821691505b6020821081036117e357634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b600083516118df8184602088016114bd565b8351908301906118f38183602088016114bd565b01949350505050565b601f82111561050557600081815260208120601f850160051c810160208610156119235750805b601f850160051c820191505b818110156119425782815560010161192f565b505050505050565b815167ffffffffffffffff81111561196457611964611612565b6119788161197284546117af565b846118fc565b602080601f8311600181146119ad57600084156119955750858301515b600019600386901b1c1916600185901b178555611942565b600085815260208120601f198616915b828110156119dc578886015182559484019460019091019084016119bd565b50858210156119fa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b8181038181111561033157610331611a0a565b8082018082111561033157610331611a0a565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611a79908301846114e1565b9695505050505050565b600060208284031215611a9557600080fd5b815161105b8161148a56fea264697066735822122007b37e5bd0b2127c4de67ce9355957f9dcf71e4ebee1eb5cb062adc857654b1b64736f6c63430008110033",
}

// KuiperNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use KuiperNFTMetaData.ABI instead.
var KuiperNFTABI = KuiperNFTMetaData.ABI

// KuiperNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KuiperNFTMetaData.Bin instead.
var KuiperNFTBin = KuiperNFTMetaData.Bin

// DeployKuiperNFT deploys a new Ethereum contract, binding an instance of KuiperNFT to it.
func DeployKuiperNFT(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _owner common.Address, _royaltyFee *big.Int, _royaltyRecipient common.Address) (common.Address, *types.Transaction, *KuiperNFT, error) {
	parsed, err := KuiperNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KuiperNFTBin), backend, _name, _symbol, _owner, _royaltyFee, _royaltyRecipient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KuiperNFT{KuiperNFTCaller: KuiperNFTCaller{contract: contract}, KuiperNFTTransactor: KuiperNFTTransactor{contract: contract}, KuiperNFTFilterer: KuiperNFTFilterer{contract: contract}}, nil
}

// KuiperNFT is an auto generated Go binding around an Ethereum contract.
type KuiperNFT struct {
	KuiperNFTCaller     // Read-only binding to the contract
	KuiperNFTTransactor // Write-only binding to the contract
	KuiperNFTFilterer   // Log filterer for contract events
}

// KuiperNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type KuiperNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KuiperNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KuiperNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KuiperNFTSession struct {
	Contract     *KuiperNFT        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KuiperNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KuiperNFTCallerSession struct {
	Contract *KuiperNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KuiperNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KuiperNFTTransactorSession struct {
	Contract     *KuiperNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KuiperNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type KuiperNFTRaw struct {
	Contract *KuiperNFT // Generic contract binding to access the raw methods on
}

// KuiperNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KuiperNFTCallerRaw struct {
	Contract *KuiperNFTCaller // Generic read-only contract binding to access the raw methods on
}

// KuiperNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KuiperNFTTransactorRaw struct {
	Contract *KuiperNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKuiperNFT creates a new instance of KuiperNFT, bound to a specific deployed contract.
func NewKuiperNFT(address common.Address, backend bind.ContractBackend) (*KuiperNFT, error) {
	contract, err := bindKuiperNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KuiperNFT{KuiperNFTCaller: KuiperNFTCaller{contract: contract}, KuiperNFTTransactor: KuiperNFTTransactor{contract: contract}, KuiperNFTFilterer: KuiperNFTFilterer{contract: contract}}, nil
}

// NewKuiperNFTCaller creates a new read-only instance of KuiperNFT, bound to a specific deployed contract.
func NewKuiperNFTCaller(address common.Address, caller bind.ContractCaller) (*KuiperNFTCaller, error) {
	contract, err := bindKuiperNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTCaller{contract: contract}, nil
}

// NewKuiperNFTTransactor creates a new write-only instance of KuiperNFT, bound to a specific deployed contract.
func NewKuiperNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*KuiperNFTTransactor, error) {
	contract, err := bindKuiperNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTTransactor{contract: contract}, nil
}

// NewKuiperNFTFilterer creates a new log filterer instance of KuiperNFT, bound to a specific deployed contract.
func NewKuiperNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*KuiperNFTFilterer, error) {
	contract, err := bindKuiperNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFilterer{contract: contract}, nil
}

// bindKuiperNFT binds a generic wrapper to an already deployed contract.
func bindKuiperNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KuiperNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KuiperNFT *KuiperNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KuiperNFT.Contract.KuiperNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KuiperNFT *KuiperNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KuiperNFT.Contract.KuiperNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KuiperNFT *KuiperNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KuiperNFT.Contract.KuiperNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KuiperNFT *KuiperNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KuiperNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KuiperNFT *KuiperNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KuiperNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KuiperNFT *KuiperNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KuiperNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KuiperNFT *KuiperNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KuiperNFT *KuiperNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KuiperNFT.Contract.BalanceOf(&_KuiperNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_KuiperNFT *KuiperNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _KuiperNFT.Contract.BalanceOf(&_KuiperNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KuiperNFT.Contract.GetApproved(&_KuiperNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _KuiperNFT.Contract.GetApproved(&_KuiperNFT.CallOpts, tokenId)
}

// GetRoyaltyFee is a free data retrieval call binding the contract method 0x820bdcdc.
//
// Solidity: function getRoyaltyFee() view returns(uint256)
func (_KuiperNFT *KuiperNFTCaller) GetRoyaltyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "getRoyaltyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoyaltyFee is a free data retrieval call binding the contract method 0x820bdcdc.
//
// Solidity: function getRoyaltyFee() view returns(uint256)
func (_KuiperNFT *KuiperNFTSession) GetRoyaltyFee() (*big.Int, error) {
	return _KuiperNFT.Contract.GetRoyaltyFee(&_KuiperNFT.CallOpts)
}

// GetRoyaltyFee is a free data retrieval call binding the contract method 0x820bdcdc.
//
// Solidity: function getRoyaltyFee() view returns(uint256)
func (_KuiperNFT *KuiperNFTCallerSession) GetRoyaltyFee() (*big.Int, error) {
	return _KuiperNFT.Contract.GetRoyaltyFee(&_KuiperNFT.CallOpts)
}

// GetRoyaltyRecipient is a free data retrieval call binding the contract method 0x95edc18c.
//
// Solidity: function getRoyaltyRecipient() view returns(address)
func (_KuiperNFT *KuiperNFTCaller) GetRoyaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "getRoyaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoyaltyRecipient is a free data retrieval call binding the contract method 0x95edc18c.
//
// Solidity: function getRoyaltyRecipient() view returns(address)
func (_KuiperNFT *KuiperNFTSession) GetRoyaltyRecipient() (common.Address, error) {
	return _KuiperNFT.Contract.GetRoyaltyRecipient(&_KuiperNFT.CallOpts)
}

// GetRoyaltyRecipient is a free data retrieval call binding the contract method 0x95edc18c.
//
// Solidity: function getRoyaltyRecipient() view returns(address)
func (_KuiperNFT *KuiperNFTCallerSession) GetRoyaltyRecipient() (common.Address, error) {
	return _KuiperNFT.Contract.GetRoyaltyRecipient(&_KuiperNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KuiperNFT *KuiperNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KuiperNFT *KuiperNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KuiperNFT.Contract.IsApprovedForAll(&_KuiperNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_KuiperNFT *KuiperNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _KuiperNFT.Contract.IsApprovedForAll(&_KuiperNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KuiperNFT *KuiperNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KuiperNFT *KuiperNFTSession) Name() (string, error) {
	return _KuiperNFT.Contract.Name(&_KuiperNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KuiperNFT *KuiperNFTCallerSession) Name() (string, error) {
	return _KuiperNFT.Contract.Name(&_KuiperNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KuiperNFT *KuiperNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KuiperNFT *KuiperNFTSession) Owner() (common.Address, error) {
	return _KuiperNFT.Contract.Owner(&_KuiperNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KuiperNFT *KuiperNFTCallerSession) Owner() (common.Address, error) {
	return _KuiperNFT.Contract.Owner(&_KuiperNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KuiperNFT.Contract.OwnerOf(&_KuiperNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_KuiperNFT *KuiperNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _KuiperNFT.Contract.OwnerOf(&_KuiperNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KuiperNFT *KuiperNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KuiperNFT *KuiperNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KuiperNFT.Contract.SupportsInterface(&_KuiperNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KuiperNFT *KuiperNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KuiperNFT.Contract.SupportsInterface(&_KuiperNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KuiperNFT *KuiperNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KuiperNFT *KuiperNFTSession) Symbol() (string, error) {
	return _KuiperNFT.Contract.Symbol(&_KuiperNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KuiperNFT *KuiperNFTCallerSession) Symbol() (string, error) {
	return _KuiperNFT.Contract.Symbol(&_KuiperNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KuiperNFT *KuiperNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _KuiperNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KuiperNFT *KuiperNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KuiperNFT.Contract.TokenURI(&_KuiperNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_KuiperNFT *KuiperNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _KuiperNFT.Contract.TokenURI(&_KuiperNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.Approve(&_KuiperNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.Approve(&_KuiperNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.Burn(&_KuiperNFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.Burn(&_KuiperNFT.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KuiperNFT *KuiperNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KuiperNFT *KuiperNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _KuiperNFT.Contract.RenounceOwnership(&_KuiperNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KuiperNFT *KuiperNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KuiperNFT.Contract.RenounceOwnership(&_KuiperNFT.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns()
func (_KuiperNFT *KuiperNFTTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "safeMint", to, uri)
}

// SafeMint is a paid mutator transaction binding the contract method 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns()
func (_KuiperNFT *KuiperNFTSession) SafeMint(to common.Address, uri string) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeMint(&_KuiperNFT.TransactOpts, to, uri)
}

// SafeMint is a paid mutator transaction binding the contract method 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) SafeMint(to common.Address, uri string) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeMint(&_KuiperNFT.TransactOpts, to, uri)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeTransferFrom(&_KuiperNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeTransferFrom(&_KuiperNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_KuiperNFT *KuiperNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_KuiperNFT *KuiperNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeTransferFrom0(&_KuiperNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SafeTransferFrom0(&_KuiperNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_KuiperNFT *KuiperNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_KuiperNFT *KuiperNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SetApprovalForAll(&_KuiperNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _KuiperNFT.Contract.SetApprovalForAll(&_KuiperNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.TransferFrom(&_KuiperNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.TransferFrom(&_KuiperNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KuiperNFT *KuiperNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KuiperNFT *KuiperNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KuiperNFT.Contract.TransferOwnership(&_KuiperNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KuiperNFT.Contract.TransferOwnership(&_KuiperNFT.TransactOpts, newOwner)
}

// UpdateRoyaltyFee is a paid mutator transaction binding the contract method 0x4e83be47.
//
// Solidity: function updateRoyaltyFee(uint256 _royaltyFee) returns()
func (_KuiperNFT *KuiperNFTTransactor) UpdateRoyaltyFee(opts *bind.TransactOpts, _royaltyFee *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.contract.Transact(opts, "updateRoyaltyFee", _royaltyFee)
}

// UpdateRoyaltyFee is a paid mutator transaction binding the contract method 0x4e83be47.
//
// Solidity: function updateRoyaltyFee(uint256 _royaltyFee) returns()
func (_KuiperNFT *KuiperNFTSession) UpdateRoyaltyFee(_royaltyFee *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.UpdateRoyaltyFee(&_KuiperNFT.TransactOpts, _royaltyFee)
}

// UpdateRoyaltyFee is a paid mutator transaction binding the contract method 0x4e83be47.
//
// Solidity: function updateRoyaltyFee(uint256 _royaltyFee) returns()
func (_KuiperNFT *KuiperNFTTransactorSession) UpdateRoyaltyFee(_royaltyFee *big.Int) (*types.Transaction, error) {
	return _KuiperNFT.Contract.UpdateRoyaltyFee(&_KuiperNFT.TransactOpts, _royaltyFee)
}

// KuiperNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KuiperNFT contract.
type KuiperNFTApprovalIterator struct {
	Event *KuiperNFTApproval // Event containing the contract specifics and raw log

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
func (it *KuiperNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KuiperNFTApproval)
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
		it.Event = new(KuiperNFTApproval)
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
func (it *KuiperNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KuiperNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KuiperNFTApproval represents a Approval event raised by the KuiperNFT contract.
type KuiperNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KuiperNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _KuiperNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTApprovalIterator{contract: _KuiperNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KuiperNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _KuiperNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KuiperNFTApproval)
				if err := _KuiperNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) ParseApproval(log types.Log) (*KuiperNFTApproval, error) {
	event := new(KuiperNFTApproval)
	if err := _KuiperNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KuiperNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the KuiperNFT contract.
type KuiperNFTApprovalForAllIterator struct {
	Event *KuiperNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KuiperNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KuiperNFTApprovalForAll)
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
		it.Event = new(KuiperNFTApprovalForAll)
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
func (it *KuiperNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KuiperNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KuiperNFTApprovalForAll represents a ApprovalForAll event raised by the KuiperNFT contract.
type KuiperNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KuiperNFT *KuiperNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KuiperNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KuiperNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTApprovalForAllIterator{contract: _KuiperNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_KuiperNFT *KuiperNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KuiperNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _KuiperNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KuiperNFTApprovalForAll)
				if err := _KuiperNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_KuiperNFT *KuiperNFTFilterer) ParseApprovalForAll(log types.Log) (*KuiperNFTApprovalForAll, error) {
	event := new(KuiperNFTApprovalForAll)
	if err := _KuiperNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KuiperNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KuiperNFT contract.
type KuiperNFTOwnershipTransferredIterator struct {
	Event *KuiperNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KuiperNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KuiperNFTOwnershipTransferred)
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
		it.Event = new(KuiperNFTOwnershipTransferred)
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
func (it *KuiperNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KuiperNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KuiperNFTOwnershipTransferred represents a OwnershipTransferred event raised by the KuiperNFT contract.
type KuiperNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KuiperNFT *KuiperNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KuiperNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KuiperNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTOwnershipTransferredIterator{contract: _KuiperNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KuiperNFT *KuiperNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KuiperNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KuiperNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KuiperNFTOwnershipTransferred)
				if err := _KuiperNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KuiperNFT *KuiperNFTFilterer) ParseOwnershipTransferred(log types.Log) (*KuiperNFTOwnershipTransferred, error) {
	event := new(KuiperNFTOwnershipTransferred)
	if err := _KuiperNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KuiperNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KuiperNFT contract.
type KuiperNFTTransferIterator struct {
	Event *KuiperNFTTransfer // Event containing the contract specifics and raw log

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
func (it *KuiperNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KuiperNFTTransfer)
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
		it.Event = new(KuiperNFTTransfer)
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
func (it *KuiperNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KuiperNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KuiperNFTTransfer represents a Transfer event raised by the KuiperNFT contract.
type KuiperNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KuiperNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _KuiperNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTTransferIterator{contract: _KuiperNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KuiperNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _KuiperNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KuiperNFTTransfer)
				if err := _KuiperNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_KuiperNFT *KuiperNFTFilterer) ParseTransfer(log types.Log) (*KuiperNFTTransfer, error) {
	event := new(KuiperNFTTransfer)
	if err := _KuiperNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
