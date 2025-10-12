#### 作业1：ERC20代币
##### 任务：参考openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的ERC20代币合约。要求：
1. 合约包含以下标准ERC20功能
2. balanceOf: 查询账户余额
3. transfer: 转账
4. approve 和 transferFrom: 授权和代扣转账
5. 使用event记录转账和授权操作
6. 提供mint函数，允许合约所有者增发代币。
##### 提示：
- 使用mapping存储账户余额和授权信息。
- 使用event定义Transfer和Approval事件。
- 部署到sepolia测试网，导入到自己的钱包

#### 作业2：在测试网上发行一个图文并茂的NFT
##### 任务目标
1. 使用Solidity编写一个符合ERC721标准的NFT合约
2. 将图文数据上传到IPFS，生成元数据链接。
3. 将合约部署到以太测试网(如Goerli或Sepolia)
4. 铸造NFT并在测试网环境中查看。
##### 任务步骤
1. 编写NFT合约
    - 使用OpenZeppelin的ERC721库编写一个NFT合约。
    - 合约应包含以下功能。
    - 构造函数：设置NFT的名称和称号。
    - mintNFT函数：允许用户铸造NFT，并关联元数据链接(tokenURI)。
    - 在RemixIDE中编译合约。
2. 准备图文数据
    - 准备一张图片，并将其上传到IPFS(可以使用Pinata或其他工具)
    - 创建一个JSON文件，描述NFT的属性(如名称、描述、图片链接等)
    - 将JSON文件上传到IPFS，获取元数据链接。
    - [JSON文件参考](https://docs.opensea.io/docs/metadata-standards)
3. 部署合约到测试网
    - 在RemixIDE中连接MetaMask，并确保MetaMask连接到Goerli或Sepolia测试网。
    - 部署NFT合约到测试网，并记录合约地址。
4. 铸造NFT
    - 使用mintNFT函数铸造NFT
    - 在recipient字段中输入你的钱包地址
    - 在tokenURI字段中输入元数据的IPFS链接
    - 在MetaMask中确认交易
5. 查看NFT
    - 打开OpenSea测试网或Etherscan测试网
    - 连接你的钱包，查看你铸造的NFT

#### 作业3：编写一个讨饭合约
##### 任务目标
1. 使用Solidity编写一个合约，允许用户向合约地址发送以太币
2. 记录每个捐赠者的地址和捐赠金额
3. 允许合约所有者提取所有捐赠的资金。

##### 任务步骤
1. 编写合约
    - 创建一个名为BeggingContract的合约。
    - 合约应包含以下功能：
        - 一个mapping来记录每个捐赠者的捐赠金额。
        - 一个donate函数，允许用户向合约发送以太币，并记录捐赠信息。
        - 一个withdraw函数，允许合约所有者提取所有资金。
        - 一个getDonation函数，允许查询某个地址的捐赠金额。
        - 使用payable修饰符和address.transfer实现支付和提款。
2. 部署合约
    - 在RemixIDE编译合约。
    - 部署合约到Goerli或Sepolia测试网。
3. 测试合约
    - 使用MetaMask向合约发送以太币，测试donate功能。
    - 调用withdraw函数，测试合约所有者是否可以提取资金。
    - 调用getDonation函数，查询某个地址的捐赠金额。

##### 任务要求：
1. 合约代码：
    - 使用mapping记录捐赠者的地址和金额
    - 使用payable修饰符实现donate和withdraw函数。
    - 使用onlyOwner修饰符限制withdraw函数只能由合约所有者调用。
2. 测试网部署：
    - 合约必须部署到Goerli或Sopolia测试网。
3. 功能测试：
    - 确保donate、withdraw和getDonation函数正常工作。

##### 提交内容
1. 合约代码：提交Solidity合约文件(如BeggingContract.sol)
2. 合约地址：提交部署到测试网的合约地址。
3. 测试截图：提交在Remix或Etherscan上测试合约的截图。

##### 额外挑战(可选)
1. 捐赠事件：添加Donation事件，记录每次捐赠的地址和金额。
2. 捐赠排行榜：实现一个功能，显式捐赠金额最多的前3个地址。
3. 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。