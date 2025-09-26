package main

import (
	"fmt"
	"slices"
	"strconv"
)

// 1. 只出现一次的数字
func SingleNumber(nums []int) int {
	var mapHash = make(map[int]int)
	for _, num := range nums {
		_, ok := mapHash[num]
		if !ok {
			mapHash[num] = 0
		}
		mapHash[num]++
	}
	for key, value := range mapHash {
		if value == 1 {
			return key
		}
	}
	return 0
}

// 2. 回文数
func IsPalindromeNum(num int) bool {
	var str string = strconv.Itoa(num)
	if len(str) == 1 {
		fmt.Println("回文数")
		return true
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			fmt.Println("非回文数")
			return false
		}
	}
	fmt.Println("回文数")
	return true
}

// 3.有效的括号"(){}[]"或者"({})
func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			//左括号没放进去,说明顺序不对
			//栈顶元素和当前的不是一对，也不匹配，必须紧挨着
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// 4. 最长的公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 5. 加一
func AddOne(nums []int) []int {
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		if nums[i] != 9 {
			nums[i]++
			for j := i + 1; j < length; j++ {
				nums[j] = 0
			}
			return nums
		}
	}
	//全为9的情况
	nums = make([]int, length+1)
	nums[0] = 1
	return nums
}

// 6. 删除有序数组的重复项,返回最新长度
func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// 7.合并区间
func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { //可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1])
		} else { //不相交,无法合并
			ans = append(ans, p)
		}
	}
	return
}

// 8.两数之和
func TwoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, num := range nums {
		if index, ok := hashMap[target-num]; ok {
			return []int{index, i}
		}
		hashMap[num] = i
	}
	return nil
}

func main() {
	ret := SingleNumber([]int{1, 3, 1, 3, 2, 5, 6, 6, 5})
	fmt.Println("寻找只出现一次的数字: ", ret)
	ret2 := IsPalindromeNum(12321)
	fmt.Println("是否为回文数: ", ret2)
	ret3 := IsValid("([])")
	fmt.Println("是否为正确的括号: ", ret3)
	ret4 := LongestCommonPrefix([]string{"abc", "abcd", "abcde"})
	fmt.Println("最长公共前缀: ", ret4)
	ret5 := AddOne([]int{3, 5, 8, 9, 8, 9})
	fmt.Println("加一: ", ret5)
	ret6 := RemoveDuplicates([]int{1, 3, 3, 5, 6, 6, 6, 7})
	fmt.Println("移除重复项: ", ret6)
	ret7 := merge([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}})
	fmt.Println("合并区间: ", ret7)
	ret8 := TwoSum([]int{8, 6, 2, 11}, 8)
	fmt.Println("两数之和: ", ret8)
}
