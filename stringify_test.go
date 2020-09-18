package qs

import (
	"testing"
)

func TestStringify(t *testing.T) {
	testMap := map[string]interface{}{
		"a": 1.123,
		"b": "string",
		"c": true,
		"d": []interface{}{1.123, "string", true},
		"e": map[string]interface{}{
			"a": 1.123,
			"b": "string",
			"c": true,
			"d": []interface{}{1.123, "string", true},
		},
	}
	str := Stringify(testMap)
	t.Log(testMap)
	t.Log(str)
}

func BenchmarkStringify(b *testing.B) {
	testMap := map[string]interface{}{
		"a": 1.123,
		"b": "string",
		"c": true,
		"d": []interface{}{1.123, "string", true},
		"e": map[string]interface{}{
			"a": 1.123,
			"b": "string",
			"c": true,
			"d": []interface{}{1.123, "string", true},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Stringify(testMap)
	}
}
