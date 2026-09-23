// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Account} from "@openzeppelin/contracts/account/Account.sol";
import {SignerEIP7702} from "@openzeppelin/contracts/utils/cryptography/signers/SignerEIP7702.sol";

/// @title ECDSADelegateAccount
/// @notice 这份合约不会被单独部署使用,它是给 EOA 通过 EIP-7702
/// authorization 签名"委托"指向的实现合约。
///
/// 一旦某个 EOA 对这份合约签署了 7702 授权(authorization),
/// 该 EOA 的地址上就会执行这里的代码,但存储和 ETH 余额上下文
/// 仍然是这个 EOA 自己的——不需要迁移地址,也不需要部署新合约。
///
/// {SignerEIP7702} 是 OpenZeppelin 提供的签名器实现,它用
/// address(this) 本身作为签名验证的对象——也就是说,验证签名时
/// 用的正是这个 EOA 自己的地址,天然契合"EOA 自己给自己的操作签名"
/// 这个 EIP-7702 的核心场景,不需要额外的 initialize 步骤。
contract ECDSADelegateAccount is Account, SignerEIP7702 {
    /// @notice 执行一次外部调用。只能由 EntryPoint(如果接入 ERC-4337 基础设施),
    /// 或者这个地址自己(比如 EOA 自己发起的 7702 交易内部逻辑)发起。
    function execute(address target, uint256 value, bytes calldata data) external {
        require(msg.sender == address(entryPoint()) || msg.sender == address(this), "not authorized");
        _call(target, value, data);
    }

    /// @notice 批量执行,典型场景:approve + swap 一次性完成
    function executeBatch(
        address[] calldata targets,
        uint256[] calldata values,
        bytes[] calldata datas
    ) external {
        require(msg.sender == address(entryPoint()) || msg.sender == address(this), "not authorized");
        require(targets.length == values.length && targets.length == datas.length, "length mismatch");
        for (uint256 i = 0; i < targets.length; i++) {
            _call(targets[i], values[i], datas[i]);
        }
    }

    function _call(address target, uint256 value, bytes calldata data) internal {
        (bool success, bytes memory ret) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(ret, 32), mload(ret))
            }
        }
    }

    // receive() external payable {}
}
