import {buildModule} from "@nomicfoundation/hardhat-ignition/modules";
export default buildModule("MetaNodeStakeModule", (m) => {
  const MetaNodeToken = m.contract("MetaNodeToken");
  // 部署 MetaNodeStake 合约，传入 MetaNodeToken 地址、质押起始区块高度、质押结束区块高度和每个区块奖励的 MetaNode token 数量作为参数
  const MetaNodeStake = m.contract("MetaNodeStake");
  const metaNodePerBlock = 100n
  const startBlock = 1n
  const blockHight = 10000n
  m.call(MetaNodeStake, "initialize", [MetaNodeToken, startBlock, startBlock + blockHight, metaNodePerBlock]);
  return { MetaNodeStake, MetaNodeToken };
});
