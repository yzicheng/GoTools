package mid

import (
	"strings"
	"testing"
)

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
//
// 示例 1：
//
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// 示例 3：
//
// 输入：s = "A", numRows = 1
// 输出："A"
func Test6(t *testing.T) {
	tests := []struct {
		name string
		s    string
		num  int
		want string
	}{
		{
			name: "case 1",
			s:    "PAYPALISHIRING",
			num:  3,
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "case 2",
			s:    "PAYPALISHIRING",
			num:  4,
			want: "PINALSIGYAHRPI",
		},
		{
			name: "case 3",
			s:    "A",
			num:  1,
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := convert(tt.s, tt.num)
			if l != tt.want {
				t.Errorf("convert() = %v, want %v", l, tt.want)
			}
		})
	}
}

// {
// name: "case 1",
// s:    "PAYPALISHIRING",
// num:  3,
// want: "PAHNAPLSIIGYIR",
// },
func convert(s string, numRows int) string {
	strMap := make(map[int][]string)
	if numRows == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i%(2*numRows-2) < numRows {
			strMap[i%(2*numRows-2)] = append(strMap[i%(2*numRows-2)], string(s[i]))
		} else {
			strMap[2*(numRows-1)-i%(2*numRows-2)] = append(strMap[2*(numRows-1)-i%(2*numRows-2)], string(s[i]))
		}
	}
	res := ""
	for i := 0; i < numRows; i++ {
		res += strings.Join(strMap[i], "")
	}
	return res
}
