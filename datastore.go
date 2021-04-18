package datastore

import "errors"

// Errors
var (
	ErrNotFound = errors.New("not found")
)

// ListFunc is used to list items in the datastore
type ListFunc func(key []byte, value []byte) error

// GetSetter is used to get or set values in a datastore
type GetSetter interface {
	Get(bucket []byte, key []byte) (value []byte, err error)
	Set(bucket []byte, key []byte, value []byte) error
	List(bucket []byte, pattern []byte, f ListFunc) error
}
