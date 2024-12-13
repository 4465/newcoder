package main

import (
	"fmt"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @param n int整型
 * @return string字符串
 */

/**
 * 输入："This is a sample",16
 * 返回值："SAMPLE A IS tHIS"
 */
func trans(s string, n int) string {
	// write code here
	items := strings.Split(s, " ")
	resLst := []string{}
	for i := len(items) - 1; i >= 0; i-- {
		a := []rune(items[i])
		for i := 0; i < len(a); i++ {
			if a[i] >= 'a' && a[i] <= 'z' {
				a[i] = a[i] - 32
			} else if a[i] >= 'A' && a[i] <= 'Z' {
				a[i] = a[i] + 32
			}
		}
		resLst = append(resLst, string(a))
	}

	return strings.Join(resLst, " ")
}

func main() {
	a := "nowcoder"
	res := trans(a, 8)
	fmt.Println(res)
}
