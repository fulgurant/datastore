package datastore

import "errors"

// Errors
var (
	ErrNotFound           = errors.New("not found")
	ErrUnsupportedPattern = errors.New("unsupported pattern")
)

type ListFunc func(key interface{}, value interface{}) error

// GetSetter is used to get or set values in a datastore
type GetSetter interface {
	Get(key interface{}) (value interface{}, err error)
	Set(key interface{}, value interface{}) error
	List(pattern interface{}, f ListFunc) error
}
