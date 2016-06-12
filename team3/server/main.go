package main

import (
	"route"

	"github.com/labstack/echo/engine/standard"
)

// 1day インターン チャットサーバ
func main() {
	router := route.Init()
	router.Run(standard.New(":1323"))
}
