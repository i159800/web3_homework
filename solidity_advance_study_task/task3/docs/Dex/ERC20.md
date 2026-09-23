## ERC20
`ERC20`是以太坊上的代币标准，来自2015年11月V神参与的[`EIP20`](https://eips.ethereum.org/EIPS/eip-20)。它实现了代币转账的基本逻辑：

- 账户余额(balanceOf())
- 代币总供给(totalSupply())
- 授权转账额度(allowance())
- 转账(transfer())
- 授权(approve())
- 授权转账(transferFrom())
- 代币信息(可选)：名称(name())，代号(symbol())，小数位数(decimals())

## IERC20

`IERC20`是`ERC20`代币标准的接口合约，规定了`ERC20`代币需要实现的函数和事件。
之所以需要定义接口，是因为有了规范后，就存在所有的`ERC20`代币都通用的函数名称，输入参数，输出参数。
在接口函数中，只需要定义函数名称，输入参数，输出参数，并不关心函数内部如何实现。
由此，函数就分为内部和外部两个内容，一个重点是实现，另一个是对外接口，约定共同数据。
这就是为什么需要`ERC20.sol`和`IERC20.sol`两个文件实现一个合约。

### 事件

`IERC20`定义了`2`个事件：`Transfer`事件和`Approval`事件，分别在转账和授权时被释放

```solidity
/**
 * @dev 释放条件：当`value`单位的货币从账户(`from`)转账到另一账户(`to`)时.
 */
event Transfer(address indexed from, address indexed to, uint256 value);

/**
 * @dev 释放条件：当`value`单位的货币从账户(`owner`)授权给另一账户(`spender`)时.
 */
event Approval(address indexed owner, address indexed spender, uint256 value);
```

### 函数

`IERC20`定义了`6`个函数，提供了转移代币的基本功能，并允许代币获得批准，以便其他链上第三方使用。

- `totalSupply()`返回代币总供给
    ```solidity
    /**
     * @dev 返回代币总供给
     */
    function totalSupply() external view returns (uint256);
    ```
- `balanceOf()`返回账户余额
    ```solidity
    /**
     * @dev 返回账户`account`所持有的代币数
     */
    function balanceOf(address account) external view returns(uint256);
    ```
- `allowance()`返回授权额度
    ```solidity
    /**
     * @dev 返回`owner`账户授权给`spender`账户的额度，默认为0
     * 当{approve}或{transferFrom}被调用时，`allowance`会改变
     */
    function allowance(address owner, address spender) external view returns (uint256);
    ```
- `transfer()`转账
    ```solidity
    /**
     * @dev 转账`amount`单位代币，从调用者账户到另一账户`to`
     * 如果成功，返回`true`
     *
     * 释放{Transfer}事件
     */
    function transfer(address to, uint256 amount) external returns(bool);
    ```

- `approve()`授权
    ```solidity
    /**
     * @dev 调用者账户给`spender`账户授权`amount`数量代币. 
     *
     * 如果成功，返回`true`
     *
     * 释放{Approval}事件
     */
    function approve(address spender, uint256 amount) external returns (bool);
    ```
- `transferFrom()`授权转账
    ```solidity
    /**
     * @dev 通过授权机制，从`from`账户向`to`账户转账`amount`数量代币。转账的部分会从调用者的`allowance`中扣除。
     *
     * 如果成功，返回`true`
     *
     * 释放{Transfer}事件
     */
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
    ```