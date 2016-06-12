.PHONY: more \
  deps clean gopath run \
  hello messages messages/*

GOPATH:=$(CURDIR)/_gopath

DB:=$(CURDIR)/_leveldb

PORT:=1323
ADDR:=http://localhost:$(PORT)

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
more:
	@awk 'BEGIN {FS = ":.*?## "} /^[\$$\(\)\/a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean:                      ## データベースを初期化する
	-rm -rf $(DB)

$(GOPATH):
	@mkdir -p $@

deps: $(GOPATH)
	GOPATH=$(GOPATH):$(CURDIR) go get -d -v ./...

run: deps $(GOPATH)         ## サーバを起動する
	GOPATH=$(GOPATH):$(CURDIR) go run main.go

hello:                      ## GET   /
	curl -sS $(ADDR)

messages:                   ## GET   /messages
	curl -sS $(ADDR)/$@

messages/post: USERNAME:=$(shell whoami)
messages/post: BODY=1day intern!
messages/post:              ## POST  /messages     BODY=メッセージ内容
	curl -sS -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"user_name": "$(USERNAME)", "body":"$(BODY)"}' $(ADDR)/messages

messages/get: ID=0
messages/get:               ## GET   /messages/:id ID=メッセージID
	curl -sS $(ADDR)/messages/$(ID)

messages/put: ID=0
messages/put: USERNAME:=$(shell whoami)
messages/put: BODY=update!
messages/put:               ## PUT   /messages/:id ID=メッセージID BODY=メッセージ内容
	curl -sS -H "Accept: application/json" -H "Content-type: application/json" -X PUT -d '{"user_name": "$(USERNAME)", "body":"$(BODY)"}' http://localhost:1323/messages/$(ID)

messages/delete: ID=0
messages/delete:            ## DELTE /messages/:id ID=メッセージID
	curl -sS -X DELETE http://localhost:1323/messages/$(ID)
