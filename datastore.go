package datastore

import "errors"

// Errors
var (
	ErrNotFound = errors.New("not found")
)

// ListFunc is used to list items in the datastore
type ListFunc func(key []byte, value []byte) error

// Handler is used to get or set values in a datastore
type Handler interface {
	Get(bucket []byte, key []byte) (value []byte, err error)
	Set(bucket []byte, key []byte, value []byte) error
	List(bucket []byte, pattern []byte, f ListFunc) error
}
