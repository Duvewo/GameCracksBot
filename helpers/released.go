package helpers

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
)

var replyMarkup = &tb.ReplyMarkup{}

var SortByButton = replyMarkup.Data("Sort By", "sort", "sort")
var InvertSortButton = replyMarkup.Data("Invert sort", "invertsort", "invertsort")
var IsTripleAButton = replyMarkup.Data("Is AAA?", "isaaa", "isaaa")
var IsReleasedButton = replyMarkup.Data("Is Released?", "isreleased", "isreleased")
var IsCrackedButton = replyMarkup.Data("Is Cracked?", "iscracked", "iscracked")
var ShowGamesButton = replyMarkup.Data("Search", "search", "search")

func IsReleased(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {
		var sortBy string
		var invert string
		var isAAA string
		var isReleased string
		var isCracked string

		replyMarkup.Inline(replyMarkup.Row(SortByButton, InvertSortButton, IsCrackedButton, IsReleasedButton, IsTripleAButton, ShowGamesButton))

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

		if isReleased == "No" {
			b.Edit(c.Message, "Search games by filter (<b>Sort By</b>: "+sortBy+", <b>Invert</b>: "+invert+", <b>Is AAA</b>: "+isAAA+", <b>Is Released</b>: Yes, <b>Is Cracked</b>: "+isCracked+")", replyMarkup, tb.ModeHTML)
		} else {
			b.Edit(c.Message, "Search games by filter (<b>Sort By</b>: "+sortBy+", <b>Invert</b>: "+invert+", <b>Is AAA</b>: "+isAAA+", <b>Is Released</b>: No, <b>Is Cracked</b>: "+isCracked+")", replyMarkup, tb.ModeHTML)
		}
	}
}