// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { IERC1271 } from "@openzeppelin/contracts/interfaces/IERC1271.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract Token_1271 is IERC1271, Ownable {
    // 定义 ERC-1271 成功返回的魔数：0x1626ba7e
    bytes4 constant internal MAGICVALUE = 0x1626ba7e;

    constructor(address owner) Ownable(owner) {}

    function isValidSignature(bytes32 hash, bytes memory signature)
        external
        view
        override
        returns (bytes4 magicValue)
    {
        // 验证签名者是否为合约的 owner
        address signer = ECDSA.recover(hash, signature);
        if (signer == owner()) {
            return MAGICVALUE;
        } else {
            return 0xffffffff;
        }
    }
}