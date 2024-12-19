package number

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func intToRoman(num int) string {
//		arr := make([]string, 0)
//		n1000 := num / 1000
//		if n1000 > 0 {
//			arr = append(arr, strings.Repeat("M", n1000))
//			num -= 1000 * n1000
//		}
//		n100 := num / 100
//		if n100 > 0 {
//			if n100 > 5 {
//				if n100 > 8 {
//					arr = append(arr, strings.Repeat("C", 10-n100)+"M")
//				} else {
//					arr = append(arr, "D"+strings.Repeat("C", n100-5))
//				}
//			} else {
//				if n100 <= 3 {
//					arr = append(arr, strings.Repeat("C", n100))
//				} else {
//					arr = append(arr, strings.Repeat("C", 5-n100)+"D")
//				}
//			}
//			num -= n100 * 100
//		}
//		n10 := num / 10
//		if n10 > 0 {
//			if n10 > 5 {
//				if n10 > 8 {
//					arr = append(arr, strings.Repeat("X", 10-n10)+"C")
//				} else {
//					arr = append(arr, "L"+strings.Repeat("X", n10-5))
//				}
//			} else {
//				if n10 <= 3 {
//					arr = append(arr, strings.Repeat("X", n10))
//				} else {
//					arr = append(arr, strings.Repeat("X", 5-n10)+"L")
//				}
//			}
//			num -= n10 * 10
//		}
//		if num > 0 {
//			if num > 5 {
//				if num > 8 {
//					arr = append(arr, strings.Repeat("I", 10-num)+"X")
//				} else {
//					arr = append(arr, "V"+strings.Repeat("I", num-5))
//				}
//			} else {
//				if num <= 3 {
//					arr = append(arr, strings.Repeat("I", num))
//				} else {
//					arr = append(arr, strings.Repeat("I", 5-num)+"V")
//				}
//			}
//		}
//		return strings.Join(arr, "")
//	}
//
//	func intToRoman(num int) string {
//		arr := make([]string, 0)
//		nums := []int{3000, 2000, 1000, 900, 800, 700, 600, 500, 400, 300, 200, 100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
//		strs := []string{"MMM", "MM", "M", "CM", "DCCC", "DCC", "DC", "D", "CD", "CCC", "CC", "C", "XC", "LXXX", "LXX", "LX", "L", "XL", "XXX", "XX", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}
//		for i, n := range nums {
//			x := num - n
//			if x >= 0 {
//				arr = append(arr, strs[i])
//				num = x
//			}
//		}
//		return strings.Join(arr, "")
//	}
func intToRoman(num int) string {
	arr := make([]string, 0)
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i, n := range nums {
		for num >= n {
			arr = append(arr, strs[i])
			num -= n
		}
		if num == 0 {
			break
		}
	}
	return strings.Join(arr, "")
}

func TestT12(t *testing.T) {
	assert.Equal(t, "MMMDCCXLIX", intToRoman(3749))
	assert.Equal(t, "LVIII", intToRoman(58))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
	assert.Equal(t, "II", intToRoman(2))
	assert.Equal(t, "XIII", intToRoman(13))
}
