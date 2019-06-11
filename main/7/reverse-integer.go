package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	if (x > 10) && (x%10 == 0) {
		reverse(x / 10)
	}

	negative := false
	if x < 0 {
		negative = true
		x = x * -1
	}
	s := strconv.Itoa(x)
	length := len(s)
	var bs []byte
	for i := length; i > 0; i-- {
		bs = append(bs, s[i-1])
	}

	const MaxUint = ^uint32(0)
	const MaxInt = int32(MaxUint >> 1)

	if string(bs) > strconv.FormatInt(int64(MaxInt), 10) && len(string(bs)) >= len(strconv.FormatInt(int64(MaxInt), 10)) {
		return 0
	}
	res, _ := strconv.Atoi(string(bs))
	if negative {
		return res * -1
	} else {
		return res
	}

}

func main() {

	fmt.Println(reverse(-123))
}
