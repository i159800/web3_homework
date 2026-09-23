// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Clones} from "@openzeppelin/contracts/proxy/Clones.sol";
import {ECDSAAccount} from "./ECDSAAccount.sol";

/// @title ECDSAAccountFactory
/// @notice 负责部署 {ECDSAAccount} 的最小代理(EIP-1167 Clone),
/// 用 CREATE2 让地址在部署之前就能被预先算出来(counterfactual address),
/// 这样 ERC-4337 的 initCode 才能在账户第一次收到 UserOperation 时按需部署。
contract ECDSAAccountFactory {
    ECDSAAccount public immutable accountImplementation;

    constructor() {
        // 部署一次逻辑合约,后续所有账户都是指向它的最小代理
        accountImplementation = new ECDSAAccount();
    }

    /// @notice 创建(或者如果已存在则直接返回)一个绑定给定 signer 的账户
    function createAccount(address signer, uint256 salt) external returns (ECDSAAccount) {
        address predicted = getAddress(signer, salt);
        if (predicted.code.length > 0) {
            return ECDSAAccount(payable(predicted));
        }
        ECDSAAccount account = ECDSAAccount(
            payable(Clones.cloneDeterministic(address(accountImplementation), _salt(signer, salt)))
        );
        account.initialize(signer);
        return account;
    }

    /// @notice 预先计算出账户地址,不需要真的部署
    function getAddress(address signer, uint256 salt) public view returns (address) {
        return Clones.predictDeterministicAddress(address(accountImplementation), _salt(signer, salt));
    }

    function _salt(address signer, uint256 salt) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(signer, salt));
    }
}
