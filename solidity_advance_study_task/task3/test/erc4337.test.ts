import { describe, it } from "node:test";
import {expect} from "chai";
import hre from "hardhat";

const {viem, ethers, networkHelpers} = await hre.network.create();

describe("ERC-4337: ECDSAAccount (基于 OpenZeppelin Account + SignerECDSA)", function () {
  it("通过 EntryPoint.handleOps 执行一笔转账 UserOperation", async function () {
    const [deployer, beneficiary] = await ethers.getSigners();

    // 模拟用户的 EOA:只用它的私钥签名 UserOperation,
    // 从头到尾都不需要它自己发起链上交易(交易由 bundler/deployer 打包提交)。
    const ownerWallet = ethers.Wallet.createRandom().connect(ethers.provider);

    // 1) 部署 EntryPoint(eth-infinitism 的参考实现,通过 Imports.sol 引入编译)
    const EntryPoint = await ethers.getContractFactory("EntryPoint");
    const entryPoint = await EntryPoint.deploy();
    await entryPoint.waitForDeployment();
    const entryPointAddress = await entryPoint.getAddress();
    console.log("entryPointAddress:", entryPointAddress);
    console.log("ownerWallet.address:", ownerWallet.address);
    console.log("deployer.address:", deployer.address);
    console.log("beneficiary.address:", beneficiary.address);

    // 2) 部署工厂,并通过工厂创建一个绑定 ownerWallet 的智能账户
    const Factory = await ethers.getContractFactory("ECDSAAccountFactory");
    const factory = await Factory.deploy();
    await factory.waitForDeployment();

    const salt = 0n;
    await (await factory.createAccount(ownerWallet.address, salt)).wait();
    const accountAddress = await factory.getAddress(ownerWallet.address, salt);
    const account = await ethers.getContractAt("ECDSAAccount", accountAddress);

    // 账户自己得知道去哪个 EntryPoint 验证/结算——这里假设 Account 基类
    // 的 entryPoint() 默认指向 ERC4337Utils 里内置的规范地址;如果你的
    // OpenZeppelin 版本要求显式传入/设置 EntryPoint 地址,按提示调整即可。
    void entryPointAddress;

    // 3) 给账户预存一点 ETH,同时往 EntryPoint 里给账户充值,
    // 用于覆盖 EntryPoint 垫付的验证/执行 Gas。
    await deployer.sendTransaction({ to: accountAddress, value: ethers.parseEther("1") });
    await entryPoint.depositTo(accountAddress, { value: ethers.parseEther("1") });

    // 4) 构造 callData:调用账户自己的 execute(),给 beneficiary 转 0.1 ETH
    const callData = account.interface.encodeFunctionData("execute", [
      beneficiary.address,
      ethers.parseEther("0.1"),
      "0x",
    ]);

    const nonce = await entryPoint.getNonce(accountAddress, 0);

    const userOp = {
      sender: accountAddress,
      nonce,
      initCode: "0x",
      callData,
      // PackedUserOperation 把 verificationGasLimit 和 callGasLimit
      // 打包进同一个 bytes32(各占 16 字节)
      accountGasLimits: ethers.solidityPacked(["uint128", "uint128"], [300000, 300000]),
      preVerificationGas: 100000,
      // 同理,maxPriorityFeePerGas 和 maxFeePerGas 打包进一个 bytes32
      gasFees: ethers.solidityPacked(["uint128", "uint128"], [1_000_000_000, 1_000_000_000]),
      paymasterAndData: "0x",
      signature: "0x",
    };

    // 5) 计算 userOpHash,并直接对这个哈希做原始 ECDSA 签名
    //    (注意:不要用 signMessage/personal_sign,那样会额外加上
    //    "\x19Ethereum Signed Message" 前缀,导致签名对不上——
    //    userOpHash 本身已经是最终的、可直接签名的摘要)
    const userOpHash: string = await entryPoint.getUserOpHash(userOp);
    const signature = ownerWallet.signingKey.sign(userOpHash).serialized;
    userOp.signature = signature;

    const beforeBalance = await ethers.provider.getBalance(beneficiary.address);

    // 6) 模拟 bundler:把这个 UserOperation 打包提交给 EntryPoint
    await (await entryPoint.handleOps([userOp], deployer.address)).wait();

    const afterBalance = await ethers.provider.getBalance(beneficiary.address);
    expect(afterBalance - beforeBalance).to.equal(ethers.parseEther("0.1"));
  });
});
