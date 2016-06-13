package bot

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"model"
	"net/url"
	// "math/rand"
	// "time"
	"strings"
	"strconv"

	"net/http"
  "io/ioutil"
	"github.com/antonholmquist/jason"
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

	UranaiProcesser struct {
	}

	WarikanProcesser struct {
		Persons int
		Money int
	}

	// フォト蔵
	PhotozouProcesser struct {
		SearchWord string
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
	txt := "http://news.biglobe.ne.jp/img/blnews/animal160128_02.jpg"
	return &model.Message{Img: txt}
	// answers := []string{"大吉","吉","凶"}
	// rand.Seed(time.Now().UnixNano())
	// txt := answers[rand.Intn(len(answers))]
	// return &model.Message{Body: txt}
}

func (p *WarikanProcesser) Process(msgIn *model.Message) *model.Message {
	r := strings.Split(msgIn.Body, " ")
	money,_ := strconv.Atoi(r[1])
	persons, _ := strconv.Atoi(r[2])

	money_per_person := money / persons

	txt := "一人当たり" + strconv.Itoa(money_per_person) + "円です．"
	return &model.Message{Body: txt}
}

// フォト蔵
func (p *PhotozouProcesser) Process(msgIn *model.Message) *model.Message {
	url := "https://api.photozou.jp/rest/search_public.json?keyword=aaa"

  resp, _ := http.Get(url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
	return &model.Message{Img: url}
}
