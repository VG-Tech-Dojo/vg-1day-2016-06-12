package bot

import (
	"model"
)

type (
	// broadcaster
	//
	// 一つのチャンネルを利用して複数botを同時に動かすためのhelper.
	// MsgInputに入ってきたmessageを
	// 各Membersに対し伝達する
	// Membersへの登録は EntryInputで可能
	broadcaster struct {
		MsgInput   chan model.Message
		EntryInput chan *bot
		Members    map[*bot]bool
	}
)

func (b *broadcaster) Run() {
	for {
		select {
		case bot := <-b.EntryInput:
			b.Members[bot] = true
		case msg := <-b.MsgInput:
			for bot, _ := range b.Members {
				bot.Input <- msg
			}
		}
	}
}

// broadcasterの生成はこの関数を使う
func NewBroadcaster() *broadcaster {
	msgIn := make(chan model.Message)
	memberIn := make(chan *bot)
	return &broadcaster{
		MsgInput:   msgIn,
		EntryInput: memberIn,
		Members:    make(map[*bot]bool),
	}
}
