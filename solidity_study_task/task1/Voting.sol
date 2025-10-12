// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract Voting {
    //定义一个投票的映射
    mapping(string username => uint256 amount) public userVoteMap;
    //定义一个动态数组，用来存储当前的用户列表
    string[] userlist;
    //进行投票
    function vote(string memory username) public {
        uint voteNum = getVotes(username);
        if (voteNum <= 0) {
            userlist.push(username);
        }
        userVoteMap[username] += 1;
    }
    //获得某个用户的投票
    function getVotes(string memory username) public view returns (uint256) {
        return userVoteMap[username];
    }

    //获得所有投票的用户
    function getUserlist() public view returns (string[] memory) {
        return userlist;
    }

    //重置所有用户的投票
    function resetVotes() public {
        uint len = userlist.length;
        for (uint i = 0; i < len; i++) {
            userVoteMap[userlist[i]] = 0;
        }
    }
}