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

// KuiperNFTFactoryMetaData contains all meta data concerning the KuiperNFTFactory contract.
var KuiperNFTFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CreatedNFTCollection\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"}],\"name\":\"createNFTCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnCollections\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"name\":\"isKuiperNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061250f806100206000396000f3fe60806040523480156200001157600080fd5b5060043610620000465760003560e01c80637a3c7a2c146200004b578063cdbe0fc31462000064578063debe518d1462000086575b600080fd5b620000626200005c366004620002e2565b620000c6565b005b6200006e620001a1565b6040516200007d919062000369565b60405180910390f35b620000b562000097366004620003b8565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020016200007d565b60008484338585604051620000db906200020c565b620000eb95949392919062000425565b604051809103906000f08015801562000108573d6000803e3d6000fd5b50336000818152602081815260408083208054600180820183559185528385200180546001600160a01b0319166001600160a01b03881690811790915584529182905291829020805460ff19169091179055519192507f63769cea406878ef0dbb67fc60c62a7c20747798c2703734c451fa2e8381ef249162000192919084908990899062000479565b60405180910390a15050505050565b33600090815260208181526040918290208054835181840281018401909452808452606093928301828280156200020257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311620001e3575b5050505050905090565b61201380620004c783390190565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200024257600080fd5b813567ffffffffffffffff808211156200026057620002606200021a565b604051601f8301601f19908116603f011681019082821181831017156200028b576200028b6200021a565b81604052838152866020858801011115620002a557600080fd5b836020870160208301376000602085830101528094505050505092915050565b80356001600160a01b0381168114620002dd57600080fd5b919050565b60008060008060808587031215620002f957600080fd5b843567ffffffffffffffff808211156200031257600080fd5b620003208883890162000230565b955060208701359150808211156200033757600080fd5b50620003468782880162000230565b935050604085013591506200035e60608601620002c5565b905092959194509250565b6020808252825182820181905260009190848201906040850190845b81811015620003ac5783516001600160a01b03168352928401929184019160010162000385565b50909695505050505050565b600060208284031215620003cb57600080fd5b620003d682620002c5565b9392505050565b6000815180845260005b818110156200040557602081850181015186830182015201620003e7565b506000602082860101526020601f19601f83011685010191505092915050565b60a0815260006200043a60a0830188620003dd565b82810360208401526200044e8188620003dd565b6001600160a01b03968716604085015260608401959095525050921660809092019190915292915050565b6001600160a01b03858116825284166020820152608060408201819052600090620004a790830185620003dd565b8281036060840152620004bb8185620003dd565b97965050505050505056fe60806040523480156200001157600080fd5b506040516200201338038062002013833981016040819052620000349162000339565b8484600062000044838262000461565b50600162000053828262000461565b505050620000706200006a6200012460201b60201c565b62000128565b612710821115620000da5760405162461bcd60e51b815260206004820152602960248201527f726f79616c7479206665652063616e2774206265206d6f7265207468616e20316044820152680c081c195c98d95b9d60ba1b60648201526084015b60405180910390fd5b6001600160a01b038116620000ee57600080fd5b6009829055600a80546001600160a01b0319166001600160a01b03831617905562000119836200017a565b50505050506200052d565b3390565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b62000184620001f9565b6001600160a01b038116620001eb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000d1565b620001f68162000128565b50565b6007546001600160a01b03163314620002555760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000d1565b565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200027f57600080fd5b81516001600160401b03808211156200029c576200029c62000257565b604051601f8301601f19908116603f01168101908282118183101715620002c757620002c762000257565b81604052838152602092508683858801011115620002e457600080fd5b600091505b83821015620003085785820183015181830184015290820190620002e9565b600093810190920192909252949350505050565b80516001600160a01b03811681146200033457600080fd5b919050565b600080600080600060a086880312156200035257600080fd5b85516001600160401b03808211156200036a57600080fd5b6200037889838a016200026d565b965060208801519150808211156200038f57600080fd5b506200039e888289016200026d565b945050620003af604087016200031c565b925060608601519150620003c6608087016200031c565b90509295509295909350565b600181811c90821680620003e757607f821691505b6020821081036200040857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200045c57600081815260208120601f850160051c81016020861015620004375750805b601f850160051c820191505b81811015620004585782815560010162000443565b5050505b505050565b81516001600160401b038111156200047d576200047d62000257565b62000495816200048e8454620003d2565b846200040e565b602080601f831160018114620004cd5760008415620004b45750858301515b600019600386901b1c1916600185901b17855562000458565b600085815260208120601f198616915b82811015620004fe57888601518255948401946001909101908401620004dd565b50858210156200051d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611ad6806200053d6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063715018a6116100b8578063a22cb4651161007c578063a22cb46514610273578063b88d4fde14610286578063c87b56dd14610299578063d204c45e146102ac578063e985e9c5146102bf578063f2fde38b146102d257600080fd5b8063715018a614610239578063820bdcdc146102415780638da5cb5b1461024957806395d89b411461025a57806395edc18c1461026257600080fd5b806342842e0e116100ff57806342842e0e146101cc57806342966c68146101df5780634e83be47146101f25780636352211e1461020557806370a082311461021857600080fd5b806301ffc9a71461013c57806306fdde0314610164578063081812fc14610179578063095ea7b3146101a457806323b872dd146101b9575b600080fd5b61014f61014a3660046114a0565b6102e5565b60405190151581526020015b60405180910390f35b61016c610337565b60405161015b919061150d565b61018c610187366004611520565b6103c9565b6040516001600160a01b03909116815260200161015b565b6101b76101b2366004611555565b6103f0565b005b6101b76101c736600461157f565b61050a565b6101b76101da36600461157f565b61053c565b6101b76101ed366004611520565b610557565b6101b7610200366004611520565b610588565b61018c610213366004611520565b6105e7565b61022b6102263660046115bb565b610647565b60405190815260200161015b565b6101b76106cd565b60095461022b565b6007546001600160a01b031661018c565b61016c6106e1565b600a546001600160a01b031661018c565b6101b76102813660046115d6565b6106f0565b6101b761029436600461169e565b6106ff565b61016c6102a7366004611520565b610737565b6101b76102ba36600461171a565b610742565b61014f6102cd36600461177c565b610779565b6101b76102e03660046115bb565b6107a7565b60006001600160e01b031982166380ac58cd60e01b148061031657506001600160e01b03198216635b5e139f60e01b145b8061033157506301ffc9a760e01b6001600160e01b03198316145b92915050565b606060008054610346906117af565b80601f0160208091040260200160405190810160405280929190818152602001828054610372906117af565b80156103bf5780601f10610394576101008083540402835291602001916103bf565b820191906000526020600020905b8154815290600101906020018083116103a257829003601f168201915b5050505050905090565b60006103d48261081d565b506000908152600460205260409020546001600160a01b031690565b60006103fb826105e7565b9050806001600160a01b0316836001600160a01b03160361046d5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061048957506104898133610779565b6104fb5760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c0000006064820152608401610464565b610505838361087c565b505050565b610515335b826108ea565b6105315760405162461bcd60e51b8152600401610464906117e9565b610505838383610949565b610505838383604051806020016040528060008152506106ff565b6105603361050f565b61057c5760405162461bcd60e51b8152600401610464906117e9565b61058581610aba565b50565b610590610ac3565b6127108111156105e25760405162461bcd60e51b815260206004820152601a60248201527f63616e2774206d6f7265207468616e2031302070657263656e740000000000006044820152606401610464565b600955565b6000818152600260205260408120546001600160a01b0316806103315760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610464565b60006001600160a01b0382166106b15760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b6064820152608401610464565b506001600160a01b031660009081526003602052604090205490565b6106d5610ac3565b6106df6000610b1d565b565b606060018054610346906117af565b6106fb338383610b6f565b5050565b61070933836108ea565b6107255760405162461bcd60e51b8152600401610464906117e9565b61073184848484610c3d565b50505050565b606061033182610c70565b61074a610ac3565b600061075560085490565b9050610765600880546001019055565b61076f8382610d78565b6105058183610d92565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6107af610ac3565b6001600160a01b0381166108145760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610464565b61058581610b1d565b6000818152600260205260409020546001600160a01b03166105855760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b6044820152606401610464565b600081815260046020526040902080546001600160a01b0319166001600160a01b03841690811790915581906108b1826105e7565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806108f6836105e7565b9050806001600160a01b0316846001600160a01b0316148061091d575061091d8185610779565b806109415750836001600160a01b0316610936846103c9565b6001600160a01b0316145b949350505050565b826001600160a01b031661095c826105e7565b6001600160a01b0316146109825760405162461bcd60e51b815260040161046490611836565b6001600160a01b0382166109e45760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610464565b6109f18383836001610e25565b826001600160a01b0316610a04826105e7565b6001600160a01b031614610a2a5760405162461bcd60e51b815260040161046490611836565b600081815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260038552838620805460001901905590871680865283862080546001019055868652600290945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b61058581610ead565b6007546001600160a01b031633146106df5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610464565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b816001600160a01b0316836001600160a01b031603610bd05760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610464565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b610c48848484610949565b610c5484848484610eed565b6107315760405162461bcd60e51b81526004016104649061187b565b6060610c7b8261081d565b60008281526006602052604081208054610c94906117af565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc0906117af565b8015610d0d5780601f10610ce257610100808354040283529160200191610d0d565b820191906000526020600020905b815481529060010190602001808311610cf057829003601f168201915b505050505090506000610d2b60408051602081019091526000815290565b90508051600003610d3d575092915050565b815115610d6f578082604051602001610d579291906118cd565b60405160208183030381529060405292505050919050565b61094184610fee565b6106fb828260405180602001604052806000815250611062565b6000828152600260205260409020546001600160a01b0316610e0d5760405162461bcd60e51b815260206004820152602e60248201527f45524337323155524953746f726167653a2055524920736574206f66206e6f6e60448201526d32bc34b9ba32b73a103a37b5b2b760911b6064820152608401610464565b6000828152600660205260409020610505828261194a565b6001811115610731576001600160a01b03841615610e6b576001600160a01b03841660009081526003602052604081208054839290610e65908490611a20565b90915550505b6001600160a01b03831615610731576001600160a01b03831660009081526003602052604081208054839290610ea2908490611a33565b909155505050505050565b610eb681611095565b60008181526006602052604090208054610ecf906117af565b1590506105855760008181526006602052604081206105859161143c565b60006001600160a01b0384163b15610fe357604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290610f31903390899088908890600401611a46565b6020604051808303816000875af1925050508015610f6c575060408051601f3d908101601f19168201909252610f6991810190611a83565b60015b610fc9573d808015610f9a576040519150601f19603f3d011682016040523d82523d6000602084013e610f9f565b606091505b508051600003610fc15760405162461bcd60e51b81526004016104649061187b565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610941565b506001949350505050565b6060610ff98261081d565b600061101060408051602081019091526000815290565b90506000815111611030576040518060200160405280600081525061105b565b8061103a84611138565b60405160200161104b9291906118cd565b6040516020818303038152906040525b9392505050565b61106c83836111cb565b6110796000848484610eed565b6105055760405162461bcd60e51b81526004016104649061187b565b60006110a0826105e7565b90506110b0816000846001610e25565b6110b9826105e7565b600083815260046020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526003845282852080546000190190558785526002909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6060600061114583611364565b600101905060008167ffffffffffffffff81111561116557611165611612565b6040519080825280601f01601f19166020018201604052801561118f576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461119957509392505050565b6001600160a01b0382166112215760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610464565b6000818152600260205260409020546001600160a01b0316156112865760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610464565b611294600083836001610e25565b6000818152600260205260409020546001600160a01b0316156112f95760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610464565b6001600160a01b038216600081815260036020908152604080832080546001019055848352600290915280822080546001600160a01b0319168417905551839291907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106113a35772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106113cf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106113ed57662386f26fc10000830492506010015b6305f5e1008310611405576305f5e100830492506008015b612710831061141957612710830492506004015b6064831061142b576064830492506002015b600a83106103315760010192915050565b508054611448906117af565b6000825580601f10611458575050565b601f01602090049060005260206000209081019061058591905b808211156114865760008155600101611472565b5090565b6001600160e01b03198116811461058557600080fd5b6000602082840312156114b257600080fd5b813561105b8161148a565b60005b838110156114d85781810151838201526020016114c0565b50506000910152565b600081518084526114f98160208601602086016114bd565b601f01601f19169290920160200192915050565b60208152600061105b60208301846114e1565b60006020828403121561153257600080fd5b5035919050565b80356001600160a01b038116811461155057600080fd5b919050565b6000806040838503121561156857600080fd5b61157183611539565b946020939093013593505050565b60008060006060848603121561159457600080fd5b61159d84611539565b92506115ab60208501611539565b9150604084013590509250925092565b6000602082840312156115cd57600080fd5b61105b82611539565b600080604083850312156115e957600080fd5b6115f283611539565b91506020830135801515811461160757600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561164357611643611612565b604051601f8501601f19908116603f0116810190828211818310171561166b5761166b611612565b8160405280935085815286868601111561168457600080fd5b858560208301376000602087830101525050509392505050565b600080600080608085870312156116b457600080fd5b6116bd85611539565b93506116cb60208601611539565b925060408501359150606085013567ffffffffffffffff8111156116ee57600080fd5b8501601f810187136116ff57600080fd5b61170e87823560208401611628565b91505092959194509250565b6000806040838503121561172d57600080fd5b61173683611539565b9150602083013567ffffffffffffffff81111561175257600080fd5b8301601f8101851361176357600080fd5b61177285823560208401611628565b9150509250929050565b6000806040838503121561178f57600080fd5b61179883611539565b91506117a660208401611539565b90509250929050565b600181811c908216806117c357607f821691505b6020821081036117e357634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b600083516118df8184602088016114bd565b8351908301906118f38183602088016114bd565b01949350505050565b601f82111561050557600081815260208120601f850160051c810160208610156119235750805b601f850160051c820191505b818110156119425782815560010161192f565b505050505050565b815167ffffffffffffffff81111561196457611964611612565b6119788161197284546117af565b846118fc565b602080601f8311600181146119ad57600084156119955750858301515b600019600386901b1c1916600185901b178555611942565b600085815260208120601f198616915b828110156119dc578886015182559484019460019091019084016119bd565b50858210156119fa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b8181038181111561033157610331611a0a565b8082018082111561033157610331611a0a565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611a79908301846114e1565b9695505050505050565b600060208284031215611a9557600080fd5b815161105b8161148a56fea264697066735822122007b37e5bd0b2127c4de67ce9355957f9dcf71e4ebee1eb5cb062adc857654b1b64736f6c63430008110033a26469706673582212201ab5814d9164b112e66bd3cba6aaa81aed3a972562bc9fbd6b07a193dd9e8b8264736f6c63430008110033",
}

// KuiperNFTFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use KuiperNFTFactoryMetaData.ABI instead.
var KuiperNFTFactoryABI = KuiperNFTFactoryMetaData.ABI

// KuiperNFTFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KuiperNFTFactoryMetaData.Bin instead.
var KuiperNFTFactoryBin = KuiperNFTFactoryMetaData.Bin

// DeployKuiperNFTFactory deploys a new Ethereum contract, binding an instance of KuiperNFTFactory to it.
func DeployKuiperNFTFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KuiperNFTFactory, error) {
	parsed, err := KuiperNFTFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KuiperNFTFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KuiperNFTFactory{KuiperNFTFactoryCaller: KuiperNFTFactoryCaller{contract: contract}, KuiperNFTFactoryTransactor: KuiperNFTFactoryTransactor{contract: contract}, KuiperNFTFactoryFilterer: KuiperNFTFactoryFilterer{contract: contract}}, nil
}

// KuiperNFTFactory is an auto generated Go binding around an Ethereum contract.
type KuiperNFTFactory struct {
	KuiperNFTFactoryCaller     // Read-only binding to the contract
	KuiperNFTFactoryTransactor // Write-only binding to the contract
	KuiperNFTFactoryFilterer   // Log filterer for contract events
}

// KuiperNFTFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KuiperNFTFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KuiperNFTFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KuiperNFTFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KuiperNFTFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KuiperNFTFactorySession struct {
	Contract     *KuiperNFTFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KuiperNFTFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KuiperNFTFactoryCallerSession struct {
	Contract *KuiperNFTFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KuiperNFTFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KuiperNFTFactoryTransactorSession struct {
	Contract     *KuiperNFTFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KuiperNFTFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KuiperNFTFactoryRaw struct {
	Contract *KuiperNFTFactory // Generic contract binding to access the raw methods on
}

// KuiperNFTFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KuiperNFTFactoryCallerRaw struct {
	Contract *KuiperNFTFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KuiperNFTFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KuiperNFTFactoryTransactorRaw struct {
	Contract *KuiperNFTFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKuiperNFTFactory creates a new instance of KuiperNFTFactory, bound to a specific deployed contract.
func NewKuiperNFTFactory(address common.Address, backend bind.ContractBackend) (*KuiperNFTFactory, error) {
	contract, err := bindKuiperNFTFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFactory{KuiperNFTFactoryCaller: KuiperNFTFactoryCaller{contract: contract}, KuiperNFTFactoryTransactor: KuiperNFTFactoryTransactor{contract: contract}, KuiperNFTFactoryFilterer: KuiperNFTFactoryFilterer{contract: contract}}, nil
}

// NewKuiperNFTFactoryCaller creates a new read-only instance of KuiperNFTFactory, bound to a specific deployed contract.
func NewKuiperNFTFactoryCaller(address common.Address, caller bind.ContractCaller) (*KuiperNFTFactoryCaller, error) {
	contract, err := bindKuiperNFTFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFactoryCaller{contract: contract}, nil
}

// NewKuiperNFTFactoryTransactor creates a new write-only instance of KuiperNFTFactory, bound to a specific deployed contract.
func NewKuiperNFTFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KuiperNFTFactoryTransactor, error) {
	contract, err := bindKuiperNFTFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFactoryTransactor{contract: contract}, nil
}

// NewKuiperNFTFactoryFilterer creates a new log filterer instance of KuiperNFTFactory, bound to a specific deployed contract.
func NewKuiperNFTFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KuiperNFTFactoryFilterer, error) {
	contract, err := bindKuiperNFTFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFactoryFilterer{contract: contract}, nil
}

// bindKuiperNFTFactory binds a generic wrapper to an already deployed contract.
func bindKuiperNFTFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KuiperNFTFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KuiperNFTFactory *KuiperNFTFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KuiperNFTFactory.Contract.KuiperNFTFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KuiperNFTFactory *KuiperNFTFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.KuiperNFTFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KuiperNFTFactory *KuiperNFTFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.KuiperNFTFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KuiperNFTFactory *KuiperNFTFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KuiperNFTFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KuiperNFTFactory *KuiperNFTFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KuiperNFTFactory *KuiperNFTFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.contract.Transact(opts, method, params...)
}

