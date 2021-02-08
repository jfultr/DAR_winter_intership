package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"./lib"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func getRequest(url string) {
	clt := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Error on new request:" + err.Error())
	}
	resp, err := clt.Do(request)
	if err != nil {
		panic("Error on get request:" + err.Error())
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error on read body:" + err.Error())
	}

	fmt.Println(url)
}

func readReq(r *http.Request) *lib.SaveJSON {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Error on read body:" + err.Error())
	}
	b := &lib.SaveJSON{}
	if err := json.Unmarshal(body, b); err != nil {
		panic("Error on parsing body:" + err.Error())
	}
	return b
}

func parseSync(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	b := readReq(r)

	for _, url := range b.Urls {
		getRequest(url)
	}

	fmt.Fprintf(w, "time spend: %s", time.Since(start))
}

func parseAsync(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	b := readReq(r)

	for _, url := range b.Urls {
		go getRequest(url)
	}

	fmt.Fprintf(w, "time spend: %s", time.Since(start))
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/parseSync", parseSync)
	http.HandleFunc("/parseAsync", parseAsync)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
