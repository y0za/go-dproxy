package dproxy

import "testing"

func TestUnmarshal(t *testing.T) {
	type item struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	v := parseJSON(`{
		"item": {
			"name": "Bob",
			"age": 20
		}
	}`)

	itemProxy := New(v).M("item")

	var i item

	err := itemProxy.Unmarshal(&i)
	if err != nil {
		t.Error(err)
	}

	if i.Name != "Bob" {
		t.Error("unexpected name:", i.Name)
	}
	if i.Age != 20 {
		t.Error("unexpected age:", i.Age)
	}
}
