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
    t.Log("Success Stringify")
    t.Log(str) // a=1.123&b=string&c=true&d[]=1.123&d[]=string&d[]=true&e[b]=string&e[c]=true&e[d][]=1.123&e[d][]=string&e[d][]=true&e[a]=1.123
}
```
### Benchmark
```text
goos: darwin
goarch: amd64
pkg: github.com/yujintang/querystring
BenchmarkStringify-8      375606              3003 ns/op
PASS
ok      github.com/yujintang/querystring        1.234s
```