// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract IntToRoman {
    
    ValueSymbols[] private symbols ;

    constructor() {
        symbols.push(ValueSymbols(1000, "M"));
        symbols.push(ValueSymbols(900,  "CM"));
        symbols.push(ValueSymbols(500, "D"));
        symbols.push(ValueSymbols(400, "CD"));
        symbols.push(ValueSymbols(100, "C"));
        symbols.push(ValueSymbols(90, "XC"));
        symbols.push(ValueSymbols(50, "L"));
        symbols.push(ValueSymbols(40, "XL"));
        symbols.push(ValueSymbols(10, "X"));
        symbols.push(ValueSymbols(5, "V"));
        symbols.push(ValueSymbols(4, "IV"));
        symbols.push(ValueSymbols(1, "I"));
    }

    struct ValueSymbols {
        uint value;
        string symbol;
    }

    function intToRoman(uint num) public view returns(string memory) {
        string memory str;
        for (uint i = 0; i < symbols.length; i++) {
            ValueSymbols memory vs = symbols[i];
            while (num >= vs.value) {
                num -= vs.value;
                str = string.concat(str, vs.symbol);
            }
            if (num == 0) {
                break;
            }
        }
        return str;
    }


}