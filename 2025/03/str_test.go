package _03

import (
	"fmt"
	"testing"
)

/*
*

	给定只包括[,],(,),{,},[,]的字符串，判断字符串是否有效

满足：
1. 左括号用对用的右括号
2. 左括号以相同顺序闭合
3. 右括号对应左括号

思路：
1. 不是偶数的false
2. 两个的判断是否对称
3. 多个的分两种：a. 单个对称 b.多个对称
*/
func TestStr(t *testing.T) {
	fmt.Println(MS("([)]"))
}

var strM = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

var strList = []string{"()", "[]", "{}"}

func MS(str string) bool {
	if len(str) == 0 || len(str)%2 > 0 {
		return false
	}
	if len(str) == 2 {
		for _, v := range strList {
			if str == v {
				return true
			}
		}
	}
	m := make(map[string]int)

	for _, v := range str {
		m[string(v)] = m[string(v)] + 1
	}

	if len(m)%2 > 0 {
		return false
	}

	for k, v := range strM {
		mk, okk := m[k]
		mv, okv := m[v]
		if okk && okv && (mk+mv)%2 == 0 {
			return true
		}
	}
	return false
}
