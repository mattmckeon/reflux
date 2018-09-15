package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestData() *RefluxData {
	return &RefluxData{
		Data:   map[string]string{"foo": "bar"},
		Counts: map[string]int{"bar": 1},
	}
}

func TestGetOp(t *testing.T) {
	db := NewTestData()
	b := strings.Builder{}
	op := &Get{
		Key:    "foo",
		Output: &b,
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Noop{}, inverse, "Unexpected inverse operation")
	assert.Equal(t, "bar\n", b.String(), "unexpected output value")
}

func TestGetOpMissingKey(t *testing.T) {
	db := NewTestData()
	b := strings.Builder{}
	op := &Get{
		Key:    "flo",
		Output: &b,
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Noop{}, inverse, "Unexpected inverse operation type")
	assert.Equal(t, "NULL\n", b.String(), "unexpected output value")
}

func TestSetOp(t *testing.T) {
	db := NewTestData()
	op := &Set{
		Key:   "foo",
		Value: "baz",
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Set{}, inverse, "Unexpected inverse operation type")
	assert.Equal(t, "baz", db.Get("foo"), "unexpected value in db")
}

func TestSetOpInverse(t *testing.T) {
	db := NewTestData()
	op := &Set{
		Key:   "foo",
		Value: "baz",
	}
	op.Apply(db).Apply(db)
	assert.Equal(t, "bar", db.Get("foo"), "Unexpected value in db after inverse set")
}

func TestSetOpInverseDelete(t *testing.T) {
	db := NewTestData()
	op := &Set{
		Key:   "flo",
		Value: "bar",
	}
	op.Apply(db).Apply(db)
	assert.False(t, db.Has("flo"))
	assert.Equal(t, 1, db.Count("bar"))
}

func TestDeleteOp(t *testing.T) {
	db := NewTestData()
	op := &Delete{
		Key: "foo",
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Set{}, inverse, "Unexpected inverse operation type")
	assert.False(t, db.Has("foo"), "Value shouldn't be present in db after delete")
}

func TestDeleteOpInverse(t *testing.T) {
	db := NewTestData()
	op := &Delete{
		Key: "foo",
	}
	op.Apply(db).Apply(db)
	assert.Equal(t, "bar", db.Get("foo"), "Value should be present in db after inverse delete")
}

func TestCountOp(t *testing.T) {
	db := NewTestData()
	b := strings.Builder{}
	op := &Count{
		Value:  "bar",
		Output: &b,
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Noop{}, inverse, "Unexpected inverse operation")
	assert.Equal(t, "1\n", b.String(), "unexpected output value")
}

func TestCountOpMissingValue(t *testing.T) {
	db := NewTestData()
	b := strings.Builder{}
	op := &Count{
		Value:  "baz",
		Output: &b,
	}
	inverse := op.Apply(db)
	assert.IsType(t, &Noop{}, inverse, "Unexpected inverse operation type")
	assert.Equal(t, "0\n", b.String(), "unexpected output value")
}
