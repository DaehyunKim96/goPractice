package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errKeyExist = errors.New("Key Exist")
)

func (d Dictionary) Search(word string) (string, error) {
	val, exists := d[word]
	if exists {
		return val, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, val string) error {
	if _, err := d.Search(key); err != nil {
		d[key] = val
		return nil
	} else {
		return errKeyExist
	}
}

func (d Dictionary) Update(key, val string) error {
	if _, err := d.Search(key); err != nil {
		return errNotFound
	} else {
		d[key] = val
		return nil
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)
	case errNotFound:
		return errNotFound
	}
	return nil
}
