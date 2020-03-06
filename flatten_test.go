package flatten_test

import (
	"testing"

	"github.com/darylnwk/flatten"
	"github.com/stretchr/testify/assert"
)

type NestedFooBar struct {
	Foo string `json:"foo"`
	Bar string `json:"bar,omitempty"`
}

type FooBar struct {
	Foo          string       `json:"foo"`
	Bar          string       `json:"bar"`
	NestedFooBar NestedFooBar `json:"nested"`
}

func TestFlatten(t *testing.T) {
	m := map[string]interface{}{}

	flatten.Struct(&FooBar{
		Foo: "foo",
		Bar: "bar",
	}, m)

	assert.Equal(t, map[string]interface{}{
		"foo": "foo",
		"bar": "bar",
	}, m)
}

func TestFlatten_Nested(t *testing.T) {
	m := map[string]interface{}{}

	flatten.Struct(&FooBar{
		Foo: "foo",
		Bar: "bar",
		NestedFooBar: NestedFooBar{
			Foo: "foo",
			Bar: "bar",
		},
	}, m)

	assert.Equal(t, map[string]interface{}{
		"foo":        "foo",
		"bar":        "bar",
		"nested.foo": "foo",
		"nested.bar": "bar",
	}, m)
}
