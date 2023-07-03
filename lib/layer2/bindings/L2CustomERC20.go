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

// L2CustomERC20MetaData contains all meta data concerning the L2CustomERC20 contract.
var L2CustomERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"}],\"name\":\"transfers\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200259c3803806200259c8339818101604052810190620000379190620001e3565b81816040518060400160405280600d81526020017f4c32437573746f6d4552433230000000000000000000000000000000000000008152506040518060400160405280600381526020017f4c32450000000000000000000000000000000000000000000000000000000000815250601282828160039081620000ba9190620004a4565b508060049081620000cc9190620004a4565b50505083600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660146101000a81548160ff021916908360ff160217905550505050505050506200058b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001ab826200017e565b9050919050565b620001bd816200019e565b8114620001c957600080fd5b50565b600081519050620001dd81620001b2565b92915050565b60008060408385031215620001fd57620001fc62000179565b5b60006200020d85828601620001cc565b92505060206200022085828601620001cc565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ac57607f821691505b602082108103620002c257620002c162000264565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200032c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002ed565b620003388683620002ed565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003856200037f620003798462000350565b6200035a565b62000350565b9050919050565b6000819050919050565b620003a18362000364565b620003b9620003b0826200038c565b848454620002fa565b825550505050565b600090565b620003d0620003c1565b620003dd81848462000396565b505050565b5b818110156200040557620003f9600082620003c6565b600181019050620003e3565b5050565b601f82111562000454576200041e81620002c8565b6200042984620002dd565b8101602085101562000439578190505b620004516200044885620002dd565b830182620003e2565b50505b505050565b600082821c905092915050565b6000620004796000198460080262000459565b1980831691505092915050565b600062000494838362000466565b9150826002028217905092915050565b620004af826200022a565b67ffffffffffffffff811115620004cb57620004ca62000235565b5b620004d7825462000293565b620004e482828562000409565b600060209050601f8311600181146200051c576000841562000507578287015190505b62000513858262000486565b86555062000583565b601f1984166200052c86620002c8565b60005b8281101562000556578489015182556001820191506020850194506020810190506200052f565b8683101562000576578489015162000572601f89168262000466565b8355505b6001600288020188555050505b505050505050565b612001806200059b6000396000f3fe6080604052600436106101095760003560e01c806376809ce311610095578063a457c2d711610064578063a457c2d714610384578063a9059cbb146103c1578063ae1f6aaf146103fe578063c01e1bd614610429578063dd62ed3e1461045457610109565b806376809ce3146102e9578063904b56401461031457806395d89b41146103305780639dc29fac1461035b57610109565b806323b872dd116100dc57806323b872dd146101de578063313ce5671461021b578063395093511461024657806340c10f191461028357806370a08231146102ac57610109565b806301ffc9a71461010e57806306fdde031461014b578063095ea7b31461017657806318160ddd146101b3575b600080fd5b34801561011a57600080fd5b50610135600480360381019061013091906113fd565b610491565b6040516101429190611445565b60405180910390f35b34801561015757600080fd5b50610160610567565b60405161016d91906114f0565b60405180910390f35b34801561018257600080fd5b5061019d600480360381019061019891906115a6565b6105f9565b6040516101aa9190611445565b60405180910390f35b3480156101bf57600080fd5b506101c861061c565b6040516101d591906115f5565b60405180910390f35b3480156101ea57600080fd5b5061020560048036038101906102009190611610565b610626565b6040516102129190611445565b60405180910390f35b34801561022757600080fd5b50610230610655565b60405161023d919061167f565b60405180910390f35b34801561025257600080fd5b5061026d600480360381019061026891906115a6565b61065e565b60405161027a9190611445565b60405180910390f35b34801561028f57600080fd5b506102aa60048036038101906102a591906115a6565b610695565b005b3480156102b857600080fd5b506102d360048036038101906102ce919061169a565b610781565b6040516102e091906115f5565b60405180910390f35b3480156102f557600080fd5b506102fe6107c9565b60405161030b919061167f565b60405180910390f35b61032e6004803603810190610329919061172c565b6107dc565b005b34801561033c57600080fd5b506103456108b0565b60405161035291906114f0565b60405180910390f35b34801561036757600080fd5b50610382600480360381019061037d91906115a6565b610942565b005b34801561039057600080fd5b506103ab60048036038101906103a691906115a6565b610a2e565b6040516103b89190611445565b60405180910390f35b3480156103cd57600080fd5b506103e860048036038101906103e391906115a6565b610aa5565b6040516103f59190611445565b60405180910390f35b34801561040a57600080fd5b50610413610ac8565b604051610420919061179b565b60405180910390f35b34801561043557600080fd5b5061043e610aee565b60405161044b919061179b565b60405180910390f35b34801561046057600080fd5b5061047b600480360381019061047691906117b6565b610b14565b60405161048891906115f5565b60405180910390f35b6000807f01ffc9a7a5cef8baa21ed3c5c0d7e23accb804b619e9333b597f47a0d84076e290506000639dc29fac60e01b6340c10f1960e01b63c01e1bd660e01b18189050817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061055e5750807bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b92505050919050565b60606003805461057690611825565b80601f01602080910402602001604051908101604052809291908181526020018280546105a290611825565b80156105ef5780601f106105c4576101008083540402835291602001916105ef565b820191906000526020600020905b8154815290600101906020018083116105d257829003601f168201915b5050505050905090565b600080610604610b9b565b9050610611818585610ba3565b600191505092915050565b6000600254905090565b600080610631610b9b565b905061063e858285610d6c565b610649858585610df8565b60019150509392505050565b60006012905090565b600080610669610b9b565b905061068a81858561067b8589610b14565b6106859190611885565b610ba3565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071c90611905565b60405180910390fd5b61072f828261106e565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161077591906115f5565b60405180910390a25050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600660149054906101000a900460ff1681565b600083905060005b838390508110156108a9578173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85858481811061081e5761081d611925565b5b9050602002016020810190610833919061169a565b60016040518363ffffffff1660e01b8152600401610852929190611999565b6020604051808303816000875af1158015610871573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089591906119ee565b5080806108a190611a1b565b9150506107e4565b5050505050565b6060600480546108bf90611825565b80601f01602080910402602001604051908101604052809291908181526020018280546108eb90611825565b80156109385780601f1061090d57610100808354040283529160200191610938565b820191906000526020600020905b81548152906001019060200180831161091b57829003601f168201915b5050505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c990611905565b60405180910390fd5b6109dc82826111c4565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610a2291906115f5565b60405180910390a25050565b600080610a39610b9b565b90506000610a478286610b14565b905083811015610a8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8390611ad5565b60405180910390fd5b610a998286868403610ba3565b60019250505092915050565b600080610ab0610b9b565b9050610abd818585610df8565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0990611b67565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7890611bf9565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610d5f91906115f5565b60405180910390a3505050565b6000610d788484610b14565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df25781811015610de4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ddb90611c65565b60405180910390fd5b610df18484848403610ba3565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610e67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5e90611cf7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ed6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ecd90611d89565b60405180910390fd5b610ee1838383611391565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610f67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f5e90611e1b565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161105591906115f5565b60405180910390a3611068848484611396565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036110dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d490611e87565b60405180910390fd5b6110e960008383611391565b80600260008282546110fb9190611885565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111ac91906115f5565b60405180910390a36111c060008383611396565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122a90611f19565b60405180910390fd5b61123f82600083611391565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156112c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112bc90611fab565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161137891906115f5565b60405180910390a361138c83600084611396565b505050565b505050565b505050565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113da816113a5565b81146113e557600080fd5b50565b6000813590506113f7816113d1565b92915050565b6000602082840312156114135761141261139b565b5b6000611421848285016113e8565b91505092915050565b60008115159050919050565b61143f8161142a565b82525050565b600060208201905061145a6000830184611436565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561149a57808201518184015260208101905061147f565b60008484015250505050565b6000601f19601f8301169050919050565b60006114c282611460565b6114cc818561146b565b93506114dc81856020860161147c565b6114e5816114a6565b840191505092915050565b6000602082019050818103600083015261150a81846114b7565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061153d82611512565b9050919050565b61154d81611532565b811461155857600080fd5b50565b60008135905061156a81611544565b92915050565b6000819050919050565b61158381611570565b811461158e57600080fd5b50565b6000813590506115a08161157a565b92915050565b600080604083850312156115bd576115bc61139b565b5b60006115cb8582860161155b565b92505060206115dc85828601611591565b9150509250929050565b6115ef81611570565b82525050565b600060208201905061160a60008301846115e6565b92915050565b6000806000606084860312156116295761162861139b565b5b60006116378682870161155b565b93505060206116488682870161155b565b925050604061165986828701611591565b9150509250925092565b600060ff82169050919050565b61167981611663565b82525050565b60006020820190506116946000830184611670565b92915050565b6000602082840312156116b0576116af61139b565b5b60006116be8482850161155b565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126116ec576116eb6116c7565b5b8235905067ffffffffffffffff811115611709576117086116cc565b5b602083019150836020820283011115611725576117246116d1565b5b9250929050565b6000806000604084860312156117455761174461139b565b5b60006117538682870161155b565b935050602084013567ffffffffffffffff811115611774576117736113a0565b5b611780868287016116d6565b92509250509250925092565b61179581611532565b82525050565b60006020820190506117b0600083018461178c565b92915050565b600080604083850312156117cd576117cc61139b565b5b60006117db8582860161155b565b92505060206117ec8582860161155b565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061183d57607f821691505b6020821081036118505761184f6117f6565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061189082611570565b915061189b83611570565b92508282019050808211156118b3576118b2611856565b5b92915050565b7f4f6e6c79204c32204272696467652063616e206d696e7420616e64206275726e600082015250565b60006118ef60208361146b565b91506118fa826118b9565b602082019050919050565b6000602082019050818103600083015261191e816118e2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000819050919050565b6000819050919050565b600061198361197e61197984611954565b61195e565b611570565b9050919050565b61199381611968565b82525050565b60006040820190506119ae600083018561178c565b6119bb602083018461198a565b9392505050565b6119cb8161142a565b81146119d657600080fd5b50565b6000815190506119e8816119c2565b92915050565b600060208284031215611a0457611a0361139b565b5b6000611a12848285016119d9565b91505092915050565b6000611a2682611570565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a5857611a57611856565b5b600182019050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611abf60258361146b565b9150611aca82611a63565b604082019050919050565b60006020820190508181036000830152611aee81611ab2565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611b5160248361146b565b9150611b5c82611af5565b604082019050919050565b60006020820190508181036000830152611b8081611b44565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611be360228361146b565b9150611bee82611b87565b604082019050919050565b60006020820190508181036000830152611c1281611bd6565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611c4f601d8361146b565b9150611c5a82611c19565b602082019050919050565b60006020820190508181036000830152611c7e81611c42565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611ce160258361146b565b9150611cec82611c85565b604082019050919050565b60006020820190508181036000830152611d1081611cd4565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611d7360238361146b565b9150611d7e82611d17565b604082019050919050565b60006020820190508181036000830152611da281611d66565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611e0560268361146b565b9150611e1082611da9565b604082019050919050565b60006020820190508181036000830152611e3481611df8565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611e71601f8361146b565b9150611e7c82611e3b565b602082019050919050565b60006020820190508181036000830152611ea081611e64565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611f0360218361146b565b9150611f0e82611ea7565b604082019050919050565b60006020820190508181036000830152611f3281611ef6565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611f9560228361146b565b9150611fa082611f39565b604082019050919050565b60006020820190508181036000830152611fc481611f88565b905091905056fea26469706673582212200acb024cd6b58e13f46c09e5aec239a51ace4ff71fb4a9197b6e369ddd1bef8464736f6c63430008120033",
}

