const hre = require("hardhat");
const {expect} = require("chai");
describe("MyToken Test", async () => {
    const {ethers} = hre;
    const initialSupply = 10000;
    let MyTokenContract;
    let account1, account2;
    beforeEach(async () => {
        [account1, account2] = await ethers.getSigners();
        console.log(account1, account2, '==accounts==');
        const MyToken = await ethers.getContractFactory("MyToken");
        MyTokenContract = await MyToken.connect(account2).deploy(initialSupply);
        MyTokenContract.waitForDeployment();
        const contractAddress = await MyTokenContract.getAddress();
        expect(contractAddress).to.length.greaterThan(0);
    });

    it("验证合约的name,symbol,decimal", async () => {
        const name = await MyTokenContract.name();
        const symbol = await MyTokenContract.symbol();
        const decimal = await MyTokenContract.decimals();
        expect(name).to.equal("MyToken");
        expect(symbol).to.equal("MTK");
        expect(decimal).to.equal(18);
    });

    it("测试转账", async () => {
        const resp = await MyTokenContract.transfer(account1, initialSupply / 2);
        console.log(resp);
        const balanceOfAccount2 = await MyTokenContract.balanceOf(account2);
        expect(balanceOfAccount2).to.equal(initialSupply / 2);
    });
});