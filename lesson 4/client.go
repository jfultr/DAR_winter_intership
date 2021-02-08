package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./lib"
)

const (
	auth = "http://127.0.0.1:8001"
)

func main() {
	clt := http.Client{}

	var urls []string

	urls = append(urls, "http://youtube.com/")
	urls = append(urls, "http://google.com/")
	urls = append(urls, "http://twitch.com/")

	message := &lib.SaveJSON{
		Urls: urls,
	}

	json, err := json.Marshal(message)

	request, err := http.NewRequest("POST", auth+"/parseAsync", bytes.NewBuffer(json))
	if err != nil {
		panic("Error on new request:" + err.Error())
	}
	resp, err := clt.Do(request)
	if err != nil {
		panic("Error on get request:" + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error on read body:" + err.Error())
	}

	fmt.Println(string(body))

}
