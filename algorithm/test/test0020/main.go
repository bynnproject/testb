package main

import "fmt"

func main() {
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	tmp := make([]byte, len(s))
	m := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	length := 0
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if !ok {
			tmp[length] = s[i]
			length++
			continue
		}
		if length == 0 {
			return false
		}
		if tmp[length-1] != v {
			return false
		}
		length--
	}
	if length != 0 {
		return false
	}
	return true
}
