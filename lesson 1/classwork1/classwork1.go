package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("INPUT.txt")
	data := make([]byte, 100)
	count, _ := file.Read(data)

	str := string(data[:count])
	nums := strings.Fields(str)

	sum := 0
	for _, v := range nums {
		v, _ := strconv.Atoi(v)
		sum += v
	}
	ssum := fmt.Sprint(sum)

	ioutil.WriteFile("OUTPUT.txt", []byte(ssum), 0755)
}
