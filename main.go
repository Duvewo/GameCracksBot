package main

import (
	"GameCracksBot/commands"
	"GameCracksBot/endpoints"
	"GameCracksBot/helpers"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main(){
	b, err := tb.NewBot(tb.Settings{Token: "1488843849:AAFtvLZw_tRIdW91I4tcpwyMSkbSiCZZQI8", Poller: &tb.LongPoller{Timeout: 10}})

	if err != nil {
		fmt.Println(err)
		return
	}

	b.Handle("/start", commands.OnStart(b))
	b.Handle("/cracks", commands.OnCracks(b))
	b.Handle("/games", commands.OnGames(b))

	b.Handle(&commands.ShowGamesButton, endpoints.OnSearchButton(b))

	//filter
	b.Handle(&commands.IsTripleAButton, helpers.IsTripleA(b))
	b.Handle(&commands.InvertSortButton, helpers.IsInvertedSort(b))
	b.Handle(&commands.IsReleasedButton, helpers.IsReleased(b))
	b.Handle(&commands.IsCrackedButton, helpers.IsCracked(b))
	b.Handle(&commands.SortByButton, helpers.SortFilter(b))

	b.Start()

}