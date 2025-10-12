// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract BinarySearch {
    uint[] arr = [1, 3, 5, 7, 9, 11, 13];

    function binarySearch(uint target) public view returns (int) {
        uint left = 0;
        uint right = arr.length - 1;
        while (left <= right) {
            uint mid = (left + right) / 2;
            if (arr[mid] == target) {
                return int(mid);
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
}