package route

import (
	"api"
	"bot"
	"model"

	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ルーティングの初期化
func Init() *echo.Echo {

	msgCh := prepareBot()

	// echo
	e := echo.New()

	e.Debug()

	// Set MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	m := new(sync.Mutex)
	e.Use(api.Mutex(m))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/messages", api.ReadMessages)                    // メッセージ一覧取得
	e.POST("/messages", api.ObservableCreateMessage(msgCh)) // メッセージ投稿

	e.GET("/messages/:id", api.ReadMessage)      // メッセージ取得
	e.PUT("/messages/:id", api.UpdateMessage)    // メッセージ更新
	e.DELETE("/messages/:id", api.DeleteMessage) // メッセージ削除

	return e
}

// bot の準備
func prepareBot() chan model.Message {

	p := bot.NewPoster(10)
	go p.Run()

	broadcaster := bot.NewBroadcaster()
	go broadcaster.Run()

	b1 := bot.NewGreetBot("simple_bot", `^hello`, p.Input)
	go b1.Run()
	broadcaster.EntryInput <- b1

	b2 := bot.NewGreetBot("another_bot", `^voyage`, p.Input)
	go b2.Run()
	broadcaster.EntryInput <- b2

	b3 := bot.NewTimelineBot(p.Input)
	go b3.Run()
	broadcaster.EntryInput <- b3

	return broadcaster.MsgInput
}
