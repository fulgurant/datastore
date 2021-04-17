package datastore

import (
	"strings"
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
func (m *Mock) Get(key interface{}) (value interface{}, err error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	value, ok := m.Values.Load(key)
	if !ok {
		return nil, ErrNotFound
	}
	return value, nil
}

// Set sets the specified key to the specified value
func (m *Mock) Set(key interface{}, value interface{}) error {
	if m.SetError != nil {
		return m.SetError
	}
	m.Values.Store(key, value)
	return nil
}

// List calls the specified function once per matching pattern
// Only string patterns are supported
func (m *Mock) List(pattern interface{}, f ListFunc) error {
	if m.ListError != nil {
		return m.ListError
	}

	p, ok := pattern.(string)
	if !ok {
		return ErrUnsupportedPattern
	}

	var err error

	m.Values.Range(func(key interface{}, value interface{}) bool {
		k, ok := key.(string)
		if !ok {
			return true
		}

		if !strings.Contains(k, p) {
			return true
		}

		err = f(key, value)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
