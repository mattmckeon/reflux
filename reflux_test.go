package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Flush(b *strings.Builder) string {
	s := b.String()
	b.Reset()
	return s
}

func TestCounts(t *testing.T) {
	db := NewRefluxDb()
	b := &strings.Builder{}
	db.DoGet("a", b)
	assert.Equal(t, "NULL\n", Flush(b), "unexpected output value")
	db.DoSet("a", "foo")
	db.DoSet("b", "foo")
	db.DoCount("foo", b)
	assert.Equal(t, "2\n", Flush(b), "unexpected output value")
	db.DoCount("bar", b)
	assert.Equal(t, "0\n", Flush(b), "unexpected output value")
	db.DoDelete("a")
	db.DoCount("foo", b)
	assert.Equal(t, "1\n", Flush(b), "unexpected output value")
	db.DoSet("b", "baz")
	db.DoCount("foo", b)
	assert.Equal(t, "0\n", Flush(b), "unexpected output value")
	db.DoGet("b", b)
	assert.Equal(t, "baz\n", Flush(b), "unexpected output value")
	db.DoGet("B", b)
	assert.Equal(t, "NULL\n", Flush(b), "unexpected output value")
}

func TestBasicTransaction(t *testing.T) {
	db := NewRefluxDb()
	b := &strings.Builder{}

	db.DoBegin()
	db.DoSet("a", "foo")
	db.DoGet("a", b)
	assert.Equal(t, "foo\n", Flush(b), "unexpected output value")
	db.DoBegin()
	db.DoSet("a", "bar")
	db.DoGet("a", b)
	assert.Equal(t, "bar\n", Flush(b), "unexpected output value")
	db.DoRollback(b)
	db.DoGet("a", b)
	assert.Equal(t, "foo\n", Flush(b), "unexpected output value")
	db.DoRollback(b)
	db.DoGet("a", b)
	assert.Equal(t, "NULL\n", Flush(b), "unexpected output value")
}

func TestRollbackWithoutTransaction(t *testing.T) {
	db := NewRefluxDb()
	b := &strings.Builder{}

	db.DoSet("a", "foo")
	db.DoRollback(b)
	assert.Equal(t, "TRANSACTION NOT FOUND\n", Flush(b), "unexpected output value")
	db.DoGet("a", b)
	assert.Equal(t, "foo\n", Flush(b), "unexpected output value")
}
