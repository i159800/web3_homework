// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
contract MergeSortedArray {
    uint[] public arr1 = [1, 3, 5];
    uint[] public arr2 = [2, 4, 6];
    uint[] public arr;

    function merge() public returns (uint[] memory) {
        uint p1 = 0;
        uint p2 = 0;
        while(true){
            if (p1 == arr1.length) {
                while (p2 < arr2.length) {
                    arr.push(arr2[p2]);
                    p2++;
                }
                break;
            }
            if (p2 == arr2.length) {
                while (p1 < arr1.length) {
                    arr.push(arr1[p1]);
                    p1++;
                }
                break;
            }
            if (arr1[p1] < arr2[p2]) {
                arr.push(arr1[p1]);
                p1++;
            } else {
                arr.push(arr2[p2]);
                p2++;
            }
        }

        return arr;
    }
}