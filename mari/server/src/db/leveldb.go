package db

import (
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

const LevelDbName = "_leveldb"

// データベースからの取得を行う
func Get(key string) ([]byte, error) {
	db, err := leveldb.OpenFile(LevelDbName, nil)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	defer db.Close()

	data, err := db.Get([]byte(key), nil)

	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}

		return nil, errors.Wrap(err, "")
	}

	return data, nil
}

// データベースへの更新を行う
func Put(key string, value []byte) error {
	db, err := leveldb.OpenFile(LevelDbName, nil)

	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Put([]byte(key), value, nil); err != nil {
		return err
	}

	return nil
}

// データベースから削除する
func Delete(key string) error {
	db, err := leveldb.OpenFile(LevelDbName, nil)

	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Delete([]byte(key), nil); err != nil {
		return err
	}

	return nil
}