// L2CustomERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use L2CustomERC20MetaData.ABI instead.
var L2CustomERC20ABI = L2CustomERC20MetaData.ABI

// L2CustomERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2CustomERC20MetaData.Bin instead.
var L2CustomERC20Bin = L2CustomERC20MetaData.Bin

// DeployL2CustomERC20 deploys a new Ethereum contract, binding an instance of L2CustomERC20 to it.
func DeployL2CustomERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _l2Bridge common.Address, _l1Token common.Address) (common.Address, *types.Transaction, *L2CustomERC20, error) {
	parsed, err := L2CustomERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2CustomERC20Bin), backend, _l2Bridge, _l1Token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2CustomERC20{L2CustomERC20Caller: L2CustomERC20Caller{contract: contract}, L2CustomERC20Transactor: L2CustomERC20Transactor{contract: contract}, L2CustomERC20Filterer: L2CustomERC20Filterer{contract: contract}}, nil
}

// L2CustomERC20 is an auto generated Go binding around an Ethereum contract.
type L2CustomERC20 struct {
	L2CustomERC20Caller     // Read-only binding to the contract
	L2CustomERC20Transactor // Write-only binding to the contract
	L2CustomERC20Filterer   // Log filterer for contract events
}

// L2CustomERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type L2CustomERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CustomERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L2CustomERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CustomERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2CustomERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CustomERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2CustomERC20Session struct {
	Contract     *L2CustomERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2CustomERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CustomERC20CallerSession struct {
	Contract *L2CustomERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2CustomERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2CustomERC20TransactorSession struct {
	Contract     *L2CustomERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2CustomERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type L2CustomERC20Raw struct {
	Contract *L2CustomERC20 // Generic contract binding to access the raw methods on
}

// L2CustomERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CustomERC20CallerRaw struct {
	Contract *L2CustomERC20Caller // Generic read-only contract binding to access the raw methods on
}

// L2CustomERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2CustomERC20TransactorRaw struct {
	Contract *L2CustomERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL2CustomERC20 creates a new instance of L2CustomERC20, bound to a specific deployed contract.
func NewL2CustomERC20(address common.Address, backend bind.ContractBackend) (*L2CustomERC20, error) {
	contract, err := bindL2CustomERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20{L2CustomERC20Caller: L2CustomERC20Caller{contract: contract}, L2CustomERC20Transactor: L2CustomERC20Transactor{contract: contract}, L2CustomERC20Filterer: L2CustomERC20Filterer{contract: contract}}, nil
}

// NewL2CustomERC20Caller creates a new read-only instance of L2CustomERC20, bound to a specific deployed contract.
func NewL2CustomERC20Caller(address common.Address, caller bind.ContractCaller) (*L2CustomERC20Caller, error) {
	contract, err := bindL2CustomERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20Caller{contract: contract}, nil
}

// NewL2CustomERC20Transactor creates a new write-only instance of L2CustomERC20, bound to a specific deployed contract.
func NewL2CustomERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*L2CustomERC20Transactor, error) {
	contract, err := bindL2CustomERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20Transactor{contract: contract}, nil
}

