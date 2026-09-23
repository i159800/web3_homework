// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import { ERC721Royalty } from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Royalty.sol";

contract Token_2982 is ERC721, ERC721Royalty {
    constructor() ERC721("MyNFT", "MNFT") {
        // 设置默认版税：给 msg.sender，比例为 5% (500 / 10000)
        _setDefaultRoyalty(msg.sender, 500);
    }
    
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Royalty)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}