## Factory

### 接口定义

```solidity
// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.24;

interface IFactory {
    struct Parameters {
        address factory;
        address tokenA;
        address tokenB;
        int24 tickLower;
        int24 tickUpper;
        uint24 fee;
    }

    function parameters()
        external
        view
        returns (
            address factory,
            address tokenA,
            address tokenB,
            int24 tickLower,
            int24 tickUpper,
            uint24 fee
        );

    event PoolCreated(
        address token0,
        address token1,
        uint32 index,
        int24 tickLower,
        int24 tickUpper,
        uint24 fee,
        address pool
    );

    function getPool(
        address tokenA,
        address tokenB,
        uint32 index
    ) external view returns (address pool);

    function createPool(
        address tokenA,
        address tokenB,
        int24 tickLower,
        int24 tickUpper,
        uint24 fee
    ) external returns (address pool);
}
```

### 具体代码实现

**`核心变量定义`**
```solidity
mapping(address => mapping(address => address[])) public pools; //流动性池映射:[tokenA][tokenB][]
```

**`sortToken`**
- 作用：代币对的排序
```solidity
function sortToken(address tokenA, address tokenB) private pure returns (address, address) {
    return tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
}
```

**`getPool`**
- 作用：获取代币A+代币B+index的流动性池地址
```solidity
function getPool(
    address tokenA,
    address tokenB,
    uint32 index
) external view override returns (address) {
    require(tokenA != tokenB, "IDENTICAL_ADDRESSES");
    require(tokenA != address(0) && tokenB != address(0), "ZERO_ADDRESS");

    // Declare token0 and token1
    address token0;
    address token1;

    (token0, token1) = sortToken(tokenA, tokenB);

    return pools[token0][token1][index];
}
```

**`createPool`**
- 作用：创建流动性池
```solidity
function createPool(
    address tokenA,
    address tokenB,
    int24 tickLower,
    int24 tickUpper,
    uint24 fee
) external override returns (address pool) {
    // validate token's individuality
    require(tokenA != tokenB, "IDENTICAL_ADDRESSES");

    // Declare token0 and token1
    address token0;
    address token1;

    // sort token, avoid the mistake of the order
    (token0, token1) = sortToken(tokenA, tokenB);

    // get current all pools
    address[] memory existingPools = pools[token0][token1];

    // check if the pool already exists
    for (uint256 i = 0; i < existingPools.length; i++) {
        IPool currentPool = IPool(existingPools[i]);

        if (
            currentPool.tickLower() == tickLower &&
            currentPool.tickUpper() == tickUpper &&
            currentPool.fee() == fee
        ) {
            return existingPools[i];
        }
    }

    // save pool info
    parameters = Parameters(
        address(this),
        token0,
        token1,
        tickLower,
        tickUpper,
        fee
    );

    // generate create2 salt
    bytes32 salt = keccak256(
        abi.encode(token0, token1, tickLower, tickUpper, fee)
    );

    // create pool
    pool = address(new Pool{salt: salt}());

    // save created pool
    pools[token0][token1].push(pool);

    // delete pool info
    delete parameters;

    emit PoolCreated(
        token0,
        token1,
        uint32(existingPools.length),
        tickLower,
        tickUpper,
        fee,
        pool
    );
}
```


### 核心接口分析

**`createPool`** 

我们通过 `pool = address(new Pool{salt: salt}());` 这一行代码创建了一个新的 `Pool` 合约，并通过 `pools[token0][token1].push(pool);` 将它的地址保存到 `pools` 中。

这里需要注意的是，我们通过添加了 `salt` 来使用CREATE2的方式来创建合约，这样的好处是创建出来的合约地址是可预测的，地址生成的逻辑是 `新地址 = hash("0xFF",创建者地址, salt, initcode)`。

而在我们的代码中 `salt` 是通过 `abi.encode(token0, token1, tickLower, tickUpper, fee)` 计算出来的，这样的好处是只要我们知道了 `token0` 和 `token1` 的地址，以及 `tickLower`、`tickUpper` 和 `fee` 这三个参数，我们就可以预测出来新合约的地址。在我们的教程设计中，这样似乎并没有什么用。但是在实际的 DeFi 场景中，这样会带来很多好处。比如其他合约可以直接计算出我们 `Pool` 合约的地址，这样可以开发出和 `Pool` 合约交互的更多的功能。

