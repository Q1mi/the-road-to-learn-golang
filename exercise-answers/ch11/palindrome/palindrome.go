package palindrome

import "unicode"

/*
编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
（回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等。）
*/

// Judge 校验输入的字符串是否为回文字符串
func Judge(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		return letters[i] == letters[j]
	}

	return true
}
