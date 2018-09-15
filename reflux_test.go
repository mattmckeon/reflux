package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	db := NewReflux()
	assert.False(t, db.Has("foo"), `Has("foo") should be false`)

	db = &Reflux{
		Data:  map[string]string{"foo": "bar"},
		Counts: map[string]int{"bar": 1},
	}
	assert.True(t, db.Has("foo"), `Has("foo") should be true`)
}

func TestGet(t *testing.T) {
	db := NewReflux()
	assert.Equal(t, "", db.Get("foo"), `Get("foo") should be ""`)

	db = &Reflux{
		Data:  map[string]string{"foo": "bar"},
		Counts: map[string]int{"bar": 1},
	}
	assert.Equal(t, "bar", db.Get("foo"), `Get("foo") should be "bar"`)

	db = &Reflux{
		Data:  map[string]string{"foo": "baz"},
		Counts: map[string]int{"baz": 1},
	}
	assert.Equal(t, "baz", db.Get("foo"), `Get("foo") should be "baz"`)
}

func TestSet(t *testing.T) {
	db := NewReflux()

	db.Set("foo", "bar")
	assert.Equal(t, "bar", db.Data["foo"], `"bar" should be value for key "foo"`)
	assert.Equal(t, 1, db.Counts["bar"], `1 should be count for value "bar"`)

	db.Set("flo", "bar")
	assert.Equal(t, "bar", db.Data["flo"], `"bar" should be value for key "flo"`)
	assert.Equal(t, 2, db.Counts["bar"], `2 should be count for value "bar"`)

	db.Set("foo", "baz")
	assert.Equal(t, "baz", db.Data["foo"], `"baz" should be value for key "foo"`)
	assert.Equal(t, 1, db.Counts["baz"], `1 should be count for value "baz"`)
	assert.Equal(t, 1, db.Counts["bar"], `1 should be count for value "bar"`)

	db.Set("flo", "bat")
	assert.Equal(t, 0, db.Counts["bar"], `0 should be count for value "bar"`)
}

func TestDelete(t *testing.T) {
	db := &Reflux{
		Data:  map[string]string{"foo": "bar", "flo": "bar", "fla": "baz"},
		Counts: map[string]int{"bar": 2, "baz": 1},
	}

	db.Delete("foo")
	assert.Equal(t, "", db.Data["foo"], `"" should be value for key "foo"`)
	assert.Equal(t, 1, db.Counts["bar"], `1 should be count for value "bar"`)
	assert.Equal(t, 1, db.Counts["baz"], `1 should be count for value "baz"`)

	db.Delete("flo")
	assert.Equal(t, "", db.Data["flo"], `"" should be value for key "foo"`)
	assert.Equal(t, 0, db.Counts["bar"], `0 should be count for value "bar"`)
}
