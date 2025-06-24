package utils

import (
	"testing"
)

func TestSortStruct(t *testing.T) {
	data := &struct {
		Name   string `json:"name"`
		Filed1 string `json:"filed1"`
		Filed2 string `json:"filed2"`
	}{
		Name:   "test",
		Filed1: "filed1",
		Filed2: "filed2",
	}
	kv := SortStruct(data)
	t.Log(kv)
}
