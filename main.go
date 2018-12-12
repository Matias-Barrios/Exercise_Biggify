package main

import (
	"fmt"
	"strconv"
)

type bigdigits struct {
	values map[string][]string
}

var b = bigdigits{values: map[string][]string{

	"0": {" ***** ",
		" *   * ",
		" *   * ",
		" *   * ",
		" ***** "},

	"1": {"   *   ",
		"   *   ",
		"   *   ",
		"   *   ",
		"   *   "},
	"2": {" ***** ",
		"     * ",
		"   **  ",
		" **    ",
		" ***** "},
	"3": {" ***** ",
		"     * ",
		"  **** ",
		"     * ",
		" ***** "},
	"4": {" *   * ",
		" *   * ",
		" ***** ",
		"     * ",
		"     * "},
	"5": {"*****",
		"*    ",
		"*****",
		"    *",
		"*****"},
	"6": {" ***** ",
		" *     ",
		" ***** ",
		" *   * ",
		" ***** "},
	"7": {" ***** ",
		"     * ",
		"     * ",
		"     * ",
		"     * "},
	"8": {"*****",
		"*   *",
		"*****",
		"*   *",
		"*****"},
	"9": {" ***** ",
		" *   * ",
		" ***** ",
		"     * ",
		" ***** "},
},
}

func bigify(i int) {
	for r := 0; r <= 4; r++ {
		for _, v := range strconv.Itoa(i) {
			fmt.Print(b.values[string(v)][r])
		}
		fmt.Print("\n")
	}
}

func main() {
	bigify(1234567890)
}
