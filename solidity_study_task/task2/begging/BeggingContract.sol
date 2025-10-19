// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract {
    mapping(address => uint256) private donations;
    address private owner;
    address[] private donors;

    uint256 public startTime;
    uint256 public endTime;

    event Donation(address indexed donor, uint256 amount);

    modifier onlyOwner() {
        require(msg.sender == owner, "no owner");
        _;
    }

    constructor(uint _startTime, uint _endTime) {
        require(_endTime > _startTime, "time invalid");
        owner = msg.sender;
        startTime = _startTime;
        endTime = _endTime;
    }

    //捐赠
    function donate() external payable {
        require(
            block.timestamp >= startTime && block.timestamp <= endTime,
            "time expired or not start"
        );
        require(msg.value > 0, "donate value must > 0");

        if (donations[msg.sender] == 0) {
            donors.push(msg.sender);
        }
        donations[msg.sender] += msg.value;

        emit Donation(msg.sender, msg.value);
    }

    //查询某个地址的捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    //合约所有者提取所有资金
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "no balance");
        payable(owner).transfer(balance);
    }

    //获取合约当前余额
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}