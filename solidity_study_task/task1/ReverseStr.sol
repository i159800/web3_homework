// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract ReverseStr {
    string public str;
    bytes public bs1 = new bytes(0);
    bytes public bs2 = new bytes(0);
    string public str1;
    function SetBytes(string memory str2) public {
        str = str2;
        bs1 = bytes(str2);
    }

    function reverse() public {
        for (uint i = bs1.length; i >= 1; i--) {
            bs2.push(bs1[i-1]);
        }
        str1 = string(bs2);
    }
}