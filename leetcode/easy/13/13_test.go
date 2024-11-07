package _3

import (
	"fmt"
	"testing"
)

//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给定一个罗马数字，将其转换成整数。
//
//
//
//示例 1:
//
//输入: s = "III"
//输出: 3
//示例 2:
//
//输入: s = "IV"
//输出: 4
//示例 3:
//
//输入: s = "IX"
//输出: 9
//示例 4:
//
//输入: s = "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//示例 5:
//
//输入: s = "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
//

func Test(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "III",
			want: 3,
		},
		{
			name: "case 2",
			s:    "IV",
			want: 4,
		},
		{
			name: "case 3",
			s:    "IX",
			want: 9,
		},
		{
			name: "case 4",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "case 5",
			s:    "MCMXCIV",
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := romanToInt(tt.s)
			if res != tt.want {
				t.Errorf("err %d,  %d", res, tt.want)
			}
		})
	}
}

func romanToInt(s string) int {
	//I             1
	//V             5
	//X             10
	//L             50
	//C             100
	//D             500
	//M             1000
	//"MCMXCIV"
	//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100,
		"D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40,
		"XC": 90, "CD": 400, "CM": 900}
	sum := 0

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		last := ""
		if i+1 < len(s) {
			last = string(s[i+1])
		}
		if romanInt, ok := roman[v+last]; ok {
			fmt.Println(v + last)
			sum += romanInt
			i++
		} else if romanInt, ok := roman[v]; ok {
			fmt.Println(v + last)
			sum += romanInt
		}
	}
	return sum
}
