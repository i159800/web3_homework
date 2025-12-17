// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Counter {
    uint256 public count;

    function get() public view returns (uint256) {
        return count;
    }

    function increment() public {
        count += 1;
    }

    function decrement() public {
        require( count > 0, "Count is already zero");
        count -= 1;
    }

    function set(uint256 value) public {
        count = value;
    }
}