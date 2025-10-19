// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract Erc721_Test is ERC721, Ownable {
    uint256 private _nextTokenId = 1;

    constructor(address initialOwner)
        ERC721("MyToken", "ZXH")
        Ownable(initialOwner)
    {}

    function _baseURI() internal pure override returns (string memory) {
        return "ipfs://bafybeifne3cnjwp764w5gcenpgmdx5aaznkbpekaworbc34ogog3sczhoy/";
    }

    function safeMint(address to) public onlyOwner returns (uint256) {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
        return tokenId;
    }
}