// GetOwnCollections is a free data retrieval call binding the contract method 0xcdbe0fc3.
//
// Solidity: function getOwnCollections() view returns(address[])
func (_KuiperNFTFactory *KuiperNFTFactoryCaller) GetOwnCollections(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KuiperNFTFactory.contract.Call(opts, &out, "getOwnCollections")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwnCollections is a free data retrieval call binding the contract method 0xcdbe0fc3.
//
// Solidity: function getOwnCollections() view returns(address[])
func (_KuiperNFTFactory *KuiperNFTFactorySession) GetOwnCollections() ([]common.Address, error) {
	return _KuiperNFTFactory.Contract.GetOwnCollections(&_KuiperNFTFactory.CallOpts)
}

// GetOwnCollections is a free data retrieval call binding the contract method 0xcdbe0fc3.
//
// Solidity: function getOwnCollections() view returns(address[])
func (_KuiperNFTFactory *KuiperNFTFactoryCallerSession) GetOwnCollections() ([]common.Address, error) {
	return _KuiperNFTFactory.Contract.GetOwnCollections(&_KuiperNFTFactory.CallOpts)
}

// IsKuiperNFT is a free data retrieval call binding the contract method 0xdebe518d.
//
// Solidity: function isKuiperNFT(address _nft) view returns(bool)
func (_KuiperNFTFactory *KuiperNFTFactoryCaller) IsKuiperNFT(opts *bind.CallOpts, _nft common.Address) (bool, error) {
	var out []interface{}
	err := _KuiperNFTFactory.contract.Call(opts, &out, "isKuiperNFT", _nft)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKuiperNFT is a free data retrieval call binding the contract method 0xdebe518d.
//
// Solidity: function isKuiperNFT(address _nft) view returns(bool)
func (_KuiperNFTFactory *KuiperNFTFactorySession) IsKuiperNFT(_nft common.Address) (bool, error) {
	return _KuiperNFTFactory.Contract.IsKuiperNFT(&_KuiperNFTFactory.CallOpts, _nft)
}

// IsKuiperNFT is a free data retrieval call binding the contract method 0xdebe518d.
//
// Solidity: function isKuiperNFT(address _nft) view returns(bool)
func (_KuiperNFTFactory *KuiperNFTFactoryCallerSession) IsKuiperNFT(_nft common.Address) (bool, error) {
	return _KuiperNFTFactory.Contract.IsKuiperNFT(&_KuiperNFTFactory.CallOpts, _nft)
}

// CreateNFTCollection is a paid mutator transaction binding the contract method 0x7a3c7a2c.
//
// Solidity: function createNFTCollection(string _name, string _symbol, uint256 _royaltyFee, address _royaltyRecipient) returns()
func (_KuiperNFTFactory *KuiperNFTFactoryTransactor) CreateNFTCollection(opts *bind.TransactOpts, _name string, _symbol string, _royaltyFee *big.Int, _royaltyRecipient common.Address) (*types.Transaction, error) {
	return _KuiperNFTFactory.contract.Transact(opts, "createNFTCollection", _name, _symbol, _royaltyFee, _royaltyRecipient)
}

// CreateNFTCollection is a paid mutator transaction binding the contract method 0x7a3c7a2c.
//
// Solidity: function createNFTCollection(string _name, string _symbol, uint256 _royaltyFee, address _royaltyRecipient) returns()
func (_KuiperNFTFactory *KuiperNFTFactorySession) CreateNFTCollection(_name string, _symbol string, _royaltyFee *big.Int, _royaltyRecipient common.Address) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.CreateNFTCollection(&_KuiperNFTFactory.TransactOpts, _name, _symbol, _royaltyFee, _royaltyRecipient)
}

// CreateNFTCollection is a paid mutator transaction binding the contract method 0x7a3c7a2c.
//
// Solidity: function createNFTCollection(string _name, string _symbol, uint256 _royaltyFee, address _royaltyRecipient) returns()
func (_KuiperNFTFactory *KuiperNFTFactoryTransactorSession) CreateNFTCollection(_name string, _symbol string, _royaltyFee *big.Int, _royaltyRecipient common.Address) (*types.Transaction, error) {
	return _KuiperNFTFactory.Contract.CreateNFTCollection(&_KuiperNFTFactory.TransactOpts, _name, _symbol, _royaltyFee, _royaltyRecipient)
}

// KuiperNFTFactoryCreatedNFTCollectionIterator is returned from FilterCreatedNFTCollection and is used to iterate over the raw logs and unpacked data for CreatedNFTCollection events raised by the KuiperNFTFactory contract.
type KuiperNFTFactoryCreatedNFTCollectionIterator struct {
	Event *KuiperNFTFactoryCreatedNFTCollection // Event containing the contract specifics and raw log

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
func (it *KuiperNFTFactoryCreatedNFTCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KuiperNFTFactoryCreatedNFTCollection)
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
		it.Event = new(KuiperNFTFactoryCreatedNFTCollection)
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
func (it *KuiperNFTFactoryCreatedNFTCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KuiperNFTFactoryCreatedNFTCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KuiperNFTFactoryCreatedNFTCollection represents a CreatedNFTCollection event raised by the KuiperNFTFactory contract.
type KuiperNFTFactoryCreatedNFTCollection struct {
	Creator common.Address
	Nft     common.Address
	Name    string
	Symbol  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatedNFTCollection is a free log retrieval operation binding the contract event 0x63769cea406878ef0dbb67fc60c62a7c20747798c2703734c451fa2e8381ef24.
//
// Solidity: event CreatedNFTCollection(address creator, address nft, string name, string symbol)
func (_KuiperNFTFactory *KuiperNFTFactoryFilterer) FilterCreatedNFTCollection(opts *bind.FilterOpts) (*KuiperNFTFactoryCreatedNFTCollectionIterator, error) {

	logs, sub, err := _KuiperNFTFactory.contract.FilterLogs(opts, "CreatedNFTCollection")
	if err != nil {
		return nil, err
	}
	return &KuiperNFTFactoryCreatedNFTCollectionIterator{contract: _KuiperNFTFactory.contract, event: "CreatedNFTCollection", logs: logs, sub: sub}, nil
}

// WatchCreatedNFTCollection is a free log subscription operation binding the contract event 0x63769cea406878ef0dbb67fc60c62a7c20747798c2703734c451fa2e8381ef24.
//
// Solidity: event CreatedNFTCollection(address creator, address nft, string name, string symbol)
func (_KuiperNFTFactory *KuiperNFTFactoryFilterer) WatchCreatedNFTCollection(opts *bind.WatchOpts, sink chan<- *KuiperNFTFactoryCreatedNFTCollection) (event.Subscription, error) {

	logs, sub, err := _KuiperNFTFactory.contract.WatchLogs(opts, "CreatedNFTCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KuiperNFTFactoryCreatedNFTCollection)
				if err := _KuiperNFTFactory.contract.UnpackLog(event, "CreatedNFTCollection", log); err != nil {
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

// ParseCreatedNFTCollection is a log parse operation binding the contract event 0x63769cea406878ef0dbb67fc60c62a7c20747798c2703734c451fa2e8381ef24.
//
// Solidity: event CreatedNFTCollection(address creator, address nft, string name, string symbol)
func (_KuiperNFTFactory *KuiperNFTFactoryFilterer) ParseCreatedNFTCollection(log types.Log) (*KuiperNFTFactoryCreatedNFTCollection, error) {
	event := new(KuiperNFTFactoryCreatedNFTCollection)
	if err := _KuiperNFTFactory.contract.UnpackLog(event, "CreatedNFTCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
