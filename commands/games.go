package commands

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

var replyMarkup = &tb.ReplyMarkup{}

var SortByButton = replyMarkup.Data("Sort By", "sort", "sort")
var InvertSortButton = replyMarkup.Data("Invert sort", "invertsort", "invertsort")
var IsTripleAButton = replyMarkup.Data("Is AAA?", "isaaa", "isaaa")
var IsReleasedButton = replyMarkup.Data("Is Released?", "isreleased", "isreleased")
var IsCrackedButton = replyMarkup.Data("Is Cracked?", "iscracked", "iscracked")
var ShowGamesButton = replyMarkup.Data("Search", "search", "search")

func OnGames(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {
		replyMarkup.Inline(replyMarkup.Row(SortByButton, InvertSortButton, IsCrackedButton, IsReleasedButton, IsTripleAButton, ShowGamesButton))
		_, err := b.Send(m.Sender, "Search games by filter (<b>Sort By</b>: Release Date, <b>Invert</b>: No, <b>Is AAA</b>: No, <b>Is Released</b>: No, <b>Is Cracked</b>: No)", replyMarkup, tb.ModeHTML)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}