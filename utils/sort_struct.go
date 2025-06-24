package utils

import (
	"reflect"
	"sort"
	"unicode"
)

type KeyValue struct {
	Key   string
	Value any
}

func SortStruct(data any) []KeyValue {
	v := reflect.ValueOf(data).Elem()
	t := v.Type()
	var kv []KeyValue
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := ToSnakeCase(t.Field(i).Name)
		if field.IsZero() {
			continue
		}
		kv = append(kv, KeyValue{Key: key, Value: field.Interface()})
	}
	sort.SliceStable(kv, func(i, j int) bool {
		return kv[i].Key < kv[j].Key
	})
	return kv
}

func ToSnakeCase(s string) string {
	var result []rune
	for i, char := range s {
		if unicode.IsUpper(char) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(char))
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
