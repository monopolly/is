package is

import (
	"fmt"
	"strings"
)

// is string
func String(key string, value string) string {
	switch key == "" {
	case true:
		return value
	default:
		return key
	}
}

// not nil
func NotNil(k, v interface{}) (r interface{}) {
	if k != nil {
		return v
	}
	return
}

func Nil(k interface{}, yes, no string) (r string) {
	if k == nil {
		return yes
	}
	return no
}
func Nils(k interface{}, yes string) (r string) {
	if k == nil {
		return yes
	}
	return
}

func NotNilString(k, v interface{}) (r string) {

	switch v.(type) {
	case string:
		if v.(string) != "" {
			return fmt.Sprint(v)
		}
	case []byte:
		if v.([]byte) != nil {
			return fmt.Sprint(v)
		}
	case int8:
		if v.(int8) > 0 {
			return fmt.Sprint(v)
		}
	case int16:
		if v.(int16) > 0 {
			return fmt.Sprint(v)
		}
	case int32:
		if v.(int32) > 0 {
			return fmt.Sprint(v)
		}
	case int:
		if v.(int) > 0 {
			return fmt.Sprint(v)
		}
	case int64:
		if v.(int64) > 0 {
			return fmt.Sprint(v)
		}
	case uint8:
		if v.(uint8) > 0 {
			return fmt.Sprint(v)
		}
	case uint16:
		if v.(uint16) > 0 {
			return fmt.Sprint(v)
		}
	case uint32:
		if v.(uint32) > 0 {
			return fmt.Sprint(v)
		}
	case uint64:
		if v.(uint64) > 0 {
			return fmt.Sprint(v)
		}
	case uint:
		if v.(uint) > 0 {
			return fmt.Sprint(v)
		}
	case bool:
		if v.(bool) {
			return fmt.Sprint(v)
		}
	default:
		if v != nil {
			return fmt.Sprint(v)
		}
	}

	return
}

func StringExist(key string) string {
	switch key {
	case "":
		return ""
	default:
		return key
	}
}

func StringIfThen(ifs, then string) string {
	switch ifs {
	case "":
		return ""
	default:
		return then
	}
}

func Int64IfThen(ifs int64, then string) string {
	switch ifs {
	case 0:
		return ""
	default:
		return then
	}
}

func Int(v int, defaults int) int {
	switch v {
	case 0:
		return defaults
	default:
		return v
	}
}

func Uint32(k, v uint32) uint32 {
	switch k == 0 {
	case true:
		return v
	default:
		return k
	}
}

func Uint64(k, v uint64) uint64 {
	switch k == 0 {
	case true:
		return v
	default:
		return k
	}
}

func BoolD(k, v bool) bool {
	switch k {
	case true:
		return k
	default:
		return v
	}
}

func Bool(value bool, trues, falses string) string {
	switch value {
	case true:
		return trues
	default:
		return falses
	}
}

func InArray(a []string, value string, trues, falses string) string {
	for _, x := range a {
		if strings.ToLower(x) == strings.ToLower(value) {
			return trues
		}
	}
	return falses
}

func UintString(value uint, r map[uint]string) string {
	return r[value]
}
