package fib

import (
	"encoding/json"
	"testing"
)

var str1 = `{"key":111}`
var str2 = `{"key":{"key":222}}`
var str3 = `{"key":{"key":{"key":333}}}`

type (
	Str1 struct {
		Key int `json:"key"`
	}
	Str2 struct {
		Key Str1
	}
	Str3 struct {
		Key Str2
	}
)

// go test json_test.go -bench=. -benchmem -benchtime=30s
func BenchmarkJson1(b *testing.B) {
	var s1 Str1
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(str1), &s1)
	}
}

func BenchmarkJson2(b *testing.B) {
	var s2 Str1
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(str2), &s2)
	}
}

func BenchmarkJson3(b *testing.B) {
	var s3 Str1
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(str3), &s3)
	}
}
