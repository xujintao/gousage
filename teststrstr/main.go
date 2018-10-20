package main

import (
	"fmt"
)

//判断str2是否为str1的子集
func strstr(str1 string, str2 string) bool {
	s1 := []byte(str1)
	s2 := []byte(str2)

	for len(s1) > 0 && len(s2) > 0 && len(s1) >= len(s2) {
		for n := 0; n <= len(s2)-1 && s1[n] == s2[n]; n++ {
			if len(s2)-1 == n {
				return true
			}
		}

		s1 = s1[1:]
	}
	return false
}

func main() {
	fmt.Println(strstr("abcdef", "def"))
	fmt.Println(strstr("abcdef", "deg"))
}
