package main

import (
	"fmt"
	"io"
)

// Operation is the base interface for database operations.
type Operation interface {
	// Apply executes this operation on given datastore, returning its inverse.
	Apply(data *RefluxData) Operation
}

// Noop is a null operation, used for non-reversible ops.
type Noop struct{}

// Apply does nothing for the Noop operation.
func (o *Noop) Apply(data *RefluxData) Operation {
	return o
}

// GlobalNoop is a singleton for Noop.
var GlobalNoop = &Noop{}

// Get represents the GET operation.
type Get struct {
	Key    string
	Output io.Writer
}

// Apply prints out the value for the given key, or NULL if it doesn't exist.
// Return a Noop for its inverse.
func (o *Get) Apply(data *RefluxData) Operation {
	if data.Has(o.Key) {
		fmt.Fprintln(o.Output, data.Get(o.Key))
	} else {
		fmt.Fprintln(o.Output, "NULL")
	}
	return GlobalNoop
}

// Set represents the SET operation.
type Set struct {
	Key   string
	Value string
}

// Apply sets the key to the given value, returning an inverse operation of
// either DELETE or SET depending on the value that exists in the datastore.
func (o *Set) Apply(data *RefluxData) Operation {
	var inverse Operation
	if data.Has(o.Key) {
		inverse = &Set{Key: o.Key, Value: data.Get(o.Key)}
	} else {
		inverse = &Delete{Key: o.Key}
	}
	data.Set(o.Key, o.Value)
	return inverse
}

// Delete represents the DELETE operation.
type Delete struct {
	Key string
}

// Apply deletes the given key from the datastore, returning an inverse
// operation of SET that restores the existing value.
func (o *Delete) Apply(data *RefluxData) Operation {
	inverse := &Set{Key: o.Key, Value: data.Get(o.Key)}
	data.Delete(o.Key)
	return inverse
}

// Count represents the COUNT operation.
type Count struct {
	Value  string
	Output io.Writer
}

// Apply prints out the number of times the given value occurs in the datastore,
// or 0 if the value doesn't exist. Returns a Noop for its inverse operation.
func (o *Count) Apply(data *RefluxData) Operation {
	fmt.Fprintln(o.Output, data.Count(o.Value))
	return GlobalNoop
}
