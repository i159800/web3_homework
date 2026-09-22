import hardhatToolboxViem from "@nomicfoundation/hardhat-toolbox-viem";
import { configVariable, defineConfig } from "hardhat/config";

import hardhatEthers from "@nomicfoundation/hardhat-ethers";
import hardhatTypechain from "@nomicfoundation/hardhat-typechain";
import hardhatMocha from "@nomicfoundation/hardhat-mocha";
import hardhatEthersChaiMatchers from "@nomicfoundation/hardhat-ethers-chai-matchers";
import hardhatNetworkHelpers from "@nomicfoundation/hardhat-network-helpers";
import hardhatUpgrades from '@openzeppelin/hardhat-upgrades';


export default defineConfig({
  plugins: [hardhatToolboxViem,
    hardhatEthers,
    hardhatTypechain,
    hardhatMocha,
    hardhatEthersChaiMatchers,
    hardhatNetworkHelpers,
    hardhatUpgrades
  ],
  
  solidity: {
    profiles: {
      default: {
        version: "0.8.28",
          settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
          viaIR: true,
        },
      },
      production: {
        version: "0.8.28",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
          viaIR: true,
        },
      },
    },
  },
  networks: {
    hardhatMainnet: {
      type: "edr-simulated",
      chainType: "l1",
       allowUnlimitedContractSize:true
    },
    hardhatOp: {
      type: "edr-simulated",
      chainType: "op",
       allowUnlimitedContractSize:true
    },
    sepolia: {
      type: "http",
      chainType: "l1",
      url: configVariable("SEPOLIA_RPC_URL"),
      accounts: [configVariable("SEPOLIA_PRIVATE_KEY")],
    },
    hardhat:{
      type:"edr-simulated",
      chainType:"l1",
      allowUnlimitedContractSize:true
    },
    default: {
      type: "edr-simulated",
      allowUnlimitedContractSize: true,
    }
  }
});
