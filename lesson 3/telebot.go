package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"./lib"
)

const (
	auth = "https://api.telegram.org/bot906366127:AAF-nQkH3DFqKXlAhQ3ChkNbh49Itw62ji4"
)

func getUpdates(clt http.Client) ([]byte, error) {
	request, err := http.NewRequest("POST", auth+"/getUpdates", nil)
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
	return body, err
}

func sendMessage(clt http.Client, ID int, answer string) {
	req, _ := http.NewRequest("GET", auth+"/sendMessage?text="+answer+"&chat_id="+fmt.Sprint(ID), nil)
	clt.Do(req)

}

// ParseRasponse model returns response object
func parseRasponse(body []byte) (*lib.GetUpdatesResponse, error) {
	b := &lib.GetUpdatesResponse{}
	if err := json.Unmarshal(body, b); err != nil {
		panic("Error on parsing body:" + err.Error())
	}
	return b, nil
}

func isAnswered(update lib.Update, answered map[int][]int) bool {
	messagesIDs, faund := answered[update.Message.Chat.ID]
	if faund {
		for _, id := range messagesIDs {
			if id == update.Message.MessageID {
				return true
			}
		}
	}
	return false
}

func tweakAnswer(update lib.Update, answered *map[int][]int) string {
	switch text := update.Message.Text; text {
	case "/start":
		return "Рабочие фразы: 1.привет 2.как дела? 3.ты живой?"
	case "/help":
		return "Напиши: 1.привет 2.как дела? 3.ты живой?"
	case "привет":
		return "Добрый день!"
	case "как дела?":
		return "Хорошо, ведь пока мой скрипт работает"
	case "ты живой?":
		return "да, пока работает скрипт"
	default:
		return "извини, я не понимаю"
	}
}

func readUpdates(clt http.Client, resp *lib.GetUpdatesResponse, answered map[int][]int) {
	for _, update := range resp.Result {
		if !isAnswered(update, answered) {
			answered[update.Message.Chat.ID] = append(answered[update.Message.Chat.ID], update.Message.MessageID)
			answer := tweakAnswer(update, &answered)
			sendMessage(clt, update.Message.Chat.ID, answer)
		}
	}

}
func main() {
	clt := http.Client{}
	answered := make(map[int][]int)
	for {
		body, _ := getUpdates(clt)
		resp, _ := parseRasponse(body)
		readUpdates(clt, resp, answered)
	}
}
