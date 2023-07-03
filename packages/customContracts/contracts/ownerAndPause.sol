// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract POContract is Pausable, Ownable {
    uint256 public counter;
    string private _name;

    event CounterIncremented(address indexed user, uint256 newValue);

    constructor(string memory name) {
        _name = name;
    }

    function getName() public view returns (string memory) {
        return _name;
    }

    function incrementCounter() public whenNotPaused {
        counter++;
        emit CounterIncremented(msg.sender, counter);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function withdrawEther() external onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    receive() external payable {}
}