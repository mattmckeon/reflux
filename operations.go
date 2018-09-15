package main

import (
	"fmt"
	"io"
)

type Operation interface {
	Apply(db *Reflux) Operation
}

type Noop struct{}

func (o *Noop) Apply(db *Reflux) Operation {
	return o
}

var GlobalNoop = &Noop{}

type Get struct {
	Key    string
	Output io.Writer
}

func (o *Get) Apply(db *Reflux) Operation {
	if db.Has(o.Key) {
		fmt.Fprintln(o.Output, db.Get(o.Key))
	} else {
		fmt.Fprintln(o.Output, "NULL")
	}
	return GlobalNoop
}

type Set struct {
	Key string;
	Value string;
}

func (o *Set) Apply(db *Reflux) Operation {
	inverse := &Set{Key: o.Key, Value: db.Get(o.Key)}
	db.Set(o.Key, o.Value)
	return inverse
}

type Delete struct {
	Key string;
}

func (o *Delete) Apply(db *Reflux) Operation {
	inverse := &Set{Key: o.Key, Value: db.Get(o.Key)}
	db.Delete(o.Key);
	return inverse;
}

type Count struct {
	Value    string
	Output io.Writer
}

func (o *Count) Apply(db *Reflux) Operation {
	fmt.Fprintln(o.Output, db.Count(o.Value))
	return GlobalNoop
}
