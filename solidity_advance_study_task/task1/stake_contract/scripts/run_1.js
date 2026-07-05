import hre from "hardhat";

const { ethers, provider } = await hre.network.create();

const address = "0xe7f1725e7734ce288f8367e1bb143e90bb3f0512";
const contract = await ethers.getContractAt("MetaNodeToken", address);

const addressStake = "0x5fbdb2315678afecb367f032d93f642f64180aa3";
const contractStake = await ethers.getContractAt("MetaNodeStake", addressStake);

const startBlock = await contractStake.startBlock();
console.log("startBlock::", startBlock.toString());