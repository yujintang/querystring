### golang querystring

### Install
```shell script
go get -u github.com/yujintang/querystring
```
### Test
```go
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
	if str == "a=1.123&b=string&c=true&d[]=1.123&d[]=string&d[]=true&e[b]=string&e[c]=true&e[d][]=1.123&e[d][]=string&e[d][]=true&e[a]=1.123" {
		t.Error("Error Stringify")
	} else {
		t.Log("Success Stringify")
	}
}
```
### Benchmark
```text
goos: darwin
goarch: amd64
pkg: github.com/yujintang/querystring
BenchmarkStringify-8      313260              3404 ns/op
PASS
ok      github.com/yujintang/querystring        1.192s
```