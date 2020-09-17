package qs

import (
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

		components += prefix + "=" + url.QueryEscape(vv)
	case bool:
		components += prefix + "=" + url.QueryEscape(strconv.FormatBool(vv))
	case float64:
		components += prefix + "=" + url.QueryEscape(strconv.FormatFloat(vv, 'G', -1, 64))
	default:
	}
	return components
}
