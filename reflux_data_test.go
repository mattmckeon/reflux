package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	data := NewRefluxData()
	assert.False(t, data.Has("foo"), `Has("foo") should be false`)

	data = &RefluxData{
		Data:   map[string]string{"foo": "bar"},
		Counts: map[string]int{"bar": 1},
	}
	assert.True(t, data.Has("foo"), `Has("foo") should be true`)
}

func TestGet(t *testing.T) {
	data := NewRefluxData()
	assert.Equal(t, "", data.Get("foo"), `Get("foo") should be ""`)

	data = &RefluxData{
		Data:   map[string]string{"foo": "bar"},
		Counts: map[string]int{"bar": 1},
	}
	assert.Equal(t, "bar", data.Get("foo"), `Get("foo") should be "bar"`)

	data = &RefluxData{
		Data:   map[string]string{"foo": "baz"},
		Counts: map[string]int{"baz": 1},
	}
	assert.Equal(t, "baz", data.Get("foo"), `Get("foo") should be "baz"`)
}

func TestSet(t *testing.T) {
	data := NewRefluxData()

	data.Set("foo", "bar")
	assert.Equal(t, "bar", data.Data["foo"], `"bar" should be value for key "foo"`)
	assert.Equal(t, 1, data.Counts["bar"], `1 should be count for value "bar"`)

	data.Set("flo", "bar")
	assert.Equal(t, "bar", data.Data["flo"], `"bar" should be value for key "flo"`)
	assert.Equal(t, 2, data.Counts["bar"], `2 should be count for value "bar"`)

	data.Set("foo", "baz")
	assert.Equal(t, "baz", data.Data["foo"], `"baz" should be value for key "foo"`)
	assert.Equal(t, 1, data.Counts["baz"], `1 should be count for value "baz"`)
	assert.Equal(t, 1, data.Counts["bar"], `1 should be count for value "bar"`)

	data.Set("flo", "bat")
	assert.Equal(t, 0, data.Counts["bar"], `0 should be count for value "bar"`)
}

func TestDelete(t *testing.T) {
	data := &RefluxData{
		Data:   map[string]string{"foo": "bar", "flo": "bar", "fla": "baz"},
		Counts: map[string]int{"bar": 2, "baz": 1},
	}

	data.Delete("foo")
	assert.Equal(t, "", data.Data["foo"], `"" should be value for key "foo"`)
	assert.Equal(t, 1, data.Counts["bar"], `1 should be count for value "bar"`)
	assert.Equal(t, 1, data.Counts["baz"], `1 should be count for value "baz"`)

	data.Delete("flo")
	assert.Equal(t, "", data.Data["flo"], `"" should be value for key "foo"`)
	assert.Equal(t, 0, data.Counts["bar"], `0 should be count for value "bar"`)
}

func TestDeleteCounts(t *testing.T) {
	data := NewRefluxData()
	assert.Equal(t, "", data.Get("a"), `"" should be value for key "a"`)
	data.Set("a", "foo")
	data.Set("b", "foo")
	assert.Equal(t, 2, data.Count("foo"), `2 should be value for value "foo"`)
	assert.Equal(t, 0, data.Count("bar"), `0 should be value for value "bar"`)
	data.Delete("a")
	assert.Equal(t, 1, data.Count("foo"), `1 should be value for value "foo"`)
	data.Set("b", "baz")
	assert.Equal(t, 0, data.Count("foo"), `0 should be value for value "foo"`)
	assert.Equal(t, "baz", data.Get("b"), `"baz" should be value for key "b"`)
	assert.Equal(t, "", data.Get("B"), `"" should be value for key "B"`)
}
