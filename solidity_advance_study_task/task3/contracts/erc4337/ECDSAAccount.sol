// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Account} from "@openzeppelin/contracts/account/Account.sol";
import {SignerECDSA} from "@openzeppelin/contracts/utils/cryptography/signers/SignerECDSA.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";

/// @title ECDSAAccount
/// @notice 一个最小化的 ERC-4337 智能合约账户。
///
/// - UserOperation 的校验流程、nonce 管理、EntryPoint 相关逻辑,
///   全部来自 OpenZeppelin 的 {Account} 基类。
/// - 具体的签名验证算法(标准 ECDSA / EOA 签名),来自 OpenZeppelin
///   的 {SignerECDSA},它实现了 {AbstractSigner-_rawSignatureValidation}。
///
/// 因为账户通常是通过工厂 + 最小代理(Clones)批量部署的,构造函数里
/// 不设置任何状态,真正的初始化放在 `initialize` 里,并用 {Initializable}
/// 防止被重复初始化。
contract ECDSAAccount is Account, SignerECDSA, Initializable {
    // SignerECDSA 现在带一个构造函数参数(初始签名者地址)。这里传 address(0)
    // 只是占位——因为这个逻辑合约只会被 Clone,本身不存真正的签名者;
    // 每个 Clone 出来的账户,会在 initialize() 里通过 _setSigner 重新
    // 绑定各自真实的签名者,覆盖掉这个占位值。
    //
    // Account 基类还继承了 EIP712(v0.8 EntryPoint 的 userOpHash 是 EIP-712
    // 摘要,需要域分隔符),这里给它起个名字和版本号,任意取,只要保证
    // 全部 Clone 共用同一个域即可(它不像签名者那样需要按账户区分)。
    constructor() SignerECDSA(address(0)) {
        // 逻辑合约本身不会被直接使用(只会被 clone),提前锁死,
        // 避免有人直接调用逻辑合约的 initialize 抢占初始化。
        _disableInitializers();
    }

    /// @notice 由工厂在 clone 出账户之后调用一次,绑定这个账户的签名者(通常是用户的 EOA 公钥地址)
    function initialize(address signer) public initializer {
        _setSigner(signer);
    }

    /// @notice 执行一次外部调用。只能由 EntryPoint,或账户自己(比如账户内部批量操作)发起。
    function execute(address target, uint256 value, bytes calldata data) external {
        require(msg.sender == address(entryPoint()) || msg.sender == address(this), "ECDSAAccount: not authorized");
        _call(target, value, data);
    }

    /// @notice 批量执行,一次 UserOperation 里完成多步调用(比如 approve + swap)
    function executeBatch(
        address[] calldata targets,
        uint256[] calldata values,
        bytes[] calldata datas
    ) external {
        require(msg.sender == address(entryPoint()) || msg.sender == address(this), "ECDSAAccount: not authorized");
        require(targets.length == values.length && targets.length == datas.length, "ECDSAAccount: length mismatch");
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

    /// @notice 允许账户接收 ETH(比如 EntryPoint 结算 Gas 差额,或者别人直接转账)
    // receive() external payable {}
}
