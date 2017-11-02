/*
Package gofcnreadarray to read a simple array with float 64 values from a file or through the terminal.
*/
package gofcnreadarray

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Termread reads an array through the terminal, where the elements are separated by space.

Example:

# Enter array:
# 2.4 6.8 7.2 10.6 5.9

return:
a []float64: array
*/
func Termread() []float64 {
	var a []float64
	var line []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter array: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	line = strings.Split(string(text), " ")
	a = lineToFloat(line)

	return a
}

/*
lineToFloat transforms an array of strings into an array of float64

param:
line: []string

return:
s: []float64
*/
func lineToFloat(line []string) []float64 {
	var s = []float64{}

	for _, l := range line {
		num, err := strconv.ParseFloat(l, 64)
		if err != nil {
			panic(err)
		}
		s = append(s, num)
	}

	return s
}
