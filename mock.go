package datastore

import (
	"bytes"
	"sync"
)

// Mock is used for quick dev and test scenarios
type Mock struct {
	Values    *sync.Map
	GetError  error
	SetError  error
	ListError error
}

// NewMock creates a new Mock instance
func NewMock() *Mock {
	return &Mock{
		Values:    &sync.Map{},
		GetError:  nil,
		SetError:  nil,
		ListError: nil,
	}
}

// Get retrieves the value stored with the specified key
func (m *Mock) Get(bucket []byte, key []byte) (value []byte, err error) {
	if m.GetError != nil {
		return nil, m.GetError
	}

	k := append(bucket, key...)
	v, ok := m.Values.Load(k)
	if !ok {
		return nil, ErrNotFound
	}
	return v.([]byte), nil
}

// Set sets the specified key to the specified value
func (m *Mock) Set(bucket []byte, key []byte, value []byte) error {
	if m.SetError != nil {
		return m.SetError
	}
	k := append(bucket, key...)
	m.Values.Store(k, value)
	return nil
}

// List calls the specified function once per matching pattern
// Only string patterns are supported
func (m *Mock) List(bucket []byte, pattern []byte, f ListFunc) error {
	if m.ListError != nil {
		return m.ListError
	}

	var err error

	m.Values.Range(func(key interface{}, value interface{}) bool {
		k := key.([]byte)
		v := value.([]byte)

		if !bytes.HasPrefix(k, bucket) {
			return true
		}

		k = k[len(bucket)+1:]
		if !bytes.Contains(k, pattern) {
			return true
		}

		err = f(k, v)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