当然，这样也会带来一个问题，这样会使得我们不能通过合约的构造函数传参来传递 `Pool` 合约的初始化参数，因为那样会导致上面新地址计算中的 `initcode` 发生变化。所以我们在代码中引入了 `parameters` 这个变量来保存 `Pool` 合约的初始化参数，这样我们就可以在 `Pool` 合约中通过 `parameters` 来获取到初始化参数。


## PoolManager

### 接口定义

```solidity
// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.24;
pragma abicoder v2;

import "./IFactory.sol";

interface IPoolManager is IFactory {
    struct PoolInfo {
        address pool;
        address token0;
        address token1;
        uint32 index;
        uint24 fee;
        uint8 feeProtocol;
        int24 tickLower;
        int24 tickUpper;
        int24 tick;
        uint160 sqrtPriceX96;
        uint128 liquidity;
    }

    struct Pair {
        address token0;
        address token1;
    }

    function getPairs() external view returns (Pair[] memory);

    function getAllPools() external view returns (PoolInfo[] memory poolsInfo);

    struct CreateAndInitializeParams {
        address token0;
        address token1;
        uint24 fee;
        int24 tickLower;
        int24 tickUpper;
        uint160 sqrtPriceX96;
    }

    function createAndInitializePoolIfNecessary(
        CreateAndInitializeParams calldata params
    ) external payable returns (address pool);
}
```

### 具体代码实现

**`getPairs()`**
- 作用：返回pairs数组的全部数据，供前端使用

```solidity
function getPairs() external view override returns (Pair[] memory) {
        return pairs;
    }
```

**`getAllPools()`**
-作用: 对Factory保存的pools数组的数据进行重新组装并返回。获得数据供前端展示。

```solidity
function getAllPools()
        external
        view
        override
        returns (PoolInfo[] memory poolsInfo)
    {
        uint32 length = 0;
        // 先算一下大小
        for (uint32 i = 0; i < pairs.length; i++) {
            length += uint32(pools[pairs[i].token0][pairs[i].token1].length);
        }

        // 再填充数据
        poolsInfo = new PoolInfo[](length);
        uint256 index;
        for (uint32 i = 0; i < pairs.length; i++) {
            address[] memory addresses = pools[pairs[i].token0][
                pairs[i].token1
            ];
            for (uint32 j = 0; j < addresses.length; j++) {
                IPool pool = IPool(addresses[j]);
                poolsInfo[index] = PoolInfo({
                    pool: addresses[j],
                    token0: pool.token0(),
                    token1: pool.token1(),
                    index: j,
                    fee: pool.fee(),
                    feeProtocol: 0,
                    tickLower: pool.tickLower(),
                    tickUpper: pool.tickUpper(),
                    tick: pool.tick(),
                    sqrtPriceX96: pool.sqrtPriceX96(),
                    liquidity: pool.liquidity()
                });
                index++;
            }
        }
        return poolsInfo;
    }
```

**`createAndInitializePoolIfNecessary`**
- 作用：创建池子，并动态维护一个`pairs`数组。

```solidity
function createAndInitializePoolIfNecessary(
        CreateAndInitializeParams calldata params
) external payable override returns (address poolAddress) {
    require(
        params.token0 < params.token1,
        "token0 must be less than token1"
    );

    poolAddress = this.createPool(
        params.token0,
        params.token1,
        params.tickLower,
        params.tickUpper,
        params.fee
    );

    IPool pool = IPool(poolAddress);

    uint256 index = pools[pool.token0()][pool.token1()].length;

    // 新创建的池子，没有初始化价格，需要初始化价格
    if (pool.sqrtPriceX96() == 0) {
        pool.initialize(params.sqrtPriceX96);

        if (index == 1) {
            // 如果是第一次添加该交易对，需要记录
            pairs.push(
                Pair({token0: pool.token0(), token1: pool.token1()})
            );
        }
    }
}
```
### 核心接口分析

**`getPairs`**

对于数组这个特殊情况，如果我们需要它返回全部的内容，就需要自己写一个合约方法将其返回。

获取到的数据供前端的token列表用。

**`getAllPools`**

由于池子的信息是在 `Factory` 合约中保存的，因此我们在返回全部池子信息的时候，还需要对 `Factory` 保存的信息进行处理，处理成我们想要的数据格式。

这部分的逻辑比较清晰，通过遍历全部的池子信息，做一些数据转换就行。

获取到的数据供前端展示所有的pool信息。

**`createAndInitializePoolIfNecessary`**

