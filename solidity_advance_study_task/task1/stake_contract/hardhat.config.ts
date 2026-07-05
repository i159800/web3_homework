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