package endpoints

import (
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"time"
)

type Response []struct {
	ID             string        `json:"_id"`
	Title          string        `json:"title"`
	Slug           string        `json:"slug"`
	ReleaseDate    string    	`json:"releaseDate"`
	Protections    []string      `json:"protections"`
	Versions       []interface{} `json:"versions"`
	Image          string        `json:"image"`
	ImagePoster    string        `json:"imagePoster"`
	IsAAA          bool          `json:"isAAA"`
	NFOsCount      int           `json:"NFOsCount"`
	CommentsCount  int           `json:"commentsCount"`
	FollowersCount int           `json:"followersCount"`
	Groups         []interface{} `json:"groups"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	URL            string        `json:"url"`
}

func OnSearchButton(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {

		var sortBy string
		var invert string
		var isAAA string
		var isReleased string
		var isCracked string

		var url = url2.Values{}
		var re = regexp.MustCompile(`(?m)Search games by filter \(Sort By: (.*), Invert: (.*), Is AAA: (.*), Is Released: (.*), Is Cracked: (.*)\)`)

		for i, match := range re.FindStringSubmatch(c.Message.Text) {
			if i == 1 {
				sortBy = match
			} else if i == 2 {
				invert = match
			} else if i == 3 {
				isAAA = match
			} else if i == 4 {
				isReleased = match
			} else if i == 5 {
				isCracked = match
			}
		}

		if sortBy == "Release Date" {
			url.Add("sort_by", "release_date")
		} else {
			url.Add("sort_by", "crack_date")
		}

		if invert == "Yes" {
			url.Add("is_sort_inverted", "true")
		}

		if isAAA == "Yes" {
			url.Add("is_aaa", "true")
		}

		if isReleased == "Yes" {
			url.Add("is_released", "true")
		}

		if isCracked == "Yes" {
			url.Add("is_cracked", "true")
		}



		resp, err := http.Get("https://api.crackwatch.com/api/games?" + url.Encode())
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

		var message = "Search Result:\n"

		for _, value := range response {
			if isReleased == "No" {
				message += " <b>" + value.Title + "</b> Release date: " + value.ReleaseDate + ", Followers: " + strconv.Itoa(value.FollowersCount) + ", Link: " + value.URL + "\n "
			} else {
				message += " <b>" + value.Title + "</b> Released, Followers: " + strconv.Itoa(value.FollowersCount) + ", Link: " + value.URL + "\n "
			}
		}

		_, err = b.Send(c.Sender, message, tb.ModeHTML)

		if err != nil {
			fmt.Println(err)
			return
		}

	}
}