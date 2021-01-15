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
	var res int = num * num

	ioutil.WriteFile("OUTPUT.txt", []byte(strconv.Itoa(res)), 0755)
}
