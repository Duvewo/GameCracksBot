package commands

import tb "gopkg.in/tucnak/telebot.v2"

func OnStart(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {
		b.Send(m.Sender, "Hello!\nAvailable commands:\n/cracks - Show to you latest cracks")
	}
}