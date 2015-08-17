package coll

import (
	"github.com/azer/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

var log = logger.New("level-collection")

var (
	Client *leveldb.DB
)

func Open(path string) error {
	log.Info("Opening %s", path)

	conn, err := leveldb.OpenFile(path, nil)
	Client = conn

	if err != nil {
		return err
	}

	return nil
}

func Set(key, value []byte) error {
	log.Info("Set", logger.Attrs{
		"key": key,
	})

	return Client.Put(key, value, nil)
}

func Get(key []byte) ([]byte, error) {
	value, err := Client.Get(key, nil)

	if err != nil {
		log.Info("Can not get", logger.Attrs{
			"key":   key,
			"error": err,
		})

		return nil, err
	}

	return value, nil
}

func Delete(key []byte) error {
	log.Info("Delete", logger.Attrs{
		"key": key,
	})

	return Client.Delete(key, nil)
}
