package main

import "fmt"

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	record := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var res []string
	var dataList [][]byte
	for _, digit := range []byte(digits) {
		tmp, ok := record[digit]
		if !ok {
			panic("invalid digit")
		}
		dataList = append(dataList, tmp)
	}
	build(dataList, &res, make([]byte, len(dataList)), 0, len(dataList))
	return res
}

func build(dataList [][]byte, res *[]string, tmp []byte, position, length int) {
	for _, d := range dataList[position] {
		tmp[position] = d
		if position == length-1 {
			*res = append(*res, string(tmp))
			continue
		}
		build(dataList, res, tmp, position+1, length)
	}
}
