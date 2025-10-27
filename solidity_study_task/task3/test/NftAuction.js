const {ethers, deployments, upgrades} = require("hardhat");
const {expect} = require("chai");
// describe("Starting", async () => {
//     const Contract = await ethers.getContractFactory("NftAuction");
//     const contract = await Contract.deploy();
//     await contract.waitForDeployment();

//     await contract.createAuction(
//         1000 * 1000,
//         ethers.parseEther("0.01"),
//         ethers.ZeroAddress,
//         1
//     );

//     const auction =  await contract.auctions(0);
//     console.log("auction:", auction);
// });

describe("Test Upgrade", async () => {
    it("Should be able to deploy", async () => {
        //1.部署业务合约
        await deployments.fixture(["deployNftAuction"]);

        const nftAuctionProxy = await deployments.get("NftAuctionProxy");

        //2.调用createAuction方法创建拍卖
        const nftAuction = await ethers.getContractAt("NftAuction", nftAuctionProxy.address);
        nftAuction.createAuction(
            1000 * 1000,
            ethers.parseEther("0.01"),
            ethers.ZeroAddress,
            1
        );
        const auction = await nftAuction.auctions(0);
        console.log("创建拍卖成功：", auction);
        const implAddress1 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address);
        //3.升级合约
        await deployments.fixture(["upgradeNftAuction"]);
        const implAddress2 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address);

        //4.读取合约的auctions[0]
        const auction2 = await nftAuction.auctions(0);
        console.log("升级合约后读取拍卖信息：", auction2);
        console.log(implAddress1, implAddress2, '==implAddress==');
        const nftAuctionV2 = await ethers.getContractAt("NftAuctionV2", nftAuctionProxy.address);
        const hello = await nftAuctionV2.testHello();
        console.log("testHello:", hello);
        // expect(auction2.startPrice).to.equal(auction.startPrice);
        expect(implAddress1).to.not.equal(implAddress2);
    })
});