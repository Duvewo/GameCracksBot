package commands

import (
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"
	"time"
)

type Response []struct {
	GroupName string    `json:"groupName"`
	Date      time.Time `json:"date"`
	Title     string    `json:"title"`
	Size      string    `json:"size"`
	Image     string    `json:"image"`
}

func OnCracks(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {
		var resp *http.Response
		var err error

		if len(m.Payload) == 0 {
			resp, err = http.Get("https://api.crackwatch.com/api/cracks")
		} else {
			resp, err = http.Get("https://api.crackwatch.com/api/cracks?page=" + m.Payload)
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			return
		}

		var response Response
		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			fmt.Println(err)
			return
		}

		var message = "Latest cracks:\n"

		for _, value := range response {
			message += " <b>" + value.Title + "</b> (" + value.GroupName + "): " + value.Date.String() + "\n"
		}

		_, err = b.Send(m.Sender, message, tb.ModeHTML)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}