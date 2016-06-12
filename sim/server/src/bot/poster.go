package bot

import (
	"model"
)

type (
	// poster
	// Inputに入ったメッセージをPostする
	poster struct {
		Input chan model.Message
	}
)

func (p *poster) Run() {
	url := "http://localhost:1323/messages"
	for m := range p.Input {
		output := model.Message{}
		go postJson(url, m, &output)
	}
}

// posterの生成はこの関数を使う
func NewPoster(buffer_size int) *poster {
	ch := make(chan model.Message, buffer_size)
	return &poster{
		Input: ch,
	}
}