// NewL2CustomERC20Filterer creates a new log filterer instance of L2CustomERC20, bound to a specific deployed contract.
func NewL2CustomERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*L2CustomERC20Filterer, error) {
	contract, err := bindL2CustomERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20Filterer{contract: contract}, nil
}

// bindL2CustomERC20 binds a generic wrapper to an already deployed contract.
func bindL2CustomERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2CustomERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CustomERC20 *L2CustomERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CustomERC20.Contract.L2CustomERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CustomERC20 *L2CustomERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.L2CustomERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CustomERC20 *L2CustomERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.L2CustomERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CustomERC20 *L2CustomERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CustomERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CustomERC20 *L2CustomERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CustomERC20 *L2CustomERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L2CustomERC20.Contract.Allowance(&_L2CustomERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L2CustomERC20.Contract.Allowance(&_L2CustomERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _L2CustomERC20.Contract.BalanceOf(&_L2CustomERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _L2CustomERC20.Contract.BalanceOf(&_L2CustomERC20.CallOpts, account)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20Caller) Decimal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "decimal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20Session) Decimal() (uint8, error) {
	return _L2CustomERC20.Contract.Decimal(&_L2CustomERC20.CallOpts)
}

// Decimal is a free data retrieval call binding the contract method 0x76809ce3.
//
// Solidity: function decimal() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20CallerSession) Decimal() (uint8, error) {
	return _L2CustomERC20.Contract.Decimal(&_L2CustomERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20Session) Decimals() (uint8, error) {
	return _L2CustomERC20.Contract.Decimals(&_L2CustomERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2CustomERC20 *L2CustomERC20CallerSession) Decimals() (uint8, error) {
	return _L2CustomERC20.Contract.Decimals(&_L2CustomERC20.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2CustomERC20 *L2CustomERC20Caller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2CustomERC20 *L2CustomERC20Session) L1Token() (common.Address, error) {
	return _L2CustomERC20.Contract.L1Token(&_L2CustomERC20.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2CustomERC20 *L2CustomERC20CallerSession) L1Token() (common.Address, error) {
	return _L2CustomERC20.Contract.L1Token(&_L2CustomERC20.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2CustomERC20 *L2CustomERC20Caller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2CustomERC20 *L2CustomERC20Session) L2Bridge() (common.Address, error) {
	return _L2CustomERC20.Contract.L2Bridge(&_L2CustomERC20.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2CustomERC20 *L2CustomERC20CallerSession) L2Bridge() (common.Address, error) {
	return _L2CustomERC20.Contract.L2Bridge(&_L2CustomERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2CustomERC20 *L2CustomERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2CustomERC20 *L2CustomERC20Session) Name() (string, error) {
	return _L2CustomERC20.Contract.Name(&_L2CustomERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2CustomERC20 *L2CustomERC20CallerSession) Name() (string, error) {
	return _L2CustomERC20.Contract.Name(&_L2CustomERC20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2CustomERC20 *L2CustomERC20Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _L2CustomERC20.Contract.SupportsInterface(&_L2CustomERC20.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2CustomERC20 *L2CustomERC20CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _L2CustomERC20.Contract.SupportsInterface(&_L2CustomERC20.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2CustomERC20 *L2CustomERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2CustomERC20 *L2CustomERC20Session) Symbol() (string, error) {
	return _L2CustomERC20.Contract.Symbol(&_L2CustomERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2CustomERC20 *L2CustomERC20CallerSession) Symbol() (string, error) {
	return _L2CustomERC20.Contract.Symbol(&_L2CustomERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2CustomERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20Session) TotalSupply() (*big.Int, error) {
	return _L2CustomERC20.Contract.TotalSupply(&_L2CustomERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2CustomERC20 *L2CustomERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _L2CustomERC20.Contract.TotalSupply(&_L2CustomERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Approve(&_L2CustomERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Approve(&_L2CustomERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Burn(&_L2CustomERC20.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Burn(&_L2CustomERC20.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.DecreaseAllowance(&_L2CustomERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.DecreaseAllowance(&_L2CustomERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.IncreaseAllowance(&_L2CustomERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2CustomERC20 *L2CustomERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.IncreaseAllowance(&_L2CustomERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Mint(&_L2CustomERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2CustomERC20 *L2CustomERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Mint(&_L2CustomERC20.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Transfer(&_L2CustomERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Transfer(&_L2CustomERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.TransferFrom(&_L2CustomERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_L2CustomERC20 *L2CustomERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.TransferFrom(&_L2CustomERC20.TransactOpts, from, to, amount)
}

// Transfers is a paid mutator transaction binding the contract method 0x904b5640.
//
// Solidity: function transfers(address _token, address[] to) payable returns()
func (_L2CustomERC20 *L2CustomERC20Transactor) Transfers(opts *bind.TransactOpts, _token common.Address, to []common.Address) (*types.Transaction, error) {
	return _L2CustomERC20.contract.Transact(opts, "transfers", _token, to)
}

// Transfers is a paid mutator transaction binding the contract method 0x904b5640.
//
// Solidity: function transfers(address _token, address[] to) payable returns()
func (_L2CustomERC20 *L2CustomERC20Session) Transfers(_token common.Address, to []common.Address) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Transfers(&_L2CustomERC20.TransactOpts, _token, to)
}

// Transfers is a paid mutator transaction binding the contract method 0x904b5640.
//
// Solidity: function transfers(address _token, address[] to) payable returns()
func (_L2CustomERC20 *L2CustomERC20TransactorSession) Transfers(_token common.Address, to []common.Address) (*types.Transaction, error) {
	return _L2CustomERC20.Contract.Transfers(&_L2CustomERC20.TransactOpts, _token, to)
}

// L2CustomERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the L2CustomERC20 contract.
type L2CustomERC20ApprovalIterator struct {
	Event *L2CustomERC20Approval // Event containing the contract specifics and raw log

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
func (it *L2CustomERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CustomERC20Approval)
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
		it.Event = new(L2CustomERC20Approval)
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
func (it *L2CustomERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CustomERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CustomERC20Approval represents a Approval event raised by the L2CustomERC20 contract.
type L2CustomERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L2CustomERC20 *L2CustomERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*L2CustomERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2CustomERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20ApprovalIterator{contract: _L2CustomERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L2CustomERC20 *L2CustomERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *L2CustomERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2CustomERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CustomERC20Approval)
				if err := _L2CustomERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_L2CustomERC20 *L2CustomERC20Filterer) ParseApproval(log types.Log) (*L2CustomERC20Approval, error) {
	event := new(L2CustomERC20Approval)
	if err := _L2CustomERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CustomERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the L2CustomERC20 contract.
type L2CustomERC20BurnIterator struct {
	Event *L2CustomERC20Burn // Event containing the contract specifics and raw log

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
func (it *L2CustomERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CustomERC20Burn)
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
		it.Event = new(L2CustomERC20Burn)
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
func (it *L2CustomERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CustomERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CustomERC20Burn represents a Burn event raised by the L2CustomERC20 contract.
type L2CustomERC20Burn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) FilterBurn(opts *bind.FilterOpts, _account []common.Address) (*L2CustomERC20BurnIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2CustomERC20.contract.FilterLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20BurnIterator{contract: _L2CustomERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *L2CustomERC20Burn, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2CustomERC20.contract.WatchLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CustomERC20Burn)
				if err := _L2CustomERC20.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) ParseBurn(log types.Log) (*L2CustomERC20Burn, error) {
	event := new(L2CustomERC20Burn)
	if err := _L2CustomERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CustomERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the L2CustomERC20 contract.
type L2CustomERC20MintIterator struct {
	Event *L2CustomERC20Mint // Event containing the contract specifics and raw log

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
func (it *L2CustomERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CustomERC20Mint)
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
		it.Event = new(L2CustomERC20Mint)
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
func (it *L2CustomERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CustomERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CustomERC20Mint represents a Mint event raised by the L2CustomERC20 contract.
type L2CustomERC20Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) FilterMint(opts *bind.FilterOpts, _account []common.Address) (*L2CustomERC20MintIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2CustomERC20.contract.FilterLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20MintIterator{contract: _L2CustomERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *L2CustomERC20Mint, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2CustomERC20.contract.WatchLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CustomERC20Mint)
				if err := _L2CustomERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2CustomERC20 *L2CustomERC20Filterer) ParseMint(log types.Log) (*L2CustomERC20Mint, error) {
	event := new(L2CustomERC20Mint)
	if err := _L2CustomERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CustomERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L2CustomERC20 contract.
type L2CustomERC20TransferIterator struct {
	Event *L2CustomERC20Transfer // Event containing the contract specifics and raw log

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
func (it *L2CustomERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CustomERC20Transfer)
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
		it.Event = new(L2CustomERC20Transfer)
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
func (it *L2CustomERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CustomERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CustomERC20Transfer represents a Transfer event raised by the L2CustomERC20 contract.
type L2CustomERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2CustomERC20 *L2CustomERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2CustomERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2CustomERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2CustomERC20TransferIterator{contract: _L2CustomERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2CustomERC20 *L2CustomERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L2CustomERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2CustomERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CustomERC20Transfer)
				if err := _L2CustomERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_L2CustomERC20 *L2CustomERC20Filterer) ParseTransfer(log types.Log) (*L2CustomERC20Transfer, error) {
	event := new(L2CustomERC20Transfer)
	if err := _L2CustomERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
