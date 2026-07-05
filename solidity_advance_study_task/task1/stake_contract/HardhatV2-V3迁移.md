## HardhatV2-V3迁移记录

### 1 操作步骤

#### 1.1 清理旧环境与依赖

##### 1.1.1 确保Node.js版本，Hardhat v3要求Node.js v22.10.0或更高版本。
```bash
node --version
```

##### 1.1.2 清理缓存
```bash
npx hardhat clean
```

##### 1.1.3 彻底卸载Hardhat v2及其旧插件：
打开`package.json`，彻底移除`hardhat`、所有以`hardhat-`、`@nomicfoundation/`、`@nomiclabs/`开头的插件包，以及`solidity-coverage`和`hardhat-gas-reporter`。

##### 1.1.4 重新安装依赖以刷新依赖树：
```bash
npm install
npm why hardhat
```

##### 1.1.5 备份并修改旧配置文件名称：
```bash
mv hardhat.config.js hardhat.config.old.js
```

#### 1.2 将项目全面转换为ESM

##### 1.2.1 修改package.json：在根级别添加`"type":"module"`。


#### 1.3 安装Hardhat v3及新版工具链

##### 1.3.1 安装Hardhat v3核心包
```bash
npm add --save-dev hardhat
```

##### 1.3.2 安装专为v3设计的官方新版Toolbox插件
```bash
npm add --save-dev @nomicfoundation/hardhat-toolbox-mocha-ethers
npm add --save-dev dotenv
npm add --save-dev @openzeppelin/hardhat-upgrades
```

#### 1.4 编写全新的v3配置文件
在根目录下创建全新的`hardhat.config.ts`。请注意，v3的配置现在使用声明式的`defineConfig`函数包围，且插件引入使用标准的`import`语法。
```json
import {defineConfig} from "hardhat/config";
import "dotenv/config";
import hardhatUpgrades from '@openzeppelin/hardhat-upgrades';
import hardhatToolboxMochaEthers from "@nomicfoundation/hardhat-toolbox-mocha-ethers";
export default defineConfig({
    plugins: [hardhatToolboxMochaEthers, hardhatUpgrades],
    solidity: {
        version: "0.8.22",
    },
    networks: {
        sepolia: {
            type: "http",
            url:process.env.SEPOLIA_RPC_URL || "",
            accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
        },
    }
});
```

#### 1.5 修改测试文件的代码格式
```
import  hre  from "hardhat";
let upgradesApi;
let ethers;
let provider;
before(async () => {
    const connection = await hre.network.create();
    ({ ethers } = connection);
    provider = await ethers.provider; 
    upgradesApi = await upgrades(hre, connection);
});
```