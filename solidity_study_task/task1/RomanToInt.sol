// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract RomanToInt {
    mapping(bytes1 => uint) public valueMap;
    function initMap() private {
        valueMap['I'] = 1;
        valueMap['V'] = 5;
        valueMap['X'] = 10;
        valueMap['L'] = 50;
        valueMap['C'] = 100;
        valueMap['D'] = 500;
        valueMap['M'] = 1000;
    }

    constructor() {
        initMap();
    }
    
    function romanToInt(string memory s) public view returns (uint) {
        uint result = 0;
        for (uint i = 0; i < bytes(s).length; i++) {
            uint currentValue = getValue(bytes(s)[i]);
            if (i < bytes(s).length - 1 && currentValue < getValue(bytes(s)[i+1])) {
                result -= currentValue;
            } else {
                result += currentValue;
            }
            
        }
        return result;
    }

    function getValue(bytes1 c) private view returns (uint) {
        return valueMap[c];
    }

}