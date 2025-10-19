// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/**
 * @dev ERC20 接口合约
 **/
interface IERC20 {
    //@dev 释放条件：当value单位的货币从账户from转账到另一个账户to时
    event Transfer(address indexed from, address indexed to, uint256 value);

    //@dev 释放条件：当value单位的货币从账户owner授权给另一个账户spender时
    event Approval(address indexed owner, address indexed spender, uint256 value);

    //@dev 返回代币总供给
    function totalSupply() external view returns (uint256);

    //@dev 返回账户account所持有的代币数
    function balanceOf(address account) external view returns (uint256);

    /**
     *@dev 转账account单位代币，从调用者账户到另一个账户to
     * 如果成功，返回true
     * 释放{Transfer}事件
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     *@dev 返回owner账户授权给spender账户的额度，默认为0
     *当{approve} 或 {transferFrom} 被调用时，allownance会改变
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     *@dev 调用者账户给spender账户授权amount数量代币
     *如果成功，返回true
     *释放{Approval}事件
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     *@dev 通过授权机制，从from账户向to账户转账amount数量代币，转账的部分会从调用者的allowance中扣除
     *如果成功，返回true
     *释放{Transfer}事件
     */
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}