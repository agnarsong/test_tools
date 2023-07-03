// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// L1CustomERC20MetaData contains all meta data concerning the L1CustomERC20 contract.
var L1CustomERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620024583803806200245883398181016040528101906200003791906200047d565b818181600390816200004a91906200074d565b5080600490816200005c91906200074d565b5050506200007f62000073620000a560201b60201c565b620000ad60201b60201c565b6200009d336b1fe9af603f173559996800006200017360201b60201c565b50506200094f565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620001e5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001dc9062000895565b60405180910390fd5b620001f960008383620002e060201b60201c565b80600260008282546200020d9190620008e6565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002c0919062000932565b60405180910390a3620002dc60008383620002e560201b60201c565b5050565b505050565b505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620003538262000308565b810181811067ffffffffffffffff8211171562000375576200037462000319565b5b80604052505050565b60006200038a620002ea565b905062000398828262000348565b919050565b600067ffffffffffffffff821115620003bb57620003ba62000319565b5b620003c68262000308565b9050602081019050919050565b60005b83811015620003f3578082015181840152602081019050620003d6565b60008484015250505050565b60006200041662000410846200039d565b6200037e565b90508281526020810184848401111562000435576200043462000303565b5b62000442848285620003d3565b509392505050565b600082601f830112620004625762000461620002fe565b5b815162000474848260208601620003ff565b91505092915050565b60008060408385031215620004975762000496620002f4565b5b600083015167ffffffffffffffff811115620004b857620004b7620002f9565b5b620004c6858286016200044a565b925050602083015167ffffffffffffffff811115620004ea57620004e9620002f9565b5b620004f8858286016200044a565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200055557607f821691505b6020821081036200056b576200056a6200050d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005d57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000596565b620005e1868362000596565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200062e620006286200062284620005f9565b62000603565b620005f9565b9050919050565b6000819050919050565b6200064a836200060d565b62000662620006598262000635565b848454620005a3565b825550505050565b600090565b620006796200066a565b620006868184846200063f565b505050565b5b81811015620006ae57620006a26000826200066f565b6001810190506200068c565b5050565b601f821115620006fd57620006c78162000571565b620006d28462000586565b81016020851015620006e2578190505b620006fa620006f18562000586565b8301826200068b565b50505b505050565b600082821c905092915050565b6000620007226000198460080262000702565b1980831691505092915050565b60006200073d83836200070f565b9150826002028217905092915050565b620007588262000502565b67ffffffffffffffff81111562000774576200077362000319565b5b6200078082546200053c565b6200078d828285620006b2565b600060209050601f831160018114620007c55760008415620007b0578287015190505b620007bc85826200072f565b8655506200082c565b601f198416620007d58662000571565b60005b82811015620007ff57848901518255600182019150602085019450602081019050620007d8565b868310156200081f57848901516200081b601f8916826200070f565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200087d601f8362000834565b91506200088a8262000845565b602082019050919050565b60006020820190508181036000830152620008b0816200086e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620008f382620005f9565b91506200090083620005f9565b92508282019050808211156200091b576200091a620008b7565b5b92915050565b6200092c81620005f9565b82525050565b600060208201905062000949600083018462000921565b92915050565b611af9806200095f6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063a457c2d711610066578063a457c2d71461029d578063a9059cbb146102cd578063dd62ed3e146102fd578063f2fde38b1461032d57610100565b8063715018a61461023b57806379cc6790146102455780638da5cb5b1461026157806395d89b411461027f57610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806340c10f19146101ef57806370a082311461020b57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d610349565b60405161011a919061111e565b60405180910390f35b61013d600480360381019061013891906111d9565b6103db565b60405161014a9190611234565b60405180910390f35b61015b6103f9565b604051610168919061125e565b60405180910390f35b61018b60048036038101906101869190611279565b610403565b6040516101989190611234565b60405180910390f35b6101a9610432565b6040516101b691906112e8565b60405180910390f35b6101d960048036038101906101d491906111d9565b61043b565b6040516101e69190611234565b60405180910390f35b610209600480360381019061020491906111d9565b610472565b005b61022560048036038101906102209190611303565b610480565b604051610232919061125e565b60405180910390f35b6102436104c8565b005b61025f600480360381019061025a91906111d9565b6104dc565b005b6102696104ea565b604051610276919061133f565b60405180910390f35b610287610514565b604051610294919061111e565b60405180910390f35b6102b760048036038101906102b291906111d9565b6105a6565b6040516102c49190611234565b60405180910390f35b6102e760048036038101906102e291906111d9565b61061d565b6040516102f49190611234565b60405180910390f35b6103176004803603810190610312919061135a565b610640565b604051610324919061125e565b60405180910390f35b61034760048036038101906103429190611303565b6106c7565b005b606060038054610358906113c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610384906113c9565b80156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b60006103ef6103e861074a565b8484610752565b6001905092915050565b6000600254905090565b60008061040e61074a565b905061041b85828561091b565b6104268585856109a7565b60019150509392505050565b60006012905090565b60008061044661074a565b90506104678185856104588589610640565b6104629190611429565b610752565b600191505092915050565b61047c8282610c1d565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6104d0610d73565b6104da6000610df1565b565b6104e68282610eb7565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610523906113c9565b80601f016020809104026020016040519081016040528092919081815260200182805461054f906113c9565b801561059c5780601f106105715761010080835404028352916020019161059c565b820191906000526020600020905b81548152906001019060200180831161057f57829003601f168201915b5050505050905090565b6000806105b161074a565b905060006105bf8286610640565b905083811015610604576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fb906114cf565b60405180910390fd5b6106118286868403610752565b60019250505092915050565b60008061062861074a565b90506106358185856109a7565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6106cf610d73565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073590611561565b60405180910390fd5b61074781610df1565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b8906115f3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610830576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082790611685565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161090e919061125e565b60405180910390a3505050565b60006109278484610640565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146109a15781811015610993576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098a906116f1565b60405180910390fd5b6109a08484848403610752565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d90611783565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7c90611815565b60405180910390fd5b610a90838383611084565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610b16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0d906118a7565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610c04919061125e565b60405180910390a3610c17848484611089565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8390611913565b60405180910390fd5b610c9860008383611084565b8060026000828254610caa9190611429565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d5b919061125e565b60405180910390a3610d6f60008383611089565b5050565b610d7b61074a565b73ffffffffffffffffffffffffffffffffffffffff16610d996104ea565b73ffffffffffffffffffffffffffffffffffffffff1614610def576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de69061197f565b60405180910390fd5b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1d90611a11565b60405180910390fd5b610f3282600083611084565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610fb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610faf90611aa3565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161106b919061125e565b60405180910390a361107f83600084611089565b505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110c85780820151818401526020810190506110ad565b60008484015250505050565b6000601f19601f8301169050919050565b60006110f08261108e565b6110fa8185611099565b935061110a8185602086016110aa565b611113816110d4565b840191505092915050565b6000602082019050818103600083015261113881846110e5565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061117082611145565b9050919050565b61118081611165565b811461118b57600080fd5b50565b60008135905061119d81611177565b92915050565b6000819050919050565b6111b6816111a3565b81146111c157600080fd5b50565b6000813590506111d3816111ad565b92915050565b600080604083850312156111f0576111ef611140565b5b60006111fe8582860161118e565b925050602061120f858286016111c4565b9150509250929050565b60008115159050919050565b61122e81611219565b82525050565b60006020820190506112496000830184611225565b92915050565b611258816111a3565b82525050565b6000602082019050611273600083018461124f565b92915050565b60008060006060848603121561129257611291611140565b5b60006112a08682870161118e565b93505060206112b18682870161118e565b92505060406112c2868287016111c4565b9150509250925092565b600060ff82169050919050565b6112e2816112cc565b82525050565b60006020820190506112fd60008301846112d9565b92915050565b60006020828403121561131957611318611140565b5b60006113278482850161118e565b91505092915050565b61133981611165565b82525050565b60006020820190506113546000830184611330565b92915050565b6000806040838503121561137157611370611140565b5b600061137f8582860161118e565b92505060206113908582860161118e565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113e157607f821691505b6020821081036113f4576113f361139a565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611434826111a3565b915061143f836111a3565b9250828201905080821115611457576114566113fa565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006114b9602583611099565b91506114c48261145d565b604082019050919050565b600060208201905081810360008301526114e8816114ac565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061154b602683611099565b9150611556826114ef565b604082019050919050565b6000602082019050818103600083015261157a8161153e565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006115dd602483611099565b91506115e882611581565b604082019050919050565b6000602082019050818103600083015261160c816115d0565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061166f602283611099565b915061167a82611613565b604082019050919050565b6000602082019050818103600083015261169e81611662565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b60006116db601d83611099565b91506116e6826116a5565b602082019050919050565b6000602082019050818103600083015261170a816116ce565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061176d602583611099565b915061177882611711565b604082019050919050565b6000602082019050818103600083015261179c81611760565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006117ff602383611099565b915061180a826117a3565b604082019050919050565b6000602082019050818103600083015261182e816117f2565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611891602683611099565b915061189c82611835565b604082019050919050565b600060208201905081810360008301526118c081611884565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006118fd601f83611099565b9150611908826118c7565b602082019050919050565b6000602082019050818103600083015261192c816118f0565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611969602083611099565b915061197482611933565b602082019050919050565b600060208201905081810360008301526119988161195c565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b60006119fb602183611099565b9150611a068261199f565b604082019050919050565b60006020820190508181036000830152611a2a816119ee565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611a8d602283611099565b9150611a9882611a31565b604082019050919050565b60006020820190508181036000830152611abc81611a80565b905091905056fea26469706673582212200a7fd52ed1dce0061a949a3ed1bed4c6f585c1fdb130d0a491e3f20381a02d0d64736f6c63430008120033",
}

