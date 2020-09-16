package qs

import (
	"testing"
)

func TestStringify(t *testing.T) {
	testMap := map[string]interface{}{
		"a":1.123,
		"b":"string",
		"c": true,
		"d": []interface{}{1.123, "string", true },
		"e": map[string]interface{}{
			"a":1.123,
			"b":"string",
			"c": true,
			"d": []interface{}{1.123, "string", true },
		},
	}
	if str, err := Stringify(testMap); err != nil{
		t.Error(err)
	}else {
		t.Log(str) // a=1.123&b=string&c=true&d[]=1.123&d[]=string&d[]=true&e[b]=string&e[c]=true&e[d][]=1.123&e[d][]=string&e[d][]=true&e[a]=1.123
	}
}

func BenchmarkStringify(b *testing.B) {
	testMap := map[string]interface{}{
		"a":1.123,
		"b":"string",
		"c": true,
		"d": []interface{}{1.123, "string", true },
		"e": map[string]interface{}{
			"a":1.123,
			"b":"string",
			"c": true,
			"d": []interface{}{1.123, "string", true },
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Stringify(testMap)
	}
}
