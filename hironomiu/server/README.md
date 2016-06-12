# サーバーサイド

[![wercker status](https://app.wercker.com/status/dd5bc401a99482c7b7e8f41cec4d5c8b/s/master "wercker status")](https://app.wercker.com/project/bykey/dd5bc401a99482c7b7e8f41cec4d5c8b)

## 主に利用しているフレームワーク

* https://labstack.com/echo
* https://github.com/syndtr/goleveldb/leveldb


## はじめに

### GOPATHを設定しよう

```shell
$ export GOPATH=`pwd`/_gopath:`pwd`
```

### 依存ライブラリをインストールしよう

```shell
$ go get -v -d ./...
```

## サーバの起動と確認

### 起動

```shell
$ go run main.go
```

### 確認

* 注意: サーバが起動してるターミナルとは、別のターミナルで実行すること

```shell
$ curl -sS http://localhost:1323
Hello, World!
```

## API

* 注意: サーバが起動してるターミナルとは、別のターミナルで実行すること

### 投稿

```shell
$ curl -sS -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"body":"1day intern!"}' http://localhost:1323/messages
{"id":0,"body":"1day intern!"}
```

### 取得

```shell
$ curl -sS http://localhost:1323/messages/0
{"id":0,"body":"1day intern!"}
```

### 一覧取得

```shell
$ curl -sS http://localhost:1323/messages
[{"id":0,"body":"1day intern!"}]
```
