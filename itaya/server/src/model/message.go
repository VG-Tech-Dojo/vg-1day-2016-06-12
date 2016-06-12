package model

import (
	"db"

	"encoding/json"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

type (
	Message struct {
		Id   int    `json:"id"`
		Body string `json:"body"`
		// CreatedAt string `json:"created_at"` // 1-1. メッセージの投稿時刻
		// Username  string `json:"user_name"`  // 1-2. ユーザ名
	}
	Messages []Message
)

// メッセージ一覧にあるメッセージの id をデータベースから全て取得する
func getMessageIds() ([]int, error) {
	data, err := db.Get("messages")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read database. key: messages")
	}

	var ids []int
	if data != nil {
		if err := json.Unmarshal(data, &ids); err != nil {
			return nil, errors.Wrapf(err, "failed to encode json. string: %s", data)
		}
	}

	sort.Sort(sort.IntSlice(ids))

	return ids, nil
}

// 新規メッセージの id をデータベースから取得する
func newMessageId() (int, error) {
	ids, err := getMessageIds()
	if err != nil {
		return -1, err
	}

	newId := 0
	if len(ids) != 0 {
		newId = ids[len(ids)-1] + 1
	}
	ids = append(ids, newId)

	data, err := json.Marshal(ids)
	if err != nil {
		return -1, errors.Wrapf(err, "failed to decode json: %v", ids)
	}

	if err := db.Put("messages", data); err != nil {
		return -1, errors.Wrapf(err, "failed to write to database: %v", data)
	}

	return newId, nil
}

// メッセージ id をデータベースから削除する
func deleteMessageId(id int) error {
	ids, err := getMessageIds()
	if err != nil {
		return err
	}

	result := []int{}
	for _, v := range ids {
		if v != id {
			result = append(result, v)
		}
	}

	data, err := json.Marshal(result)
	if err != nil {
		return errors.Wrapf(err, "failed to decode json. data: %v", result)
	}

	if err := db.Put("messages", data); err != nil {
		return errors.Wrapf(err, "failed to write database. key: messages, value: %v", data)
	}

	return nil
}

// メッセージをつくる
// 1-2. ユーザ名を受け取ってメッセージをつくる
func NewMessage(body string) (*Message, error) {
	id, err := newMessageId()
	if err != nil {
		return nil, err
	}

	return &Message{
		Id:   id,
		Body: body,
		// 1-1. CreatedAt に時刻をセットする
		// ヒント: https://golang.org/pkg/time/
		// 1-2. Username にユーザ名をセットする
	}, nil
}

// データベースにメッセージを保存する
func (m *Message) SaveMessage() error {
	key := fmt.Sprintf("message-%d", m.Id)

	json, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := db.Put(key, json); err != nil {
		return err
	}

	return nil
}

// データベースからメッセージを取得する
func (m *Message) LoadMessage(id int) error {
	key := fmt.Sprintf("message-%d", id)
	data, err := db.Get(key)
	if err != nil {
		return errors.Wrapf(err, "faild to read database. key: %s", key)
	}

	if data != nil {
		if err := json.Unmarshal([]byte(data), m); err != nil {
			return errors.Wrapf(err, "faild to encode json. data: %v", []byte(data))
		}
	}

	return nil
}

// 1-4. データベースからメッセージを削除する
func (m *Message) DeleteMessage() error {
	// key を作る
	// メッセージ id = 1 のメッセージの key は "message-1"

	// key を使ってデータベースからメッセージを削除する
	// ヒント db.Delete()

	if err := deleteMessageId(m.Id); err != nil {
		return err
	}

	return nil
}

// データベースからメッセージ一覧を取得する
func LoadMessages() (Messages, error) {
	ids, err := getMessageIds()
	if err != nil {
		return nil, err
	}

	var (
		m  Message
		ms Messages
	)
	for _, v := range ids {
		if err := m.LoadMessage(v); err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}
