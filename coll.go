package coll

import (
	"errors"
	"fmt"
)

type Coll struct {
	Name   string
	Parent *Coll
}

func NewColl(name string) *Coll {
	return &Coll{
		Name: name,
	}
}

func NewChildColl(name string, parent *Coll) *Coll {
	return &Coll{
		Name:   name,
		Parent: parent,
	}
}

func (coll *Coll) Key(fields ...string) (string, error) {
	if len(fields) == 0 {
		return "", errors.New(fmt.Sprintf("Not enough parameters to access %s", coll.Name))
	}

	result := fmt.Sprintf("%s:%s", coll.Name, fields[len(fields)-1])

	if coll.Parent != nil {
		key, err := coll.Parent.Key(fields[0 : len(fields)-1]...)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("%s:%s", key, result)
	}

	return result, nil
}

func (coll *Coll) Select(fields ...string) *ReadWrite {
	key, err := coll.Key(fields...)
	return &ReadWrite{[]byte(key), err}
}