在 `Factory` 合约中，每次创建完成一个池子，都会记录下它的信息，因此这个信息我们不需要再记录，我们需要记录的是交易对的种类，即在获得一个新的交易对时，动态地维护一个 `pairs` 数组。

需要注意的是，虽然 `createPool` 的入参 `tokenA` 和 `tokenB` 没有顺序要求，但是在 `createAndInitializePoolIfNecessary` 中我们创建的时候要求 `token0 < token1`。因为在这个方法中需要传入初始化的价格，而在交易池中价格是按照 `token0/token1` 的方式计算的，做这个限制可以避免 LP 不小心初始化错误的价格。在后续的代码和测试中，我们也约定了 `tokenA` 和 `tokenB` 是未排序的，而 `token0` 和 `token1` 是排序的，这样也便于我们理解代码。


## Pool

### 接口定义

```solidity
// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.24;

interface IMintCallback {
    function mintCallback(
        uint256 amount0Owed,
        uint256 amount1Owed,
        bytes calldata data
    ) external;
}

interface ISwapCallback {
    function swapCallback(
        int256 amount0Delta,
        int256 amount1Delta,
        bytes calldata data
    ) external;
}

interface IPool {
    function factory() external view returns (address);

    function token0() external view returns (address);

    function token1() external view returns (address);

    function fee() external view returns (uint24);

    function tickLower() external view returns (int24);

    function tickUpper() external view returns (int24);

    function sqrtPriceX96() external view returns (uint160);

    function tick() external view returns (int24);

    function liquidity() external view returns (uint128);

    function initialize(uint160 sqrtPriceX96) external;

    /// feeGrowthGlobal0X128 记录从创建到现在，每个流动性累计产生的 token0 的手续费
    /// @notice The fee growth as a Q128.128 fees of token0 collected per unit of liquidity for the entire life of the pool
    /// @dev This value can overflow the uint256
    function feeGrowthGlobal0X128() external view returns (uint256);

    /// feeGrowthGlobal1X128 记录从创建到现在，每个流动性累计产生的 token1 的手续费
    /// @notice The fee growth as a Q128.128 fees of token1 collected per unit of liquidity for the entire life of the pool
    /// @dev This value can overflow the uint256
    function feeGrowthGlobal1X128() external view returns (uint256);

    function getPosition(
        address owner
    )
        external
        view
        returns (
            uint128 _liquidity,
            uint256 feeGrowthInside0LastX128,
            uint256 feeGrowthInside1LastX128,
            uint128 tokensOwed0,
            uint128 tokensOwed1
        );

    event Mint(
        address sender,
        address indexed owner,
        uint128 amount,
        uint256 amount0,
        uint256 amount1
    );

    function mint(
        address recipient,
        uint128 amount,
        bytes calldata data
    ) external returns (uint256 amount0, uint256 amount1);

    event Collect(
        address indexed owner,
        address recipient,
        uint128 amount0,
        uint128 amount1
    );

    function collect(
        address recipient,
        uint128 amount0Requested,
        uint128 amount1Requested
    ) external returns (uint128 amount0, uint128 amount1);

    event Burn(
        address indexed owner,
        uint128 amount,
        uint256 amount0,
        uint256 amount1
    );

    function burn(
        uint128 amount
    ) external returns (uint256 amount0, uint256 amount1);

    event Swap(
        address indexed sender,
        address indexed recipient,
        int256 amount0,
        int256 amount1,
        uint160 sqrtPriceX96,
        uint128 liquidity,
        int24 tick
    );

    function swap(
        address recipient,
        bool zeroForOne,
        int256 amountSpecified,
        uint160 sqrtPriceLimitX96,
        bytes calldata data
    ) external returns (int256 amount0, int256 amount1);
}
```

### 具体代码实现

**`核心变量定义`**

