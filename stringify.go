package qs

import (
	"fmt"
	"net/url"
	"strconv"
)

func Stringify(hash map[string]interface{}) (string, error) {
	return buildNestedQuery(hash, "", 10)
}

func buildNestedQuery(value interface{}, prefix string, deep int) (string, error) {
	components := ""

	switch vv := value.(type) {
	case []interface{}:
		if deep < 0 {
			return "", nil
		}
		for i, v := range vv {
			component, err := buildNestedQuery(v, prefix+"[]", deep-1)

			if err != nil {
				return "", err
			}

			components += component

			if i < len(vv)-1 {
				components += "&"
			}
		}

	case map[string]interface{}:
		if deep < 0 {
			return "", nil
		}

		for k, v := range vv {
			childPrefix := ""

			if prefix != "" {
				childPrefix = prefix + "[" + url.QueryEscape(k) + "]"
			} else {
				childPrefix = url.QueryEscape(k)
			}

			component, err := buildNestedQuery(v, childPrefix, deep-1)

			if err != nil {
				return "", err
			}

			if len(component) > 0 && len(components) > 0 {
				components += "&"
			}
			components += component
		}

	case string:
		if prefix == "" {
			return "", fmt.Errorf("value must be a map[string]interface{}")
		}

		components += prefix + "=" + url.QueryEscape(vv)
	case bool:
		components += prefix + "=" + url.QueryEscape(strconv.FormatBool(vv))
	case float64:
		components += prefix + "=" + url.QueryEscape(strconv.FormatFloat(vv, 'G', -1, 64))
	default:
	}
	return components, nil
}
