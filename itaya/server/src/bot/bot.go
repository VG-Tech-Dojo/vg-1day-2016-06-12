package bot

import (
	"model"
)

type (
	// bot
	//
	// Inputに入ってきたmessageが
	// Checkerを満たす場合
	// Processerで追加投稿用messageを作成し
	// Outputに流す
	bot struct {
		Input     chan model.Message
		Output    chan model.Message
		Checker   Checker
		Processer Processer
	}
)

func (b bot) Run() {
	for m := range b.Input {
		if !b.Checker.Check(m) {
			continue
		}
		msg := b.Processer.Process(&m)
		if msg == nil {
			continue
		}
		b.Output <- *msg
	}
}

func NewBot(c Checker, p Processer, out chan model.Message) (b *bot) {
	ch := make(chan model.Message)
	return &bot{
		Input:     ch,
		Output:    out,
		Checker:   c,
		Processer: p,
	}
}

func NewGreetBot(name string, pattern string, out chan model.Message) (b *bot) {
	checker := &RegexpChecker{
		Pattern: pattern,
	}
	processer := &GreetProcesser{
		Name: name,
	}
	return NewBot(checker, processer, out)
}

func NewTimelineBot(out chan model.Message) (b *bot) {
	checker := &RegexpChecker{
		Pattern: `^timeline`,
	}
	processer := &TimelineProcesser{}
	processer.Init()
	return NewBot(checker, processer, out)
}
