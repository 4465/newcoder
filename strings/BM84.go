package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param strs string字符串一维数组
 * @return string字符串
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	// write code here
	if len(strs) == 0 {
		return ""
	}
	var minLen int = len(strs[0])
	for _, str := range strs {
		minLen = min(minLen, len(str))
	}
	fmt.Println(minLen)
	var i int
	for i = 0; i < minLen; i++ {
		// 当最长公共前缀为最短的字符串时， 循环是不会走到return的，所以i的生命周期要扩大
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:i]
}

func main() {
	strs := []string{"flower", "flow", "flowht"}
	fmt.Println(longestCommonPrefix(strs))
}
