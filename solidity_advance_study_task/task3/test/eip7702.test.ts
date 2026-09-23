import { expect } from "chai";
import hre from "hardhat";
const {viem} = await hre.network.create();
import { parseEther, encodeFunctionData } from "viem";

describe("EIP-7702: ECDSADelegateAccount (基于 OpenZeppelin Account + SignerEIP7702)", function () {
  it("EOA 对实现合约签署 7702 authorization 后,用自己的地址直接发起批量调用", async function () {
    const publicClient = await viem.getPublicClient();
    const [relayerClient, eoaClient, beneficiaryClient] = await viem.getWalletClients();

    // 1) 部署委托目标合约(实现逻辑),注意:这份合约本身不存资金、
    //    不会被用户直接交互,它只是被"借用"代码的模板。
    const delegate = await viem.deployContract("ECDSADelegateAccount");

    // 2) EOA 对这份实现合约签署 EIP-7702 authorization。
    //    executor: "self" 表示接下来这笔 Type 4 交易由 EOA 自己
    //    签名并广播、自己付 Gas(对应"用户自己升级自己的钱包"场景;
    //    如果换成不传 executor,则默认由另一个账户代付 Gas广播,
    //    对应"赞助商/中继商代付"场景)。
    const authorization = await eoaClient.signAuthorization({
      contractAddress: delegate.address,
      executor: "self",
    });

    // 3) 给这个 EOA 打点 ETH,用来支付 Gas,以及后续要转出去的金额
    await relayerClient.sendTransaction({
      to: eoaClient.account!.address,
      value: parseEther("1"),
    });

    const beneficiary = beneficiaryClient.account!.address;
    const before = await publicClient.getBalance({ address: beneficiary });

    // 4) 构造 calldata:调用(委托生效后的)自己地址上的 execute(),
    //    给 beneficiary 转 0.1 ETH
    const callData = encodeFunctionData({
      abi: delegate.abi,
      functionName: "execute",
      args: [beneficiary, parseEther("0.1"), "0x"],
    });

    // 5) 发送 Type 4(Set Code)交易:携带 authorizationList,
    //    发往 EOA 自己的地址——委托生效后,这次调用会执行
    //    ECDSADelegateAccount 的 execute() 逻辑,但存储/余额
    //    上下文仍然是这个 EOA 自己的。
    const txHash = await eoaClient.sendTransaction({
      authorizationList: [authorization],
      to: eoaClient.account!.address,
      data: callData,
    });
    await publicClient.waitForTransactionReceipt({ hash: txHash });

    const after = await publicClient.getBalance({ address: beneficiary });
    expect(after - before).to.equal(parseEther("0.1"));

    // 6) 验证委托确实生效了:EOA 地址上现在应该能读到委托指示符
    //    (0xef0100 + 实现合约地址),eth_getCode 不再返回空字节
    const code = await publicClient.getCode({ address: eoaClient.account!.address });
    expect(code).to.not.equal("0x");
  });
});
