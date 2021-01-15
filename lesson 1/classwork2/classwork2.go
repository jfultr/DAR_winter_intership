package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, _ := file.Read(data)
	fmt.Printf("bytes: %q\n", data[:count])

	num, _ := strconv.Atoi(string(data[:count]))
	var res int = 0
	for i := 0; i <= num; i++ {
		res += i
	}

	ioutil.WriteFile("OUTPUT.txt", []byte(strconv.Itoa(res)), 0755)
}
