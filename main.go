package main

import (
	"fmt"
	"strings"
)

func main() {
	var init = make([][]string, 4)
	for y := 0; y < 4; y++ {
		var tmp = make([]string, 4)
		for x := 0; x < 4; x++ {
			tmp[x] = " "
		}
		init[y] = tmp
	}
	init[0][2] = "■"
	var res string
	for _, t := range init {
		res = res + strings.Join(t, "") + "\n"
	}
	fmt.Printf("%v", strings.TrimRight(res, "\n"))
}
