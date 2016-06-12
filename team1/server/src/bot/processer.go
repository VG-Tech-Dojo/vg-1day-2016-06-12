package bot

import (
	"fmt"
	"math/rand"
	"model"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type (
	// Processer
	// もととなるmessageをもとに，新規投稿するためのmessageを作成する
	Processer interface {
		Process(*model.Message) *model.Message
	}

	// EchoProcesser
	// 入力値のテキストそのままで返すProcesser
	EchoProcesser struct {
	}

	// Greetprocesser
	// 自身の名前で挨拶するProcesser
	GreetProcesser struct {
		Name string
	}

	// TimelineProcesser
	// homeのtimelineのtweetを1つ取得するProcesser
	TimelineProcesser struct {
		Api *anaconda.TwitterApi
	}

	// UranaiProcesser
	// 占い結果を返す
	UranaiProcesser struct {
	}

	// WarikanProcesser
	// 割り勘の結果を返す
	WarikanProcesser struct {
	}
)

func (p *EchoProcesser) Process(msgIn *model.Message) *model.Message {
	return &model.Message{Body: "[echo] " + msgIn.Body}
}

func (p *GreetProcesser) Process(msgIn *model.Message) *model.Message {
	txt := "[greet] nice to meet you! my name is " + p.Name
	return &model.Message{Body: txt}
}

func (p *TimelineProcesser) Init() {
	p.Api = getTwitterApi()
}

func (p *TimelineProcesser) Process(msgIn *model.Message) *model.Message {
	if p.Api == nil {
		return &model.Message{Body: "[timeline] api can not available"}
	}
	query := url.Values{}
	query.Add("count", "1")

	timeline, _ := p.Api.GetHomeTimeline(query)
	if len(timeline) < 1 {
		return &model.Message{Body: "[timeline] no result found"}
	}

	tweet := timeline[0]
	return &model.Message{Body: fmt.Sprintf("[timeline:%s] %s", tweet.User.Name, tweet.Text)}
}

func (p *UranaiProcesser) Process(msgIn *model.Message) *model.Message {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)
	result := []string{"大吉", "吉", "凶"}
	return &model.Message{Body: result[randNum]}
}

func (p *WarikanProcesser) Process(msgIn *model.Message) *model.Message {
	warikan := strings.Split(msgIn.Body, " ")
	amount, _ := strconv.Atoi(warikan[1])
	member, _ := strconv.Atoi(warikan[2])
	warikanResult := amount / member
	var answer string = "1人" + strconv.Itoa(warikanResult) + "円です。"
	return &model.Message{Body: answer}
}
