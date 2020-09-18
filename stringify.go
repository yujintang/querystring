package qs

import (
	"fmt"
	"net/url"
	"strconv"
)

func Stringify(hash map[string]interface{}) string {
	return buildNestedQuery(hash, "", 10)
}

func buildNestedQuery(value interface{}, prefix string, deep int) string {
	components := ""

	switch vv := value.(type) {
	case []interface{}:
		if deep < 0 {
			return ""
		}
		for i, v := range vv {
			component := buildNestedQuery(v, prefix+"[]", deep-1)

			components += component

			if i < len(vv)-1 {
				components += "&"
			}
		}

	case map[string]interface{}:
		if deep < 0 {
			return ""
		}

		for k, v := range vv {
			childPrefix := ""

			if prefix != "" {
				childPrefix = prefix + "[" + url.QueryEscape(k) + "]"
			} else {
				childPrefix = url.QueryEscape(k)
			}

			component := buildNestedQuery(v, childPrefix, deep-1)

			if len(component) > 0 && len(components) > 0 {
				components += "&"
			}
			components += component
		}

	case string:
		if prefix == "" {
			return ""
		}
		components += prefix + "=" + vv
	case bool:
		components += prefix + "=" + strconv.FormatBool(vv)
	case int:
		components += prefix + "=" + strconv.Itoa(vv)
	case int8:
		components += prefix + "=" + strconv.FormatInt(int64(vv), 10)
	case int16:
		components += prefix + "=" + strconv.FormatInt(int64(vv), 10)
	case int32:
		components += prefix + "=" + strconv.FormatInt(int64(vv), 10)
	case int64:
		components += prefix + "=" + strconv.FormatInt(vv, 10)
	case float32:
		components += prefix + "=" + strconv.FormatFloat(float64(vv), 'f', -1, 32)
	case float64:
		components += prefix + "=" + strconv.FormatFloat(vv, 'f', -1, 64)
	default:
		components += prefix + "=" + fmt.Sprintf("%v",vv)
	}
	return components
}
