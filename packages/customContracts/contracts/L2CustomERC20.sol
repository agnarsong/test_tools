// SPDX-License-Identifier: MIT
pragma solidity >=0.5.16 <0.9.0;

import { L2StandardERC20 } from "@mantleio/contracts/standards/L2StandardERC20.sol";

contract L2CustomERC20 is L2StandardERC20 {
  constructor(
    address _l2Bridge,
    address _l1Token
  )
  
  L2StandardERC20(_l2Bridge, _l1Token, "L2CustomERC20", "L2E", 18){}

  function decimals() public view override returns (uint8) {
      return 18;
  }
  
  function transfers(address _token,address[] calldata to) public payable {
    L2CustomERC20 token = L2CustomERC20(_token);
    for (uint256 i = 0; i < to.length; i++) {
        token.transfer(to[i], 1);
    }
  }
}