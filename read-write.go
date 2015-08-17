package coll

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type ReadWrite struct {
	Key   []byte
	Error error
}

func (rw *ReadWrite) Attr(name string) *ReadWrite {
	rw.Key = []byte(fmt.Sprintf("%s:%s", rw.Key, name))
	return rw
}

func (rw *ReadWrite) Read() (string, error) {
	if rw.Error != nil {
		return "", rw.Error
	}

	value, err := Get(rw.Key)

	if err != nil {
		return "", err
	}

	return string(value), nil
}

func (rw *ReadWrite) ReadByte() ([]byte, error) {
	if rw.Error != nil {
		return nil, rw.Error
	}

	return Get(rw.Key)
}

func (rw *ReadWrite) Write(value string) error {
	if rw.Error != nil {
		return rw.Error
	}

	return Set(rw.Key, []byte(value))
}

func (rw *ReadWrite) WriteByte(value []byte) error {
	if rw.Error != nil {
		return rw.Error
	}

	return Set(rw.Key, value)
}

func (rw *ReadWrite) Delete() error {
	if rw.Error != nil {
		return rw.Error
	}

	return Delete(rw.Key)
}

func (rw *ReadWrite) Iter() iterator.Iterator {
	return Client.NewIterator(util.BytesPrefix(rw.Key), nil)
}
