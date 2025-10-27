// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract NftAuctionV2 is Initializable{
    //结构体
    struct Auction {
        address seller; //卖家
        uint256 duration; //拍卖持续时间
        uint256 startPrice; //起始价格
        uint256 startTime; //拍卖开始时间
        bool ended; //是否结束
        address highestBidder; //最高出价者
        uint256 highestBid; //最高价格
        address nftContract; //NFT合约地址
        uint256 nftId; //NFT ID
    }

    //状态变量
    mapping(uint256 => Auction) public auctions;
    //下一个拍卖ID
    uint256 public nextAuctionId;
    //管理员地址
    address public admin;

    function initialize() public initializer {
        admin = msg.sender;
    }

    //创建拍卖
    function createAuction(uint256 _duration, uint256 _startPrice, address _nftAddress, uint256 _tokenId) public {
        //只有管理员才可以创建拍卖
        require(msg.sender == admin, "Only admin can create auction");
        //检查参数
        require(_duration > 1000 * 60, "Duration must be greater than 0");
        require(_startPrice > 0, "Start price must be greater than 0");
        auctions[nextAuctionId] = Auction ({
            seller:msg.sender,
            duration:_duration,
            startPrice:_startPrice,
            ended:false,
            highestBidder:address(0),
            highestBid:0,
            startTime:block.timestamp,
            nftContract:_nftAddress,
            nftId:_tokenId
        });
        nextAuctionId++;
    }

    //买家参与买单
    function placeBid(uint256 _auctionID) external payable {
        Auction storage auction = auctions[_auctionID];
        //检查拍卖是否结束
        require(block.timestamp < auction.startTime + auction.duration, "Auction has ended");
        //检查出价是否高于起始价格和当前最高价
        require(msg.value >= auction.startPrice, "Bid must be at least the start price");
        require(msg.value > auction.highestBid, "Bid must be higher than current highest bid");

        //如果有之前的最高出价者，退还其出价
        if (auction.highestBidder != address(0)) {
            payable(auction.highestBidder).transfer(auction.highestBid);
        }

        //更新最高出价者和最高价
        auction.highestBidder = msg.sender;
        auction.highestBid = msg.value;
    }

    function testHello() public pure returns (string memory) {
        return "Hello, World!";
    }
}