// L1CustomERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use L1CustomERC20MetaData.ABI instead.
var L1CustomERC20ABI = L1CustomERC20MetaData.ABI

// L1CustomERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1CustomERC20MetaData.Bin instead.
var L1CustomERC20Bin = L1CustomERC20MetaData.Bin

// DeployL1CustomERC20 deploys a new Ethereum contract, binding an instance of L1CustomERC20 to it.
func DeployL1CustomERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *L1CustomERC20, error) {
	parsed, err := L1CustomERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1CustomERC20Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1CustomERC20{L1CustomERC20Caller: L1CustomERC20Caller{contract: contract}, L1CustomERC20Transactor: L1CustomERC20Transactor{contract: contract}, L1CustomERC20Filterer: L1CustomERC20Filterer{contract: contract}}, nil
}

// L1CustomERC20 is an auto generated Go binding around an Ethereum contract.
type L1CustomERC20 struct {
	L1CustomERC20Caller     // Read-only binding to the contract
	L1CustomERC20Transactor // Write-only binding to the contract
	L1CustomERC20Filterer   // Log filterer for contract events
}

// L1CustomERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type L1CustomERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L1CustomERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1CustomERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1CustomERC20Session struct {
	Contract     *L1CustomERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1CustomERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1CustomERC20CallerSession struct {
	Contract *L1CustomERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1CustomERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1CustomERC20TransactorSession struct {
	Contract     *L1CustomERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1CustomERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type L1CustomERC20Raw struct {
	Contract *L1CustomERC20 // Generic contract binding to access the raw methods on
}

// L1CustomERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1CustomERC20CallerRaw struct {
	Contract *L1CustomERC20Caller // Generic read-only contract binding to access the raw methods on
}

// L1CustomERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1CustomERC20TransactorRaw struct {
	Contract *L1CustomERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL1CustomERC20 creates a new instance of L1CustomERC20, bound to a specific deployed contract.
func NewL1CustomERC20(address common.Address, backend bind.ContractBackend) (*L1CustomERC20, error) {
	contract, err := bindL1CustomERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20{L1CustomERC20Caller: L1CustomERC20Caller{contract: contract}, L1CustomERC20Transactor: L1CustomERC20Transactor{contract: contract}, L1CustomERC20Filterer: L1CustomERC20Filterer{contract: contract}}, nil
}

// NewL1CustomERC20Caller creates a new read-only instance of L1CustomERC20, bound to a specific deployed contract.
func NewL1CustomERC20Caller(address common.Address, caller bind.ContractCaller) (*L1CustomERC20Caller, error) {
	contract, err := bindL1CustomERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20Caller{contract: contract}, nil
}

// NewL1CustomERC20Transactor creates a new write-only instance of L1CustomERC20, bound to a specific deployed contract.
func NewL1CustomERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*L1CustomERC20Transactor, error) {
	contract, err := bindL1CustomERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20Transactor{contract: contract}, nil
}

