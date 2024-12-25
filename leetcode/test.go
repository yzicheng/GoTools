package leetcode

import (
	list2 "container/list"
	"errors"
	"strings"
)

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
func IP(s string) {
	stack := list2.New()
	//strconv.Itoa(1)
	//strconv.Atoi("1")
	stack.PushBack(strList[0])
	i := 1
	for stack.Len() > 0 {
		back := stack.Back().Value.(string)
	}
	ch := make(chan struct{})
	if ch == nil {

	}
}

// ActiveIP 每个整数位于 0 到 255 之间组成，且不能含有前导 0
// "0.011.255.245"
// "192.168.1.312"
func ActiveIP(s string) error {
	ips := strings.Split(s, ".")
	if len(ips) != 4 {
		return errors.New("err ip")
	}
	//for i := 1; i < len(ips); i++ {
	//	ips[]
	//}
}