```solidity
//Factory工厂地址
address public immutable override factory;
//token0代币地址
address public immutable override token0;
//token1代币地址
address public immutable override token1;
//交易手续费
uint24 public immutable override fee;
//池子的价格下限
int24 public immutable override tickLower;
//池子的价格上限
int24 public immutable override tickUpper;

//当前池子的价格的开平方*2的96次方
uint160 public override sqrtPriceX96;
//当前池子的价格对应的tick
int24 public override tick;
//当前池子的流动性
uint128 public override liquidity;

//token0代币上积累的全部手续费
uint256 public override feeGrowthGlobal0X128;
//token1代币上积累的全部手续费
uint256 public override feeGrowthGlobal1X128;

struct Position {
    //LP在池子中拥有的流动性
    uint128 liquidity;
    // 可提取的 token0手续费数量
    uint128 tokensOwed0;
    // 可提取的 token1手续费数量
    uint128 tokensOwed1;
    // 上次提取手续费时的 feeGrowthGlobal0X128
    uint256 feeGrowthInside0LastX128;
    // 上次提取手续费是的 feeGrowthGlobal1X128
    uint256 feeGrowthInside1LastX128;
}

// 用一个 mapping 来存放所有 Position 的信息
mapping(address => Position) public positions;

//交易中需要临时存储的变量
struct SwapState {
    int256 amountSpecifiedRemaining;
    int256 amountCalculated;
    uint160 sqrtPriceX96;
    uint256 feeGrowthGlobalX128;
    uint256 amountIn;
    uint256 amountOut;
    uint256 feeAmount;
}
```

**`_modifyPosition()`**
- 参数：
    - ModifyPositionParams{address: owner, int128: liquidityDelta}
- 具体实现：
    1. 计算amount0:$$amount0=liq*q96*\frac{(sqrtPriceB-sqrtPriceA)}{sqrtPriceB*sqrtPriceA}$$
    2. 计算amount1:$$amount1=liq*\frac{(sqrtPriceB-sqrtPriceA)}{q96}$$
    3. 提取之前的amount0手续费:$$tokensOwed0=\frac{(feeGrowthGlobal0X128-position.feeGrowthInside0LastX128)*position.liquidity}{q128}$$
    4. 提取之前的amount1手续费:$$tokensOwed1=\frac{(feeGrowthGlobal1X128-position.feeGrowthInside1LastX128)*position.liquidity}{q128}$$
    5. 更新position提取手续费的记录，同步为当前最新
    6. 把可以提取的手续费追加到position.tokensOwed0和position.tokensOwed1上
    7. 把本次mint的liquidity追加到position.liquidity上。

**`mint`**
- 作用：添加流动性

- 具体实现：
    1. 调用`_modifyPosition()方法，计算需要添加的token0和token1数量，同时更新LP的position信息`
    2. 调用mintCallback将代币转给当前的Pool合约

```solidity
function mint(
        address recipient,
        uint128 amount,
        bytes calldata data
) external override returns (uint256 amount0, uint256 amount1) {
    require(amount > 0, "Mint amount must be greater than 0");
    // 基于 amount 计算出当前需要多少 amount0 和 amount1
    (int256 amount0Int, int256 amount1Int) = _modifyPosition(
        ModifyPositionParams({
            owner: recipient,
            liquidityDelta: int128(amount)
        })
    );
    amount0 = uint256(amount0Int);
    amount1 = uint256(amount1Int);

    uint256 balance0Before;
    uint256 balance1Before;
    if (amount0 > 0) balance0Before = balance0();
    if (amount1 > 0) balance1Before = balance1();
    // 回调 mintCallback
    IMintCallback(msg.sender).mintCallback(amount0, amount1, data);

    if (amount0 > 0)
        require(balance0Before.add(amount0) <= balance0(), "M0");
    if (amount1 > 0)
        require(balance1Before.add(amount1) <= balance1(), "M1");

    emit Mint(msg.sender, recipient, amount, amount0, amount1);
}
```

**`swap()`*
- 作用：代币交换

- 参数解读：
    1. recipient:代币交换的发起者地址
    2. zeroForOne:是否token0->token1 或 token1->token0
    3. amountSpcified:指定交易的代币数量
    4. sqrtPriceLimitX96:限定价格，超过此价格，交易结束
    5. calldata data:执行对应的代币转移

- 具体实现：
    1. 执行前验证：
        1. token0->token1，价格下降，要求限定价格低于当前价格
        2. token1->token0，价格上升，要求限定价格高于当前价格
    2. 确认`exactInput`是否为true，即代表是确定输入值
    3. 初始化SwapState变量
    4. 计算用户交易价格的限制：
        1. 如果是zeroForOne=true，`sqrtPriceX96PoolLimit`=sqrtPriceX96Lower
        2. 如果是zeroForOne=false,`sqrtPriceX96PoolLimit`=sqrtPriceX96Upper
    5. 计算交易的具体数值
        5.1 使用`SwapMath.computeSwapStep()`
            1. 计算实际消耗的代币数量:$amountRemainingLessFee=\frac{amountRemaining*(10^6-fee)}{10^6}$
            2. 计算amount0Delta=$$$$