// NewL1CustomERC20Filterer creates a new log filterer instance of L1CustomERC20, bound to a specific deployed contract.
func NewL1CustomERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*L1CustomERC20Filterer, error) {
	contract, err := bindL1CustomERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20Filterer{contract: contract}, nil
}

// bindL1CustomERC20 binds a generic wrapper to an already deployed contract.
func bindL1CustomERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1CustomERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1CustomERC20 *L1CustomERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1CustomERC20.Contract.L1CustomERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1CustomERC20 *L1CustomERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.L1CustomERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1CustomERC20 *L1CustomERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.L1CustomERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1CustomERC20 *L1CustomERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1CustomERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1CustomERC20 *L1CustomERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1CustomERC20 *L1CustomERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L1CustomERC20.Contract.Allowance(&_L1CustomERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L1CustomERC20.Contract.Allowance(&_L1CustomERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _L1CustomERC20.Contract.BalanceOf(&_L1CustomERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _L1CustomERC20.Contract.BalanceOf(&_L1CustomERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1CustomERC20 *L1CustomERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1CustomERC20 *L1CustomERC20Session) Decimals() (uint8, error) {
	return _L1CustomERC20.Contract.Decimals(&_L1CustomERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1CustomERC20 *L1CustomERC20CallerSession) Decimals() (uint8, error) {
	return _L1CustomERC20.Contract.Decimals(&_L1CustomERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1CustomERC20 *L1CustomERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1CustomERC20 *L1CustomERC20Session) Name() (string, error) {
	return _L1CustomERC20.Contract.Name(&_L1CustomERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1CustomERC20 *L1CustomERC20CallerSession) Name() (string, error) {
	return _L1CustomERC20.Contract.Name(&_L1CustomERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomERC20 *L1CustomERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomERC20 *L1CustomERC20Session) Owner() (common.Address, error) {
	return _L1CustomERC20.Contract.Owner(&_L1CustomERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomERC20 *L1CustomERC20CallerSession) Owner() (common.Address, error) {
	return _L1CustomERC20.Contract.Owner(&_L1CustomERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1CustomERC20 *L1CustomERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1CustomERC20 *L1CustomERC20Session) Symbol() (string, error) {
	return _L1CustomERC20.Contract.Symbol(&_L1CustomERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1CustomERC20 *L1CustomERC20CallerSession) Symbol() (string, error) {
	return _L1CustomERC20.Contract.Symbol(&_L1CustomERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1CustomERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20Session) TotalSupply() (*big.Int, error) {
	return _L1CustomERC20.Contract.TotalSupply(&_L1CustomERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1CustomERC20 *L1CustomERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _L1CustomERC20.Contract.TotalSupply(&_L1CustomERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Approve(&_L1CustomERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Approve(&_L1CustomERC20.TransactOpts, spender, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.BurnFrom(&_L1CustomERC20.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.BurnFrom(&_L1CustomERC20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.DecreaseAllowance(&_L1CustomERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.DecreaseAllowance(&_L1CustomERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.IncreaseAllowance(&_L1CustomERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1CustomERC20 *L1CustomERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.IncreaseAllowance(&_L1CustomERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Mint(&_L1CustomERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_L1CustomERC20 *L1CustomERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Mint(&_L1CustomERC20.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1CustomERC20 *L1CustomERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1CustomERC20 *L1CustomERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _L1CustomERC20.Contract.RenounceOwnership(&_L1CustomERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1CustomERC20 *L1CustomERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1CustomERC20.Contract.RenounceOwnership(&_L1CustomERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Transfer(&_L1CustomERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.Transfer(&_L1CustomERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.TransferFrom(&_L1CustomERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L1CustomERC20 *L1CustomERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.TransferFrom(&_L1CustomERC20.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1CustomERC20 *L1CustomERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1CustomERC20 *L1CustomERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.TransferOwnership(&_L1CustomERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1CustomERC20 *L1CustomERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomERC20.Contract.TransferOwnership(&_L1CustomERC20.TransactOpts, newOwner)
}

// L1CustomERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the L1CustomERC20 contract.
type L1CustomERC20ApprovalIterator struct {
	Event *L1CustomERC20Approval // Event containing the contract specifics and raw log

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
func (it *L1CustomERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomERC20Approval)
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
		it.Event = new(L1CustomERC20Approval)
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
func (it *L1CustomERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomERC20Approval represents a Approval event raised by the L1CustomERC20 contract.
type L1CustomERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*L1CustomERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L1CustomERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20ApprovalIterator{contract: _L1CustomERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *L1CustomERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L1CustomERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomERC20Approval)
				if err := _L1CustomERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) ParseApproval(log types.Log) (*L1CustomERC20Approval, error) {
	event := new(L1CustomERC20Approval)
	if err := _L1CustomERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1CustomERC20 contract.
type L1CustomERC20OwnershipTransferredIterator struct {
	Event *L1CustomERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1CustomERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomERC20OwnershipTransferred)
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
		it.Event = new(L1CustomERC20OwnershipTransferred)
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
func (it *L1CustomERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomERC20OwnershipTransferred represents a OwnershipTransferred event raised by the L1CustomERC20 contract.
type L1CustomERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1CustomERC20 *L1CustomERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1CustomERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1CustomERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20OwnershipTransferredIterator{contract: _L1CustomERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1CustomERC20 *L1CustomERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1CustomERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1CustomERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomERC20OwnershipTransferred)
				if err := _L1CustomERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1CustomERC20 *L1CustomERC20Filterer) ParseOwnershipTransferred(log types.Log) (*L1CustomERC20OwnershipTransferred, error) {
	event := new(L1CustomERC20OwnershipTransferred)
	if err := _L1CustomERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L1CustomERC20 contract.
type L1CustomERC20TransferIterator struct {
	Event *L1CustomERC20Transfer // Event containing the contract specifics and raw log

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
func (it *L1CustomERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomERC20Transfer)
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
		it.Event = new(L1CustomERC20Transfer)
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
func (it *L1CustomERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomERC20Transfer represents a Transfer event raised by the L1CustomERC20 contract.
type L1CustomERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1CustomERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1CustomERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomERC20TransferIterator{contract: _L1CustomERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L1CustomERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1CustomERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomERC20Transfer)
				if err := _L1CustomERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1CustomERC20 *L1CustomERC20Filterer) ParseTransfer(log types.Log) (*L1CustomERC20Transfer, error) {
	event := new(L1CustomERC20Transfer)
	if err := _L1CustomERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
